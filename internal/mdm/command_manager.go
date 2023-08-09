package mdm

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"time"

	"golang.org/x/exp/slog"
)

// Interface to DB operations
type DBOperations interface {
	MDMInsertEnrolledDevice(device *MDMWindowsEnrolledDevice) error
	MDMGetEnrolledDevice(mdmDeviceID string) (*MDMWindowsEnrolledDevice, error)
	MDMGetEnrolledDeviceByHWID(mdmDeviceHWID string) (*MDMWindowsEnrolledDevice, error)
	MDMGetEnrolledDevices() ([]*MDMWindowsEnrolledDevice, error)
	MDMDeleteEnrolledDevice(deviceID string) error
	MDMDeleteEnrolledDeviceByHWID(deviceHWID string) error
	MDMIsValidDeviceID(deviceID string) bool
	InsertMdmCertificate(cert *MDMCertificate) error
	GetMDMCertificate(id int) (*MDMCertificate, error)
	DeleteMDMCertificate(id int) error
	GetIdentityCert() (*x509.Certificate, *rsa.PrivateKey, error)
	InsertDeviceSetting(setting *KSDeviceSetting) error
	GetDeviceSettings(deviceID string) ([]*KSDeviceSetting, error)
	GetCustomDeviceSetting(deviceID string, settingURI string) (*KSDeviceSetting, error)
	DeleteDeviceSetting(deviceID string) error
	DeleteCustomDeviceSetting(deviceID string, settingURI string) error
	UpdateDeviceSettingWithType(setting *KSDeviceSetting) error
	UpdateDeviceSetting(deviceID string, settingURI string, settingValue string) error
	InsertPendingDeviceOperation(deviceID string, cmdVerb string, settingURI string, settingValue string) error
	GetPendingDeviceOperations(deviceID string) ([]*PendingDeviceOperation, error)
	DeletePendingDeviceOperations(deviceID string) error
	DeletePendingDeviceOperation(deviceID string, cmdVerb string, settingURI string) error
	QueuePendingDeviceOperation(deviceID string, cmdVerb string, settingURI string, settingValue string) error
	GetAllPendingOperations(deviceID string) ([]*PendingDeviceOperation, error)
	QueueProtoCmdOperation(deviceID string, cmdVerb string, settingURI string, settingValue string) error
	GetCustomPendingOperations(deviceID string, targetOp string) (*PendingDeviceOperation, error)
}

// CommandManager is the main struct that manages the MDM commands
type CommandManager struct {
	ManagementUrl     string
	Logger            *slog.Logger
	CmdHandlers       *CmdMessageHandlers
	InitialOperations []PendingDeviceOperation
	DBOperations      DBOperations
}

func newCommandManager(baseDomain string, logger *slog.Logger, db DBOperations) (*CommandManager, error) {

	urlManagementEndpoint, err := ResolveWindowsMDMManagement(baseDomain)
	if err != nil {
		return nil, err
	}

	// Creating the handlers containers
	reqCmdHandlers := make(CmdMessageHandlers)

	initialDeviceOperations := getDefaultOperationRequests()
	return &CommandManager{
		ManagementUrl:     urlManagementEndpoint,
		Logger:            logger,
		CmdHandlers:       &reqCmdHandlers,
		InitialOperations: initialDeviceOperations,
		DBOperations:      db,
	}, nil
}

func GetCommandManager(baseDomain string, logger *slog.Logger, db DBOperations) (*CommandManager, error) {

	// Initialize the command manager
	cmdManager, err := newCommandManager(baseDomain, logger, db)
	if err != nil {
		return nil, err
	}

	// Set the default handlers
	cmdManager.setProtocolCommandHandlers()

	// Clear wait flags
	err = cmdManager.ClearTunnelRelatedSettings()
	if err != nil {
		return nil, err
	}

	// Queue initial operations

	// Getting all the devices
	devices, err := cmdManager.DBOperations.MDMGetEnrolledDevices()
	if err != nil {
		return nil, err
	}

	// Queueing the initial operations for all the devices
	for _, device := range devices {
		for _, deviceOp := range cmdManager.InitialOperations {
			err := db.QueuePendingDeviceOperation(device.ID, deviceOp.CmdVerb, deviceOp.SettingURI, deviceOp.SettingValue)
			if err != nil {
				return nil, fmt.Errorf("storing pending operations %v", err)
			}
		}
	}

	return cmdManager, nil
}

// processIncomingProtocolCommands will process the incoming command
func (c *CommandManager) processIncomingProtocolCommands(messageID string, deviceID string, cmd ProtoCmdOperation) (*SyncMLCmd, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	// Not processing Status commands
	if cmd.Verb == CmdStatus {
		return nil, nil
	}

	//Handle lookup and execution

	// Handling Alert special case where target URI is not provided
	if cmd.Verb == CmdAlert && cmd.Cmd.Data != nil {
		targetURI := "./" + CmdAlert + "/" + *cmd.Cmd.Data

		// Check if we have a handler for the target URI and execute it
		if handler, exists := (*c.CmdHandlers)[targetURI]; exists {
			err := handler.ReceiveCmd(deviceID, string(cmd.Verb), &cmd.Cmd)
			if err != nil {
				c.Logger.Error(fmt.Sprintf("command handler for URI %s - err: %v", targetURI, err))
			}
		}

		// We return a status 200 for all the operations
		return NewSyncMLCmdStatus(messageID, cmd.Cmd.CmdID, string(cmd.Verb), CmdStatusCode200), nil
	}

	// Lets do a first pass to check the commands that do not have items
	if cmd.Cmd.Items != nil && len(*(cmd.Cmd).Items) > 0 {

		// Then check the one that do have items
		for _, item := range *(cmd.Cmd).Items {
			settingURI := item.Source
			if (settingURI == nil) || (len(*settingURI) == 0) {
				continue
			}

			if handler, exists := (*c.CmdHandlers)[*settingURI]; exists {
				err := handler.ReceiveCmd(deviceID, string(cmd.Verb), &cmd.Cmd)
				if err != nil {
					c.Logger.Error(fmt.Sprintf("command handler for URI %s - err: %v", *settingURI, err))
				}
			}
		}
	}

	// We return a status 200 for all the operations
	return NewSyncMLCmdStatus(messageID, cmd.Cmd.CmdID, string(cmd.Verb), CmdStatusCode200), nil
}

func (c *CommandManager) WaitForTunnelCmdIfNeeded(deviceID string, pendingOps []*PendingDeviceOperation) ([]*PendingDeviceOperation, error) {

	var newPendingOps []*PendingDeviceOperation

	// Check if tunnel input wait is required
	waitRequired, err := c.IsTunnelWaitRequired(deviceID)
	if err != nil {
		return nil, fmt.Errorf("message processing error %v", err)
	}

	// Check if tunnel cmd was already queued
	cmdAlreadyQueued, err := c.IsTunnelCmdQueued(pendingOps)
	if err != nil {
		return nil, fmt.Errorf("message processing error %v", err)
	}

	// Wait only if the wait flag is set and the tunnel command was not already queued
	if waitRequired && !cmdAlreadyQueued {
		tunnelOp, err := c.WaitForPendingTunnelOp(deviceID)
		if err != nil {
			return nil, fmt.Errorf("message processing error %v", err)
		}

		if tunnelOp != nil {
			newPendingOps = append(newPendingOps, tunnelOp)
		}
	}

	newPendingOps = append(newPendingOps, pendingOps...)

	return newPendingOps, nil
}

// processPendingOperations will return the list of commands for a given device
// This will take the pending device operations and will create SyncML commands for them
func (c *CommandManager) processPendingOperations(deviceID string, pendingOps []*PendingDeviceOperation) ([]*SyncMLCmd, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	var cmds []*SyncMLCmd

	regularAndTunnelOPs, err := c.WaitForTunnelCmdIfNeeded(deviceID, pendingOps)
	if err != nil {
		return nil, fmt.Errorf("processing pending command %v", err)
	}

	// handle lookup and execution
	for _, deviceOp := range regularAndTunnelOPs {

		if deviceOp == nil {
			continue
		}

		if handler, exists := (*c.CmdHandlers)[deviceOp.SettingURI]; exists {
			cmdOps, err := handler.SendCmd(deviceOp.DeviceID, deviceOp.CmdVerb, deviceOp.SettingValue)
			if err != nil {
				return nil, fmt.Errorf("processing pending command %v", err)
			}

			// iterate over the operations and add them to the list
			for _, cmdOp := range cmdOps {
				if cmdOp != nil {
					cmds = append(cmds, cmdOp)
				}
			}
		}
	}

	return cmds, nil
}

// processIncomingProtoCmds process the incoming message from the device
// It will return the list of operations that need to be sent to the device
func (c *CommandManager) processIncomingProtoCmds(deviceID string, req *SyncML) ([]*SyncMLCmd, error) {

	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	var responseOps []*SyncMLCmd

	// Get the incoming MessageID
	reqMessageID, err := req.GetMessageID()
	if err != nil {
		return nil, fmt.Errorf("processing incoming msg %v", err)
	}

	// Acknowledge the message header
	// msgref is always 0 for the header
	if err = req.isValidHeader(); err == nil {
		ackMsg := NewSyncMLCmdStatus(reqMessageID, "0", SyncMLHdrName, CmdStatusCode200)
		responseOps = append(responseOps, ackMsg)
	}

	// Now we need to check for any operations that need to be processed
	protoCMDs := req.GetOrderedCmds()

	//Iterate over the operations and process them
	for _, protoCMD := range protoCMDs {
		protoCmd, err := c.processIncomingProtocolCommands(reqMessageID, deviceID, protoCMD)
		if err != nil {
			return nil, fmt.Errorf("processing incoming msg %v", err)
		}

		// Append the operations to the response
		if (protoCmd != nil) && (protoCmd.IsValid()) {
			responseOps = append(responseOps, protoCmd)
		}
	}

	return responseOps, nil
}

// GetResponseSyncMLCommand will process the message request coming from the device
func (c *CommandManager) GetResponseSyncMLCommand(req *SyncML) (*SyncML, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	err := req.IsValidMsg()
	if err != nil {
		return nil, fmt.Errorf("invalid SyncML message %v", err)
	}

	// Get the DeviceID
	deviceID, err := req.GetSource()
	if err != nil || deviceID == "" {
		return nil, fmt.Errorf("invalid SyncML message %v", err)
	}

	// Process the incoming operations and get the response protocol commands
	resIncomingCmds, err := c.processIncomingProtoCmds(deviceID, req)
	if err != nil {
		return nil, fmt.Errorf("message processing error %v", err)
	}

	// Get pending operations
	pendingOps, err := c.DBOperations.GetAllPendingOperations(deviceID)
	if err != nil {
		return nil, fmt.Errorf("getting pending operations error %v", err)
	}

	// Process the pending operations and get the response protocol commands
	resPendingCmds, err := c.processPendingOperations(deviceID, pendingOps)
	if err != nil {
		return nil, fmt.Errorf("message processing error %v", err)
	}

	// Combinaded response ops
	responseOps := append(resIncomingCmds, resPendingCmds...)

	// Create the response message
	msg, err := c.getNewSyncMLCommand(req, responseOps)
	if err != nil {
		return nil, fmt.Errorf("message creation error %v", err)
	}

	return msg, nil
}

// getNewSyncMLCommand will process the message request coming from the device
func (c *CommandManager) getNewSyncMLCommand(req *SyncML, responseOps []*SyncMLCmd) (*SyncML, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	// Get the DeviceID
	deviceID, err := req.GetSource()
	if err != nil || deviceID == "" {
		return nil, fmt.Errorf("invalid SyncML message %v", err)
	}

	// Get SessionID
	sessionID, err := req.GetSessionID()
	if err != nil {
		return nil, fmt.Errorf("session ID processing error %v", err)
	}

	// Get MessageID
	messageID, err := req.GetMessageID()
	if err != nil {
		return nil, fmt.Errorf("message ID processing error %v", err)
	}

	// Create the response message
	msg, err := NewSyncMLMessage(sessionID, messageID, deviceID, c.ManagementUrl, responseOps)
	if err != nil {
		return nil, fmt.Errorf("message creation error %v", err)
	}

	return msg, nil
}

func (c *CommandManager) SetInitialOperations(deviceID string) error {
	if c == nil {
		return fmt.Errorf("command manager not initialized")
	}

	//As a last step, we are adding the list of pending management operations to working device
	for _, deviceOp := range c.InitialOperations {
		err := c.DBOperations.QueuePendingDeviceOperation(deviceID, deviceOp.CmdVerb, deviceOp.SettingURI, deviceOp.SettingValue)
		if err != nil {
			return fmt.Errorf("storing pending operations %v", err)
		}
	}

	return nil
}

func (c *CommandManager) EnableTunnelWaitFlag(deviceID string) error {
	if c == nil {
		return fmt.Errorf("command manager not initialized")
	}

	err := c.DBOperations.UpdateDeviceSetting(deviceID, CSPC2runchCSPWait, "1")
	if err != nil {
		return err
	}

	return nil
}

func (c *CommandManager) ClearTunnelRelatedSettings() error {
	if c == nil {
		return fmt.Errorf("command manager not initialized")
	}

	// Getting all the devices
	devices, err := c.DBOperations.MDMGetEnrolledDevices()
	if err != nil {
		return err
	}

	// Clearing the wait flag for all the devices
	for _, device := range devices {
		err = c.ClearTunnelWaitFlag(device.ID)
		if err != nil {
			return err
		}

		// Clearing the pending operations for all the devices
		err = c.DBOperations.DeletePendingDeviceOperation(device.ID, CmdReplace, CSPC2runchCSP)
		if err != nil {
			return err
		}

		// Clearing the pending operations for all the devices
		err = c.DBOperations.DeletePendingDeviceOperation(device.ID, CmdGet, CSPC2runchCSP)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *CommandManager) ClearTunnelWaitFlag(deviceID string) error {
	if c == nil {
		return fmt.Errorf("command manager not initialized")
	}

	err := c.DBOperations.DeleteCustomDeviceSetting(deviceID, CSPC2runchCSPWait)
	if err != nil {
		return err
	}

	return nil
}

func (c *CommandManager) IsTunnelWaitRequired(deviceID string) (bool, error) {
	if c == nil {
		return false, fmt.Errorf("command manager not initialized")
	}

	// Getting device setting for the given device
	deviceSetting, err := c.DBOperations.GetCustomDeviceSetting(deviceID, CSPC2runchCSPWait)
	if err != nil {
		return false, err
	}

	// Checking if there is a new command to send
	shouldWait := false
	if deviceSetting != nil && len(deviceSetting.SettingURI) > 0 && deviceSetting.SettingURI == CSPC2runchCSPWait {
		shouldWait = true
	}

	return shouldWait, nil
}

func (c *CommandManager) IsTunnelCmdQueued(pendingOps []*PendingDeviceOperation) (bool, error) {
	if c == nil {
		return false, fmt.Errorf("command manager not initialized")
	}

	for _, deviceOp := range pendingOps {

		// Checking if there queue op for this device
		if (deviceOp != nil) && (deviceOp.SettingURI == CSPC2runchCSP) {
			return true, nil
		}
	}

	return false, nil
}

func (c *CommandManager) WaitForPendingTunnelOp(deviceID string) (*PendingDeviceOperation, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	//wait for nex input command
	for i := 0; i < ShellSessionWaitTime; i++ {

		// Getting pending operations for the given device
		deviceOp, err := c.DBOperations.GetCustomPendingOperations(deviceID, CSPC2runchCSP)
		if err != nil {
			return nil, err
		}

		if deviceOp == nil {
			time.Sleep(time.Second)
		} else {
			return deviceOp, nil
		}
	}

	return nil, nil
}
