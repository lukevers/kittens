// learn.js
//
// Learn/unlearn things! 
// Syntax:
//
// !learn "x" "y" [global] [overwrite]
// !unlearn "y"

var util = require('util');
var fs = require('fs');

var commands = ['!learn', '!unlearn'];

module.exports = function(bot) {
    var book = require('../plugins.json')['book'];
    
    bot.addListener('message', function(from, to, text, message) {
	var channel = message.args[0];
	var msg = message.args[1];
		
	if (msg.indexOf('!learn') == 0) {
	    var overwrite = (msg.indexOf('overwrite') > -1);
	    var global = (msg.indexOf('global') > -1);
	    msg = msg.replace(/\\"/g, "'");
	    var learn = msg.split('"');

	    // Delete un-used stuff
	    delete learn[0];
	    for (var i = 1; i < learn.length; i++) {
		if (learn[i].indexOf('overwrite') > -1 || 
		    learn[i].indexOf('global') > -1 ||
		    learn[i] == " as" || learn[i] == " as " ||
		    learn[i] == "as")
		    delete learn[i]
	    }

	    var remember = learn[1];
	    var as;
	    for (var i = 2; i < learn.length; i++) {
		if (learn[i] == " " || learn[i] == ""
		   || typeof learn[i] == 'undefined') continue;
		else as = learn[i];
	    }

	    if (typeof remember == 'undefined' ||
		typeof as == 'undefined') {
		bot.say(channel, from+': !learn example: !learn "kittens are cool" as "cool" [global] [overwrite]');
	    } else {
		if (global){
		    if (typeof book['global'] == 'undefined') {
			book['global'] = {};
		    }
		    if (typeof book['global'][as] == 'undefined' || overwrite) {
		 	book['global'][as] = remember;
			bot.say(channel, 'Okay, I will (globally) remember '+remember+' as '+as);
			writeFile(book);
		    } else {
			bot.say(channel, 'I already know '+as+' as something else! Use overwrite if you want to overwrite it.');
		    }
		} else {
		    if (typeof book[channel] == 'undefined') {
			book[channel] = {};
		    }
		    if (typeof book[channel][as] == 'undefined' || overwrite) {
			bot.say(channel, 'Okay, I will remember '+remember+' as '+as);
			book[channel][as] = remember;
			writeFile(book);
		    } else {
			bot.say(channel, 'I already know '+as+' as something else! Use overwrite if you want to overwrite it.');
		    }
		}
	    }
	} else if (msg.indexOf('!unlearn') == 0) {
	    var learned = msg.split(' ');
	    var global = (msg.indexOf('global') > -1);
	    
	    if (global) {
		if (typeof book['global'][learned[1]] == 'undefined') {
		    bot.say(channel, 'I don\'t know '+learned[1]+' as anything globally!');
		} else {
		    delete book['global'][learned[1]];
		    bot.say(channel, 'Okay, I will forget '+learned[1]);
		    writeFile(book);
		}
	    } else {
		if (typeof book[channel][learned[1]] == 'undefined') {
		    bot.say(channel, 'I don\'t know '+learned[1]+' as anything!');
		} else {
		    delete book[channel][learned[1]];
		    bot.say(channel, 'Okay, I will forget '+learned[1]);
		    writeFile(book);
		}
	    }
	} else {
	    var check = msg.split(' ')[0];
	    if (typeof book[channel][check] != 'undefined') {
		bot.say(channel, book[channel][check]);
	    } else if (typeof book['global'][check] != 'undefined') {
		bot.say(channel, book['global'][check]);
	    }
	}
    });
    
    
    return commands;
}

function writeFile(book) {
    var file = require('../plugins.json');
    file['book'] = book;
    fs.writeFile('./plugins.json', JSON.stringify(file, null, 4), function(err) {
	if(err) {
	    util.log(err);
	} else {
	    util.log('The plugins.json file was updated!');
	}
    });
}