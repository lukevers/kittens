//
//	Kittens
//  GitHub.js
//	
//	Modular plugin for Kittens for the use of 
//	fetching HEAD <title> data about issues
//	
//	Written by Luke Evers.
//

var util = require('util');
var fs = require('fs');
var request = require('request');

var commands = ['!setRepo'];

module.exports = function(bot) {
	
	var users = readFile('users');
	var repos = readFile('repos');
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
			var owner = repos[channel];
			owner = owner.substring('https://github.com/'.length);
			var url = 'https://api.github.com/repos/'+owner+'/issues/'+number;
			url += '?client_id='+conf.client_id+'&client_secret='+conf.client_secret;
			if (!isNaN(parseInt(number))) postLink(bot, url, channel, parseInt(number));
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
		if ((typeof msg != 'undefined') && msg.indexOf('API Rate Limit Exceeded') > -1) {
			bot.say(channel, 'API Rate Limit Exceeded for the hour.');
		} else bot.say(channel, '#'+number+' - \u0002'+str+'\u000f');
	});
}

function writeFile(config) {
	fs.writeFile('./repos.json', JSON.stringify(config), function(err) {
		if(err) {
			util.log(err);
		} else {
			util.log('The repos.json file was updated!');
		}
	}); 
}
	
function readFile(which) {
	return require('../'+which+'.json');
}