FROM golang:1.19-alpine3.16 as build
WORKDIR /app 
COPY . . 
RUN go build -o main . 

FROM alpine:3.14
WORKDIR /app
COPY --from=build /app/main .
COPY config.yml .
COPY wait-for.sh .
RUN chmod +x wait-for.sh
ENTRYPOINT [ "/app/main" ]
EXPOSE 8081
CMD ["serve"]
