package gothulearn

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type NoticeService service

type NoticeList struct {
	Result string       `json:"result"`
	Msg    string       `json:"msg"`
	Object NoticeObject `json:"object"`
}

type NoticeObject struct {
	ITotalRecords        string   `json:"iTotalRecords"`
	ITotalDisplayRecords string   `json:"iTotalDisplayRecords"`
	IDisplayStart        string   `json:"iDisplayStart"`
	IDisplayLength       string   `json:"iDisplayLength"`
	SEcho                string   `json:"sEcho"`
	SSearch              string   `json:"sSearch"`
	AaData               []Notice `json:"aaData"`
}

type Notice struct {
	Id       string `json:"id"`       // ID=公告ID+学号
	Ggid     string `json:"ggid"`     // 公告ID
	Bt       string `json:"bt"`       // 标题
	Xh       string `json:"xh"`       // 学号
	Wlkcid   string `json:"wlkcid"`   //
	Fbr      string `json:"fbr"`      // 发布人
	Fbrxm    string `json:"fbrxm"`    // 发布人姓名
	Fbsj     string `json:"fbsj"`     // 发布时间
	FbsjStr  string `json:"fbsjStr"`  // 发布时间 Str
	Ydsj     string `json:"ydsj"`     // 阅读时间
	Sfqd     string `json:"sfqd"`     // 是否确定？// FLAG
	Ggnr     string `json:"ggnr"`     // 公告内容
	GgnrStr  string `json:"ggnrStr"`  // 公告内容 Str
	GgnrMini string `json:"ggnrMini"` // 公告内容 Mini
	Fjmc     string `json:"fjmc"`     // 附件名称
	Sfyd     string `json:"sfyd"`     // 是否阅读
}

func (n *NoticeService) GetNoticeList(courseID string) (*NoticeList, error) {
	// Create request
	var payload url.Values = url.Values{}
	payload.Add("aoData", fmt.Sprintf("[{\"name\":\"sEcho\",\"value\":1},{\"name\":\"iColumns\",\"value\":3},{\"name\":\"sColumns\",\"value\":\",,\"},{\"name\":\"iDisplayStart\",\"value\":0},{\"name\":\"iDisplayLength\",\"value\":\"-1\"},{\"name\":\"mDataProp_0\",\"value\":\"bt\"},{\"name\":\"bSortable_0\",\"value\":true},{\"name\":\"mDataProp_1\",\"value\":\"fbr\"},{\"name\":\"bSortable_1\",\"value\":true},{\"name\":\"mDataProp_2\",\"value\":\"fbsj\"},{\"name\":\"bSortable_2\",\"value\":true},{\"name\":\"iSortingCols\",\"value\":0},{\"name\":\"wlkcid\",\"value\":\"%s\"}]", courseID))
	resp, err := n.client.Request(context.Background(), http.MethodPost, addCSRFTokenToUrl(learnNotificationList(courseID), n.client.csrf), strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, err
	}

	// Handle response
	bytesResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response NoticeList
	json.Unmarshal(bytesResponse, &response)

	return &response, nil
}
