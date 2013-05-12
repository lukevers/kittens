//
//	Kittens
//	irc.js
//	
//	A modular IRC bot written in Node.js.
//	
//	Written by Luke Evers.
//

var util = require("util");
var irc = require("irc");
var request = require("request");
var c = require("./config");
var fs = require('fs');

util.log("Configured "+c.config.botName);
util.log("Creating "+c.config.botName);

var bot = new irc.Client(c.config.server, c.config.botName, {
	channels: c.config.channels
});

util.log("Created "+c.config.botName);
util.log("Connecting to "+c.config.server);

fs.readdir('./packages', function(err, files) {
	for (var i = 0; i < files.length; i++) {
		
		//
		// TODO: create a JSON list of packages people want and then
		//       check it against each one
		//
		
		util.log('Loading package: '+files[i].substring(0, files[i].length-3));
		require('./packages/'+files[i])(bot);
	}
});
