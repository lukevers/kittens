# kittens

Kittens is an IRC bot written in Go. I just recently started rewritting kittens in Go, so it lacks a lot of features currently. If you think you can help, please feel free to contribute! Currently I have an old version of kittens on the branch [coffee-depreciated](https://github.com/lukevers/kittens/tree/coffee-depreciated) written in CoffeeScript. If you want to check it out go ahead, but when the Go version has more features I'll be deleting that branch (although it will still be in the git history, it just won't have a branch).

[![Build Status](https://travis-ci.org/lukevers/kittens.png?branch=master)](https://travis-ci.org/lukevers/kittens)

# Building

#### 0. Before You Build

Make sure you have [Go](http://golang.org/) installed. In order to compile the LESS/JS the preferred way is to use [Gulp](http://gulpjs.com/). To install Gulp you need to have [NPM](https://www.npmjs.org/) installed. Once you have NPM installed you can install Gulp via NPM:

```bash
npm install -g gulp
```

Once everything is installed make sure you have set your [$GOPATH](http://golang.org/doc/code.html#GOPATH) properly, or it will prove difficult to build.

#### 1. Get the Code

Start by cloning the repository and getting all the dependencies.

```bash
git clone https://github.com/lukevers/kittens
go get
```

#### 2. Build LESS/JS

Before we can run Gulp we need to make sure we install all of the necessary modules:
```bash
npm update
```

Building our webserver CSS/JS files is easy with Gulp.

```bash
gulp
```

When developing you can run `gulp watch` instead of running `gulp` every time you make changes.

If you'd rather use your own way of compiling LESS to CSS and concating all the CSS files into one file and JS files into one file, feel free. You can checkout `gulpfile.js` in the root of the directory to find out where these files are located and where they end up.

#### 3. Build the Source

```bash
go build
```

# Configuring

An example configuration file is included in the repo ([example.config.json](example.config.json)) if you'd rather jump right into it instead of reading the documentation. Reading the source is also a really good way to find out what's going on.

#### "Debug"

Debug is a bool set to `true` by default currently because Kittens is still in active development. If debug is set to `true` then templates will be recompiled each time the page is refreshed. Also verbose logging information will be logged to stdout.

#### "Interface"

Interface is a string that defines what interface the web interface should bind to. The default is `0.0.0.0` but another common one is `127.0.0.1`.

#### "Port"

Port is an integer that defines what port the web interface listens on. The default port is set to `3000`.

#### "Username"

Username is the username to be used for a user to sign in to the web interface. *THIS IS ONLY TEMPORARY WHILE THERE IS NO DATABASE CONNECTION AND I DO NOT RECOMMEND USING THIS FOR PRODUCTION YET.*

#### "Password"

Password is the password to be used for a user to sign in to the web interface. *THIS IS ONLY TEMPORARY WHILE THERE IS NO DATABASE CONNECTION AND I DO NOT RECOMMEND USING THIS FOR PRODUCTION YET.*

#### "Servers"

Servers is an array of servers to be connected to. The following are fields that each server has:

##### "Nick"

Nick is a string which is the nickname that the bot will use while connected to this server.

##### "RealName"

RealName is a string which defines the "real name" that the bot will have while connected to this server.

##### "Host"

Host is a string which defines the host that the bot will have while connected to this server.

##### "ServerName"

ServerName is a string that is used throughout the web interface to show which server is which. 

##### "Network"

Network is a string of the link that the bot uses to connect to.

##### "Port"

Port is an integer that defines that port that the bot uses to connect to the server.

##### "SSL"

SSL is set to `true` if the bot is connecting via SSL, and set to `false` if the bot is not connecting via SSL.

##### "Password"

Password is a string that is only used if connecting to a network that requires a password. If a password is given then it will be used, otherwise it will not be used.

##### "Enabled"

Enabled is a bool that is set to `true` if the bot is enabled and set to `false` if the bot is not enabled. If the bot is enabled that means the bot will be connected (or trying to connect) to the server that is given. If it is not enabled then it won't try to connect at all.

##### "Connected"

Connected is a bool that is set to `true` when a bot connects to a server and set to `false` when it disconnects/is not connected.

##### "Channels"

Channels is an array of channels to be connected to. The following are fields that all channels have:

###### "Name"

Name is the actual name of the channel that we're connected to. `#go-nuts` for example.


# Screenshots

![Dashboard](http://i.imgur.com/1vRVYLH.png)

![Update Server](http://i.imgur.com/LOyuwyT.png)
