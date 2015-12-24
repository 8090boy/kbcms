package main
import (
    "testing"   
    "fmt" 
)
 
type User struct{
	id int
	name string
	
}
  
func Test_A1(t *testing.T) {
var u = User{12,"zhangsan"}
var v1,p1 interface{} = u,&u
p1.(*User).name = "jack"
fmt.Printf("%v\n",v1)
fmt.Printf("%v\n",p1)
}