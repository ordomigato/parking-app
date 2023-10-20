FROM golang:latest
RUN go install github.com/cosmtrek/air@latest
WORKDIR /app
COPY . .
RUN go build -o main .
COPY ./store/migrations ./store/migrations
ENTRYPOINT ["air"]