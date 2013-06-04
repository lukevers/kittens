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

var commands = ['+op', '+deop'];

module.exports = function(bot) {
	
	var users = readFile();
	bot.addListener('message', function(from, to, text, message) {
		if (typeof users[from] == 'undefined') {
			isMaster = false;
		} else isMaster = users[from].master;
		if (message.args[1].indexOf('+op') == 0) {
			if (message.args[1].replace(/ /g, '') == '+op') {
				bot.say(message.args[0], from+': The command +op requires a user to be specified. By +op\'ing a user, the bot will remember to op them every time they sign in.');
			} else {
				if (isMaster) op(from, message, message.args[1].split(' ')[1]);
				else bot.say(message.args[0], from+': you do not have permission to do that!');
			}
		} else if (message.args[1].indexOf('+deop') == 0) {
			if (message.args[1].replace(/ /g, '') == '+deop') {
				// HELP
				bot.say(message.args[0], from+': The command +deop requires a user to be specified. By +deoping\'ing a user, the bot will not remember to op them every time they sign in anymore.');
			} else {
				if (isMaster) deop(from, message, message.args[1].split(' ')[1]);
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
		if (userinfo.host == userhost && userinfo.mode == '+o') {
			bot.send(':'+nick+'!'+userhost, 'MODE', channel, userinfo.mode, nick);
			util.log(userinfo.mode+' '+nick+' in '+channel);
		}
	});
	
	function op(from, message, user) {
		if (typeof users[user] == 'undefined') {
			bot.whois(user, function(info) {
				users[user] = {'mode':'+o', 'host':info.user+'@'+info.host};
				bot.send(':'+user+'!'+info.user+'@'+info.host, 'MODE', message.args[0], '+o', user);
				writeFile(users);
				});
			} 
		else {		
		if (users[user].mode == '+o') {
			bot.say(message.args[0], from+': '+user+' already has mode +o!');
		} else {
			users[user].mode = '+o';
			bot.send(':'+user+'!'+users[user].host, 'MODE', message.args[0], '+o', user);
			writeFile(users);
			}
		}
	}
	
	function deop(from, message, user) {
		if (typeof users[user] == 'undefined') {
			bot.say(message.args[0], from+': '+user+' already does not have mode +o!');
			return;
		}
		if (users[user].mode == '+o') {
			bot.send(':'+user+'!'+users[user].host, 'MODE', message.args[0], '-o', user);
			delete users[user];
			writeFile(users);
		} else {
			bot.say(message.args[0], from+': '+user+' already does not have mode +o!');
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
	return require('./users.json');
}