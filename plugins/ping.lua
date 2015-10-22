on("PRIVMSG", function (event)
    if string.find(event["message"], "^@ping$") then
        say(event["channel"], "pong")
    end
end)
