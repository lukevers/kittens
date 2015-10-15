on("PRIVMSG", "echo")

function echo(channel, message)
    if string.find(message, "^@echo") then
        say(channel, string.gsub(message, "^@echo", ""))
    end
end
