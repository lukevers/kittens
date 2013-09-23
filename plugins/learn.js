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