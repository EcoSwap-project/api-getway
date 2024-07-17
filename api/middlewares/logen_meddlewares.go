package middlewares

import (
	"fmt"
	"strings"

	"api_getway/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}


func JaWTMiddlewares(ctx *gin.Context) {
	jwtKey := []byte(config.Load().SECRET_KEY)

	tokenString := ctx.GetHeader("Authorization")

	// Check if the Authorization header is present
	if tokenString == "" {
		ctx.JSON(401, gin.H{"error": "Authorization header is missing"})
		ctx.Abort()
		return
	}

	// Remove "Bearer " prefix
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	} else {
		ctx.JSON(401, gin.H{"error": "Invalid Authorization header format"})
		ctx.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		ctx.JSON(401, gin.H{"error": "parsing token failed: " + err.Error()})
		ctx.Abort()
		return
	}

	if !token.Valid {
		ctx.JSON(401, gin.H{"error": "invalid token"})
		ctx.Abort()
		return
	}

	ctx.Next()
}

func JWTMiddlewares(ctx *gin.Context) {
	jwtKey := []byte(config.Load().SECRET_KEY)
	
	if ctx.Request.URL.Path == "/auth/login" || ctx.Request.URL.Path == "/auth/register" ||ctx.Request.URL.Path == "/auth/reset-password-access"||ctx.Request.URL.Path == "/auth/reset-password" || strings.HasPrefix(ctx.Request.URL.Path, "/swagger") {
	  ctx.Next()
	  return
	}
  
	tokenStr := ctx.GetHeader("Authorization")
	if tokenStr == "" {
	  ctx.JSON(401, gin.H{"error": "Authorization header is missing"})
	  ctx.Abort()
	  return
	}
  
	if !strings.HasPrefix(tokenStr, "Bearer ") {
	  ctx.JSON(401, gin.H{"error": "Invalid token format"})
	  ctx.Abort()
	  return
	}
  
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
  
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
	  if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	  }
	  return jwtKey, nil
	})
  
	if err != nil {
	  ctx.JSON(401, gin.H{"error": fmt.Sprintf("Invalid token: %v", err)})
	  ctx.Abort()
	  return
	}
  
	if !token.Valid {
	  ctx.JSON(401, gin.H{"error": "Token is not valid"})
	  ctx.Abort()
	  return
	}
  
	ctx.Set("user_id", claims.UserID)
	ctx.Set("username", claims.Username)
  
	ctx.Next()
  }
  