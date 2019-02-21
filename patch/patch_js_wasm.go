// Code generated by webidl-bind. DO NOT EDIT.

package patch

import "syscall/js"

// using following types:

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

// dictionary: MouseEventInit
type MouseEventInit struct {
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *MouseEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	return out
}

// MouseEventInitFromJS is allocating a new
// MouseEventInit object and copy all values from
// input javascript object
func MouseEventInitFromJS(value js.Wrapper) *MouseEventInit {
	var out MouseEventInit
	var ()
	return &out
}

// interface: FormData
type FormData struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FormData) JSValue() js.Value {
	return _this.Value_JS
}

// FormDataFromJS is casting a js.Wrapper into FormData.
func FormDataFromJS(value js.Wrapper) *FormData {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FormData{}
	ret.Value_JS = input
	return ret
}

// interface: MediaSource
type MediaSource struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *MediaSource) JSValue() js.Value {
	return _this.Value_JS
}

// MediaSourceFromJS is casting a js.Wrapper into MediaSource.
func MediaSourceFromJS(value js.Wrapper) *MediaSource {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &MediaSource{}
	ret.Value_JS = input
	return ret
}

// interface: MediaStream
type MediaStream struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *MediaStream) JSValue() js.Value {
	return _this.Value_JS
}

// MediaStreamFromJS is casting a js.Wrapper into MediaStream.
func MediaStreamFromJS(value js.Wrapper) *MediaStream {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &MediaStream{}
	ret.Value_JS = input
	return ret
}

// interface: ServiceWorker
type ServiceWorker struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *ServiceWorker) JSValue() js.Value {
	return _this.Value_JS
}

// ServiceWorkerFromJS is casting a js.Wrapper into ServiceWorker.
func ServiceWorkerFromJS(value js.Wrapper) *ServiceWorker {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ServiceWorker{}
	ret.Value_JS = input
	return ret
}

// interface: SVGScriptElement
type SVGScriptElement struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *SVGScriptElement) JSValue() js.Value {
	return _this.Value_JS
}

// SVGScriptElementFromJS is casting a js.Wrapper into SVGScriptElement.
func SVGScriptElementFromJS(value js.Wrapper) *SVGScriptElement {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &SVGScriptElement{}
	ret.Value_JS = input
	return ret
}

// interface: SVGImageElement
type SVGImageElement struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *SVGImageElement) JSValue() js.Value {
	return _this.Value_JS
}

// SVGImageElementFromJS is casting a js.Wrapper into SVGImageElement.
func SVGImageElementFromJS(value js.Wrapper) *SVGImageElement {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &SVGImageElement{}
	ret.Value_JS = input
	return ret
}

// interface: Uint8ClampedArray
type Uint8ClampedArray struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Uint8ClampedArray) JSValue() js.Value {
	return _this.Value_JS
}

// Uint8ClampedArrayFromJS is casting a js.Wrapper into Uint8ClampedArray.
func Uint8ClampedArrayFromJS(value js.Wrapper) *Uint8ClampedArray {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Uint8ClampedArray{}
	ret.Value_JS = input
	return ret
}

// interface: ByteString
type ByteString struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *ByteString) JSValue() js.Value {
	return _this.Value_JS
}

// ByteStringFromJS is casting a js.Wrapper into ByteString.
func ByteStringFromJS(value js.Wrapper) *ByteString {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ByteString{}
	ret.Value_JS = input
	return ret
}

// interface: MouseEvent
type MouseEvent struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *MouseEvent) JSValue() js.Value {
	return _this.Value_JS
}

// MouseEventFromJS is casting a js.Wrapper into MouseEvent.
func MouseEventFromJS(value js.Wrapper) *MouseEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &MouseEvent{}
	ret.Value_JS = input
	return ret
}

// interface: ReadableStream
type ReadableStream struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *ReadableStream) JSValue() js.Value {
	return _this.Value_JS
}

// ReadableStreamFromJS is casting a js.Wrapper into ReadableStream.
func ReadableStreamFromJS(value js.Wrapper) *ReadableStream {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ReadableStream{}
	ret.Value_JS = input
	return ret
}
