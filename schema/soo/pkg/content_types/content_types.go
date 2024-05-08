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

package content_types ;import (_fg "encoding/xml";_c "fmt";_fc "github.com/unidoc/unioffice";_fe "github.com/unidoc/unioffice/common/logger";_a "regexp";);const ST_ContentTypePattern ="\u005e\\\u0070{\u004c\u0061\u0074\u0069\u006e\u007d\u002b\u002f\u002e\u002a\u0024";
func (_caa *Default )UnmarshalXML (d *_fg .Decoder ,start _fg .StartElement )error {_caa .CT_Default =*NewCT_Default ();for _ ,_fgd :=range start .Attr {if _fgd .Name .Local =="\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"{_bfe ,_bce :=_fgd .Value ,error (nil );
if _bce !=nil {return _bce ;};_caa .ExtensionAttr =_bfe ;continue ;};if _fgd .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_fea ,_dd :=_fgd .Value ,error (nil );if _dd !=nil {return _dd ;};_caa .ContentTypeAttr =_fea ;continue ;
};};for {_ece ,_feg :=d .Token ();if _feg !=nil {return _c .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u003a\u0020\u0025\u0073",_feg );};if _ea ,_eg :=_ece .(_fg .EndElement );_eg &&_ea .Name ==start .Name {break ;
};};return nil ;};

// ValidateWithPath validates the CT_Default and its children, prefixing error messages with path
func (_agd *CT_Default )ValidateWithPath (path string )error {if !ST_ExtensionPatternRe .MatchString (_agd .ExtensionAttr ){return _c .Errorf ("\u0025s\u002f\u006d.\u0045\u0078\u0074\u0065n\u0073\u0069\u006fn\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074 m\u0061\u0074\u0063h\u0020\u0027%\u0073\u0027\u0020\u0028\u0068\u0061v\u0065\u0020%\u0076\u0029",path ,ST_ExtensionPatternRe ,_agd .ExtensionAttr );
};if !ST_ContentTypePatternRe .MatchString (_agd .ContentTypeAttr ){return _c .Errorf ("\u0025\u0073/\u006d\u002e\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0027\u0025\u0073\u0027\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029",path ,ST_ContentTypePatternRe ,_agd .ContentTypeAttr );
};return nil ;};

// Validate validates the CT_Types and its children
func (_ed *CT_Types )Validate ()error {return _ed .ValidateWithPath ("\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073");};type Default struct{CT_Default };type CT_Types struct{Default []*Default ;Override []*Override ;};type CT_Default struct{ExtensionAttr string ;
ContentTypeAttr string ;};func (_de *Default )MarshalXML (e *_fg .Encoder ,start _fg .StartElement )error {return _de .CT_Default .MarshalXML (e ,start );};var ST_ExtensionPatternRe =_a .MustCompile (ST_ExtensionPattern );

// ValidateWithPath validates the Default and its children, prefixing error messages with path
func (_agf *Default )ValidateWithPath (path string )error {if _egc :=_agf .CT_Default .ValidateWithPath (path );_egc !=nil {return _egc ;};return nil ;};type Override struct{CT_Override };

// Validate validates the CT_Override and its children
func (_eef *CT_Override )Validate ()error {return _eef .ValidateWithPath ("C\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065");};func (_cfe *Override )MarshalXML (e *_fg .Encoder ,start _fg .StartElement )error {return _cfe .CT_Override .MarshalXML (e ,start );
};func NewOverride ()*Override {_gge :=&Override {};_gge .CT_Override =*NewCT_Override ();return _gge };

// Validate validates the Default and its children
func (_acc *Default )Validate ()error {return _acc .ValidateWithPath ("\u0044e\u0066\u0061\u0075\u006c\u0074");};func (_aae *CT_Override )UnmarshalXML (d *_fg .Decoder ,start _fg .StartElement )error {_aae .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";
for _ ,_gac :=range start .Attr {if _gac .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_ee ,_ad :=_gac .Value ,error (nil );if _ad !=nil {return _ad ;};_aae .ContentTypeAttr =_ee ;continue ;};if _gac .Name .Local =="\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"{_cd ,_gc :=_gac .Value ,error (nil );
if _gc !=nil {return _gc ;};_aae .PartNameAttr =_cd ;continue ;};};for {_bdb ,_gab :=d .Token ();if _gab !=nil {return _c .Errorf ("\u0070\u0061\u0072si\u006e\u0067\u0020\u0043\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u003a\u0020\u0025\u0073",_gab );
};if _bc ,_fcg :=_bdb .(_fg .EndElement );_fcg &&_bc .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the CT_Types and its children, prefixing error messages with path
func (_efc *CT_Types )ValidateWithPath (path string )error {for _gcf ,_df :=range _efc .Default {if _gfe :=_df .ValidateWithPath (_c .Sprintf ("\u0025\u0073\u002f\u0044\u0065\u0066\u0061\u0075\u006ct\u005b\u0025\u0064\u005d",path ,_gcf ));_gfe !=nil {return _gfe ;
};};for _aag ,_bda :=range _efc .Override {if _cge :=_bda .ValidateWithPath (_c .Sprintf ("\u0025s\u002fO\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u005b\u0025\u0064\u005d",path ,_aag ));_cge !=nil {return _cge ;};};return nil ;};const ST_ExtensionPattern ="\u0028\u005b\u0021\u0024\u0026\u0027\\\u0028\u005c\u0029\u005c\u002a\\\u002b\u002c\u003a\u003d\u005d\u007c(\u0025\u005b\u0030\u002d\u0039a\u002d\u0066\u0041\u002d\u0046\u005d\u005b\u0030\u002d\u0039\u0061\u002df\u0041\u002d\u0046\u005d\u0029\u007c\u005b\u003a\u0040\u005d\u007c\u005b\u0061\u002d\u007aA\u002d\u005a\u0030\u002d\u0039\u005c\u002d\u005f~\u005d\u0029\u002b";


// ValidateWithPath validates the CT_Override and its children, prefixing error messages with path
func (_ca *CT_Override )ValidateWithPath (path string )error {if !ST_ContentTypePatternRe .MatchString (_ca .ContentTypeAttr ){return _c .Errorf ("\u0025\u0073/\u006d\u002e\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0027\u0025\u0073\u0027\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029",path ,ST_ContentTypePatternRe ,_ca .ContentTypeAttr );
};return nil ;};var ST_ContentTypePatternRe =_a .MustCompile (ST_ContentTypePattern );

// Validate validates the Override and its children
func (_ff *Override )Validate ()error {return _ff .ValidateWithPath ("\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065");};

// Validate validates the Types and its children
func (_eba *Types )Validate ()error {return _eba .ValidateWithPath ("\u0054\u0079\u0070e\u0073")};func (_af *CT_Default )UnmarshalXML (d *_fg .Decoder ,start _fg .StartElement )error {_af .ExtensionAttr ="\u0078\u006d\u006c";_af .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";
for _ ,_gf :=range start .Attr {if _gf .Name .Local =="\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"{_cf ,_ga :=_gf .Value ,error (nil );if _ga !=nil {return _ga ;};_af .ExtensionAttr =_cf ;continue ;};if _gf .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_b ,_ag :=_gf .Value ,error (nil );
if _ag !=nil {return _ag ;};_af .ContentTypeAttr =_b ;continue ;};};for {_gaf ,_aa :=d .Token ();if _aa !=nil {return _c .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u003a\u0020%\u0073",_aa );
};if _fgc ,_fa :=_gaf .(_fg .EndElement );_fa &&_fgc .Name ==start .Name {break ;};};return nil ;};type CT_Override struct{ContentTypeAttr string ;PartNameAttr string ;};

// ValidateWithPath validates the Override and its children, prefixing error messages with path
func (_gfea *Override )ValidateWithPath (path string )error {if _cag :=_gfea .CT_Override .ValidateWithPath (path );_cag !=nil {return _cag ;};return nil ;};func (_eae *Types )UnmarshalXML (d *_fg .Decoder ,start _fg .StartElement )error {_eae .CT_Types =*NewCT_Types ();
_cfa :for {_bfef ,_bg :=d .Token ();if _bg !=nil {return _bg ;};switch _bbf :=_bfef .(type ){case _fg .StartElement :switch _bbf .Name {case _fg .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u0044e\u0066\u0061\u0075\u006c\u0074"}:_agg :=NewDefault ();
if _gcg :=d .DecodeElement (_agg ,&_bbf );_gcg !=nil {return _gcg ;};_eae .Default =append (_eae .Default ,_agg );case _fg .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}:_ba :=NewOverride ();
if _bac :=d .DecodeElement (_ba ,&_bbf );_bac !=nil {return _bac ;};_eae .Override =append (_eae .Override ,_ba );default:_fe .Log .Debug ("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn \u0054\u0079\u0070e\u0073 \u0025\u0076",_bbf .Name );
if _ddf :=d .Skip ();_ddf !=nil {return _ddf ;};};case _fg .EndElement :break _cfa ;case _fg .CharData :};};return nil ;};func NewCT_Default ()*CT_Default {_g :=&CT_Default {};_g .ExtensionAttr ="\u0078\u006d\u006c";_g .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";
return _g ;};func (_db *CT_Types )UnmarshalXML (d *_fg .Decoder ,start _fg .StartElement )error {_cg :for {_ef ,_bdd :=d .Token ();if _bdd !=nil {return _bdd ;};switch _gca :=_ef .(type ){case _fg .StartElement :switch _gca .Name {case _fg .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u0044e\u0066\u0061\u0075\u006c\u0074"}:_ac :=NewDefault ();
if _cab :=d .DecodeElement (_ac ,&_gca );_cab !=nil {return _cab ;};_db .Default =append (_db .Default ,_ac );case _fg .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}:_eca :=NewOverride ();
if _fgeb :=d .DecodeElement (_eca ,&_gca );_fgeb !=nil {return _fgeb ;};_db .Override =append (_db .Override ,_eca );default:_fe .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073\u0020\u0025\u0076",_gca .Name );
if _agb :=d .Skip ();_agb !=nil {return _agb ;};};case _fg .EndElement :break _cg ;case _fg .CharData :};};return nil ;};func NewTypes ()*Types {_aac :=&Types {};_aac .CT_Types =*NewCT_Types ();return _aac };func NewCT_Types ()*CT_Types {_d :=&CT_Types {};
return _d };func (_agfe *Types )MarshalXML (e *_fg .Encoder ,start _fg .StartElement )error {start .Attr =append (start .Attr ,_fg .Attr {Name :_fg .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s"});
start .Attr =append (start .Attr ,_fg .Attr {Name :_fg .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0054\u0079\u0070e\u0073";return _agfe .CT_Types .MarshalXML (e ,start );};

// ValidateWithPath validates the Types and its children, prefixing error messages with path
func (_ggd *Types )ValidateWithPath (path string )error {if _dea :=_ggd .CT_Types .ValidateWithPath (path );_dea !=nil {return _dea ;};return nil ;};func (_fee *CT_Override )MarshalXML (e *_fg .Encoder ,start _fg .StartElement )error {start .Attr =append (start .Attr ,_fg .Attr {Name :_fg .Name {Local :"C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"},Value :_c .Sprintf ("\u0025\u0076",_fee .ContentTypeAttr )});
start .Attr =append (start .Attr ,_fg .Attr {Name :_fg .Name {Local :"\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"},Value :_c .Sprintf ("\u0025\u0076",_fee .PartNameAttr )});e .EncodeToken (start );e .EncodeToken (_fg .EndElement {Name :start .Name });
return nil ;};func (_ec *CT_Types )MarshalXML (e *_fg .Encoder ,start _fg .StartElement )error {e .EncodeToken (start );if _ec .Default !=nil {_bf :=_fg .StartElement {Name :_fg .Name {Local :"\u0044e\u0066\u0061\u0075\u006c\u0074"}};for _ ,_ggf :=range _ec .Default {e .EncodeElement (_ggf ,_bf );
};};if _ec .Override !=nil {_dg :=_fg .StartElement {Name :_fg .Name {Local :"\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}};for _ ,_fge :=range _ec .Override {e .EncodeElement (_fge ,_dg );};};e .EncodeToken (_fg .EndElement {Name :start .Name });
return nil ;};func NewDefault ()*Default {_ae :=&Default {};_ae .CT_Default =*NewCT_Default ();return _ae };

// Validate validates the CT_Default and its children
func (_bb *CT_Default )Validate ()error {return _bb .ValidateWithPath ("\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074");};func (_gg *CT_Default )MarshalXML (e *_fg .Encoder ,start _fg .StartElement )error {start .Attr =append (start .Attr ,_fg .Attr {Name :_fg .Name {Local :"\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"},Value :_c .Sprintf ("\u0025\u0076",_gg .ExtensionAttr )});
start .Attr =append (start .Attr ,_fg .Attr {Name :_fg .Name {Local :"C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"},Value :_c .Sprintf ("\u0025\u0076",_gg .ContentTypeAttr )});e .EncodeToken (start );e .EncodeToken (_fg .EndElement {Name :start .Name });
return nil ;};func NewCT_Override ()*CT_Override {_bd :=&CT_Override {};_bd .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";return _bd ;};func (_fb *Override )UnmarshalXML (d *_fg .Decoder ,start _fg .StartElement )error {_fb .CT_Override =*NewCT_Override ();
for _ ,_ab :=range start .Attr {if _ab .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_def ,_cfg :=_ab .Value ,error (nil );if _cfg !=nil {return _cfg ;};_fb .ContentTypeAttr =_def ;continue ;};if _ab .Name .Local =="\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"{_aee ,_accb :=_ab .Value ,error (nil );
if _accb !=nil {return _accb ;};_fb .PartNameAttr =_aee ;continue ;};};for {_ade ,_fdb :=d .Token ();if _fdb !=nil {return _c .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u004f\u0076\u0065r\u0072\u0069\u0064\u0065: \u0025\u0073",_fdb );};if _cae ,_dff :=_ade .(_fg .EndElement );
_dff &&_cae .Name ==start .Name {break ;};};return nil ;};type Types struct{CT_Types };func init (){_fc .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073",NewCT_Types );
_fc .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074",NewCT_Default );
_fc .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","C\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065",NewCT_Override );
_fc .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0054\u0079\u0070e\u0073",NewTypes );
_fc .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0044e\u0066\u0061\u0075\u006c\u0074",NewDefault );
_fc .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065",NewOverride );
};