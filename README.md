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

# Screenshots

![Dashboard](http://i.imgur.com/1vRVYLH.png)

![Update Server](http://i.imgur.com/LOyuwyT.png)
