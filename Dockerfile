FROM golang:1.16.3 AS base
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* .
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

FROM base as build
RUN go build -o /bin/bot ./cmd/bot.main.go

FROM scratch as bin
COPY --from=build /bin/bot /
