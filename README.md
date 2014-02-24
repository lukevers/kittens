# Kittens

An IRC bot written in [CoffeeScript](http://coffeescript.org/). котята!

[![Build Status](https://travis-ci.org/lukevers/kittens.png?branch=master)](https://travis-ci.org/lukevers/kittens)

## Install

### 0. Install dependencies

Kittens needs a few things in order to run. [Git](http://git-scm.com/), in order to download the repository, and [Node.js](http://nodejs.org/) to run it.

### 1. Get the source code

Clone the repository from GitHub and change to that directory:

```bash
git clone https://github.com/lukevers/kittens.git
cd kittens
```

### 2. Install node modules

```bash
npm install
npm install -g coffee-script
```

### 3. Build

```bash
cake build
```

If nothing is wrong, the output should be `Compiled with no errors`.

## Setup

Upon running Kittens for the first time, you'll be prompted to enter some information about what server to connect. If you leave any blank, the default value will be used. The default values are as follows:

```
Bot Name: kittens
Username: kittens
Real Name: kittens
Server: irc.hypeirc.net
ServerName: server
Port: 6667
Command Symbol: !
Channels:

Add another server? (yes/no): no
```

To add multiple channels, user commas in between channels. For example, `#a #b #c` would put you in the channel `#a#b#c`, but `#a, #b, #c` would put you in the channels `#a`, `#b`, and `#c`.

## Startup

Starting up Kittens is farily easy. Once you're inside the direcotry of the repository, you just run `coffee kittens.coffee`. You can have the process run in the background, but you can use stdin to run commands.

## Commands

This is a list of commands that you can run from stdin after starting up Kittens.

#### help

`help` prints out a list of all available commands.

#### quit

`quit` disconnects from all servers and shuts down kittens.

#### disconnect

`disconnect all` disconnects from all servers.

`disconnect [serverName]` disconnects from a certain server, assuming it is a connected server.

#### connect

`connect all` connects all servers.

`connect [serverName]` connects a certain server, assuming it is in the configuration file.

#### join

`join [serverName] [channel]` joins the channel specified on the server specified.

#### part

`part [serverName] [channel]` parts the channel specified on the server specified.

#### whois

`whois [serverName] [nick]` runs a whois query on the nick specified on the server specified.

#### say

`say [server] [channel/user] [message]` sends a message to the channel or user specified on the server specified.

#### set

`set [serverName] server [newserver]` sets a new server domain on the server specified.

`set [serverName] serverName [newname]` sets a new server name on the server specified

`set [serverName] port [port]` sets a new port on the server specified.

`set [serverName] commandSymbol [symbol]` sets a new command symbol on the server specified.

`set [serverName] nick [nick]` sets a new nickname on the server specified.

`set [serverName] user [user]` sets a new username on the server specified.

`set [serverName] name [name]` sets a new real name on the server specified.

`set [serverName] mode [channel] [nick] [mode]` sets the mode for a user in a channel on the server specified.

`set [serverName] host [channel] [nick] [host]` sets the host for a user in a channel on the server specified.

#### get

`get [serverName] mode [channel] [user]` gets the mode for a user in a channel on the server specified.

`get [serverName] host [channel] [user]` gets the host for a user in a channel on the server specified.

## Contributing

If you see something that bothers you, or have an idea that would make Kittens better, feel free to open an issue or submit a pull request! All are welcome, and I encourage them! See [CONTRIBUTING.md](./CONTRIBUTING.md) for full details.

## Copyright & License

Kittens is licensed under GPLv3. See [LICENSE](./LICENSE) for full details.
