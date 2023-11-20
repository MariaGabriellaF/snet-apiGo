FROM golang:1.21 AS build

WORKDIR /app

COPY . /app
COPY config.toml /app/ 
COPY go.mod ./
COPY go.sum ./

RUN CGO_ENABLED=0 GOOS=linux go build -o api-snet-go main.go
RUN ls -la /app

FROM scratch

WORKDIR /app

COPY --from=build /app/api-snet-go ./
COPY --from=build /app/config.toml ./

EXPOSE 8080

CMD [ "./api-snet-go" ]