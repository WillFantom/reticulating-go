ARG GO_VERSION=1.16
FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /src
COPY ./go.mod ./go.mod
RUN go mod download
COPY ./cmd/reticulate .
RUN CGO_ENABLED=0 go build -ldflags -o reticulate .

FROM scratch

LABEL maintainer="Will Fantom <willf@ntom.dev>"

WORKDIR /sims-loading-messages
COPY --from=builder /src/reticulate .

EXPOSE 8080
ENTRYPOINT [ "./reticulate", "api", "--port", "8080" ]
