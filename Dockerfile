
## Build
FROM golang:buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /api

## Deploy
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /api /api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/api"]