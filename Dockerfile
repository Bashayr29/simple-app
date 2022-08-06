
FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /api

EXPOSE 8080

CMD [ "/api" ]
