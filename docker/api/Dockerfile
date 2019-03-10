FROM golang:latest as builder
WORKDIR /repos
COPY . .
WORKDIR /repos/src/api
RUN go get github.com/jinzhu/gorm github.com/go-sql-driver/mysql github.com/nlopes/slack github.com/gin-gonic/gin
RUN GOOS=linux CGO_ENABLED=0 go build

# runtime image
FROM alpine
RUN apk add --no-cache ca-certificates\
    mysql-client
COPY --from=builder /repos/src/api/api /api
EXPOSE 8080
ENTRYPOINT ["./api"]