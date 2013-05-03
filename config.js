// 
// Config.js
//
// Instead of having a config file, it's easier to just to have this
// file be manipulated for personal use.
//

// Most user configuration will go on here. Further customization,
// however, is possible below.
var config = {
	userName: "kittens",
	realName: "Kitten IRC Bot",
	autoRejoin: true,
	autoConnect: true,
	channels: ["#marylandmesh"],
	server: "irc.efnet.org",
	botName: "kittens",
	usersFile: "./users.json"
};

// Set random quotes.
var quotes = ["how many kittens?!", "you've got to be kitten me!", "I've got a good feline about this!", "that's the cat's pajamas!"];

// Set greetings to be said to people.
var greetings = ["hi", "hello", "howdy", "hai", "bonjour"];

// Set farewells to be said to people.
var farewells = ["bye", "bai", "farewell", "have a nice day", "cherrio"];

// Set key-word threats.
var keyThreats = ["castrate", "kill", "destroy", "murder", "fuck", "fry", "bake", "hurt", "burn", "slice", "chop"];

// Set responses to threats.
var threats = ["I will claw your eyes out", "I will pee on everything you love", "I will claw your walls",
               "I will throw hairballs at you", "I will fart on you while you sleep",
               "I will burn the \u0002heart\u000f out of you."];

users = require(config.usersFile);

// Set the variables for the main IRC file to be able to use.
exports.quotes = quotes;
exports.config = config;
exports.keyThreats = keyThreats;
exports.threats = threats;
exports.greetings = greetings;
exports.farewells = farewells;
exports.users = users;
