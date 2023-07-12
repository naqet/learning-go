package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/naqet/learning-go/jwt/initializers"
	"github.com/naqet/learning-go/jwt/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
    var body struct {
        Email string
        Password string
    }

    if c.Bind(&body) != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"});
        
        return;
    }

    hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10);

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash the password"});
        return;
    }

    user := models.User{Email: body.Email, Password: string(hash)};

    result := initializers.DB.Create(&user);

    if result.Error != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Failed to create a user"});
        return;
    }

    c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
    var body struct {
        Email string
        Password string
    }

    if c.Bind(&body) != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"});
        return;
    }
    var user models.User

    result := initializers.DB.First(&user, "email = ?", body.Email);

    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect email or password"})
        return
    }

    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect email or password"})
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
        "sub": user.ID,
        "exp": time.Now().Add(time.Hour).Unix(),
    })

    tokenString, err := token.SignedString([]byte(os.Getenv("JWT")))

    
    if err != nil {
        fmt.Println(err)
        c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Failed to create a token"});
        return;
    }

    c.SetSameSite(http.SameSiteLaxMode)
    c.SetCookie("Auth", tokenString, 60 * 60 * 60, "", "", true, true);
    c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Logged in",
    })
}
