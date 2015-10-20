on("PRIVMSG", "alot")

function alot(channel, message)

    math.randomseed(os.time())
    local alots = { 
        "http://4.bp.blogspot.com/_D_Z-D2tzi14/S8TRIo4br3I/AAAAAAAACv4/Zh7_GcMlRKo/s400/ALOT.png",
        "http://1.bp.blogspot.com/_D_Z-D2tzi14/S8TflwXvTgI/AAAAAAAACxI/qgd1wYcTWV8/s320/ALOT12.png",
        "http://3.bp.blogspot.com/_D_Z-D2tzi14/S8TffVGLElI/AAAAAAAACxA/trH1ch0Y3tI/s320/ALOT6.png",
        "http://2.bp.blogspot.com/_D_Z-D2tzi14/S8TiTtIFjpI/AAAAAAAACxQ/HXLdiZZ0goU/s320/ALOT14.png",
        "http://4.bp.blogspot.com/_D_Z-D2tzi14/S8TfVzrqKDI/AAAAAAAACw4/AaBFBmKK3SA/s320/ALOT5.png",
        "http://2.bp.blogspot.com/_D_Z-D2tzi14/S8Tdnn-NE0I/AAAAAAAACww/khYjZePN50Y/s400/ALOT4.png",
        "http://1.bp.blogspot.com/_D_Z-D2tzi14/S8TZcKXqR-I/AAAAAAAACwg/F7AqxDrPjhg/s320/ALOT13.png",
        "http://3.bp.blogspot.com/_D_Z-D2tzi14/S8TW0Y2bL_I/AAAAAAAACwY/MGdywFA2tbg/s320/ALOT8.png",
        "http://3.bp.blogspot.com/_D_Z-D2tzi14/S8TWtWhXOfI/AAAAAAAACwQ/vCeUMPnMXno/s320/ALOT9.png",
        "http://3.bp.blogspot.com/_D_Z-D2tzi14/S8TWUJ0APWI/AAAAAAAACwI/014KRxexoQ0/s320/ALOT3.png",
        "http://3.bp.blogspot.com/_D_Z-D2tzi14/S8TTPQCPA6I/AAAAAAAACwA/ZHZH-Bi8OmI/s400/ALOT2.png"
    }

    if string.find(message, "alot") then
        say(channel, alots[math.random(1, #alots)])
    end
end
