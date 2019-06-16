FROM golang:1.12.6-alpine3.9
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main.exe .
CMD ["/app/main.exe"]