FROM golang:1.23-alpine AS builder

WORKDIR /
COPY . .

RUN go build ./cmd/ingest/

FROM scratch
# Copy our static executable.
COPY --from=builder /ingest /ingest
# Run the hello binary.
ENTRYPOINT ["/ingest"]