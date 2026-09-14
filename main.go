package main
import ("fmt"; "net/http")
func main(){
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
   w.Header().Set("Access-Control-Allow-Origin","*")
   fmt.Fprint(w, "Go Broken Link Checker is Live! Use /check?url=...")
 })
 fmt.Println("Server running...")
 http.ListenAndServe(":8080", nil)
}
