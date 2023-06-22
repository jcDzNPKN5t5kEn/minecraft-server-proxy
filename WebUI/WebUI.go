package WebUI

import (
	// "io"
	// "net"
	"io/ioutil"
	"minecraft-proxy/Config"
	"minecraft-proxy/Utils"
	"minecraft-proxy/WebUI/API"
	"net/http"
	"strconv"
	"strings"
	// "os"
)

func Init(port string) {
	http.HandleFunc("/", HandleFunc)
	err := http.ListenAndServe(port, nil)
	if err != nil {

	}
}

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	// ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	// io.WriteString(w, ip)
	cookie, err := r.Cookie("_u")
	if err != nil {
		notLoggedInHandle(w, r)
		return
	}
	auth, _ := AuthWebUI(cookie.Value,"")
	if auth {
		if(strings.HasPrefix(r.URL.Path,"/api/")){
			API.HandleWebAPI(w,r)
		}else{
			CopyFileToResponse(w, "./WebUI" + r.URL.Path + ".htm", r, HandleWeb)
		}

	} else {
		notLoggedInHandle(w, r)
	}

}
func CopyFileToResponse(w http.ResponseWriter, fileName string, r *http.Request, handle func(string, *http.Request) string) {
        // 1. 打开文件并读取其内容到内存中
        content, err := ioutil.ReadFile(fileName)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

	replacedContent := string(content)
        if(handle != nil){
		 replacedContent = handle(replacedContent, r)
	}


        // 3. 将替换后的内容写到 http.ResponseWriter
        _, err = w.Write([]byte(replacedContent))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
}

func AddBanner (ori string) string {
	banner,_ := Utils.OpenFile("./WebUI/Banner.htm")
	return strings.Replace(ori,"<!--  banner  -->",string(banner),1)
}

func HandleWeb (ori string, r *http.Request) string {
	returnText := ori
	if r.URL.Path == "/Whitelist" {
		returnText = strings.Replace(ori,"<!--ALLOWED_NAME-->",API.GetAllowedName(),1)
	}
	if r.URL.Path == "/Panel" {
		returnText = strings.Replace(ori,"update-upstream-replace",Config.CurConfig.Remote,1)
		returnText = strings.Replace(returnText,"update-overwrite-host-replace",Config.CurConfig.OverwriteHost,1)
		returnText = strings.Replace(returnText,"update-overwrite-port-replace",strconv.Itoa(Config.CurConfig.OverwritePort),1)
	}
	returnText = AddBanner(returnText)
	return returnText
}
