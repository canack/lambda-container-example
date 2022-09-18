FROM golang:1.19 as builder

COPY . /app
WORKDIR /app

RUN go mod tidy && \
    CGO_ENABLED=0 go build -ldflags "-s -w" -o /app/text-to-base64

FROM gcr.io/distroless/static

COPY --from=builder /app/text-to-base64 /app/text-to-base64

CMD ["/app/text-to-base64"]