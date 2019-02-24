FROM golang:1.11
WORKDIR /app
# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV SRC_DIR=/go/src/verisart
ENV GO111MODULE=on
# Add the source code:
ADD . $SRC_DIR
# Build it:
RUN cd $SRC_DIR; go build -o verisart; cp verisart /app/
ENTRYPOINT ["./verisart"]
