FROM golang:1.23.6 AS builder

ARG ARCH=amd64
ARG GO111MODULE=on

WORKDIR /code

COPY go.mod go.sum /code/
COPY . /code

RUN CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} go build \
        -o $GOPATH/bin/cipher \
		-gcflags "all=-trimpath=$GOPATH" \
        ./cmd/cipher/.

FROM alpine:3.15

COPY --from=builder /go/bin/cipher /bin/cipher
RUN chmod +x /bin/cipher

ENTRYPOINT ["/bin/cipher"]