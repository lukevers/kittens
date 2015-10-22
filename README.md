Hi! If you're reading this then you're currently looking at the refactor branch. A lot is changing from the old branch. Feel free to help out if you want. Checkout the master branch (or in the git history) if you want to see the old code.

---

# Kittens

An IRC bot hub.

## History

Kittens started as a really basic IRC bot written in [Node.js](https://nodejs.org/en/) when I wanted to learn some server-side JavaScript Then it evolved into [CoffeeScript](http://coffeescript.org/) when I decided I wanted to learn CoffeeScript. Eventually it transformed into [Go](http://golang.org/) when I wanted to learn Go! (See a pattern here?) I'm most likely not going to be re-writing Kittens in any other language anymore. Instead, I decided to refactor it and transform it into an IRC bot hub instead of an IRC bot.

## An IRC Bot Hub?

Yes!

When I started writting Kittens I just wanted a nice IRC bot. When I rewrote Kittens for the second time, I gave it a web user-interface--but it was never completed. This time I've decided to make Kittens more user-oriented. By default when you run Kittens for the first time it will create a default user/password (which you may--and really, really should since the password is in the code--change once logged in for the first time) for you to use. There's an environment variable you can set to allow other users to register if you wish to allow multiple users.

Users create bots that connect to networks which are controlled via their user on the web interface. Each bot can be configured however the user wants with any number of plugins.

## License

The old code was licensed under [GPL-3.0](http://opensource.org/licenses/GPL-3.0), but as of the most recent refactor I've decided to relicense Kittens under [MIT](http://opensource.org/licenses/MIT). See the [LICENSE](LICENSE.md) file for more information regarding the license.

## Installing

TODO LATER WHEN THE REFACTOR IS DONE.

## Environment Variables

TODO LATER WHEN THE REFACTOR IS DONE.

## Plugins

Plugins are written in [GopherLua](https://github.com/yuin/gopher-lua) which is almost identical to [Lua](http://www.lua.org/), but with slight differences. Take a look at the [differences between Lua and GopherLua](https://github.com/yuin/gopher-lua#differences-between-lua-and-gopherlua) to get started.

Let's look at the structure of a plugin:

```lua
on("EVENTCODE", function (event)
    -- implementation here
end)
```

The function `on` is a global function that binds a function to an irc event code. The event code is expected to be a string. The callback function is passed one parameter which contains all of the event information in a table. The event information includes the following data:

1. `message` - The IRC message.
2. `channel` - The IRC channel that this event happened at.
3. `nick` - The nickname of the user that is related to the event.
4. `host` - The hostname of the user that is related to the event.
5. `user` - The user that is related to the event.
6. `source` - The full host (`<nick>!<usr>@<host>`) of the user that is related to the event.
7. `raw` - The raw IRC message.

### Global Functions

### On

TODO

#### Say

TODO

---

This guide is still being written as the codebase has been changing drastically. Feel free to open any issue you want.
