FROM golang:1 AS build

ENV CGO_ENABLED=0

RUN update-ca-certificates

RUN apt-get update && apt-get install -y wamerican

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build ./... && go build -o cbojar.telegrambot


FROM scratch

ENV TELEGRAM_BOT_KEY=''

COPY --from=build /usr/share/dict/american-english /usr/share/dict/american-english

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app

COPY --from=build --chown=1001 /build/cbojar.telegrambot ./cbojar.telegrambot

USER 1001

ENTRYPOINT ["/app/cbojar.telegrambot"]
