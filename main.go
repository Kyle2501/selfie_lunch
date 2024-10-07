 // # ~ . ~ Selfie Lunch Software
// ,  ° . +
package main

import (
    "os"
    "log"
    
    "fmt"
    "strings"
		
	"text/template"
	"net/http"
	
    "time"

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

func getPageData(w http.ResponseWriter, r *http.Request) {
//	pagePath := r.URL.Path
pagePath := r.URL.Path
    var pathLayers = strings.Split(pagePath, "/")
    data := pathLayers[2]

      if data == "Focaccia-Hands" {
    data = `
    <p><ul>
      <li>_ People &nbsp; <button onclick="getPageInfo('People')">Open</button></li>
      <li>_ Product &nbsp; <button onclick="getPageInfo('Product')">Open</button></li>
      <li>_ Brand &nbsp; <button onclick="getPageInfo('Brand')">Open</button></li>
      <li>_ Booth &nbsp; <button onclick="getPageInfo('Focaccia-Hands')">Open</button></li>
      <li>_ Software &nbsp; <button onclick="getPageInfo('Software')">Open</button></li>
      <li>_ Sales &nbsp; <button onclick="getPageInfo('Sales')">Open</button></li>
    </ul></p>
    `
}
   if data == "Kitchen-Page" {
    data = `</p><ul>
        <li>_ Cafe Needs &nbsp; <button onclick="getPageInfo('Cafe-Needs')">Open</button></li>
        <li>_ Timer &nbsp; <button onclick="getPageInfo('Timer')">Open</button></li>
        <li>_ Amount Conversion &nbsp; <button onclick="getPageInfo('Amount-Conversion')">Open</button></li>
        <li>_ Recipe Stuff &nbsp; <button onclick="getPageInfo('Recipe-Stuff')">Open</button></li>
        <li>_ Camera &nbsp; <button onlick="getPageInfo('Camera')">Open</button></li>
        <li>_ Notes &nbsp; <button onclick="getPageInfo('Notes')">Open</button></li>
        <li>_ Session Log &nbsp; <button onclick="getPageInfo('Session-Log')">Open</button></li>
    </ul></p>`
}

  if data == "Cafe-Needs" {
  data = "test"
}

  if data == "Timer" {
  date := time.Now() //.UTC().MarshalText()
  data = date.Format(time.RFC3339) + " !# - &nbsp;" + date.Format(time.Kitchen)
}

  if data == "Amount-Conversion" {
  data = "test2"
}

  if data == "Recipe-Stuff" {
  data = "test3"
}

    fmt.Fprintf(w, "Hello!")
    fmt.Fprintf(w, data)
}

// pagePath := 'test'

func getNavList(x string) string {
  data := x
  if data == "navList_Courses" {
    data = `<div><p><ul>
        <li>_ CSS Layout &nbsp; <button onclick="getPage('/courses/CSS-Layout')">Open</button></li>
        <li>_ Culniary Manager&nbsp; <button onclick="getPage('/courses/Culniary-Manager')">Open</button></li>
        <li>_ Food Handler&nbsp; <button onclick="getPage('/courses/Food-Handler')">Open</button></li>
        <li>_ System Administrator&nbsp; <button onclick="getPage('/courses/System-Administrator')">Open</button></li>
    </ul></p></div>`
  }
  if data == "navList_Concepts" {
    data = `<div><p><ul>
      <li>_ Loco Moco Bus&nbsp; <button onclick="getPage('/concepts/Loco-Moco-Bus')">Open</button></li>
      <li>_ Kitchen Page&nbsp; <button onclick="getPage('/concepts/Kitchen-Page')">Open</button></li>
      <li>_ Honolulu Lemonade&nbsp; <button onclick="getPage('/concepts/Honolulu-Lemonade')">Open</button></li>
      <li>_ Focaccia Hands&nbsp; <button onclick="getPage('/concepts/Focaccia-Hands')">Open</button></li>
    </ul></p></div>`
  }
  if data == "Focaccia-Hands" {
    data = `
    <p><ul>
      <li>_ People</li>
      <li>_ Product</li>
      <li>_ Brand</li>
      <li>_ Booth</li>
      <li>_ Software</li>
      <li>_ Sales</li>
    </ul></p>
    `
  }
   
  return data
} // - 

func getPageInfo(x string) string {
	data := "hi" + x
    data = "hello" + data
	
	if x == "/concepts/Focaccia-Hands" {
      data = "Focaccia-Hands Concept Page"
    } 
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

.bottomSheet { outline: 1px solid #aaa; width: 95%; margin: 0 auto; }
    
</style>
</head>
    `, pagePath)
    
    topBar := `<body>
    <div class="top_bar">
      <p><b>#! - Selfie Lunch Concepts</b></p>
      <hr />
    </div><!-- - . top_bar - -->`
    
    fmt.Fprintf(w, topBar)
    var sectionBreak = "<br /><hr /><br />"
    
    fmt.Fprintf(w, "World! %s", pagePath)
    
    pageInfo := getPageInfo(pagePath)
    fmt.Fprintf(w, pageInfo)
    
    fmt.Fprintf(w, sectionBreak)
    
    getData_Button := `<div id="demo">
  <h2>Let AJAX change this text</h2>
  <button type="button" onclick="loadDoc()">Change Content</button>
</div>
<div id="info"><code>- info</code></div>`

    fmt.Fprintf(w, getData_Button)
    fmt.Fprintf(w, sectionBreak)
    
    var pathLayers = strings.Split(pagePath, "/")
    fmt.Fprintf(w, pathLayers[1])
    fmt.Fprintf(w, sectionBreak)
    fmt.Fprintf(w, pathLayers[2])
   
    fmt.Fprintf(w, sectionBreak)
    
    getPage := `function getPage(x) {
     window.location.href = x }`
    fmt.Fprintf(w, "<script>%s</script>", getPage)
    
    pageID :=  fmt.Sprintf("var pageID = '%s'", pathLayers[2])
    fmt.Fprintf(w, "<script>%s</script>", pageID)
    
    
    navList := getNavList("navList_Courses") + getNavList("navList_Concepts")
    
    navData := "<button onclick='getPageData(pageID)'>Page Open</button>"
    
    navSpace := navData + navList
    
    bottomSheet := fmt.Sprintf("<div class='bottomSheet'> -%s</div>", navSpace)
    
    fmt.Fprintf(w, bottomSheet)
    
    getPageData_Request := `function getPageData(x) {
  const xhttp = new XMLHttpRequest();
  xhttp.onload = function() {
    document.getElementById("demo").innerHTML = this.responseText;
    }
  xhttp.open("GET", "/getData/" + x, true);
  xhttp.send();
}`

    fmt.Fprintf(w, "<script>%s</script>", getPageData_Request)
    
    getPageInfo_Request := `function getPageInfo(x) {
  const xhttp = new XMLHttpRequest();
  xhttp.onload = function() {
    document.getElementById("info").innerHTML = this.responseText;
    }
  xhttp.open("GET", "/getData/" + x, true);
  xhttp.send();
}`

    fmt.Fprintf(w, "<script>%s</script>", getPageInfo_Request)
    
    getData_Request := `function loadDoc() {
  const xhttp = new XMLHttpRequest();
  xhttp.onload = function() {
    document.getElementById("demo").innerHTML = this.responseText;
    }
  xhttp.open("GET", "/hello", true);
  xhttp.send();
}`

    fmt.Fprintf(w, "<script>%s</script>", getData_Request)
    
    fmt.Fprintf(w, sectionBreak)
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
    http.HandleFunc("/getData/Focaccia-Hands", getPageData)
    http.HandleFunc("/getData/Loco-Moco-Bus", getPageData)
    http.HandleFunc("/getData/Kitchen-Page", getPageData)
    http.HandleFunc("/getData/Honolulu-Lemonade", getPageData)
    http.HandleFunc("/getData/CSS-Layout", getPageData)
    http.HandleFunc("/getData/Culniary-Manager", getPageData)
    http.HandleFunc("/getData/Food-Handler", getPageData)
    http.HandleFunc("/getData/System-Administrator", getPageData)
    
   http.HandleFunc("/getData/People", getPageData)
 http.HandleFunc("/getData/Product", getPageData)
    http.HandleFunc("/getData/Brand", getPageData)
    http.HandleFunc("/getData/Booth", getPageData)
    http.HandleFunc("/getData/Software", getPageData)
    http.HandleFunc("/getData/Sales", getPageData)
    
       http.HandleFunc("/getData/Cafe-Needs", getPageData)
 http.HandleFunc("/getData/Timer", getPageData)
    http.HandleFunc("/getData/Amount-Conversion", getPageData)
    http.HandleFunc("/getData/Recipe-Stuff", getPageData)
    http.HandleFunc("/getData/Camera", getPageData)
    http.HandleFunc("/getData/Notes", getPageData)
    http.HandleFunc("/getData/Session-Log", getPageData)
   
          
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