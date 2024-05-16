
<h1 align="center"> 🌬️ Venture </h1>

<h1 align="center"> Somos segurança, velocidade e tecnologia. Somos Venture </h1>

<p align="center">
  <img src="https://i.imgur.com/yieDOSJ.png"/>
</p>

# 👋 Olá, esta é a documentação oficial da primeira versão do Venture Backend!

### ❓ Quem somos?

Somos os criadores do aplicativo Venture, disponível para Android e IOS. O aplicativo que promete facilitar algo que você continua tendo dor de cabeça. Prometemos fazer o melhor para todos os lados, proporcionando mais avanço na educação.

- Como ajudamos os motoristas?

Nosso foco foi em realmente, facilitar a vida dos motoristas, em primeira instância. Mas, depois analisamos e verificamos minunciosamente, a possibilidade de ajudar a todos.

O motorista tem sérios problemas para receber em todos os meses do ano e com data fixa. Eliminamos este problema, solicitando pagamentos automáticos via Boletos, Cartão de Débito e Crédito, com plataformas externas

Além de ter que todo dia verificar quais são as crianças que vão ou não vão no dia seguinte. E montar uma rota completa de acordo com as alterações.

Este problema também foi sanado, o pai poderá adicionar um dia antes, se seu filho irá ou não para a escola no dia seguinte, junto de uma justificação (opcional).

Montar a rota, também não é mais um problema do motorista. Junto com uma API externa fazemos todo esse papel pra ele, já verificando, quem vai ou não para a escola.

Os motoristas, precisam aceitar um convite de uma escola para fazer parte dela, criando uma parceria. Para que quando o pai adicionar o filho na escola, verificar todos os motoristas disponíveis naquela escola. Mas, somente a escola pode enviar o convite e gerenciar se o motorista ainda faz parte ou não da escola.

- Como ajudamos os pais e filhos?

Os pais poderão registrar todos os seus filhos no aplicativo. E registra-los em escolas diferentes, com tios diferentes. Ao colocar a escola, poderá verificar todos os motoristas que tem parceria com aquela escola. 

O responsável poderá verificar o perfil do motorista, as avaliações e as vantagens, além de claro. O preço da mensalidade.

No futuro pretendemos adicionar uma verificação em tempo real, a rota do motorista, enquanto estiver atuando com seu filho.

- Como ajudamos as escolas?

Com o aplicativo, facilitamos a gerência de motoristas. Para que as escolas possam convida-los e que eles sejam "contratados", como um negócio, de fato.

Também será mais simples e seguro, compreender qual é o motorista responsável pela criança no trajeto.

- Atualmente, ainda estamos em construção, então fique preocupado, caso o serviço não esteja completo.

# 🛠️ Dependências e linguagens

### 🔵 Linguagem 

Por ser a stack com menor curva de aprendizado e mais ferramentas web, leitura simplificada, rápida compilação. Escolhemos Go, a linguagem do google.

---

> Antes de abordarmos sobre as dependências e tentar rodar este serviço de maneira isolada. Recomendamos que vá até o repositório oficial do Venture, rodando a infraestrutura completa. Além disso você também pode rodar diretamente o Dockerfile, ou usar os comandos Go para instalação automática da libs. Entretanto, caso queira realmente verificar as nossas dependências, aqui estão:

### 🫁 Dependências

- AWS SDK Go
- JWT Go
- Gin Gonic
- Go Redis
- Postgresql
- Gin Swagger
- Files Swagger
- Yaml V2
- Kakfa Go
- QRCode GO
- Base64x
- TOML Go

# 🚀 Inicializando

Para inicializar o Venture, via Dockerfile, recomendamos que inicie a infraestrutura completa do Venture, como já citado anteriormente. Portanto.

```sh
docker compose up --build -d --scale venture-bff=3
```

Agora para rodar o serviço manualmente, lembramos que é extremamente necessário possuir um arquivo chamado `yaml.config` dentro do diretório `/config`. Com a seguinte estrutura e tendo todos os recursos disponíveis.

```yaml
name: venture
development: true

server:
  host: "host-database"
  port: 8787
  secret: "secret"

database:
  dbuser: "user-of-db-postgres"
  dbport: 5432
  dbhost: "host-or-ip-database"
  dbpassword: "password-of-user-in-database"
  dbname: "venture_staging"
  schema: "database/schema.sql"

cloud:
  region: "region-cloud"
  accesskey: "access-key-aws"
  secretkey: "secret-key-aws"
  token: "optional" # yes, it are optional
  source: "email-source"
  bucket: "bucket-name"

cache:
  address: "address-of-redis:port-redis"
  password: "password-redis"

messaging:
  broker: "broker-ip:port-of-kafka"
  topic: "email.send.topic"
  partition: 0

```

- ⚠️ Necessário ter todos esses recursos disponíveis, nem que sejam locais ou utilizando LocalStack. Os valores, que estão definidos de fato, devem permanecer: Nome do Banco de Dados; Porta da aplicação (a alteração é opcional mas não se esqueça de alterar também no Dockerfile); Schema e Diretório; Nome do Tópico Kafka e sua partição.

## ⚙️ API Endpoints

Todas as rotas possuem `api/v1`. Antecendo, como prefixo da rota.

### GET /ping

Retorna uma simples mensagem de "pong" para validar funcionamento da aplicação.

**Resposta**

```json
{
    "message": "pong"
}
```
---

### Middleware - Driver

Verifica se você tem CNH válido na plataforma e que foi gerado pela rota de login.

**Resposta 200**

```json
{
    "is.Authenticated": true,
    "cnh": "xxx"
}
```

**Resposta 401**

```json
{
    "message": "Erro ao verificar token de sessão"
}
```

---

### Middleware - School

Verifica se você tem CNPJ válido na plataforma e que foi gerado pela rota de login.

**Resposta 200**

```json
{
    "is.Authenticated": true,
    "cnpj": "xxx"
}
```

**Resposta 401**

```json
{
    "message": "Erro ao verificar token de sessão"
}
```

---

### POST /driver

Crie conta de motorista na plataforma.

**Parâmetros**

| Nome       | Local | Tipo   | Descrição            |
|------------|-------|--------|----------------------|
| `name`     | body  | string | Nome do motorista.   |
| `email`    | body  | string | E-mail do motorista.|
| `password` | body  | string | Senha do motorista. |
| `cpf`      | body  | string | CPF do motorista.    |
| `cnh`      | body  | string | CNH do motorista.    |
| `street`   | body  | string | Rua do motorista.    |
| `number`   | body  | string | Número do motorista. |
| `zip`      | body  | string | CEP do motorista.    |
| `complement`| body | string | Complemento do endereço do motorista. |
     

**Resposta**

```json
{
    "driver": {
      "id": 123,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "password": "s3cr3tP@ssw0rd",
      "cpf": "123.456.789-00",
      "cnh": "123456789",
      "street": "123 Main Street",
      "number": "456",
      "zip": "12345-678",
      "complement": "Apt 101"
    } 
}
```

---

### GET /driver/:cnh

Verifique uma conta de um motorista.

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `cnh` | uri | string  | CNH do motorista. |     

**Resposta**

```json
{
    "driver": {
      "id": 123,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "cpf": "123.456.789-00",
      "cnh": "123456789",
      "street": "123 Main Street",
      "number": "456",
      "zip": "12345-678",
      "complement": "Apt 101"
    } 
}
```

---

### PATCH /driver

Altere os próprios dados, de acordo com o campo que alterar (para motorista). Retorno do JSON com dados atualizados.

**Parâmetros**

| Nome       | Local | Tipo   | Descrição            |
|------------|-------|--------|----------------------|
| `name`     | body  | string | Nome do motorista.   |
| `email`    | body  | string | E-mail do motorista.|
| `password` | body  | string | Senha do motorista. |
| `cpf`      | body  | string | CPF do motorista.    |
| `cnh`      | body  | string | CNH do motorista.    |
| `street`   | body  | string | Rua do motorista.    |
| `number`   | body  | string | Número do motorista. |
| `zip`      | body  | string | CEP do motorista.    |
| `complement`| body | string | Complemento do endereço do motorista. |    

**Resposta**

```json
{
    "driver": {
      "id": 123,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "cpf": "123.456.789-00",
      "cnh": "123456789",
      "street": "123 Main Street",
      "number": "456",
      "zip": "12345-678",
      "complement": "Apt 101"
    } 
}
```

---

### DELETE /driver

Motorista deletando sua própria conta.     

**Resposta**

```json
{
    "message": "Conta deletada com sucesso"
}
```

---

### POST /login/driver

Cria uma conta na API

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `cnh`| body | string  | CNH do motorista. |
| `password`| body | string  | Senha do usuário. |      

**Resposta**

```json
{
    "driver": {
      "id": 123,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "cpf": "123.456.789-00",
      "cnh": "123456789",
      "street": "123 Main Street",
      "number": "456",
      "zip": "12345-678",
      "complement": "Apt 101"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}
```

---

### GET /driver/partners

Verificar todas as escolas parceiras.    

**Resposta**

```json
{
  "school": [{
    "id": 456,
    "name": "Escola Estadual",
    "cnpj": "12.345.678/0001-90",
    "email": "escola@exemplo.com",
    "street": "Rua da Escola, 123",
    "number": "456",
    "zip": "12345-678"
  },
  {
    "id": 456,
    "name": "Escola Estadual",
    "cnpj": "12.345.678/0001-90",
    "email": "escola@exemplo.com",
    "street": "Rua da Escola, 123",
    "number": "456",
    "zip": "12345-678"
  },
  {
    "id": 456,
    "name": "Escola Estadual",
    "cnpj": "12.345.678/0001-90",
    "email": "escola@exemplo.com",
    "street": "Rua da Escola, 123",
    "number": "456",
    "zip": "12345-678"
  }]
}
```

---

### POST /school

Cria uma conta de escola na plataforma

**Parâmetros**

| Nome     | Local | Tipo   | Descrição            |
|----------|-------|--------|----------------------|
| `id`     | body  | int    | ID da escola.        |
| `name`   | body  | string | Nome da escola.      |
| `cnpj`   | body  | string | CNPJ da escola.      |
| `email`  | body  | string | E-mail da escola.    |
| `password` | body | string | Senha da escola.    |
| `street` | body  | string | Rua da escola.       |
| `number` | body  | string | Número da escola.    |
| `zip`    | body  | string | CEP da escola.       |    

**Resposta**

```json
{
  "school": {
    "id": 456,
    "name": "Escola Estadual",
    "cnpj": "12.345.678/0001-90",
    "email": "escola@exemplo.com",
    "password": "segredo123",
    "street": "Rua da Escola, 123",
    "number": "456",
    "zip": "12345-678"
  }
}
```

---

### POST /school/:cnpj

Cria uma conta na API

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `nome` | body | string  | Nome do usuário. |
| `email`| body | string  | E-mail do usuário. |
| `senha`| body | string  | Senha do usuário. |      

**Resposta**

```json
{
    "message": "Usuário criado com sucesso"
}
```

---

### POST /school

Cria uma conta na API

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `nome` | body | string  | Nome do usuário. |
| `email`| body | string  | E-mail do usuário. |
| `senha`| body | string  | Senha do usuário. |      

**Resposta**

```json
{
    "message": "Usuário criado com sucesso"
}
```

---

### POST /school

Cria uma conta na API

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `nome` | body | string  | Nome do usuário. |
| `email`| body | string  | E-mail do usuário. |
| `senha`| body | string  | Senha do usuário. |      

**Resposta**

```json
{
    "message": "Usuário criado com sucesso"
}
```

---

### POST /school

Cria uma conta na API

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `nome` | body | string  | Nome do usuário. |
| `email`| body | string  | E-mail do usuário. |
| `senha`| body | string  | Senha do usuário. |      

**Resposta**

```json
{
    "message": "Usuário criado com sucesso"
}
```

---

### POST /login/school

Cria uma conta na API

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `nome` | body | string  | Nome do usuário. |
| `email`| body | string  | E-mail do usuário. |
| `senha`| body | string  | Senha do usuário. |      

**Resposta**

```json
{
    "message": "Usuário criado com sucesso"
}
```

---

### POST /school/employees

Cria uma conta na API

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `nome` | body | string  | Nome do usuário. |
| `email`| body | string  | E-mail do usuário. |
| `senha`| body | string  | Senha do usuário. |      

**Resposta**

```json
{
    "message": "Usuário criado com sucesso"
}
```

---


### POST /school/employees

Cria uma conta na API

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `nome` | body | string  | Nome do usuário. |
| `email`| body | string  | E-mail do usuário. |
| `senha`| body | string  | Senha do usuário. |      

**Resposta**

```json
{
    "message": "Usuário criado com sucesso"
}
```

---

### POST /invite

Cria uma conta na API

**Parâmetros**

| Nome       | Local | Tipo   | Descrição            |
|------------|-------|--------|----------------------|
| `requester`| body  | string | Solicitante do convite.|
| `guest`    | body  | string | Convidado pelo convite.|     

**Resposta**

```json
{
  "requester": "João",
  "guest": "Maria"
}
```

---

### POST /invite

Cria uma conta na API

**Parâmetros**

| Nome       | Local | Tipo   | Descrição            |
|------------|-------|--------|----------------------|
| `requester`| body  | string | Solicitante do convite.|
| `guest`    | body  | string | Convidado pelo convite.|     

**Resposta**

```json
{
  "requester": "João",
  "guest": "Maria"
}
```

---

### GET /invite

Verificar todos os invites de um motorista    

**Resposta**

```json
{
    "id": "1",
    "school": "E.E Estadual",
    "driver": "Maria João José",
    "status": "pending"
}
```

---

### PATCH /invite/:id

Aceite um invite para ser parceiro de uma escola

**Parâmetros**

| Nome       | Local | Tipo   | Descrição            |
|------------|-------|--------|----------------------|
| `id`       | url  | int    | ID do convite.       | 

**Resposta**

```json
{
    "id": "1",
    "school": "E.E Estadual",
    "driver": "Maria João José",
    "status": "pending"
}
```

---

### DELETE /invite/:id

Recuse o convite de uma escola

**Parâmetros**

| Nome       | Local | Tipo   | Descrição            |
|------------|-------|--------|----------------------|
| `id`       | url  | int    | ID do convite.       |    

**Resposta**

```json
{
    "message": "invite was deleted w/ sucessfully"
}
```

---

# Arquitetura

![8137cb5a-3b22-4e0c-b6a0-c28f890e132e](https://github.com/gurodrigues-dev/venture/assets/83222330/fa2c214a-88dd-41a1-ba43-1b41bb564e61)
