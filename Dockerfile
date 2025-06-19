# docker build -t protoc-gen .
FROM golang:1.24 as builder

RUN apt-get update && apt-get install -y unzip curl git

ENV PROTOC_VERSION=25.1
RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip \
    && unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /usr/local \
    && rm protoc-${PROTOC_VERSION}-linux-x86_64.zip

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ENV PATH="/go/bin:$PATH"

WORKDIR /app
COPY . .

# Добавляем права на исполнение
RUN chmod +x ./protoc-gen.sh

CMD ["./protoc-gen.sh"]
