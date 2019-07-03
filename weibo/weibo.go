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
	b.Notify(weibo.Id)
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
	Attach(bFans FansInterface)
	Detach(bFans FansInterface)
	Notify(id int)
	GetWeiBo(wbid int) *PostContent
}

func (b *Blogger)Attach(bFans FansInterface){
	b.Fans=append(b.Fans,bFans)
}

func (b *Blogger)Detach(bFans FansInterface){
	for i:=0;i<len(b.Fans);i++{
		if b.Fans[i]==bFans{
			b.Fans=append(b.Fans[:i],b.Fans[i+1:]...)
		}
	}
}

func (b *Blogger)Notify(wbid int) {
	for _,fan:=range b.Fans {
		fan.Update(b,wbid)
	}
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
	Update(bloggerI BloggerInterface,wbid int)
	Action(bloggerI BloggerInterface,wbid int)
}

func (b *Blogger)GetWeiBo(wbid int) *PostContent{
	for _,blog:=range b.WeiBos {
		if blog.Id==wbid {
			return blog
		}
	}
	return nil
}

//真爱粉
type FriedFans struct {
	Fans
}

func (f *FriedFans)Update(bloggerI BloggerInterface,wbid int){
	fmt.Printf("%s,你所关注的博主发布了一个微博",f.Name)
	f.Action(bloggerI,wbid)
}
func (f *FriedFans)Action(bloggerI BloggerInterface, wbid int){
	weibo:=bloggerI.GetWeiBo(wbid)
	fmt.Println(weibo)
}

//黑粉
type BadFans struct {
	Fans
}

func (f *BadFans)Update(bloggerI BloggerInterface,wbid int){
	fmt.Printf("%s,你所关注的博主发布了一个微博",f.Name)
	f.Action(bloggerI,wbid)
}
func (f *BadFans)Action(bloggerI BloggerInterface,wbid int){
	weibo:=bloggerI.GetWeiBo(wbid)
	fmt.Println(weibo)
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
	// blg.Detach(friedFans)

	blg.PostWeiBo("今天天气很好",1)
}