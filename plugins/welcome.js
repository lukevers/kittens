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

var commands = ['!setWelcomeMessage'];

module.exports = function(bot) {
	var users = require('../users.json');
	var file = require('../welcome.json');
	bot.addListener('message', function(from, to, text, message) {
		var channel = message.args[0];
		if (typeof users[from] == 'undefined') {
			users[from] = {};	
		}
		if (typeof users[from][channel] == 'undefined') {
			isOP = false;
		} else isOP = (users[from][channel].mode == '+o');
		
		if (message.args[1].indexOf('!setWelcomeMessage') == 0) {
			if (message.args[1].replace(/ /g, '') == '!setWelcomeMessage') {
				bot.say(channel, from+': The command !setWelcomeMessage requires a new welcome message to be specified. Every time a new user logs on, they will recieve the welcome message.');
			} else {
				if (isOP) setWelcomeMessage(channel, message.args[1].substring(19), from);
				else bot.say(channel, from+': you do not have permission to do that!');
			}
		}
	});
	
	bot.addListener('join', function(channel, nick, message) {
		if (typeof file[channel] == 'undefined') {
			file[channel] = {'old': '', 'message': 'Welcome to {channel}, {nick}!'};
		}
		if (file[channel].old.indexOf(nick) == -1) {			
			file[channel].old = file[channel].old+nick;
			var msg = parseMessage(file[channel].message, channel, nick);
			bot.say(channel, msg);
			util.log(nick+' joined '+channel+' for the first time and was given the message: '+msg);
			writeFile(file);
		}
	});
	
	// will do in a bit
	function setWelcomeMessage(channel, message, setby) {
		file[channel].message = message;
		bot.say(channel, 'The new welcome message, set by ' + setby + ', is '+message);
		writeFile(file);
	}
	
	return commands;
}
 
//  -- Parse Message --
//  
//  Replaces certain words with 
//  channel names or nick names
//  
//  {channel} -> channel name
//  {nick}    -> nick name
//  
function parseMessage(message, channel, nick) {	
	return message.replace(/{channel}/g, channel).replace(/{nick}/g, nick);
}

function writeFile(file) {
	fs.writeFile('./welcome.json', JSON.stringify(file), function(err) {
		if(err) {
			util.log(err);
		} else {
			util.log('The welcome.json file was updated!');
		}
	}); 
}