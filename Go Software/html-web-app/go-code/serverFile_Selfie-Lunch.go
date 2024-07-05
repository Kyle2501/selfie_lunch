# Selfie_Lunch ~

package main

// .
  
import (
  "os"
  "log"
		
  "text/template"
  "net/http"
		
  "cloud.google.com/go"
    
    
)



type htmlPageData struct {
    pageTitle string
    pagePath string
    pageList []pageNav
    
}


type pageNav struct {
    pageTitle string
    pageLink string
}



func app_welcome_center_page() {
    
    
}



// . appHandler
func appHandler(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/app" {
    	http.NotFound(w, r)
    	return
    }
    
// ,

  pageTitle := "~ Selfie_Lunch // - Website App"
  pagePath := r.URL.Path
  
  pageType := ".."
  
  
  pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
      pageList: []pageNav {
          { pageTitle: "one", pageLink: "one"},
          { pageTitle: "two", pageLink: "two"},
          { pageTitle: "three", pageLink: "three"},
      },
  	
  }  //. .  pageData
  
  
  if pagePath == "/user" {
      pageTitle = "User Page"
      pageList = pageList
  }
  
  if pagePath == "/account" {
      pageTitle = "Account Page"
      pageList = pageList
  }
  
  if pagePath == "/profile" {
      pageTitle = "Profile Page"
      pageList = pageList
  }
  
  
  
  if pagePath == "/portfolio" {
      pageTitle = "Portfolio Page"
      pageList = pageList
  }
  
  if pagePath == "/resume" {
      pageTitle = "Resume Page"
      pageList = pageList
  }
  
  if pagePath == "/settings" {
      pageTitle = "Settings Page"
      pageList = pageList
  }
  

 
  
  
  
  pageFilePath := template.Must(template.ParseFiles("layout_main_page.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  appHandler




//  .  html url routes 
//  .  as well as terminal cli logs

func main() {

  appName := "~ Selfie_Lunch // - Website App"


  http.HandleFunc("/", indexHandler)
  
  http.HandleFunc("/user", indexHandler)
  http.HandleFunc("/account", indexHandler)
  http.HandleFunc("/profile", indexHandler)
  
  http.HandleFunc("/portfolio", indexHandler)
  http.HandleFunc("/resume", indexHandler)
  
  http.HandleFunc("/settings", indexHandler)
  




// -- -
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
    log.Printf("Loading _webapp with default port")
  }
  
  
  log.Printf("_webapp is active and Listening on port %s", port)

  log.Printf("// -- - %s", appName)
  log.Printf("_webapp now loaded and running at http://localhost:%s", port)

// -- - 
  if err := http.ListenAndServe(":"+port, nil); err != nil {
    log.Fatal("Error Starting the HTTP Server :", err)
    return
  }


}
