FROM golang:1.23-alpine AS builder

WORKDIR /
COPY . .
RUN apk --no-cache add --no-check-certificate ca-certificates && update-ca-certificates


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/skylink-ingest 

FROM scratch
# Copy our static executable.
COPY --from=builder /etc/ssl/ /etc/ssl/ 
COPY --from=builder /skylink-ingest /skylink-ingest

ENTRYPOINT ["./skylink-ingest"]