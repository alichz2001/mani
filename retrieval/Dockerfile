FROM golang:latest as builder
WORKDIR /app
COPY ../retrieval/go.mod .
COPY ../retrieval/go.sum .
ENV GOPROXY=https://goproxy.io,direct
RUN go mod download
COPY ../retrieval .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gateway .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/gateway .
EXPOSE 9090
CMD ["./gateway"]