FROM golang:1.24 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o weather ./cmd/server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/weather .
ENV WEATHER_API_KEY=${WEATHER_API_KEY}
ENTRYPOINT [ "./weather" ]