//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package powerpoint ;import (_g "encoding/xml";_d "fmt";_da "github.com/unidoc/unioffice";);type CT_Rel struct{IdAttr *string ;};

// Validate validates the Textdata and its children
func (_cge *Textdata )Validate ()error {return _cge .ValidateWithPath ("\u0054\u0065\u0078\u0074\u0064\u0061\u0074\u0061");};func (_ede *Iscomment )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {_ede .CT_Empty =*NewCT_Empty ();for {_edc ,_daa :=d .Token ();
if _daa !=nil {return _d .Errorf ("p\u0061\u0072\u0073\u0069ng\u0020I\u0073\u0063\u006f\u006d\u006de\u006e\u0074\u003a\u0020\u0025\u0073",_daa );};if _ca ,_fg :=_edc .(_g .EndElement );_fg &&_ca .Name ==start .Name {break ;};};return nil ;};func NewTextdata ()*Textdata {_db :=&Textdata {};
_db .CT_Rel =*NewCT_Rel ();return _db };func (_ag *CT_Rel )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {if _ag .IdAttr !=nil {start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0069\u0064"},Value :_d .Sprintf ("\u0025\u0076",*_ag .IdAttr )});
};e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};type Textdata struct{CT_Rel };

// Validate validates the CT_Empty and its children
func (_ee *CT_Empty )Validate ()error {return _ee .ValidateWithPath ("\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079");};func (_bg *CT_Empty )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for {_c ,_a :=d .Token ();if _a !=nil {return _d .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fE\u006d\u0070\u0074\u0079: \u0025\u0073",_a );
};if _dd ,_eb :=_c .(_g .EndElement );_eb &&_dd .Name ==start .Name {break ;};};return nil ;};type Iscomment struct{CT_Empty };

// ValidateWithPath validates the CT_Empty and its children, prefixing error messages with path
func (_be *CT_Empty )ValidateWithPath (path string )error {return nil };func (_ffa *Textdata )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});
start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061";return _ffa .CT_Rel .MarshalXML (e ,start );};

// ValidateWithPath validates the CT_Rel and its children, prefixing error messages with path
func (_ba *CT_Rel )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the Iscomment and its children, prefixing error messages with path
func (_ce *Iscomment )ValidateWithPath (path string )error {if _cb :=_ce .CT_Empty .ValidateWithPath (path );_cb !=nil {return _cb ;};return nil ;};func (_ff *Iscomment )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});
start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0069s\u0063\u006f\u006d\u006d\u0065\u006et";return _ff .CT_Empty .MarshalXML (e ,start );};func NewCT_Empty ()*CT_Empty {_e :=&CT_Empty {};return _e };func (_cbb *Textdata )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {_cbb .CT_Rel =*NewCT_Rel ();
for _ ,_gb :=range start .Attr {if _gb .Name .Local =="\u0069\u0064"{_bab ,_cf :=_gb .Value ,error (nil );if _cf !=nil {return _cf ;};_cbb .IdAttr =&_bab ;continue ;};};for {_cac ,_fe :=d .Token ();if _fe !=nil {return _d .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0054\u0065\u0078t\u0064\u0061\u0074\u0061: \u0025\u0073",_fe );
};if _eeg ,_ac :=_cac .(_g .EndElement );_ac &&_eeg .Name ==start .Name {break ;};};return nil ;};func (_ddc *CT_Rel )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for _ ,_eg :=range start .Attr {if _eg .Name .Local =="\u0069\u0064"{_gg ,_bgb :=_eg .Value ,error (nil );
if _bgb !=nil {return _bgb ;};_ddc .IdAttr =&_gg ;continue ;};};for {_dg ,_dae :=d .Token ();if _dae !=nil {return _d .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0043T\u005f\u0052e\u006c\u003a\u0020\u0025\u0073",_dae );};if _f ,_ab :=_dg .(_g .EndElement );
_ab &&_f .Name ==start .Name {break ;};};return nil ;};

// Validate validates the CT_Rel and its children
func (_cg *CT_Rel )Validate ()error {return _cg .ValidateWithPath ("\u0043\u0054\u005f\u0052\u0065\u006c");};func NewIscomment ()*Iscomment {_cgb :=&Iscomment {};_cgb .CT_Empty =*NewCT_Empty ();return _cgb };func NewCT_Rel ()*CT_Rel {_gc :=&CT_Rel {};return _gc };


// ValidateWithPath validates the Textdata and its children, prefixing error messages with path
func (_ad *Textdata )ValidateWithPath (path string )error {if _aa :=_ad .CT_Rel .ValidateWithPath (path );_aa !=nil {return _aa ;};return nil ;};type CT_Empty struct{};

// Validate validates the Iscomment and its children
func (_cgg *Iscomment )Validate ()error {return _cgg .ValidateWithPath ("\u0049s\u0063\u006f\u006d\u006d\u0065\u006et");};func (_ed *CT_Empty )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });
return nil ;};func init (){_da .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079",NewCT_Empty );
_da .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0052\u0065\u006c",NewCT_Rel );
_da .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0069s\u0063\u006f\u006d\u006d\u0065\u006et",NewIscomment );
_da .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061",NewTextdata );
};