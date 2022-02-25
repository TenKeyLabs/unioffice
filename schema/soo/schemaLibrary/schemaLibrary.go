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

package schemaLibrary ;import (_b "encoding/xml";_bf "fmt";_e "github.com/unidoc/unioffice";_g "github.com/unidoc/unioffice/common/logger";);

// Validate validates the SchemaLibrary and its children
func (_dae *SchemaLibrary )Validate ()error {return _dae .ValidateWithPath ("\u0053\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079");};func (_ea *CT_Schema )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {if _ea .UriAttr !=nil {start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u006d\u0061\u003a\u0075\u0072\u0069"},Value :_bf .Sprintf ("\u0025\u0076",*_ea .UriAttr )});};if _ea .ManifestLocationAttr !=nil {start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u006d\u0061\u003a\u006dan\u0069\u0066\u0065\u0073\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"},Value :_bf .Sprintf ("\u0025\u0076",*_ea .ManifestLocationAttr )});};if _ea .SchemaLocationAttr !=nil {start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u006d\u0061\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u006f\u0063a\u0074\u0069\u006f\u006e"},Value :_bf .Sprintf ("\u0025\u0076",*_ea .SchemaLocationAttr )});};if _ea .SchemaLanguageAttr !=nil {start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u006d\u0061\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0061\u006eg\u0075\u0061\u0067\u0065"},Value :_bf .Sprintf ("\u0025\u0076",*_ea .SchemaLanguageAttr )});};e .EncodeToken (start );e .EncodeToken (_b .EndElement {Name :start .Name });return nil ;};func (_be *CT_SchemaLibrary )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_ed :for {_gag ,_fa :=d .Token ();if _fa !=nil {return _fa ;};switch _edc :=_gag .(type ){case _b .StartElement :switch _edc .Name {case _b .Name {Space :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e",Local :"\u0073\u0063\u0068\u0065\u006d\u0061"}:_dc :=NewCT_Schema ();if _ab :=d .DecodeElement (_dc ,&_edc );_ab !=nil {return _ab ;};_be .Schema =append (_be .Schema ,_dc );default:_g .Log .Debug ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u0020\u0025v",_edc .Name );if _ba :=d .Skip ();_ba !=nil {return _ba ;};};case _b .EndElement :break _ed ;case _b .CharData :};};return nil ;};func (_eb *CT_SchemaLibrary )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {e .EncodeToken (start );if _eb .Schema !=nil {_fd :=_b .StartElement {Name :_b .Name {Local :"\u006da\u003a\u0073\u0063\u0068\u0065\u006da"}};for _ ,_dad :=range _eb .Schema {e .EncodeElement (_dad ,_fd );};};e .EncodeToken (_b .EndElement {Name :start .Name });return nil ;};func NewCT_Schema ()*CT_Schema {_ga :=&CT_Schema {};return _ga };func (_d *CT_Schema )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {for _ ,_gb :=range start .Attr {if _gb .Name .Local =="\u0075\u0072\u0069"{_de ,_c :=_gb .Value ,error (nil );if _c !=nil {return _c ;};_d .UriAttr =&_de ;continue ;};if _gb .Name .Local =="\u006d\u0061n\u0069\u0066\u0065s\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"{_bc ,_gg :=_gb .Value ,error (nil );if _gg !=nil {return _gg ;};_d .ManifestLocationAttr =&_bc ;continue ;};if _gb .Name .Local =="\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u006f\u0063a\u0074\u0069\u006f\u006e"{_a ,_da :=_gb .Value ,error (nil );if _da !=nil {return _da ;};_d .SchemaLocationAttr =&_a ;continue ;};if _gb .Name .Local =="\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0061\u006eg\u0075\u0061\u0067\u0065"{_ca ,_dag :=_gb .Value ,error (nil );if _dag !=nil {return _dag ;};_d .SchemaLanguageAttr =&_ca ;continue ;};};for {_dec ,_ef :=d .Token ();if _ef !=nil {return _bf .Errorf ("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0053\u0063\u0068e\u006d\u0061\u003a\u0020\u0025\u0073",_ef );};if _db ,_gc :=_dec .(_b .EndElement );_gc &&_db .Name ==start .Name {break ;};};return nil ;};type CT_Schema struct{UriAttr *string ;ManifestLocationAttr *string ;SchemaLocationAttr *string ;SchemaLanguageAttr *string ;};func NewSchemaLibrary ()*SchemaLibrary {_ae :=&SchemaLibrary {};_ae .CT_SchemaLibrary =*NewCT_SchemaLibrary ();return _ae ;};

// Validate validates the CT_Schema and its children
func (_fg *CT_Schema )Validate ()error {return _fg .ValidateWithPath ("\u0043T\u005f\u0053\u0063\u0068\u0065\u006da");};

// ValidateWithPath validates the CT_Schema and its children, prefixing error messages with path
func (_gdg *CT_Schema )ValidateWithPath (path string )error {return nil };func (_ad *SchemaLibrary )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"});start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u006d\u0061"},Value :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"});start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u006d\u0061:\u0073\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079";return _ad .CT_SchemaLibrary .MarshalXML (e ,start );};type SchemaLibrary struct{CT_SchemaLibrary };type CT_SchemaLibrary struct{Schema []*CT_Schema ;};func (_eg *SchemaLibrary )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_eg .CT_SchemaLibrary =*NewCT_SchemaLibrary ();_cea :for {_dbd ,_dg :=d .Token ();if _dg !=nil {return _dg ;};switch _dbdd :=_dbd .(type ){case _b .StartElement :switch _dbdd .Name {case _b .Name {Space :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e",Local :"\u0073\u0063\u0068\u0065\u006d\u0061"}:_ega :=NewCT_Schema ();if _bcd :=d .DecodeElement (_ega ,&_dbdd );_bcd !=nil {return _bcd ;};_eg .Schema =append (_eg .Schema ,_ega );default:_g .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0053\u0063\u0068\u0065m\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079 \u0025\u0076",_dbdd .Name );if _fe :=d .Skip ();_fe !=nil {return _fe ;};};case _b .EndElement :break _cea ;case _b .CharData :};};return nil ;};

// Validate validates the CT_SchemaLibrary and its children
func (_beb *CT_SchemaLibrary )Validate ()error {return _beb .ValidateWithPath ("\u0043\u0054_\u0053\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079");};func NewCT_SchemaLibrary ()*CT_SchemaLibrary {_dbf :=&CT_SchemaLibrary {};return _dbf };

// ValidateWithPath validates the SchemaLibrary and its children, prefixing error messages with path
func (_gdb *SchemaLibrary )ValidateWithPath (path string )error {if _bag :=_gdb .CT_SchemaLibrary .ValidateWithPath (path );_bag !=nil {return _bag ;};return nil ;};

// ValidateWithPath validates the CT_SchemaLibrary and its children, prefixing error messages with path
func (_cf *CT_SchemaLibrary )ValidateWithPath (path string )error {for _gab ,_edce :=range _cf .Schema {if _ce :=_edce .ValidateWithPath (_bf .Sprintf ("\u0025\u0073\u002f\u0053\u0063\u0068\u0065\u006d\u0061\u005b\u0025\u0064\u005d",path ,_gab ));_ce !=nil {return _ce ;};};return nil ;};func init (){_e .RegisterConstructor ("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e","\u0043T\u005f\u0053\u0063\u0068\u0065\u006da",NewCT_Schema );_e .RegisterConstructor ("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e","\u0043\u0054_\u0053\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079",NewCT_SchemaLibrary );_e .RegisterConstructor ("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e","\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079",NewSchemaLibrary );};