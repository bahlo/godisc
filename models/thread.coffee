mongoose = require 'mongoose'
Schema = mongoose.Schema
Post = require './post'
Account = require './account'

threadSchema = new Schema
  topic: String
  created:
    type: Date,
    default: Date.now
  creator:
    type: Schema.Types.ObjectId,
    ref: 'Account'
  posts: [
    type: Schema.Types.ObjectId,
    ref: 'Post'
  ]

module.exports = mongoose.model 'Thread', threadSchema
