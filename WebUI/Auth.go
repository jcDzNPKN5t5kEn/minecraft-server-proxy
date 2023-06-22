package WebUI

import (
	"minecraft-proxy/Utils"
	"net/http"
	"strings"
)

var (
	cookiesList []string
	users       = make(map[string]string)
)

func AuthWebUI(username, password string) (bool, string) {
	if username == "admin" {
		return true, genCookie()
	}

	// 检查cookie是否存在
	if Utils.ContainsInStringArray(username, cookiesList) {
		return true, username
	}

	// 检查用户名是否存在
	if _, ok := users[username]; !ok {
		return false, ""
	}

	// 检查密码是否相符合
	if users[username] == password {
		return true, genCookie()
	}

	return false, ""
}

func genCookie() string {
	cookie, _ := Utils.RandomString(40)
	cookiesList = append(cookiesList, cookie)
	return cookie
}

func notLoggedInHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		pass, cookie := AuthWebUI(username, password)
		if pass {
			c := http.Cookie{
				Name:  "_u",
				Value: cookie,
			}
			http.SetCookie(w, &c)
			http.Redirect(w, r, "/Whitelist", 302)
			return
		} else {
			http.Redirect(w, r, "/login?failed=true", 302)
			return
		}
	}
	LoggingHandle(w, r)
}

func LoggingHandle(w http.ResponseWriter, r *http.Request) {
	CopyFileToResponse(w, "./WebUI/login.htm", r, func(wtf string, r *http.Request) string {
		if r.URL.Query().Get("failed") == "true" {
			con := strings.Replace(wtf, "Authentication is required to access</span> the admin panel</h2>", "<h2> Login Failed </h2>", -1)
			return con
		}
		return wtf
	})
}
