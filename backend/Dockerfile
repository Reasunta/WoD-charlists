FROM golang:1.24.2 as builder
WORKDIR /backend/app
COPY go.mod go.sum /backend/app
RUN go mod download
COPY . /backend/app

RUN CGO_ENABLED=0 go build -o wod-backend app.go

FROM alpine:3.21
WORKDIR /
COPY --from=builder /backend/app/wod-backend .