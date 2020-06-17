# Fase build
FROM golang:1.13.11-alpine AS builder

WORKDIR /go

COPY server /go/src/server

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -ldflags="-w -s" server
RUN CGO_ENABLED=0 go test server
RUN chmod +x /go/bin/server

# Fase imagem final
FROM scratch

COPY --from=builder /go/bin/server /go/bin/server

EXPOSE 8000

ENTRYPOINT ["/go/bin/server"]
