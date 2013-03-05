// 
// Kittens - IRC Bot
//

var irc = require("irc");
var request = require("request");
var l = require("./log");
var c = require("./config");

l.appendLog("Configured "+c.config.botName);

// Create the bot
var bot = new irc.Client(c.config.server, c.config.botName, {
	channels: c.config.channels
});

l.appendLog("Created "+c.config.botName);
l.appendLog("Connecting to "+c.config.server);

// Listen for topic changes on channels,
// And when there is a topic change, the
// Bot will announce the new topic.
bot.addListener("topic", function(channel, topic, nick, message){
	l.appendLog("The new topic on "+channel+" is \""+topic+"\"");
	bot.say(channel, "The new topic on "+channel+" is \"\u0002"+topic+"\u000f\"");
});

// Listen for people to join the channels,
// And if they're supposed to be an OP and
// They're not already autooped, then they
// Should be op'ed. Same for auto-voice.
bot.addListener("join", function(channel, nick, message){
	l.appendLog(nick+" joined "+channel);
	autoOP(nick, channel);
	autoVoice(nick, channel);
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
		postLink(findUrl(message), from, message.args[0]);
	} 
	
	// If someone says meow, then meow
	// Back at them!
	else if (String(message.args[1]).toLowerCase().indexOf("meow") > -1) {
		bot.say(message.args[0], from+": meow!");
	}

	// If someone says "kittens"
	else if (String(message.args[1]).toLowerCase().indexOf(c.config.botName) > -1) {
		var msg = String(message.args[1]).toLowerCase();
		// If someone just says a lone number,
		// Get the relevant xkcd comic.
		if (!isNaN(msg.substring(c.config.botName.length+1).trim())) {
			postLink("http://xkcd.com/"+msg.substring(c.config.botName.length+1).trim(), from, message.args[0]);
		} 
		// If someone says "kittens" but none
		// Of the other conditions apply, the
		// Bot should just send the channel a
		// Random quote.
		else {
			bot.say(message.args[0], from+": "+RandomQuote());
		}
	}
});

// --------------------------------------------------------------------------- //
// --------------------------------------------------------------------------- //
// --------------------------------------------------------------------------- //

// The function RandomQuote gets a random
// Quote to be said back to a user in the
// IRC channel if nothing else is said to
// The bot when parsing.
function RandomQuote() {
	return c.quotes[Math.floor(Math.random()*c.quotes.length)];
}

// The function findURL searches through
// A message that someone says, and then
// It finds just the URL from the String
// And returns it.
function findUrl(message) {
	if (String(message.args[1]).indexOf("https") > -1) return findUrlHTTPS(message);
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
	return url;
}

// The function findURLHTTPS is called
// When the function findURL finds out
// That what it's searching for is not
// An HTTP call.
function findUrlHTTPS(message) {
	var before = String(message.args[1].substring(0, String(message.args[1]).toLowerCase().indexOf("https")));
	var msgAtURL = message.args[1].substring(before.length);
	var after = msgAtURL.substring(msgAtURL.indexOf(" "));
	var url = msgAtURL.substring(0, msgAtURL.indexOf(after));
	if (url == "") url = after;
	var host = url;
	var path = "/";
	if (url.substring(8).indexOf("/") > -1) {
		host = url.substring(8, (url.substring(8).indexOf("/")+8));
		path = url.substring(host.length+8);
	}
	return url;
}

// The function postLink gets a certain
// Link that someone said and then gets
// The title of the link and relays the
// Information back to the channel.
function postLink(url, from, channel) {
	l.appendLog("GET request for ["+url+"] from "+from);
	
	request({
		uri: url,
	}, function(err, res, body) {
		var title = /<title>(.*)<\/title>/.exec(body);
		if (title != null) {
			l.appendLog(url+" - "+title[1]);
			bot.say(channel, url+" - \u0002"+title[1]+"\u000f");
		}
	});
}

// The function autoOP cycles through
// The list of people that are always
// Going to be OP'd, and then if said
// Person is on the list then they'll
// Be OP'd.
function autoOP(nick, channel) {
	for (var i = 0; i < c.op.length; i++) {
		if (c.op[i] == nick) {
			bot.send(":"+nick+"!"+c.jop[[nick]],"MODE", channel, "+o", nick);
			l.appendLog(":"+nick+"!"+c.jop[[nick]]+" MODE "+channel+" +o "+nick);
		}
	}
}

// The function autoVoice cycles through
// The list of people that should always
// Be voiced, and if the person is on it
// Then they'll be voiced.
function autoVoice(nick, channel) {
	for (var i = 0; i < c.voice.length; i++) {
		if (c.voice[i] == nick) {
			bot.send(":"+nick+"!"+c.jvoice[[nick]],"MODE", channel, "+v", nick);
			l.appendLog(":"+nick+"!"+c.jvoice[[nick]]+" MODE "+channel+" +v "+nick);
			l.appendLog("Voiced "+nick);
		}
	}
}