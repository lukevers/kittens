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

There are two JSON files that you must set up. There is an example of each JSON file called `example.*.json` in the repo. You will have to create your own of each of the following configuration files:

1. config.json
2. plugins.json

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

TODO