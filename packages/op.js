//
//	op.js
//	
//	Modular package for Kittens for the use of 
//	OPing and deOPing people.
//
//	op.js requires the use of the users.js
//	package.
//	
//	Written by Luke Evers.
//

var users_j = require('./users.js');

function op(from, message, user) {
	var users = users_j.readFile();
	if (typeof users[user] == 'undefined') {
		bot.whois(user, function(info) {
			users[user] = {'mode':'+o', 'host':info.user+'@'+info.host};
			bot.send(':'+user+'!'+info.user+'@'+info.host, 'MODE', message.args[0], '+o', user);
			users_j.writeFile(users);
			});
		} 
	else {		
	if (users[user].mode == '+o') {
		bot.say(message.args[0], from+': '+user+' already has mode +o!');
	} else {
		users[user].mode = '+o';
		bot.send(':'+user+'!'+c.users[user].host, 'MODE', message.args[0], '+o', user);
		writeFile(users);
		}
	}
}

function deop(from, message, user) {
	if (c.users[user].mode == '+o') {
		bot.send(':'+user+'!'+c.users[user].host, 'MODE', message.args[0], '-o', user);
		delete c.users[user];
		writeFile(JSON.stringify(c.users));
	} else {
		bot.say(message.args[0], from+': '+user+' already does not have mode +o!');
	}
}