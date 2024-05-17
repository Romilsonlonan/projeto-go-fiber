<p align="center">
  <img class="align-items: center" src="https://github.com/Romilsonlonan/projeto-go-fiber/assets/90980220/a2c8e078-efcf-4ba9-9cb2-4b9c70059c33" alt="Imagem Golang" style="width: 400px; height: 350px;">
</p>
<br>


# DOCUMENTAÇÃO

https://pkg.go.dev/github.com/gofiber/fiber/v2@v2.52.4#App.Listen


# CRIANDO UM AMBIENTE VIRTUAL

Bash
go mod venv myenv

Bash
source myenv/bin/activate 

# LIVE RELOAD PARA APLICAÇÕES GO 

Bash
go install github.com/cosmtrek/air@latest

Bash
go get -u github.com/cosmtrek/air


# CRIANDO UM PROJETO SIMPLES COM GO E FIBER

# Criação do Diretório do Projeto:

Bash
mkdir fiber-project

* Este comando cria um diretório chamado fiber-project para armazenar os arquivos do seu projeto.

# Navegação no Diretório:

Bash
cd fiber-project

* Este comando navega para o diretório fiber-project para que você possa executar os próximos comandos a partir dele.

# Criação do Arquivo Principal:

Bash
touch main.go

* Este comando cria um arquivo vazio chamado main.go que conterá o código principal da sua aplicação web.

# Inicialização do Módulo Go:

Bash
go mod init fiber-project

* Este comando inicializa um módulo Go para o seu projeto e cria um arquivo go.mod que gerencia as dependências do seu projeto.

# Instalação do Fiber:

Bash
go get github.com/gofiber

* Este comando instala o framework Fiber do repositório GitHub oficial e o torna disponível para uso no seu projeto.

# Conteúdo do Arquivo main.go:

package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    app.Listen(":3000")
}
 

* Este código define um pacote main, importa o pacote Fiber e define uma aplicação web básica. A função main cria uma nova instância do Fiber, define uma rota / que retorna a mensagem "Hello, World!" e inicia o servidor na porta 3000.

# Execução da Aplicação:

Bash
go run main.go

* Este comando compila e executa o código Go no arquivo main.go. Você deve ver a mensagem "Hello, World!" impressa no seu terminal quando acessar http://localhost:3000/ em seu navegador web.



# INSTALAÇÃO MYSQL 

Instalação e Configuração Segura do MySQL
Este tutorial explica como instalar o MySQL no Fedora e realizar a configuração inicial de segurança.

1. Instalação do MySQL Server:

Comando: dnf install mysql-community-server
Explicação: Este comando utiliza o gerenciador de pacotes dnf do Fedora para instalar o pacote mysql-community-server, que é a versão gratuita e de código aberto do servidor MySQL.

2. Verificar o Status do Serviço:

Comando: systemctl status mysql
Explicação: Este comando verifica o status do serviço mysql. Ele mostrará se o serviço está em execução, parado, habilitado para iniciar automaticamente ou desabilitado.

3. Iniciar o Serviço MySQL:

Comando: systemctl start mysqld
Explicação: Este comando inicia o serviço mysqld, que é o processo principal do servidor MySQL.

4. Habilitar Inicialização Automática do Serviço MySQL:

Comando: systemctl enable mysqld
Explicação: Este comando configura o serviço mysqld para iniciar automaticamente na inicialização do sistema. Isso garante que o servidor MySQL esteja sempre em execução após a reinicialização do computador, sem a necessidade de iniciar o serviço manualmente cada vez.

5. Parar o Serviço MySQL:

Comando: systemctl stop mysqld
Explicação: Este comando finaliza o processo mysqld, que é o serviço principal do servidor MySQL. Isso interrompe a execução do servidor MySQL e o torna indisponível para conexões.

6. Conectar ao MySQL (sem senha inicial):

Comando: sudo mysql
Explicação: Este comando tenta se conectar ao servidor MySQL usando o usuário root e sem senha (por padrão, o MySQL não possui senha definida durante a instalação).

7. Definir Senha para o Usuário root:

Comando: alter user 'root'@'localhost' identified by 'senha@123';
Explicação: Este comando altera a senha do usuário root do MySQL para 'senha@123' (substitua 'senha@123' por uma senha forte e complexa de sua escolha).

8. Atualizar Permissões:

Comando: flush privileges;
Explicação: Este comando atualiza as tabelas de permissões do MySQL para refletir a nova senha definida para o usuário root.

9. Sair do Shell do MySQL:

Comando: exit
Explicação: Este comando finaliza a conexão com o shell do MySQL.

10. Script de Segurança MySQL (opcional):

Comando: sudo mysql_secure_installation
Explicação: Este comando executa o script de segurança do MySQL, que permite definir configurações adicionais de segurança, como:
Remover contas de teste padrão do MySQL.
Remover acesso remoto para o usuário root.
Criar um novo usuário MySQL e conceder privilégios.
Atualizar o sistema de autenticação do MySQL.
Remover o banco de dados de teste padrão do MySQL.
Observações sobre o Script de Segurança:

As instruções na tela orientam por cada etapa do processo.
Digite a senha do usuário root criada anteriormente quando solicitado.
Responda às perguntas do script de segurança de acordo com suas necessidades (por exemplo, você pode optar por manter o acesso remoto para o usuário root se necessário).

11. Conectar ao MySQL com Senha:

Comando: mysql -u root -h localhost -p
Explicação: Este comando tenta se conectar ao servidor MySQL usando o usuário root, o hostname localhost (que indica a conexão local) e solicita a senha definida anteriormente.
Enter password: Digite a senha do usuário root criada no passo 5 e pressione Enter.
Depois de seguir esses passos, você terá o MySQL instalado e configurado com segurança no seu Fedora. Você poderá conectar ao servidor MySQL usando o usuário root e a senha definida e começar a gerenciar seus bancos de dados.

show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
4 rows in set (0,03 sec)


# INSTALAÇÃO DO GORM 

1. Baixar e Instalar o GORM:

go get -u gorm.io/gorm: Este comando utiliza a ferramenta go get do Go para baixar e instalar o pacote gorm do repositório gorm.io/gorm. O -u indica que a ferramenta deve atualizar o pacote para a versão mais recente.

2. Baixar e Instalar o Driver MySQL:

go get -u gorm.io/driver/mysql: Este comando utiliza a ferramenta go get para baixar e instalar o pacote driver/mysql do repositório gorm.io/driver/mysql. O -u indica que a ferramenta deve atualizar o pacote para a versão mais recente.


## Conexão com Mysql ok!

![Captura de tela de 2024-05-15 22-02-48](https://github.com/Romilsonlonan/projeto-go-fiber/assets/90980220/59dccefd-6af6-4261-a873-3bc9350e00aa)

<hr>

```
package controllers // Declara pacote `controllers`

import ( // Importa dependências
    "fiber-project/models" // Importa pacote `models` do diretório `fiber-project`
    "github.com/gofiber/fiber/v2" // Importa pacote `fiber` da versão 2
)

func Register(c *fiber.Ctx) error { // Define função `Register` que recebe um contexto do Fiber e retorna um erro
    // Cria uma nova instância de `User` do pacote `models`
    user := models.User{
        FirstName: "Toti", // Define o nome do usuário como "Toti"
    }

    // Define o sobrenome do usuário como "Cavalcante"
    user.LastName = "Cavalcante"

    // Retorna o JSON da estrutura `user` no contexto da requisição
    return c.JSON(user)
}

```

![Captura de tela de 2024-05-16 08-28-13](https://github.com/Romilsonlonan/projeto-go-fiber/assets/90980220/eb45d396-0237-4ed6-8a90-79c294dc8ff9)

#Register e Login 

package controllers // Declara pacote `controllers`

import ( // Importa dependências
	"fiber-project/database" // Importa
	"fiber-project/models"   // Importa pacote `models` do diretório `fiber-project`

	"github.com/gofiber/fiber/v2" // Importa pacote `fiber` da versão 2
	"golang.org/x/crypto/bcrypt"
)

// Função `Register` que recebe o contexto da requisição (`c`) e retorna um erro (`error`)
func Register(c *fiber.Ctx) error {
    // Cria um mapa de string para string para armazenar os dados do formulário
    var data map[string]string

    // Tenta ler os dados do corpo da requisição e verifica se há erro
    if err := c.BodyParser(&data); err != nil {
        // Retorna o erro caso a leitura falhe
        return err
    }

    // Define o status da resposta para "Requisição Ruim" (400 Bad Request)
    c.Status(400)

    // Verifica se as senhas informadas nos campos "password" e "confirm_password" são iguais
    if data["password"] != data["confirm_password"] {
        // Retorna uma mensagem de erro caso as senhas não sejam iguais
        return c.JSON(fiber.Map{
            "message": "A senha está incorreta",
        })
    }

    // Gera um hash seguro para a senha utilizando o algoritmo bcrypt
    password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

    // Cria uma nova estrutura `User` a partir dos dados do formulário
    user := models.User{
        FirstName: data["first_name"],
        LastName:  data["last_name"],
        Email:     data["email"],
        Password:  password,
    }

    // Insere o novo usuário no banco de dados
    database.DB.Create(&user)

    // Retorna o usuário recém-criado como JSON na resposta
    return c.JSON(user)
}

// Função `Login` que recebe o contexto da requisição (`c`) e retorna um erro (`error`)
func Login(c *fiber.Ctx) error {
    // Cria um mapa de string para string para armazenar os dados do formulário
    var data map[string]string

    // Tenta ler os dados do corpo da requisição e verifica se há erro
    if err := c.BodyParser(&data); err != nil {
        // Retorna o erro caso a leitura falhe
        return err
    }

    // Cria uma variável para armazenar o usuário obtido do banco de dados
    var user models.User

    // Busca o usuário no banco de dados pelo e-mail informado
    database.DB.Where("email = ?", data["email"]).First(&user)

    // Retorna o usuário encontrado como JSON na resposta
    return c.JSON(user)
}
  

