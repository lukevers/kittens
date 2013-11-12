# Load  modules
fs = require 'fs'

# Load the configuration file
Config = if fs.existsSync './config.json' then require './config.json' else require './example.json'

###########
### IRC ###
###########

# Load modules for IRC
irc = require 'irc'

# Create client
client = new irc.Client Config.IRC.server, Config.IRC.botName, Config.IRC

# Commands
commands = [Config.IRC.commandSymbol+'help']

# Add error handling
client.addListener 'error', (message) ->
        console.log 'error: ', message

# Listen for the help command
client.addListener 'message', (from, to, text, message) ->
        cmds = commands.join(' '+Config.IRC.commandSymbol).replace(/,|  /g, ' '+Config.IRC.commandSymbol)
        if message.args[1].indexOf(Config.IRC.commandSymbol + 'help') is 0
                client.say message.args[0], from + ': ' + cmds

# Load plugins
fs.readdir './lib/plugins', (err, files) ->
        for i in [0..files.length-1] by 1
                p = require('./lib/plugins/' + files[i]) client
                for j in [0..p.length-1] by 1
                        commands.push p[j] unless typeof p[j] is 'undefined'

###########
### LOG ###
###########

require('./lib/logging') client
