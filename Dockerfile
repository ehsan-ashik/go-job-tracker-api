FROM golang:1.23.4

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go mod tidy

CMD ["air", "-c", ".air.toml"]