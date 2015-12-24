package models

import (	
	"strconv"
 
)  
// 某条数据在某目录(分类)下的配置
type Extend struct{
	Id				int64
	DataId			int64
	TreeId			int64
	IsTop			int64
    Recommend		int64
    Visible			int64
    Amount			int64
	Sort 			string
	Content			string
} 


func (cate *Extend)Add( ) error {
   _, err := Eng.Insert(cate)
    if err !=nil{
		return err
	}
	return nil
}

func  (data *Extend)Del(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
    data.Id = cid
	_, err = Eng.Delete(data)
	return err
}

func  (cate *Extend)ById(id string)(*Extend,error){
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


func  (e *Extend)All(startId ,count int) ( []Extend, error) {
	
	 sql := "select * from extend where id >= ? limit ?"
	resultsSlice, err := Eng.Query(sql,startId,count)	
		if err != nil {
			//fmt.Println(err)
		}
	
    
	extends := make([]Extend, count)

	for i,Slice := range resultsSlice {
		data := new(Extend)
		for f,bb := range Slice {
			bbb := string(bb)
			
			 switch f{
				case "id":				
				data.Id ,err = strconv.ParseInt(bbb,10,64)
				case "data_id":
				data.DataId ,err = strconv.ParseInt(bbb,10,64)
				case "tree_id":
				data.TreeId  ,err = strconv.ParseInt(bbb,10,64)
				case "is_top":
				data.IsTop ,err = strconv.ParseInt(bbb,10,64)
				case "recommend":
				data.Recommend  ,err = strconv.ParseInt(bbb,10,64)
				case "visible":
				data.Visible ,err = strconv.ParseInt(bbb,10,64)
				case "amount":
				data.Amount ,err = strconv.ParseInt(bbb,10,64)
				case "content":
				data.Content = bbb
				case "sort":
				data.Sort  = bbb
				 			
			}
				if err !=nil{
					break
				}
		}
		 
		 if(data.Id >= 1){			
			extends[i] = *data
		}
		 
		
	}
     
		return extends ,nil 
}

func (data *Extend)Edit()error{
	if  data.Id==0{
		return nil
	} 
	_, err := Eng.Id(data.Id).Update(&data)
	if err == nil {
		return err
	}
  return nil
}