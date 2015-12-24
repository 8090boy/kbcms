package models

import (	
	"strconv" 
) 

// 数据
type Data struct{
  Id				int64
  ExtendId			int64
  TreeId			int64 //第一次发布需要这个ID，也可以不关联
  Title				string
  Url				string
  Attachment		string
  ImgUrl			string
  Origin			string
  Summary			string      
  Note				string     
  Content			string		`orm:"size(20000)"`
  Published 		string  
  ReferenceWords	string 
} 

func (cate *Data)Add() error {
	_, err := Eng.Insert(cate)
    if err !=nil{
		return err		
	}	 
	return nil
}



func (data *Data)Del(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
    data.Id = cid
	_, err = Eng.Delete(data)
	return err
}

func (d *Data)All(startId ,count int) ([]*Data, error) {
 
    sql := "select * from data where id >= ? limit ?"
	resultsSlice, err := Eng.Query(sql,startId,count)	
		if err != nil {
			//fmt.Println(err)
		}
	
 	datas := make([]*Data, count)
	
	for i,Slice := range resultsSlice {	
		data := new(Data)		
		for f,bb := range Slice {
			tmpStr := string(bb)
			
			 switch f{
				case "id":
				data.Id , err = strconv.ParseInt(tmpStr,10,64)
				case "extend_id":
				data.ExtendId, err= strconv.ParseInt(tmpStr,10,64)
				case "tree_id":
				data.TreeId , err= strconv.ParseInt(tmpStr,10,64)
				case "title":
				data.Title = tmpStr
				case "url":
				data.Url = tmpStr
				case "attachment":
				data.Attachment = tmpStr
				case "img_url":
				data.ImgUrl = tmpStr
				case "content":
				data.Content = tmpStr
				case "reference_words":
				data.ReferenceWords = tmpStr
				case "origin":
				data.Origin = tmpStr
				case "summary":
				data.Summary = tmpStr
				case "note":
				data.Note = tmpStr
				case "published":
				//data.Published,err = time.Parse(TimeFormat,tmpStr)				
				data.Published = tmpStr			
			}
			
				if err !=nil{
					break
				} 
		}
		if(data.Id > 0){
			 datas[i] = data
		 } 
	 
	}
		return datas ,nil 
}

func (cate *Data)ById(id string)(*Data,error){
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



func (data *Data)Edit()error{
	if  data.Id==0{
		return nil
	} 
	_, err := Eng.Id(data.Id).Update(&data)
	if err == nil {
		return err
	}
  return nil
}