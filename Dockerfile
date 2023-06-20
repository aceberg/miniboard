FROM golang:alpine AS builder

RUN apk add build-base
COPY cmd /src/cmd
COPY internal /src/internal
COPY go.mod /src/
COPY go.sum /src/
RUN cd /src/cmd/miniboard/ && CGO_ENABLED=0 go build -o /miniboard .


FROM scratch

WORKDIR /data/miniboard
WORKDIR /app

COPY --from=builder /miniboard /app/

ENTRYPOINT ["./miniboard"]