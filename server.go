package main 

import ("net/http"
		"fmt"	
		//"runtime"
		) 



func HelloServer(w http.ResponseWriter, req *http.Request) { 
		
		fmt.Fprintf(w, "Hello, world")
		//fmt.Println("reg")		
} 


func main() { 

	fmt.Println("Go server başlatılıyor...")
		
	//numCPU := runtime.NumCPU()
	//fmt.Println(numCPU," Adet CPU var okadar process oluşturuluyor...")
	
  	//runtime.GOMAXPROCS(2) 
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      	HelloServer(w, r)
    }) 
	http.ListenAndServe(":8080", nil) 
} 