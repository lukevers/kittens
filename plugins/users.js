// users.js
//
// Keep track of OP and Voice in channels by their nickname and their
// host name.

var util = require('util');
var fs = require('fs');

var commands = ['!op', '!deop', '!voice', '!devoice'];

module.exports = function(bot) {
    var users = require('../plugins.json')['users'];

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
		bot.say(channel, from+': The command !op requires a user[i] to be specified. By !op\'ing a user[i], the bot will remember to op them every time they sign in.');
	    } else {
		if (isOP == '+o') op(from, message, message.args[1].split(' '), channel);
		else bot.say(channel, from+': you do not have permission to do that!');
	    }
	} else if (message.args[1].indexOf('!deop') == 0) {
	    if (message.args[1].replace(/ /g, '') == '!deop') {
		bot.say(channel, from+': The command !deop requires a user[i] to be specified. By !deoping\'ing a user[i], the bot will not remember to op them every time they sign in anymore.');
	    } else {
		if (isOP == '+o') deop(from, message, message.args[1].split(' '), channel);
		else bot.say(channel, from+': you do not have permission to do that!');
	    }
	} else	if (message.args[1].indexOf('!voice') == 0) {
	    if (message.args[1].replace(/ /g, '') == '!voice') {
		bot.say(channel, from+': The command !voice requires a user[i] to be specified. By !voice\'ing a user[i], the bot will remember to voice them every time they sign in.');
	    } else {
		if (isOP == '+o') voice(from, message, message.args[1].split(' '), channel);
		else bot.say(channel, from+': you do not have permission to do that!');
	    }
	} else if (message.args[1].indexOf('!devoice') == 0) {
	    if (message.args[1].replace(/ /g, '') == '!devoice') {
		bot.say(channel, from+': The command !devoice requires a user[i] to be specified. By !devoice\'ing a user[i], the bot will not remember to voice them every time they sign in anymore.');
	    } else {
		if (isOP == '+o') devoice(from, message, message.args[1].split(' '), channel);
		else bot.say(channel, from+': you do not have permission to do that!');
	    }
	} else return;
    });

    function op(from, message, user, channel) {
	for (var i = 1; i < user.length-1; i++) {
	    if (typeof users[user[i]] == 'undefined') {
		users[user[i]] = {};
		bot.whois(user[i], function(info) {
		    users[user[i]][channel] = {'mode':'+o', 'host':info.user[i]+'@'+info.host};
		    bot.send(':'+user[i]+'!'+info.user[i]+'@'+info.host, 'MODE', channel, '+o', user[i]);
		    writeFile(users);
		});
	    } else if (typeof users[user[i]][channel] == 'undefined') {
		bot.whois(user[i], function(info) {
		    users[user[i]][channel] = {'mode':'+o', 'host':info.user[i]+'@'+info.host};
		    bot.send(':'+user[i]+'!'+info.user[i]+'@'+info.host, 'MODE', channel, '+o', user[i]);
		    writeFile(users);
		});
	    } else {
		if (users[user[i]][channel].mode == '+o') {
		    bot.send(':'+user[i]+'!'+users[user[i]][channel].host, 'MODE', channel, '+o', user[i]);
		    writeFile(users);
		} else {
		    users[user[i]][channel].mode = '+o';
		    bot.send(':'+user[i]+'!'+users[user[i]][channel].host, 'MODE', channel, '+o', user[i]);
		    writeFile(users);
		}
	    }
	}
    }
    
    function deop(from, message, user, channel) {
	for (var i = 1; i < user.length-1; i++) {
	    if (typeof users[user[i]] == 'undefined') {
		bot.say(channel, from+': '+user[i]+' already does not have mode +o!');
		return;
	    }
	    if (typeof users[user[i]][channel] == 'undefined') {
		bot.say(channel, from+': '+user[i]+' already does not have mode +o!');
		return;
	    }
	    if (users[user[i]][channel].mode == '+o') {
		bot.send(':'+user[i]+'!'+users[user[i]][channel].host, 'MODE', channel, '-o', user[i]);
		delete users[user[i]][channel];
	    writeFile(users);
	    } else {
		bot.say(channel, from+': '+user[i]+' already does not have mode +o!');
	    }
	}
    }

    function voice(from, message, user, channel) {
	for (var i = 1; i < user.length-1; i++) {
	    if (typeof users[user[i]] == 'undefined') {
		users[user[i]] = {};
		bot.whois(user[i], function(info) {
		    users[user[i]][channel] = {'mode':'+v', 'host':info.user[i]+'@'+info.host};
		    bot.send(':'+user[i]+'!'+info.user[i]+'@'+info.host, 'MODE', channel, '+v', user[i]);
		    writeFile(users);
		});
	    } else if (typeof users[user[i]][channel] == 'undefined') {
		bot.whois(user[i], function(info) {
		    users[user[i]][channel] = {'mode':'+v', 'host':info.user[i]+'@'+info.host};
		    bot.send(':'+user[i]+'!'+info.user[i]+'@'+info.host, 'MODE', channel, '+v', user[i]);
		    writeFile(users);
		});
	    } else {
		if (users[user[i]][channel].mode == '+v') {
		    bot.send(':'+user[i]+'!'+users[user[i]][channel].host, 'MODE', channel, '+v', user[i]);
		    writeFile(users);
		} else {
		    users[user[i]][channel].mode = '+v';
		    bot.send(':'+user[i]+'!'+users[user[i]][channel].host, 'MODE', channel, '+v', user[i]);
		    writeFile(users);
		}
	    }
	}
    }

    function devoice(from, message, user, channel) {
	for (var i = 1; i < user.length-1; i++) {
	    if (typeof users[user[i]] == 'undefined') {
		bot.say(channel, from+': '+user[i]+' already does not have mode +v!');
		return;
	    }
	    if (typeof users[user[i]][channel] == 'undefined') {
		bot.say(channel, from+': '+user[i]+' already does not have mode +v!');
		return;
	    }
	    if (users[user[i]][channel].mode == '+v') {
		bot.send(':'+user[i]+'!'+users[user[i]][channel].host, 'MODE', channel, '-v', user[i]);
		delete users[user[i]][channel];
		writeFile(users);
	    } else {
		bot.say(channel, from+': '+user[i]+' already does not have mode +v!');
	    }
	}
    }

    return commands;
}

function writeFile(users) {
    var file = require('../plugins.json');
    file['users'] = users;
    fs.writeFile('./plugins.json', JSON.stringify(file, null, 4), function(err) {
	if(err) {
	    util.log(err);
	} else {
	    util.log('The plugins.json file was updated!');
	}
    }); 
}
