# Build stage
FROM golang:1.20-alpine3.17 AS build
WORKDIR /app
COPY . .
RUN go build -o main .

# Run stage  
FROM scratch
WORKDIR /app
COPY ./store/migrations ./store/migrations
COPY --from=build /app/main .
ENTRYPOINT ["/app/main"]