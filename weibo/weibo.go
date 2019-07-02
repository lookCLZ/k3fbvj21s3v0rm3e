package main 

import "time"
import "fmt"


type Base struct {
	Id int
	Name string
}
//博主类
type Blogger struct {
	Base
	WeiBos []*PostContent
	Comments map[int][]*PostContent
}
//发布微博的方法实现
func (b *Blogger)PostWeiBo(content string,wbType int) {
	weibo:=new(PostContent)
	weibo.Id=1
	weibo.Content=content
	weibo.Type=wbType
	weibo.CommentTime=time.Now()
	weibo.PostMan=b.Name

	b.WeiBos=append(b.WeiBos,weibo)
	for _,v:=range b.WeiBos{
		fmt.Println(*v)
	}
}

//博主接口
type BloggerInterface interface{
	Attach()
	Detach()
	Notify()
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

type FansInterface interface{
	Update()
	Action()
}

//真爱粉
type FriedFans struct {
	Fans
}

//黑粉
type BadFans struct {
	Fans
}

func NewBlogger(name string) *Blogger {
	blg:=new(Blogger)
	blg.Name=name
	blg.Comments=make(map[int][]*PostContent)
	blg.WeiBos=make([]*PostContent,0)

	return blg
}

func main(){
	blg:=NewBlogger("张三")
	blg.PostWeiBo("今天天气很好",1)
}