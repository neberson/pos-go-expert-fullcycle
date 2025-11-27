FROM golang:1.24 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ratelimit ./cmd/ratelimit/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/ratelimit .
ENV ratelimit_API_KEY=${ratelimit_API_KEY}
ENTRYPOINT [ "./ratelimit" ]