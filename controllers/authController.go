package controllers // Declara pacote `controllers`

import ( // Importa dependências
	"fiber-project/database" // Importa
	"fiber-project/models"   // Importa pacote `models` do diretório `fiber-project`
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2" // Importa pacote `fiber` da versão 2
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
    jwt.StandardClaims
}

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

    // Verifica se o usuário foi encontrado
    if user.Id == 0 {
        // Define o status HTTP para "Não encontrado" (404)
        c.Status(404)

        // Retorna uma mensagem de erro ao cliente
        return c.JSON(fiber.Map{
            "message": "Usuário não encontrado",
        })
    }

    // Compara a senha informada com a senha do usuário no banco de dados
    if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
        // Define o status HTTP para "Requisição incorreta" (400)
        c.Status(400)

        // Retorna uma mensagem de erro ao cliente
        return c.JSON(fiber.Map{
            "message": "Senha Incorreta!",
        })
    }

    // Cria as claims do token JWT
    claims := jwt.RegisteredClaims{
        Issuer: strconv.Itoa(int(user.Id)), // Define o ID do usuário como o ID do token
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Define a data de expiração do token para 24 horas a partir de agora
    }

    // Cria um novo token JWT com as claims e o método de assinatura ES256
    jwtToken := *jwt.NewWithClaims(jwt.SigningMethodES256, claims)

    // Assina o token JWT com a chave secreta
    token, err := jwtToken.SignedString([]byte("secret"))

    // Verifica se ocorreu algum erro durante a assinatura do token
    if err != nil {
        // Retorna um erro interno do servidor
        return c.SendStatus(fiber.StatusInternalServerError)
    }

    // Cria um cookie para armazenar o token JWT
    cookie := fiber.Cookie{
        // Nome do cookie
        Name: "jwt",

        // Valor do cookie (o token JWT)
        Value: token,

        // Data de expiração do cookie (24 horas a partir de agora)
        Expires: time.Now().Add(24 * time.Hour),

        // Indica que o cookie só pode ser acessado pelo servidor (evita acesso por JavaScript)
        HTTPOnly: true,
    }

    // Adiciona o cookie ao contexto da requisição
    c.Cookie(&cookie)

    // Retorna o token JWT para o cliente em formato JSON
    return c.JSON(fiber.Map{
        "jwt": token,
    })

}


func User(c *fiber.Ctx) error {
    cookie := c.Cookies("jwt")
    token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte("secret"), nil
    })
    if err != nil || !token.Valid{
        c.Status(fiber.StatusUnauthorized)
        return c.JSON(fiber.Map{
            "message": "unauthenticated",
        })
    }    
    
    claims := token.Claims.(*Claims)
    id := claims.Issuer
    var user models.User 
    database.DB.Where("id = ?", id).First(&user)
    return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
    
    cookie := fiber.Cookie{
        Name: "jwt",
        Value: "",
        Expires: time.Now().Add(-time.Hour),
        HTTPOnly: true,
    }

    c.Cookie(&cookie)

    return c.JSON(fiber.Map{
        "message": "Sucesso!!!",
    })
}