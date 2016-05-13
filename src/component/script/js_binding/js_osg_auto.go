package js_binding

import (
	"github.com/robertkrimen/otto"
	"common/logger"
	"io/ioutil"
	"path/filepath"
	"time"
)

var (
	globalPackagePath = "script/js/package.json"
	globalScriptPath = "./script/js"

	// Globally registered modules
	globalModules map[string]ModuleLoader = make(map[string]ModuleLoader)

	// Globally registered paths (paths to search for modules)
	globalPaths []string
)

type _timer struct {
	timer    *time.Timer
	duration time.Duration
	interval bool
	call     otto.FunctionCall
}

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

	// Run will execute the given JavaScript, continuing to run until all timers have finished executing (if any).
	// The VM has the following functions available:
	//
	//      <timer> = setTimeout(<function>, <delay>, [<arguments...>])
	//      <timer> = setInterval(<function>, <delay>, [<arguments...>])
	//      clearTimeout(<timer>)
	//      clearInterval(<timer>)
	//
	go func() {
		registry := map[*_timer]*_timer{}
		ready := make(chan *_timer)
		newTimer := func(call otto.FunctionCall, interval bool) (*_timer, otto.Value) {
			delay, _ := call.Argument(1).ToInteger()
			if 0 >= delay {
				delay = 1
			}

			timer := &_timer{
				duration: time.Duration(delay) * time.Millisecond,
				call:     call,
				interval: interval,
			}
			registry[timer] = timer

			timer.timer = time.AfterFunc(timer.duration, func() {
				ready <- timer
			})

			value, err := call.Otto.ToValue(timer)
			if err != nil {
				panic(err)
			}

			return timer, value
		}

		setTimeout := func(call otto.FunctionCall) otto.Value {
			_, value := newTimer(call, false)
			return value
		}
		DefaultScript.Set("setTimeout", setTimeout)

		setInterval := func(call otto.FunctionCall) otto.Value {
			_, value := newTimer(call, true)
			return value
		}
		DefaultScript.Set("setInterval", setInterval)

		clearTimeout := func(call otto.FunctionCall) otto.Value {
			timer, _ := call.Argument(0).Export()
			if timer, ok := timer.(*_timer); ok {
				timer.timer.Stop()
				delete(registry, timer)
			}
			return otto.UndefinedValue()
		}
		DefaultScript.Set("clearTimeout", clearTimeout)
		DefaultScript.Set("clearInterval", clearTimeout)

		for {
			select {
			case timer := <-ready:
				var arguments []interface{}
				if len(timer.call.ArgumentList) > 2 {
					tmp := timer.call.ArgumentList[2:]
					arguments = make([]interface{}, 2 + len(tmp))
					for i, value := range tmp {
						arguments[i + 2] = value
					}
				} else {
					arguments = make([]interface{}, 1)
				}
				arguments[0] = timer.call.ArgumentList[0]
				_, err := DefaultScript.Call(`Function.call.call`, nil, arguments...)
				if err != nil {
					for _, timer := range registry {
						timer.timer.Stop()
						delete(registry, timer)
						logger.Error("script timer: %s", err.Error())
					}
				}
				if timer.interval {
					timer.timer.Reset(timer.duration)
				} else {
					delete(registry, timer)
				}
			default:
			// Escape valve!
			// If this isn't here, we deadlock...
			}
			if len(registry) == 0 {
				break
			}
			logger.Debug("timer : %d", uint32(time.Now().Unix()))
		}
	}()

	// Provide the "require" method in the JsScript scope.
	jsRequire := func(call otto.FunctionCall) otto.Value {
		jsModuleName := call.Argument(0).String()

		moduleValue, err := DefaultScript.Require(jsModuleName, globalScriptPath)
		if err != nil {
			jsException(&DefaultScript, "Error", "JsScript: " + err.Error())
		}

		return moduleValue
	}
	DefaultScript.Object(`osg = {}`)
	DefaultScript.Set("require", jsRequire)

	return &DefaultScript
}

func (self *JsScript) SetGlobalPackagePath(path string) {
	globalPackagePath = path;
}

func (self *JsScript) ExecuteScriptFile(file string) {
	script, err := ioutil.ReadFile(file)
	if err != nil {
		logger.Error("script: ReadFile %s, Err : %s", file, err.Error())
	}
	if _,err = self.Otto.Run(string(script)); err != nil {
		logger.Error("script: ExecuteScriptFile %s, Err : %s", file, err.Error())
	}
}

func (self *JsScript) UpdateScriptFile() {
	self.moduleCache = make(map[string]otto.Value)
	entryPoint, err := parsePackageEntryPoint(globalPackagePath)
	if err != nil {
		logger.Error("script: ReadFile %s, Err : %s", globalPackagePath, err.Error())
	}
	self.ExecuteScriptFile("script/js/" + entryPoint)
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