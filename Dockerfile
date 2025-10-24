# Etapa 1 - build
FROM golang:1.23-alpine AS builder
WORKDIR /app


# Copia os arquivos de dependências primeiro (para cache eficiente)
COPY go.mod go.sum ./
RUN go mod download


# Copia o restante do código e constrói o binário
COPY . .
RUN go build -o main ./cmd/server


# Erapa 2 - Runtime (imagem leve)
FROM alpine:latest
WORKDIR /root/


# Copia o binario do estágio anterior
COPY --from=builder /app/main .


# Expoe a porta 
EXPOSE 8080


# COmando de inicialização
CMD ["./main"]