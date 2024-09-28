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

// ,  ° . +
type SOSPageData struct {
    PageTitle string
    PagePath string
    PageName string
    SOSNav     []navList
}


// . indexHandler,  ~ for Public Pages °
func indexHandler(w http.ResponseWriter, r *http.Request) {
// ,  ° . +
    if r.URL.Path != "/" {
    	http.NotFound(w, r)
    	return
    }

// , ° . +
  pageTitle := "~ . - // - Website App"
  pagePath := r.URL.Path
  // pageType := ".."
  
 
// ,  ° . +
  pageFilePath := template.Must(
    template.ParseFiles(pageHTML))
  pageFilePath.Execute(w, pageData)
  
}  //  .  indexHandler


func hello(w http.ResponseWriter, r *http.Request) {
//	pagePath := r.URL.Path
    fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
//	pagePath := r.URL.Path
    fmt.Fprintf(w, "World!")
}

// ,  ° . +
func kitchenPage_world(w http.ResponseWriter, r *http.Request) {
	pagePath := r.URL.Path
   // fmt.Fprintf(w, "World!")
   // fmt.Fprintf(w, pagePath)
    
    pageData := pagePath
    pageName := "hi test"
    
    if pagePath == "/page/timer" {
      pageName = "Timer Page"
 //     pageList = pageList
  }
  
  if pagePath == "/page/amount_conversion" {
    pageName = "Amount Conversion Page"
    //     pageList = pageList
  }


    if pagePath == "/needs_list-a/Equiptment" {
      pageName = "Equiptment Page"
 //     pageList = pageList
  }

    if pagePath == "/needs_list-a/Employees" {
      pageName = "Employees Page"
 //     pageList = pageList
  }
  
     if pagePath == "/needs_list-a/Ingredients" {
      pageName = "Ingredients Page"
 //     pageList = pageList
  } 
  
      if pagePath == "/needs_list-a/Process" {
      pageName = "Process Page"
 //     pageList = pageList
  }
  
      if pagePath == "/needs_list-a/Menu" {
      pageName = "Menu Page"
 //     pageList = pageList
  }
  
        if pagePath == "/needs_list-a/Vehicle" {
      pageName = "Vehicle Page"
 //     pageList = pageList
  }
  
        if pagePath == "/needs_list-a/Software" {
      pageName = "Software Page"
 //     pageList = pageList
  }
  
        if pagePath == "/needs_list-a/Brand" {
      pageName = "Brand Page"
 //     pageList = pageList
  }
  
        if pagePath == "/needs_list-a/Sales" {
      pageName = "Sales Page"
 //     pageList = pageList
  }

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
  
 // ,  ° . +
    data := SOSPageData{
            PageTitle: pageData,
            PagePath: pageData,
            PageName: pageName,
            SOSNav: []navList{
                {Title: "Timer", Done: false},
                {Title: "Amount Conversion", Done: true},
                {Title: "Recipe Stuff", Done: true},
                {Title: "Camera", Done: true},
                {Title: "Notes", Done: true},
                {Title: "Session Log", Done: true},
        
            },
        }
 
 pageHTML := "layout_main_page.html";
 
  // ,  ° . +
  pageFilePath := template.Must(
    template.ParseFiles(pageHTML))
  pageFilePath.Execute(w, data)
}
   
   
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