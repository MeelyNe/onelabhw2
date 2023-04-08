FROM golang:1.19.8-alpine3.17 as build

WORKDIR /app/onelab

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 go build -o ./bin/app ./cmd/main.go

FROM alpine
COPY --from=build /app/onelab/bin/app /app/onelab/bin/app
COPY --from=build /app/onelab/.env /app/onelab/bin/.env

RUN chmod +x /app/onelab/bin/app

EXPOSE 8080
ENTRYPOINT ["/app/onelab/bin/app"]