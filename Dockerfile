FROM golang

RUN	go get github.com/tbalthazar/onesignal-go

RUN go get github.com/gorilla/mux

WORKDIR /go/src/github.com/oms-services/onesignal

ADD . /go/src/github.com/oms-services/onesignal

RUN go install github.com/oms-services/onesignal

ENTRYPOINT onesignal

EXPOSE 3000