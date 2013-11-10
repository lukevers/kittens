# Load modules
fs = require 'fs'

module.exports = (client) ->

        # Create the log directory if it does not already exist
        fs.mkdirSync './logging' if !fs.existsSync './logging'

        # Listen for normal messages, and log them
        client.addListener 'message', (from, to, text, message) ->
                channel = message.args[0]
                msg = '<' + from + '> ' + message.args[1]
                writeFile channel, msg

        # Listen for join events
        client.addListener 'join', (channel, nick, message) ->
                msg = nick + ' joined ' + channel
                writeFile channel, msg

        # Listen for topic changes
        client.addListener 'topic', (channel, topic, nick, message) ->
                msg = 'The topic for ' + channel + ' is ' + topic
                writeFile channel, msg

        # Listen for part events
        client.addListener 'part', (channel, nick, reason, message) ->
                msg = nick + ' has quit (' + reason + ')'
                writeFile channel, msg

        # Listen for kick events
        client.addListener 'kick', (channel, nick, byy, reason, message) ->
                msg = nick + ' was kicked by ' + byy + ' (' + reason + ')'
                writeFile channel, msg

# Initialze channel log file if it does not already exist
initChannel = (channel) ->
        fs.writeFileSync './logging/' + channel, '' if !fs.existsSync './logging/' + channel

# Write to file
writeFile = (channel, message) ->
        initChannel channel
        fs.appendFileSync './logging/' + channel, message.timestamp() + '\n'

# Append timestamp to message
String::timestamp = ->
        time = new Date().toTimeString().replace(/.*(\d{2}:\d{2}:\d{2}).*/, "$1")
        time = '[' + time + '] ' + this;
        time