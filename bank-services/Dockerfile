FROM golang:1.24 AS build


RUN apt update && apt install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
ENV PATH="$PATH:$(go env GOPATH)/bin"
ENV CGO_ENABLED=0

WORKDIR /src

COPY go.mod go.sum ./
COPY proto ./proto
COPY Makefile ./Makefile
COPY mail ./mail
COPY pkg ./pkg


RUN go mod tidy

RUN make protobuf

RUN go build mail/main.go

FROM alpine:3.20

WORKDIR /app

COPY --from=build /src/main mail

EXPOSE 8081

CMD ["./mail"]