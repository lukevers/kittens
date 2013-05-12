// 
// Kittens - IRC Bot
//

var util = require("util");
var irc = require("irc");
var request = require("request");
var c = require("./config");
var fs = require("fs");

util.log("Configured "+c.config.botName);

// Create the bot.
var bot = new irc.Client(c.config.server, c.config.botName, {
	channels: c.config.channels
});

util.log("Created "+c.config.botName);
util.log("Connecting to "+c.config.server);

// Listen for topic changes on channels, and when there is a topic
// change, the bot will announce the new topic.
bot.addListener("topic", function(channel, topic, nick, message){
	util.log("The new topic on "+channel+" is \""+topic+"\"");
	bot.say(channel, "The new topic on "+channel+" is \"\u0002"+topic+"\u000f\"");
});

// Listen for for joins to the channel so that the relevant people can
// be autooped or autovoiced.
bot.addListener("join", function(channel, nick, message){
	util.log(nick+" joined "+channel);

	// Use the 'users' map to apply the appropriate mode, if
	// applicable.
	userinfo = c.users[[nick]]
	if (typeof userinfo == "undefined") {
		return
	}
	
	userhost = message.user+"@"+message.host
	if (userinfo.host == userhost) {
		bot.send(":"+nick+"!"+userhost, "MODE", channel, userinfo.mode, nick);
		util.log(userinfo.mode+" "+nick+" in "+channel);
	}
});

// Listen for any message said on channels first, it logs the message,
// and then it parses the message to see what it is to do next.
bot.addListener("message", function(from, to, text, message) {
	// Log anything and everything just to have it.
	util.log(from+": "+String(message.args[1]));
	var msg = String(message.args[1]).toLowerCase();
	var channel = message.args[0];
	
	
	
});

// --------------------------------------------------------------------------- //
// --------------------------------------------------------------------------- //
// --------------------------------------------------------------------------- //
