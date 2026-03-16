FROM golang:tip-alpine3.22 
COPY . .
CMD ["go","run","main.go"]