module.exports = (app, client, config) ->

        app.get '/', (req, res) ->
                res.render 'index'

        app.get '/admin', (req, res) ->
                console.log config
                res.render 'admin', conf: config

        app.get '*', (req, res) ->
                res.status 404
                res.send '404 Not Found'