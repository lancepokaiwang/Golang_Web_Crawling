FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY errors ./errors
COPY server ./server
COPY proto ./proto

RUN go mod download

COPY . /app
RUN go build -o /docker-gs-ping

EXPOSE 8000

CMD [ "/docker-gs-ping" ]