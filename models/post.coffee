mongoose = require 'mongoose'
Schema = mongoose.Schema
Account = require './account'

postSchema = new Schema
  body: String
  author:
    type: ObjectId
    ref: 'Account'

module.exports = mongoose.model 'Post', postSchema
