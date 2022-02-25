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

package elements ;import (_d "encoding/xml";_da "fmt";_f "github.com/unidoc/unioffice";_ed "github.com/unidoc/unioffice/common/logger";);func NewAny ()*Any {_g :=&Any {};_g .SimpleLiteral =*NewSimpleLiteral ();return _g };func (_fc *SimpleLiteral )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {e .EncodeToken (start );e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the Any and its children, prefixing error messages with path
func (_ac *Any )ValidateWithPath (path string )error {if _c :=_ac .SimpleLiteral .ValidateWithPath (path );_c !=nil {return _c ;};return nil ;};func (_gbe *ElementsGroupChoice )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {if _gbe .Any !=nil {_dc :=_d .StartElement {Name :_d .Name {Local :"\u0064\u0063\u003a\u0061\u006e\u0079"}};for _ ,_dgf :=range _gbe .Any {e .EncodeElement (_dgf ,_dc );};};return nil ;};func (_df *ElementsGroup )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {if _df .Choice !=nil {for _ ,_ad :=range _df .Choice {_ad .MarshalXML (e ,_d .StartElement {});};};return nil ;};type ElementsGroupChoice struct{Any []*Any ;};type SimpleLiteral struct{};

// ValidateWithPath validates the ElementContainer and its children, prefixing error messages with path
func (_ab *ElementContainer )ValidateWithPath (path string )error {for _eg ,_ba :=range _ab .Choice {if _bac :=_ba .ValidateWithPath (_da .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_eg ));_bac !=nil {return _bac ;};};return nil ;};func (_ge *ElementsGroup )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_cf :for {_fge ,_ff :=d .Token ();if _ff !=nil {return _ff ;};switch _ec :=_fge .(type ){case _d .StartElement :switch _ec .Name {case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_cfe :=NewElementsGroupChoice ();if _gae :=d .DecodeElement (&_cfe .Any ,&_ec );_gae !=nil {return _gae ;};_ge .Choice =append (_ge .Choice ,_cfe );default:_ed .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006de\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070 \u0025\u0076",_ec .Name );if _ef :=d .Skip ();_ef !=nil {return _ef ;};};case _d .EndElement :break _cf ;case _d .CharData :};};return nil ;};type ElementsGroup struct{Choice []*ElementsGroupChoice ;};

// Validate validates the ElementsGroup and its children
func (_cgf *ElementsGroup )Validate ()error {return _cgf .ValidateWithPath ("\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070");};

// Validate validates the Any and its children
func (_ag *Any )Validate ()error {return _ag .ValidateWithPath ("\u0041\u006e\u0079")};type ElementContainer struct{Choice []*ElementsGroupChoice ;};

// Validate validates the ElementContainer and its children
func (_cd *ElementContainer )Validate ()error {return _cd .ValidateWithPath ("\u0045\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072");};func (_gf *Any )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_gf .SimpleLiteral =*NewSimpleLiteral ();for {_ga ,_a :=d .Token ();if _a !=nil {return _da .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0041\u006e\u0079\u003a\u0020\u0025\u0073",_a );};if _gb ,_fg :=_ga .(_d .EndElement );_fg &&_gb .Name ==start .Name {break ;};};return nil ;};func (_dae *ElementContainer )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_fe :for {_dg ,_de :=d .Token ();if _de !=nil {return _de ;};switch _bb :=_dg .(type ){case _d .StartElement :switch _bb .Name {case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_ca :=NewElementsGroupChoice ();if _gfg :=d .DecodeElement (&_ca .Any ,&_bb );_gfg !=nil {return _gfg ;};_dae .Choice =append (_dae .Choice ,_ca );default:_ed .Log .Debug ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0025v",_bb .Name );if _fd :=d .Skip ();_fd !=nil {return _fd ;};};case _d .EndElement :break _fe ;case _d .CharData :};};return nil ;};func NewSimpleLiteral ()*SimpleLiteral {_eb :=&SimpleLiteral {};return _eb };

// ValidateWithPath validates the ElementsGroupChoice and its children, prefixing error messages with path
func (_af *ElementsGroupChoice )ValidateWithPath (path string )error {for _beg ,_dd :=range _af .Any {if _ea :=_dd .ValidateWithPath (_da .Sprintf ("\u0025\u0073\u002f\u0041\u006e\u0079\u005b\u0025\u0064\u005d",path ,_beg ));_ea !=nil {return _ea ;};};return nil ;};func NewElementsGroup ()*ElementsGroup {_feb :=&ElementsGroup {};return _feb };

// Validate validates the ElementsGroupChoice and its children
func (_gbf *ElementsGroupChoice )Validate ()error {return _gbf .ValidateWithPath ("\u0045\u006c\u0065\u006den\u0074\u0073\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u006f\u0069\u0063\u0065");};type Any struct{SimpleLiteral };

// ValidateWithPath validates the ElementsGroup and its children, prefixing error messages with path
func (_cdf *ElementsGroup )ValidateWithPath (path string )error {for _adf ,_be :=range _cdf .Choice {if _gea :=_be .ValidateWithPath (_da .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_adf ));_gea !=nil {return _gea ;};};return nil ;};func (_bd *ElementsGroupChoice )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_edc :for {_cgd ,_fa :=d .Token ();if _fa !=nil {return _fa ;};switch _aga :=_cgd .(type ){case _d .StartElement :switch _aga .Name {case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_dgd :=NewAny ();if _aa :=d .DecodeElement (_dgd ,&_aga );_aa !=nil {return _aa ;};_bd .Any =append (_bd .Any ,_dgd );default:_ed .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072ou\u0070\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076",_aga .Name );if _fdb :=d .Skip ();_fdb !=nil {return _fdb ;};};case _d .EndElement :break _edc ;case _d .CharData :};};return nil ;};func (_eag *SimpleLiteral )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for {_gg ,_efb :=d .Token ();if _efb !=nil {return _da .Errorf ("\u0070a\u0072\u0073\u0069\u006eg\u0020\u0053\u0069\u006d\u0070l\u0065L\u0069t\u0065\u0072\u0061\u006c\u003a\u0020\u0025s",_efb );};if _bed ,_egd :=_gg .(_d .EndElement );_egd &&_bed .Name ==start .Name {break ;};};return nil ;};func NewElementContainer ()*ElementContainer {_eda :=&ElementContainer {};return _eda };func (_agb *ElementContainer )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Name .Local ="\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072";e .EncodeToken (start );if _agb .Choice !=nil {for _ ,_gbg :=range _agb .Choice {_gbg .MarshalXML (e ,_d .StartElement {});};};e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func (_b *Any )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {return _b .SimpleLiteral .MarshalXML (e ,start );};

// Validate validates the SimpleLiteral and its children
func (_gfd *SimpleLiteral )Validate ()error {return _gfd .ValidateWithPath ("\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c");};

// ValidateWithPath validates the SimpleLiteral and its children, prefixing error messages with path
func (_geb *SimpleLiteral )ValidateWithPath (path string )error {return nil };func NewElementsGroupChoice ()*ElementsGroupChoice {_abc :=&ElementsGroupChoice {};return _abc };func init (){_f .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c",NewSimpleLiteral );_f .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072",NewElementContainer );_f .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0061\u006e\u0079",NewAny );_f .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070",NewElementsGroup );};