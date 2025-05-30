FROM golang:1.22.12-alpine3.20 AS builder

ENV GOCACHE=/root/.cache/go-build

WORKDIR /appsource

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./go.mod ./


RUN go mod tidy

RUN --mount=type=cache,target="/root/.cache/go-build" \
   go build -o applinux ./cmd/main.go


FROM alpine:3.20

WORKDIR /myapp

COPY --from=builder /appsource/applinux ./

EXPOSE 6080

CMD [ "/myapp/applinux" ]