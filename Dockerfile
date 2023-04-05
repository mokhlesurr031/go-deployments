FROM golang:1.19-alpine3.16
WORKDIR /app 
COPY . . 
RUN go build -o main . 
CMD ["./main", "serve"]
