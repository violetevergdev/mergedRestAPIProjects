package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"muffin_rest/internal/database/postgres"
	"muffin_rest/internal/models"

	"github.com/dgrijalva/jwt-go"
)

func GetUSDFuncsShares(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow cross domain access
	w.Header().Set("Content-Type", "application/json") // Set the content type

	var ArrShares = myCurrentFunds("ETF")
	json.NewEncoder(w).Encode(ArrShares) // Encode the data into JSON format
}

func myCurrentFunds(fundType string) []models.Funds {
	var amountShares []models.Funds 

	db := postgres.DBConnect() // Connect to the database
	defer db.Close() // Close the connection

	rows, err := db.Query("SELECT * FROM fundsusd WHERE type=$1",fundType)
	if err != nil {
		log.Fatal("Query error: ", err)
	}
	defer rows.Close()
	for rows.Next() {
		f := models.Funds{}

		err = rows.Scan(&f.Id, &f.Name, &f.Ticker, &f.Amount, &f.PricePerItem, &f.PurchasePrice, &f.PriceCurrent, &f.PercentChanges, &f.YearlyInvestment, &f.ClearMoney, &f.DatePurchase, &f.DateLastUpdate, &f.Type)
		if err!= nil {
			log.Fatal("Scan error: ", err)
		}

		amountShares = append(amountShares, f)
	}
	return amountShares
}


// Login

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow cross domain access
	w.Header().Set("Content-Type", "application/json") // Set the content type

	var u models.User
	json.NewDecoder(r.Body).Decode(&u)

	checkLogin(u)
}

func checkLogin(u models.User) string {
	if models.UserMock.Username != u.Username || models.UserMock.Password != u.Password {
		fmt.Println("Login error")
		err := "error"
		return err
	}

	validTocken, err := generateJWT() // Generate a new token for the user
	fmt.Println(validTocken)
	if err!= nil {
		fmt.Println("Error generating token")
		return "error"
	}

	return validTocken
}

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * 1000).Unix()
	claims["username"] = "John"

	tokenStr, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Fatal("Error signing token")
	}
	return tokenStr, err
}

//Check token (Auth)
func CheckToken(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        if r.Header["Token"] != nil {

            token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                    return nil, fmt.Errorf("There was an error")
                }
                return []byte("secret"), nil
            })

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            if token.Valid {
                endpoint(w, r)
            }
        } else {

            fmt.Fprintf(w, "Not Authorized")
        }
    })
}