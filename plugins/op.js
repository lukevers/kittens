//
//	Kittens
//	op.js
//	
//	Modular plugin for Kittens for the use of 
//	OPing and deOPing people.
//	
//	Written by Luke Evers.
//

var util = require('util');
var fs = require('fs');

var commands = ['!op', '!deop'];

module.exports = function(bot) {
	
	var users = readFile();
	bot.addListener('message', function(from, to, text, message) {
		var channel = message.args[0];
		if (typeof users[from] == 'undefined') {
			users[from] = {};	
		}
		if (typeof users[from][channel] == 'undefined') {
			isOP = false;
		} else isOP = users[from][channel].mode;
		
		if (message.args[1].indexOf('!op') == 0) {
			if (message.args[1].replace(/ /g, '') == '!op') {
				bot.say(channel, from+': The command !op requires a user to be specified. By !op\'ing a user, the bot will remember to op them every time they sign in.');
			} else {
				if (isOP) op(from, message, message.args[1].split(' ')[1], channel);
				else bot.say(channel, from+': you do not have permission to do that!');
			}
		} else if (message.args[1].indexOf('!deop') == 0) {
			if (message.args[1].replace(/ /g, '') == '!deop') {
				// HELP
				bot.say(channel, from+': The command !deop requires a user to be specified. By !deoping\'ing a user, the bot will not remember to op them every time they sign in anymore.');
			} else {
				if (isOP) deop(from, message, message.args[1].split(' ')[1], channel);
				else bot.say(channel, from+': you do not have permission to do that!');
			}
		} else return;
	});
	
	bot.addListener('join', function(channel, nick, message) {
		util.log(nick+' joined '+channel);
		var file = readFile();
		if (typeof file[[nick]] == 'undefined' || typeof file[[nick]][channel] == 'undefined') return;
		var userinfo = file[[nick]][channel];
		util.log(userinfo);
		var userhost = message.user+'@'+message.host;
		if (typeof userinfo.host == 'undefined' || typeof userinfo.mode == 'undefined') return;
		if (userinfo.host == userhost && userinfo.mode == '+o') {
			bot.send(':'+nick+'!'+userhost, 'MODE', channel, userinfo.mode, nick);
			util.log(userinfo.mode+' '+nick+' in '+channel);
		}
	});
	
	function op(from, message, user, channel) {
		if (typeof users[user] == 'undefined') {
			users[user] = {};
			bot.whois(user, function(info) {
				users[user][channel] = {'mode':'+o', 'host':info.user+'@'+info.host};
				bot.send(':'+user+'!'+info.user+'@'+info.host, 'MODE', channel, '+o', user);
				writeFile(users);
				});
		} else if (typeof users[user][channel] == 'undefined') {
			bot.whois(user, function(info) {
				users[user][channel] = {'mode':'+o', 'host':info.user+'@'+info.host};
				bot.send(':'+user+'!'+info.user+'@'+info.host, 'MODE', channel, '+o', user);
				writeFile(users);
				});
		} else {		
			if (users[user][channel].mode == '+o') {
				bot.say(channel, from+': '+user+' already has mode +o!');
			} else {
				users[user][channel].mode = '+o';
				bot.send(':'+user+'!'+users[user][channel].host, 'MODE', channel, '+o', user);
				writeFile(users);
			}
		}
	}
	
	function deop(from, message, user, channel) {
		if (typeof users[user][channel] == 'undefined') {
			bot.say(channel, from+': '+user+' already does not have mode +o!');
			return;
		}
		if (users[user][channel].mode == '+o') {
			bot.send(':'+user+'!'+users[user][channel].host, 'MODE', channel, '-o', user);
			delete users[user][channel];
			writeFile(users);
		} else {
			bot.say(channel, from+': '+user+' already does not have mode +o!');
		}
	}
	
	return commands;
}

function writeFile(users) {
	fs.writeFile('./users.json', JSON.stringify(users), function(err) {
		if(err) {
			util.log(err);
		} else {
			util.log('The users.json file was updated!');
		}
	}); 
}
	
function readFile() {
	return require('../users.json');
}