# 開発用ステージ
FROM golang:1.21-alpine AS development

RUN apk add --no-cache git make
WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.40.4

CMD ["air", "-c", ".air.toml"]

# 本番用ステージ
FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLE=0 GOOS=linux go build -o /go/bin/app

# 最終的なイメージ
FROM alpine:latest AS production

RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/

COPY --from=build /go/bin/app .

EXPOSE 8080
CMD ["./app"]
