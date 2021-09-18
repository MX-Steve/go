package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type User struct {
	UserName string
	Password string
	Age      int
}

func info(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./info.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	// add self method
	kuaFunc := func(arg string) (string, error) {
		return arg + "beautiful", nil
	}
	// template.New("info").Funcs(template.FuncMap{"kua": kuaFunc}).Parse(string(data))
	t, err := template.New("info").Funcs(template.FuncMap{"kua": kuaFunc}).Parse(string(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	// data := "<li>《我的世界》</li>"
	// t.Execute(w, data)

	// user := User{
	// 	"张三",
	// 	"1233",
	// 	18,
	// }
	// t.Execute(w, user)

	userMap := map[int]User{
		1: {"张三", "li12", 18},
		2: {"莉丝", "sl11", 28},
		3: {"王武", "1223", 8},
		4: {"照留", "xxxx", 38},
	}
	t.Execute(w, userMap)

	// num := rand.Intn(10)
	// dataStr := string(data)
	// if num > 5 {
	// 	dataStr = strings.Replace(dataStr, "{ooxx}", "《我的世界》", 1)
	// } else {
	// 	dataStr = strings.Replace(dataStr, "{ooxx}", "《Go开发》", 1)
	// }
	// w.Write([]byte(dataStr))

}

func main() {
	http.HandleFunc("/info", info)
	panic(http.ListenAndServe("0.0.0.0:8081", nil))
}
