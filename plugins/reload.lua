on("PRIVMSG", function (event)
    if string.find(event["message"], "^@reload$") then
        say(event["channel"], "Reloading plugins...")
        reload()
    end
end)
