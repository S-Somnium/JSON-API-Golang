FROM golang:latest  
WORKDIR /app
COPY . .
RUN go mod download
WORKDIR /app/api
RUN go build -o ./server
RUN ./server
EXPOSE 8080


