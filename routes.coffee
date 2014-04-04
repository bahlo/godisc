fs = require 'fs'
path = require 'path'
passport = require 'passport'
async = require 'async'
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
  app.get '/', loggedIn, (req, res) ->
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

  app.get '/settings', loggedIn, (req, res) ->
    res.render 'settings', user: req.user

  app.post '/settings', loggedIn, (req, res) ->
    saveUser = (data) ->
      req.user.update data, (err) ->
        req.user.save()
        res.redirect '/settings'

    if picture = req.files?.picture
      fs.readFile picture.path, (err, data) ->
        newPath = "#{__dirname}/public/uploads/#{picture.originalFilename}"
        console.log "Writing to #{newPath}"
        fs.writeFile newPath, data, (err) ->
          saveUser
            displayName: req.body.displayname
            picture: "/uploads/#{picture.originalFilename}"
    else
      saveUser displayName: req.body.displayname

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
      # Just think of it as
      # thread.posts.populate 'author', (err, posts) ->
      # More info at https://github.com/LearnBoost/mongoose/issues/601
      async.map thread.posts, (post, done) ->
        post.populate 'author', done
      , (err, posts) ->
        res.render 'thread',
          user: req.user
          thread: thread
          posts: posts

  app.get '/thread/:threadId/remove', loggedIn, (req, res) ->
    Thread.findOneAndRemove
       _id: req.params.threadId
       creator: req.user._id
    , (err, thread) ->
      res.redirect "/"

  # Posts
  app.post '/thread/:threadId/post/new', loggedIn, (req, res) ->
    if req.body.body?
      Thread.findById req.params.threadId, (err, thread) ->
        post = new Post
          body: req.body.body
          author: req.user
        post.save()

        thread.posts.push post
        thread.save()

        res.redirect "/thread/#{thread.id}"

  app.get '/thread/:threadId/post/:postId/remove', loggedIn, (req, res) ->
    Post.findOneAndRemove
       _id: req.params.postId
       author: req.user._id
    , (err, thread) ->
      res.redirect "/thread/#{req.params.threadId}"
