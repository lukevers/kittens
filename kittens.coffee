####################
### LOAD MODULES ###
####################

fs       = require 'fs'
irc      = require 'irc'
readline = require 'readline'

##############
### COLORS ###
##############

green = `'\033[0;32m'`
reset = `'\033[0m'`
red = `'\033[0;31m'`

####################
### KITTENS LOGO ###
####################

console.log green

console.log " _     _  _    _"
console.log "| |   (_)| |  | |"
console.log "| | __ _ | |_ | |_  ___  _ __   ___"
console.log "| |/ /| || __|| __|/ _ \\| '_ \\ / __|"
console.log "|   < | || |_ | |_|  __/| | | |\\__ \\"
console.log "|_|\\_\\|_| \\__| \\__|\\___||_| |_||___/"

console.log reset

###################
### INIT CONFIG ###
###################

config =
        'botName': 'kittens'
        'username': 'kittens'
        'realName': 'kittens'
        'server': 'irc.hypeirc.net'
        'port': 6667
        'autoConnect': true
        'commandSymbol': '!'
        'channels': []

####################
### CHECK CONFIG ###
####################

# Ask the user a question
ask = (q, next) ->
        line = readline.createInterface(process.stdin, process.stdout)
        line.question q+': ', (answer) ->
                line.close()
                next(answer)

# Create config
init = ->
        fs.writeFileSync './config.json', JSON.stringify(config)
        console.log green + 'Config file generated!' + reset

# Get channels, go to init
channels = (a) ->
        a = a.replace(/\ /g, '')
        config.channels = a.split(',') if a != ''
        init()

# Get commandsymbol, go to channels
commandsymbol = (a) ->
        config.commandSymbol = a if a != ''
        ask 'Channels', channels

# Get port, go to commandsymbol
port = (a) ->
        config.port = parseInt(a) if a != ''
        ask 'Command Symbol', commandsymbol

# Get server, go to port
server = (a) ->
        config.server = a if a != ''
        ask 'Port', port

# Get realname, go to server
realname = (a) ->
        config.realName = a if a != ''
        ask 'Server', server

# Get username, go to realname
username = (a) ->
        config.username = a if a != ''
        ask 'Real Name', realname

# Get botname, go to username
botname = (a) ->
        config.botName = a if a != ''
        ask 'Username', username

# If `./config.json` does not exist, make a new one.
if !fs.existsSync './config.json'
        console.log red + '\nNo config file was found!' + reset + ' Creating a new one.'
        console.log 'Leave each blank for their default value'
        ask 'Bot Name', botname

####################
### START UP IRC ###
####################

# Create client
client = new irc.Client config.server config.botName config

# Add error handling
client.addListener 'error', (message) ->
        console.log 'error: ' + message

