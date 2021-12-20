FROM golang:1.17

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN ls

RUN go mod download
RUN go build -o messanger ./cmd/centrifuge_lib/main.go

CMD ["./messanger"]