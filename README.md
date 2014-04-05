# nodisc

A _small_ forum in node.js

## Install
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
