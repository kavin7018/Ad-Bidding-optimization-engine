FROM golang:1.23-alpine

WORKDIR /usr/src/app/

COPY . .

RUN go build -a -v .
EXPOSE 2112
ENTRYPOINT [ "./app" ]