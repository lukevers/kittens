// 
// Config.js
//
// Instead of having a config file,
// It's easier to just to have this
// File be manipulated for personal
// Use.
//


// Set random quotes
var quotes = ["how many kittens?!", "you've got to be kitten me!", "I've got a good feline about this!", "that's the cat's pajamas!"];

// Set greetings to be said to people
var greetings = ["hi", "hello", "howdy", "hai", "bonjour"];

// Set farewells to be said to people
var farewells = ["bye", "bai", "farewell", "have a nice day", "cherrio"];

// Set auto-ops
var op = ["lukevers", "DuoNoxSol", "Dan68_", "thefinn93", "derpz", "werecat"];
var jop = {"lukevers":"~luke@li557-106.members.linode.com", "DuoNoxSol":"~duonoxsol@li557-106.members.linode.com",
           "Dan68_":"~Dan@199.83.100.24", "derpz":"~derp@2600:3c01::f03c:91ff:fedf:a466", "thefinn93":"~thefinn93@thefinn93.com", 
		   "werecat":"~werecatd@2002:44ae:bc22:0:1e6f:65ff:fea7:d09f"};

// Set auto-voices
var voice = ["dylwhich", "inhies"];
var jvoice = {"dylwhich":"~dylwhich@li557-106.members.linode.com", "inhies":"~inhies@8.23.56.82"};

// Set key-word threats
var keyThreats = ["castrate", "kill", "destroy", "murder", "fuck", "fry", "bake", "hurt", "burn", "slice", "chop"];

// Set responses to threats
var threats = ["I will claw your eyes out", "I will pee on everything you love", "I will claw your walls",
               "I will throw hairballs at you", "I will fart on you while you sleep",
               "I will burn the \u0002heart\u000f out of you."];

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

// Set the variables for the main IRC
// File to be able to use.
exports.quotes = quotes;
exports.op = op;
exports.jop = jop;
exports.voice = voice;
exports.jvoice = jvoice;
exports.config = config;
exports.keyThreats = keyThreats;
exports.threats = threats;
exports.greetings = greetings;
exports.farewells = farewells;
