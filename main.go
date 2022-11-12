package main

import (
	_interface "ClientServer/interface"
	"ClientServer/repositories"
	_ "database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"strings"
)

var (
	pgUser         = "secret"
	pgPassword     = "secret"
	pgHost         = "localhost"
	pgPort         = "5432"
	pgDB           = "jwt"
	pgUserTable    = "users"
	pgLoginField   = "login"
	pgRoleTable    = "roles"
	pgRoleIdField  = "role_id"
	userRepository _interface.UserRepository
	hmacSecret     = []byte("c4bd7d88edb4fa1817abb11707958924384f7933e5facfd707dc1d1429af9936")
	port           = 9096
)

type Token struct {
	token  string
	host   string
	siteId string
}

func main() {

	var err error

	userRepository, err = repositories.NewUserRepositoryPG(pgUser, pgPassword, pgHost, pgPort, pgDB)

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/api/v1/auth/reg", RegUser)
	http.HandleFunc("/api/v1/user/pay_token", PayToken)
	http.HandleFunc("/api/v1/user/pay_token/status", PayTokenStatus)
	http.HandleFunc("/api/v1/user/phone/sms", PhoneSms)
	http.HandleFunc("/api/v1/user/pay", Pay)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func RegUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	phone := r.FormValue("phone")
	passwords := r.FormValue("passwords")

	if phone == "" || passwords == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if !userRepository.CheckOccupancyPhone(phone) {
		result := userRepository.UserRegistration(phone, passwords)
		if !result {
			http.Error(w, "", http.StatusBadGateway)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"mes": "successfully",
		})
	}

	http.Error(w, "", http.StatusUnprocessableEntity)
}

func PayToken(w http.ResponseWriter, r *http.Request) {

	//id := r.Header.Get("User-Id")
	//
	//var re Repository.UserRepositoryPG
	//var st, err = re.GetTokenbyId(id)
	//
	//if err != nil {
	//	return nil
	//}
	//
	//var datebd string
	//t, _ := time.Parse(st.Expired_date, datebd)
	//t2 := time.Now()
	//dur := t2.Sub(t)
	//if dur*time.Hour > 24 {
	//	return err
	//}
	//
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"id":  id,
	//	"key": strconv.FormatInt(time.Now().Unix(), 10) + id,
	//	"exp": time.Now().Add(5 * time.Minute).Unix(),
	//})
	//_, _ = token.SignedString(hmacSecret)
	//
	//w.Header().Set("Content-Type", "application/json")
	//_ = json.NewEncoder(w).Encode(map[string]interface{}{
	//	"token":       0,
	//	"expiredDate": "string",
	//})
	//
	//return nil
}

func PayTokenStatus(w http.ResponseWriter, r *http.Request) {

	//if r.Method != http.MethodPost {
	//	http.Error(w, "", http.StatusMethodNotAllowed)
	//	return
	//}
	//
	//id := r.FormValue("id")
	//amount := r.FormValue("amount")
	//status := r.FormValue("status")
	//
	//if id == "" || amount == "" || status == "" {
	//	http.Error(w, "", http.StatusBadRequest)
	//	return
	//}
	//var re Repository.UserRepositoryPG
	//
	//_, err := re.UserRegistration(id, token, expresion)
	//if err != nil {
	//	http.Error(w, "", http.StatusForbidden)
	//	return
	//}
}

func PhoneSms(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Content-Type", "application/json")
	//
	//_ = json.NewEncoder(w).Encode(map[string]interface{}{
	//	"message": tokenString,
	//})
	//http.Error(w, "", http.StatusMethodNotAllowed)
	//return
}

func Pay(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//
	//_ = json.NewEncoder(w).Encode(map[string]interface{}{
	//	"message": tokenString,
	//})
	//http.Error(w, "", http.StatusMethodNotAllowed)
	//return
}

func getToken(r *http.Request) (string, int) {

	if r.Method != http.MethodPost {
		return "", http.StatusMethodNotAllowed
	}

	token := r.Header.Get("Authorization")

	if token == "" {
		return "", http.StatusUnauthorized
	}
	extractedToken := strings.Split(token, "Bearer ")

	if len(extractedToken) < 2 {
		return "", http.StatusUnauthorized
	}

	return extractedToken[1], http.StatusOK
}