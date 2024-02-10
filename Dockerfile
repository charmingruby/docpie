# build step
FROM golang:1.18 AS build

WORKDIR /app

COPY . /app/

RUN CGO_ENABLED=0 GOOS=linux go build -o make-it-survey ./cmd/server/main.go

# image step
FROM scratch

WORKDIR /app

COPY --from=build /app/make-it-survey /app/

EXPOSE 8000

CMD [ "./make-it-survey" ]