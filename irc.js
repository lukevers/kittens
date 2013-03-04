// 
// IRC bot
//

var irc = require("irc");
var request = require("request");
var l = require("./log");

// Create Random Quotes
var quotes = ["how many kittens?!", "you've got to be kitten me!"];

// OP's and Voices
var op = ["lukevers", "derpz", "werecat", "DuoNoxSol", "Dan68_", "thefinn93"];
var voice = ["inhies", "unicorn", "dylwhich", "snap1"];

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
	bot.say(config.channels[0], "The new topic on "+channel+" is \"\u0002"+topic+"\u000f\"");
});

// Listen for people to join the channels,
// And if they're supposed to be an OP and
// They're not already autooped, then they
// Should be op'ed. Same for auto-voice.
bot.addListener("join", function(channel, nick, message){
	l.appendLog(nick+" joined "+channel);
	
	for (var i = 0; i < op.length; i++) {
		if (op[i] == nick) {
			bot.say(config.channels[0], "/op "+nick);
			l.appendLog("OP'd "+nick);
		} else if (voice[i] == nick) {
			bot.say(config.channels[0], "/voice "+nick);
			l.appendLog("Voiced "+nick);
		}
	}
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
		postLink(from, to, text, message);
	} 
	
	// If someone says "kittens"
	else if (String(message.args[1]).toLowerCase().indexOf(config.botName) > -1) {
		var msg = String(message.args[1]).toLowerCase();
		// If someone says ...
		if (msg.indexOf("orangesiadahoiwd") > -1) {
			
		} 
		// If someone just says a lone number,
		// Get the relevant xkcd comic.
		else if (!isNaN(msg.substring(8).trim())) {
			l.appendLog("http://xkcd.com/");
			bot.say(config.channels[0], "http://xkcd.com/"+msg.substring(8).trim()+"/");
		} 
		else {
			bot.say(config.channels[0], from+": "+RandomQuote());
		}
	}
});

// The function RandomQuote gets a random
// Quote to be said back to a user in the
// IRC channel if nothing else is said to
// The bot when parsing.
function RandomQuote() {
	return quotes[Math.floor(Math.random()*quotes.length)];
}

// The function postLink gets a certain
// Link that someone said and then gets
// The title of the link and relays the
// Information back to the channel.
function postLink(from, to, text, message) {
	var before = String(message.args[1].substring(0, String(message.args[1]).toLowerCase().indexOf("http")));
	var msgAtURL = message.args[1].substring(before.length);
	var after = msgAtURL.substring(msgAtURL.indexOf(" "));
	var url = msgAtURL.substring(0, msgAtURL.indexOf(after));
	if (url == "") url = after;
	var host = url;
	var path = "/";
	if (url.substring(7).indexOf("/") > -1) {
		host = url.substring(7, (url.substring(7).indexOf("/")+7));
		path = url.substring(host.length+7);
	}
	
	l.appendLog("url : "+url);
	l.appendLog("host: "+host);
	l.appendLog("path: "+path);
	
	l.appendLog("GET request for ["+url+"] from "+from);
	
	request({
		uri: url,
	}, function(err, res, body) {
		var title = /<title>(.*)<\/title>/.exec(body);
		if (title != null) {
			l.appendLog(url+" - "+title[1]);
			bot.say(config.channels[0], url+" - \u0002"+title[1]+"\u000f");
		}
	});
}