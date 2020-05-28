FROM node:8.12.0

WORKDIR /build

COPY ./client .

WORKDIR /build/scale

RUN npm install -g npm-install-peers

RUN npm install

RUN npm run build

FROM golang:1.13.5-alpine3.10 AS builder

WORKDIR /build
#RUN adduser -u 10001 -D app-runner

COPY ./server/go.mod .
COPY ./server/go.sum .
RUN go mod download

COPY ./server .
RUN ls
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o ssr .

FROM alpine:3.10 AS final

WORKDIR /app
COPY --from=builder /build/conf/ /app/conf/
COPY --from=builder /build/static/ /app/static/
COPY --from=builder /build/views/ /app/views/
COPY --from=0 /build/scale/dist/ /app/dist/
COPY --from=builder /build/ssr /app/

#USER app-runner
ENTRYPOINT ["/app/ssr"]