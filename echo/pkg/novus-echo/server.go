package novus_echo

import (
	"encoding/json"
	"net"
	"net/http"
)

type Response struct {  
	IP string `json:"ip"`  
} 

func GetIP(w http.ResponseWriter, r *http.Request) { 
	if r.URL.Path != "/ip" {  
        http.NotFound(w, r)  
        return  
    }
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {  
		http.Error(w, "Invalid address", http.StatusBadRequest)  
		return  
	}
	
	response := Response{IP: ip}  
  
	jsonResponse, err := json.Marshal(response)  
	if err != nil {  
		http.Error(w, err.Error(), http.StatusInternalServerError)  
		return  
	}  
  
	w.Header().Set("Content-Type", "application/json")  
	w.Write(jsonResponse) 
} 
