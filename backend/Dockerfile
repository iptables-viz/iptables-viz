# syntax=docker/dockerfile:1

FROM golang:1.20

WORKDIR /app

RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl && \
    chmod +x kubectl && \
    mv kubectl /usr/local/bin/

RUN curl -LO https://github.com/kellyjonbrazil/jc/releases/download/v1.23.1/jc-1.23.1-linux-x86_64.tar.gz && \
    tar -xvf jc-1.23.1-linux-x86_64.tar.gz && \
    chmod +x jc && \
    mv jc /usr/local/bin/

COPY . .
RUN go mod download

RUN env GOOS=linux GOARCH=amd64 go build -o iptables-viz-backend /app

EXPOSE 8080

CMD ["./iptables-viz-backend", "-platform", "kubernetes"]
