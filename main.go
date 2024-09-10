 // # ~ . ~
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


func world(w http.ResponseWriter, r *http.Request) {
	pagePath := r.URL.Path
   // fmt.Fprintf(w, "World!")
   // fmt.Fprintf(w, pagePath)
    
    pageData := pagePath
    pageName := "hi test"
    
    if pagePath == "/page/classSchedule" {
      pageName = "classSchedule Page"
 //     pageList = pageList
  }
  
  if pagePath == "/page/studentAttenance" {
    pageName = "studentAttenance Page"
    //     pageList = pageList
  }
  
    if pagePath == "/page/studentProfiles" {
    pageName = "studentProfiles Page"
    //     pageList = pageList
  }
  
    if pagePath == "/page/paymentInformation" {
    pageName = "paymentInformation Page"
    //     pageList = pageList
  }
  
      if pagePath == "/page/classEvent" {
    pageName = "classEvent Page"
    //     pageList = pageList
  }
  
      if pagePath == "/page/studentHomework" {
    pageName = "studentHomework Page"
    //     pageList = pageList
  }
  
      if pagePath == "/page/classQuestions" {
    pageName = "classQuestions Page"
    //     pageList = pageList
  }
  
        if pagePath == "/page/studentFeildtrip" {
    pageName = "studentFeildtrip Page"
    //     pageList = pageList
  }
  
        if pagePath == "/page/classMovies" {
    pageName = "classMovies Page"
    //     pageList = pageList
  }
  
        if pagePath == "/page/studentWorkshops" {
    pageName = "studentWorkshops Page"
    //     pageList = pageList
  }
  
       if pagePath == "/page/listClass" {
    pageName = "listClass Page"
    //     pageList = pageList
  }
  
       if pagePath == "/page/loadClass" {
    pageName = "loadClass Page"
    //     pageList = pageList
  }
  
       if pagePath == "/page/joinClass" {
    pageName = "joinClass Page"
    //     pageList = pageList
  }
  
         if pagePath == "/page/leaveClass" {
    pageName = "leaveClass Page"
    //     pageList = pageList
  }
  
         if pagePath == "/page/addHomework" {
    pageName = "addHomework Page"
    //     pageList = pageList
  }
  
         if pagePath == "/page/startClass" {
    pageName = "startClass Page"
    //     pageList = pageList
  }
  
           if pagePath == "/page/meetTeacher" {
    pageName = "meetTeacher Page"
    //     pageList = pageList
  }
  
           if pagePath == "/page/askQuestions" {
    pageName = "askQuestions Page"
    //     pageList = pageList
  }
  
           if pagePath == "/page/meetStudents" {
    pageName = "meetStudents Page"
    //     pageList = pageList
  }


// ,  ° . +
    data := SOSPageData{
            PageTitle: pageData,
            PagePath: pageData,
            PageName: pageName,
            SOSNav: []navList{
                {Title: "classSchedule", Done: false},
                {Title: "studentProfiles", Done: true},
                {Title: "paymentInformation", Done: true},
                {Title: "studentAttenance", Done: true},
                {Title: "classEvent", Done: true},
                {Title: "studentHomework", Done: true},
                {Title: "studentFeildtrip", Done: true},
                {Title: "classQuestions", Done: true},
                {Title: "classMovies", Done: true},
                {Title: "studentWorkshops", Done: true},
                {Title: "listClass", Done: true},
                {Title: "loadClass", Done: true},
                {Title: "joinClass", Done: true},
                {Title: "leaveClass", Done: true},
                {Title: "addHomework", Done: true},
                {Title: "startClass", Done: true},
                {Title: "meetTeacher", Done: true},
                {Title: "askQuestions", Done: true},
                {Title: "meetStudents", Done: true},
        
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
  appName := "~ . - // - Website App"
  
  http.HandleFunc("/hello", hello)
    http.HandleFunc("/world", world)
  // ,  ° . +
    http.HandleFunc("/page/classSchedule", world)
    http.HandleFunc("/page/studentFeildtrip", world)
    http.HandleFunc("/page/classMovies", world)
    http.HandleFunc("/page/studentProfiles", world)
    http.HandleFunc("/page/paymentInformation", world)
    http.HandleFunc("/page/studentAttenance", world)
    http.HandleFunc("/page/classEvent", world)
    http.HandleFunc("/page/studentHomework", world)
      http.HandleFunc("/page/classQuestions", world)
      http.HandleFunc("/page/studentWorkshops", world)
      http.HandleFunc("/page/listClass", world)
      http.HandleFunc("/page/loadClass", world)
      http.HandleFunc("/page/joinClass", world)
      http.HandleFunc("/page/leaveClass", world)
      http.HandleFunc("/page/addHomework", world)
      http.HandleFunc("/page/startClass", world)
         http.HandleFunc("/page/meetTeacher", world)
          http.HandleFunc("/page/askQuestions", world)
         http.HandleFunc("/page/meetStudents", world)
          

// ,  ° . +
  http.HandleFunc("/", indexHandler)
  
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