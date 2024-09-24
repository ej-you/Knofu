package serializers

import (
	"fmt"

	echo "github.com/labstack/echo/v4"
	validate "github.com/gobuffalo/validate/v3"

	"github.com/Danil-114195722/Knofu/user/models"
	"github.com/Danil-114195722/Knofu/user/services"
	coreDB "github.com/Danil-114195722/Knofu/core/db"
	coreValidator "github.com/Danil-114195722/Knofu/core/validator"
)


// структура для входных данных регистрации юзера
type RegisterUserIn struct {
	Email 		string `json:"email" validate:"required|email"`
	FirstName 	string `json:"firstName" validate:"required"`
	LastName 	string `json:"lastName" validate:"required"`
	Password 	string `json:"password" validate:"required|min:8|max:50"`
}

// базовая валидация полей по тегам
func (self *RegisterUserIn) IsValid(errors *validate.Errors) {
	coreValidator.BaseValidator(self, errors)
}

// более глубокая валидация с возвратом ошибок валидации
func (self *RegisterUserIn) Validate() error {
	// базовая валидация полей по тегам
	var validateErrors *validate.Errors = validate.Validate(self)

	if len(validateErrors.Errors) > 0 {
		// словарь для ошибок
		errMap := make(map[string]string, len(validateErrors.Errors))

		for key, value := range validateErrors.Errors {
			errMap[key] = value[0]
		}
		// возвращаем *echo.HTTPError
		httpError := echo.NewHTTPError(400, errMap)
		return httpError
	}

	// TODO: ДОБАВИТЬ ПРОВЕРКУ НА НАЛИЧИЕ ТАКОГО ЮЗЕРА В БД

	return nil
}

// создание нового юзера в БД
func (self *RegisterUserIn) Create() (models.User, error) {
	// получаем хэш пароля
	hashPasswd, err := services.EncodePassword(self.Password)
	if err != nil {
		return models.User{}, err
	}

	newUser := models.User{
		Email: self.Email,
		FirstName: self.FirstName,
		LastName: self.LastName,
		Password: hashPasswd,
	}

	dbConnect, err := coreDB.GetConnection()
	if err != nil {
		return models.User{}, echo.NewHTTPError(500, map[string]string{"dbConnect": "Failed to connect to DB"})
	}

	// TODO: ДОБАВИТЬ ПРОВЕРКУ НА ДУБЛИРОВАНИЕ EMAIL

	createResult := dbConnect.Create(&newUser)
	if err = createResult.Error; err != nil {
		fmt.Println(err)
		return models.User{}, err
	}


	return newUser, nil
}


// структура для выходных данных регистрации юзера
type RegisterUserOut struct {
	ID			uint64 `json:"id"`
	Email 		string `json:"email"`
	FirstName 	string `json:"firstName"`
	LastName 	string `json:"lastName"`
	Token		string `json:"token"`
}

// формирование структуры для ответа
func GetOutStruct(newUser models.User) (RegisterUserOut, error) {
	// получение токена для юзера
	encryptedToken, err := services.GetJWTToken(newUser.ID)
	if err != nil {
		return RegisterUserOut{}, err
	}

	userOut := RegisterUserOut{
		ID: newUser.ID,
		Email: newUser.Email,
		FirstName: newUser.FirstName,
		LastName: newUser.LastName,
		Token: encryptedToken,
	}

	return userOut, nil
}
