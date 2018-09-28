FROM golang:latest
WORKDIR /Users/kayrules/Projects/go/gopath/src/github.com/kayrules/jobads-api

RUN go get -d -v github.com/facebookgo/grace/gracehttp && \
    go get -d -v github.com/labstack/echo && \
    go get -d -v github.com/labstack/echo/middleware && \
    go get -d -v github.com/labstack/gommon/log && \
    go get -d -v github.com/swaggo/echo-swagger && \
    go get -d -v github.com/asaskevich/govalidator && \
    go get -d -v github.com/globalsign/mgo/bson

ADD ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /Users/kayrules/Projects/go/gopath/src/github.com/kayrules/jobads-api/main .
CMD ["./main"]