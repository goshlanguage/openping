FROM golang:1.12.4-alpine3.9

COPY . /go/src/github.com/ryanhartje/openping/
WORKDIR /go/src/github.com/ryanhartje/openping/
RUN go build -o /go/bin/openping ./cmd/openping/openping.go
RUN GOARCH=arm64 go build -o /go/bin/openping-arm64 ./cmd/openping/openping.go


ENTRYPOINT ["/go/bin/openping"]

