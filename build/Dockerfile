# build the gookla speedtest app
FROM golang:alpine as builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build

COPY go.mod .
COPY go.sum .
COPY . .
RUN go build -o /goookla-influx

# final image
FROM alpine:latest
RUN apk --no-cache add wget  && \
    wget https://bintray.com/ookla/download/download_file?file_path=ookla-speedtest-1.0.0-x86_64-linux.tgz -O /var/tmp/speedtest.tar.gz && \
    tar xf /var/tmp/speedtest.tar.gz -C /var/tmp && \
    mv /var/tmp/speedtest /usr/local/bin && \
    rm /var/tmp/speedtest.tar.gz
RUN addgroup -S speedtest && adduser -S speedtest -G speedtest
COPY --from=builder /goookla-influx /goookla-influx
RUN chmod +x /goookla-influx
USER speedtest
ENTRYPOINT ["/goookla-influx"]