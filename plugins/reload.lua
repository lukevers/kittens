on("PRIVMSG", "parse")

function parse(channel, message)
    if string.find(message, "^@reload$") then
        say(channel, "Reloading plugins...")
        reload()
    end
end
