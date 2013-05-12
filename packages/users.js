//
//	users.js
//	
//	Modular package for Kittens for the use of 
//	reading/writing to users.json.
//
//	users.js requires the use of the users.json
//	configuration file.
//	
//	Written by Luke Evers.
//

var util = require('util');

exports.writeFile = function writeFile(users) {
	fs.writeFile('../users.json', JSON.stringify(users), function(err) {
		if(err) {
			util.log(err);
		} else {
			util.log('The users.json file was updated!');
		}
	}); 
}

exports.readFile = function readFile() {
	return require('../users.json');
}
