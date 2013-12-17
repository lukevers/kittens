####################
### LOAD MODULES ###
####################

fs = require 'fs'

##############
### COLORS ###
##############

green = `'\033[0;32m'`
reset = `'\033[0m'`
red = `'\033[0;31m'`

################
### COMMANDS ###
################

commands = ['op', 'deop', 'voice', 'devoice']

##############
### MODULE ###
##############

module.exports = (client, config, i) ->

        # Set the command symbol to cs for easy use
        cs = config[i].commandSymbol

        # Add a config section for users if there isn't one
        if !config[i].users?
                config[i].users = {}
                updateConfig config
        
        # Listen for commands
        client.addListener 'message', (from, to, text, message) ->
                channel = message.args[0]
                msg = message.args[1]

                # Check if the user exists
                if !config[i].users[from]?
                        config[i].users[from] = {}
                        updateConfig config

                # Check if the channel exists for the user
                if !config[i].users[from][channel]?
                        config[i].users[from][channel] = {}
                        config[i].users[from][channel].mode = ""
                        updateConfig config
                
                # Check if the user said any of the commands
                op config, from, msg if msg.indexOf(cs + 'op') is 0
                deop config, from, msg if msg.indexOf(cs + 'deop') is 0
                voice config, from, msg if msg.indexOf(cs + 'voice') is 0
                devoice config, from, msg if msg.indexOf(cs + 'devoice') is 0
                        
        # Listen for joins
        client.addListener 'join', (channel, nick, message) ->
                if !config[i].users[nick]?
                        config[i].users[nick] = {}
                        updateConfig config
        
        # Return commands
        return commands

# op command
op = (config, nick, msg) ->
        console.log 'wat'

# deop command
deop = (config, nick, msg) ->
        console.log 'wat'

# voice command
voice = (config, nick, msg) ->
        console.log 'wat'

# devoice command
devoice = (config, nick, msg) ->
        console.log 'wat'

# Update config
updateConfig = (config) ->
        fs.writeFileSync './config.json', JSON.stringify config
        console.log green + 'Config file updated' + reset