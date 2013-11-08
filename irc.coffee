# Load  modules
fs = require 'fs'

# Load the configuration file
Config = if fs.existsSync './config.json' then require './config.json' else require './example.json'

###########
### WEB ###
###########

# Load modules for web server
http = require 'http'
express = require 'express'

# Configure the web server
app = express()
app.use 'static', express.static(__dirname + '/static')
app.set 'views', __dirname + '/views'
app.set 'view engine', 'coffee'
app.engine 'coffee', require('coffeecup').__express

# Load the router for the web server
require('./lib/router') app

# Serve the web server
http.createServer(app).listen(Config.Web.Port, '::')

###########
### IRC ###
###########

# Load modules for IRC
irc = require 'irc'

# Create client
client = new irc.Client Config.IRC.server, Config.IRC.botName, Config.IRC

# Add error handling
client.addListener 'error', (message) ->
        console.log 'error: ', message