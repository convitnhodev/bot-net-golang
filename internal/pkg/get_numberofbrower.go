package pkg

import (
	"os/exec"
	"regexp"
	"strings"
)

func GetBrowser() (int, []string) {
	out, err := exec.Command("cmd", "/c", "reg", "query", `HKEY_LOCAL_MACHINE\SOFTWARE\Clients\StartMenuInternet`).Output()
	if err != nil {
		panic(err)
	}

	// Regex to extract the browser name from the registry key
	r := regexp.MustCompile(`StartMenuInternet\\(.+)$`)

	// Convert the byte slice to a string and split it by newlines
	browsers := strings.Split(string(out), "\n")

	// Iterate through each browser name and extract it using the regex
	result := make([]string, 0)
	for _, b := range browsers {
		if m := r.FindStringSubmatch(b); m != nil {
			result = append(result, m[1])
		}
	}
	return len(result), result
}
