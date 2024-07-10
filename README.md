Criar uma aplicação com os seguintes requisitos:

1. Usar golang para o backend
2. Usar um framework de ORM para acesso ao banco de dados - recomendo o https://gorm.io/index.html 
3. Criar o frontend em qualquer framework - HTML + CSS + Bootstrap + JS, React, Angular, etc
4. O projeto deve ser feito em dupla
5. No final ele deve ser hospedado em alguma alternativa de hospedagem gratuita - recomendo o https://www.oracle.com/br/cloud/
O dono do sistema pode cadastrar em um sistema web: professores, turmas, alunos, atividades e notas. Os dados do professor são: Nome, e-mail, CPF. Os dados da turma são: nome da turma, semestre, ano, qual professor a leciona. Os dados do aluno são: nome, matrícula, em quais turmas ele está cadastrado. Os dados da atividade são: turma, valor e data. Deve-se restringir para que todas as atividades somadas não ultrapassem 100 pontos. A nota deve ser realizada da seguinte maneira: se escolhe uma turma e uma atividade. Após isso, todos os alunos são apresentados na tela e o professor pode digitar sua nota. A nota deve ser validada de acordo com o máximo permitido da disciplina. 

backend    
    models
        professor.go
        turma.go
        aluno.go
        atividade.go
    routes
        professor.go
        turma.go
    config
        db.go
        config.yaml
    main.go
