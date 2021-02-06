# banking-auth

##### Run `./start.sh` to download the dependencies and run the the application

To run the application, you have to define the environment variables, default values of the variables are defined inside `start.sh`

- SERVER_ADDRESS    `[IP Address of the machine]`
- SERVER_PORT       `[Port of the machine]`
- DB_USER           `[Database username]`
- DB_PASSWD         `[Database password]`
- DB_ADDR           `[IP address of the database]`
- DB_PORT           `[Port of the database]`
- DB_NAME           `[Name of the database]`




1º acesso a camada de serviço pra realização do login
2º acesso ao repositorio onde executa o sql pra retornar os dados do usuário caso ele exista
3º acessa o metodo gerar token presente na classe de login
   1. verifica se é retornado o campo account e customer diferente de nulo
      1.1 - se diferente de nulo -> é criado um token pra usuário
      1.2 - se ambos nulo -> é criado token de administrador 
4º É invocado o método jwt.NewWithClaims pra criação das claims com dados do usuário
5º É retornado a numeração do token no formato String por base nas clams geradas
6º 