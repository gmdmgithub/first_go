############################
# STEP 1 build executable binary
############################
FROM golang:1.12.6-alpine3.9

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN mkdir /app
ADD . /app
WORKDIR /app


# Download all the dependencies
# Using go get.
RUN go get -d -v
# # https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
# RUN go get -d -v ./...

# Install the package
# RUN go install -v ./...
EXPOSE 8083
RUN go build -o main .
CMD ["/app/main -docker=true"]