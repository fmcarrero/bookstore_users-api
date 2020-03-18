# build stage
FROM golang as builder

# Configure the repo url so we can configure our work directory:
ENV REPO_URL=github.com/fmcarrero/bookstore_users-api

# Setup out $GOPATH
ENV GOPATH=/app

ENV GO111MODULE=on

ENV APP_PATH=$GOPATH/src/$REPO_URL

# Copy the entire source code from the current directory to $WORKPATH
ENV WORKPATH=$APP_PATH/src

COPY go.mod $WORKPATH
COPY go.sum $WORKPATH
COPY src $WORKPATH

WORKDIR $WORKPATH


RUN go mod download

RUN go build -o users-api .

# final stages
FROM alpine:3.7

COPY --from=builder $APP_PATH/users-api .
ENTRYPOINT ["./users-api"]