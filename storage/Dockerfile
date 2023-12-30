FROM golang:latest as builder
WORKDIR /app
COPY ../storage/go.mod .
COPY ../storage/go.sum .
ENV GOPROXY=https://goproxy.io,direct
RUN go mod download
COPY ../storage .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o storage .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/storage .
EXPOSE 9091
RUN mkdir -p "/opt/mani/data"
CMD ["./storage"]
