package usecases


import (
    "fmt"
    "errors"
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/Reach-Insurance/go/db"
)

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request){
	var user db.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	database.Create(&user)
	respondToClient(w, 200, user, "User Registered successfully.")
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	_ = json.NewDecoder(r.Body).Decode(&credentials)
	
    var user db.User
    result := database.Where("username = ? AND password = ?", credentials.Username, credentials.Password).First(&user)
    rows := result.RowsAffected
    if rows > 0 {
        fmt.Printf("\nSigned in succeffully.\n")
        respondToClient(w, 200, user, "Sign in successful.")  
    }else{
        fmt.Println("Signed in failed.")
        respondToClient(w, 403, user, "Access denied.")  
    }
}

func userExists (identifier string) (bool, db.User, error) {
    //the identifier can be ID, phone, email, username
    var user db.User
    response := database.Where("id = ? OR username = ?", identifier, identifier).First(&user)                   
    numberOfRowsFound := response.RowsAffected
    userExists := numberOfRowsFound > 0
    
    if !userExists {
        if id, err := strconv.Atoi(identifier); err == nil {
            resp := database.Where("id = ?", uint(id)).First(&user)
            rowsFound := resp.RowsAffected
            exists := rowsFound > 0
            return exists, user, response.Error
        }else{
            return false, user, errors.New("user id must be a number")
        } 
    }else{
        return userExists, user, response.Error
    }
}

func ReadUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    identifier := params["id"]
    
    ok, user, err := userExists(identifier)
    if err != nil {
        respondToClient(w, 400, nil, err.Error())
    }
    
    if !ok {
        respondToClient(w, 404, nil, "Specified User not found")
    }
    
    respondToClient(w, 200, user, "")
}

func ReadAllUsers(w http.ResponseWriter, r *http.Request) {
    var users []db.User
    response := database.Find(&users)
    numberOfRowsFound := response.RowsAffected
    msg := fmt.Sprintf("Found %d users", numberOfRowsFound)
    respondToClient(w, 200, users, msg)
}

func GetAllCustomers (w http.ResponseWriter, r *http.Request) {
    var users []db.User
    response := database.Where("usertype = ?", "customer").Find(&users)
    numberOfRowsFound := response.RowsAffected
    msg := fmt.Sprintf("Found %d users", numberOfRowsFound)
    respondToClient(w, 200, users, msg)
}



