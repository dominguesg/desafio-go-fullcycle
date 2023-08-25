FROM golang:1.17

WORKDIR /app

COPY . /app

RUN go get github.com/go-chi/chi
RUN go get github.com/go-sql-driver/mysql

CMD ["go", "run", "api/main.go"]
