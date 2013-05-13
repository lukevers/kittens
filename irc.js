//
//	Kittens
//	irc.js
//	
//	A modular IRC bot written in Node.js.
//	
//	Written by Luke Evers.
//

var util = require('util');
var irc = require('irc');
var fs = require('fs');

util.log('Configuring bot');

var conf = require('./config.json');

var config = {
	botName: conf.botName,
	realName: conf.realName,
	autoRejoin: conf.autoRejoin,
	autoConnect: conf.autoConnect,
	channels: conf.channels,
	server: conf.server,
	port: conf.port,
	usersFile: conf.usersFile
};

util.log('Configured '+config.botName);
util.log('Creating '+config.botName);

var bot = new irc.Client(config.server, config.botName, config);

util.log('Created '+config.botName);

fs.readdir('./plugins', function(err, files) {
	util.log('Loading plugins');
	var plugin = require('./plugins.json');
	for (var i = 0; i < files.length; i++) {
		for (var key in plugin) {
			if ([[key]] == files[i].substring(0, files[i].length-3)) {
				util.log('Loading plugin: '+files[i].substring(0, files[i].length-3));
				require('./plugins/'+files[i])(bot);
			}
		}
	}
});

util.log('Connecting to '+config.server);
