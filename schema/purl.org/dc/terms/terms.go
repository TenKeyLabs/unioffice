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

package terms ;import (_b "encoding/xml";_ca "fmt";_a "github.com/unidoc/unioffice";_e "github.com/unidoc/unioffice/common/logger";_ea "github.com/unidoc/unioffice/schema/purl.org/dc/elements";);

// Validate validates the RFC1766 and its children
func (_gef *RFC1766 )Validate ()error {return _gef .ValidateWithPath ("\u0052F\u0043\u0031\u0037\u0036\u0036");};func (_dg *DDC )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for {_eg ,_fb :=d .Token ();if _fb !=nil {return _ca .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0044\u0044\u0043\u003a\u0020\u0025\u0073",_fb );
};if _db ,_ede :=_eg .(_b .EndElement );_ede &&_db .Name ==start .Name {break ;};};return nil ;};

// Validate validates the LCC and its children
func (_fdc *LCC )Validate ()error {return _fdc .ValidateWithPath ("\u004c\u0043\u0043")};

// Validate validates the ElementsAndRefinementsGroup and its children
func (_bbb *ElementsAndRefinementsGroup )Validate ()error {return _bbb .ValidateWithPath ("E\u006c\u0065\u006d\u0065\u006e\u0074s\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065m\u0065\u006e\u0074s\u0047r\u006f\u0075\u0070");};type TGN struct{};func NewElementsAndRefinementsGroup ()*ElementsAndRefinementsGroup {_ddd :=&ElementsAndRefinementsGroup {};
return _ddd ;};func (_gcb *ISO3166 )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Name .Local ="\u0049S\u004f\u0033\u0031\u0036\u0036";e .EncodeToken (start );e .EncodeToken (_b .EndElement {Name :start .Name });return nil ;};func (_ada *Point )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for {_gcfc ,_fdeg :=d .Token ();
if _fdeg !=nil {return _ca .Errorf ("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0050\u006f\u0069\u006et\u003a\u0020\u0025\u0073",_fdeg );};if _gaeg ,_dca :=_gcfc .(_b .EndElement );_dca &&_gaeg .Name ==start .Name {break ;};};return nil ;};func NewDCMIType ()*DCMIType {_ag :=&DCMIType {};
return _ag };func (_ba *Box )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for {_f ,_fd :=d .Token ();if _fd !=nil {return _ca .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0042\u006f\u0078\u003a\u0020\u0025\u0073",_fd );};if _ad ,_af :=_f .(_b .EndElement );
_af &&_ad .Name ==start .Name {break ;};};return nil ;};type ElementOrRefinementContainer struct{Choice []*ElementsAndRefinementsGroupChoice ;};func (_dfgc *RFC1766 )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Name .Local ="\u0052F\u0043\u0031\u0037\u0036\u0036";
e .EncodeToken (start );e .EncodeToken (_b .EndElement {Name :start .Name });return nil ;};type MESH struct{};func (_cfd *ElementOrRefinementContainer )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_aa :for {_cad ,_agc :=d .Token ();if _agc !=nil {return _agc ;
};switch _gd :=_cad .(type ){case _b .StartElement :switch _gd .Name {case _b .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_dc :=NewElementsAndRefinementsGroupChoice ();
if _fe :=d .DecodeElement (&_dc .Any ,&_gd );_fe !=nil {return _fe ;};_cfd .Choice =append (_cfd .Choice ,_dc );default:_e .Log .Debug ("\u0073k\u0069\u0070\u0070\u0069\u006e\u0067\u0020un\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006de\u006e\u0074 \u006f\u006e\u0020E\u006c\u0065\u006d\u0065\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065n\u0074\u0043on\u0074\u0061\u0069n\u0065\u0072\u0020\u0025\u0076",_gd .Name );
if _ge :=d .Skip ();_ge !=nil {return _ge ;};};case _b .EndElement :break _aa ;case _b .CharData :};};return nil ;};

// Validate validates the DDC and its children
func (_ae *DDC )Validate ()error {return _ae .ValidateWithPath ("\u0044\u0044\u0043")};

// Validate validates the ISO639_2 and its children
func (_dfb *ISO639_2 )Validate ()error {return _dfb .ValidateWithPath ("\u0049\u0053\u004f\u0036\u0033\u0039\u005f\u0032");};type Point struct{};func (_ged *ElementsAndRefinementsGroupChoice )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_bdb :for {_egb ,_gdf :=d .Token ();
if _gdf !=nil {return _gdf ;};switch _ebc :=_egb .(type ){case _b .StartElement :switch _ebc .Name {case _b .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_ef :=_ea .NewAny ();
if _bg :=d .DecodeElement (_ef ,&_ebc );_bg !=nil {return _bg ;};_ged .Any =append (_ged .Any ,_ef );default:_e .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0041\u006ed\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006fu\u0070\u0043\u0068o\u0069\u0063\u0065\u0020\u0025\u0076",_ebc .Name );
if _dgf :=d .Skip ();_dgf !=nil {return _dgf ;};};case _b .EndElement :break _bdb ;case _b .CharData :};};return nil ;};func (_bd *ElementOrRefinementContainer )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Name .Local ="\u0065\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072";
e .EncodeToken (start );if _bd .Choice !=nil {for _ ,_caa :=range _bd .Choice {_caa .MarshalXML (e ,_b .StartElement {});};};e .EncodeToken (_b .EndElement {Name :start .Name });return nil ;};

// Validate validates the RFC3066 and its children
func (_bed *RFC3066 )Validate ()error {return _bed .ValidateWithPath ("\u0052F\u0043\u0033\u0030\u0036\u0036");};func (_cb *LCC )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Name .Local ="\u004c\u0043\u0043";e .EncodeToken (start );
e .EncodeToken (_b .EndElement {Name :start .Name });return nil ;};func (_dfg *IMT )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for {_bc ,_aaf :=d .Token ();if _aaf !=nil {return _ca .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0049\u004d\u0054\u003a\u0020\u0025\u0073",_aaf );
};if _gab ,_afg :=_bc .(_b .EndElement );_afg &&_gab .Name ==start .Name {break ;};};return nil ;};

// Validate validates the ISO3166 and its children
func (_gaa *ISO3166 )Validate ()error {return _gaa .ValidateWithPath ("\u0049S\u004f\u0033\u0031\u0036\u0036");};type IMT struct{};

// Validate validates the LCSH and its children
func (_ddc *LCSH )Validate ()error {return _ddc .ValidateWithPath ("\u004c\u0043\u0053\u0048")};

// ValidateWithPath validates the W3CDTF and its children, prefixing error messages with path
func (_fda *W3CDTF )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the IMT and its children, prefixing error messages with path
func (_cde *IMT )ValidateWithPath (path string )error {return nil };func NewUDC ()*UDC {_dddf :=&UDC {};return _dddf };

// ValidateWithPath validates the Box and its children, prefixing error messages with path
func (_ce *Box )ValidateWithPath (path string )error {return nil };func NewIMT ()*IMT {_be :=&IMT {};return _be };

// Validate validates the TGN and its children
func (_edc *TGN )Validate ()error {return _edc .ValidateWithPath ("\u0054\u0047\u004e")};type ElementsAndRefinementsGroupChoice struct{Any []*_ea .Any ;};

// Validate validates the URI and its children
func (_gdd *URI )Validate ()error {return _gdd .ValidateWithPath ("\u0055\u0052\u0049")};func (_cff *ElementsAndRefinementsGroup )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {if _cff .Choice !=nil {for _ ,_bb :=range _cff .Choice {_bb .MarshalXML (e ,_b .StartElement {});
};};return nil ;};func (_gce *LCSH )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Name .Local ="\u004c\u0043\u0053\u0048";e .EncodeToken (start );e .EncodeToken (_b .EndElement {Name :start .Name });return nil ;};

// Validate validates the IMT and its children
func (_ac *IMT )Validate ()error {return _ac .ValidateWithPath ("\u0049\u004d\u0054")};type ISO639_2 struct{};type UDC struct{};

// ValidateWithPath validates the DCMIType and its children, prefixing error messages with path
func (_df *DCMIType )ValidateWithPath (path string )error {return nil };func NewDDC ()*DDC {_ee :=&DDC {};return _ee };func NewRFC1766 ()*RFC1766 {_agb :=&RFC1766 {};return _agb };

// ValidateWithPath validates the TGN and its children, prefixing error messages with path
func (_cdg *TGN )ValidateWithPath (path string )error {return nil };func NewElementOrRefinementContainer ()*ElementOrRefinementContainer {_ec :=&ElementOrRefinementContainer {};return _ec ;};type URI struct{};type DDC struct{};func (_edd *IMT )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Name .Local ="\u0049\u004d\u0054";
e .EncodeToken (start );e .EncodeToken (_b .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the ElementsAndRefinementsGroup and its children, prefixing error messages with path
func (_ceb *ElementsAndRefinementsGroup )ValidateWithPath (path string )error {for _gca ,_dge :=range _ceb .Choice {if _abc :=_dge .ValidateWithPath (_ca .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_gca ));
_abc !=nil {return _abc ;};};return nil ;};

// Validate validates the ElementsAndRefinementsGroupChoice and its children
func (_bgc *ElementsAndRefinementsGroupChoice )Validate ()error {return _bgc .ValidateWithPath ("\u0045\u006c\u0065\u006d\u0065\u006et\u0073\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006et\u0073\u0047\u0072\u006f\u0075\u0070\u0043h\u006f\u0069\u0063\u0065");
};type Period struct{};

// ValidateWithPath validates the MESH and its children, prefixing error messages with path
func (_aeb *MESH )ValidateWithPath (path string )error {return nil };func (_bfde *URI )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Name .Local ="\u0055\u0052\u0049";e .EncodeToken (start );e .EncodeToken (_b .EndElement {Name :start .Name });
return nil ;};

// Validate validates the Point and its children
func (_dec *Point )Validate ()error {return _dec .ValidateWithPath ("\u0050\u006f\u0069n\u0074")};func (_edee *LCSH )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for {_gea ,_fcd :=d .Token ();if _fcd !=nil {return _ca .Errorf ("\u0070\u0061r\u0073\u0069\u006eg\u0020\u004c\u0043\u0053\u0048\u003a\u0020\u0025\u0073",_fcd );
};if _gcf ,_fde :=_gea .(_b .EndElement );_fde &&_gcf .Name ==start .Name {break ;};};return nil ;};type W3CDTF struct{};type RFC3066 struct{};func (_cfg *URI )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for {_ecce ,_ffad :=d .Token ();
if _ffad !=nil {return _ca .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0055\u0052\u0049\u003a\u0020\u0025\u0073",_ffad );};if _dfd ,_cbg :=_ecce .(_b .EndElement );_cbg &&_dfd .Name ==start .Name {break ;};};return nil ;};func NewTGN ()*TGN {_cab :=&TGN {};
return _cab };type DCMIType struct{};func (_ddcb *MESH )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Name .Local ="\u004d\u0045\u0053\u0048";e .EncodeToken (start );e .EncodeToken (_b .EndElement {Name :start .Name });return nil ;};
func (_afgg *MESH )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for {_efg ,_bbc :=d .Token ();if _bbc !=nil {return _ca .Errorf ("\u0070\u0061r\u0073\u0069\u006eg\u0020\u004d\u0045\u0053\u0048\u003a\u0020\u0025\u0073",_bbc );};if _afe ,_ffd :=_efg .(_b .EndElement );
_ffd &&_afe .Name ==start .Name {break ;};};return nil ;};func NewPoint ()*Point {_fdce :=&Point {};return _fdce };

// ValidateWithPath validates the ElementOrRefinementContainer and its children, prefixing error messages with path
func (_eag *ElementOrRefinementContainer )ValidateWithPath (path string )error {for _gf ,_abg :=range _eag .Choice {if _aac :=_abg .ValidateWithPath (_ca .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_gf ));
_aac !=nil {return _aac ;};};return nil ;};

// Validate validates the Period and its children
func (_dac *Period )Validate ()error {return _dac .ValidateWithPath ("\u0050\u0065\u0072\u0069\u006f\u0064");};

// ValidateWithPath validates the Point and its children, prefixing error messages with path
func (_efc *Point )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the ISO3166 and its children, prefixing error messages with path
func (_ecg *ISO3166 )ValidateWithPath (path string )error {return nil };func (_bba *RFC1766 )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for {_agg ,_bfc :=d .Token ();if _bfc !=nil {return _ca .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0052\u0046\u0043\u0031\u0037\u0036\u0036\u003a\u0020\u0025\u0073",_bfc );
};if _agcg ,_dea :=_agg .(_b .EndElement );_dea &&_agcg .Name ==start .Name {break ;};};return nil ;};func (_fg *DDC )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Name .Local ="\u0044\u0044\u0043";e .EncodeToken (start );e .EncodeToken (_b .EndElement {Name :start .Name });
return nil ;};

// Validate validates the Box and its children
func (_ed *Box )Validate ()error {return _ed .ValidateWithPath ("\u0042\u006f\u0078")};func NewLCC ()*LCC {_fab :=&LCC {};return _fab };func (_cd *Box )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Name .Local ="\u0042\u006f\u0078";
e .EncodeToken (start );e .EncodeToken (_b .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the RFC1766 and its children, prefixing error messages with path
func (_fdca *RFC1766 )ValidateWithPath (path string )error {return nil };func (_edcg *W3CDTF )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Name .Local ="\u0057\u0033\u0043\u0044\u0054\u0046";e .EncodeToken (start );e .EncodeToken (_b .EndElement {Name :start .Name });
return nil ;};

// ValidateWithPath validates the ElementsAndRefinementsGroupChoice and its children, prefixing error messages with path
func (_ecc *ElementsAndRefinementsGroupChoice )ValidateWithPath (path string )error {for _dcg ,_ff :=range _ecc .Any {if _eeg :=_ff .ValidateWithPath (_ca .Sprintf ("\u0025\u0073\u002f\u0041\u006e\u0079\u005b\u0025\u0064\u005d",path ,_dcg ));_eeg !=nil {return _eeg ;
};};return nil ;};type LCC struct{};func (_gc *DCMIType )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for {_ab ,_eb :=d .Token ();if _eb !=nil {return _ca .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0044\u0043\u004dI\u0054\u0079\u0070\u0065: \u0025\u0073",_eb );
};if _abb ,_cac :=_ab .(_b .EndElement );_cac &&_abb .Name ==start .Name {break ;};};return nil ;};func NewISO3166 ()*ISO3166 {_cda :=&ISO3166 {};return _cda };func NewISO639_2 ()*ISO639_2 {_afa :=&ISO639_2 {};return _afa };

// Validate validates the UDC and its children
func (_aed *UDC )Validate ()error {return _aed .ValidateWithPath ("\u0055\u0044\u0043")};

// Validate validates the W3CDTF and its children
func (_cg *W3CDTF )Validate ()error {return _cg .ValidateWithPath ("\u0057\u0033\u0043\u0044\u0054\u0046");};

// ValidateWithPath validates the Period and its children, prefixing error messages with path
func (_cebc *Period )ValidateWithPath (path string )error {return nil };func NewRFC3066 ()*RFC3066 {_adfd :=&RFC3066 {};return _adfd };type RFC1766 struct{};func (_ddf *ElementsAndRefinementsGroupChoice )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {if _ddf .Any !=nil {_bff :=_b .StartElement {Name :_b .Name {Local :"\u0064\u0063\u003a\u0061\u006e\u0079"}};
for _ ,_da :=range _ddf .Any {e .EncodeElement (_da ,_bff );};};return nil ;};func NewURI ()*URI {_cec :=&URI {};return _cec };func (_eae *LCC )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for {_gg ,_eee :=d .Token ();if _eee !=nil {return _ca .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u004c\u0043\u0043\u003a\u0020\u0025\u0073",_eee );
};if _cdag ,_fc :=_gg .(_b .EndElement );_fc &&_cdag .Name ==start .Name {break ;};};return nil ;};func NewLCSH ()*LCSH {_bdbg :=&LCSH {};return _bdbg };func (_ggg *Period )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Name .Local ="\u0050\u0065\u0072\u0069\u006f\u0064";
e .EncodeToken (start );e .EncodeToken (_b .EndElement {Name :start .Name });return nil ;};

// Validate validates the DCMIType and its children
func (_d *DCMIType )Validate ()error {return _d .ValidateWithPath ("\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065");};type LCSH struct{};

// Validate validates the ElementOrRefinementContainer and its children
func (_de *ElementOrRefinementContainer )Validate ()error {return _de .ValidateWithPath ("\u0045\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072");};func NewPeriod ()*Period {_gec :=&Period {};
return _gec };

// Validate validates the MESH and its children
func (_ffa *MESH )Validate ()error {return _ffa .ValidateWithPath ("\u004d\u0045\u0053\u0048")};func (_fga *UDC )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Name .Local ="\u0055\u0044\u0043";e .EncodeToken (start );e .EncodeToken (_b .EndElement {Name :start .Name });
return nil ;};func (_ggf *UDC )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for {_ebb ,_dab :=d .Token ();if _dab !=nil {return _ca .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0055\u0044\u0043\u003a\u0020\u0025\u0073",_dab );};if _bbd ,_ddg :=_ebb .(_b .EndElement );
_ddg &&_bbd .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the DDC and its children, prefixing error messages with path
func (_dd *DDC )ValidateWithPath (path string )error {return nil };func NewElementsAndRefinementsGroupChoice ()*ElementsAndRefinementsGroupChoice {_bf :=&ElementsAndRefinementsGroupChoice {};return _bf ;};func (_afd *W3CDTF )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for {_dee ,_gad :=d .Token ();
if _gad !=nil {return _ca .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u00573\u0043\u0044T\u0046\u003a\u0020\u0025\u0073",_gad );};if _debf ,_efec :=_dee .(_b .EndElement );_efec &&_debf .Name ==start .Name {break ;};};return nil ;};func (_efgd *RFC3066 )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Name .Local ="\u0052F\u0043\u0033\u0030\u0036\u0036";
e .EncodeToken (start );e .EncodeToken (_b .EndElement {Name :start .Name });return nil ;};func (_efe *Period )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for {_bfd ,_ded :=d .Token ();if _ded !=nil {return _ca .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0050e\u0072\u0069o\u0064\u003a\u0020\u0025\u0073",_ded );
};if _dcc ,_gcg :=_bfd .(_b .EndElement );_gcg &&_dcc .Name ==start .Name {break ;};};return nil ;};func (_aaa *ISO639_2 )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Name .Local ="\u0049\u0053\u004f\u0036\u0033\u0039\u002d\u0032";
e .EncodeToken (start );e .EncodeToken (_b .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the URI and its children, prefixing error messages with path
func (_efa *URI )ValidateWithPath (path string )error {return nil };type ElementsAndRefinementsGroup struct{Choice []*ElementsAndRefinementsGroupChoice ;};

// ValidateWithPath validates the ISO639_2 and its children, prefixing error messages with path
func (_aad *ISO639_2 )ValidateWithPath (path string )error {return nil };func (_cdb *Point )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Name .Local ="\u0050\u006f\u0069n\u0074";e .EncodeToken (start );e .EncodeToken (_b .EndElement {Name :start .Name });
return nil ;};func NewBox ()*Box {_cf :=&Box {};return _cf };

// ValidateWithPath validates the RFC3066 and its children, prefixing error messages with path
func (_gfcb *RFC3066 )ValidateWithPath (path string )error {return nil };func (_g *DCMIType )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Name .Local ="\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065";e .EncodeToken (start );e .EncodeToken (_b .EndElement {Name :start .Name });
return nil ;};

// ValidateWithPath validates the UDC and its children, prefixing error messages with path
func (_bcgc *UDC )ValidateWithPath (path string )error {return nil };type ISO3166 struct{};type Box struct{};func NewMESH ()*MESH {_gae :=&MESH {};return _gae };func (_fgd *ElementsAndRefinementsGroup )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_bbf :for {_ga ,_abgc :=d .Token ();
if _abgc !=nil {return _abgc ;};switch _fa :=_ga .(type ){case _b .StartElement :switch _fa .Name {case _b .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_adf :=NewElementsAndRefinementsGroupChoice ();
if _bda :=d .DecodeElement (&_adf .Any ,&_fa );_bda !=nil {return _bda ;};_fgd .Choice =append (_fgd .Choice ,_adf );default:_e .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006ce\u006d\u0065\u006e\u0074\u0073\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006et\u0073\u0047\u0072\u006f\u0075\u0070\u0020\u0025\u0076",_fa .Name );
if _cca :=d .Skip ();_cca !=nil {return _cca ;};};case _b .EndElement :break _bbf ;case _b .CharData :};};return nil ;};func NewW3CDTF ()*W3CDTF {_gga :=&W3CDTF {};return _gga };

// ValidateWithPath validates the LCC and its children, prefixing error messages with path
func (_ced *LCC )ValidateWithPath (path string )error {return nil };func (_dgfb *TGN )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Name .Local ="\u0054\u0047\u004e";e .EncodeToken (start );e .EncodeToken (_b .EndElement {Name :start .Name });
return nil ;};func (_cdf *TGN )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for {_bgf ,_feg :=d .Token ();if _feg !=nil {return _ca .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0054\u0047\u004e\u003a\u0020\u0025\u0073",_feg );};if _fcf ,_bga :=_bgf .(_b .EndElement );
_bga &&_fcf .Name ==start .Name {break ;};};return nil ;};func (_fdb *RFC3066 )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for {_gda ,_fdf :=d .Token ();if _fdf !=nil {return _ca .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0052\u0046\u0043\u0033\u0030\u0036\u0036\u003a\u0020\u0025\u0073",_fdf );
};if _gfc ,_cdd :=_gda .(_b .EndElement );_cdd &&_gfc .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the LCSH and its children, prefixing error messages with path
func (_ebe *LCSH )ValidateWithPath (path string )error {return nil };func (_bcg *ISO639_2 )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for {_cea ,_deb :=d .Token ();if _deb !=nil {return _ca .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0049\u0053\u004f6\u0033\u0039\u005f\u0032: \u0025\u0073",_deb );
};if _faf ,_fgda :=_cea .(_b .EndElement );_fgda &&_faf .Name ==start .Name {break ;};};return nil ;};func (_bcc *ISO3166 )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for {_babf ,_eed :=d .Token ();if _eed !=nil {return _ca .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0049\u0053\u004f\u0033\u0031\u0036\u0036\u003a\u0020\u0025\u0073",_eed );
};if _dfe ,_cag :=_babf .(_b .EndElement );_cag &&_dfe .Name ==start .Name {break ;};};return nil ;};func init (){_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u004c\u0043\u0053\u0048",NewLCSH );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u004d\u0045\u0053\u0048",NewMESH );_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0044\u0044\u0043",NewDDC );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u004c\u0043\u0043",NewLCC );_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0055\u0044\u0043",NewUDC );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0050\u0065\u0072\u0069\u006f\u0064",NewPeriod );_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0057\u0033\u0043\u0044\u0054\u0046",NewW3CDTF );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065",NewDCMIType );_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0049\u004d\u0054",NewIMT );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0055\u0052\u0049",NewURI );_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0049\u0053\u004f\u0036\u0033\u0039\u002d\u0032",NewISO639_2 );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0052F\u0043\u0031\u0037\u0036\u0036",NewRFC1766 );_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0052F\u0043\u0033\u0030\u0036\u0036",NewRFC3066 );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0050\u006f\u0069n\u0074",NewPoint );_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0049S\u004f\u0033\u0031\u0036\u0036",NewISO3166 );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0042\u006f\u0078",NewBox );_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0054\u0047\u004e",NewTGN );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0065\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072",NewElementOrRefinementContainer );
_a .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","e\u006c\u0065\u006d\u0065\u006e\u0074s\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065m\u0065\u006e\u0074s\u0047r\u006f\u0075\u0070",NewElementsAndRefinementsGroup );
};