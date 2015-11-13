FROM golang:latest

RUN apt-get install git

# Get revel dependencies
RUN go get github.com/revel/cmd/revel && \
    go get github.com/revel/revel

# Get correct versions
RUN cd $GOPATH/src/github.com/revel/revel && git checkout v0.9.1 && \
    cd $GOPATH/src/github.com/revel/cmd && git checkout a11342d

# Create directory and copy code
WORKDIR $GOPATH/src/godisc
COPY . $GOPATH/src/godisc

CMD revel run godisc
