function echo(channel, message)
    say(channel, message)
end

on("PRIVMSG", "echo")
