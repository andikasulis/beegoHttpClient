package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	req := httplib.Post("https://ptsv2.com/t/ad9oc-1627574830/post")
	req.Body("{\n  \"ack\": \"00\",\n  \"message\": \"Succcess\",\n  \"data\": {\n    \"nik\": \"5413435434354\",\n    \"nama\": \"Jonson\"\n  }\n}")
	str, err := req.String()
	if err != nil {
		//t.Fatal(err)
	}
	fmt.Println(str)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
