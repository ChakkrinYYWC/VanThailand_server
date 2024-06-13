FROM golang:1.22.2

WORKDIR /app

COPY . .

RUN go get

CMD ["go", "run", "."]

EXPOSE 8080