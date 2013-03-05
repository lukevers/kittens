// 
// Kittens - IRC Bot
//

var irc = require("irc");
var request = require("request");
var l = require("./log");

// Create random quotes
var quotes = ["how many kittens?!", "you've got to be kitten me!"];

// Set auto-ops
var op = ["lukevers", "DuoNoxSol", "Dan68_", "thefinn93", "derpz", "werecat"];
var jop = {"lukevers":"~luke@li557-106.members.linode.com", "DuoNoxSol":"~duonoxsol@li557-106.members.linode.com",
           "Dan68_":"~Dan@199.83.100.24", "derpz":"~derp@2600:3c01::f03c:91ff:fedf:a466", "thefinn93":"~thefinn93@thefinn93.com", 
		   "werecat":"~werecatd@2002:44ae:bc22:0:1e6f:65ff:fea7:d09f"};

// Set auto-voices
var voice = ["dylwhich", "inhies"];
var jvoice = {"dylwhich":"~dylwhich@li557-106.members.linode.com", "inhies":"~inhies@8.23.56.82"};

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
	bot.say(config.channels[0], "The new topic on "+channel+" is \"\u0002"+topic+"\u000f\"");
});

// Listen for people to join the channels,
// And if they're supposed to be an OP and
// They're not already autooped, then they
// Should be op'ed. Same for auto-voice.
bot.addListener("join", function(channel, nick, message){
	l.appendLog(nick+" joined "+channel);
	autoOp(nick, channel);
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
		postLink(findUrl(message), from);
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
			postLink("http://xkcd.com/"+msg.substring(8).trim()+"/", from);
		} 
		// If someone says "kittens" but none
		// Of the other conditions apply, the
		// Bot should just send the channel a
		// Random quote.
		else {
			bot.say(config.channels[0], from+": "+RandomQuote());
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
	return quotes[Math.floor(Math.random()*quotes.length)];
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
function postLink(url, from) {
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

// The function autoOP cycles through
// The list of people that are always
// Going to be OP'd, and then if said
// Person is on the list then they'll
// Be OP'd.
function autoOP(nick, channel) {
	for (var i = 0; i < op.length; i++) {
		if (op[i] == nick) {
			bot.send(":"+nick+"!"+jop[[nick]],"MODE", channel, "+o", nick);
			l.appendLog(":"+nick+"!"+jop[[nick]]+" MODE "+channel+" +o "+nick);
		}
	}
}

// The function autoVoice cycles through
// The list of people that should always
// Be voiced, and if the person is on it
// Then they'll be voiced.
function autoVoice(nick, channel) {
	for (var i = 0; i < voice.length; i++) {
		if (voice[i] == nick) {
			bot.send(":"+nick+"!"+jvoice[[nick]],"MODE", channel, "+v", nick);
			l.appendLog(":"+nick+"!"+jvoice[[nick]]+" MODE "+channel+" +v "+nick);
			l.appendLog("Voiced "+nick);
		}
	}
}