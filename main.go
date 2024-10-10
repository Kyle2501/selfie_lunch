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
  data = `<div><p><b>Cafe Needs Page</b><ul>
     <li>_ Menu &nbsp; <button onclick="getPageInfo('Menu')">Open</button></li>
     <li>_ Location &nbsp; <button onclick="getPageInfo('Location')">Open</button></li>
      <li>_ Equiptment &nbsp; <button onclick="getPageInfo('Equiptment')">Open</button></li>
      <li>_ Staff &nbsp; <button onclick="getPageInfo('Staff')">Open</button></li>
      <li>_ Services &nbsp; <button onclick="getPageInfo('Services')">Open</button></li>
     <li>_ Music &nbsp; <button onclick="getPageInfo('Music')">Open</button></li>
     <li>_ Setting &nbsp; <button onclick="getPageInfo('Setting')">Open</button></li>
       <li>_ Sources &nbsp; <button onclick="getPageInfo('Sources')">Open</button></li>
      <li>_ Marketing &nbsp; <button onclick="getPageInfo('Marketing')">Open</button></li>
      <li>_ Operations &nbsp; <button onclick="getPageInfo('Operations')">Open</button></li>
     <li>_ Software &nbsp; <button onclick="getPageInfo('Software')">Open</button></li>
      <li>_ Events &nbsp; <button onclick="getPageInfo('Events')">Open</button></li>
       <li>_ Seating &nbsp; <button onclick="getPageInfo('Seating')">Open</button></li>
     <li>_ Hours &nbsp; <button onclick="getPageInfo('Hours')">Open</button></li>
  `
}

  if data == "Timer" {
  date := time.Now() //.UTC().MarshalText()
  data = date.Format(time.RFC3339) + " !# - &nbsp;" + date.Format(time.Kitchen)
}

  if data == "Amount-Conversion" {
  data = " Amount Conversion Page"
}

  if data == "Recipe-Stuff" {
  data = "Recipe Stuff Page"
}

  if data == "Camera" {
  data = "Camera Page"
}

  if data == "Notes" {
  data = "Notes Page"
}

  if data == "Session-Log" {
    data = "Session Log Page"
}

if data == "CSS-Layout" {
  data = `
    <div>
    <p>Intro and Welcome &nbsp; <button onclick="getPageInfo('Intro-and-Welcome')">Open</button></p>
    <p>Setup Environment &nbsp; <button onclick="getPageInfo('Setup-Environment')">Open</button></p>
    <p>Text and Color &nbsp; <button onclick="getPageInfo('Text-and-Color')">Open</button></p>
    <p>Box Model Concept &nbsp; <button onclick="getPageInfo('Box-Model-Concepts')">Open</button></p>
    <p>Element Class Names &nbsp; <button onclick="getPageInfo('Element-Class-Names')">Open</button></p>
    <p>Design Principles &nbsp; <button onclick="getPageInfo('Design-Principles')">Open</button></p>
    </div>
  `
}

  if data == "Intro-and-Welcome" {
  data = `<p><b>Intro and Welcome Page</b></p>
  <p>--</p>
  <div>
    <p><b>_ Weekly Schedule</b><ul>
      <li>_ Day 01 &nbsp; // on Mon ~</li>
      <li>_ Day 02 &nbsp; // on Wed ~</li>
      <li>_ Day 03 &nbsp; // on Fri ~</li>
    </ul></p>
    <p><b>_ Extra Help Days</b><ul>
      <li>_ Day 01 &nbsp; // on Tues ~</li>
      <li>_ Day 02 &nbsp; // on Thurs ~</li>
    </ul></p>
    <p><b>_ Weekend Classes</b><ul>
      <li>_ Day 01 &nbsp; // on Sat ~</li>
      <li>_ Day 02 &nbsp; // on Sun ~</li>
    </ul</p>
    <p><b>_ Meeting Times</b><ul>
      <li>_ Early Class &nbsp; // at 10:am ~</li>
      <li>_ Later Class &nbsp; // at 2:pm ~</li>
      <li>_ Evening Class &nbsp; // at 7:pm ~</li>
    </ul></p>
  </div>
  `
}

  if data == "Setup-Environment" {
  data =  `<p><b>Setup Environment Page</b></p>
    <p><ul>
      <li>_ Hardware Setup  &nbsp; <button onclick="getPageInfo('Hardware-Setup')">Open</button></li>
      <li>_ Software Setup  &nbsp; <button onclick="getPageInfo('Software-Setup')">Open</button></li>
      <li>_ Online Accounts Setup  &nbsp; <button onclick="getPageInfo('Online-Accounts')">Open</button></li>
  </ul</p>
  `
}

  if data == "Hardware-Setup" {
  data = "Hardware Setup Page"
}
  if data == "Software-Setup" {
  data = "Software Setup Page"
}
  if data == "Online-Accounts" {
  data = "Online Accounts Setup Page"
}


  if data == "Text-and-Color" {
  data = `<b>Text and Color Page</b>
  <p>_ Vocab<ul>
    <li>_ text-align</li>
    <li>_ font-size</li>
    <li>_ font-family</li>
    <li>_ font-wight</li>
    <li>_ color</li>
  
  </ul></p>
  
  `
}

  if data == "Box-Model-Concept" {
  data = `<p><b>Box Model Concept Page</b></p>
  <p>_ Vocab<ul>
    <li>_ Padding</li>
    <li>_ Margin</li>
    <li>_ Border</li>
    <li>_ Outline</li>
    <li>_ Width</li>
    <li>_ Height</li>
    <li>_ Left</li>
    <li>_ Top</li>
    <li>_ Right</li>
    <li>_ Bottom</li>
    <li>_ Background-Color</li>
    <li>_ Position</li>
    <li>_ z-index</li>
  </ul></p>
  `
}

  if data == "Element-Class-Names" {
  data = "Element Class Names Page"
}

if data == "Design-Principles" {
  data = `<div><p><b>Design Principles Page</b><ul>
    <li>_ Contrast  &nbsp; <button onclick="getPageInfo('Contrast')">Open</button></li>
    <li>_ Movement  &nbsp; <button onclick="getPageInfo('Repetition')">Open</button></li>
    <li>_ Repetition  &nbsp; <button onclick="getPageInfo('People')">Open</button></li>
    <li>_ Empasis  &nbsp; <button onclick="getPageInfo('Empasis')">Open</button></li>
    <li>_ Balance  &nbsp; <button onclick="getPageInfo('Balance')">Open</button></li>
    <li>_ Hierarchy  &nbsp; <button onclick="getPageInfo('Hierarchy')">Open</button></li>
    <li>_ Proportion  &nbsp; <button onclick="getPageInfo('Proportion')">Open</button></li>
    <li>_ Unity  &nbsp; <button onclick="getPageInfo('Unity')">Open</button></li>
    <li>_ White Space  &nbsp; <button onclick="getPageInfo('White-Space')">Open</button></li>
    <li>_ Rhythm  &nbsp; <button onclick="getPageInfo('Rhythm')">Open</button></li>
    <li>_ Variety &nbsp; <button  onclick="getPageInfo('Variety')">Open</button></li>
    <li>_ Color  &nbsp; <button onclick="getPageInfo('Color')">Open</button></li>
    <li>_ Pattern &nbsp; <button  onclick="getPageInfo('Pattern')">Open</button></li>
    <li>_ Scale  &nbsp; <button onclick="getPageInfo('Scale')">Open</button></li>
    <li>_ Visual Hierarchy  &nbsp; <button onclick="getPageInfo('Visual-Hierarchy')">Open</button></li>
    <li>_ Aethetics  &nbsp; <button onclick="getPageInfo('Aethetics')">Open</button></li>
    <li>_ Context &nbsp; <button  onclick="getPageInfo('Context')">Open</button></li>
    <li>_ Gestalt Principles  &nbsp; <button onclick="getPageInfo('Gestalt')">Open</button></li>
    <li>_ Space  &nbsp; <button onclick="getPageInfo('Space')">Open</button></li>
    <li>_ Symmetry &nbsp; <button  onclick="getPageInfo('Symmetry')">Open</button></li>
    <li>_ Alignment &nbsp; <button  onclick="getPageInfo('Alignment')">Open</button></li>
    <li>_ Prevent Errors if Possible  &nbsp; <button onclick="getPageInfo('Prevent-Errors-if-Possible')">Open</button></li>
  
  </ul></p></div>`
}


    if data == "Contrast" { data = "Contrast Page" }
    if data == "Movement" { data = "Movement Page" }
    if data == "Repetition" { data = "Repetition Page" }
    if data == "Empasis" { data = "Empasis Page" }
    if data == "Balance" { data = "Balance Page" }
    if data == "Hierarchy" { data = "Hierarchy Page" }
    if data == "Proportion" { data = "Proportion Page" }
    if data == "Unity" { data = "Unity Page" }
    if data == "White-Space" { data = "White Space Page" }
    if data == "Rhythm" { data = "Rhythm Page" }
    if data == "Variety" { data = "Variety Page" }
    if data == "Color" { data = "Color Page" }
    if data == "Pattern" { data = "Pattern Page" }
     if data == "Scale" { data = "Scale Page" }
    if data == "Visual-Hierarchy" { data = "Visual Hierarchy Page" }
    if data == "Aethetics" { data = "Aethetics Page" }
    if data == "Context" { data = "Context Page" }
    if data == "Gestalt-Principles" { data = "Gestalt Page" }
    if data == "Space" { data = "Space Page" }
    if data == "Symmetry" { data = "Symmetry Page" }
    if data == "Alignment" { data = "Alignment Page" }
    if data == "Prevent-Errors-if-Possible" { data = "Pervent Errors if Possible Page" }

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
   
   http.HandleFunc("/getData/Intro-and-Welcome", getPageData)
   http.HandleFunc("/getData/Setup-Environment", getPageData)
   http.HandleFunc("/getData/Text-and-Color", getPageData)
   http.HandleFunc("/getData/Box-Model-Concept", getPageData)
   http.HandleFunc("/getData/Element-Class-Names", getPageData)
   http.HandleFunc("/getData/Design-Principles", getPageData)
   
   http.HandleFunc("/getData/Hardware-Setup", getPageData)
   http.HandleFunc("/getData/Software-Setup", getPageData)
   http.HandleFunc("/getData/Online-Accounts", getPageData)
   
   // ,  ° . +
       http.HandleFunc("/getData/Contrast", getPageData)
    http.HandleFunc("/getData/Movement", getPageData)
    http.HandleFunc("/getData/Repetition", getPageData)
    http.HandleFunc("/getData/Empasis", getPageData)
    http.HandleFunc("/getData/Balance", getPageData)
    http.HandleFunc("/getData/Hierarchy", getPageData)
    http.HandleFunc("/getData/Proportion", getPageData)
    http.HandleFunc("/getData/Unity", getPageData)
    http.HandleFunc("/getData/White-Space", getPageData)
    http.HandleFunc("/getData/Rhythm", getPageData)
    http.HandleFunc("/getData/Variety", getPageData)
    http.HandleFunc("/getData/Color", getPageData)
    http.HandleFunc("/getData/Pattern", getPageData)
    http.HandleFunc("/getData/Scale", getPageData)
    http.HandleFunc("/getData/Visual-Hierarchy", getPageData)
    http.HandleFunc("/getData/Aethetics", getPageData)
    http.HandleFunc("/getData/Context", getPageData)
    http.HandleFunc("/getData/Gestalt-Principles", getPageData)
    http.HandleFunc("/getData/Space", getPageData)
    http.HandleFunc("/getData/Symmetry", getPageData)
    http.HandleFunc("/getData/Alignment", getPageData)
    http.HandleFunc("/getData/Prevent-Errors-if-Possible", getPageData)
          
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