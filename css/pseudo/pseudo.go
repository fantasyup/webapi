// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package pseudo

import js "github.com/gowebapi/webapi/core/failjs"

import (
	"github.com/gowebapi/webapi/css/animations/webani"
	"github.com/gowebapi/webapi/css/cssom"
	"github.com/gowebapi/webapi/css/cssom/view"
	"github.com/gowebapi/webapi/dom"
	"github.com/gowebapi/webapi/dom/geometry"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// cssom.CSSStyleDeclaration
// dom.Element
// geometry.DOMPoint
// geometry.DOMPointInit
// geometry.DOMQuad
// geometry.DOMQuadInit
// geometry.DOMRectReadOnly
// javascript.Object
// view.BoxQuadOptions
// view.ConvertCoordinateOptions
// webani.Animation

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

// interface: CSSPseudoElement
type CSSPseudoElement struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *CSSPseudoElement) JSValue() js.Value {
	return _this.Value_JS
}

// CSSPseudoElementFromJS is casting a js.Wrapper into CSSPseudoElement.
func CSSPseudoElementFromJS(value js.Wrapper) *CSSPseudoElement {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSPseudoElement{}
	ret.Value_JS = input
	return ret
}

// Type returning attribute 'type' with
// type string (idl: DOMString).
func (_this *CSSPseudoElement) Type() string {
	var ret string
	value := _this.Value_JS.Get("type")
	ret = (value).String()
	return ret
}

// Element returning attribute 'element' with
// type dom.Element (idl: Element).
func (_this *CSSPseudoElement) Element() *dom.Element {
	var ret *dom.Element
	value := _this.Value_JS.Get("element")
	ret = dom.ElementFromJS(value)
	return ret
}

// Style returning attribute 'style' with
// type cssom.CSSStyleDeclaration (idl: CSSStyleDeclaration).
func (_this *CSSPseudoElement) Style() *cssom.CSSStyleDeclaration {
	var ret *cssom.CSSStyleDeclaration
	value := _this.Value_JS.Get("style")
	ret = cssom.CSSStyleDeclarationFromJS(value)
	return ret
}

func (_this *CSSPseudoElement) GetBoxQuads(options *view.BoxQuadOptions) (_result []*geometry.DOMQuad) {
	var (
		_args [1]interface{}
		_end  int
	)
	if options != nil {
		_p0 := options.JSValue()
		_args[0] = _p0
		_end++
	}
	_returned := _this.Value_JS.Call("getBoxQuads", _args[0:_end]...)
	var (
		_converted []*geometry.DOMQuad // javascript: sequence<DOMQuad> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]*geometry.DOMQuad, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *geometry.DOMQuad
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = geometry.DOMQuadFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

func (_this *CSSPseudoElement) ConvertQuadFromNode(quad *geometry.DOMQuadInit, from *Union, options *view.ConvertCoordinateOptions) (_result *geometry.DOMQuad) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := quad.JSValue()
	_args[0] = _p0
	_end++
	_p1 := from.JSValue()
	_args[1] = _p1
	_end++
	if options != nil {
		_p2 := options.JSValue()
		_args[2] = _p2
		_end++
	}
	_returned := _this.Value_JS.Call("convertQuadFromNode", _args[0:_end]...)
	var (
		_converted *geometry.DOMQuad // javascript: DOMQuad _what_return_name
	)
	_converted = geometry.DOMQuadFromJS(_returned)
	_result = _converted
	return
}

func (_this *CSSPseudoElement) ConvertRectFromNode(rect *geometry.DOMRectReadOnly, from *Union, options *view.ConvertCoordinateOptions) (_result *geometry.DOMQuad) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := rect.JSValue()
	_args[0] = _p0
	_end++
	_p1 := from.JSValue()
	_args[1] = _p1
	_end++
	if options != nil {
		_p2 := options.JSValue()
		_args[2] = _p2
		_end++
	}
	_returned := _this.Value_JS.Call("convertRectFromNode", _args[0:_end]...)
	var (
		_converted *geometry.DOMQuad // javascript: DOMQuad _what_return_name
	)
	_converted = geometry.DOMQuadFromJS(_returned)
	_result = _converted
	return
}

func (_this *CSSPseudoElement) ConvertPointFromNode(point *geometry.DOMPointInit, from *Union, options *view.ConvertCoordinateOptions) (_result *geometry.DOMPoint) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := point.JSValue()
	_args[0] = _p0
	_end++
	_p1 := from.JSValue()
	_args[1] = _p1
	_end++
	if options != nil {
		_p2 := options.JSValue()
		_args[2] = _p2
		_end++
	}
	_returned := _this.Value_JS.Call("convertPointFromNode", _args[0:_end]...)
	var (
		_converted *geometry.DOMPoint // javascript: DOMPoint _what_return_name
	)
	_converted = geometry.DOMPointFromJS(_returned)
	_result = _converted
	return
}

func (_this *CSSPseudoElement) Animate(keyframes *javascript.Object, options *Union) (_result *webani.Animation) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := keyframes.JSValue()
	_args[0] = _p0
	_end++
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("animate", _args[0:_end]...)
	var (
		_converted *webani.Animation // javascript: Animation _what_return_name
	)
	_converted = webani.AnimationFromJS(_returned)
	_result = _converted
	return
}

func (_this *CSSPseudoElement) GetAnimations() (_result []*webani.Animation) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getAnimations", _args[0:_end]...)
	var (
		_converted []*webani.Animation // javascript: sequence<Animation> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]*webani.Animation, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *webani.Animation
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = webani.AnimationFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

// interface: CSSPseudoElementList
type CSSPseudoElementList struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *CSSPseudoElementList) JSValue() js.Value {
	return _this.Value_JS
}

// CSSPseudoElementListFromJS is casting a js.Wrapper into CSSPseudoElementList.
func CSSPseudoElementListFromJS(value js.Wrapper) *CSSPseudoElementList {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSPseudoElementList{}
	ret.Value_JS = input
	return ret
}

// Length returning attribute 'length' with
// type uint (idl: unsigned long).
func (_this *CSSPseudoElementList) Length() uint {
	var ret uint
	value := _this.Value_JS.Get("length")
	ret = (uint)((value).Int())
	return ret
}

func (_this *CSSPseudoElementList) Item(index uint) (_result *CSSPseudoElement) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("item", _args[0:_end]...)
	var (
		_converted *CSSPseudoElement // javascript: CSSPseudoElement _what_return_name
	)
	_converted = CSSPseudoElementFromJS(_returned)
	_result = _converted
	return
}

func (_this *CSSPseudoElementList) GetByType(_type string) (_result *CSSPseudoElement) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("getByType", _args[0:_end]...)
	var (
		_converted *CSSPseudoElement // javascript: CSSPseudoElement _what_return_name
	)
	_converted = CSSPseudoElementFromJS(_returned)
	_result = _converted
	return
}
