
# # production

# FROM golang:1.21.4 AS builder
# WORKDIR /app
# COPY . .
# RUN CGO_ENABLED=0 go build -o /hello-world

# FROM scratch
# COPY --from=builder /hello-world /hello-world
# EXPOSE 8080
# CMD ["/hello-world"]

# development

FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o /app/bin/gizoffer ./cmd/app/main.go

FROM alpine:latest
RUN apk add --no-cache bash
RUN apk add go
COPY --from=builder /app/bin/gizoffer /gizoffer
EXPOSE 80
CMD ["/gizoffer"]
