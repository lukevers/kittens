####################
### LOAD MODULES ###
####################

fs       = require 'fs'
readline = require 'readline'

##############
### COLORS ###
##############

green = `'\033[0;32m'`
reset = `'\033[0m'`
red   = `'\033[0;31m'`

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

i = 0

if !fs.existsSync './config.json'
        config = []
        tpconf =
                'botName': 'kittens'
                'userName': 'kittens'
                'realName': 'kittens'
                'server': 'irc.hypeirc.net'
                'port': 6667
                'autoConnect': true
                'commandSymbol': '!'
                'channels': []
else
        config = require './config.json'

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
        fs.writeFileSync './config.json', JSON.stringify config
        console.log green + 'Config file generated' + reset
        require('./lib/irc')(config)

# Again?
again = (a) ->
        if a.substring(0, 1) is 'y' or a.substring(0, 1) is 'Y'
                i++
                addServer()
        else
                init()

# Get channels, go to init
channels = (a) ->
        a = a.replace(/\ /g, '')
        config[i].channels = a.split(',') if a != ''
        ask 'Add another server? (yes/no)', again

# Get commandsymbol, go to channels
commandsymbol = (a) ->
        config[i].commandSymbol = a if a != ''
        config[i].commandSymbol = '!' if a is ''
        ask 'Channels', channels

# Get port, go to commandsymbol
port = (a) ->
        config[i].port = parseInt(a) if a != ''
        config[i].port = 6667 if a is ''
        ask 'Command Symbol', commandsymbol

# Get server, go to port
server = (a) ->
        config[i].server = a if a != ''
        config[i].server = 'irc.hypeirc.net' if a is ''
        ask 'Port', port

# Get realname, go to server
realname = (a) ->
        config[i].realName = a if a != ''
        config[i].realName = 'kittens' if a is ''
        ask 'Server', server

# Get username, go to realname
username = (a) ->
        config[i].userName = a if a != ''
        config[i].userName = 'kittens' if a is ''
        ask 'Real Name', realname

# Get botname, go to username
botname = (a) ->
        config[i].botName = a if a != ''
        config[i].botName = 'kittens' if a is ''
        ask 'Username', username

clone = (obj) ->
        if not obj? or typeof obj isnt 'object'
                return obj
                
        newInstance = new obj.constructor()

        for key of obj
                newInstance[key] = clone obj[key]
        
        return newInstance

# Add server
addServer = ->
        config[i] = clone(tpconf)
        ask 'Bot Name', botname

# If `./config.json` does not exist, make a new one.
if !fs.existsSync './config.json'
        console.log red + '\nNo config file was found!' + reset + ' Creating a new one.'
        console.log 'Leave each blank for their default value'
        addServer()
else
        require('./lib/irc')(config)