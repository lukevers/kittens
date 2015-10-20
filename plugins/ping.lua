on("PRIVMSG", "ping")

function ping(channel, message)
    if string.find(message, "^@ping$") then
        say(channel, "pong")
    end
end
