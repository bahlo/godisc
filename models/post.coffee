mongoose = require 'mongoose'
Schema = mongoose.Schema
Account = require './account'

postSchema = new Schema
  body: String
  created:
    type: Date,
    default: Date.now
  author:
    type: Schema.Types.ObjectId
    ref: 'Account'

module.exports = mongoose.model 'Post', postSchema
