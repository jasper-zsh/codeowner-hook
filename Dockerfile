FROM golang:alpine AS builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .
COPY conf /app/conf
RUN go build -o /app/codeowner-hook main.go

FROM --platform=linux/amd64 alpine:latest

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata

ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app /app

CMD ["/app/codeowner-hook"]
