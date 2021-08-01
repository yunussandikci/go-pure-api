FROM golang:1.16-alpine as builder
ADD .. /src
RUN cd /src && go build -ldflags="-s -w" -o go-pure-api

FROM alpine
WORKDIR /app
COPY --from=builder /src/go-pure-api /app/
ENTRYPOINT ./go-pure-api