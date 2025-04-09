FROM golang:1.23 AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o deposit .

FROM alpine:latest

WORKDIR /app

COPY --from=build /src/deposit .

EXPOSE 8081

CMD ["./deposit"]