FROM golang:alpine

# Set the working directory to /app
WORKDIR /go/src/app

COPY . .

RUN apk add --no-cache git \
    && go get github.com/gorilla/mux \
    && go get github.com/dmirou/tribonacci \
    && apk del git

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["app"]