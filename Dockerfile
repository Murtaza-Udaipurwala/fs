FROM golang:alpine as builder

RUN apk add --update --no-cache ca-certificates git

COPY . ./fs

RUN cd fs && go build

ENTRYPOINT ["/go/fs/fs"]

FROM alpine:latest

COPY --from=builder /go/fs/fs main

ENTRYPOINT [ "./main" ]
