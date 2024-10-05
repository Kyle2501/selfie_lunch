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
type htmlPageData struct {
    pageTitle string
    pagePath string
    pageList []pageNav
    
}

type pageNav struct {
    pageTitle string
    pageLink string
}

type pageList struct {
    pageTitle string
    pageLink string
}


type navList struct {
    Title string
    Done  bool
}


// ,  ° . +
type SOSPageData struct {
    PageTitle string
    PagePath string
    PageName string
    PageLayout string
    SOSNav     []navList
}




// ,  ° . +
func app_welcome_center_page() {


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
pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
    //  pageList: []pageNav {
  //        { pageTitle: "one", pageLink: "one"},
//          { pageTitle: "two", pageLink: "two"},
   //       { pageTitle: "three", pageLink: "three"},
    //  },
  	
  }  //. .  pageData
  
  pageHTML := "layout_main_page.html";
  if pagePath == "/" {
      pageTitle = "Index Page"
 //     pageList = pageList
  }
  
    if pagePath == "/page/classSchedule" {
      pageTitle = "classSchedule Page"
 //     pageList = pageList
 pageHTML = "/Go-Software/html-web-app/html-files/layout_sosbot_classSchedule.html";
  }


// ,  ° . +
  pageFilePath := template.Must(
    template.ParseFiles(pageHTML))
  pageFilePath.Execute(w, pageData)
  
}  //  .  indexHandler

// . testHandler,  ~ for Public Pages °
func testHandler(w http.ResponseWriter, r *http.Request) {
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
pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
    //  pageList: []pageNav {
  //        { pageTitle: "one", pageLink: "one"},
//          { pageTitle: "two", pageLink: "two"},
   //       { pageTitle: "three", pageLink: "three"},
    //  },
  	
  }  //. .  pageData
  
  pageHTML := "layout_main_page.html";
  if pagePath == "/" {
      pageTitle = "Index Page"
 //     pageList = pageList
  }
 


// ,  ° . +
  pageFilePath := template.Must(
    template.ParseFiles(pageHTML))
  pageFilePath.Execute(w, pageData)
  
}  //  .  testHandler

func hello(w http.ResponseWriter, r *http.Request) {
//	pagePath := r.URL.Path
    fmt.Fprintf(w, "Hello!")
}

// pagePath := 'test'

func getPageInfo(x string) string {
  data := "hi" + x
  data = "hello" + data
  return data
}

func world(w http.ResponseWriter, r *http.Request) {
   pagePath := "test"
   if r.URL.Path != "/" {
   	pagePath = r.URL.Path
   }
    
    fmt.Fprintf(w, `<!doctype html>
<html>
<head>
<title>~ %s</title>
<meta name="viewport" content="width=device-width, initial-scale=1" />

<style>
    
body { padding-bottom: 175px; }
    
</style>
</head>
    `, pagePath)
    
    topBar := `<body>
    <div class="top_bar">
      <p><b>#! - Selfie Lunch Concepts</b></p>
      <hr />
    </div><!-- - . top_bar - -->`
    
    fmt.Fprintf(w, topBar)
    
    fmt.Fprintf(w, "World! %s", pagePath)
    
    pageInfo := getPageInfo("hey")
    fmt.Fprintf(w, pageInfo)
    
    fmt.Fprintf(w, `<footer>
       <code>- footer</code>
     </footer>
</body>
</html>`)
    
    
} // - end world

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
  
    if pagePath == "/page/recipe_stuff" {
    pageName = "Recipe Stuff Page"
    //     pageList = pageList
  }
  
    if pagePath == "/page/camera" {
    pageName = "Camera Page"
    //     pageList = pageList
  }
  
      if pagePath == "/page/nbotes" {
    pageName = "Notes Page"
    //     pageList = pageList
  }
  
      if pagePath == "/page/session_log" {
    pageName = "Session Log Page"
    //     pageList = pageList
  }
  

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
 
 pageHTML := "layout_kitchen_page.html";
 
  // ,  ° . +
  pageFilePath := template.Must(
    template.ParseFiles(pageHTML))
  pageFilePath.Execute(w, data)
}


// - . -
func worldLoader(w http.ResponseWriter, r *http.Request) {
	pagePath := r.URL.Path
    // fmt.Fprintf(w, "World!")
    pageName := "test"
    
    pageLayout := `
        <div>hi test .</div>
    `
    
  if pagePath == "/page/Menu" {
    pageName = "Menu Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
  if pagePath == "/page/Location" {
    pageName = "Location Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
  if pagePath == "/page/Equiptment" {
    pageName = "Equiptment Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
  if pagePath == "/page/Staff" {
    pageName = "Staff Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
  if pagePath == "/page/Services" {
    pageName = "Services Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
  if pagePath == "/page/Music" {
    pageName = "Music Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
  if pagePath == "/page/Setting" {
    pageName = "Setting Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
  if pagePath == "/page/Sources" {
    pageName = "Sources Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
  if pagePath == "/page/Marketing" {
    pageName = "Marketing Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
  if pagePath == "/page/Operations" {
    pageName = "Operations Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
  if pagePath == "/page/Software" {
    pageName = "Software Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
  if pagePath == "/page/Events" {
    pageName = "Events Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
  if pagePath == "/page/Seating" {
    pageName = "Seating Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
  if pagePath == "/page/Hours" {
    pageName = "Hours Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
    // ,  ° . +
  //  http.HandleFunc("/concepts/index", worldLoader)
  if pagePath == "/concepts/index" {
    pageName = "Hours Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
    //  http.HandleFunc("/concepts/Loco-Moco-Bus", worldLoader)
  if pagePath == "/concepts/Loco-Moco-Bus" {
    pageName = "Loco-Moco-Bus Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
  //  http.HandleFunc("/concepts/Kitchen-Page", worldLoader)
  if pagePath == "/concepts/Kitchen-Page" {
    pageName = "Kitchen-Page Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
 //   http.HandleFunc("/concepts/Honolulu-Lemonade", worldLoader)
  if pagePath == "/concepts/Honolulu-Lemonade" {
    pageName = "Honolulu-Lemonade Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
 //   http.HandleFunc("/concepts/Focaccia-Hands", worldLoader)
      if pagePath == "/concepts/Focaccia-Hands" {
    pageName = "Focaccia-Hands Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
    
   // ,  ° . +
 //   http.HandleFunc("/courses/index", worldLoader)
  if pagePath == "/courses/index" {
    pageName = "Hours Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
 //   http.HandleFunc("/courses/CSS-Layout", worldLoader)
  if pagePath == "/courses/CSS-Layout" {
    pageName = "CSS-Layout Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
 //   http.HandleFunc("/courses/Culniary-Manager", worldLoader)
  if pagePath == "/courses/Culniary-Manager" {
    pageName = "Culniary-Manager Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
 //   http.HandleFunc("/courses/Food-Handler", worldLoader)
   if pagePath == "/courses/Food-Handler" {
    pageName = "Food-Handler Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
//   http.HandleFunc("/courses/System-Administrator", worldLoader)
    if pagePath == "/courses/System-Administrator" {
    pageName = "System-Administrator Page"
    //     pageList = pageList
    pageLayout = `
        <div>hi test .</div>
    `
  }
  
  
  // ,  ° . +
    data := SOSPageData{
            PageTitle: pageName,
            PagePath: pageName,
            PageName: pageName,
            PageLayout: pageLayout,
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
    http.HandleFunc("/page/kitchen", kitchenPage_world)
    
          http.HandleFunc("/", testHandler)
          
  // ,  ° . +
    http.HandleFunc("/concepts/index", world)
    http.HandleFunc("/concepts/Loco-Moco-Bus", world)
    http.HandleFunc("/concepts/Kitchen-Page", world)
    http.HandleFunc("/concepts/Honolulu-Lemonade", world)
    http.HandleFunc("/concepts/Focaccia-Hands", world)
    
    
   // ,  ° . +
    http.HandleFunc("/courses/index", world)
    http.HandleFunc("/courses/CSS-Layout", world)
    http.HandleFunc("/courses/Culniary-Manager", world)
    http.HandleFunc("/courses/Food-Handler", world)
    http.HandleFunc("/courses/System-Administrator", world)

// ,  ° . +
  http.HandleFunc("/Menu", world)
  http.HandleFunc("/Location", world)
  http.HandleFunc("/Equiptment", world)
  http.HandleFunc("/Staff", world)
  http.HandleFunc("/Services", world)
  http.HandleFunc("/Music", world)
  http.HandleFunc("/Setting", world)
  http.HandleFunc("/Sources", world)
  http.HandleFunc("/Marketing", world)
  http.HandleFunc("/Operations", world)
  http.HandleFunc("/Software", world)
  http.HandleFunc("/Events", world)
  http.HandleFunc("/Seating", world)
  http.HandleFunc("/Hours", world)
  
  http.HandleFunc("/page/Hours", world)
  

  
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