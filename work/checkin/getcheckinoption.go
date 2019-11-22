package checkin

import (
	"fmt"
	"io"
	"net/http"

	"github.com/golib2020/wechat/internal/response"
)

type OptionResponse struct {
	Info []struct {
		Userid string `json:"userid"`
		Group  struct {
			Grouptype   int `json:"grouptype"`
			Groupid     int `json:"groupid"`
			Checkindate []struct {
				Workdays    []int `json:"workdays"`
				Checkintime []struct {
					WorkSec          int `json:"work_sec"`
					OffWorkSec       int `json:"off_work_sec"`
					RemindWorkSec    int `json:"remind_work_sec"`
					RemindOffWorkSec int `json:"remind_off_work_sec"`
				} `json:"checkintime"`
				FlexTime       int  `json:"flex_time"`
				NoneedOffwork  bool `json:"noneed_offwork"`
				LimitAheadtime int  `json:"limit_aheadtime"`
			} `json:"checkindate"`
			SpeWorkdays []struct {
				Timestamp   int    `json:"timestamp"`
				Notes       string `json:"notes"`
				Checkintime []struct {
					WorkSec          int `json:"work_sec"`
					OffWorkSec       int `json:"off_work_sec"`
					RemindWorkSec    int `json:"remind_work_sec"`
					RemindOffWorkSec int `json:"remind_off_work_sec"`
				} `json:"checkintime"`
			} `json:"spe_workdays"`
			SpeOffdays []struct {
				Timestamp   int           `json:"timestamp"`
				Notes       string        `json:"notes"`
				Checkintime []interface{} `json:"checkintime"`
			} `json:"spe_offdays"`
			SyncHolidays bool   `json:"sync_holidays"`
			Groupname    string `json:"groupname"`
			NeedPhoto    bool   `json:"need_photo"`
			WifimacInfos []struct {
				Wifiname string `json:"wifiname"`
				Wifimac  string `json:"wifimac"`
			} `json:"wifimac_infos"`
			NoteCanUseLocalPic     bool `json:"note_can_use_local_pic"`
			AllowCheckinOffworkday bool `json:"allow_checkin_offworkday"`
			AllowApplyOffworkday   bool `json:"allow_apply_offworkday"`
			LocInfos               []struct {
				Lat       int    `json:"lat"`
				Lng       int    `json:"lng"`
				LocTitle  string `json:"loc_title"`
				LocDetail string `json:"loc_detail"`
				Distance  int    `json:"distance"`
			} `json:"loc_infos"`
		} `json:"group"`
	} `json:"info"`
}

//GetCheckinOption 获取打卡规则
func getCheckinOption(token string, body io.Reader) (*OptionResponse, error) {
	apiUrl := fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckinoption?access_token=%s`, token)
	resp, err := http.Post(apiUrl, "application/json", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res := new(OptionResponse)
	if err := response.ReadBody(resp.Body, res); err != nil {
		return nil, err
	}
	return res, nil
}
