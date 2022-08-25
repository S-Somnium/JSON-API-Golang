FROM golang:latest  
WORKDIR /app
COPY . .
RUN go mod download && go mod verify

RUN go build -o ./server
EXPOSE 8080
CMD ["./server"]



