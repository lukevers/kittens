################
### COMMANDS ###
################

commands = ['op', 'deop', 'voice', 'devoice']

##############
### MODULE ###
##############

module.exports = (client, pc) ->
        # Listen for commands
        client.addListener 'message', (from, to, text, message) ->
                channel = message.args[0]

        # Listen for joins
        client.addListener 'join', (channel, nick, message) ->
                
                


        
        # Return commands
        return commands