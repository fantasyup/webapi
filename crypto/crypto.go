// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package crypto

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// javascript.Object
// javascript.Promise

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

// enum: KeyFormat
type KeyFormat int

const (
	RawKeyFormat KeyFormat = iota
	SpkiKeyFormat
	Pkcs8KeyFormat
	JwkKeyFormat
)

var keyFormatToWasmTable = []string{
	"raw", "spki", "pkcs8", "jwk",
}

var keyFormatFromWasmTable = map[string]KeyFormat{
	"raw": RawKeyFormat, "spki": SpkiKeyFormat, "pkcs8": Pkcs8KeyFormat, "jwk": JwkKeyFormat,
}

// JSValue is converting this enum into a java object
func (this *KeyFormat) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this KeyFormat) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(keyFormatToWasmTable) {
		return keyFormatToWasmTable[idx]
	}
	panic("unknown input value")
}

// KeyFormatFromJS is converting a javascript value into
// a KeyFormat enum value.
func KeyFormatFromJS(value js.Value) KeyFormat {
	key := value.String()
	conv, ok := keyFormatFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: KeyType
type KeyType int

const (
	PublicKeyType KeyType = iota
	PrivateKeyType
	SecretKeyType
)

var keyTypeToWasmTable = []string{
	"public", "private", "secret",
}

var keyTypeFromWasmTable = map[string]KeyType{
	"public": PublicKeyType, "private": PrivateKeyType, "secret": SecretKeyType,
}

// JSValue is converting this enum into a java object
func (this *KeyType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this KeyType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(keyTypeToWasmTable) {
		return keyTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// KeyTypeFromJS is converting a javascript value into
// a KeyType enum value.
func KeyTypeFromJS(value js.Value) KeyType {
	key := value.String()
	conv, ok := keyTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: KeyUsage
type KeyUsage int

const (
	EncryptKeyUsage KeyUsage = iota
	DecryptKeyUsage
	SignKeyUsage
	VerifyKeyUsage
	DeriveKeyKeyUsage
	DeriveBitsKeyUsage
	WrapKeyKeyUsage
	UnwrapKeyKeyUsage
)

var keyUsageToWasmTable = []string{
	"encrypt", "decrypt", "sign", "verify", "deriveKey", "deriveBits", "wrapKey", "unwrapKey",
}

var keyUsageFromWasmTable = map[string]KeyUsage{
	"encrypt": EncryptKeyUsage, "decrypt": DecryptKeyUsage, "sign": SignKeyUsage, "verify": VerifyKeyUsage, "deriveKey": DeriveKeyKeyUsage, "deriveBits": DeriveBitsKeyUsage, "wrapKey": WrapKeyKeyUsage, "unwrapKey": UnwrapKeyKeyUsage,
}

// JSValue is converting this enum into a java object
func (this *KeyUsage) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this KeyUsage) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(keyUsageToWasmTable) {
		return keyUsageToWasmTable[idx]
	}
	panic("unknown input value")
}

// KeyUsageFromJS is converting a javascript value into
// a KeyUsage enum value.
func KeyUsageFromJS(value js.Value) KeyUsage {
	key := value.String()
	conv, ok := keyUsageFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// dictionary: JsonWebKey
type JsonWebKey struct {
	Kty    string
	Use    string
	KeyOps []string
	Alg    string
	Ext    bool
	Crv    string
	X      string
	Y      string
	D      string
	N      string
	E      string
	P      string
	Q      string
	Dp     string
	Dq     string
	Qi     string
	Oth    []*RsaOtherPrimesInfo
	K      string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *JsonWebKey) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Kty
	out.Set("kty", value0)
	value1 := _this.Use
	out.Set("use", value1)
	value2 := js.Global().Get("Array").New(len(_this.KeyOps))
	for __idx2, __seq_in2 := range _this.KeyOps {
		__seq_out2 := __seq_in2
		value2.SetIndex(__idx2, __seq_out2)
	}
	out.Set("key_ops", value2)
	value3 := _this.Alg
	out.Set("alg", value3)
	value4 := _this.Ext
	out.Set("ext", value4)
	value5 := _this.Crv
	out.Set("crv", value5)
	value6 := _this.X
	out.Set("x", value6)
	value7 := _this.Y
	out.Set("y", value7)
	value8 := _this.D
	out.Set("d", value8)
	value9 := _this.N
	out.Set("n", value9)
	value10 := _this.E
	out.Set("e", value10)
	value11 := _this.P
	out.Set("p", value11)
	value12 := _this.Q
	out.Set("q", value12)
	value13 := _this.Dp
	out.Set("dp", value13)
	value14 := _this.Dq
	out.Set("dq", value14)
	value15 := _this.Qi
	out.Set("qi", value15)
	value16 := js.Global().Get("Array").New(len(_this.Oth))
	for __idx16, __seq_in16 := range _this.Oth {
		__seq_out16 := __seq_in16.JSValue()
		value16.SetIndex(__idx16, __seq_out16)
	}
	out.Set("oth", value16)
	value17 := _this.K
	out.Set("k", value17)
	return out
}

// JsonWebKeyFromJS is allocating a new
// JsonWebKey object and copy all values from
// input javascript object
func JsonWebKeyFromJS(value js.Wrapper) *JsonWebKey {
	input := value.JSValue()
	var out JsonWebKey
	var (
		value0  string                // javascript: DOMString {kty Kty kty}
		value1  string                // javascript: DOMString {use Use use}
		value2  []string              // javascript: sequence<DOMString> {key_ops KeyOps keyOps}
		value3  string                // javascript: DOMString {alg Alg alg}
		value4  bool                  // javascript: boolean {ext Ext ext}
		value5  string                // javascript: DOMString {crv Crv crv}
		value6  string                // javascript: DOMString {x X x}
		value7  string                // javascript: DOMString {y Y y}
		value8  string                // javascript: DOMString {d D d}
		value9  string                // javascript: DOMString {n N n}
		value10 string                // javascript: DOMString {e E e}
		value11 string                // javascript: DOMString {p P p}
		value12 string                // javascript: DOMString {q Q q}
		value13 string                // javascript: DOMString {dp Dp dp}
		value14 string                // javascript: DOMString {dq Dq dq}
		value15 string                // javascript: DOMString {qi Qi qi}
		value16 []*RsaOtherPrimesInfo // javascript: sequence<RsaOtherPrimesInfo> {oth Oth oth}
		value17 string                // javascript: DOMString {k K k}
	)
	value0 = (input.Get("kty")).String()
	out.Kty = value0
	value1 = (input.Get("use")).String()
	out.Use = value1
	__length2 := input.Get("key_ops").Length()
	__array2 := make([]string, __length2, __length2)
	for __idx2 := 0; __idx2 < __length2; __idx2++ {
		var __seq_out2 string
		__seq_in2 := input.Get("key_ops").Index(__idx2)
		__seq_out2 = (__seq_in2).String()
		__array2[__idx2] = __seq_out2
	}
	value2 = __array2
	out.KeyOps = value2
	value3 = (input.Get("alg")).String()
	out.Alg = value3
	value4 = (input.Get("ext")).Bool()
	out.Ext = value4
	value5 = (input.Get("crv")).String()
	out.Crv = value5
	value6 = (input.Get("x")).String()
	out.X = value6
	value7 = (input.Get("y")).String()
	out.Y = value7
	value8 = (input.Get("d")).String()
	out.D = value8
	value9 = (input.Get("n")).String()
	out.N = value9
	value10 = (input.Get("e")).String()
	out.E = value10
	value11 = (input.Get("p")).String()
	out.P = value11
	value12 = (input.Get("q")).String()
	out.Q = value12
	value13 = (input.Get("dp")).String()
	out.Dp = value13
	value14 = (input.Get("dq")).String()
	out.Dq = value14
	value15 = (input.Get("qi")).String()
	out.Qi = value15
	__length16 := input.Get("oth").Length()
	__array16 := make([]*RsaOtherPrimesInfo, __length16, __length16)
	for __idx16 := 0; __idx16 < __length16; __idx16++ {
		var __seq_out16 *RsaOtherPrimesInfo
		__seq_in16 := input.Get("oth").Index(__idx16)
		__seq_out16 = RsaOtherPrimesInfoFromJS(__seq_in16)
		__array16[__idx16] = __seq_out16
	}
	value16 = __array16
	out.Oth = value16
	value17 = (input.Get("k")).String()
	out.K = value17
	return &out
}

// dictionary: RsaOtherPrimesInfo
type RsaOtherPrimesInfo struct {
	R string
	D string
	T string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *RsaOtherPrimesInfo) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.R
	out.Set("r", value0)
	value1 := _this.D
	out.Set("d", value1)
	value2 := _this.T
	out.Set("t", value2)
	return out
}

// RsaOtherPrimesInfoFromJS is allocating a new
// RsaOtherPrimesInfo object and copy all values from
// input javascript object
func RsaOtherPrimesInfoFromJS(value js.Wrapper) *RsaOtherPrimesInfo {
	input := value.JSValue()
	var out RsaOtherPrimesInfo
	var (
		value0 string // javascript: DOMString {r R r}
		value1 string // javascript: DOMString {d D d}
		value2 string // javascript: DOMString {t T t}
	)
	value0 = (input.Get("r")).String()
	out.R = value0
	value1 = (input.Get("d")).String()
	out.D = value1
	value2 = (input.Get("t")).String()
	out.T = value2
	return &out
}

// interface: Crypto
type Crypto struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Crypto) JSValue() js.Value {
	return _this.Value_JS
}

// CryptoFromJS is casting a js.Wrapper into Crypto.
func CryptoFromJS(value js.Wrapper) *Crypto {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Crypto{}
	ret.Value_JS = input
	return ret
}

// Subtle returning attribute 'subtle' with
// type SubtleCrypto (idl: SubtleCrypto).
func (_this *Crypto) Subtle() *SubtleCrypto {
	var ret *SubtleCrypto
	value := _this.Value_JS.Get("subtle")
	ret = SubtleCryptoFromJS(value)
	return ret
}

func (_this *Crypto) GetRandomValues(array *Union) (_result *Union) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := array.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("getRandomValues", _args[0:_end]...)
	var (
		_converted *Union // javascript: Union _what_return_name
	)
	_converted = UnionFromJS(_returned)
	_result = _converted
	return
}

// interface: CryptoKey
type CryptoKey struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *CryptoKey) JSValue() js.Value {
	return _this.Value_JS
}

// CryptoKeyFromJS is casting a js.Wrapper into CryptoKey.
func CryptoKeyFromJS(value js.Wrapper) *CryptoKey {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CryptoKey{}
	ret.Value_JS = input
	return ret
}

// Type returning attribute 'type' with
// type KeyType (idl: KeyType).
func (_this *CryptoKey) Type() KeyType {
	var ret KeyType
	value := _this.Value_JS.Get("type")
	ret = KeyTypeFromJS(value)
	return ret
}

// Extractable returning attribute 'extractable' with
// type bool (idl: boolean).
func (_this *CryptoKey) Extractable() bool {
	var ret bool
	value := _this.Value_JS.Get("extractable")
	ret = (value).Bool()
	return ret
}

// Algorithm returning attribute 'algorithm' with
// type javascript.Object (idl: object).
func (_this *CryptoKey) Algorithm() *javascript.Object {
	var ret *javascript.Object
	value := _this.Value_JS.Get("algorithm")
	ret = javascript.ObjectFromJS(value)
	return ret
}

// Usages returning attribute 'usages' with
// type javascript.Object (idl: object).
func (_this *CryptoKey) Usages() *javascript.Object {
	var ret *javascript.Object
	value := _this.Value_JS.Get("usages")
	ret = javascript.ObjectFromJS(value)
	return ret
}

// interface: SubtleCrypto
type SubtleCrypto struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *SubtleCrypto) JSValue() js.Value {
	return _this.Value_JS
}

// SubtleCryptoFromJS is casting a js.Wrapper into SubtleCrypto.
func SubtleCryptoFromJS(value js.Wrapper) *SubtleCrypto {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &SubtleCrypto{}
	ret.Value_JS = input
	return ret
}

func (_this *SubtleCrypto) Encrypt(algorithm *Union, key *CryptoKey, data *Union) (_result *javascript.Promise) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := algorithm.JSValue()
	_args[0] = _p0
	_end++
	_p1 := key.JSValue()
	_args[1] = _p1
	_end++
	_p2 := data.JSValue()
	_args[2] = _p2
	_end++
	_returned := _this.Value_JS.Call("encrypt", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *SubtleCrypto) Decrypt(algorithm *Union, key *CryptoKey, data *Union) (_result *javascript.Promise) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := algorithm.JSValue()
	_args[0] = _p0
	_end++
	_p1 := key.JSValue()
	_args[1] = _p1
	_end++
	_p2 := data.JSValue()
	_args[2] = _p2
	_end++
	_returned := _this.Value_JS.Call("decrypt", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *SubtleCrypto) Sign(algorithm *Union, key *CryptoKey, data *Union) (_result *javascript.Promise) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := algorithm.JSValue()
	_args[0] = _p0
	_end++
	_p1 := key.JSValue()
	_args[1] = _p1
	_end++
	_p2 := data.JSValue()
	_args[2] = _p2
	_end++
	_returned := _this.Value_JS.Call("sign", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *SubtleCrypto) Verify(algorithm *Union, key *CryptoKey, signature *Union, data *Union) (_result *javascript.Promise) {
	var (
		_args [4]interface{}
		_end  int
	)
	_p0 := algorithm.JSValue()
	_args[0] = _p0
	_end++
	_p1 := key.JSValue()
	_args[1] = _p1
	_end++
	_p2 := signature.JSValue()
	_args[2] = _p2
	_end++
	_p3 := data.JSValue()
	_args[3] = _p3
	_end++
	_returned := _this.Value_JS.Call("verify", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *SubtleCrypto) Digest(algorithm *Union, data *Union) (_result *javascript.Promise) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := algorithm.JSValue()
	_args[0] = _p0
	_end++
	_p1 := data.JSValue()
	_args[1] = _p1
	_end++
	_returned := _this.Value_JS.Call("digest", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *SubtleCrypto) GenerateKey(algorithm *Union, extractable bool, keyUsages []KeyUsage) (_result *javascript.Promise) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := algorithm.JSValue()
	_args[0] = _p0
	_end++
	_p1 := extractable
	_args[1] = _p1
	_end++
	_p2 := js.Global().Get("Array").New(len(keyUsages))
	for __idx2, __seq_in2 := range keyUsages {
		__seq_out2 := __seq_in2.JSValue()
		_p2.SetIndex(__idx2, __seq_out2)
	}
	_args[2] = _p2
	_end++
	_returned := _this.Value_JS.Call("generateKey", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *SubtleCrypto) DeriveKey(algorithm *Union, baseKey *CryptoKey, derivedKeyType *Union, extractable bool, keyUsages []KeyUsage) (_result *javascript.Promise) {
	var (
		_args [5]interface{}
		_end  int
	)
	_p0 := algorithm.JSValue()
	_args[0] = _p0
	_end++
	_p1 := baseKey.JSValue()
	_args[1] = _p1
	_end++
	_p2 := derivedKeyType.JSValue()
	_args[2] = _p2
	_end++
	_p3 := extractable
	_args[3] = _p3
	_end++
	_p4 := js.Global().Get("Array").New(len(keyUsages))
	for __idx4, __seq_in4 := range keyUsages {
		__seq_out4 := __seq_in4.JSValue()
		_p4.SetIndex(__idx4, __seq_out4)
	}
	_args[4] = _p4
	_end++
	_returned := _this.Value_JS.Call("deriveKey", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *SubtleCrypto) DeriveBits(algorithm *Union, baseKey *CryptoKey, length uint) (_result *javascript.Promise) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := algorithm.JSValue()
	_args[0] = _p0
	_end++
	_p1 := baseKey.JSValue()
	_args[1] = _p1
	_end++
	_p2 := length
	_args[2] = _p2
	_end++
	_returned := _this.Value_JS.Call("deriveBits", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *SubtleCrypto) ImportKey(format KeyFormat, keyData *Union, algorithm *Union, extractable bool, keyUsages []KeyUsage) (_result *javascript.Promise) {
	var (
		_args [5]interface{}
		_end  int
	)
	_p0 := format.JSValue()
	_args[0] = _p0
	_end++
	_p1 := keyData.JSValue()
	_args[1] = _p1
	_end++
	_p2 := algorithm.JSValue()
	_args[2] = _p2
	_end++
	_p3 := extractable
	_args[3] = _p3
	_end++
	_p4 := js.Global().Get("Array").New(len(keyUsages))
	for __idx4, __seq_in4 := range keyUsages {
		__seq_out4 := __seq_in4.JSValue()
		_p4.SetIndex(__idx4, __seq_out4)
	}
	_args[4] = _p4
	_end++
	_returned := _this.Value_JS.Call("importKey", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *SubtleCrypto) ExportKey(format KeyFormat, key *CryptoKey) (_result *javascript.Promise) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := format.JSValue()
	_args[0] = _p0
	_end++
	_p1 := key.JSValue()
	_args[1] = _p1
	_end++
	_returned := _this.Value_JS.Call("exportKey", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *SubtleCrypto) WrapKey(format KeyFormat, key *CryptoKey, wrappingKey *CryptoKey, wrapAlgorithm *Union) (_result *javascript.Promise) {
	var (
		_args [4]interface{}
		_end  int
	)
	_p0 := format.JSValue()
	_args[0] = _p0
	_end++
	_p1 := key.JSValue()
	_args[1] = _p1
	_end++
	_p2 := wrappingKey.JSValue()
	_args[2] = _p2
	_end++
	_p3 := wrapAlgorithm.JSValue()
	_args[3] = _p3
	_end++
	_returned := _this.Value_JS.Call("wrapKey", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *SubtleCrypto) UnwrapKey(format KeyFormat, wrappedKey *Union, unwrappingKey *CryptoKey, unwrapAlgorithm *Union, unwrappedKeyAlgorithm *Union, extractable bool, keyUsages []KeyUsage) (_result *javascript.Promise) {
	var (
		_args [7]interface{}
		_end  int
	)
	_p0 := format.JSValue()
	_args[0] = _p0
	_end++
	_p1 := wrappedKey.JSValue()
	_args[1] = _p1
	_end++
	_p2 := unwrappingKey.JSValue()
	_args[2] = _p2
	_end++
	_p3 := unwrapAlgorithm.JSValue()
	_args[3] = _p3
	_end++
	_p4 := unwrappedKeyAlgorithm.JSValue()
	_args[4] = _p4
	_end++
	_p5 := extractable
	_args[5] = _p5
	_end++
	_p6 := js.Global().Get("Array").New(len(keyUsages))
	for __idx6, __seq_in6 := range keyUsages {
		__seq_out6 := __seq_in6.JSValue()
		_p6.SetIndex(__idx6, __seq_out6)
	}
	_args[6] = _p6
	_end++
	_returned := _this.Value_JS.Call("unwrapKey", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}
