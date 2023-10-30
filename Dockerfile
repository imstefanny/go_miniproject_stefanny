FROM golang:1.21.0-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go_miniproject_stefanny

EXPOSE 8000

CMD [ "/go_miniproject_stefanny" ]
