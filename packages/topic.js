//
//	Kittens
//	topic.js
//	
//	Modular package for Kittens for the use of 
//	stating the topic when it changes. This is
//	mainly for logging purposes.
//	
//	Written by Luke Evers.
//

var util = require('util');

module.exports = function(bot) {
	bot.addListener("topic", function(channel, topic, nick, message){
		util.log('The new topic on '+channel+' is "'+topic+'"');
		bot.say(channel, 'The new topic on '+channel+' is "\u0002'+topic+'\u000f"');
	});
}
