# Используем официальный образ Go
FROM golang:1.22-alpine AS builder

# Устанавливаем необходимые инструменты
RUN apk add --no-cache \
    git \
    make \
    protobuf \
    protobuf-dev

# Устанавливаем protoc-gen-go и protoc-gen-go-grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Добавляем PATH для protoc-gen-* инструментов
ENV PATH $PATH:/go/bin

# Рабочая директория
WORKDIR /app

# Копируем файлы проекта
COPY . .

# Команда для генерации кода из .proto файлов
RUN mkdir -p gen/go && \
    protoc --proto_path=proto \
           --go_out=gen/go \
           --go-grpc_out=gen/go \
           proto/**/*.proto

# Финальный образ (для работы с protoc)
FROM golang:1.22-alpine

# Устанавливаем protoc и необходимые инструменты
RUN apk add --no-cache \
    protobuf \
    protobuf-dev

# Копируем бинарники protoc-gen-go и protoc-gen-go-grpc
COPY --from=builder /go/bin/protoc-gen-go /go/bin/protoc-gen-go
COPY --from=builder /go/bin/protoc-gen-go-grpc /go/bin/protoc-gen-go-grpc

# Добавляем PATH для protoc-gen-* инструментов
ENV PATH $PATH:/go/bin

WORKDIR /app

# Копируем сгенерированные файлы из builder
COPY --from=builder /app/gen/go ./gen/go

# Копируем исходные .proto файлы
COPY proto ./proto