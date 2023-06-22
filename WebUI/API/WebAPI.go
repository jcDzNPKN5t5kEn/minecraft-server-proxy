package API

import (
	"fmt"
	"minecraft-proxy/Config"
	"net/http"
	"strings"
	"strconv"
)

func HandleWebAPI(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/whitelist-update" && r.Method == "POST" {
		whitelist := r.FormValue("whitelist")
		allowedNames := strings.Split(whitelist, "\r\n")
		for _, name := range allowedNames {
			if !InWhiteList(name) {
				AddAllowedName(name)
			}
		}
		http.Redirect(w, r, "/Whitelist", 302)
	}
	if r.URL.Path == "/api/logout" {
		http.SetCookie(w, &http.Cookie{
			Name:  "_u",
			Value: "",
		})
		http.Redirect(w, r, "/login", 302)
	}
	if r.URL.Path == "/api/update-overwrite-host" {
		Config.CurConfig.OverwriteHost = r.FormValue("update-overwrite-host")
		fmt.Println(Config.CurConfig.OverwriteHost)
		http.Redirect(w, r, "/Panel", 302)
	}
	if r.URL.Path == "/api/update-upstream" {
		Config.CurConfig.Remote = r.FormValue("update-upstream")
		fmt.Println(Config.CurConfig.Remote)
		http.Redirect(w, r, "/Panel", 302)
	}
	if r.URL.Path == "/api/update-overwrite-port" {
		port, err := strconv.Atoi(r.FormValue("update-overwrite-port"))
		if(err != nil){
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		Config.CurConfig.OverwritePort = port
		fmt.Println(Config.CurConfig.OverwritePort)
		http.Redirect(w, r, "/Panel", 302)
	}
}
