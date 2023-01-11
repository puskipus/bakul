package authcontroller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hanifhahn/bakul/config"
	"github.com/hanifhahn/bakul/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	// data user dari input
	var userLogin model.User
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// cek user dari input == user dari database
	var user model.User
	if err := model.DB.Where("email = ?", userLogin.Email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Username atau Password Salah"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	// cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username atau Password Salah"})
		return
	}

	// buat jwt
	expTime := time.Now().Add(time.Hour * 24)
	claims := &config.JWTClaim{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// buat algoritma login
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		// log.Fatal("ERROR")
		return
	}

	// set token ke cookie
	c.SetCookie("token", token, 24*3600, "/", "localhost", false, true)

	// respone success
	c.JSON(http.StatusOK, gin.H{"message": "login berhasil"})

}

func Register(c *gin.Context) {
	// Bind user dari input
	var newUser model.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// create hash password
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(hashPass)

	// buat user baru
	if err := model.DB.Create(&newUser).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// response user berhasil di Add
	c.JSON(http.StatusOK, gin.H{"message": "Registrasi Berhasil", "data": newUser})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "logout berhasil"})
}
