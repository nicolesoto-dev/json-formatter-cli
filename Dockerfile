FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o /app/json-formatter-cli .

FROM gcr.io/distroless/static:nonroot

ARG USER_HOME=/home/nonroot
ENV HOME=${USER_HOME}

WORKDIR /app
COPY --from=builder /app/json-formatter-cli .

USER 65532:65532

ENTRYPOINT ["/app/json-formatter-cli"]