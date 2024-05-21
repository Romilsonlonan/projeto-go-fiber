package controllers // Declara pacote `controllers`

import ( // Importa dependências
	"fiber-project/database" // Importa
	// Importa pacote `models` do diretório `fiber-project`
	"fiber-project/models"
	"math/rand"

	"github.com/gofiber/fiber/v2" // Importa pacote `fiber` da versão 2
)

// Função `Register` que recebe o contexto da requisição (`c`) e retorna um erro (`error`)
func Forgot(c *fiber.Ctx) error {
    // Cria um mapa de string para string para armazenar os dados do formulário
    var data map[string]string

    // Tenta ler os dados do corpo da requisição e verifica se há erro
    if err := c.BodyParser(&data); err != nil {
        // Retorna o erro caso a leitura falhe
        return err
    }

    token := RandStringRunes(12)
	passwordReset := models.PasswordReset{
		Email: data["email"],
		Token: token,

	}

	database.DB.Create(&passwordReset)

    return c.JSON(fiber.Map{
        "message": "Sucesso!!!",
    })
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

