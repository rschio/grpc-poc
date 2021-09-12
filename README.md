# Exemplo de API gRPC

## Build
```sh
go mod tidy
```
### Nos diretórios `cmd/client`, `cmd/server`, `cmd/gateway`:
```sh
go build
```

## Teste
```sh
go test ./...
```
```sh
go test -tags=integration ./...
```
## Executando
### gRPC
Execute em terminais distintos:
```sh
./cmd/server/server
```
```sh
./cmd/client/client
```
### gRPC-Gateway
Importe a documentação da API para o Postman:
`tracker/openapiv2/tracker.swagger.json`

Edite a váriavel de ambiente __baseURL__ da coleção do Postman para: __http://localhost:8081__.

Alguns endpoints necessitam de autenticação, para isso use o Bearer Token: __consegue__.

Execute em terminais distintos:
```sh
./cmd/server/server
```
```sh
./cmd/gateway/gateway
```
Tente pesquisar pelo trackingCode: BR000000000BR
