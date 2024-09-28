 // # ~ . ~ Selfie Lunch Software
   // ,  ° . +
   package main
   
   import (
       "os"
       "log"
       
       "fmt"
   		
   	"text/template"
   	"net/http"
   	
   //	"time"
   
   )
   
   
   <div class="concept_needs_list">
   	<p><b>_ Concept Needs List A</b><ul>
   		<li>_ Equiptment</li>
   		<li>_ Employees</li>
   		<li>_ Ingredients </li>
   		<li>_ Process</li>
   		<li>_ Menu</li>
   		<li>_ Vehicle</li>
   		<li>_ Software</li>
   		<li>_ Brand</li>
   		<li>_ Sales</li>
   	</ul></p>
   </div><!-- - . concept_needs_list - -->
   	
   <div class="concept_needs_list">
     <p><b>_ Concept Needs List B</b><ul>
      <li>_ People</li>
     <li>_ Booth</li>
     <li>_ Product</li>
       <li>_ Brand</li>
       <li>_ Software</li>
   <li>_ Sales</li>	
      </ul></p>
      
      
      
      //  .  html url routes 
   //  .  as well as terminal cli logs
   
   func main() {
   // ,  ° . +
     appName := "~ . - Selfie Lunch Software // - Website App"
     
     http.HandleFunc("/hello", hello)
       http.HandleFunc("/world", world)
     // ,  ° . +
        
     // ,  ° . +
     http.HandleFunc("/one", testHandler)
   
   // . ° ~ +
    // http.HandleFunc("/page/classSchedule", indexHandler)
   
   // -- -
     port := os.Getenv("PORT")
     if port == "" {
       port = "8080"
       log.Printf("Loading _webapp with default port")
     }
     
   // ,  ° . +
     log.Printf("_webapp is active and Listening on port %s", port)
   
     log.Printf("// -- - %s", appName)
     log.Printf("_webapp now loaded and running at http://localhost:%s", port)
   
   // -- - 
     if err := http.ListenAndServe(":"+port, nil); err != nil {
       log.Fatal("Error Starting the HTTP Server :", err)
       return
     }}