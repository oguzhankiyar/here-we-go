FROM golang:1.16.5-alpine as builder

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o web-sample cmd/web-sample/main.go

FROM scratch as sc

WORKDIR /app

COPY --from=builder /app/web-sample ./
COPY --from=builder /app/configs ./configs

ENTRYPOINT [ "./web-sample" ]