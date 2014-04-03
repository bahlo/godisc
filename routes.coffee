passport = require 'passport'
config = (require 'cson').parseFileSync 'config.cson'
Account = require './models/account'
Thread = require './models/thread'
Post = require './models/post'

# Custom middleware to check if logged in
loggedIn = (req, res, next) ->
  if req.user
    next()
  else
    res.redirect '/login'


module.exports = (app) ->
  # Dashboard
  app.get '/', (req, res) ->
    ((Thread.find()).populate 'creator').exec (err, threads) ->
      res.render 'index',
        user: req.user
        threads: threads

  # User stuff
  app.get '/register', (req, res) ->
    if config?.registration is false
      res.render 'register',
        user: req.user
        err: 'Registrierung ist vorerst geschlossen!'
    else
      res.render 'register', user: req.user

  app.post '/register', (req, res) ->
    if config?.registration is false
      res.redirect '/register'
    else
      Account.register (new Account username: req.body.username )
      , req.body.password, (err, account) ->
        if err?
          return res.render 'register', account: account

        (passport.authenticate 'local') req, res, ->
          res.redirect '/'

  app.get '/login', (req, res) ->
    res.render 'login', user: req.user

  app.post '/login', (passport.authenticate 'local'), (req, res) ->
    res.redirect '/'

  app.get '/logout', (req, res) ->
    req.logout()
    res.redirect '/'

  # Thread
  app.get '/thread/new', loggedIn, (req, res) ->
    res.render 'thread_new', user: req.user

  app.post '/thread/new', loggedIn, (req, res) ->
    if req.body.topic? and req.user?
      thread = new Thread
        topic: req.body.topic
        creator: req.user._id
      thread.save()
      res.redirect "/thread/#{thread.id}"

  app.get '/thread/:id', loggedIn, (req, res) ->
    ((Thread.findById req.params.id).populate ['creator', 'posts']).exec (err, thread) ->
      res.render 'thread',
        user: req.user
        thread: thread

  app.post '/thread/:id/post/new', loggedIn, (req, res) ->
    if req.body.body?
      Thread.findById req.params.id, (err, thread) ->
        post = new Post
          body: req.body.body
          author: req.user
        post.save()

        thread.posts.push post
        thread.save()

        res.redirect "/thread/#{thread.id}"
