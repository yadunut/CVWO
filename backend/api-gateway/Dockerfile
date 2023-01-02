FROM golang:1.19-bullseye AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

RUN go build -o api-gateway cmd/server/main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /
COPY --from=build /app/api-gateway /api-gateway

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/api-gateway"]
