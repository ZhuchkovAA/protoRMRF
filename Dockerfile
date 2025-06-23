# Используем многоэтапную сборку
FROM golang:1.24 as builder

# Устанавливаем системные зависимости
RUN apt-get update && apt-get install -y \
    unzip \
    curl \
    git \
    # Для Node.js
    ca-certificates \
    gnupg

# Устанавливаем Node.js (LTS версию)
RUN mkdir -p /etc/apt/keyrings \
    && curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key | gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg \
    && echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_20.x nodistro main" > /etc/apt/sources.list.d/nodesource.list \
    && apt-get update && apt-get install -y nodejs

# Устанавливаем protoc
ENV PROTOC_VERSION=25.1
RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip \
    && unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /usr/local \
    && rm protoc-${PROTOC_VERSION}-linux-x86_64.zip

# Устанавливаем Go-плагины
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Копируем package.json отдельно для кэширования
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install

# Копируем остальные файлы
COPY . .

# Делаем скрипт исполняемым
RUN chmod +x ./protoc-gen.sh

# Добавляем пути к плагинам в PATH
ENV PATH="/go/bin:/app/node_modules/.bin:$PATH"

CMD ["./protoc-gen.sh"]