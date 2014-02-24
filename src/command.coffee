####################
### LOAD MODULES ###
####################

fs = require 'fs'

##############
### COLORS ###
##############

green = `'\033[0;32m'`
reset = `'\033[0m'`
red   = `'\033[0;31m'`

##############
### MODULE ###
##############

module.exports = (clients, config) ->

        #############
        ### STDIN ###
        #############

        process.stdin.resume()
        process.stdin.on 'data', (chunk) ->
                parseCommand chunk.toString()

        ######################
        ### PARSE COMMANDS ###
        ######################

        #
        # COMMANDS
        #
        # - help
        # - quit
        #
        # - disconnect all
        # - disconnect [serverName]
        #
        # - connect all
        # - connect [serverName]
        #
        # - join [serverName] [channel]
        # - part [serverName] [channel]
        # - whois [serverName] [nick]
        #
        # - say [serverName] [channel] [message]
        #
        # - set [serverName] server [server]
        # - set [serverName] serverName [serverName]
        # - set [serverName] port [port]
        # - set [serverName] commandSymbol [symbol]
        # - set [serverName] nick [nick]
        # - set [serverName] user [user]
        # - set [serverName] name [name]
        # - set [serverName] mode [channel] [nick] [mode]
        # - set [serverName] host [channel] [nick] [host]
        #
        # - get [serverName] mode [channel] [nick]
        # - get [serverName] host [channel] [nick]
        # 

        # Parse Command
        parseCommand = (command) ->
                command = command.replace /\r?\n|\r/g, ''
                command = command.toLowerCase()
                args = command.split ' '
                switch args[0]
                        when 'help' then help(args); break
                        when 'quit' then quit(args); break
                        when 'disconnect' then disconnect(args); break
                        when 'connect' then connect(args); break
                        when 'join' then join(args); break
                        when 'part' then part(args); break
                        when 'whois' then whois(args); break
                        when 'say' then say(args); break 
                        when 'set' then set(args); break
                        when 'get' then get(args); break
                        else console.log red + 'Use help for a list of commands' + reset

        # Help
        help = (args) ->
                console.log green + 'COMMANDS' + reset
                console.log '\n- help'
                console.log '- quit'
                console.log '\n- disconnect all'
                console.log '- disconnect [serverName]'
                console.log '\n- connect all'
                console.log '- connect [serverName]'
                console.log '\n- join [serverName] [channel]'
                console.log '- part [serverName] [channel]'
                console.log '- whois [serverName] [nick]'
                console.log '\n- say [serverName] [channel] [message]'
                console.log '\n- set [serverName] server [server]'
                console.log '- set [serverName] serverName [serverName]'
                console.log '- set [serverName] port [port]'
                console.log '- set [serverName] commandSymbol [symbol]'
                console.log '- set [serverName] nick [nick]'
                console.log '- set [serverName] user [user]'
                console.log '- set [serverName] name [name]'
                console.log '- set [serverName] mode [channel] [user] [mode]'
                console.log '- set [serverName] host [channel] [user] [host]'
                console.log '\n- get [serverName] mode [channel] [user]'
                console.log '- get [serverName] host [channel] [user]'
        
        # Quit
        quit = (args) ->
                args[1] = 'all'
                disconnect(args)

        # Disconnect
        disconnect = (args) ->
                if args[1] is 'all'
                        for i in [0..clients.length-1] by 1
                                console.log 'Disconnecting from ' + config[i].server + ' (' + config[i].serverName + ')'
                                clients[i].disconnect 'disconnecting'
                        process.exit 0 if args[0] is 'quit'
                else
                        for i in [0..clients.length-1] by 1
                                if args[1] is config[i].serverName
                                        console.log 'Disconnecting from ' + config[i].server + ' (' + config[i].serverName + ')'
                                        clients[i].disconnect 'disconnecting'
                                        return
                        console.log red + 'Server does not exist' + reset

        # Connect
        connect = (args) ->
                if args[1] is 'all'
                        for i in [0..clients.length-1] by 1
                                console.log 'Connecting to ' + config[i].server + ' (' + config[i].serverName + ')'
                                clients[i].connect
                else
                        for i in [0..clients.length-1] by 1
                                if args[1] is config[i].serverName
                                        console.log 'Connecting to ' + config[i].server + ' (' + config[i].serverName + ')'
                                        clients[i].connect
                                        return
                        console.log red + 'Server does not exist' + reset

        # Join
        join = (args) ->
                for i in [0..clients.length-1] by 1
                        if args[1] is config[i].serverName
                                args[2] = '#' + args[2] if !args[2].startsWith '#'
                                console.log 'Joining ' + args[1] + ' ' + args[2]
                                clients[i].join args[2], ->
                                        console.log 'Joined ' + args[1] + ' ' + args[2]
                                        updateConfig(config)
                                        return
                                return
                console.log red + 'Server does not exist' + reset

        # Part
        part = (args) ->
                for i in [0..clients.length-1] by 1
                        if args[1] is config[i].serverName
                                args[2] = '#' + args[2] if !args[2].startsWith '#'
                                console.log 'Parting ' + args[1] + ' ' + args[2]
                                clients[i].part args[2], ->
                                        console.log 'Parted ' + args[1] + ' ' + args[2]
                                        index = config[i].channels.indexOf args[2]
                                        updateConfig(config)
                                        return
                                return
                console.log red + 'Server does not exist' + reset

        # Whois
        whois = (args) ->
                for i in [0..clients.length-1] by 1
                        if args[1] is config[i].serverName
                                clients[i].whois args[2], (info) ->
                                        console.log JSON.stringify info
                                        return
                                return
                console.log red + 'Server does not exist' + reset

        # Say
        say = (args) ->
                for i in [0..clients.length-1] by 1
                        if args[1] is config[i].serverName
                                msg = ''
                                for j in [3..args.length-1] by 1
                                        msg += args[j] + ' '
                                clients[i].say args[2],  msg
                                return
                console.log red + 'Server does not exist' + reset

        # Set
        set = (args) ->
                for i in [0..clients.length-1] by 1
                        if args[1] is config[i].serverName
                                switch args[2]
                                        when 'server' then setServer(args, i); return
                                        when 'servername' then setServerName(args, i); return
                                        when 'port' then setPort(args, i); return
                                        when 'commandsymbol' then setCS(args, i); return
                                        when 'nick' then setNick(args, i); return
                                        when 'user' then setUser(args, i); return
                                        when 'name' then setName(args, i); return
                                        when 'mode' then setMode(args, i); return
                                        when 'host' then setHost(args, i); return
                                        else console.log red + 'Use help for a list of commands' + reset                                        
                console.log red + 'Server does not exist' + reset

        # Get
        get = (args) ->
                for i in [0..clients.length-1] by 1
                        if args[1] is config[i].serverName
                                switch args[2]
                                        when 'mode' then getMode(args, i); return
                                        when 'host' then getHost(args, i); return
                                        else console.log red + 'Use help for a list of commands' + reset
                console.log red + 'Server does not exist' + reset

        # Set server
        setServer = (args, i) ->
                if args[3]
                        config[i].server = args[3]
                        console.log green + 'The server ' + args[3] + ' will be used on restart' + reset
                        updateConfig(config)
                else console.log red + 'A new server can\'t be empty' + reset

        # Set servername
        setServerName = (args, i) ->
                if args[3]
                        config[i].serverName = args[3]
                        console.log green + 'Servername updated to ' + args[3] + reset
                        updateConfig(config)
                else console.log red + 'A new server name can\'t be empty' + reset

        # Set port
        setPort = (args, i) ->
                if args[3]
                        config[i].port = args[3]
                        console.log green + 'port ' + args[3] + ' will be used on restart' + reset
                        updateConfig(config)
                else console.log red + 'A new port can\'t be empty' + reset

        # Set command symbol
        setCS = (args, i) ->
                if args[3]
                        config[i].commandSymbol = args[3]
                        updateConfig(config)
                        console.log green + 'Command symbol updated to ' + args[3] + reset
                else console.log red + 'A new command symbol can\'t be empty' + reset
                
        # Set nickname
        setNick = (args, i) ->
                if args[3]
                        config[i].botName = args[3]
                        updateConfig(config)
                        clients[i].send 'NICK', args[3]
                        console.log green + 'Nick updated to ' + args[3] + reset
                else console.log red + 'A new nick can\'t be empty' + reset

        # Set username
        setUser = (args, i) ->
                if args[3]
                        config[i].userName = args[3]
                        updateConfig(config)
                        console.log green + 'The username ' + args[3] + ' will be used on restart' + reset
                else console.log red + 'A new user can\'t be empty' + reset

        # Set realname
        setName = (args, i) ->
                if args[3]
                        config[i].realName = args[3]
                        updateConfig(config)
                        console.log green + 'The realname ' + args[3] + ' will be used on restart' + reset
                else console.log red + 'A new name can\'t be empty' + reset

        # Set mode
        setMode = (args, i) ->
                if args.length > 5
                        config[i].users[args[4]][args[3]].mode = args[5]
                        updateConfig(config)
                else console.log red + 'Use help for a list of commands' + reset

        # Set host
        setHost = (args, i) ->
                if args.length > 5
                        config[i].users[args[4]][args[3]].host = args[5]
                        updateConfig(config)
                else console.log red + 'Use help for a list of commands' + reset

        # Get mode
        getMode = (args, i) ->
                if args.length > 4
                        console.log 'The mode for ' + args[4] + ' in ' + args[3] + ' is ' + green + config[i].users[args[4]][args[3]].mode + reset
                else console.log red + 'Use help for a list of commands' + reset
        # Get host
        getHost = (args, i) ->
                if args.length > 4
                        console.log 'The host for ' + args[4] + ' in ' + args[3] + ' is ' + green + config[i].users[args[4]][args[3]].host + reset
                else console.log red + 'Use help for a list of commands' + reset
                
# Starts with
String::startsWith = (it) ->
        @slice(0, it.length) is it

# Update config
updateConfig = (config) ->
        fs.writeFileSync './config.json', JSON.stringify config
        console.log green + 'Config file updated' + reset
        