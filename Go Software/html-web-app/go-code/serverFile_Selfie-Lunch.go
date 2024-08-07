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



// . pageHandler, ~ for User Pages °
func pageHandler(w http.ResponseWriter, r *http.Request) {
// ,  ° . +

  pageTitle := "~ Selfie_Lunch - // - Website App"
  pagePath := r.URL.Path
  pageType := ".."


// ,  ° . +
pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
      pageList: []pageNav {
          { pageTitle: "one", pageLink: "one"},
          { pageTitle: "two", pageLink: "two"},
          { pageTitle: "three", pageLink: "three"},
      },
  	
  }  //. .  pageData
  
  
  // ,  ° . +
  if pagePath == "/user" {
      pageTitle = "User Page"
      pageList = pageList
  }
  
  // ,  ° . +
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
  
  
// ,  ° . +
  if pagePath == "/portfolio" {
      pageTitle = "Portfolio Page"
      pageList = pageList
  }
  
  if pagePath == "/resume" {
      pageTitle = "Resume Page"
      pageList = pageList
  }
  
  
    // ,  ° . + ` <p>Cafe Needs<ul>
    
    
   // ,  ° . +  - _ ` Menu ~
  // ,  ° . +
  if pagePath == "/Menu" {
      pageTitle = "Menu Page"
      pageList = pageList
  }
  
  
    // ,  ° . + - _ ` Location ~
  // ,  ° . +
  if pagePath == "/Location" {
      pageTitle = "Location Page"
      pageList = pageList
  }
  
  
   // ,  ° . + - _ ` Equiptment ~
  // ,  ° . +
  if pagePath == "/Equiptment" {
      pageTitle = "Equiptment Page"
      pageList = pageList
  }
  
  
    // ,  ° . +- _ ` Staff ~
  // ,  ° . +
  if pagePath == "/Staff" {
      pageTitle = "Staff Page"
      pageList = pageList
  }
  
  
    // ,  ° . + - _ ` Services ~
  // ,  ° . +
  if pagePath == "/Services" {
      pageTitle = "Services Page"
      pageList = pageList
  }
  
  
    // ,  ° . + - _ ` Music ~
  // ,  ° . +
  if pagePath == "/Music" {
      pageTitle = "Music Page"
      pageList = pageList
  }
  
  
    // ,  ° . + - _ ` Setting ~
  // ,  ° . +
  if pagePath == "/Setting" {
      pageTitle = "Setting Page"
      pageList = pageList
  }
  
  
    // ,  ° . + - _ ` Sources ~
  // ,  ° . +
  if pagePath == "/Sources" {
      pageTitle = "Sources Page"
      pageList = pageList
  }
  
  
   // ,  ° . +  - _ ` Marketing ~
  // ,  ° . +
  if pagePath == "/Marketing" {
      pageTitle = "Marketing Page"
      pageList = pageList
  }
  
  
  // ,  ° . +  - _ ` Operations ~
  // ,  ° . +
  if pagePath == "/Operations" {
      pageTitle = "Operations Page"
      pageList = pageList
  }
  
  
    // ,  ° . + - _ ` Software ~
  // ,  ° . +
  if pagePath == "/Software" {
      pageTitle = "Software Page"
      pageList = pageList
  }
  
  
    // ,  ° . + - _ ` Events ~
  // ,  ° . +
  if pagePath == "/Events" {
      pageTitle = "Events Page"
      pageList = pageList
  }
  
  
   // ,  ° . +  - _ ` Seating ~
  // ,  ° . +
  if pagePath == "/Seating" {
      pageTitle = "Seating Page"
      pageList = pageList
  }
  
  
   // ,  ° . +  - _ ` Hours ~
  // ,  ° . +
  if pagePath == "/Hours" {
      pageTitle = "Hours Page"
      pageList = pageList
  }


// ,  ° . +
pageFilePath := template.Must(template.ParseFiles("layout_main_page.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  pageHandler


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
  
  // ,  ° . +` <p>Cafe Needs<ul>
http.HandleFunc("/Menu", indexHandler)
  // ,  ° . + - _ ` Menu ~
  
  http.HandleFunc("/Location", indexHandler)
  // ,  ° . + - _ ` Location ~
  
  http.HandleFunc("/Equiptment", indexHandler)
  // ,  ° . + - _ ` Equiptment ~
  
  http.HandleFunc("/Staff", indexHandler)
 // ,  ° . + - _ ` Staff ~
  
  http.HandleFunc("/Services", indexHandler)
 // ,  ° . + - _ ` Services ~
  
  http.HandleFunc("/Music", indexHandler)
  // ,  ° . +- _ ` Music ~
  
  http.HandleFunc("/Setting", indexHandler)
 // ,  ° . + - _ ` Setting ~
  
  http.HandleFunc("/Sources", indexHandler)
 // ,  ° . + - _ ` Sources ~
  
  http.HandleFunc("/Marketing", indexHandler)
  // ,  ° . +- _ ` Marketing ~
  
  http.HandleFunc("/Operations", indexHandler)
  // ,  ° . +- _ ` Operations ~
  
  http.HandleFunc("/Software", indexHandler)
// ,  ° . +  - _ ` Software ~
  
  http.HandleFunc("/Events", indexHandler)
 // ,  ° . + - _ ` Events ~
  
  http.HandleFunc("/Seating", indexHandler)
  // ,  ° . +- _ ` Seating ~
  
  http.HandleFunc("/Hours", indexHandler)
 // ,  ° . + - _ ` Hours ~



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
  }


}
