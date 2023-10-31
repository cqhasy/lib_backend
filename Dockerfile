FROM golang:1.21
WORKDIR /app
#COPY . .
#RUN go build -o app
#EXPOSE 8080
#CMD ["./app"]
CMD ["sh", "-c", "while true; do echo 'Running...'; sleep 5; done"]