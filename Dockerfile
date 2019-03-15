FROM golang

RUN	go get github.com/tbalthazar/onesignal-go

RUN go get github.com/gorilla/mux

WORKDIR /go/src/github.com/heaptracetechnology/microservice-onesignal

ADD . /go/src/github.com/heaptracetechnology/microservice-onesignal

RUN go install github.com/heaptracetechnology/microservice-onesignal

ENTRYPOINT microservice-onesignal

EXPOSE 3000