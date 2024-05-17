package controllers // Declara pacote `controllers`

import ( // Importa dependências
	"fiber-project/database" // Importa
	"fiber-project/models"   // Importa pacote `models` do diretório `fiber-project`

	"github.com/gofiber/fiber/v2" // Importa pacote `fiber` da versão 2
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error { // Define função `Register` que recebe um contexto do Fiber e retorna um erro
    var data map[string]string

    if err := c.BodyParser(&data); err != nil {
        return err
    }

    c.Status(400)
    if data["password"] != data["confirm_password"]{
        return c.JSON(fiber.Map{
            "message": "Password is incorrect",
        })
    }

    password,_:= bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

    user := models.User{ 
        FirstName: data["first_name"],
        LastName: data["last_name"],
        Email: data["email"],
        Password: password,
    }
    
    database.DB.Create(&user)
    // Retorna o JSON da estrutura `user` no contexto da requisição
    return c.JSON(user)
}
