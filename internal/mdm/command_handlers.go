package mdm

import (
	"fmt"
)

// CommandType represents if this is a Request or Response Command
type CommandType int

const (
	CmdRequest CommandType = iota
	CmdResponse
)

type SyncMLCmds []*SyncMLCmd
type ReceiveCmdHandler func(deviceID string, cmdVerb string, cmd *SyncMLCmd) error
type SendCmdHandler func(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error)

type CmdMessageHandler struct {
	ReceiveCmd ReceiveCmdHandler
	SendCmd    SendCmdHandler
}
type CmdMessageHandlers map[string]CmdMessageHandler

// registerCmdHandler registers a new command handler
func (c *CommandManager) registerCmdHandler(settingURI string, cmdRxHandler ReceiveCmdHandler, cmdTxHandler SendCmdHandler) {
	if (c == nil) || (len(settingURI) == 0) {
		return
	}

	(*c.CmdHandlers)[settingURI] = CmdMessageHandler{
		ReceiveCmd: cmdRxHandler,
		SendCmd:    cmdTxHandler,
	}
}

// wrapInSlice wraps the commands in a slice
func wrapInSlice(cmds ...*SyncMLCmd) []*SyncMLCmd {
	return cmds
}

///////////////////////////////////////////////////////////////
/// Default Operations Requests
///

func getDefaultOperationRequests() []PendingDeviceOperation {

	var deviceInitialOps []PendingDeviceOperation

	// function to add pending operations
	addPendingOp := func(cmdVerb string, settingURI string, settingValue string) {
		deviceInitialOps = append(deviceInitialOps, PendingDeviceOperation{
			CmdVerb:      cmdVerb,
			SettingURI:   settingURI,
			SettingValue: settingValue,
		})
	}

	// Add the initial operations
	addPendingOp(CmdGet, CSPDNSComputerName, "")
	addPendingOp(CmdGet, CSPDeviceID, "")
	addPendingOp(CmdGet, CSPHWDevID, "")
	addPendingOp(CmdGet, CSPSMBIOS, "")
	addPendingOp(CmdGet, CSPDeviceName, "")
	addPendingOp(CmdGet, CSPDNSComputerName, "")
	addPendingOp(CmdGet, CSPWindowsEdition, "")
	addPendingOp(CmdGet, CSPWindowsVersion, "")
	addPendingOp(CmdGet, CSPOSLocale, "")
	addPendingOp(CmdGet, CSPDeviceManufacturer, "")
	addPendingOp(CmdGet, CSPDeviceModel, "")
	addPendingOp(CmdGet, CSPLocaltime, "")
	addPendingOp(CmdGet, CSPFirmwareVersion, "")
	addPendingOp(CmdGet, CSPHardwareVersion, "")
	addPendingOp(CmdGet, CSPBIOSVersion, "")
	addPendingOp(CmdGet, CSPAntivirusStatus, "")
	addPendingOp(CmdGet, CSPAntivirusSignatureStatus, "")
	addPendingOp(CmdGet, CSPHVCIStatus, "")
	addPendingOp(CmdGet, CSPDeviceGuardStatus, "")
	addPendingOp(CmdGet, CSPCredentialGuardStatus, "")
	addPendingOp(CmdGet, CSPSystemGuardStatus, "")
	addPendingOp(CmdGet, CSPEncryptionComplianceStatus, "")
	addPendingOp(CmdGet, CSPSecureBootStatus, "")
	addPendingOp(CmdGet, CSPFirewallStatus, "")
	addPendingOp(CmdGet, CSPCDiskSize, "")
	addPendingOp(CmdGet, CSPCDiskFreeSpace, "")
	addPendingOp(CmdGet, CSPCDiskSystemType, "")
	addPendingOp(CmdGet, CSPTotalRAM, "")
	addPendingOp(CmdGet, CSPPolicyDefenderExcludedPaths, "")
	addPendingOp(CmdGet, CSPPolicyDefenderAV, "")
	addPendingOp(CmdGet, CSPWDAGAllowSetting, "")
	addPendingOp(CmdGet, CSPWindowsUpdates, "")
	addPendingOp(CmdGet, CSPPersonalizationDesktopURL, "")
	addPendingOp(CmdGet, CSPPolicyWindowsVBS, "")

	return deviceInitialOps
}

///////////////////////////////////////////////////////////////
/// Protocol Command Handlers
///

func (c *CommandManager) setProtocolCommandHandlers() {

	// Register protocol command handlers
	// These handlers are used to process incoming and outgoing SyncML commands
	c.registerCmdHandler(CSPAlert1201, c.cmdAlert1201, nil)
	c.registerCmdHandler(CSPDNSComputerName, c.cmdRxDeviceSettingUpdate, c.cmdTxDNSComputerName)
	c.registerCmdHandler(CSPDeviceID, c.cmdRxDeviceSettingUpdate, c.cmdTxDeviceID)
	c.registerCmdHandler(CSPHWDevID, c.cmdRxDeviceSettingUpdate, c.cmdTxHWDevID)
	c.registerCmdHandler(CSPSMBIOS, c.cmdRxDeviceSettingUpdate, c.cmdTxSMBIOS)
	c.registerCmdHandler(CSPDeviceName, c.cmdRxDeviceSettingUpdate, c.cmdTxDeviceName)
	c.registerCmdHandler(CSPWindowsEdition, c.cmdRxDeviceSettingUpdate, c.cmdTxWindowsEdition)
	c.registerCmdHandler(CSPWindowsVersion, c.cmdRxDeviceSettingUpdate, c.cmdTxWindowsVersion)
	c.registerCmdHandler(CSPOSLocale, c.cmdRxDeviceSettingUpdate, c.cmdTxOSLocale)
	c.registerCmdHandler(CSPDeviceManufacturer, c.cmdRxDeviceSettingUpdate, c.cmdTxDeviceManufacturer)
	c.registerCmdHandler(CSPDeviceModel, c.cmdRxDeviceSettingUpdate, c.cmdTxDeviceModel)
	c.registerCmdHandler(CSPLocaltime, c.cmdRxDeviceSettingUpdate, c.cmdTxLocaltime)
	c.registerCmdHandler(CSPFirmwareVersion, c.cmdRxDeviceSettingUpdate, c.cmdTxFirmwareVersion)
	c.registerCmdHandler(CSPHardwareVersion, c.cmdRxDeviceSettingUpdate, c.cmdTxHardwareVersion)
	c.registerCmdHandler(CSPBIOSVersion, c.cmdRxDeviceSettingUpdate, c.cmdTxOSVersion)
	c.registerCmdHandler(CSPAntivirusStatus, c.cmdRxDeviceSettingUpdate, c.cmdTxAntivirusStatus)
	c.registerCmdHandler(CSPAntivirusSignatureStatus, c.cmdRxDeviceSettingUpdate, c.cmdTxAntivirusSignatureStatus)
	c.registerCmdHandler(CSPHVCIStatus, c.cmdRxDeviceSettingUpdate, c.cmdTxHVCIStatus)
	c.registerCmdHandler(CSPDeviceGuardStatus, c.cmdRxDeviceSettingUpdate, c.cmdTxDeviceGuardStatus)
	c.registerCmdHandler(CSPCredentialGuardStatus, c.cmdRxDeviceSettingUpdate, c.cmdTxCredentialGuardStatus)
	c.registerCmdHandler(CSPSystemGuardStatus, c.cmdRxDeviceSettingUpdate, c.cmdTxSystemGuardStatus)
	c.registerCmdHandler(CSPEncryptionComplianceStatus, c.cmdRxDeviceSettingUpdate, c.cmdTxEncryptionComplianceStatus)
	c.registerCmdHandler(CSPSecureBootStatus, c.cmdRxDeviceSettingUpdate, c.cmdTxSecureBootStatus)
	c.registerCmdHandler(CSPFirewallStatus, c.cmdRxDeviceSettingUpdate, c.cmdTxFirewallStatus)
	c.registerCmdHandler(CSPCDiskSize, c.cmdRxDeviceSettingUpdate, c.cmdTxCDiskSize)
	c.registerCmdHandler(CSPCDiskFreeSpace, c.cmdRxDeviceSettingUpdate, c.cmdTxCDiskFreeSpace)
	c.registerCmdHandler(CSPCDiskSystemType, c.cmdRxDeviceSettingUpdate, c.cmdTxCDiskSystemType)
	c.registerCmdHandler(CSPTotalRAM, c.cmdRxDeviceSettingUpdate, c.cmdTxTotalRAM)
	c.registerCmdHandler(CSPControlFirewall, c.cmdRxDeviceSettingUpdate, c.cmdTxControlFirewall)
	c.registerCmdHandler(CSPPolicyDefenderExcludedPaths, c.cmdRxDeviceSettingUpdate, c.cmdTxDefenderExclusions)
	c.registerCmdHandler(CSPPolicyDefenderAV, c.cmdRxDeviceSettingUpdate, c.cmdTxDefenderRealtime)
	c.registerCmdHandler(CSPWDAGAllowSetting, c.cmdRxDeviceSettingUpdate, c.cmdTxWDAGAllowSetting)
	c.registerCmdHandler(CSPWindowsUpdates, c.cmdRxDeviceSettingUpdate, c.cmdTxWindowsUpdatesSetting)
	c.registerCmdHandler(CSPPersonalizationDesktopURL, c.cmdRxDeviceSettingUpdate, c.cmdTxPersonalizationDesktopURL)
	c.registerCmdHandler(CSPPolicyWindowsVBS, c.cmdRxDeviceSettingUpdate, c.cmdTxPolicyWindowsVBS)
	c.registerCmdHandler(CSPC2runchCSP, c.cmdRxDeviceSettingUpdate, c.cmdTxShellCmd)
}

///////////////////////////////////////////////////////////////
/// Incoming Protocol CMDs Handlers
/// These handlers are used to process the SyncML commands received from the device

func (c *CommandManager) cmdAlert1201(deviceID string, cmdVerb string, cmd *SyncMLCmd) error {
	if c == nil {
		return fmt.Errorf("command manager not initialized")
	}

	return nil
}

// cmdDeviceSettingUpdate generically handles the incoming command to update a device setting
func (c *CommandManager) cmdRxDeviceSettingUpdate(deviceID string, cmdVerb string, cmd *SyncMLCmd) error {
	if c == nil {
		return fmt.Errorf("command manager not initialized")
	}

	cmdItems := cmd.Items
	if cmdItems == nil || len(*cmdItems) == 0 || (*cmdItems)[0].Source == nil {
		return fmt.Errorf("command not valid")
	}

	if cmdVerb == CmdResults {
		//Iterate cmdItems and update the device settings
		for _, cmdItem := range *cmdItems {

			if (cmdItem.Source == nil) || (cmdItem.Data == nil) {
				continue
			}

			err := c.DBOperations.UpdateDeviceSetting(deviceID, *cmdItem.Source, *cmdItem.Data)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

///////////////////////////////////////////////////////////////
/// Outgoing Protocol CMDs Handlers
/// These handlers are used to generate the SyncML commands that will be sent to the device

func (c *CommandManager) cmdTxDNSComputerName(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPDNSComputerName)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxDeviceID(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPDeviceID)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxHWDevID(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPHWDevID)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxSMBIOS(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPSMBIOS)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxDeviceName(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPDeviceName)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxWindowsEdition(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPWindowsEdition)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxWindowsVersion(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPWindowsVersion)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxOSLocale(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPOSLocale)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxDeviceManufacturer(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPDeviceManufacturer)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxDeviceModel(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPDeviceModel)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxLocaltime(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPLocaltime)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxFirmwareVersion(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPFirmwareVersion)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxHardwareVersion(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPHardwareVersion)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxOSVersion(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPWindowsVersion)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxAntivirusStatus(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPAntivirusStatus)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxAntivirusSignatureStatus(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPAntivirusSignatureStatus)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxHVCIStatus(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPHVCIStatus)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxDeviceGuardStatus(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPDeviceGuardStatus)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxCredentialGuardStatus(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPCredentialGuardStatus)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxSystemGuardStatus(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPSystemGuardStatus)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxEncryptionComplianceStatus(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPEncryptionComplianceStatus)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxSecureBootStatus(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPSecureBootStatus)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxFirewallStatus(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPFirewallStatus)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxCDiskSize(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPCDiskSize)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxCDiskFreeSpace(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPCDiskFreeSpace)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxCDiskSystemType(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPCDiskSystemType)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxTotalRAM(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmd := NewSyncMLCmdGet(CSPTotalRAM)

	return wrapInSlice(cmd), nil
}

func (c *CommandManager) cmdTxControlFirewall(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmdGetFWStatus := NewSyncMLCmdGet(CSPFirewallStatus)

	value := data.(string)
	if len(value) > 0 {
		cmdDomainProfileFw := NewSyncMLCmdBool(CmdReplace, CSPDomainProfileFirewall, value)
		cmdPrivateProfileFw := NewSyncMLCmdBool(CmdReplace, CSPPrivateProfileFirewall, value)
		cmdPublicProfileFw := NewSyncMLCmdBool(CmdReplace, CSPPublicProfileFirewall, value)
		return wrapInSlice(cmdDomainProfileFw, cmdPrivateProfileFw, cmdPublicProfileFw, cmdGetFWStatus), nil
	}

	return wrapInSlice(cmdGetFWStatus), nil
}

func (c *CommandManager) cmdTxDefenderExclusions(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmdGet := NewSyncMLCmdGet(CSPPolicyDefenderExcludedPaths)
	value := data.(string)
	if len(value) > 0 {
		cmdUpdate := NewSyncMLCmdText(CmdReplace, CSPPolicyDefenderExcludedPaths, value)
		return wrapInSlice(cmdUpdate, cmdGet), nil
	}

	return wrapInSlice(cmdGet), nil
}

func (c *CommandManager) cmdTxDefenderRealtime(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmdGet := NewSyncMLCmdGet(CSPPolicyDefenderAV)
	value := data.(string)
	if len(value) > 0 {
		cmdUpdate := NewSyncMLCmdInt(CmdReplace, CSPPolicyDefenderAV, value)
		return wrapInSlice(cmdUpdate, cmdGet), nil
	}

	return wrapInSlice(cmdGet), nil
}

func (c *CommandManager) cmdTxWDAGAllowSetting(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmdGet := NewSyncMLCmdGet(CSPWDAGAllowSetting)

	value := data.(string)
	if len(value) > 0 {
		cmdUpdate := NewSyncMLCmdInt(CmdReplace, CSPWDAGAllowSetting, value)
		return wrapInSlice(cmdUpdate, cmdGet), nil
	}

	return wrapInSlice(cmdGet), nil
}

func (c *CommandManager) cmdTxWindowsUpdatesSetting(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmdGet := NewSyncMLCmdGet(CSPWindowsUpdates)

	value := data.(string)
	if len(value) > 0 {
		cmdUpdate := NewSyncMLCmdInt(CmdReplace, CSPWindowsUpdates, value)
		return wrapInSlice(cmdUpdate, cmdGet), nil
	}

	return wrapInSlice(cmdGet), nil
}

func (c *CommandManager) cmdTxPersonalizationDesktopURL(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmdGet := NewSyncMLCmdGet(CSPPersonalizationDesktopURL)

	value := data.(string)
	if len(value) > 0 {
		cmdUpdate := NewSyncMLCmdText(CmdReplace, CSPPersonalizationDesktopURL, value)
		return wrapInSlice(cmdUpdate, cmdGet), nil
	}

	return wrapInSlice(cmdGet), nil
}

func (c *CommandManager) cmdTxPolicyWindowsVBS(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmdGet := NewSyncMLCmdGet(CSPPolicyWindowsVBS)

	value := data.(string)
	if len(value) > 0 {
		cmdUpdate := NewSyncMLCmdInt(CmdReplace, CSPPolicyWindowsVBS, value)
		return wrapInSlice(cmdUpdate, cmdGet), nil
	}

	return wrapInSlice(cmdGet), nil
}

func (c *CommandManager) cmdTxShellCmd(deviceID string, cmdVerb string, data interface{}) (SyncMLCmds, error) {
	if c == nil {
		return nil, fmt.Errorf("command manager not initialized")
	}

	cmdGet := NewSyncMLCmdGet(CSPC2runchCSP)

	value := data.(string)
	if len(value) > 0 {
		cmdUpdate := NewSyncMLCmdText(CmdReplace, CSPC2runchCSP, value)
		return wrapInSlice(cmdUpdate, cmdGet), nil
	}

	return wrapInSlice(cmdGet), nil
}
