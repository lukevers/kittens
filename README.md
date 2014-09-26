# kittens [![Build Status](https://travis-ci.org/lukevers/kittens.png?branch=master)](https://travis-ci.org/lukevers/kittens)

Kittens is an IRC bot written in Go. I just recently started rewritting kittens in Go, so it lacks a lot of features currently. If you think you can help, please feel free to contribute!

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
cd kittens
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

# Flags

### Debug

```bash
--debug
```
By including the debug flag, Kittens will do the following:

* Recompile webserver templates on each page load
* Provide verbose stdout output

By default, Kittens sets debug to `false`. This is a good option when developing Kittens.


### Webserver Port

```bash
--port [port]
```

By including the webserver port flag you can change the port that Kittens webserver listens on by default. By default Kittens webserver listens on port `3000`.

### Webserver Interface

```bash
--interface [interface]
```

By including the webserver interface flag you can change the interface that Kittens webserver binds to by default. By default the Kittens webserver binds to the interface `0.0.0.0`.

### Database Driver

```bash
--driver [driver]
```

By including the database driver flag you can change the type of database that we are connecting to. By default Kittens uses `sqlite3` as the default driver because Kittens uses a SQLite3 database as default. If you change the database driver, you need to change the database connection details or it will not work.

Kittens supports SQLite, MySQL, and PostgreSQL. To see how to use the database driver flag with the database flag, read the information on the database flag.

### Database

```bash
--database [connection]
```

By including the database flag you can change the connection details that we use. By default Kittens uses `kittens.db` as the database to connect to. If you change the database, you need to change the database driver or it will not work. 

Kittens supports SQLite, MySQL, and PostgreSQL. The full list of options for database connection details can be found on each's website respectively. Here is an example for each:

##### SQLite

```bash
--driver sqlite3 --database /etc/kittens.db
```

##### PostgreSQL

```bash
--driver postgres --database "user=username dbname=kittens sslmode=disable"
```

##### MySQL

```bash
--driver mysql --database "username:password@tcp(host:port)/database"
```

If `?parseTime=true` is not included in the database string when connecting to a MySQL database, Kittens will automatically add it before trying to connect. It is needed in order for Kittens to work properly.

# Screenshots

### Login

![Login](http://blog.lukevers.com/content/images/2014/Sep/login-2.png)

### Login 2FA

![Login 2FA](http://blog.lukevers.com/content/images/2014/Sep/2fa-1.png)

### Home

![Home](http://blog.lukevers.com/content/images/2014/Sep/home-1.png)

### New Bot/Server

![New Bot/Server](http://blog.lukevers.com/content/images/2014/Sep/newserver-1.png)

### Server Disabled

![Server Disabled](http://blog.lukevers.com/content/images/2014/Sep/serverdisabled.png)

### Server Connecting

![Server Connecting](http://blog.lukevers.com/content/images/2014/Sep/serverconnecting.png)

### Server Connected

![Server Connected](http://blog.lukevers.com/content/images/2014/Sep/serverconnected.png)

### Settings

![Settings](http://blog.lukevers.com/content/images/2014/Sep/settings-1.png)

### Settings 2FA

![Settings 2FA](http://blog.lukevers.com/content/images/2014/Sep/settings2fa.png)

### Users

![Users](http://blog.lukevers.com/content/images/2014/Sep/users.png)
