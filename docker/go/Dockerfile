FROM library/golang:1.16.2-alpine3.13

RUN apk add make gcc libc-dev bash

RUN go get -u github.com/pressly/goose/cmd/goose
RUN go get -u github.com/guiyomh/charlatan
RUN go get -u github.com/swaggo/swag/cmd/swag

COPY ./wait-for-it.sh /usr/bin/wait-for-it.sh

ENTRYPOINT ["wait-for-it.sh", "-s" , "-t", "120", "mysql:3306", "--"]
CMD ["echo", "MySQL is available now !"]
CMD ["tail", "-f", "/dev/null"]
