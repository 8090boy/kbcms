package main

import (
    "github.com/ant0ine/go-json-rest/rest"  
    "log"
    "fmt"
    "net/http"
	model "kbcms/models"
    controller "kbcms/controllers"
)
func init() {
	model.RegisterDB()	
}
func main(){
   
    api := rest.NewApi()    
    api.Use(rest.DefaultDevStack...)
    api.Use(&rest.CorsMiddleware{
        RejectNonCorsRequests: false,
        OriginValidator: func(origin string, request *rest.Request) bool {
            stat := false 
              switch origin {
				case "http://localhost":
				stat = true
				case "http://www.jahurd.com":
				stat = true
				case "http://localhost.jahurd.com":
				stat = true
				case "http://www.kingbloc.com":
				stat = true
				case "http://api.kingbloc.com":
				stat = true
			  }
            return stat
        },
        AllowedMethods: []string{"GET", "POST"},
        AllowedHeaders: []string{"Accept", "Content-Type", "X-Custom-Header", "Origin"},
        AccessControlAllowCredentials: true,
        AccessControlMaxAge:           3600,
    })
    api.Use(&rest.JsonpMiddleware{
        CallbackNameKey: "ok",
    })

    d := controller.DataController{}
    e := controller.ExtendController{}
    t := controller.TreeController{}
 	cc := controller.CommonController{}
	 

    router, err := rest.MakeRouter(
	 
       rest.Post("/cc", cc.MethodDispatcher ),
        //
        rest.Get("/d/get/:id", d.Get),
        rest.Get("/d/del/:id", d.Delete),
        rest.Get("/d/all/:s/:c", d.All),
        rest.Post("/d/add", d.Add),
        rest.Post("/d/edit", d.Update),
		//
        rest.Get("/e/get/:id", e.Get),
        rest.Get("/e/del/:id", e.Delete),
        rest.Get("/e/all/:s/:c", e.All),
        rest.Post("/e/add", e.Add),
        rest.Post("/e/edit", e.Update),
		//
        rest.Get("/t/get/:id", t.Get),
        rest.Get("/t/del/:id", t.Delete),
        rest.Get("/t/all/:s/:c", t.All),
        rest.Post("/t/add", t.Add),
        rest.Post("/t/edit", t.Update),
		 
    )
    if err != nil {
 		log.Fatal(err)
		return
       
    }
    api.SetApp(router)
 
    http.HandleFunc("/", index )
    http.Handle("/v1/", http.StripPrefix("/v1", api.MakeHandler()))
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("."))))
	fmt.Println(http.ListenAndServe(":100", nil))
}

func  index(w http.ResponseWriter, r *http.Request) {

//	isLogin := true
//	if isLogin {
		http.Redirect(w, r, "/static/default/index.html", 200)
//	} else {
//		 http.Redirect(w, r, "/static/default/login.html", 200)
//	}
}
