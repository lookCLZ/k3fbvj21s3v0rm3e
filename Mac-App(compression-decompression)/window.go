package main

import (
    "github.com/andlabs/ui"
	"fmt"
	"os"
	"io"
	"archive/zip"
)

type MacWindow interface {
	ShowWindow()
}

type ComWindow struct {
	MacWindow
	uiWin *ui.Window
}

type LabWindow struct {
	MacWindow
}

var labelElement = ui.NewLabel("...欢迎使用...")

func (c *ComWindow)ShowWindow() {
	ComWindowObj:=new(ComWindow)
	ComWindowObj.uiWin=ui.NewWindow("文件压缩",600,230,false)
	
	ComWindowObj.uiWin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})

	ComWindowObj.uiWin.SetMargined(true)

	hbox := ui.NewHorizontalBox()

	ComWindowObj.uiWin.SetChild(hbox)

	group1 := ui.NewGroup("输入区")
	group2 := ui.NewGroup("选择文件按钮区")
	group3:=ui.NewGroup("程序按钮区")
	group4:=ui.NewGroup("消息区")

	vbox1 := ui.NewVerticalBox()
	vbox2 := ui.NewVerticalBox()
	vbox3 := ui.NewVerticalBox()
	vbox4:=labelElement

	group1.SetChild(vbox1)
	group2.SetChild(vbox2)
	group3.SetChild(vbox3)
	group4.SetChild(vbox4)

	hbox.Append(group1, false)
	hbox.Append(group2, false)
	hbox.Append(group3,false)
	hbox.Append(group4,false)

	// 输入框
	decompressInput:=ui.NewEntry()
	saveDecoInput:=ui.NewEntry()
	compressInput:=ui.NewEntry()
	saveCoInput:=ui.NewEntry()

	// 按钮
	decompressSelect:=ui.NewButton("选择解压文件")
	saveDecoSelect:=ui.NewButton("选择解压文件保存位置")
	compressSelect:=ui.NewButton("选择压缩文件")
	saveCoSelect:=ui.NewButton("选择压缩文件保存位置")

	startDeco:=ui.NewButton("开始解压")
	startCo:=ui.NewButton("开始压缩")

	// 点击 选择解压文件
	decompressSelect.OnClicked(func(*ui.Button) {
		filename:=OpenFileManager(ComWindowObj.uiWin)
		decompressInput.SetText(filename)
	})

	// 点击 选择解压文件保存位置
	saveDecoSelect.OnClicked(func(*ui.Button) {
		filename:=SaveFileManager(ComWindowObj.uiWin)
		saveDecoInput.SetText(filename)
	})

	// 点击 选择压缩文件
	compressSelect.OnClicked(func(*ui.Button) {
		filename:=OpenFileManager(ComWindowObj.uiWin)
		compressInput.SetText(filename)
	})

	// 点击 选择压缩文件保存位置
	saveCoSelect.OnClicked(func(*ui.Button) {
		filename:=SaveFileManager(ComWindowObj.uiWin)
		saveCoInput.SetText(filename)
	})

	// 点击  开始解压
	startDeco.OnClicked(func(*ui.Button) {
		res:=c.StartToUnZip(decompressInput.Text(),saveDecoInput.Text())
		if(res) {
			labelElement.SetText("解压成功")
		}else{
			labelElement.SetText("解压失败")
		}
	})

	startCo.OnClicked(func(*ui.Button){
		res:=c.StartToZip(compressInput.Text(),saveCoInput.Text())
		if(res) {
			labelElement.SetText("压缩成功")
		}else{
			labelElement.SetText("压缩失败")
		}
	})

	vbox1.Append(decompressInput,false)
	vbox1.Append(saveDecoInput,false)
	vbox1.Append(compressInput,false)
	vbox1.Append(saveCoInput,false)

	vbox2.Append(decompressSelect,false)
	vbox2.Append(saveDecoSelect,false)
	vbox2.Append(compressSelect,false)
	vbox2.Append(saveCoSelect,false)

	vbox3.Append(startDeco,false)
	vbox3.Append(startCo,false)

	ComWindowObj.uiWin.Show()
}

// 打开文件选择对话框
func OpenFileManager(w *ui.Window) string {
	filename:=ui.OpenFile(w)
	return filename
}
func SaveFileManager(w *ui.Window) string {
	filename:=ui.SaveFile(w)
	fmt.Println(filename)
	return filename
}

// 解压文件
func (c *ComWindow)StartToUnZip(file string,saveFile string)(resBool bool) {
	labelElement.SetText("正在解压")
	resBool=true
	reader,err:=zip.OpenReader(file)
	if err!=nil{
		resBool=false
		fmt.Println("001")
		fmt.Println(err)
		return
	}
	defer reader.Close()

	for _,file:=range reader.File {
		fmt.Println("file:   ",file.Name)
		rc,err:=file.Open()
		if err!=nil{
			resBool=false
			fmt.Println(err)
			return
		}
		defer rc.Close()

		newName:=saveFile+"/"+file.Name

		// 如果是文件夹
		if file.FileInfo().IsDir() {
			os.MkdirAll(newName,os.ModePerm)
			if err!=nil{
				resBool=false
				fmt.Println(err)
				return
			}
		} else {
			// 否则是文件
			f,err:=os.Create(newName)
			if err!=nil{
				resBool=false
				fmt.Println(err)
				return
			}
			_,err=io.Copy(f,rc)
			if err!=nil{
				resBool=false
				fmt.Println(err)
				return
			}
		}
	}
	return
}

// 压缩文件
func (c *ComWindow)StartToZip(file string,saveFile string)(resBool bool) {
	resBool=true
	// 创建一个文件
	d,err:=os.Create(saveFile)

	if err!=nil{
		fmt.Println(err)
	}
	defer d.Close()

	// 打开要压缩文件
	fileP,err:=os.Open(file)
	if err!=nil{
		fmt.Println(err)
	}
	defer fileP.Close()

	// 获取文件信息
	info,err:=fileP.Stat()
	if err!=nil{
		fmt.Println(err)
	}

	// 返回压缩所需的文件头信息
	header,err:=zip.FileInfoHeader(info)
	if err!=nil{
		fmt.Println(err)
	}

	w:=zip.NewWriter(d)
	defer w.Close()

	// 根据文件头信息创建一个zip操作指针
	writer,err:=w.CreateHeader(header)
	if err!=nil{
		fmt.Println(err)
	}

	// 把文件内容复制到zip操作指针所指向的内存中
	io.Copy(writer,fileP)

	return
}
func (c *LabWindow)ShowWindow() {

}

// 创建界面类对象
func Show(window_Type string) {
	var Win MacWindow
	switch window_Type {
	case "main_window":
		Win=&ComWindow{}
	case "lab_window":
		Win=&LabWindow{}
	default:
		fmt.Println("参数传递错误")
	}
	Win.ShowWindow()
}