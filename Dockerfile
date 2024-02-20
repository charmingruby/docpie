# build step
FROM golang:1.18 AS build

WORKDIR /app

COPY . /app/

RUN CGO_ENABLED=0 GOOS=linux go build -o owler ./cmd/server/main.go

# image step
FROM scratch

WORKDIR /app

COPY --from=build /app/owler /app/

EXPOSE 8000

CMD [ "./owler" ]