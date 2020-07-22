FROM golang:1.14.6 as builder

WORKDIR $GOPATH/src/github.com/apuigsech/circuit/models/namespace
COPY . .
RUN go build -o bin/namespace .

FROM: alpine

COPY --from=builder /bin/namespaces /bin/namespaces

ENTRYPOINT ["/bin/namespaces"]