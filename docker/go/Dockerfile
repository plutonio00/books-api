FROM library/golang:1.16.2-alpine3.13

RUN apk add make gcc libc-dev

RUN go get -u github.com/pressly/goose/cmd/goose
RUN go get -u github.com/guiyomh/charlatan
RUN go get -u github.com/swaggo/swag/cmd/swag

