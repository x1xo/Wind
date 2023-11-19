FROM golang:1.21.3-alpine3.18

WORKDIR /app

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /wind

EXPOSE 3000

# Run
CMD ["/wind"]