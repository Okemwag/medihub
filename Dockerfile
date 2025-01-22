FROM golang:1.23.4-alpine3.21 AS builder
RUN apk add --no-cache gcc musl-dev git build-base pkgconfig libsodium-dev

ENV GOOS=linux

WORKDIR /etc/medihub/

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
  go build -o medihub cmd/medihub/main.go

FROM alpine:3.21
RUN apk add libsodium-dev
COPY --from=builder /etc/medihub/medihub .
COPY migrations migrations


ARG GIT_COMMIT
ENV GIT_COMMIT=$GIT_COMMIT

CMD ["./medihub"]
