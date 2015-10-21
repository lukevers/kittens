on("PRIVMSG", "echo")

function echo(channel, message)
    if string.find(message, "^@echo%s") then
        say(channel, string.gsub(message, "^@echo%s", ""))
    end
end
