FROM golang:1.22.0

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o bin

EXPOSE 8080

ENTRYPOINT [ "/app/bin" ]