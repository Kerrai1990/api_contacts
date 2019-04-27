FROM golang:1.11-alpine3.8 as builder

WORKDIR /usr/src/app

RUN apk add --no-cache ca-certificates gcc git libc-dev \
  && go get -u golang.org/x/lint/golint

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"'

FROM scratch

COPY --from=builder /go/bin/api_contacts /api_contacts
COPY --from=builder /etc/ssl/certs /etc/ssl/certs

ENTRYPOINT ["/api_contacts"] 