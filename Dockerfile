# === Builder ===
FROM golang:1.25 AS builder
WORKDIR /app
ADD . /app

RUN go mod download
RUN go build -o influxmcstat

# === Runner ===
FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=builder /app/influxmcstat /

CMD ["/influxmcstat"]