package js_binding

import (
	"github.com/robertkrimen/otto"
	"common/logger"
	"io/ioutil"
)

type JsScript struct {
	vm           *otto.Otto
}

var DefaultScript JsScript

func NewScript() *JsScript {

	vm := otto.New()
	DefaultScript = JsScript{vm : vm}

	return &DefaultScript
}

func (self *JsScript) ExecuteScriptFile(file string) {
	script, err := ioutil.ReadFile(file)
	if err != nil {
		logger.Fatal("script: ReadFile %s, Err : %s", file, err.Error());
	}
	if _,err = self.vm.Run(string(script)); err != nil {
		logger.Fatal("script: ExecuteScriptFile %s, Err : %s", file, err.Error());
	}
}