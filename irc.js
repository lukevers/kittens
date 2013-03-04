// 
// IRC bot
//

var irc = require("irc");
var http = require('http');
var l = require("./log");

// Create Random Quotes
var quotes = ["how many kittens?!", "you've got to be kitten me!"];

// Configure the bot
var config = {
	userName: "kittens",
	realName: "Kitten IRC Bot",
	autoRejoin: true,
	autoConnect: true,
	channels: ["#herpderpbot"],
	server: "irc.efnet.org",
	botName: "kittens"
};

l.appendLog("Configured "+config.botName);

// Create the bot
var bot = new irc.Client(config.server, config.botName, {
	channels: config.channels
});

l.appendLog("Created "+config.botName);
l.appendLog("Connecting to "+config.server);

// Listen for topic changes on channels,
// And when there is a topic change, the
// Bot will announce the new topic.
bot.addListener("topic", function(channel, topic, nick, message){
	l.appendLog("The new topic on "+channel+" is \""+topic+"\"");
	bot.say(config.channels[0], "The new topic on "+channel+" is \""+topic+"\"");
});

// Listen for any message said on channels
// First, it logs the message, and then it
// Parses the message to see what it is to
// Do next.
bot.addListener("message", function(from, to, text, message) {
	// Log anything and everything just to have it
	l.appendLog(from+": "+String(message.args[1]));
	
	// Check if someone posted a link. If so, then
	// Get some information about the posted link.
	if (String(message.args[1]).toLowerCase().indexOf("http") > -1) {
		var before = String(message.args[1].substring(0, String(message.args[1]).toLowerCase().indexOf("http")));
		var msgAtURL = message.args[1].substring(before.length);
		var after = msgAtURL.substring(msgAtURL.indexOf(" "));
		var url = msgAtURL.substring(0, msgAtURL.indexOf(after));
		if (url == "") url = after;
		
	} 
	
	// If none of the other things have been done,
	// Then let's just send them a random quote if
	// They said "kittens." 
	else if (String(message.args[1]).toLowerCase().indexOf(config.botName) > -1) {
		bot.say(config.channels[0], from+": "+RandomQuote());
	}
});

// The function RandomQuote gets a random
// Quote to be said back to a user in the
// IRC channel if nothing else is said to
// The bot when parsing.
function RandomQuote() {
	return quotes[Math.floor(Math.random()*quotes.length)];
}