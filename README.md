# kittens

An IRC bot written in Go. I just recently started rewritting kittens in Go, so it lacks a lot of features currently. If you want a working version of my IRC bot checkout my CoffeScript bot on the master branch in this repository. This branch will be merged into master and takeover the master branch when it is ready. It is currently still in early development.

[![Build Status](https://travis-ci.org/lukevers/kittens.png?branch=go_new)](https://travis-ci.org/lukevers/kittens)

# Building

#### 0. Before You Build

Make sure you have [Go](http://golang.org/) installed. In order to compile the LESS/JS the preferred way is to use [Gulp](http://gulpjs.com/). To install Gulp you need to have [NPM](https://www.npmjs.org/) installed. Once you have NPM installed you can install Gulp via NPM:

```bash
npm install -g gulp
```


#### 1. Get the Code

Start by cloning the repository and getting all the dependencies.

```bash
git clone https://github.com/lukevers/kittens
go get
```

#### 2. Build LESS/JS

Building our webserver CSS/JS files is easy with Gulp.

```bash
gulp less && gulp
```

If you'd rather use your own way of compiling LESS to CSS and concating all the CSS files into one file and JS files into one file, feel free. You can checkout `gulpfile.js` in the root of the directory to find out where these files are located and where they end up.

#### 3. Build the Source

```bash
go build
```

# Screenshots

![Dashboard](http://i.imgur.com/1vRVYLH.png)

![Update Server](http://i.imgur.com/LOyuwyT.png)
