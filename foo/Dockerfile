FROM ghcr.io/shipwright-io/shipwright-samples/golang:1.22 AS build

COPY ../common/main.go .
ENV CGO_ENABLED=0
RUN go build \
    -ldflags "-s -w -extldflags '-static'" \
    -o /tmp/helloworld \
    main.go

FROM scratch
COPY --from=build /tmp/helloworld ./helloworld
ENV NAME=foo
ENTRYPOINT [ "./helloworld" ]
EXPOSE 8080
