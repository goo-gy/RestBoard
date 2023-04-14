FROM golang:latest

WORKDIR /app
COPY . .

RUN go build main.go

EXPOSE 8000
CMD ["./main"]