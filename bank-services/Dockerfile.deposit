FROM golang:1.24 AS build


RUN apt update && apt install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
ENV PATH="$PATH:$(go env GOPATH)/bin"
ENV CGO_ENABLED=0

WORKDIR /src

COPY go.mod go.sum /src/
COPY proto /src/proto
COPY deposit /src/deposit
COPY pkg /src/pkg

RUN go mod tidy

RUN  protoc --go_out=.      --go_opt=paths=source_relative \
           --go-grpc_out=. --go-grpc_opt=paths=source_relative \
           proto/deposit/deposit.proto


RUN GOOS=linux go build deposit/main.go


FROM alpine:3.20

WORKDIR /app

COPY --from=build /src/main deposit

EXPOSE 8081

CMD ["./deposit"]