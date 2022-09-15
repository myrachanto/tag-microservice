#build stage
FROM golang:alpine AS builder

WORKDIR /app
# COPY go.mod .
# COPY go.sum .
COPY . .
RUN go mod download

RUN go build -o single main.go

#run stage
FROM alpine 
WORKDIR /app
COPY --from=builder /app/single .
COPY app.env .
COPY .env .

EXPOSE 2200
CMD ["/app/single"]
