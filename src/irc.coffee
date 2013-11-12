module.exports = (config) ->
        
        ####################
        ### LOAD MODULES ###
        ####################
        
        irc = require 'irc'
        
        #############
        ### START ###
        #############
        
        # Create client
        client = new irc.Client config.server, config.botName, config
        console.log 'Connecting to ' + config.server
        
        # Listen for connection just for logging purposes 
        client.addListener 'registered', (message) ->
                console.log 'Connected to ' + config.server
        
        # Add error handling
        client.addListener 'error', (message) ->
                console.log 'error: ' + message

        ###############
        ### LOGGING ###
        ###############

        require('./logging') client