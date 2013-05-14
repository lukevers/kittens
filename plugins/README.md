# Writing a plugin for Kittens

Kittens is an IRC bot written in Node.js. To learn the basics of Kittens, read the [readme file](../README.md) in the main directory of the repo.

## API

Kittens uses an IRC package for node.js called [node-irc](https://github.com/martynsmith/node-irc). In order to write a plugin for Kittens, all that is really needed is the [API](https://node-irc.readthedocs.org/en/latest/API.html). The most important part of the API is the section on [events](https://node-irc.readthedocs.org/en/latest/API.html#events).

## Types of plugins

There are two types of plugins: plugins that contain commands, and plugins that do not contain commands. A plugin that does has commands must create an array, generally denoted by `var commands = ["+command", "+command2", ...]`. The plugin must
