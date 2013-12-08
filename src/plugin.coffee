####################
### LOAD MODULES ###
####################

fs = require 'fs'

##############
### MODULE ###
##############

cmds = ['help']

module.exports = (client, server) ->
        console.log 'Loading plugins for ' + server.server
        fs.readdir './lib/plugins', (err, files) ->
                for i in [0..files.length-1] by 1
                        console.log 'Loading plugin: ' + files[i].substring 0, files[i].length - 3
                        c = require('./plugins/' + files[i])(client)
                        cmds.push c unless typeof c is 'undefined'

        # Add listener for `help` command
        client.addListener 'message', (from, to, text, message) ->
                # TODO