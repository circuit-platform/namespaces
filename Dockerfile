FROM golang:1.14.6 as builder

WORKDIR $GOPATH/src/github.com/apuigsech/circuit/models/namespaces
COPY . .
RUN go build -o /bin/namespaces .

FROM alpine

COPY --from=builder /bin/namespaces /bin/namespaces

ENTRYPOINT ["/bin/namespaces"]