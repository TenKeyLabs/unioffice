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

package powerpoint ;import (_e "encoding/xml";_g "fmt";_d "github.com/unidoc/unioffice";);

// Validate validates the CT_Rel and its children
func (_cbe *CT_Rel )Validate ()error {return _cbe .ValidateWithPath ("\u0043\u0054\u005f\u0052\u0065\u006c");};func (_gff *CT_Rel )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {if _gff .IdAttr !=nil {start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0069\u0064"},Value :_g .Sprintf ("\u0025\u0076",*_gff .IdAttr )});
};e .EncodeToken (start );e .EncodeToken (_e .EndElement {Name :start .Name });return nil ;};func (_cf *CT_Rel )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {for _ ,_b :=range start .Attr {if _b .Name .Local =="\u0069\u0064"{_ge ,_be :=_b .Value ,error (nil );
if _be !=nil {return _be ;};_cf .IdAttr =&_ge ;continue ;};};for {_gg ,_ag :=d .Token ();if _ag !=nil {return _g .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0043T\u005f\u0052e\u006c\u003a\u0020\u0025\u0073",_ag );};if _gbb ,_fc :=_gg .(_e .EndElement );
_fc &&_gbb .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the Iscomment and its children, prefixing error messages with path
func (_dd *Iscomment )ValidateWithPath (path string )error {if _edd :=_dd .CT_Empty .ValidateWithPath (path );_edd !=nil {return _edd ;};return nil ;};type Iscomment struct{CT_Empty };func NewCT_Rel ()*CT_Rel {_fe :=&CT_Rel {};return _fe };func NewIscomment ()*Iscomment {_ef :=&Iscomment {};
_ef .CT_Empty =*NewCT_Empty ();return _ef };func (_ec *Iscomment )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});
start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0069s\u0063\u006f\u006d\u006d\u0065\u006et";return _ec .CT_Empty .MarshalXML (e ,start );};type CT_Empty struct{};

// Validate validates the Iscomment and its children
func (_ecd *Iscomment )Validate ()error {return _ecd .ValidateWithPath ("\u0049s\u0063\u006f\u006d\u006d\u0065\u006et");};

// ValidateWithPath validates the Textdata and its children, prefixing error messages with path
func (_cd *Textdata )ValidateWithPath (path string )error {if _cba :=_cd .CT_Rel .ValidateWithPath (path );_cba !=nil {return _cba ;};return nil ;};

// Validate validates the Textdata and its children
func (_dfe *Textdata )Validate ()error {return _dfe .ValidateWithPath ("\u0054\u0065\u0078\u0074\u0064\u0061\u0074\u0061");};func (_ff *Iscomment )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_ff .CT_Empty =*NewCT_Empty ();for {_cg ,_cc :=d .Token ();
if _cc !=nil {return _g .Errorf ("p\u0061\u0072\u0073\u0069ng\u0020I\u0073\u0063\u006f\u006d\u006de\u006e\u0074\u003a\u0020\u0025\u0073",_cc );};if _ed ,_fef :=_cg .(_e .EndElement );_fef &&_ed .Name ==start .Name {break ;};};return nil ;};func NewTextdata ()*Textdata {_fa :=&Textdata {};
_fa .CT_Rel =*NewCT_Rel ();return _fa };type CT_Rel struct{IdAttr *string ;};

// Validate validates the CT_Empty and its children
func (_df *CT_Empty )Validate ()error {return _df .ValidateWithPath ("\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079");};func (_af *CT_Empty )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {e .EncodeToken (start );e .EncodeToken (_e .EndElement {Name :start .Name });
return nil ;};func NewCT_Empty ()*CT_Empty {_a :=&CT_Empty {};return _a };type Textdata struct{CT_Rel };

// ValidateWithPath validates the CT_Rel and its children, prefixing error messages with path
func (_eg *CT_Rel )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the CT_Empty and its children, prefixing error messages with path
func (_f *CT_Empty )ValidateWithPath (path string )error {return nil };func (_gf *CT_Empty )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {for {_ea ,_gc :=d .Token ();if _gc !=nil {return _g .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fE\u006d\u0070\u0074\u0079: \u0025\u0073",_gc );
};if _gb ,_ee :=_ea .(_e .EndElement );_ee &&_gb .Name ==start .Name {break ;};};return nil ;};func (_de *Textdata )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});
start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061";return _de .CT_Rel .MarshalXML (e ,start );};func (_ffg *Textdata )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_ffg .CT_Rel =*NewCT_Rel ();for _ ,_edf :=range start .Attr {if _edf .Name .Local =="\u0069\u0064"{_def ,_efa :=_edf .Value ,error (nil );
if _efa !=nil {return _efa ;};_ffg .IdAttr =&_def ;continue ;};};for {_bee ,_bf :=d .Token ();if _bf !=nil {return _g .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0054\u0065\u0078t\u0064\u0061\u0074\u0061: \u0025\u0073",_bf );};if _gba ,_dg :=_bee .(_e .EndElement );
_dg &&_gba .Name ==start .Name {break ;};};return nil ;};func init (){_d .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079",NewCT_Empty );
_d .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0052\u0065\u006c",NewCT_Rel );
_d .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0069s\u0063\u006f\u006d\u006d\u0065\u006et",NewIscomment );
_d .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061",NewTextdata );
};