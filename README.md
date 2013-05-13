# kittens

*kittens* is an IRC bot, written in Javascript that runs on a Node.js server. For more information about Node.js, visit their [website](http://nodejs.org/). 

## Install

### Node and NPM

If you already have [`node`](http://nodejs.org/) and [`npm`](https://npmjs.org/) installed, skip this section. Otherwise, visit [here](https://github.com/joyent/node/wiki/Installing-Node.js-via-package-manager) to find out how to install properly. 

### Install

```bash
git clone https://github.com/lukevers/kittens.git
cd kittens
npm install
```

## Configuration

There are three JSON files that you must set up. There is an example of each JSON file called `example.*.json` in the repo. You will have to create your own of each of the following configuration files:

1. config.json
2. plugins.json
3. users.json

### Config.json

Although many of the variables already set in `config.json` do not have to be changed, there are a few items that have to be changed. These are some of the defaults (see `example.config.json` for all of the defaults), and generally, all of these will have to be changed. Unless, of course, you want your bot to connect to `irc.efnet.org` and join `#channel`, `#channel_2`, and `#channel_3`.

```json
{
  "botName": "ircbot",
  "server": "irc.efnet.org",
  "channels": [
    "#channel",
    "#channel_2",
    "#channel_3"
  ]
}
```

### Plugins.json

Every plugin is automatically loaded if you use the default `plugins.json` file. If you do not want a certain plugin to be used by your instance of the bot, just change "true" to "false" in your `plugins.json` file.

```json
{
  "op": "true",
  "voice": "true",
  "topic": "true"
}
```

### Users.json

The following packages use the `users.json` file:

1. Op
2. Voice

Whomever is in charge of your bot should be given the master title. You do not have to put `"master": false` for people who are not masters, for that is a given. Multiple people can have the master title given to them, and the master title may be used by some plugins to determine if a person is entitled to make these changes. The following is just an example of what a `users.json` file could look like:

```json
{
  "lukevers": {
    "master": true,
    "mode": "+o",
    "host": "~lukevers@dylwhich.xen.prgmr.com"
  },
  "oranges": {
    "mode": "+o",
    "host": "~oranges@herpderp.com"
  },
  "humans": {
    "mode": "+v",
    "host": "~humans@192.168.1.1"
  }
}
```
