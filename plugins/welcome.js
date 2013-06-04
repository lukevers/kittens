//
//	Kittens
//	welcome.js
//	
//	Modular plugin for Kittens for the use of 
//	welcoming new users to channels.
//	
//	Written by Luke Evers.
//

var util = require('util');
var fs = require('fs');

var commands = ['+setWelcomeMessage'];

module.exports = function(bot) {
	var users = require('../users.json');
	// If 'welcome.json' exists, load it
	// Else we need to create it, then load it
	fs.exists('./welcome.json', function(exists) {
		if (!exists) {
			fs.writeFile('./welcome.json', function(err) {
				if (err) util.log(err);
				else util.log('welcome.json was created.');
			});
		}
	});
	var file = require('./welcome.json');
	
	bot.addListener('message', function(from, to, text, message) {
		if (typeof users[from] == 'undefined') {
			isOP = false;
		} else isOP = (users[from].mode == '+o');
		if (message.args[1].indexOf('+setWelcomeMessage') == 0) {
			if (isOP) setWelcomeMessage(message.args[0], message.args[1].substring(19));
			else bot.say(message.args[0], from+': you do not have permission to do that!');
		}
	});
	
	bot.addListener('join', function(channel, nick, message) {
		// Check to see if 
		
		
	});
	
	// will do in a bit
	function setWelcomeMessage(channel, message) {
		bot.say(channel, message);
	}
	
	return commands;
}