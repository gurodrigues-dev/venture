---
title: Gin
description: A Gin server
tags:
  - gin
  - golang
---

## 🚀 Inicializando

As dependências são encontradas no `go.mod`, basta baixar o repo para possui-los. 

Com todas corretamente instaladas, inicie a aplicação.

```sh
go run main.go
```

## ⚙️ API Endpoints

Você pode definir nas rotas uma porta específica, claro, se souber. Mas se não souber fique tranquilo. Ela (API) inicia na porta 8080.

Caso não funcione certifique-se de que não há nada rodando na porta 8080.

### GET /health

Retorna o status de saúde da API e seu uso de recursos.

**Resposta**

```json
{
    "cpu": "3.2",
    "envs": "load environments ok!",
    "mem": "59.9",
    "message": "pong",
    "uptime": "7h52m34s"
}
```
---

### POST /users

Criando um novo usuário.

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `nome`       | form | string  | Nome do usuário. |
| `email`      | form | string  | E-mail do usuário. |
| `senha`      | form | string  | Senha do usuário. |
| `cpf`        | form | string  | CPF do usuário. |
| `rg`         | form | string  | RG do usuário. | 
| `cnh`        | form | string  | CNH do usuário. |  
| `rua`        | form | string  | Logradouro do usuário. | 
| `numero`     | form | string  | Numero referente Endereço. | 
| `complemento`| form | string  | Complemento referente ao Endereço. | 
| `cep`        | form | string  | CEP do Endreço. | 
| `cidade`     | form | string  | Cidade do Endereço. | 
| `estado`     | form | string  | Estado do Endereço. | 

**Resposta**

```json
{
    "requestID": "1f9167c5-eb52-440d-80a4-eb28fc496295",
    "s3bucketurl": "https://<bucket-name>.s3.amazonaws.com/qrcodes/<cpf>.png",
    "status": "user created successfully"
}
```

---

### GET /users/\<cpf>

Busca um usuário e todos seus dados.

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `nome`       | body | string  | Nome do usuário. |
| `email`      | body | string  | E-mail do usuário. |
| `qrcode`     | body | string  | Link do QRCode. |
| `cpf`        | body | string  | CPF do usuário. |
| `rg`         | body | string  | RG do usuário. | 
| `cnh`        | body | string  | CNH do usuário. |  
| `rua`        | body | string  | Logradouro do usuário. | 
| `numero`     | body | string  | Numero referente Endereço. | 
| `complemento`| body | string  | Complemento referente ao Endereço. | 
| `cep`        | body | string  | CEP do Endreço. | 
| `cidade`     | body | string  | Cidade do Endereço. | 
| `estado`     | body | string  | Estado do Endereço. | 

**Resposta**

```json
{
    "requestID": "f093e965-8cb3-4889-9863-152909b019ae",
    "userData": {
      "CPF": "93404833082",
      "RG": "552386347",
      "Name": "Gustavo Rodrigues",
      "CNH": "28053612377",
      "Email": "gustavorodrigueslima2004@gmail.com",
      "URL": "https://<bucket-name>.s3.amazonaws.com/qrcodes/93404833082.png",
      "Endereco": {
        "Rua": "rua cubatao",
        "Numero": "77",
        "Complemento": "apto 5",
        "Cidade": "sao paulo",
        "Estado": "sp",
        "CEP": "08132450"
		}
	}
}
```
---

### PUT /users/\<cpf>

Em desenvolvimento...

---

### DELETE /users/\<cpf>

Deleta um usuário, consequentemente seu usuário e seu qrcode.

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `cpf`        | body | string  | CPF do usuário. |

**Resposta**

```json
{
    "message": "User deleted w/ success",
    "requestID": "d4920f0f-6433-4726-a014-21cdb4aed024"
}
```

---

### POST /users/login

Em desenvolvimento...
