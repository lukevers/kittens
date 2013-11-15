module.exports = (clients, config) ->

        ##############
        ### COLORS ###
        ##############

        green = `'\033[0;32m'`
        reset = `'\033[0m'`
        red   = `'\033[0;31m'`

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
        # - disconnect [server]
        #
        # - connect all
        # - connect [server]
        #
        # - cycle all
        # - cycle [server] all
        # - cycle [server] [channel]
        # 
        # - join [server] [channel]
        # - part [server] [channel]
        # - whois [server] [nick]
        #
        # - say [server] [channel] [message]
        #
        # - set [server] server [server]
        # - set [server] port [port]
        # - set [server] commandSymbol [symbol]
        # - set [server] nick [nick]
        # - set [server] user [user]
        # - set [server] name [name]
        #

        parseCommand = (command) ->
                command = command.replace /\r?\n|\r/g, ''
                command = command.toLowerCase()
                args = command.split ' '
                switch args[0]
                        when 'help' then help(args); break
                        when 'quit' then quit(args); break
                        when 'disconnect' then disconnect(args); break
                        when 'connect' then connect(args); break
                        when 'cycle' then cycle(args); break
                        when 'join' then join(args); break
                        when 'part' then part(args); break
                        when 'say' then say(args); break 
                        when 'set' then set(args); break
                        else console.log red + 'Use help for a list of commands' + reset

        help = (args) ->
                console.log 'help'

        quit = (args) ->
                args[1] = 'all'
                disconnect(args)

        disconnect = (args) ->
                if args[1] is 'all'
                        for i in [0..clients.length-1] by 1
                                console.log 'Disconnecting from ' + config[i].server
                                clients[i].disconnect 'disconnecting'
                        process.exit 0 if args[0] is 'quit'
                else
                        for i in [0..clients.length-1] by 1
                                if args[1] is config[i].server
                                        console.log 'Disconnecting from ' + config[i].server
                                        clients[i].disconnect 'disconnecting'
                                        return
                        console.log red + 'Server does not exist' + reset

        connect = (args) ->
                if args[1] is 'all'
                        for i in [0..clients.length-1] by 1
                                console.log 'Connecting to ' + config[i].server
                                clients[i].connect
                else
                        for i in [0..clients.length-1] by 1
                                if args[1] is config[i].server
                                        console.log 'Connecting to ' + config[i].server
                                        clients[i].connect
                                        return
                        console.log red + 'Server does not exist' + reset

        cycle = (args) ->
                console.log 'cycle'

        join = (args) ->
                console.log 'join'

        part = (args) ->
                console.log 'part'

        say = (args) ->
                console.log 'say'

        set = (args) ->
                console.log 'set'

String::startsWith = (it) ->
        @slice(0, it.length) is it