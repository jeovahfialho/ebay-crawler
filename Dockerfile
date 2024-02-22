# Usar Alpine Linux como imagem base por ser leve e já conter os certificados necessários
FROM golang:1.18 as builder

# Definir o diretório de trabalho para garantir que todos os comandos sejam executados a partir deste local
WORKDIR /app

# Copiar os arquivos go.mod e go.sum e instalar as dependências
COPY go.* ./
RUN go mod download

# Copiar o restante dos arquivos do projeto
COPY . .

# Compilar o programa. Ajuste feito para o arquivo main.go dentro do diretório cmd
RUN CGO_ENABLED=0 GOOS=linux go build -v -o myapp cmd/main.go

FROM alpine:latest  

# Instalar ca-certificates
RUN apk --no-cache add ca-certificates

# Agora podemos criar a pasta data com as permissões corretas
RUN mkdir data && chmod 777 data

# Copiar o binário compilado do builder para a imagem final
COPY --from=builder /app/myapp myapp

ENTRYPOINT ["/myapp"]
