package main

import (
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "html/template"
    "net/http"
    "strconv"
    "sync"
    "os"
    "log"
)

var (
    err       error
    wg        sync.WaitGroup
    users     []*User
    servers   []*Server
    templates *template.Template
)

func main() {
    // Parse ENV
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Initialize database
    InitDatabase()

    // Remove old sessions
    CleanSessions()

    info("Starting webserver")

    // Web server routes
    route = mux.NewRouter()
    AddRoutes()

    // Initialize templates
    templates = RefreshTemplates(nil)

    // Get all users
    db.Find(&users, &User{})
    for _, user := range users {
        // Get all of the servers for each user
        db.Table("servers").Where("user_id = ?", user.Id).Find(&user.Servers)

        // Create servers
        for _, server := range user.Servers {
            // Get all of the channels for each server
            db.Table("channels").Where("server_id = ?", server.Id).Find(&server.Channels)

            // Add to slice of server
            servers = append(servers, server)

            // Add to wait group
            wg.Add(1)

            // Start our goroutine
            go server.Create()
        }
    }

    http.Handle("/", route)
    http.ListenAndServe(*interfaceFlag+":"+strconv.Itoa(*portFlag), nil)
    infof("Webserver running on port %s", *portFlag)

    wg.Wait()
}
