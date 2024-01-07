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
| `password`      | form | string  | Senha do usuário. |
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

Altere as informações do usuário.

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `email`        | body | string  | Email do usuário |
| `rua`        | body | string  | Logradouro do usuário |
| `numero`        | body | string  | Número do logradouro |
| `complemento`        | body | string  | Complemento do logradouro |
| `cidade`        | body | string  | Cidade do logradouro |
| `estado`        | body | string  | Estado do logradouro |
| `CEP`        | body | string  | CEP do logradouro |

**Resposta**

```json
{
    "message": "User updated success",
    "requestID": "d4920f0f-6433-4726-a014-21cdb4aed024"
}
```

---

### DELETE /users/\<cpf>

Deleta um usuário, consequentemente seu endereço e seu qrcode.

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

Se autentique através do login.

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `cpf`        | form | string  | CPF do usuário. |
| `password`    | form | string  | Senha do usuário. |

**Resposta**

```json
{
    "message": "User deleted w/ success",
    "requestID": "d4920f0f-6433-4726-a014-21cdb4aed024",
    "tokenJwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT"
}
```

---

### POST /password/recovery

Recuperação de senha.

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `email`        | form | string  |  do usuário. |

**Resposta**

```json
{
    "message":   "Token generated successfully",
    "redis-log": "key and value received",
    "email-log": "email sended success",
    "requestid": "d4920f0f-6433-4726-a014-21cdb4aed024"
}
```

---

### POST /password/verify

Verificando identidade da recuperação de senha.

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `email`        | body | string  |  do usuário. |
| `token`    | form | string  | Senha do usuário. |

**Resposta**

```json
{
    "message":   "redis authenticated token",
    "requestid": "d4920f0f-6433-4726-a014-21cdb4aed024"
}
```

---

### POST /password/change

Alterando senha pós comprovação e recuperação de senha.

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `email`        | body | string  |  Email do usuário. |
| `hashpassword`    | form | string  | Hash da Senha do usuário. |

**Resposta**

```json
{
    "message":   "password updated w/ sucess",
    "requestid": "d4920f0f-6433-4726-a014-21cdb4aed024"
}
```

---

