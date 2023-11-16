FROM golang:1.21 AS build

WORKDIR /app

COPY . /app
COPY go.mod ./

RUN CGO_ENABLED=0 GOOS=linux go build -o api-snet-go main.go

FROM scratch

WORKDIR /app

COPY --from=build /app/api-snet-go ./

EXPOSE 8080

CMD [ "./api-snet-go" ]