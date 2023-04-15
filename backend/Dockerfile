# syntax=docker/dockerfile:1

FROM golang:1.20

WORKDIR /app

RUN apt-get update && \
    apt-get install -y apt-transport-https ca-certificates curl gnupg2 software-properties-common

RUN curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - && \
    echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | tee /etc/apt/sources.list.d/kubernetes.list

RUN apt-get update && \
    apt-get install -y kubectl

RUN apt-get install -y jc

COPY . .
RUN go mod download

RUN env GOOS=linux GOARCH=amd64 go build -o iptables-viz-backend /app

EXPOSE 8080

CMD ["./iptables-viz-backend", "-platform", "kubernetes"]
