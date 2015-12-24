package models
import (
    "testing"
      "time" 
    "strconv" 
    "fmt" 
    
)
 
type TestType struct{
	
 	Id			int64
  Title				string
 Published     	   	time.Time
}

func Test_A(t *testing.T) {
	RegisterDB()
	sql := "select * from testtype limit 0, 10"
	ttSlice , err := Eng.Query(sql)
	fmt.Println( "-----------" )
	
 	datas := make([]*TestType, 5)
	
	for i,Slice := range ttSlice {	
		data := new(TestType)		
		for f,bb := range Slice {
			tmpStr := string(bb)
			
			 switch f{
				case "id":
				data.Id , err = strconv.ParseInt(tmpStr,10,0)
				case "title":
				data.Title = tmpStr
				case "published":
				data.Published,err = time.Parse( TimeFormat ,tmpStr)				
						
			}
			
				if err !=nil{
					break
				} 
		}
fmt.Println( *data )
		if(data.Id > 0){
			 datas[i] = data
		 }
	
}
 fmt.Println( datas[0].Published.Before( datas[1].Published ) ) //d0 在 d1 之前 true
 fmt.Println( datas[1].Published.Before( datas[0].Published ) ) //d1 在 d0 之前 false
fmt.Println( datas[1].Published.After( datas[0].Published ) ) //d1 在 d0 之后 true
fmt.Println( "-----------" )
 }
 