FROM golang:alpine AS build

ADD . /src

RUN apk -U add git && \
    cd /src && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM scratch

COPY --from=build /src/app /

EXPOSE 8080

ENTRYPOINT ["/app"]