// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package counterstyles

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/css/cssom"
)

// using following types:
// cssom.CSSRule

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

// interface: CSSCounterStyleRule
type CSSCounterStyleRule struct {
	cssom.CSSRule
}

// CSSCounterStyleRuleFromJS is casting a js.Wrapper into CSSCounterStyleRule.
func CSSCounterStyleRuleFromJS(value js.Wrapper) *CSSCounterStyleRule {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSCounterStyleRule{}
	ret.Value_JS = input
	return ret
}

// Name returning attribute 'name' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) Name() string {
	var ret string
	value := _this.Value_JS.Get("name")
	ret = (value).String()
	return ret
}

// SetName setting attribute 'name' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) SetName(value string) {
	input := value
	_this.Value_JS.Set("name", input)
}

// System returning attribute 'system' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) System() string {
	var ret string
	value := _this.Value_JS.Get("system")
	ret = (value).String()
	return ret
}

// SetSystem setting attribute 'system' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) SetSystem(value string) {
	input := value
	_this.Value_JS.Set("system", input)
}

// Symbols returning attribute 'symbols' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) Symbols() string {
	var ret string
	value := _this.Value_JS.Get("symbols")
	ret = (value).String()
	return ret
}

// SetSymbols setting attribute 'symbols' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) SetSymbols(value string) {
	input := value
	_this.Value_JS.Set("symbols", input)
}

// AdditiveSymbols returning attribute 'additiveSymbols' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) AdditiveSymbols() string {
	var ret string
	value := _this.Value_JS.Get("additiveSymbols")
	ret = (value).String()
	return ret
}

// SetAdditiveSymbols setting attribute 'additiveSymbols' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) SetAdditiveSymbols(value string) {
	input := value
	_this.Value_JS.Set("additiveSymbols", input)
}

// Negative returning attribute 'negative' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) Negative() string {
	var ret string
	value := _this.Value_JS.Get("negative")
	ret = (value).String()
	return ret
}

// SetNegative setting attribute 'negative' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) SetNegative(value string) {
	input := value
	_this.Value_JS.Set("negative", input)
}

// Prefix returning attribute 'prefix' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) Prefix() string {
	var ret string
	value := _this.Value_JS.Get("prefix")
	ret = (value).String()
	return ret
}

// SetPrefix setting attribute 'prefix' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) SetPrefix(value string) {
	input := value
	_this.Value_JS.Set("prefix", input)
}

// Suffix returning attribute 'suffix' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) Suffix() string {
	var ret string
	value := _this.Value_JS.Get("suffix")
	ret = (value).String()
	return ret
}

// SetSuffix setting attribute 'suffix' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) SetSuffix(value string) {
	input := value
	_this.Value_JS.Set("suffix", input)
}

// Range returning attribute 'range' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) Range() string {
	var ret string
	value := _this.Value_JS.Get("range")
	ret = (value).String()
	return ret
}

// SetRange setting attribute 'range' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) SetRange(value string) {
	input := value
	_this.Value_JS.Set("range", input)
}

// Pad returning attribute 'pad' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) Pad() string {
	var ret string
	value := _this.Value_JS.Get("pad")
	ret = (value).String()
	return ret
}

// SetPad setting attribute 'pad' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) SetPad(value string) {
	input := value
	_this.Value_JS.Set("pad", input)
}

// SpeakAs returning attribute 'speakAs' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) SpeakAs() string {
	var ret string
	value := _this.Value_JS.Get("speakAs")
	ret = (value).String()
	return ret
}

// SetSpeakAs setting attribute 'speakAs' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) SetSpeakAs(value string) {
	input := value
	_this.Value_JS.Set("speakAs", input)
}

// Fallback returning attribute 'fallback' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) Fallback() string {
	var ret string
	value := _this.Value_JS.Get("fallback")
	ret = (value).String()
	return ret
}

// SetFallback setting attribute 'fallback' with
// type string (idl: DOMString).
func (_this *CSSCounterStyleRule) SetFallback(value string) {
	input := value
	_this.Value_JS.Set("fallback", input)
}
