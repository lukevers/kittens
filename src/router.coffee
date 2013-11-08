module.exports = (app, client) ->

        app.get '/', (req, res) ->
                res.send 'TEST'

        app.get '*', (req, res) ->
                res.status 404
                res.send '404 Not Found'