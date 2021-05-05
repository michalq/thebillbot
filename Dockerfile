FROM golang:1.14-stretch as modules

# Force go modules.
ENV GO111MODULE=on
# GOPATH must be empty with go modules.
ENV GOPATH=

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM modules as builder

COPY . .

RUN go build -o /app/bot ./cmd/bot/main.go

FROM debian:stretch

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update \
  && apt-get install -y --no-install-recommends \
    ca-certificates \
  && apt-get clean -y \
  && apt-get autoremove -y \
  && rm -rf /tmp/* /var/tmp/* \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/bot .
ENTRYPOINT ["./bot"]