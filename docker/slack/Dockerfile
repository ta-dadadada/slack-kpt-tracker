FROM golang:latest as builder
WORKDIR /repos
COPY . .
WORKDIR /repos/src/slack
RUN go get github.com/nlopes/slack
RUN GOOS=linux CGO_ENABLED=0 go build

# runtime image
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /repos/src/slack/slack /slack
COPY --from=builder /repos/docker/slack/entrypoint.sh .
RUN pwd && ls -ltr
#ENTRYPOINT ["SLACK_TOKEN=``","./slack"]
ENTRYPOINT ["./entrypoint.sh"]
