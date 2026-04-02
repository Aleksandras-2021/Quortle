FROM golang:1.25

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o app


EXPOSE 443 80 8080 5432

CMD ["./app"]