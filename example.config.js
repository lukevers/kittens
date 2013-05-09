// 
// Config.js
//
// Instead of having a config file, it's easier to just to have this
// file be manipulated for personal use.
//

// Most user configuration will go on here. Further customization,
// however, is possible below.
var config = {
	userName: "username",
	realName: "Kittens IRC Bot",
	autoRejoin: true,
	autoConnect: true,
	channels: ["#lolchannel", "#lolchannel2"],
	server: "irc.efnet.org",
	botName: "botname",
	usersFile: "./users.json"
};

// Set random quotes.
var quotes = ["how many kittens?!", "you've got to be kitten me!", "I've got a good feline about this!", "that's the cat's pajamas!"];

// Set greetings to be said to people.
var greetings = ["hi", "hello", "howdy", "hai", "bonjour"];

// Set farewells to be said to people.
var farewells = ["bye", "bai", "farewell", "have a nice day", "cherrio"];

var users = require(config.usersFile);

// Set the variables for the main IRC file to be able to use.
exports.quotes = quotes;
exports.config = config;
exports.greetings = greetings;
exports.farewells = farewells;
exports.users = users;
