<div class="justify-content: center; align-items: center;">
  <img class="align-items: center" src="https://github.com/Romilsonlonan/projeto-go-fiber/assets/90980220/a2c8e078-efcf-4ba9-9cb2-4b9c70059c33" alt="Imagem Golang" style="width: 400px; height: 350px;">
</div>

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



