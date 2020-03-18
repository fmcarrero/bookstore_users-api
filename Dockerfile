# build stage
FROM golang as builder

ENV GO111MODULE=on

WORKDIR /go/src/github.com/fmcarrero/bookstore_users-api/


COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .


RUN  CGO_ENABLED=0 go build  -o users_api

# final stage
FROM alpine:3.7

COPY --from=builder /go/src/github.com/fmcarrero/bookstore_users-api/users_api .

# Expose port 8081 to the world:
EXPOSE 8081

CMD ["./users_api"]
