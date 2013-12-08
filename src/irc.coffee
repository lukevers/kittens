##############
### COLORS ###
##############

green = `'\033[0;32m'`
reset = `'\033[0m'`
red = `'\033[0;31m'`

##############
### MODULE ###
##############

module.exports = (config) ->
        
        ####################
        ### LOAD MODULES ###
        ####################
        
        irc = require 'irc'
        clients = []

        for i in [0..config.length-1] by 1
                client = new irc.Client config[i].server, config[i].botName, config[i]
                console.log 'Connecting to ' + config[i].server

                client.addListener 'registered', (message) ->
                        console.log 'Connected to ' + message.server

                client.addListener 'error', (message) ->
                        console.log red + 'error from ' + message.server + ': ' + JSON.stringify(message) + reset

                require('./logging') client, config[i].server
                require('./plugin') client, config[i]
                clients[i] = client

        require('./command') clients, config
        