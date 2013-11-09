module.exports = (app, client) ->

        app.get '/', (req, res) ->
                res.render 'index'

        app.get '/admin', (req, res) ->
                res.render 'admin'

        app.get '*', (req, res) ->
                res.status 404
                res.send '404 Not Found'