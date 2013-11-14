module.exports = (config) ->

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
        #
        # - disconnect all
        # - disconnect [server]
        #
        # - reconnect all
        # - reconnect [server]
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
                args = command.split ' '
                

String::startsWith = (it) ->
        @slice(0, it.length) is it