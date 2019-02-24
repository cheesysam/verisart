FROM golang:1.11
WORKDIR /app
ENV SRC_DIR=/go/src/verisart
ENV GO111MODULE=on
ADD . $SRC_DIR
RUN cd $SRC_DIR; go build -o verisart; cp verisart /app/
ENTRYPOINT ["./verisart"]
