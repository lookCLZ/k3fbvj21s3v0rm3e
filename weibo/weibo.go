package main 

import "time"

type Base struct {
	Id int
	Name string
}
//博主类
type Blogger struct {
	Base
	WeiBos []*PostContent
	Comments []*PostContent
}

type PostContent struct {
	Id int
	Content string
	CommentTime time.Time 
	Type int
	PostMan string
	To string 
}

//粉丝
type Fans struct{
	Base
}

//真爱粉
type FriedFans struct {
	Fans
}

//黑粉
type BadFans struct {
	Fans
}

func main(){

}