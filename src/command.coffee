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

        parseCommand = (command) ->
                console.log command