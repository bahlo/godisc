# nodisc

A lightweight forum in node (WIP, unstable)

## Goal

> Forums are so bloated today, I just want a self-hosted Facebook group.

I've heard people say that - and found not one good, open-soure, small and
easy-to-install forum, so I'm building one.

## Technologie

It's built completly in [CoffeeScript](http://coffeescript.org/) upon the
[Express](http://expressjs.com/) framework for [node.js](http://nodejs.org).

## Roadmap
* Translations
* A WYSIWYG, I'm tending to use [scribe](https://github.com/guardian/scribe)
* Crop-functionality for profile pictures
* An _own_ layout
* Proper error handling

## Installation
### Requirements

You'll need [node.js](http://nodejs.org) with [npm](http://npmjs.org) and
[MongoDB](https://www.mongodb.org) installed.

On a Mac, this is done in one command ([Homebrew](http://brew.sh) required):
`brew install node mongodb`.

### Get started

Clone this repository and install npm dependencies:

```bash
git clone https://github.com/bahlo/nodisc.git
cd nodisc
npm install
```

### Run
Now start the MongoDB deamon (the `&` puts it in the background) and then your
app:

```bash
mongod &
npm start
```

You should be able to navigate to <http://localhost:3000> and register.
