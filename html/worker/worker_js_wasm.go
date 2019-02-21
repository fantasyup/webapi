// Code generated by webidl-bind. DO NOT EDIT.

package worker

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/fetch"
	"github.com/gowebapi/webapi/html/canvas"
	"github.com/gowebapi/webapi/html/channel"
	"github.com/gowebapi/webapi/html/htmlevent"
	"github.com/gowebapi/webapi/javascript"
	"github.com/gowebapi/webapi/patch"
	"github.com/gowebapi/webapi/webidl"
)

// using following types:
// canvas.ImageBitmapOptions
// channel.MessagePort
// channel.PostMessageOptions
// domcore.EventHandler
// domcore.EventTarget
// fetch.RequestCredentials
// fetch.RequestInit
// htmlevent.FrameRequestCallback
// htmlevent.OnErrorEventHandler
// javascript.FrozenArray
// javascript.Object
// javascript.Promise
// patch.ByteString
// webidl.VoidFunction

// ReleasableApiResource is used to release underlaying
// allocated resources.
type ReleasableApiResource interface {
	Release()
}

type releasableApiResourceList []ReleasableApiResource

func (a releasableApiResourceList) Release() {
	for _, v := range a {
		v.Release()
	}
}

// workaround for compiler error
func unused(value interface{}) {
	// TODO remove this method
}

type Union struct {
	Value js.Value
}

func (u *Union) JSValue() js.Value {
	return u.Value
}

func UnionFromJS(value js.Value) *Union {
	return &Union{Value: value}
}

// enum: WorkerType
type WorkerType int

const (
	ClassicWorkerType WorkerType = iota
	ModuleWorkerType
)

var workerTypeToWasmTable = []string{
	"classic", "module",
}

var workerTypeFromWasmTable = map[string]WorkerType{
	"classic": ClassicWorkerType, "module": ModuleWorkerType,
}

// JSValue is converting this enum into a java object
func (this *WorkerType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this WorkerType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(workerTypeToWasmTable) {
		return workerTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// WorkerTypeFromJS is converting a javascript value into
// a WorkerType enum value.
func WorkerTypeFromJS(value js.Value) WorkerType {
	key := value.String()
	conv, ok := workerTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// dictionary: WorkerOptions
type WorkerOptions struct {
	Type        WorkerType
	Credentials fetch.RequestCredentials
	Name        string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *WorkerOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Type.JSValue()
	out.Set("type", value0)
	value1 := _this.Credentials.JSValue()
	out.Set("credentials", value1)
	value2 := _this.Name
	out.Set("name", value2)
	return out
}

// WorkerOptionsFromJS is allocating a new
// WorkerOptions object and copy all values from
// input javascript object
func WorkerOptionsFromJS(value js.Wrapper) *WorkerOptions {
	input := value.JSValue()
	var out WorkerOptions
	var (
		value0 WorkerType               // javascript: WorkerType {type Type _type}
		value1 fetch.RequestCredentials // javascript: RequestCredentials {credentials Credentials credentials}
		value2 string                   // javascript: DOMString {name Name name}
	)
	value0 = WorkerTypeFromJS(input.Get("type"))
	out.Type = value0
	value1 = fetch.RequestCredentialsFromJS(input.Get("credentials"))
	out.Credentials = value1
	value2 = (input.Get("name")).String()
	out.Name = value2
	return &out
}

// interface: WorkerGlobalScope
type WorkerGlobalScope struct {
	domcore.EventTarget
}

// WorkerGlobalScopeFromJS is casting a js.Wrapper into WorkerGlobalScope.
func WorkerGlobalScopeFromJS(value js.Wrapper) *WorkerGlobalScope {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &WorkerGlobalScope{}
	ret.Value_JS = input
	return ret
}

// Self returning attribute 'self' with
// type WorkerGlobalScope (idl: WorkerGlobalScope).
func (_this *WorkerGlobalScope) Self() *WorkerGlobalScope {
	var ret *WorkerGlobalScope
	value := _this.Value_JS.Get("self")
	ret = WorkerGlobalScopeFromJS(value)
	return ret
}

// Location returning attribute 'location' with
// type WorkerLocation (idl: WorkerLocation).
func (_this *WorkerGlobalScope) Location() *WorkerLocation {
	var ret *WorkerLocation
	value := _this.Value_JS.Get("location")
	ret = WorkerLocationFromJS(value)
	return ret
}

// Navigator returning attribute 'navigator' with
// type WorkerNavigator (idl: WorkerNavigator).
func (_this *WorkerGlobalScope) Navigator() *WorkerNavigator {
	var ret *WorkerNavigator
	value := _this.Value_JS.Get("navigator")
	ret = WorkerNavigatorFromJS(value)
	return ret
}

// Onerror returning attribute 'onerror' with
// type htmlevent.OnErrorEventHandler (idl: OnErrorEventHandlerNonNull).
func (_this *WorkerGlobalScope) Onerror() htmlevent.OnErrorEventHandlerFunc {
	var ret htmlevent.OnErrorEventHandlerFunc
	value := _this.Value_JS.Get("onerror")
	if value.Type() != js.TypeNull {
		ret = htmlevent.OnErrorEventHandlerFromJS(value)
	}
	return ret
}

// SetOnerror setting attribute 'onerror' with
// type htmlevent.OnErrorEventHandler (idl: OnErrorEventHandlerNonNull).
func (_this *WorkerGlobalScope) SetOnerror(value *htmlevent.OnErrorEventHandler) {
	var __callback3 js.Value
	if value != nil {
		__callback3 = (*value).Value
	} else {
		__callback3 = js.Null()
	}
	input := __callback3
	_this.Value_JS.Set("onerror", input)
}

// Onlanguagechange returning attribute 'onlanguagechange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WorkerGlobalScope) Onlanguagechange() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onlanguagechange")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnlanguagechange setting attribute 'onlanguagechange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WorkerGlobalScope) SetOnlanguagechange(value *domcore.EventHandler) {
	var __callback4 js.Value
	if value != nil {
		__callback4 = (*value).Value
	} else {
		__callback4 = js.Null()
	}
	input := __callback4
	_this.Value_JS.Set("onlanguagechange", input)
}

// Onoffline returning attribute 'onoffline' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WorkerGlobalScope) Onoffline() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onoffline")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnoffline setting attribute 'onoffline' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WorkerGlobalScope) SetOnoffline(value *domcore.EventHandler) {
	var __callback5 js.Value
	if value != nil {
		__callback5 = (*value).Value
	} else {
		__callback5 = js.Null()
	}
	input := __callback5
	_this.Value_JS.Set("onoffline", input)
}

// Ononline returning attribute 'ononline' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WorkerGlobalScope) Ononline() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("ononline")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnonline setting attribute 'ononline' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WorkerGlobalScope) SetOnonline(value *domcore.EventHandler) {
	var __callback6 js.Value
	if value != nil {
		__callback6 = (*value).Value
	} else {
		__callback6 = js.Null()
	}
	input := __callback6
	_this.Value_JS.Set("ononline", input)
}

// Onrejectionhandled returning attribute 'onrejectionhandled' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WorkerGlobalScope) Onrejectionhandled() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onrejectionhandled")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnrejectionhandled setting attribute 'onrejectionhandled' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WorkerGlobalScope) SetOnrejectionhandled(value *domcore.EventHandler) {
	var __callback7 js.Value
	if value != nil {
		__callback7 = (*value).Value
	} else {
		__callback7 = js.Null()
	}
	input := __callback7
	_this.Value_JS.Set("onrejectionhandled", input)
}

// Onunhandledrejection returning attribute 'onunhandledrejection' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WorkerGlobalScope) Onunhandledrejection() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onunhandledrejection")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnunhandledrejection setting attribute 'onunhandledrejection' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WorkerGlobalScope) SetOnunhandledrejection(value *domcore.EventHandler) {
	var __callback8 js.Value
	if value != nil {
		__callback8 = (*value).Value
	} else {
		__callback8 = js.Null()
	}
	input := __callback8
	_this.Value_JS.Set("onunhandledrejection", input)
}

// Origin returning attribute 'origin' with
// type string (idl: USVString).
func (_this *WorkerGlobalScope) Origin() string {
	var ret string
	value := _this.Value_JS.Get("origin")
	ret = (value).String()
	return ret
}

func (_this *WorkerGlobalScope) ImportScripts(urls ...string) {
	var (
		_args []interface{} = make([]interface{}, 0+len(urls))
		_end  int
	)
	for _, __in := range urls {
		__out := __in
		_args[_end] = __out
		_end++
	}
	_this.Value_JS.Call("importScripts", _args[0:_end]...)
	return
}

func (_this *WorkerGlobalScope) Btoa(data string) (_result string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := data
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("btoa", _args[0:_end]...)
	var (
		_converted string // javascript: DOMString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

func (_this *WorkerGlobalScope) Atob(data string) (_result *patch.ByteString) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := data
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("atob", _args[0:_end]...)
	var (
		_converted *patch.ByteString // javascript: ByteString _what_return_name
	)
	_converted = patch.ByteStringFromJS(_returned)
	_result = _converted
	return
}

func (_this *WorkerGlobalScope) SetTimeout(handler *Union, timeout *int, arguments ...interface{}) (_result int) {
	var (
		_args []interface{} = make([]interface{}, 2+len(arguments))
		_end  int
	)
	_p0 := handler.JSValue()
	_args[0] = _p0
	_end++
	if timeout != nil {
		_p1 := timeout
		_args[1] = _p1
		_end++
	}
	for _, __in := range arguments {
		__out := __in
		_args[_end] = __out
		_end++
	}
	_returned := _this.Value_JS.Call("setTimeout", _args[0:_end]...)
	var (
		_converted int // javascript: long _what_return_name
	)
	_converted = (_returned).Int()
	_result = _converted
	return
}

func (_this *WorkerGlobalScope) ClearTimeout(handle *int) {
	var (
		_args [1]interface{}
		_end  int
	)
	if handle != nil {
		_p0 := handle
		_args[0] = _p0
		_end++
	}
	_this.Value_JS.Call("clearTimeout", _args[0:_end]...)
	return
}

func (_this *WorkerGlobalScope) SetInterval(handler *Union, timeout *int, arguments ...interface{}) (_result int) {
	var (
		_args []interface{} = make([]interface{}, 2+len(arguments))
		_end  int
	)
	_p0 := handler.JSValue()
	_args[0] = _p0
	_end++
	if timeout != nil {
		_p1 := timeout
		_args[1] = _p1
		_end++
	}
	for _, __in := range arguments {
		__out := __in
		_args[_end] = __out
		_end++
	}
	_returned := _this.Value_JS.Call("setInterval", _args[0:_end]...)
	var (
		_converted int // javascript: long _what_return_name
	)
	_converted = (_returned).Int()
	_result = _converted
	return
}

func (_this *WorkerGlobalScope) ClearInterval(handle *int) {
	var (
		_args [1]interface{}
		_end  int
	)
	if handle != nil {
		_p0 := handle
		_args[0] = _p0
		_end++
	}
	_this.Value_JS.Call("clearInterval", _args[0:_end]...)
	return
}

func (_this *WorkerGlobalScope) QueueMicrotask(callback *webidl.VoidFunction) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if callback != nil {
		__callback0 = (*callback).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("queueMicrotask", _args[0:_end]...)
	return
}

func (_this *WorkerGlobalScope) CreateImageBitmap(image *Union, options *canvas.ImageBitmapOptions) (_result *javascript.Promise) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := image.JSValue()
	_args[0] = _p0
	_end++
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("createImageBitmap", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *WorkerGlobalScope) CreateImageBitmap2(image *Union, sx int, sy int, sw int, sh int, options *canvas.ImageBitmapOptions) (_result *javascript.Promise) {
	var (
		_args [6]interface{}
		_end  int
	)
	_p0 := image.JSValue()
	_args[0] = _p0
	_end++
	_p1 := sx
	_args[1] = _p1
	_end++
	_p2 := sy
	_args[2] = _p2
	_end++
	_p3 := sw
	_args[3] = _p3
	_end++
	_p4 := sh
	_args[4] = _p4
	_end++
	if options != nil {
		_p5 := options.JSValue()
		_args[5] = _p5
		_end++
	}
	_returned := _this.Value_JS.Call("createImageBitmap", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *WorkerGlobalScope) Fetch(input *Union, init *fetch.RequestInit) (_result *javascript.Promise) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := input.JSValue()
	_args[0] = _p0
	_end++
	if init != nil {
		_p1 := init.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("fetch", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

// interface: DedicatedWorkerGlobalScope
type DedicatedWorkerGlobalScope struct {
	WorkerGlobalScope
}

// DedicatedWorkerGlobalScopeFromJS is casting a js.Wrapper into DedicatedWorkerGlobalScope.
func DedicatedWorkerGlobalScopeFromJS(value js.Wrapper) *DedicatedWorkerGlobalScope {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &DedicatedWorkerGlobalScope{}
	ret.Value_JS = input
	return ret
}

// Name returning attribute 'name' with
// type string (idl: DOMString).
func (_this *DedicatedWorkerGlobalScope) Name() string {
	var ret string
	value := _this.Value_JS.Get("name")
	ret = (value).String()
	return ret
}

// Onmessage returning attribute 'onmessage' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *DedicatedWorkerGlobalScope) Onmessage() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onmessage")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnmessage setting attribute 'onmessage' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *DedicatedWorkerGlobalScope) SetOnmessage(value *domcore.EventHandler) {
	var __callback1 js.Value
	if value != nil {
		__callback1 = (*value).Value
	} else {
		__callback1 = js.Null()
	}
	input := __callback1
	_this.Value_JS.Set("onmessage", input)
}

// Onmessageerror returning attribute 'onmessageerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *DedicatedWorkerGlobalScope) Onmessageerror() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onmessageerror")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnmessageerror setting attribute 'onmessageerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *DedicatedWorkerGlobalScope) SetOnmessageerror(value *domcore.EventHandler) {
	var __callback2 js.Value
	if value != nil {
		__callback2 = (*value).Value
	} else {
		__callback2 = js.Null()
	}
	input := __callback2
	_this.Value_JS.Set("onmessageerror", input)
}

func (_this *DedicatedWorkerGlobalScope) PostMessage(message interface{}, transfer []*javascript.Object) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := message
	_args[0] = _p0
	_end++
	_p1 := js.Global().Get("Array").New(len(transfer))
	for __idx1, __seq_in1 := range transfer {
		__seq_out1 := __seq_in1.JSValue()
		_p1.SetIndex(__idx1, __seq_out1)
	}
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("postMessage", _args[0:_end]...)
	return
}

func (_this *DedicatedWorkerGlobalScope) PostMessage2(message interface{}, options *channel.PostMessageOptions) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := message
	_args[0] = _p0
	_end++
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_this.Value_JS.Call("postMessage", _args[0:_end]...)
	return
}

func (_this *DedicatedWorkerGlobalScope) Close() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("close", _args[0:_end]...)
	return
}

func (_this *DedicatedWorkerGlobalScope) RequestAnimationFrame(callback *htmlevent.FrameRequestCallback) (_result uint) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if callback != nil {
		__callback0 = (*callback).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("requestAnimationFrame", _args[0:_end]...)
	var (
		_converted uint // javascript: unsigned long _what_return_name
	)
	_converted = (uint)((_returned).Int())
	_result = _converted
	return
}

func (_this *DedicatedWorkerGlobalScope) CancelAnimationFrame(handle uint) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := handle
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("cancelAnimationFrame", _args[0:_end]...)
	return
}

// interface: SharedWorkerGlobalScope
type SharedWorkerGlobalScope struct {
	WorkerGlobalScope
}

// SharedWorkerGlobalScopeFromJS is casting a js.Wrapper into SharedWorkerGlobalScope.
func SharedWorkerGlobalScopeFromJS(value js.Wrapper) *SharedWorkerGlobalScope {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &SharedWorkerGlobalScope{}
	ret.Value_JS = input
	return ret
}

// Name returning attribute 'name' with
// type string (idl: DOMString).
func (_this *SharedWorkerGlobalScope) Name() string {
	var ret string
	value := _this.Value_JS.Get("name")
	ret = (value).String()
	return ret
}

// Onconnect returning attribute 'onconnect' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *SharedWorkerGlobalScope) Onconnect() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onconnect")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnconnect setting attribute 'onconnect' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *SharedWorkerGlobalScope) SetOnconnect(value *domcore.EventHandler) {
	var __callback1 js.Value
	if value != nil {
		__callback1 = (*value).Value
	} else {
		__callback1 = js.Null()
	}
	input := __callback1
	_this.Value_JS.Set("onconnect", input)
}

func (_this *SharedWorkerGlobalScope) Close() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("close", _args[0:_end]...)
	return
}

// interface: Worker
type Worker struct {
	domcore.EventTarget
}

// WorkerFromJS is casting a js.Wrapper into Worker.
func WorkerFromJS(value js.Wrapper) *Worker {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Worker{}
	ret.Value_JS = input
	return ret
}

func NewWorker(scriptURL string, options *WorkerOptions) (_result *Worker) {
	_klass := js.Global().Get("Worker")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := scriptURL
	_args[0] = _p0
	_end++
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *Worker // javascript: Worker _what_return_name
	)
	_converted = WorkerFromJS(_returned)
	_result = _converted
	return
}

// Onmessage returning attribute 'onmessage' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Worker) Onmessage() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onmessage")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnmessage setting attribute 'onmessage' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Worker) SetOnmessage(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onmessage", input)
}

// Onmessageerror returning attribute 'onmessageerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Worker) Onmessageerror() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onmessageerror")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnmessageerror setting attribute 'onmessageerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Worker) SetOnmessageerror(value *domcore.EventHandler) {
	var __callback1 js.Value
	if value != nil {
		__callback1 = (*value).Value
	} else {
		__callback1 = js.Null()
	}
	input := __callback1
	_this.Value_JS.Set("onmessageerror", input)
}

// Onerror returning attribute 'onerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Worker) Onerror() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onerror")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnerror setting attribute 'onerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Worker) SetOnerror(value *domcore.EventHandler) {
	var __callback2 js.Value
	if value != nil {
		__callback2 = (*value).Value
	} else {
		__callback2 = js.Null()
	}
	input := __callback2
	_this.Value_JS.Set("onerror", input)
}

func (_this *Worker) Terminate() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("terminate", _args[0:_end]...)
	return
}

func (_this *Worker) PostMessage(message interface{}, transfer []*javascript.Object) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := message
	_args[0] = _p0
	_end++
	_p1 := js.Global().Get("Array").New(len(transfer))
	for __idx1, __seq_in1 := range transfer {
		__seq_out1 := __seq_in1.JSValue()
		_p1.SetIndex(__idx1, __seq_out1)
	}
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("postMessage", _args[0:_end]...)
	return
}

func (_this *Worker) PostMessage2(message interface{}, options *channel.PostMessageOptions) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := message
	_args[0] = _p0
	_end++
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_this.Value_JS.Call("postMessage", _args[0:_end]...)
	return
}

// interface: SharedWorker
type SharedWorker struct {
	domcore.EventTarget
}

// SharedWorkerFromJS is casting a js.Wrapper into SharedWorker.
func SharedWorkerFromJS(value js.Wrapper) *SharedWorker {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &SharedWorker{}
	ret.Value_JS = input
	return ret
}

func NewSharedWorker(scriptURL string, options *Union) (_result *SharedWorker) {
	_klass := js.Global().Get("SharedWorker")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := scriptURL
	_args[0] = _p0
	_end++
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *SharedWorker // javascript: SharedWorker _what_return_name
	)
	_converted = SharedWorkerFromJS(_returned)
	_result = _converted
	return
}

// Port returning attribute 'port' with
// type channel.MessagePort (idl: MessagePort).
func (_this *SharedWorker) Port() *channel.MessagePort {
	var ret *channel.MessagePort
	value := _this.Value_JS.Get("port")
	ret = channel.MessagePortFromJS(value)
	return ret
}

// Onerror returning attribute 'onerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *SharedWorker) Onerror() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onerror")
	if value.Type() != js.TypeNull {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnerror setting attribute 'onerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *SharedWorker) SetOnerror(value *domcore.EventHandler) {
	var __callback1 js.Value
	if value != nil {
		__callback1 = (*value).Value
	} else {
		__callback1 = js.Null()
	}
	input := __callback1
	_this.Value_JS.Set("onerror", input)
}

// interface: WorkerNavigator
type WorkerNavigator struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *WorkerNavigator) JSValue() js.Value {
	return _this.Value_JS
}

// WorkerNavigatorFromJS is casting a js.Wrapper into WorkerNavigator.
func WorkerNavigatorFromJS(value js.Wrapper) *WorkerNavigator {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &WorkerNavigator{}
	ret.Value_JS = input
	return ret
}

// AppCodeName returning attribute 'appCodeName' with
// type string (idl: DOMString).
func (_this *WorkerNavigator) AppCodeName() string {
	var ret string
	value := _this.Value_JS.Get("appCodeName")
	ret = (value).String()
	return ret
}

// AppName returning attribute 'appName' with
// type string (idl: DOMString).
func (_this *WorkerNavigator) AppName() string {
	var ret string
	value := _this.Value_JS.Get("appName")
	ret = (value).String()
	return ret
}

// AppVersion returning attribute 'appVersion' with
// type string (idl: DOMString).
func (_this *WorkerNavigator) AppVersion() string {
	var ret string
	value := _this.Value_JS.Get("appVersion")
	ret = (value).String()
	return ret
}

// Platform returning attribute 'platform' with
// type string (idl: DOMString).
func (_this *WorkerNavigator) Platform() string {
	var ret string
	value := _this.Value_JS.Get("platform")
	ret = (value).String()
	return ret
}

// Product returning attribute 'product' with
// type string (idl: DOMString).
func (_this *WorkerNavigator) Product() string {
	var ret string
	value := _this.Value_JS.Get("product")
	ret = (value).String()
	return ret
}

// ProductSub returning attribute 'productSub' with
// type string (idl: DOMString).
func (_this *WorkerNavigator) ProductSub() string {
	var ret string
	value := _this.Value_JS.Get("productSub")
	ret = (value).String()
	return ret
}

// UserAgent returning attribute 'userAgent' with
// type string (idl: DOMString).
func (_this *WorkerNavigator) UserAgent() string {
	var ret string
	value := _this.Value_JS.Get("userAgent")
	ret = (value).String()
	return ret
}

// Vendor returning attribute 'vendor' with
// type string (idl: DOMString).
func (_this *WorkerNavigator) Vendor() string {
	var ret string
	value := _this.Value_JS.Get("vendor")
	ret = (value).String()
	return ret
}

// VendorSub returning attribute 'vendorSub' with
// type string (idl: DOMString).
func (_this *WorkerNavigator) VendorSub() string {
	var ret string
	value := _this.Value_JS.Get("vendorSub")
	ret = (value).String()
	return ret
}

// Oscpu returning attribute 'oscpu' with
// type string (idl: DOMString).
func (_this *WorkerNavigator) Oscpu() string {
	var ret string
	value := _this.Value_JS.Get("oscpu")
	ret = (value).String()
	return ret
}

// Language returning attribute 'language' with
// type string (idl: DOMString).
func (_this *WorkerNavigator) Language() string {
	var ret string
	value := _this.Value_JS.Get("language")
	ret = (value).String()
	return ret
}

// Languages returning attribute 'languages' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *WorkerNavigator) Languages() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("languages")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

// OnLine returning attribute 'onLine' with
// type bool (idl: boolean).
func (_this *WorkerNavigator) OnLine() bool {
	var ret bool
	value := _this.Value_JS.Get("onLine")
	ret = (value).Bool()
	return ret
}

// HardwareConcurrency returning attribute 'hardwareConcurrency' with
// type int (idl: unsigned long long).
func (_this *WorkerNavigator) HardwareConcurrency() int {
	var ret int
	value := _this.Value_JS.Get("hardwareConcurrency")
	ret = (value).Int()
	return ret
}

func (_this *WorkerNavigator) TaintEnabled() (_result bool) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("taintEnabled", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

// interface: WorkerLocation
type WorkerLocation struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *WorkerLocation) JSValue() js.Value {
	return _this.Value_JS
}

// WorkerLocationFromJS is casting a js.Wrapper into WorkerLocation.
func WorkerLocationFromJS(value js.Wrapper) *WorkerLocation {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &WorkerLocation{}
	ret.Value_JS = input
	return ret
}

// Href returning attribute 'href' with
// type string (idl: USVString).
func (_this *WorkerLocation) Href() string {
	var ret string
	value := _this.Value_JS.Get("href")
	ret = (value).String()
	return ret
}

// Origin returning attribute 'origin' with
// type string (idl: USVString).
func (_this *WorkerLocation) Origin() string {
	var ret string
	value := _this.Value_JS.Get("origin")
	ret = (value).String()
	return ret
}

// Protocol returning attribute 'protocol' with
// type string (idl: USVString).
func (_this *WorkerLocation) Protocol() string {
	var ret string
	value := _this.Value_JS.Get("protocol")
	ret = (value).String()
	return ret
}

// Host returning attribute 'host' with
// type string (idl: USVString).
func (_this *WorkerLocation) Host() string {
	var ret string
	value := _this.Value_JS.Get("host")
	ret = (value).String()
	return ret
}

// Hostname returning attribute 'hostname' with
// type string (idl: USVString).
func (_this *WorkerLocation) Hostname() string {
	var ret string
	value := _this.Value_JS.Get("hostname")
	ret = (value).String()
	return ret
}

// Port returning attribute 'port' with
// type string (idl: USVString).
func (_this *WorkerLocation) Port() string {
	var ret string
	value := _this.Value_JS.Get("port")
	ret = (value).String()
	return ret
}

// Pathname returning attribute 'pathname' with
// type string (idl: USVString).
func (_this *WorkerLocation) Pathname() string {
	var ret string
	value := _this.Value_JS.Get("pathname")
	ret = (value).String()
	return ret
}

// Search returning attribute 'search' with
// type string (idl: USVString).
func (_this *WorkerLocation) Search() string {
	var ret string
	value := _this.Value_JS.Get("search")
	ret = (value).String()
	return ret
}

// Hash returning attribute 'hash' with
// type string (idl: USVString).
func (_this *WorkerLocation) Hash() string {
	var ret string
	value := _this.Value_JS.Get("hash")
	ret = (value).String()
	return ret
}
