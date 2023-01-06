FROM golang:1.19-bullseye AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /out/gateway ./gateway/cmd/server
RUN go build -o /out/auth ./auth/cmd/server
RUN go build -o /out/user ./user/cmd/server
RUN go build -o /out/thread ./thread/cmd/server

FROM gcr.io/distroless/base-debian11

WORKDIR /
COPY --from=build /out/* .

EXPOSE 8080

USER nonroot:nonroot
