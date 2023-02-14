FROM golang:1.19-alpine
ENV APP_ROOT=/app

WORKDIR ${APP_ROOT}

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./application ./application.go

EXPOSE 8080

CMD ["./application", "app"]
