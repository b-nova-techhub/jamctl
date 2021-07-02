FROM golang:1.16.5-alpine

WORKDIR $GOPATH/src/github.com/b-nova-openhub/stapagen

COPY . .

RUN go get -d -v ./... \
    && go build -o bin/stapagen cmd/stapagen/main.go \
    && go install -v ./... \
    && chmod +x stapagen.sh

EXPOSE 8080

CMD ["sh", "stapagen.sh"]