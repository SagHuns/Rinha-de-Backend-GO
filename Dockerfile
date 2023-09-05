FROM golang:1.21.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/main .

EXPOSE 8080
EXPOSE 5432

CMD ["/app/main"]