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
        # - quit
        # - join [channel]
        # - part [channel]
        # - whois [nick]
        #
        # - say [channel] [message]
        #
        # - set nick [nick]
        # - set user [user]
        # - set name [name]
        # 
        parseCommand = (command) ->
                args = command.split ' '
                

String::startsWith = (it) ->
        @slice(0, it.length) is it