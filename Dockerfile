FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o bin/anagram-finder .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/anagram-finder .
COPY --from=builder /app/anagrams.txt .

ENTRYPOINT ["./anagram-finder", "anagrams.txt"]
