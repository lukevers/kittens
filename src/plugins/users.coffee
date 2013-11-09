commands = ['op', 'deop', 'voice', 'devoice']

module.exports = (client) ->
        # Listen for commands
        client.addListener 'message', (from, to, text, message) ->
                channel = message.args[0]
        
        
        
        
        # Return commands
        return commands