package userAuthentication

import 
(
"encoding/json" 
"net/http" 
)

func OauthGoogleLogin(w http.ResponseWriter, r *http.Request) {
  //Root

        successTest:="It Worked"  
        
 
   json.NewEncoder(w).Encode(successTest)

 } 

 func OauthGoogleCallback(w http.ResponseWriter, r *http.Request) {
    //Root
  
          successTest:="It Worked"  
          
   
     json.NewEncoder(w).Encode(successTest)
  
   } 