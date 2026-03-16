FROM golang:tip-alpine3.22 
WORKDIR /app/go-app
COPY . .
ENV PORT=8020
EXPOSE $PORT 
CMD ["go","run","main.go"]