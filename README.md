# Sistema de Gestão Acadêmica

Este projeto foi desenvolvido como parte do curso de Análise e Desenvolvimento de Sistemas. O objetivo é fornecer uma aplicação web que permite o cadastro de professores, turmas, alunos, atividades e notas. A aplicação valida as notas das atividades para garantir que o total de pontos não exceda 100.

## Tecnologias Utilizadas

- **Backend**: Golang
- **ORM**: GORM (https://gorm.io/)
- **Frontend**: (HTML + CSS + Bootstrap + JS, React, Angular, etc)
- **Hospedagem**: Oracle Cloud (https://www.oracle.com/br/cloud/)

## Funcionalidades

1. **Professores**: Cadastro de nome, e-mail e CPF.
2. **Turmas**: Cadastro de nome da turma, semestre, ano e professor responsável.
3. **Alunos**: Cadastro de nome, matrícula e turmas.
4. **Atividades**: Cadastro de turma, valor e data.
   - Restrição: O valor total das atividades de uma turma não pode ultrapassar 100 pontos.
5. **Notas**: Atribuição de notas para os alunos em uma atividade específica, com validação do valor máximo permitido.

## Estrutura do Projeto

### Backend

- **Golang**: Utilizado para o desenvolvimento do backend.
- **GORM**: Framework de ORM para acesso ao banco de dados.
- **Gin**: Framework web para criar APIs RESTful.

### Frontend

- Framework utilizado: Em desenvolvimento
### Hospedagem

- **Oracle Cloud**: Serviço de hospedagem gratuito utilizado para disponibilizar a aplicação.

## Como Executar

### Pré-requisitos

- Golang instalado: https://golang.org/dl/
- Node.js instalado (se o frontend utilizar Node.js): https://nodejs.org/

### Passos

1. Clone o repositório:
   ```sh
   git clone https://github.com/Gabrielcamargos28/nome-do-repositorio.git
   cd nome-do-repositorio
