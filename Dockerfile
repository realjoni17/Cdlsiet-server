FROM golang:1.20.5-alpine3.18
WORKDIR /test
COPY . /test/
RUN go build /test
EXPOSE 8080
ENTRYPOINT [ "./Cdlsiet" ]