FROM golang:1.13.5-alpine3.10 AS builder

WORKDIR /build
#RUN adduser -u 10001 -D app-runner

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN ls
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o ssr .

FROM alpine:3.10 AS final

WORKDIR /app
COPY --from=builder /build/conf/ /app/conf/
COPY --from=builder /build/static/ /app/static/
COPY --from=builder /build/views/ /app/views/
COPY --from=builder /build/ssr /app/
RUN ls

#COPY --from=builder /build/config /app/config
#COPY --from=builder /etc/passwd /etc/passwd
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

#USER app-runner
ENTRYPOINT ["/app/ssr"]