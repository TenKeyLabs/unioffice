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

package activeX ;import (_b "encoding/xml";_bf "fmt";_ba "github.com/unidoc/unioffice";_f "github.com/unidoc/unioffice/common/logger";);func (_cfd ST_Persistence )String ()string {switch _cfd {case 0:return "";case 1:return "\u0070e\u0072s\u0069\u0073\u0074\u0050\u0072o\u0070\u0065r\u0074\u0079\u0042\u0061\u0067";
case 2:return "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061\u006d";case 3:return "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061m\u0049\u006e\u0069\u0074";case 4:return "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074o\u0072\u0061\u0067\u0065";
};return "";};type CT_Ocx struct{ClassidAttr string ;LicenseAttr *string ;IdAttr *string ;PersistenceAttr ST_Persistence ;OcxPr []*CT_OcxPr ;};func NewCT_Font ()*CT_Font {_fg :=&CT_Font {};return _fg };

// ValidateWithPath validates the CT_OcxPrChoice and its children, prefixing error messages with path
func (_bcf *CT_OcxPrChoice )ValidateWithPath (path string )error {if _bcf .Font !=nil {if _fabd :=_bcf .Font .ValidateWithPath (path +"\u002f\u0046\u006fn\u0074");_fabd !=nil {return _fabd ;};};if _bcf .Picture !=nil {if _eab :=_bcf .Picture .ValidateWithPath (path +"\u002f\u0050\u0069\u0063\u0074\u0075\u0072\u0065");
_eab !=nil {return _eab ;};};return nil ;};type CT_OcxPrChoice struct{Font *CT_Font ;Picture *CT_Picture ;};func (_d *CT_Font )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {if _d .PersistenceAttr !=ST_PersistenceUnset {_g ,_gd :=_d .PersistenceAttr .MarshalXMLAttr (_b .Name {Local :"\u0061\u0078\u003a\u0070\u0065\u0072\u0073\u0069\u0073t\u0065\u006e\u0063\u0065"});
if _gd !=nil {return _gd ;};start .Attr =append (start .Attr ,_g );};if _d .IdAttr !=nil {start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0072\u003a\u0069\u0064"},Value :_bf .Sprintf ("\u0025\u0076",*_d .IdAttr )});};e .EncodeToken (start );
if _d .OcxPr !=nil {_eb :=_b .StartElement {Name :_b .Name {Local :"\u0061\u0078\u003a\u006f\u0063\u0078\u0050\u0072"}};for _ ,_gg :=range _d .OcxPr {e .EncodeElement (_gg ,_eb );};};e .EncodeToken (_b .EndElement {Name :start .Name });return nil ;};func (_ff *CT_Font )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for _ ,_de :=range start .Attr {if _de .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"&&_de .Name .Local =="\u0069\u0064"||_de .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073"&&_de .Name .Local =="\u0069\u0064"{_c ,_bc :=_de .Value ,error (nil );
if _bc !=nil {return _bc ;};_ff .IdAttr =&_c ;continue ;};if _de .Name .Local =="p\u0065\u0072\u0073\u0069\u0073\u0074\u0065\u006e\u0063\u0065"{_ff .PersistenceAttr .UnmarshalXMLAttr (_de );continue ;};};_dc :for {_df ,_ee :=d .Token ();if _ee !=nil {return _ee ;
};switch _cf :=_df .(type ){case _b .StartElement :switch _cf .Name {case _b .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u006f\u0063\u0078P\u0072"}:_gf :=NewCT_OcxPr ();
if _ag :=d .DecodeElement (_gf ,&_cf );_ag !=nil {return _ag ;};_ff .OcxPr =append (_ff .OcxPr ,_gf );default:_f .Log .Debug ("\u0073\u006b\u0069p\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u0046\u006f\u006e\u0074\u0020\u0025\u0076",_cf .Name );
if _gga :=d .Skip ();_gga !=nil {return _gga ;};};case _b .EndElement :break _dc ;case _b .CharData :};};return nil ;};

// ValidateWithPath validates the CT_Picture and its children, prefixing error messages with path
func (_cea *CT_Picture )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the CT_Font and its children, prefixing error messages with path
func (_bfg *CT_Font )ValidateWithPath (path string )error {if _bg :=_bfg .PersistenceAttr .ValidateWithPath (path +"\u002f\u0050e\u0072\u0073\u0069s\u0074\u0065\u006e\u0063\u0065\u0041\u0074\u0074\u0072");_bg !=nil {return _bg ;};for _bb ,_fd :=range _bfg .OcxPr {if _eg :=_fd .ValidateWithPath (_bf .Sprintf ("\u0025\u0073\u002fO\u0063\u0078\u0050\u0072\u005b\u0025\u0064\u005d",path ,_bb ));
_eg !=nil {return _eg ;};};return nil ;};func (_gcg *ST_Persistence )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_add ,_bae :=d .Token ();if _bae !=nil {return _bae ;};if _gba ,_bd :=_add .(_b .EndElement );_bd &&_gba .Name ==start .Name {*_gcg =1;
return nil ;};if _adc ,_dcd :=_add .(_b .CharData );!_dcd {return _bf .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_add );}else {switch string (_adc ){case "":*_gcg =0;
case "\u0070e\u0072s\u0069\u0073\u0074\u0050\u0072o\u0070\u0065r\u0074\u0079\u0042\u0061\u0067":*_gcg =1;case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061\u006d":*_gcg =2;case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061m\u0049\u006e\u0069\u0074":*_gcg =3;
case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074o\u0072\u0061\u0067\u0065":*_gcg =4;};};_add ,_bae =d .Token ();if _bae !=nil {return _bae ;};if _dgf ,_fde :=_add .(_b .EndElement );_fde &&_dgf .Name ==start .Name {return nil ;};return _bf .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_add );
};func (_dce *CT_OcxPrChoice )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {if _dce .Font !=nil {_bgb :=_b .StartElement {Name :_b .Name {Local :"\u0061x\u003a\u0066\u006f\u006e\u0074"}};e .EncodeElement (_dce .Font ,_bgb );};if _dce .Picture !=nil {_gbe :=_b .StartElement {Name :_b .Name {Local :"\u0061\u0078\u003a\u0070\u0069\u0063\u0074\u0075\u0072\u0065"}};
e .EncodeElement (_dce .Picture ,_gbe );};return nil ;};func NewCT_Ocx ()*CT_Ocx {_cb :=&CT_Ocx {};_cb .PersistenceAttr =ST_Persistence (1);return _cb };func (_fdee ST_Persistence )ValidateWithPath (path string )error {switch _fdee {case 0,1,2,3,4:default:return _bf .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_fdee ));
};return nil ;};func (_bagd *Ocx )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058"});
start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0061\u0078"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058"});
start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0072"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"});
start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0061\u0078\u003a\u006f\u0063\u0078";return _bagd .CT_Ocx .MarshalXML (e ,start );};type CT_OcxPr struct{NameAttr string ;ValueAttr *string ;Choice *CT_OcxPrChoice ;};

// Validate validates the CT_Ocx and its children
func (_fab *CT_Ocx )Validate ()error {return _fab .ValidateWithPath ("\u0043\u0054\u005f\u004f\u0063\u0078");};func (_cfbe ST_Persistence )MarshalXMLAttr (name _b .Name )(_b .Attr ,error ){_gfdb :=_b .Attr {};_gfdb .Name =name ;switch _cfbe {case ST_PersistenceUnset :_gfdb .Value ="";
case ST_PersistencePersistPropertyBag :_gfdb .Value ="\u0070e\u0072s\u0069\u0073\u0074\u0050\u0072o\u0070\u0065r\u0074\u0079\u0042\u0061\u0067";case ST_PersistencePersistStream :_gfdb .Value ="\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061\u006d";
case ST_PersistencePersistStreamInit :_gfdb .Value ="\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061m\u0049\u006e\u0069\u0074";case ST_PersistencePersistStorage :_gfdb .Value ="\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074o\u0072\u0061\u0067\u0065";
};return _gfdb ,nil ;};type CT_Font struct{PersistenceAttr ST_Persistence ;IdAttr *string ;OcxPr []*CT_OcxPr ;};func (_ddd *CT_OcxPr )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0061x\u003a\u006e\u0061\u006d\u0065"},Value :_bf .Sprintf ("\u0025\u0076",_ddd .NameAttr )});
if _ddd .ValueAttr !=nil {start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0061\u0078\u003a\u0076\u0061\u006c\u0075\u0065"},Value :_bf .Sprintf ("\u0025\u0076",*_ddd .ValueAttr )});};e .EncodeToken (start );if _ddd .Choice !=nil {_ddd .Choice .MarshalXML (e ,_b .StartElement {});
};e .EncodeToken (_b .EndElement {Name :start .Name });return nil ;};func NewCT_OcxPr ()*CT_OcxPr {_bgd :=&CT_OcxPr {};return _bgd };

// Validate validates the CT_OcxPrChoice and its children
func (_fdf *CT_OcxPrChoice )Validate ()error {return _fdf .ValidateWithPath ("\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072\u0043h\u006f\u0069\u0063\u0065");};func NewCT_Picture ()*CT_Picture {_ebf :=&CT_Picture {};return _ebf };func (_bdd ST_Persistence )Validate ()error {return _bdd .ValidateWithPath ("")};
func NewOcx ()*Ocx {_fga :=&Ocx {};_fga .CT_Ocx =*NewCT_Ocx ();return _fga };type Ocx struct{CT_Ocx };func (_fb *CT_OcxPrChoice )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_bce :for {_fda ,_cc :=d .Token ();if _cc !=nil {return _cc ;};
switch _bfc :=_fda .(type ){case _b .StartElement :switch _bfc .Name {case _b .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u0066\u006f\u006e\u0074"}:_fb .Font =NewCT_Font ();
if _ffe :=d .DecodeElement (_fb .Font ,&_bfc );_ffe !=nil {return _ffe ;};case _b .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u0070i\u0063\u0074\u0075\u0072\u0065"}:_fb .Picture =NewCT_Picture ();
if _age :=d .DecodeElement (_fb .Picture ,&_bfc );_age !=nil {return _age ;};default:_f .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072\u0043\u0068o\u0069c\u0065\u0020\u0025\u0076",_bfc .Name );
if _aga :=d .Skip ();_aga !=nil {return _aga ;};};case _b .EndElement :break _bce ;case _b .CharData :};};return nil ;};

// Validate validates the Ocx and its children
func (_gdce *Ocx )Validate ()error {return _gdce .ValidateWithPath ("\u004f\u0063\u0078")};func (_ge *CT_Ocx )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0061\u0078\u003a\u0063\u006c\u0061\u0073\u0073\u0069\u0064"},Value :_bf .Sprintf ("\u0025\u0076",_ge .ClassidAttr )});
if _ge .LicenseAttr !=nil {start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0061\u0078\u003a\u006c\u0069\u0063\u0065\u006e\u0073\u0065"},Value :_bf .Sprintf ("\u0025\u0076",*_ge .LicenseAttr )});};if _ge .IdAttr !=nil {start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0072\u003a\u0069\u0064"},Value :_bf .Sprintf ("\u0025\u0076",*_ge .IdAttr )});
};_ac ,_agf :=_ge .PersistenceAttr .MarshalXMLAttr (_b .Name {Local :"\u0061\u0078\u003a\u0070\u0065\u0072\u0073\u0069\u0073t\u0065\u006e\u0063\u0065"});if _agf !=nil {return _agf ;};start .Attr =append (start .Attr ,_ac );e .EncodeToken (start );if _ge .OcxPr !=nil {_ffg :=_b .StartElement {Name :_b .Name {Local :"\u0061\u0078\u003a\u006f\u0063\u0078\u0050\u0072"}};
for _ ,_bag :=range _ge .OcxPr {e .EncodeElement (_bag ,_ffg );};};e .EncodeToken (_b .EndElement {Name :start .Name });return nil ;};

// Validate validates the CT_OcxPr and its children
func (_acc *CT_OcxPr )Validate ()error {return _acc .ValidateWithPath ("\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072");};func (_ged *CT_Ocx )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_ged .PersistenceAttr =ST_Persistence (1);for _ ,_faf :=range start .Attr {if _faf .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"&&_faf .Name .Local =="\u0069\u0064"||_faf .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073"&&_faf .Name .Local =="\u0069\u0064"{_dd ,_ad :=_faf .Value ,error (nil );
if _ad !=nil {return _ad ;};_ged .IdAttr =&_dd ;continue ;};if _faf .Name .Local =="\u0063l\u0061\u0073\u0073\u0069\u0064"{_acd ,_gde :=_faf .Value ,error (nil );if _gde !=nil {return _gde ;};_ged .ClassidAttr =_acd ;continue ;};if _faf .Name .Local =="\u006ci\u0063\u0065\u006e\u0073\u0065"{_cfg ,_cff :=_faf .Value ,error (nil );
if _cff !=nil {return _cff ;};_ged .LicenseAttr =&_cfg ;continue ;};if _faf .Name .Local =="p\u0065\u0072\u0073\u0069\u0073\u0074\u0065\u006e\u0063\u0065"{_ged .PersistenceAttr .UnmarshalXMLAttr (_faf );continue ;};};_cg :for {_ga ,_fgg :=d .Token ();if _fgg !=nil {return _fgg ;
};switch _dag :=_ga .(type ){case _b .StartElement :switch _dag .Name {case _b .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u006f\u0063\u0078P\u0072"}:_ggb :=NewCT_OcxPr ();
if _ae :=d .DecodeElement (_ggb ,&_dag );_ae !=nil {return _ae ;};_ged .OcxPr =append (_ged .OcxPr ,_ggb );default:_f .Log .Debug ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u0063\u0078\u0020\u0025\u0076",_dag .Name );
if _ca :=d .Skip ();_ca !=nil {return _ca ;};};case _b .EndElement :break _cg ;case _b .CharData :};};return nil ;};type ST_Persistence byte ;type CT_Picture struct{IdAttr *string ;};func (_be *CT_OcxPr )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for _ ,_aa :=range start .Attr {if _aa .Name .Local =="\u006e\u0061\u006d\u0065"{_cd ,_dg :=_aa .Value ,error (nil );
if _dg !=nil {return _dg ;};_be .NameAttr =_cd ;continue ;};if _aa .Name .Local =="\u0076\u0061\u006cu\u0065"{_cdg ,_eag :=_aa .Value ,error (nil );if _eag !=nil {return _eag ;};_be .ValueAttr =&_cdg ;continue ;};};_dab :for {_ed ,_fgd :=d .Token ();if _fgd !=nil {return _fgd ;
};switch _fed :=_ed .(type ){case _b .StartElement :switch _fed .Name {case _b .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u0066\u006f\u006e\u0074"}:_be .Choice =NewCT_OcxPrChoice ();
if _egd :=d .DecodeElement (&_be .Choice .Font ,&_fed );_egd !=nil {return _egd ;};case _b .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u0070i\u0063\u0074\u0075\u0072\u0065"}:_be .Choice =NewCT_OcxPrChoice ();
if _ggf :=d .DecodeElement (&_be .Choice .Picture ,&_fed );_ggf !=nil {return _ggf ;};default:_f .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072\u0020\u0025\u0076",_fed .Name );
if _cga :=d .Skip ();_cga !=nil {return _cga ;};};case _b .EndElement :break _dab ;case _b .CharData :};};return nil ;};func NewCT_OcxPrChoice ()*CT_OcxPrChoice {_fge :=&CT_OcxPrChoice {};return _fge };const (ST_PersistenceUnset ST_Persistence =0;ST_PersistencePersistPropertyBag ST_Persistence =1;
ST_PersistencePersistStream ST_Persistence =2;ST_PersistencePersistStreamInit ST_Persistence =3;ST_PersistencePersistStorage ST_Persistence =4;);

// Validate validates the CT_Picture and its children
func (_gcb *CT_Picture )Validate ()error {return _gcb .ValidateWithPath ("\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065");};

// ValidateWithPath validates the CT_OcxPr and its children, prefixing error messages with path
func (_gc *CT_OcxPr )ValidateWithPath (path string )error {if _gc .Choice !=nil {if _daf :=_gc .Choice .ValidateWithPath (path +"\u002fC\u0068\u006f\u0069\u0063\u0065");_daf !=nil {return _daf ;};};return nil ;};func (_bge *ST_Persistence )UnmarshalXMLAttr (attr _b .Attr )error {switch attr .Value {case "":*_bge =0;
case "\u0070e\u0072s\u0069\u0073\u0074\u0050\u0072o\u0070\u0065r\u0074\u0079\u0042\u0061\u0067":*_bge =1;case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061\u006d":*_bge =2;case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061m\u0049\u006e\u0069\u0074":*_bge =3;
case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074o\u0072\u0061\u0067\u0065":*_bge =4;};return nil ;};func (_cae ST_Persistence )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {return e .EncodeElement (_cae .String (),start );};

// Validate validates the CT_Font and its children
func (_fe *CT_Font )Validate ()error {return _fe .ValidateWithPath ("\u0043T\u005f\u0046\u006f\u006e\u0074");};

// ValidateWithPath validates the CT_Ocx and its children, prefixing error messages with path
func (_eed *CT_Ocx )ValidateWithPath (path string )error {if _eed .PersistenceAttr ==ST_PersistenceUnset {return _bf .Errorf ("\u0025\u0073\u002f\u0050\u0065\u0072\u0073\u0069\u0073\u0074\u0065\u006e\u0063e\u0041\u0074\u0074\u0072\u0020\u0069s\u0020\u0061\u0020\u006d\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020f\u0069\u0065\u006c\u0064",path );
};if _geb :=_eed .PersistenceAttr .ValidateWithPath (path +"\u002f\u0050e\u0072\u0073\u0069s\u0074\u0065\u006e\u0063\u0065\u0041\u0074\u0074\u0072");_geb !=nil {return _geb ;};for _db ,_gb :=range _eed .OcxPr {if _dcb :=_gb .ValidateWithPath (_bf .Sprintf ("\u0025\u0073\u002fO\u0063\u0078\u0050\u0072\u005b\u0025\u0064\u005d",path ,_db ));
_dcb !=nil {return _dcb ;};};return nil ;};func (_eae *Ocx )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_eae .CT_Ocx =*NewCT_Ocx ();for _ ,_dgc :=range start .Attr {if _dgc .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"&&_dgc .Name .Local =="\u0069\u0064"||_dgc .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073"&&_dgc .Name .Local =="\u0069\u0064"{_ada ,_eac :=_dgc .Value ,error (nil );
if _eac !=nil {return _eac ;};_eae .IdAttr =&_ada ;continue ;};if _dgc .Name .Local =="\u0063l\u0061\u0073\u0073\u0069\u0064"{_gef ,_ccab :=_dgc .Value ,error (nil );if _ccab !=nil {return _ccab ;};_eae .ClassidAttr =_gef ;continue ;};if _dgc .Name .Local =="\u006ci\u0063\u0065\u006e\u0073\u0065"{_bfb ,_gdc :=_dgc .Value ,error (nil );
if _gdc !=nil {return _gdc ;};_eae .LicenseAttr =&_bfb ;continue ;};if _dgc .Name .Local =="p\u0065\u0072\u0073\u0069\u0073\u0074\u0065\u006e\u0063\u0065"{_eae .PersistenceAttr .UnmarshalXMLAttr (_dgc );continue ;};};_bab :for {_gefc ,_bgc :=d .Token ();
if _bgc !=nil {return _bgc ;};switch _bef :=_gefc .(type ){case _b .StartElement :switch _bef .Name {case _b .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u006f\u0063\u0078P\u0072"}:_fbe :=NewCT_OcxPr ();
if _bcfg :=d .DecodeElement (_fbe ,&_bef );_bcfg !=nil {return _bcfg ;};_eae .OcxPr =append (_eae .OcxPr ,_fbe );default:_f .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006fn\u0020\u004fc\u0078\u0020\u0025\u0076",_bef .Name );
if _egdd :=d .Skip ();_egdd !=nil {return _egdd ;};};case _b .EndElement :break _bab ;case _b .CharData :};};return nil ;};

// ValidateWithPath validates the Ocx and its children, prefixing error messages with path
func (_dbd *Ocx )ValidateWithPath (path string )error {if _ddg :=_dbd .CT_Ocx .ValidateWithPath (path );_ddg !=nil {return _ddg ;};return nil ;};func (_gfd *CT_Picture )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for _ ,_ggfc :=range start .Attr {if _ggfc .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"&&_ggfc .Name .Local =="\u0069\u0064"||_ggfc .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073"&&_ggfc .Name .Local =="\u0069\u0064"{_gag ,_cca :=_ggfc .Value ,error (nil );
if _cca !=nil {return _cca ;};_gfd .IdAttr =&_gag ;continue ;};};for {_cbb ,_ce :=d .Token ();if _ce !=nil {return _bf .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065\u003a\u0020%\u0073",_ce );
};if _cgae ,_fee :=_cbb .(_b .EndElement );_fee &&_cgae .Name ==start .Name {break ;};};return nil ;};func (_fba *CT_Picture )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {if _fba .IdAttr !=nil {start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0072\u003a\u0069\u0064"},Value :_bf .Sprintf ("\u0025\u0076",*_fba .IdAttr )});
};e .EncodeToken (start );e .EncodeToken (_b .EndElement {Name :start .Name });return nil ;};func init (){_ba .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058","\u0043\u0054\u005f\u004f\u0063\u0078",NewCT_Ocx );
_ba .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058","\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072",NewCT_OcxPr );
_ba .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058","\u0043T\u005f\u0046\u006f\u006e\u0074",NewCT_Font );
_ba .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058","\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065",NewCT_Picture );
_ba .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058","\u006f\u0063\u0078",NewOcx );
};