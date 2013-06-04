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

var commands = ['+voice', '+devoice'];

module.exports = function(bot) {
		
	var users = readFile();
	bot.addListener('message', function(from, to, text, message) {
		if (typeof users[from] == 'undefined') {
			isMaster = false;
		} else isMaster = users[from].master;

		if (message.args[1].indexOf('+voice') == 0) {
			if (message.args[1].replace(/ /g, '') == '+voice') {
				bot.say(message.args[0], from+': The command +voice requires a user to be specified. By +voice\'ing a user, the bot will remember to voice them every time they sign in.');
			} else {
				if (isMaster) voice(from, message, message.args[1].split(' ')[1]);
				else bot.say(message.args[0], from+': you do not have permission to do that!');
			}
		} else if (message.args[1].indexOf('+devoice') == 0) {
			if (message.args[1].replace(/ /g, '') == '+devoice') {
				bot.say(message.args[0], from+': The command +devoice requires a user to be specified. By +devoice\'ing a user, the bot will not remember to voice them every time they sign in anymore.');
			} else {
				if (isMaster) devoice(from, message, message.args[1].split(' ')[1]);
				else bot.say(message.args[0], from+': you do not have permission to do that!');
			}
		} else return;
	});
	
	bot.addListener('join', function(channel, nick, message) {
		util.log(nick+' joined '+channel);
		var file = readFile();
		var userinfo = file[[nick]];
		if (typeof userinfo == 'undefined') return;
		var userhost = message.user+'@'+message.host;
		if (userinfo.host == userhost && userinfo.mode == '+v') {
			bot.send(':'+nick+'!'+userhost, 'MODE', channel, userinfo.mode, nick);
			util.log(userinfo.mode+' '+nick+' in '+channel);
		}
	});
	
	function voice(from, message, user) {
		if (typeof users[user] == 'undefined') {
			bot.whois(user, function(info) {
				users[user] = {'mode':'+v', 'host':info.user+'@'+info.host};
				bot.send(':'+user+'!'+info.user+'@'+info.host, 'MODE', message.args[0], '+v', user);
				writeFile(users);
			});
		} else {
			if (users[user].mode == '+o') {
				bot.say(message.args[0], from+': '+user+' already has mode +o!');
			} else if (users[user].mode == '+v') {
				bot.say(message.args[0], from+': '+user+' already has mode +v!');
			} else {
				users[user].mode = '+v';
				bot.send(':'+user+'!'+users[user].host, 'MODE', message.args[0], '+v', user);
				writeFile(users);
			}
		}
	}
	
	function devoice(from, message, user) {
		if (typeof users[user] == 'undefined') {
			bot.say(message.args[0], from+': '+user+' already does not have mode +v!');
			return;
		}
		if (users[user].mode == '+v') {
			bot.send(':'+user+'!'+users[user].host, 'MODE', message.args[0], '-v', user);
			delete users[user];
			writeFile(users);
		} else {
			bot.say(message.args[0], from+': '+user+' already does not have mode +v!');
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