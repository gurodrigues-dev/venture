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

- [X] Entregar arquitetura do backend 
    - [X] Estudar uma api de mapas e que forme rotas.
        - Candidatos (em ordem de preferência):
            1. Google Maps

    - [ ] Estudar formas de pagamento.
        - Canditatos (em ordem de preferência):
            1. Pagar.me
            2. Pagseguro
            3. Stripe

- [X] Entregar arquitetura do projeto
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

- [X] Adicionar Kafka
    - Somente a criação de usuários.

- [X] Adicionar testes


# S4

- [ ] Fila para criação de Usuários

- [ ] Tabela de child com user como FK 
    - Nome
    - RG (PK)
    - Nome do Responsável (User)
    - Endereço do Responsável
    - Horario do Aluno

    Dados que não serão registrados de imediato.
    - Nome da Escola 
    - Nome do Driver

- [ ] CRUD de child (sem fila)

- [ ] Deve ser impossível se registrar não residindo no estado de SP

# S5 

- [ ] CRUD de escolas (verificando cnpj)
    - CPNJ 
    - Endereço

- [ ] Tabela intermediaria de Registro de Driver na escolas
    - Numero de matricula | ID da Escola | Driver
    - Gerenciado por conta de Escola
        - Adicionando ou Removendo
        - Adicionar pelo CPF 
        - Aceito pelo Driver

# S6

> Tudo que seja relacionado a pagamento vai ser mockado, até implantarmos de fato, um pagamento.

- [ ] Tabela intermediaria de Registro de alunos no Driver junto do horário
    - ID do registro | Driver | Aluno | Horario

    (Sem gerenciamento, feito automático)

    * Mudou de escola: Se o driver não trabalhar na mesma escola, necessário pagamento de multa de 20% do valor restante da parcela.

    * Mudou de escola: Driver trabalha na mesma escola, nada se altera, preços devem ser fixos.

    * Quero mudar de transporte: Multa de 20% do restante das parcelas, novo registro do tio. 

# S7

- [ ] Criar funcionalidade de montar rota para o Driver, de acordo com os forms
    - Horário
    - Ida ou Volta?

    - [ ] Essa funcionalidade vai permitir que todo dia o user, adicione se seu filho irá ou não para escola, assim influenciando ou removendo alguém da lista do respectivo dia.








    