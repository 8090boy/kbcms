package controllers

import (
	"fmt"
	//"net/http"
	"encoding/json"
	"io/ioutil"
    "github.com/ant0ine/go-json-rest/rest"
)

type CommonController struct{
	DataController *DataController
	ExtendController *ExtendController
	TreeController *TreeController
}
var result map[string]string
func (cc *CommonController)MethodDispatcher(rep rest.ResponseWriter,req *rest.Request){
   bp, err := ioutil.ReadAll(req.Body)
   err = json.Unmarshal(bp, &result)
	if err !=nil{ 
	  fmt.Println(err)
 	  return
	} 
    method  := result["method"]
    params  := result["params"]
	switch method {
		case "delDatasForId":
		cc.DataController.DeleteAll(params)		
		case "delExtendsForId":
		cc.ExtendController.DeleteAll(params)
		case "delTreesForId":
		cc.TreeController.DeleteAll(params)
	}
	
}

 
 