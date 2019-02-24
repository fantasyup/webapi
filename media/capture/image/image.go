// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package image

import js "github.com/gowebapi/webapi/core/failjs"

import (
	"github.com/gowebapi/webapi/javascript"
	"github.com/gowebapi/webapi/media/capture/local"
	"github.com/gowebapi/webapi/media/mediatype"
)

// using following types:
// javascript.FrozenArray
// javascript.Promise
// local.MediaStreamTrack
// mediatype.MediaSettingsRange
// mediatype.Point2D

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

// enum: FillLightMode
type FillLightMode int

const (
	AutoFillLightMode FillLightMode = iota
	OffFillLightMode
	FlashFillLightMode
)

var fillLightModeToWasmTable = []string{
	"auto", "off", "flash",
}

var fillLightModeFromWasmTable = map[string]FillLightMode{
	"auto": AutoFillLightMode, "off": OffFillLightMode, "flash": FlashFillLightMode,
}

// JSValue is converting this enum into a java object
func (this *FillLightMode) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this FillLightMode) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(fillLightModeToWasmTable) {
		return fillLightModeToWasmTable[idx]
	}
	panic("unknown input value")
}

// FillLightModeFromJS is converting a javascript value into
// a FillLightMode enum value.
func FillLightModeFromJS(value js.Value) FillLightMode {
	key := value.String()
	conv, ok := fillLightModeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: RedEyeReduction
type RedEyeReduction int

const (
	NeverRedEyeReduction RedEyeReduction = iota
	AlwaysRedEyeReduction
	ControllableRedEyeReduction
)

var redEyeReductionToWasmTable = []string{
	"never", "always", "controllable",
}

var redEyeReductionFromWasmTable = map[string]RedEyeReduction{
	"never": NeverRedEyeReduction, "always": AlwaysRedEyeReduction, "controllable": ControllableRedEyeReduction,
}

// JSValue is converting this enum into a java object
func (this *RedEyeReduction) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this RedEyeReduction) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(redEyeReductionToWasmTable) {
		return redEyeReductionToWasmTable[idx]
	}
	panic("unknown input value")
}

// RedEyeReductionFromJS is converting a javascript value into
// a RedEyeReduction enum value.
func RedEyeReductionFromJS(value js.Value) RedEyeReduction {
	key := value.String()
	conv, ok := redEyeReductionFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// dictionary: ConstrainPoint2DParameters
type ConstrainPoint2DParameters struct {
	Exact []*mediatype.Point2D
	Ideal []*mediatype.Point2D
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ConstrainPoint2DParameters) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := js.Global().Get("Array").New(len(_this.Exact))
	for __idx0, __seq_in0 := range _this.Exact {
		__seq_out0 := __seq_in0.JSValue()
		value0.SetIndex(__idx0, __seq_out0)
	}
	out.Set("exact", value0)
	value1 := js.Global().Get("Array").New(len(_this.Ideal))
	for __idx1, __seq_in1 := range _this.Ideal {
		__seq_out1 := __seq_in1.JSValue()
		value1.SetIndex(__idx1, __seq_out1)
	}
	out.Set("ideal", value1)
	return out
}

// ConstrainPoint2DParametersFromJS is allocating a new
// ConstrainPoint2DParameters object and copy all values from
// input javascript object
func ConstrainPoint2DParametersFromJS(value js.Wrapper) *ConstrainPoint2DParameters {
	input := value.JSValue()
	var out ConstrainPoint2DParameters
	var (
		value0 []*mediatype.Point2D // javascript: sequence<Point2D> {exact Exact exact}
		value1 []*mediatype.Point2D // javascript: sequence<Point2D> {ideal Ideal ideal}
	)
	__length0 := input.Get("exact").Length()
	__array0 := make([]*mediatype.Point2D, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *mediatype.Point2D
		__seq_in0 := input.Get("exact").Index(__idx0)
		__seq_out0 = mediatype.Point2DFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	value0 = __array0
	out.Exact = value0
	__length1 := input.Get("ideal").Length()
	__array1 := make([]*mediatype.Point2D, __length1, __length1)
	for __idx1 := 0; __idx1 < __length1; __idx1++ {
		var __seq_out1 *mediatype.Point2D
		__seq_in1 := input.Get("ideal").Index(__idx1)
		__seq_out1 = mediatype.Point2DFromJS(__seq_in1)
		__array1[__idx1] = __seq_out1
	}
	value1 = __array1
	out.Ideal = value1
	return &out
}

// dictionary: PhotoSettings
type PhotoSettings struct {
	FillLightMode   FillLightMode
	ImageHeight     float64
	ImageWidth      float64
	RedEyeReduction bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *PhotoSettings) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.FillLightMode.JSValue()
	out.Set("fillLightMode", value0)
	value1 := _this.ImageHeight
	out.Set("imageHeight", value1)
	value2 := _this.ImageWidth
	out.Set("imageWidth", value2)
	value3 := _this.RedEyeReduction
	out.Set("redEyeReduction", value3)
	return out
}

// PhotoSettingsFromJS is allocating a new
// PhotoSettings object and copy all values from
// input javascript object
func PhotoSettingsFromJS(value js.Wrapper) *PhotoSettings {
	input := value.JSValue()
	var out PhotoSettings
	var (
		value0 FillLightMode // javascript: FillLightMode {fillLightMode FillLightMode fillLightMode}
		value1 float64       // javascript: double {imageHeight ImageHeight imageHeight}
		value2 float64       // javascript: double {imageWidth ImageWidth imageWidth}
		value3 bool          // javascript: boolean {redEyeReduction RedEyeReduction redEyeReduction}
	)
	value0 = FillLightModeFromJS(input.Get("fillLightMode"))
	out.FillLightMode = value0
	value1 = (input.Get("imageHeight")).Float()
	out.ImageHeight = value1
	value2 = (input.Get("imageWidth")).Float()
	out.ImageWidth = value2
	value3 = (input.Get("redEyeReduction")).Bool()
	out.RedEyeReduction = value3
	return &out
}

// interface: ImageCapture
type ImageCapture struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *ImageCapture) JSValue() js.Value {
	return _this.Value_JS
}

// ImageCaptureFromJS is casting a js.Wrapper into ImageCapture.
func ImageCaptureFromJS(value js.Wrapper) *ImageCapture {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ImageCapture{}
	ret.Value_JS = input
	return ret
}

func NewImageCapture(videoTrack *local.MediaStreamTrack) (_result *ImageCapture) {
	_klass := js.Global().Get("ImageCapture")
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := videoTrack.JSValue()
	_args[0] = _p0
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *ImageCapture // javascript: ImageCapture _what_return_name
	)
	_converted = ImageCaptureFromJS(_returned)
	_result = _converted
	return
}

// Track returning attribute 'track' with
// type local.MediaStreamTrack (idl: MediaStreamTrack).
func (_this *ImageCapture) Track() *local.MediaStreamTrack {
	var ret *local.MediaStreamTrack
	value := _this.Value_JS.Get("track")
	ret = local.MediaStreamTrackFromJS(value)
	return ret
}

func (_this *ImageCapture) TakePhoto(photoSettings *PhotoSettings) (_result *javascript.Promise) {
	var (
		_args [1]interface{}
		_end  int
	)
	if photoSettings != nil {
		_p0 := photoSettings.JSValue()
		_args[0] = _p0
		_end++
	}
	_returned := _this.Value_JS.Call("takePhoto", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *ImageCapture) GetPhotoCapabilities() (_result *javascript.Promise) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getPhotoCapabilities", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *ImageCapture) GetPhotoSettings() (_result *javascript.Promise) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getPhotoSettings", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *ImageCapture) GrabFrame() (_result *javascript.Promise) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("grabFrame", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

// interface: PhotoCapabilities
type PhotoCapabilities struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PhotoCapabilities) JSValue() js.Value {
	return _this.Value_JS
}

// PhotoCapabilitiesFromJS is casting a js.Wrapper into PhotoCapabilities.
func PhotoCapabilitiesFromJS(value js.Wrapper) *PhotoCapabilities {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PhotoCapabilities{}
	ret.Value_JS = input
	return ret
}

// RedEyeReduction returning attribute 'redEyeReduction' with
// type RedEyeReduction (idl: RedEyeReduction).
func (_this *PhotoCapabilities) RedEyeReduction() RedEyeReduction {
	var ret RedEyeReduction
	value := _this.Value_JS.Get("redEyeReduction")
	ret = RedEyeReductionFromJS(value)
	return ret
}

// ImageHeight returning attribute 'imageHeight' with
// type mediatype.MediaSettingsRange (idl: MediaSettingsRange).
func (_this *PhotoCapabilities) ImageHeight() *mediatype.MediaSettingsRange {
	var ret *mediatype.MediaSettingsRange
	value := _this.Value_JS.Get("imageHeight")
	ret = mediatype.MediaSettingsRangeFromJS(value)
	return ret
}

// ImageWidth returning attribute 'imageWidth' with
// type mediatype.MediaSettingsRange (idl: MediaSettingsRange).
func (_this *PhotoCapabilities) ImageWidth() *mediatype.MediaSettingsRange {
	var ret *mediatype.MediaSettingsRange
	value := _this.Value_JS.Get("imageWidth")
	ret = mediatype.MediaSettingsRangeFromJS(value)
	return ret
}

// FillLightMode returning attribute 'fillLightMode' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *PhotoCapabilities) FillLightMode() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("fillLightMode")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}
