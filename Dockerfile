FROM golang:latest AS build

COPY . .
RUN go build \
    -trimpath \
    -ldflags "-s -w -extldflags '-static'" \
    -o /usr/local/bin/helloworld \
    main.go

FROM alpine:latest
COPY --from=build /usr/local/bin/helloworld /usr/local/bin/helloworld
ENTRYPOINT [ "/usr/local/bin/helloworld" ]
EXPOSE 8080
