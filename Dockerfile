FROM golang:latest AS build

WORKDIR /go/src/github.com/heavywombat/dummy-repo
COPY . .
RUN go build \
    -trimpath \
    -ldflags "-s -w -extldflags '-static'" \
    -o /tmp/helloworld \
    main.go

FROM scratch
COPY --from=build /tmp/helloworld ./helloworld
ENTRYPOINT ["./helloworld"]
EXPOSE 8080
