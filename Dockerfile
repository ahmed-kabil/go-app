FROM golang:tip-alpine3.22 AS builder
WORKDIR /app/
COPY . .
ARG PORT=8020
ENV PORT=${PORT}
EXPOSE $PORT 
RUN go build -o go-app main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/go-app .
COPY --from=builder /app/static ./static
CMD ["./go-app"]