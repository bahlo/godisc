mongoose = require 'mongoose'
Schema = mongoose.Schema
Account = require './account'
Post = require './post'

threadSchema = new Schema
  topic: String
  created:
    type: Date,
    default: Date.now
  creator:
    type: ObjectID,
    ref: 'Account'
  posts: [
    type: ObjectId,
    ref: 'Post'
  ]

module.exports = mongoose.model 'Thread', threadSchema
