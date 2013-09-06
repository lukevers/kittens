// GitHub.js
//
// Fetches issue information from GitHub when a user types an issue
// number with a # in front of the number.

var util = require('util');
var fs = require('fs');
var request = require('request');

var commands = ['!setRepo'];

module.exports = function(bot) {
	
	var users = readFile('plugins')['users'];
	var repos = readFile('plugins')['repos'];
	var conf = readFile('config');
	
	bot.addListener('message', function(from, to, text, message) {
		var channel = message.args[0];
		if (message.args[1].indexOf('!setRepo') == 0) {
			if (typeof users[from][channel] == 'undefined') {
				isOP = false;
			} else isOP = users[from][channel].mode;
			if (message.args[1].replace(/ /g, '') == '!setRepo') {
				bot.say(channel, from+': The command !setRepo requires a repo to be specified.');
			} else {
				repos[channel] = message.args[1].substring(9);
				writeFile(repos);
			}
		}
		if (message.args[1].indexOf('#') > -1) {
			var number = message.args[1].substring(message.args[1].indexOf('#'));
			number = number.substring(1, number.indexOf(' ', 0));
			if (number == '#') {
				number = message.args[1].substring(message.args[1].indexOf('#')+1);
			}
			if (!(typeof repos[channel] == 'undefined')) {
				var owner = repos[channel];
				owner = owner.substring('https://github.com/'.length);
				var url = 'https://api.github.com/repos/'+owner+'/issues/'+number;
				url += '?client_id='+conf.client_id+'&client_secret='+conf.client_secret;
				if (!isNaN(parseInt(number))) postLink(bot, url, channel, parseInt(number));
			}
		}
	});
	
	return commands;
}

function postLink(bot, url, channel, number) {	
	request({
		uri: url,
	}, function(err, res, b) {
		b = JSON.parse(b);
		str = b["title"];
		msg = b["message"];
		state = b["state"];
		if (state == "closed") {
			state = ' [\u0002'+state+'\u000f] ';
		} else state = ' ';
		if ((typeof msg != 'undefined') && msg.indexOf('API Rate Limit Exceeded') > -1) {
			bot.say(channel, 'API Rate Limit Exceeded for the hour.');
		} else bot.say(channel, '#'+number+state+'- \u0002'+str+'\u000f');
	});
}

function writeFile(config) {
	var file = require('../plugins.json');
	file['repos'] = config;
	fs.writeFile('./plugins.json', JSON.stringify(file, null, 4), function(err) {
		if(err) {
			util.log(err);
		} else {
			util.log('The plugins.json file was updated!');
		}
	}); 
}
	
function readFile(which) {
	return require('../'+which+'.json');
}