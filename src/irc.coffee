module.exports = (config) ->
        
        ####################
        ### LOAD MODULES ###
        ####################
        
        irc = require 'irc'


        for i in [0..config.length-1] by 1
                client = new irc.Client config[i].server, config[i].botName, config[i]
                console.log 'Connecting to ' + config[i].server

                client.addListener 'registered', (message) ->
                        console.log 'Connected to ' + message.server

                client.addListener 'error', (message) ->
                        console.log 'error from ' + message.server + ': ' + JSON.stringify message

                require('./logging') client, server