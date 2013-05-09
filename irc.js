// 
// Kittens - IRC Bot
//

var util = require("util");
var irc = require("irc");
var request = require("request");
var c = require("./config");

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
	if (userinfo == null) {
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
	
	// Check if someone posted a link. If so, then get some
	// information about the posted link.
	if (msg.indexOf("http") > -1) {
		postLink(findUrl(message), from, message.args[0]);
	} 
	
	// Parse every + request
	else if (msg.indexOf("+") > -1) {
		parseCommands(from, message);
	}
	
	// If someone says meow, then meow back at them!
	else if (msg.indexOf("meow") > -1) {
		bot.say(message.args[0], from+": meow!");
	}

	// If someone says "kittens"
	else if (msg.indexOf(c.config.botName) > -1) {
		// If someone says hello to the bot, then the bot should say
		// Hello back to them!
		if (containsGreeting(msg)) {
			bot.say(message.args[0], from+": "+RandomGreeting());
		}
		// If someone says goodbye to the bot then the bot should say
		// it back to them!
		else if (containsFarewell(msg)) {
			bot.say(message.args[0], from+": "+RandomFarewell());
		}
		// If someone threatens the bot it can't just sit around and
		// not do anything! Fight back!
		else if (isThreatened(msg)) {
			bot.say(message.args[0], from+": "+RandomThreat());
		}
		// If someone just says a lone number, get the relevant xkcd
		// comic.
		else if (!isNaN(msg.substring(c.config.botName.length+1).trim())) {
			var np = /\d+/g;
			var is = msg.substring(c.config.botName.length+1).trim().match(np);
			// if it's not a number, just send a fun message!
			if (is == "" || is == " " || is == null || is == "\n") {
				bot.say(message.args[0], from+": "+RandomQuote());
			} else {
				postLink("http://xkcd.com/"+msg.substring(c.config.botName.length+1).trim(), from, message.args[0]);	
			}
		} 
		// If someone says "kittens" but none of the other conditions
		// apply, the bot should just send the channel a random quote.
		else {
			bot.say(message.args[0], from+": "+RandomQuote());
		}
		
	}
});

// --------------------------------------------------------------------------- //
// --------------------------------------------------------------------------- //
// --------------------------------------------------------------------------- //

// RandomQuote gets a random quote to be said back to a user in the
// IRC channel.
function RandomQuote() {
	return c.quotes[Math.floor(Math.random()*c.quotes.length)];
}

// RandomGreeting gets a random qreeting to be said back to a user in
// the IRC channel.
function RandomGreeting() {
	return c.greetings[Math.floor(Math.random()*c.greetings.length)];
}

// The function RandomFarewell gets a random farewell to be said back
// to a user in the IRC channel.
function RandomFarewell() {
	return c.farewells[Math.floor(Math.random()*c.farewells.length)];
}

// RandomThreat gets a random threat to be said back to a user in the
// IRC channel.
function RandomThreat() {
	return (Math.round(Math.random()) % 2 == 0) ? "I will "+c.keyThreats[Math.floor(Math.random()*c.keyThreats.length)]+" you" : c.threats[Math.floor(Math.random()*c.threats.length)];
}

// findURL searches through a message that someone says, and then it
// finds just the URL from the String and returns it.
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

// findURLHTTPS is called then the function findURL finds out that
// what it's searching for is not an HTTP call.
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

// postLink gets a certain link that someone said and then gets the
// title of the link and relays the information back to the channel.
function postLink(url, from, channel) {
	util.log("GET request for ["+url+"] from "+from);
	
	request({
		uri: url,
	}, function(err, res, body) {
		var title = /<title>(.*)<\/title>/.exec(body);
		if (title != null) {
			util.log(url+" - "+title[1]);
			bot.say(channel, url+" - \u0002"+title[1]+"\u000f");
		}
	});
}

// isThreatened will check a message that is sent to Kittens if it is
// threatening the bot.
function isThreatened(msg) {
	for (var i = 0; i < c.keyThreats.length; i++) {
		if (msg.indexOf(c.keyThreats[i]) > -1) {
			return true;
		}
	}
	return false;
}

// containsGreeting will check to see if the phrase said to the bot
// contains a greeting.
function containsGreeting(msg) {
	for (var i = 0; i < c.greetings.length; i++) {
		if (msg.indexOf(c.greetings[i]) > -1) {
			return true;
		}
	}
	return false;
}

// containsFarewell will check to see if the phrase said to the bot
// contains a farewell.
function containsFarewell(msg) {
	for (var i = 0; i < c.farewells.length; i++) {
		if (msg.indexOf(c.farewells[i]) > -1) {
			return true;
		}
	}
	return false;
}

// parseCommands will check to see if the command is a real command
// and if so then it will do stuff depending on what the command is
function parseCommands(from, message) {
	var commands = ["help", "op", "deop", "voice", "devoice"];
	if (message.args[1].indexOf("+help") > -1) {
		m = from+": ";
		for (var i = 0; i < commands.length; i++) 
			m+= " +"+commands[i];
		bot.say(message.args[0], m);
	} 
} 