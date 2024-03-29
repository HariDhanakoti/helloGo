FROM 438991312594.dkr.ecr.ap-southeast-2.amazonaws.com/golang:alpine

WORKDIR /go/src/app

COPY . .

RUN apk add git

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["app"]