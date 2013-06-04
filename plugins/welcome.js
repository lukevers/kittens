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
	var file = require('../welcome.json');
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
		if (JSON.stringify(file.old).indexOf(nick) == -1) {			
			var msg = parseMessage(JSON.stringify(file.message), channel, nick);
			bot.say(channel, msg);
			util.log(msg);
		}
	});
	
	// will do in a bit
	function setWelcomeMessage(channel, message) {
		bot.say(channel, message);
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
	return message.substring(1, message.length-1).replace(/{channel}/g, channel).replace(/{nick}/g, nick);
}