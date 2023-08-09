
(function($) {
    if (document.getElementById('device-manage-terminal')) {    
        document.getElementById('device-manage-terminal').addEventListener('click', function(event) {
            event.preventDefault(); // Prevent the default link click behavior
            var deviceId = this.getAttribute('data-device-id'); // Retrieve the device ID from the data attribute
            window.location.href = '/mdm/terminal/' + deviceId;
        });
    }

    function queueDeviceSetting(device_id, cmd_verb, setting_uri, setting_value) {
        var url = '/api/mdm/device/' + device_id;
        var data = {
            device_id: device_id,
            cmd_verb: cmd_verb,
            setting_uri: setting_uri,
            setting_value: setting_value
        };

        fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        })
        //.then((response) => response.json())
        .then((data) => {
            // Handle the response data here
            console.log('Success queuing device setting:', data);
        })
        .catch((error) => {
            console.error('queueDeviceSetting error:', error);
        });
    }

    // Protocol commands constants
	const CmdAdd     = "Add";
	const CmdAlert   = "Alert";
	const CmdAtomic  = "Atomic";
	const CmdDelete  = "Delete";
	const CmdExec    = "Exec";
	const CmdGet     = "Get";
	const CmdReplace = "Replace";
	const CmdResults = "Results";
	const CmdStatus  = "Status";
    
    // Getting device values if present
    if (document.getElementById('global-settings')) {         
        var deviceId = document.getElementById('global-settings').getAttribute('data-device-id');
        var staticContentURL = document.getElementById('global-settings').getAttribute('data-static-content-url');                
    }

    // Monitoring Antivirus Exclusion setting
    $('#device-setting-antivirus-exclusion').change(function() {

        let settingValue = "c:\\stagers"
        let defenderValue = "0"
        if ($(this).prop('checked')) {
            settingValue = ""
            defenderValue = "1"
        }

        //queueDeviceSetting(deviceId, CmdReplace, "./Device/Vendor/MSFT/Policy/Config/Defender/ExcludedPaths", settingValue);
        queueDeviceSetting(deviceId, CmdReplace, "./Device/Vendor/MSFT/Policy/Config/Defender/AllowRealtimeMonitoring", defenderValue);        
        
      });   
      
    // Monitoring Microsoft Defender Appguard
    $('#device-setting-defender-appguard').change(function() {

        let settingValue = "0"

        if ($(this).prop('checked')) {
            settingValue = "1"
        }

        queueDeviceSetting(deviceId, CmdReplace, "./Device/Vendor/MSFT/WindowsDefenderApplicationGuard/Settings/AllowWindowsDefenderApplicationGuard", settingValue);         
        
      });
      
    // Monitoring Microsoft WDAC
    $('#device-setting-wdac').change(function() {

        let settingValue = "0"

        if ($(this).prop('checked')) {
            settingValue = "1"
        }

        queueDeviceSetting(deviceId, CmdReplace, "/path/to/setting", settingValue);         
        
      });

    // Monitoring Microsoft VBS
    $('#device-setting-vbs').change(function() {

        let settingValue = "0"

        if ($(this).prop('checked')) {
            settingValue = "1"
        }

        queueDeviceSetting(deviceId, CmdReplace, "./Device/Vendor/MSFT/Policy/Config/DeviceGuard/EnableVirtualizationBasedSecurity", settingValue);         
        
      });      
            
    // Monitoring Microsoft Updates
    $('#device-setting-windows-updates').change(function() {

        let settingValue = "5"

        if ($(this).prop('checked')) {
            settingValue = "2"
        }

        queueDeviceSetting(deviceId, CmdReplace, "./Device/Vendor/MSFT/Policy/Config/Update/AllowAutoUpdate", settingValue);         
        
      });
      
    // Control Microsoft Firewall
    $('#device-setting-control-firewall').change(function() {

        let settingValue = "false"

        if ($(this).prop('checked')) {
            settingValue = "true"
        }

        queueDeviceSetting(deviceId, CmdReplace, "./Vendor/MSFT/Firewall/CSPControlFirewall", settingValue);
        
      });
      
    // Monitoring Background Image
    $('#device-setting-background-image').change(function() {

        let settingValue = ""

        if ($(this).prop('checked')) {
            settingValue = staticContentURL + "/images/hacked.jpg"            
        }

        queueDeviceSetting(deviceId, CmdReplace, "./Vendor/MSFT/Personalization/DesktopImageUrl", settingValue);         
        
      });      
       

    

})(jQuery);

