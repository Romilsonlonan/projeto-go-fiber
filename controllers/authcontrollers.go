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
