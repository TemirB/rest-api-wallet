FROM golang:1.23.5

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o /app/rest-api-wallet cmd/rest-api/main.go

EXPOSE 8080

CMD ["/app/rest-api-wallet"]