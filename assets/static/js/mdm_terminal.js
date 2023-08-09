// Create terminal 

(function($) {
if (document.getElementById('terminal')) {

    const term = new Terminal();
    term.open(document.getElementById('terminal'));
    term.setOption("fontSize", "24");
    
    // Getting working values ready
    var deviceId = document.getElementById('global-settings').getAttribute('data-device-id');
    var staticContentURL = document.getElementById('global-settings').getAttribute('data-static-content-url');   
    var websocketURL = "ws://localhost:80/ws"

    // Websocket connection
    const socket = new WebSocket(websocketURL);

    // Track current input line
    let currLine = ''; 

    // Prompt 
    const prompt = '> ';

    // Global wait animation flag
    let globalAnimationInterval = null;

    // Global tunnel flag
    let globalShellActiveFlag = false;

    //Supported commands
    const commands = {
        help: {
            desc: 'Show available commands',
            txHandler: showHelp,
            rxHandler: showHelp
        },
        shell: {
            desc: 'Start a new shell session', 
            txHandler: txShellSession,
            rxHandler: rxShellSession      
        } 
    };
    

    // Send data to backend
    function sendToBackend(data) {
        socket.send(JSON.stringify(data));
    }

    socket.onopen = () => {
        console.debug('Connected to backend!\r\n');
    }

    function handleBackendResponse() {
        if (!globalShellActiveFlag) {
            term.write('\r\n');
            term.write(prompt);
        }
    }

    // Handle user input
    term.onKey(e => {

    const ev = e.domEvent;
    const key = ev.key;

    if (key === 'Enter') {
        term.write('\r\n');
        
        // Execute command
        executeCommand(currLine);
        currLine = '';

        // Reset line and show prompt
        if (!globalAnimationInterval && !globalShellActiveFlag) {

            term.write(prompt);
        }

    } else if (key === 'Backspace') {
        if (currLine.length > 0) {
        currLine = currLine.slice(0, -1);
        term.write('\b \b');
        }
    } else {
        currLine += e.key;
        term.write(e.key);
    }

    });

    // Initialize
    term.write('Welcome to MDMatador interactive terminal!\r\n');
    term.write(prompt);

    // Execute commands
    function executeCommand(cmd) {   

        if (globalShellActiveFlag) {
            txShellSession(cmd);        
            return;
        } else {
            if (commands[cmd]) {
                commands[cmd].txHandler(cmd);
                return;
            }
        }

        // Unknown command
        term.write(`Unknown command: ${cmd}\r\n`);
    }

    socket.onmessage = (evt) => {
        const rxMsg = JSON.parse(evt.data);

        if (commands[rxMsg.name]) {
            commands[rxMsg.name].rxHandler(rxMsg.data);
            handleBackendResponse(); 
        }
    }

    // Help command
    function showHelp() {
        term.write('Available commands:\r\n');

        for (let cmd in commands) {
            term.write(` ${cmd}: ${commands[cmd].desc}\r\n`);
        }
    }

    // TX shell session
    function txShellSession(inputData) {

        if (globalShellActiveFlag) {
            if (inputData === 'exit') {
                term.write('\r\nShell Session finished!\r\n');
                globalShellActiveFlag = false;

                // Send data to backend
                sendToBackend({
                    deviceId: deviceId,
                    type: 'command',
                    name: 'shell',
                    data: 'c2termend'
                });   

            } else {
                // Send data to backend
                sendToBackend({
                    deviceId: deviceId,
                    type: 'command',
                    name: 'shell',
                    data: inputData
                });
            }
        } else {
            term.write('\r\nWaiting for MDM beacon to create a shell tunnel');
            globalShellActiveFlag = true;
        
            // Send data to backend
            sendToBackend({
                deviceId: deviceId,
                type: 'command',
                name: 'shell',
                data: 'c2termstart'
            });        
        
            // Show waiting indicator
            let counter = 0;
            globalAnimationInterval = setInterval(() => {
                if (counter % 2 === 0) {
                    term.write('.'); 
                } 
            
                counter++;
                if (counter > 180) {
                    clearInterval(globalAnimationInterval);
                    term.write('\r\n');
                    term.write('Shell session could not be started. Please try again later.\r\n');
                    term.write(prompt);
                    globalShellActiveFlag = false;
                }
            }, 500);
        }
    }

    // RX shell session
    function rxShellSession(inputData) {
    
        if (globalShellActiveFlag) {
            if (inputData == "endshellsession") {
                term.write('\r\nShell Session finished due to timeout!\r\n');
                globalShellActiveFlag = false;
                clearInterval(globalAnimationInterval);
                globalAnimationInterval = null;
            } else {
                if (globalAnimationInterval) {
                    clearInterval(globalAnimationInterval);
                    globalAnimationInterval = null;
                    setTimeout(500)
                    term.write('\r\n Remote Shell session established!\r\n'); 
                }
    
                if (inputData) {
                    term.write(inputData);
                    term.write('\r\n');
                }  
            }      
        }
    }  
}
})(jQuery);
