####################
### LOAD MODULES ###
####################

fs = require 'fs'

##############
### MODULE ###
##############

cmds = ['  help']

module.exports = (client, config, n) ->
        server = config[n]
        console.log 'Loading plugins for ' + server.serverName
        fs.readdir './lib/plugins', (err, files) ->
                for i in [0..files.length-1] by 1
                        console.log 'Loading plugin: ' + files[i].substring 0, files[i].length - 3
                        c = require('./plugins/' + files[i])(client, config, n)
                        cmds.push c unless typeof c is 'undefined'

        # Add listener for `help` command
        client.addListener 'message', (from, to, text, message) ->
                if message.args[1].split(' ')[0] == server.commandSymbol + 'help'
                        client.say message.args[0], from + ':' + cmds.unique().join(' ' + server.commandSymbol).replace(/,|  /g, ' ' + server.commandSymbol)

Array::unique = ->
        o = {}
        i = undefined
        l = @length
        r = []
        i = 0
        while i < l
                o[this[i]] = this[i]
                i += 1
        for i of o
                r.push o[i]
        r
