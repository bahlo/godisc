http = require 'http'
path = require 'path'
express = require 'express'
mongoose = require 'mongoose'
passport = require 'passport'
moment = require 'moment'
config = (require 'cson').parseFileSync 'config.cson'
LocalStrategy = (require 'passport-local').Strategy

Account = require './models/account'
routes = require './routes'

# Configure express
app = express()
app.set 'port', process.env.PORT || 3000
app.set 'views', path.join __dirname, 'views'
app.use express.bodyParser keepExtensions: true, uploadDir: 'uploads'
app.set 'view engine', 'jade'
app.use express.favicon()
app.use express.logger 'dev'
app.use express.json()
app.use express.urlencoded()
app.use express.methodOverride()
app.use express.cookieParser 'your secret here'
app.use express.session()
app.use passport.initialize()
app.use passport.session()
app.use app.router
app.use (require 'less-middleware') src: path.join __dirname, 'public'
app.use express.static path.join __dirname, 'public'

if 'development' is app.get 'env'
  app.use express.errorHandler()

# Configure locals
app.locals.moment = moment
app.locals.config = config

# Configure passport
passport.use new LocalStrategy Account.authenticate()
passport.serializeUser Account.serializeUser()
passport.deserializeUser Account.deserializeUser()

mongoose.connect config.dbUrl

# Routing
routes app

(http.createServer app).listen (app.get 'port'), ->
  console.log "Express server listening on port #{app.get 'port'}"
