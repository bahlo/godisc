mongoose = require 'mongoose'
Schema = mongoose.Schema
passportLocalMongoose = require 'passport-local-mongoose'

accountSchema = new Schema
  email: String
  picture: String

accountSchema.plugin passportLocalMongoose

module.exports = mongoose.model 'Account', accountSchema
