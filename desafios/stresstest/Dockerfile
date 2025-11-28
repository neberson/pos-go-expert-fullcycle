# syntax=docker/dockerfile:1
FROM golang:1.22-alpine as builder
WORKDIR /app
COPY . .
RUN cd cmd/stresstest && go build -o /stresstest

FROM alpine:latest
COPY --from=builder /stresstest /stresstest
ENTRYPOINT ["/stresstest"]
