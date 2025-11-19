FROM golang:1.24 AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o weather-a ./cmd/server-a/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/weather-a .
ENV WEATHER_API_KEY=${WEATHER_API_KEY}
ENTRYPOINT [ "./weather-a" ]