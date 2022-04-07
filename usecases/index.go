package usecases


import (
	"fmt"
    "net/http"
	"gorm.io/gorm"
	"encoding/json"
	"github.com/mahani-software-engineering/website/db"
)



var database *gorm.DB

func Init(){
    database, _ = db.Connect()
}


func respondToClient(w http.ResponseWriter, statusCode uint, edata interface{}, simpleMessage string){
    w.Header().Set("Content-Type","application/json")
    switch {
        case 200 <= statusCode && statusCode < 300:
            fmt.Println("200 <= statusCode && statusCode < 300")
            w.WriteHeader(http.StatusOK)
            if edata != nil && simpleMessage != "" {
                json.NewEncoder(w).Encode(struct{Message string; Data interface{}}{ Message: simpleMessage, Data:edata })
            }else if edata != nil && simpleMessage == "" {
                json.NewEncoder(w).Encode(struct{Data interface{}}{ Data:edata})
            }else if edata == nil && simpleMessage != "" {
                json.NewEncoder(w).Encode(struct{Message string}{ Message: simpleMessage})
            }else{
                json.NewEncoder(w).Encode(struct{Message string}{ Message: "Oops! Unexpected error occured." })
            }
        case statusCode == 403:
            fmt.Println("statusCode == 403")
            w.WriteHeader(http.StatusForbidden)
            json.NewEncoder(w).Encode(struct{Message string}{ Message: simpleMessage })
        case 400 <= statusCode && statusCode < 500:
            fmt.Println("400 <= statusCode && statusCode < 500")
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(struct{Message string}{ Message: simpleMessage })
        case 500 <= statusCode && statusCode < 600:
            fmt.Println("500 <= statusCode && statusCode < 600")
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(struct{Message string}{ Message: "Server error!" })
        default: 
            fmt.Println("default")
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(struct{Message string}{ Message: "!Fatal error!" })
    }
}

func SendSMS(msg, targetPhone string) {
   smsAPI := fmt.Sprintf("http://www.egosms.co/api/v1/plain/?number=%s&message=%s&username=%s&password=%s&sender=Egosms", targetPhone, msg, "Marvin", "0700617624")
   resp, er := http.Get(smsAPI)
   if er != nil {
      fmt.Println("Error while sending the SMS ->", er)
   }
   defer resp.Body.Close()
   fmt.Println("resp = ", resp)
}

var metaDefault string = "1"



/* === AfricasTalking ===
const url = `https://content.africastalking.com/version1/messaging`;
rq.headers({
            apiKey: "77ead081e39515784eaea566760321e85f5025a1a00dc688641bdec22b89c356",
            "Content-Type": "application/x-www-form-urlencoded"
        });
        
const body = {
                to: params.to,                      // +256702740380   or +256706123303,+256702740380
                from: params.from,                  // +256781224508
                message: params.message,            // Hello from system
                username: _self.options.username    // AkaboxiUg
            };
if(resp.status === 201) {
    // API returns CREATED on success!? 
}

*/
            
//===========sending SMS ====
   
