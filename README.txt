- Escrever uma aplicação em Django onde eu (como usuários logado) possa criar enquetes
- Cada enquete deve ter uma pagina com URL amigável
- Cada usuario pode votar apenas uma vez (lembrando que estou rodando em rede interna e só tem eu como usuários, ou seja, deve ser limitado o voto por session), não precisa estar logado para votar
- Na pagina da enquete deve conter a enquete para ser respondida, após votar deve aparecer um gráfico com a porcentagem de voto em cada opção
- Abrir uma API (respondendo no /rest) escrita em Go onde eu consiga
— votar (lembrando que precisa validar a sessão na hora de votar)
— ver resultado dos votos

Como a vaga é para full stack, vamos pensar no layout (quero ver sua criatividade)

banco de dados: PostgreSQL
servidor http: nginx
servidor WSGI: gunicorn

Quero subir o ambiente com docker (usando docker-compose up)
