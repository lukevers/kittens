on("PRIVMSG", function (event)
    if string.find(event["message"], "^@echo%s") then
        say(event["channel"], string.gsub(event["message"], "^@echo%s", ""))
    end
end)
