# build step
FROM golang:1.18 AS build

WORKDIR /app

COPY . /app/

RUN CGO_ENABLED=0 GOOS=linux go build -o upl ./cmd/server/main.go

# image step
FROM scratch

WORKDIR /app

COPY --from=build /app/upl /app/

EXPOSE 8000

CMD [ "./upl" ]