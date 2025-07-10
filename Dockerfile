FROM golang:1.24-alpine AS builder

WORKDIR /src
COPY . .

RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o /app/json-formatter-cli

FROM gcr.io/distroless/static:nonroot

ARG USER_HOME=/home/nonroot
ENV HOME=${USER_HOME}

WORKDIR /
COPY --from=builder /app/ .

USER 65532:65532

ENTRYPOINT ["./app/json-formatter-cli"]