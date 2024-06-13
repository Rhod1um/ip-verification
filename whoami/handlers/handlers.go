package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
) 


func Handler(w http.ResponseWriter, r *http.Request) {  
	if r.URL.Path != "/ip" {  
        http.NotFound(w, r)  
        return  
    }

	fmt.Printf("Handler called with URL: %s\n", r.URL.Path)  

	resp, err := http.Get("http://app1.example.com:30001/ip") // replace with app1's address and port. Use this when port-forwarding: k port-forward svc/app1service 8888:6666
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)  
		return  
	}  
	defer resp.Body.Close()  
  
	fmt.Printf("HTTP response status: %s\n", resp.Status)
	for name, headers := range resp.Header {
		for _, h := range headers {
			fmt.Printf("%v: %v\n", name, h)
		}
	}
  
	// Read the JSON body into a map  
	body := map[string]interface{}{}  
	err = json.NewDecoder(resp.Body).Decode(&body)  
	if err != nil {  
		http.Error(w, err.Error(), http.StatusInternalServerError)  
		return  
	}  
  
	// Modify the map  
	body["newKey"] = "newValue"  
  
	// Write the modified map back out as JSON  
	header := w.Header()  
	header.Set("Content-Type", "application/json")  
	err = json.NewEncoder(w).Encode(body)  
	if err != nil {  
		http.Error(w, err.Error(), http.StatusInternalServerError)  
		return  
	}   
}

func TestHandler(w http.ResponseWriter, r *http.Request) {  
	if r.URL.Path != "/test" {  
        http.NotFound(w, r)  
        return  
    }
	body := map[string]string{"test": "test"} 
	header := w.Header()  
	header.Set("Content-Type", "application/json")  
	err := json.NewEncoder(w).Encode(body)  
	if err != nil {  
		http.Error(w, err.Error(), http.StatusInternalServerError)  
		return  
	}
}