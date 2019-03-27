# Build with modules
FROM golang:1.12.1-alpine as builder
RUN apk add --no-cache ca-certificates curl git build-base

# Get dependencies first for docker caching
WORKDIR /fhir
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy source
COPY . .
ARG GIT_COMMIT=dev

# Build fhir-server
WORKDIR /fhir/fhir-server
RUN go build -ldflags "-X main.gitCommit=$GIT_COMMIT"

# Copy to light-weight runtime image
FROM alpine:3.7
RUN apk add --no-cache ca-certificates tini
COPY --from=builder /fhir/fhir-server/fhir-server /
COPY --from=builder /fhir/fhir-server/config/ /config
COPY --from=builder /fhir/conformance/ /conformance

ENV MONGODB_URI mongodb://fhir-mongo:27017/?replicaSet=rs0
CMD ["sh", "-c", "/fhir-server -port 3001 -disableSearchTotals -enableXML -databaseName fhir -mongodbURI $MONGODB_URI"]
