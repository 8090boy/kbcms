package models

import (
"strconv" 
"fmt" 
)
// 目录(分类)
type Tree struct{
	Id				int64
	ParentId		int64
	Name			string
	RefInfo			string
}
func (data *Tree)Add( ) error { 
	 _, err := Eng.Insert(data)
    if err !=nil{
		return err		
	}	 
	return nil
}

func  (data *Tree)Del(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
    data.Id = cid
	_, err = Eng.Delete(data)
	return err
}

func  (cate *Tree)ById(id string)(*Tree,error){
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil,err
	}
	_, err = Eng.Id(cid).Get(cate)
	if err != nil {
			return nil,err
		}
	return cate,nil
}

func (data *Tree)Edit()error{
	if  data.Id==0{
		return nil
	} 
	_, err := Eng.Id(data.Id).Update(&data)
	if err == nil {
		return err
	}
  return nil
}

func  (t *Tree)All(startId ,count int) ([]Tree, error) {
	 
	 sql := "select * from tree where id >= ? limit ?"
	resultsSlice, err := Eng.Query(sql,startId,count)	
		if err != nil {
			//fmt.Println(err)
		}
 
 	datas := make([]Tree, count)

	for i,Slice := range resultsSlice {
		data := Tree{}
		for f,bb := range Slice {
			bbb := string(bb) 
			 switch f{
				case "id":				
				data.Id  , err = strconv.ParseInt(bbb,10,64)
				case "parent_id":
				data.ParentId , err = strconv.ParseInt(bbb,10,64)
				case "name":
				data.Name = bbb
				case "ref_info":
				data.RefInfo = bbb			 
			}
				if err !=nil{
					break
				}
		}
		if(data.Id > 0){
			 datas[i] = data
		 } 
		
	}
      fmt.Println(datas)
		return datas ,nil 
	
	
}

 
