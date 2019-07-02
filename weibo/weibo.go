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
	Fans []FansInterface
}
//发布微博的方法实现
func (b *Blogger)PostWeiBo(content string,wbType int) {
	weibo:=new(PostContent)
	weibo.Id=b.GetId()
	weibo.Content=content
	weibo.Type=wbType
	weibo.CommentTime=time.Now()
	weibo.PostMan=b.Name

	b.WeiBos=append(b.WeiBos,weibo)
	for _,v:=range b.WeiBos{
		fmt.Println(*v)
	}
}

//获取微博编号
func (b *Blogger)GetId() int {
	if len(b.WeiBos)==0{
		return 0
	}
	return b.WeiBos[len(b.WeiBos)-1].Id+1
}

//博主接口
type BloggerInterface interface{
	Attach()
	Detach()
	Notify()
}

func (b *Blogger)Attach(bFans FansInterface){
	b.Fans=append(b.Fans,bFans)
	fmt.Println(b.Fans)
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

func (f *FriedFans)Update(){

}
func (f *FriedFans)Action(){

}

//黑粉
type BadFans struct {
	Fans
}

func (f *BadFans)Update(){

}
func (f *BadFans)Action(){

}

func NewBlogger(name string) *Blogger {
	blg:=new(Blogger)
	blg.Name=name
	// blg.Comments=make(map[int][]*PostContent)
	// blg.WeiBos=make([]*PostContent,0)

	return blg
}



func main(){
	blg:=NewBlogger("张三")

	friedFans:=new(FriedFans)
	friedFans.Id=1
	friedFans.Name="李四"

	blg.Attach(friedFans)
	blg.PostWeiBo("今天天气很好",1)
}