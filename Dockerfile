ARG GO_VERSION=1.16
FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /src
ARG APP_VERSION=
RUN test -n "$APP_VERSION"
COPY ./go.mod ./go.mod
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-X github.com/willfantom/loading-messages/version.appVersion=${APP_VERSION}" -o loading-messages .

FROM scratch

LABEL maintainer="Will Fantom <willf@ntom.dev>"

WORKDIR /sims-loading-messages
COPY --from=builder /src/loading-messages .
COPY --from=builder /src/gamedata.json .
ENTRYPOINT [ "./loading-messages" ]
