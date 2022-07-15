#BUILD
FROM golang:1.18 as BUILD
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build main.go 

#RUN
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=BUILD /app/main ./main
COPY public ./public
CMD ["./main"]