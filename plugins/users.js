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
		bot.say(channel, from+': The command !op requires a thisUser to be specified. By !op\'ing a thisUser, the bot will remember to op them every time they sign in.');
	    } else {
		if (isOP == '+o') add(from, message, message.args[1].split(' '), channel, '+o');
		else bot.say(channel, from+': you do not have permission to do that!');
	    }
	} else if (message.args[1].indexOf('!deop') == 0) {
	    if (message.args[1].replace(/ /g, '') == '!deop') {
		bot.say(channel, from+': The command !deop requires a thisUser to be specified. By !deoping\'ing a thisUser, the bot will not remember to op them every time they sign in anymore.');
	    } else {
		if (isOP == '+o') remove(from, message, message.args[1].split(' '), channel, '-o');
		else bot.say(channel, from+': you do not have permission to do that!');
	    }
	} else	if (message.args[1].indexOf('!voice') == 0) {
	    if (message.args[1].replace(/ /g, '') == '!voice') {
		bot.say(channel, from+': The command !voice requires a thisUser to be specified. By !voice\'ing a thisUser, the bot will remember to voice them every time they sign in.');
	    } else {
		if (isOP == '+o') add(from, message, message.args[1].split(' '), channel, '+v');
		else bot.say(channel, from+': you do not have permission to do that!');
	    }
	} else if (message.args[1].indexOf('!devoice') == 0) {
	    if (message.args[1].replace(/ /g, '') == '!devoice') {
		bot.say(channel, from+': The command !devoice requires a thisUser to be specified. By !devoice\'ing a thisUser, the bot will not remember to voice them every time they sign in anymore.');
	    } else {
		if (isOP == '+o') remove(from, message, message.args[1].split(' '), channel, '-v');
		else bot.say(channel, from+': you do not have permission to do that!');
	    }
	} else return;
    });
    
    function add(from, message, user, channel, which) {
    	for (var i = 1; i < user.length; i++) {
    	    util.log(which+" for "+user[i]);
    	    var thisUser = user[i];
    	    if (typeof users[thisUser] == 'undefined') {
	        users[thisUser] = {};
	        bot.whois(thisUser, function(info) {
		    users[thisUser][channel] = {'mode':which, 'host':info.user+'@'+info.host};
		    bot.send(':'+thisUser+'!'+info.user+'@'+info.host, 'MODE', channel, which, thisUser);
		    writeFile(users);
	 	});
    	    } else if (typeof users[thisUser][channel] == 'undefined') {
		bot.whois(thisUser, function(info) {
		    users[thisUser][channel] = {'mode':which, 'host':info.user+'@'+info.host};
		    bot.send(':'+thisUser+'!'+info.user+'@'+info.host, 'MODE', channel, which, thisUser);
		    writeFile(users);
		});
	    } else {
		if (users[thisUser][channel].mode == which) {
		    bot.send(':'+thisUser+'!'+users[thisUser][channel].host, 'MODE', channel, which, thisUser);
		    writeFile(users);
		} else {
		    users[thisUser][channel].mode = which;
		    bot.send(':'+thisUser+'!'+users[thisUser][channel].host, 'MODE', channel, which, thisUser);
		    writeFile(users);
		}
	    }
    	}
    }

    function remove(from, message, user, channel, which) {
	for (var i = 1; i < user.length; i++) {
	    var thisUser = user[i];
	    if (typeof users[thisUser] == 'undefined') {
		bot.say(channel, from+': '+thisUser+' already does not have mode!');
		continue;
	    }
	    if (typeof users[thisUser][channel] == 'undefined') {
		bot.say(channel, from+': '+thisUser+' already does not have mode!');
		continue;
	    }
	    if (users[thisUser][channel].mode == which) {
		bot.send(':'+thisUser+'!'+users[thisUser][channel].host, 'MODE', channel, which, thisUser);
		delete users[thisUser][channel];
	        writeFile(users);
	    } else {
		bot.say(channel, from+': '+thisUser+' already does not have mode!');
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
