package controllers

import (
	"encoding/json"
	"io/ioutil"
    "fmt"
	"strconv"
"strings"
	model "kbcms/models"
	"github.com/ant0ine/go-json-rest/rest"
)
type ExtendController struct{}
func (d *ExtendController) DeleteAll(idStr string){
    idStr = strings.Replace(idStr,"|",".",-1)
    idStr = strings.Replace(idStr,";",".",-1)
    idStr = strings.Replace(idStr,"；",".",-1)
    idStr = strings.Replace(idStr,",",".",-1)
    idStr = strings.Replace(idStr,"，",".",-1)
	ids := strings.Split(idStr,".")
	//
	data := new(model.Extend)
	fail := make([]string,len(ids),len(ids))	
	for i:=0;i<len(ids);i++{
		if  ids[i]!=""{
			err := data.Del(ids[i])
			if err != nil {
				fail[i]	= ids[i]				
			}
		}
	}
	fmt.Println("----fail--%v----", fail )
	
}

func (d *ExtendController) Delete(w rest.ResponseWriter, req *rest.Request) {
	data := new(model.Extend)
	err := data.Del(req.PathParam("id"))
	if err != nil {
		w.WriteJson(0)
		return
	}
	// TODO: 将extend中dataId的数据删除
	w.WriteJson(1)
}

func (d *ExtendController) Add(w rest.ResponseWriter, req *rest.Request) {
	data := new(model.Extend)
	c, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(c, &data)
	err = data.Add()
	if err != nil {
		w.WriteJson(0)
	}	
	w.WriteJson(data)
}

func (d *ExtendController) Update(w rest.ResponseWriter, req *rest.Request) {
	data := new(model.Extend)
	c, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(c, &data)
	err = data.Edit()
	if err != nil {
		w.WriteJson(0)
	}
	w.WriteJson(data)
}

func (d *ExtendController) Get(w rest.ResponseWriter, req *rest.Request) {
	data := new(model.Extend)
	data.ById(req.PathParam("id"))
	w.WriteJson(data)
}
func (d *ExtendController) All(w rest.ResponseWriter, req *rest.Request) {
	s := req.PathParam("s")
	c := req.PathParam("c")
	startId, _ := strconv.Atoi(s)
	count, _ := strconv.Atoi(c)
	data := new(model.Extend)
	datas, err := data.All(startId, count)
	if err != nil {
		fmt.Println(err)
		w.WriteJson(err)
		return
	}
	
	w.WriteJson(datas)
}

 