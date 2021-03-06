//文件 bundle.go   （文件中无此句）
// auto-generated
// Code generated by '$ fyne bundle'. DO NOT EDIT.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"gui-license/theme"
	"net/http"
	"strings"
	"time"
)

type  requestData struct {
	Data []*requestParam `json:"data"`
	Sign string `json:"sign"`
}
type requestParam struct {
	SnCode string `json:"snCode"`
	SnId   string `json:"snId"`
	Time   int64  `json:"time"`
	Name   string `json:"name"`
}

//var requestUrl = "http://192.168.0.32:9999/api/license/setlicense"
var requestUrl = "http://www.v-secure.cn/api/license/setlicense"


func main() {
	t:= &theme.MyTheme{}
	t.SetFonts("../theme/Alibaba-PuHuiTi-Regular.ttf", "")
	a := app.New()
	a.Settings().SetTheme(t)
	w := a.NewWindow("手动修改路径工具")

	title := widget.NewLabel("snCode，多个使用，隔开")
	entry := widget.NewEntry()
	title2 := widget.NewLabel("单位缩写，按照snCode数量进行匹配，使用‘，’隔开")
	entry2 := widget.NewEntry()
	titleData := widget.NewLabel("到期时间，按照snCode数量进行匹配，使用‘，’隔开")
	entryData := widget.NewEntry()
	titleSnID := widget.NewLabel("snid，按照snCode数量进行匹配，使用‘，’隔开")
	entrySnID := widget.NewEntry()
	w.SetContent(container.NewVBox(
		title,
		entry,
		title2,
		entry2,
		titleData,
		entryData,
		titleSnID,
		entrySnID,
		widget.NewButton("yes", func() {
			re := UpdateOfficialWebsite(entry.Text, entry2.Text, entryData.Text, entrySnID.Text)
			dialog.ShowInformation("Title", re, w)
		}),
	))
	size := fyne.NewSize(400, 500)
	w.Resize(size)
	w.ShowAndRun()


}

func UpdateOfficialWebsite(snCode, name, exData, snId string) string {
	snCodes := strings.Split(snCode, ",")
	names := strings.Split(name, ",")
	exDatas := strings.Split(exData, ",")
	snIds := strings.Split(snId, ",")
	if len(snCodes) != len(names) || len(snCodes) != len(exDatas) || len(snCodes) != len(snIds) || len(snCodes) != len(names) {
		return "Sn does not match the abbreviation"
	}
	parm := make([]*requestParam, 0, len(snCodes))
	for i := 0; i < len(snCodes); i++ {
		reqp := new(requestParam)
		reqp.SnCode = snCodes[i]
		reqp.Name = names[i]
		reqp.Time = TimeToUnix2("2006-01-02", exDatas[i])
		reqp.SnId = snIds[i]
		parm = append(parm, reqp)
	}
	reqData := new(requestData)
	reqData.Data = parm
	reqData.Sign = "ZBmx3MNOlbknRgHoC9DTb6uhJS0L0wG4lozo7+2cwTsRylVjTsatrDfpt68L+OuGEaifwlSrVKuhc+qnzFs9NbGuW9OfV/LXuTeNekCjvzmgq4n3gZUKbNmf72PtOXQ3NEs4wZdMAE2K+8azhD2XM2oznOUrfFekypt8FvyF7J7ahkvOkt9/PMvsYic9hee4u0stRP6k5yyx2u+sojRniAQEzF4LhOUNykHgCb9buyveMsrRBTzSHzLyLbU4nI1/hxsf+9ZnuVnEXx3P8wB/luqQuSeOidWTtSZ4ABGpSX8mA36ezQiq7zg/U7EQmR6oDYAwp2vkE5+5ob54GWWVow=="
	data, err := json.Marshal(reqData)
	if err != nil {
		return "404 request failed "
	}
	//return string(data)
	request, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(data))
	if err != nil {
		return "500 " + err.Error()
	}
	request.Header.Set("Content-Type", "application/json")
	var client = http.DefaultClient
	_, err = client.Do(request)
	if err != nil {
		return fmt.Sprintf("404 request failed %v", err.Error())
	}
	return "ok"
}

//时间字符串转时间戳
func TimeToUnix2(format string, timeValue string) int64 {
	loc, _ := time.LoadLocation("Local")                       //重要：获取时区
	theTime, _ := time.ParseInLocation(format, timeValue, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()
	return sr
}
