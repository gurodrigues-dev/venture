```bash

Summary:
    # -> Sprint
    # - [] -> Storie
    # - Task 

```

# S1 - Project and future

- [X] Escrever stories and tasks 
- [X] Refactor de ponteirs and issues
    - [X] Refactor of pointers
    - [X] Issues
        - [X] Email bug
        - [X] Qrcode bug

- [ ] Entregar arquitetura do backend 
    - [ ] Estudar uma api de mapas e que forme rotas.
        - Candidatos (em ordem de preferência):
            1. Google Maps

    - [ ] Estudar formas de pagamento.
        - Canditatos (em ordem de preferência):
            1. Pagar.me
            2. Pagseguro
            3. Stripe

- [ ] Entregar arquitetura do projeto
    - O volume deve ser alto, portanto devemos usar filas.

# S2 - Criando separação entre as categorias de Users e Drivers

- [X] Adaptar users para drivers 
    - Criar coluna no database chamada categoria e para motoristas adicionar a categoria "drivers" e para comum "user". 
    - Um user, pode se tornar driver
    - Um driver, deve poder acessar tudo referente a user.
    - Quando um user for registrado, não deve ter qrcode
    - Quando um user virar driver, ele deve receber um qrcode.

# S3 - Logs, Testes e Filas

- [X] Implantação de logs

- [ ] Adicionar Kafka
    - Somente a criação de usuários.

- [X] Adicionar testes


# S4 - Drivers

- [ ] Verificação de documentos.

- [ ] Deve ser impossível se registrar não residindo no estado de SP
 




    