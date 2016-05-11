package js_binding

import (
	"github.com/robertkrimen/otto"
	"common/logger"
	"io/ioutil"
	"path/filepath"
)

var (
	globalScriptPath = "./script/js"

	// Globally registered modules
	globalModules map[string]ModuleLoader = make(map[string]ModuleLoader)

	// Globally registered paths (paths to search for modules)
	globalPaths []string
)

type JsScript struct {
	*otto.Otto

	// Modules that registered for current vm
	modules map[string]ModuleLoader

	// Location to search for modules
	paths []string

	// Onece a module is required by vm, the exported value is cached for further
	// use.
	moduleCache map[string]otto.Value
}

var DefaultScript JsScript

// Create a JsScript vm instance.
func NewScript() *JsScript {

	DefaultScript = JsScript{otto.New(), make(map[string]ModuleLoader), nil, make(map[string]otto.Value)}

	AddPath(globalScriptPath)

	// Provide the "require" method in the JsScript scope.
	jsRequire := func(call otto.FunctionCall) otto.Value {
		jsModuleName := call.Argument(0).String()

		moduleValue, err := DefaultScript.Require(jsModuleName, globalScriptPath)
		if err != nil {
			jsException(&DefaultScript, "Error", "JsScript: " + err.Error())
		}

		return moduleValue
	}

	DefaultScript.Set("require", jsRequire)

	return &DefaultScript
}

func (self *JsScript) ExecuteScriptFile(file string) {
	script, err := ioutil.ReadFile(file)
	if err != nil {
		logger.Fatal("script: ReadFile %s, Err : %s", file, err.Error());
	}
	if _,err = self.Otto.Run(string(script)); err != nil {
		logger.Fatal("script: ExecuteScriptFile %s, Err : %s", file, err.Error());
	}
}

// Run a module or file
func (self *JsScript) Run(name string) (otto.Value, error) {
	if ok, _ := isFile(name); ok {
		name, _ = filepath.Abs(name)
	}

	return self.Require(name, ".")
}

// Require a module with cache
func (self *JsScript) Require(id, pwd string) (otto.Value, error) {
	if cache, ok := self.moduleCache[id]; ok {
		return cache, nil
	}

	loader, ok := self.modules[id]
	if !ok {
		loader, ok = globalModules[id]
	}

	if loader != nil {
		value, err := loader(self)
		if err != nil {
			return otto.UndefinedValue(), err
		}

		self.moduleCache[id] = value
		return value, nil
	}

	filename, err := FindFileModule(id, pwd, append(self.paths, globalPaths...))
	if err != nil {
		return otto.UndefinedValue(), err
	}

	// resove id
	id = filename

	if cache, ok := self.moduleCache[id]; ok {
		return cache, nil
	}

	v, err := CreateLoaderFromFile(id)(self)

	if err != nil {
		return otto.UndefinedValue(), err
	}

	// cache
	self.moduleCache[id] = v

	return v, nil
}

// Register a new module to current vm.
func (self *JsScript) AddModule(id string, loader ModuleLoader) {
	self.modules[id] = loader
}


// Add paths to search for modules.
func (self *JsScript) AddPath(paths ...string) {
	self.paths = append(self.paths, paths...)
}

// Register a global module
func AddModule(id string, m ModuleLoader) {
	globalModules[id] = m
}

// Register global path.
func AddPath(paths ...string) {
	globalPaths = append(globalPaths, paths...)
}