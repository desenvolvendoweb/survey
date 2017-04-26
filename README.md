Este teste é voltado para aprimoramento e estudos envolvendo várias tecnologias, cujo:
- **Docker ( Creation and management of isolated environments -> docker-compose );**
- **Nginx ( server proxy HTTP and reverse );**
- **Gunicorn ( server wsgi native );**
- **Python ( High-level programming language );**
- **Django ( Web framework written in python );**
- **Golang ( Programming language big scale );**
- **PostgreSQL ( Object-relational database management system );**

A ideia do teste ( aplicado por uma empresa real ) é realizar as seguintes tarefas:
- **Escrever uma aplicação em Django onde o usuário ( logado ) possa criar enquetes;**
- **Cada enquete deve ter uma página com URL ( amigável );**
- **Cada usuário pode votar apenas uma vez ( lembrando que o suário está rodando em uma rede interna e só tem um usuário usando, ou seja, deve ser limitado o voto por sessão ). Não é necessário estar logado;**
- **Na página da enquete, deve conter as alternativas a serem votadas e após votar, deve aparecer um gráfico com a porcentagem de voto para cada opção;**
- **Abrir uma API ( respondendo no /rest ) escrita em Golang onde eu consiga:**
  - Votar ( lembrando que precisa validar a sessão );
  - Ver o resultado dos votos.
- **Pensar em um layout criativo no front-end para apresentar a enquete;**
- **Banco de dados a ser usado: PostgreSQL;**
- **Servidor HTTP a ser usado: Nginx;**
- **Servidor WSGI a ser usado: Gunicorn;**
- **Subir o ambiente com docker ( usando apenas docker-compose up ).**

Vamos para prática:

Para baixar o teste no link:

    git clone https://github.com/desenvolvendoweb/survey.git

Entre no diretório do projeto:

    cd survey

Para subir o ambiente isolado ( [docker compose install documentation](https://docs.docker.com/compose/install/) ):

    docker-compose up

Pronto ! O comando irá fazer todo os resto. Irá baixar as imagens, copiar o código para os devidos ambientes, instalar as apps necessárias e subir os serviços.

Para acessar o ambiente web para adicionar as enquetes, basta acessar link:

[http://localhost/admin](http://localhost/admin/) login: admin senha: admin

Cadastre suas enquetes, com as alternativas correspondentes a cada enquete e para acessar no front-end, acesse:

[http://localhost](http://localhost)

Navegue pelo menu lateral esquerdo, e após votar, note que o gráfico de pizza assume o papel no local da enquete ( porque somente poderá votar uma vez por enquete ). Se quiser votar novamente, para testar, é só abrir em uma aba anônima :).

Outro recurso bem interessante aplicado neste teste é API Restfull escrita em Golang.

Para acessar via navegador as boas vindas, acesse:

[http://localhost/rest](http://localhost/rest/)

Agora, para acessar as demais url's, é nessário instalar um interpretador de json para visualizar no próprio navegador. Se estiver usando o Google-Chrome, segue o link da extensão para ser instalada [JSONView](https://chrome.google.com/webstore/detail/jsonview/chklaanhfefbnpoihckbnefhakgolnmc?utm_source=chrome-app-launcher-info-dialog).

Para visualizar usando navegador, acesse:

[http://localhost/rest/surveys](http://localhost/rest/surveys)

Irá mostrar todas as enquetes em lista e suas alternativas.

Para visualizar uma enquete especifica, acesse:

[http://localhost/rest/survey/(ID)](http://localhost/rest/survey/1)

Irá mostrar a enquete especifica e suas alternativas.

Para votar ( neste caso não será possível por navegador ) basta usar um script que está no seguinte diretório:

    /rest/go/src/api/vote.sh

Entre no diretório, dê permissão de execução ( chmod +x vote.sh ) e rode ele usando shell ( ./vote.sh ). O script irá votar usando método POST e retornar objetos em JSON informando se foi sucesso ou falha. lembrando que você só pode votar uma vez, caso contrário, será informado da falha !

- **Algumas dicas:**
  - Para rodar o programa docker-compose up, será preciso ter acesso root;
  - A porta 80 será usada pelo programa Nginx, então antes de subir, favor liberá-la;
  - Para maiores informações de como está configurado o Nginx, [clique Aqui](./tree/master/nginx);
  - Para maiores informações de como está configurado o PostgreSQL, [clique aqui](./tree/master/postgres);
  - Para maiores informações de como está configurado e como foi programado o Restfull, [clique aqui](./tree/master/rest);
  - Para maiores informações de como está configurado e como foi programado o serviço Web, [clique aqui](./tree/master/web).

Fim.
