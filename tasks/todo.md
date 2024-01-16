```bash

Summary:
    # -> Sprint
    # - [] -> Storie
    # - Task 

```

# S1 - Project and future

- [ ] Escrever stories and tasks 
- [ ] Refactor de ponteirs and issues
    - [X] Refactor of pointers
    - [ ] Issues
        - [ ] Email bug
        - [ ] Qrcode bug
- [ ] Entregar arquitetura do backend 
    - Estudar uma api de mapas e que forme rotas.
    - Formas de pagamento.
- [ ] Entregar arquitetura do projeto
    - O volume deve ser alto, portanto devemos usar filas.

# S2 - Criando separação entre as categorias de Users e Drivers

- [ ] Adaptar users para drivers - #SGR3
    - Criar coluna no database chamada categoria e para motoristas adicionar a categoria "drivers" e para comum "user". 
    - Um user, pode se tornar driver
    - Um driver, deve poder acessar tudo referente a user.
    - Quando um user for registrado, não deve ter qrcode
    - Quando um user virar driver, ele deve receber um qrcode.

# S3 - Drivers

- [ ] Verificação de documentos.
- [ ] Ele deve setar uma região de trabalho (independente do endereço) a qual irá trabalhar, e as escolas e endereços elegiveis, só entrarão em um raio de 10 km.
- [ ] Ele deve ter uma lista de alunos, registrados em seu carro, com período atrelado.
- [ ] Conforme o período, devemos montar a melhor rota para ele baseada no caminho de sua casa em conjunto.
- [ ] Para cada motorista, deve existir uma aba de avaliação. Entretanto, apenas podem avaliar quem já foi cliente do motorista em algum momento.
- [ ] Deve ser impossível se registrar não residindo no estado de SP

# S4 - Users

- [ ] Ele deve registrar seu filhos e a escola em que estudam.
    - Podendo ter mais de um filho
    - Podendo mudar a escola de seu filho
    - Podendo alterar de motorista
- [ ] Ele deve poder pesquisar a escola e os motoristas da escola
- [ ] Ele pode assinar com o motorista por um preço FIXO e parcelar na quantidade que quiser.
    - Com desistência impondo multa.
    - Validar quando plano esta se encerrando.
- [ ] Ele deve poder verificar os documentos do motorista.
- [ ] Deve ser impossível se registrar não residindo no estado de SP




    