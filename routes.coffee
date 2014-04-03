passport = require 'passport'
Account = require './models/account'

module.exports = (app) ->
  app.get '/', (req, res) ->
    res.render 'index', user : req.user

  app.get '/register', (req, res) ->
    res.render 'register', user: req.user

  app.post '/register', (req, res) ->
    Account.register (new Account username : req.body.username )
    , req.body.password, (err, account) ->
      if err?
        return res.render('register', { account : account });

      (passport.authenticate 'local') req, res, ->
        res.redirect '/'

  app.get '/login', (req, res) ->
    res.render 'login', user : req.user

  app.post '/login', (passport.authenticate 'local'), (req, res) ->
    res.redirect '/'

  app.get '/logout', (req, res) ->
    req.logout()
    res.redirect '/'
