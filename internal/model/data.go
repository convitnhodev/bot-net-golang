package model

type Data struct {
	Ten           string   `json:"ten"`
	UID           string   `json:"uid"`
	Email         string   `json:"email"`
	Birthday      string   `json:"birthday"`
	Location      string   `json:"location"`
	Fa            string   `json:"fa"`
	Step          string   `json:"step"`
	TrinhDuyet    string   `json:"trinhDuyet"`
	Cookie        string   `json:"cookie"`
	Password      string   `json:"password"`
	Usergmail     string   `json:"usergmail"`
	Passgmail     string   `json:"passgmail"`
	Useroutlook   string   `json:"useroutlook"`
	Passoutlook   string   `json:"passoutlook"`
	HanhDong      string   `json:"hanhDong"`
	Token         string   `json:"token"`
	Businesses    []string `json:"businesses"`
	Fanpage       []string `json:"fanpage"`
	TaiKhoanAds   []string `json:"taiKhoanAds"`
	CookieOutlook string   `json:"cookieOutlook"`
	CookieGg      string   `json:"cookieGg"`
}
