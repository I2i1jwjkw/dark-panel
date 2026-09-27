FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/dark-panel ./cmd
FROM alpine:3.21
RUN adduser -D -H app
COPY --from=build /out/dark-panel /usr/local/bin/dark-panel
USER app
EXPOSE 8080
ENTRYPOINT ["dark-panel"]
