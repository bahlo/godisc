FROM golang:latest

# Get dependencies
RUN go get github.com/coopernurse/gorp \
      github.com/go-sql-driver/mysql \
      code.google.com/p/go.crypto/bcrypt \
      code.google.com/p/go.net/websocket \
      github.com/howeyc/fsnotify \
      github.com/streadway/simpleuuid \
      github.com/agtorre/gocolorize \
      github.com/robfig/config \
      github.com/robfig/pathtree

# Get revel as gopath + correct versions
RUN git clone https://github.com/revel/revel \
      $GOPATH/src/github.com/revel/revel &&\
    cd $GOPATH/src/github.com/revel/revel && git checkout v0.9.1 && \
    git clone https://github.com/revel/cmd $GOPATH/src/github.com/revel/cmd &&\
    cd $GOPATH/src/github.com/revel/cmd && git checkout a11342d

# Install revel cmd
RUN go install github.com/revel/cmd/revel

# Set workdir
WORKDIR $GOPATH/src/godisc
COPY . $GOPATH/src/godisc

# Install bower components
RUN curl -sL https://deb.nodesource.com/setup_5.x | bash - && \
    apt-get install -y nodejs && \
    npm install -g bower && \
    bower --allow-root install

CMD revel run godisc
