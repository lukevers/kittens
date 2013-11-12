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

####################
### LOAD MODULES ###
####################

fs       = require 'fs'
readline = require 'readline'

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

ask = (q, next) ->
        line = readline.createInterface(process.stdin, process.stdout)
        line.question q+': ', (answer) ->
                line.close()
                next(answer)

init = ->
        fs.writeFileSync './config.json', JSON.stringify(config)
        console.log green + 'Config file generated!' + reset

channels = (a) ->
        a = a.replace(/\ /g, '')
        config.channels = a.split(',') if a != ''
        init()

commandsymbol = (a) ->
        config.commandSymbol = a if a != ''
        ask 'Channels', channels

port = (a) ->
        config.port = parseInt(a) if a != ''
        ask 'Command Symbol', commandsymbol

server = (a) ->
        config.server = a if a != ''
        ask 'Port', port

realname = (a) ->
        config.realName = a if a != ''
        ask 'Server', server

username = (a) ->
        config.username = a if a != ''
        ask 'Real Name', realname

botname = (a) ->
        config.botName = a if a != ''
        ask 'Username', username

if !fs.existsSync './config.json'
        console.log red + '\nNo config file was found!' + reset + ' Creating a new one.'
        console.log 'Leave each blank for their default value'
        ask 'Bot Name', botname

################
### START UP ###
################

