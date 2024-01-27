package controllers

import (
	"gin/models"
	"gin/repository"
	"gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDriver(c *gin.Context) {

	requestID, _ := c.Get("RequestID")

	emailExist, err := repository.CheckExistsEmailInDrivers(c.PostForm("email"))

	if emailExist {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Este email já existe.",
			"error":   err.Error(),
		})

		return

	}

	_, err = utils.SendMessageOfVerifyToEmailInAwsSes(c.PostForm("email"))

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao encontrar email.",
			"error":   err.Error(),
		})

		return

	}

	respOfAwsBucket, err := utils.SaveQRCodeOfDriver(c.PostForm("cpf"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": respOfAwsBucket,
			"error":   err.Error(),
		})

		return
	}

	driver, endereco := utils.GetDriverAndAdressFromRequest(c, respOfAwsBucket)

	validateDocs, documentError := utils.ValidateDocsDriver(driver, endereco)

	if !validateDocs {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "type and try insert your documents again, please.",
			"error":   documentError,
		})

		return

	}

	_, err = repository.SaveDriver(driver, endereco)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error when inserting into database",
			"error":   err.Error(),
		})

		return

	}

	c.JSON(http.StatusCreated, gin.H{
		"requestID":   requestID,
		"status":      "driver created successfully",
		"s3bucketurl": respOfAwsBucket,
		"email":       "Por favor, confirme o email.",
	})

	return

}

func GetDriver(c *gin.Context) {

	requestID, _ := c.Get("RequestID")

	cpf := c.Param("cpf")

	driver, err := repository.FindDriverByCpf(cpf)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"requestID": requestID,
			"error":     err.Error(),
			"message":   "Error while searching in database",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"requestID":    requestID,
		"dataOfDriver": driver,
	})

}

func UpdateDriver(c *gin.Context) {

	requestID, _ := c.Get("RequestID")

	resp, _ := utils.VerifyCpf(c)

	if !resp {
		c.JSON(http.StatusBadRequest, gin.H{
			"requestID": requestID,
			"error":     "Security breach, intruder account trying to delete account.",
			"message":   "Invalid Cpf",
		})

		return

	}

	var update models.UpdateDriver

	update.Email = c.PostForm("email")
	update.Endereco.Rua = c.PostForm("rua")
	update.Endereco.Numero = c.PostForm("numero")
	update.Endereco.Complemento = c.PostForm("complemento")
	update.Endereco.Cidade = c.PostForm("cidade")
	update.Endereco.Estado = c.PostForm("estado")
	update.Endereco.CEP = c.PostForm("CEP")

	resp, err := repository.UpdateDriver(c, &update)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"requestID": requestID,
			"error":     err.Error(),
			"message":   "Error whiling update client",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"requestID": requestID,
		"message":   "user updated success.",
	})

}

func DeleteDriver(c *gin.Context) {

	requestID, _ := c.Get("RequestID")

	resp, _ := utils.VerifyCpf(c)

	if !resp {
		c.JSON(http.StatusBadRequest, gin.H{
			"requestID": requestID,
			"error":     "Security breach, intruder account trying to delete account.",
			"message":   "Invalid Cpf",
		})

		return

	}

	cpf := c.Param("cpf")

	emailOfUserToDeleteInAwsSes, err := repository.DeleteDriverByCpf(cpf)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"requestID": requestID,
			"error":     err.Error(),
			"message":   "Error while deleting in database",
		})

		return
	}

	_, err = utils.DeleteQRCodeOfUser(cpf)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"requestID": requestID,
			"error":     err.Error(),
			"message":   "Error when deleting qrcode of user",
		})

		return
	}

	_, err = utils.DeleteEmailFromAwsSes(emailOfUserToDeleteInAwsSes)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"requestID": requestID,
			"error":     err.Error(),
			"message":   "Error when deleting user email of SES",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"requestID": requestID,
		"message":   "User deleted w/ success",
	})

}
