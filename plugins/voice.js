//
//	Kittens
//	voice.js
//	
//	Modular plugin for Kittens for the use of 
//	Voicing and devoicing people.
//	
//	Written by Luke Evers.
//

var util = require('util');
var fs = require('fs');

var commands = ['!voice', '!devoice'];

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

		if (message.args[1].indexOf('!voice') == 0) {
			if (message.args[1].replace(/ /g, '') == '!voice') {
				bot.say(channel, from+': The command !voice requires a user to be specified. By !voice\'ing a user, the bot will remember to voice them every time they sign in.');
			} else {
				if (isOP) voice(from, message, message.args[1].split(' ')[1], channel);
				else bot.say(channel, from+': you do not have permission to do that!');
			}
		} else if (message.args[1].indexOf('!devoice') == 0) {
			if (message.args[1].replace(/ /g, '') == '!devoice') {
				bot.say(channel, from+': The command !devoice requires a user to be specified. By !devoice\'ing a user, the bot will not remember to voice them every time they sign in anymore.');
			} else {
				if (isOP) devoice(from, message, message.args[1].split(' ')[1], channel);
				else bot.say(channel, from+': you do not have permission to do that!');
			}
		} else return;
	});
	
	bot.addListener('join', function(channel, nick, message) {
		util.log(nick+' joined '+channel);
		var file = readFile();
		if (typeof file[[nick]] == 'undefined') return;
		var userinfo = file[[nick]][channel];
		var userhost = message.user+'@'+message.host;
		if (userinfo.host == userhost && userinfo.mode == '+v') {
			bot.send(':'+nick+'!'+userhost, 'MODE', channel, userinfo.mode, nick);
			util.log(userinfo.mode+' '+nick+' in '+channel);
		}
	});
	
	function voice(from, message, user, channel) {
		if (typeof users[user] == 'undefined') {
			users[user] = {};
			bot.whois(user, function(info) {
				users[user][channel] = {'mode':'+v', 'host':info.user+'@'+info.host};
				bot.send(':'+user+'!'+info.user+'@'+info.host, 'MODE', channel, '+v', user);
				writeFile(users);
			});
		}
		if (typeof users[user][channel] == 'undefined') {
			bot.whois(user, function(info) {
				users[user][channel] = {'mode':'+v', 'host':info.user+'@'+info.host};
				bot.send(':'+user+'!'+info.user+'@'+info.host, 'MODE', channel, '+v', user);
				writeFile(users);
			});
		} else {
			if (users[user][channel].mode == '+o') {
				bot.say(channel, from+': '+user+' already has mode +o!');
			} else if (users[user][channel].mode == '+v') {
				bot.say(channel, from+': '+user+' already has mode +v!');
			} else {
				users[user][channel].mode = '+v';
				bot.send(':'+user+'!'+users[user][channel].host, 'MODE', channel, '+v', user);
				writeFile(users);
			}
		}
	}
	
	function devoice(from, message, user, channel) {
		if (typeof users[user][channel] == 'undefined') {
			bot.say(channel, from+': '+user+' already does not have mode +v!');
			return;
		}
		if (users[user][channel].mode == '+v') {
			bot.send(':'+user+'!'+users[user][channel].host, 'MODE', channel, '-v', user);
			delete users[user][channel];
			writeFile(users);
		} else {
			bot.say(channel, from+': '+user+' already does not have mode +v!');
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