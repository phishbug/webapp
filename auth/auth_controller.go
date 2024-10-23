package auth

import (
    "fmt"
    "encoding/json"
    "net/http"
    "webapp/types"
    "webapp/helpers"
    "time"

    "github.com/dgrijalva/jwt-go"
)

// Secret key used to sign the JWT
var secretKey = []byte("thisisphishbugwewillbegettingsucessin2015")

// GenerateJWT creates a new JWT token
func GenerateJWT(username string) (string, error) {
    claims := jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expiration time
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}

// ValidateJWT checks the token and returns the claims
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return secretKey, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }
    return nil, fmt.Errorf("invalid token")
}

// AuthMiddleware returns a middleware function that validates JWT
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Authorization header required", http.StatusUnauthorized)
            return
        }

        claims, err := ValidateJWT(token)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Optionally, set the username in the context for use in the handler
        r.Header.Set("Username", claims["username"].(string))
        next(w, r) // Call the next handler
    }
}


// Handler to demonstrate generating a token
func LoginHandler(w http.ResponseWriter, r *http.Request) {

    var data types.LoginData

    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    
    if data.Username != "kunalthool23031986@phishbug.com"  && data.Password != "kunalthool23031986@phishbug.com"{
        // Write the JSON response
        helpers.SendErrorResponse(w, http.StatusInternalServerError,  "Invalid Auth User")
        return
    }

    token, err := GenerateJWT(data.Username)
    if err != nil {
        http.Error(w, "Could not generate token", http.StatusInternalServerError)
        return
    }

    helpers.SendOkResponse(w, "Request successful", map[string]string{"token": token, "admin_name": "Kunal"})
}

// Protected handler that requires JWT authentication
func protectedHandler(w http.ResponseWriter, r *http.Request) {
    username := r.Header.Get("Username")
    w.Write([]byte(fmt.Sprintf("Welcome, %s!", username)))
}