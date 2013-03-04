// 
// IRC bot
//

var irc = require("irc");
var l = require("./log");

// Create Random Quotes
var quotes = ["how many kittens?!", "you've got to be kitten me!"];

// Configure the bot
var config = {
	userName: "kittens",
	realName: "Kitten IRC Bot",
	autoRejoin: true,
	autoConnect: true,
	channels: ["#marylandmesh"],
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
	l.appendLog(from+": "+String(message.args[1]));
	
	
	if (String(message.args).indexOf(config.botName) > -1) {
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