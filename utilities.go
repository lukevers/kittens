package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
    )

// Helper func so less code is duplicated here.
func GetServerFromRequest(req *http.Request) (server *Server, id uint16) {
    // Figure out what {id} is in "/server/{id}"
    id_, err := strconv.ParseUint(mux.Vars(req)["id"], 10, 16)
    if err != nil {
        warnf("Error converting server id: %s", err)
    } else {
        id = uint16(id_)
    }

    // Get our server from our slice of servers
    for _, s := range clients {
        if s.ID == uint16(id) {
            server = s
        }
    }

    return
}
