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

package core_properties ;import (_d "encoding/xml";_db "fmt";_ae "github.com/unidoc/unioffice";_b "github.com/unidoc/unioffice/common/logger";_a "time";);

// ValidateWithPath validates the CT_Keywords and its children, prefixing error messages with path
func (_af *CT_Keywords )ValidateWithPath (path string )error {for _beba ,_eb :=range _af .Value {if _beeg :=_eb .ValidateWithPath (_db .Sprintf ("\u0025\u0073\u002fV\u0061\u006c\u0075\u0065\u005b\u0025\u0064\u005d",path ,_beba ));_beeg !=nil {return _beeg ;};};return nil ;};func (_ag *CT_Keywords )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for _ ,_bcg :=range start .Attr {if _bcg .Name .Space =="\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"&&_bcg .Name .Local =="\u006c\u0061\u006e\u0067"{_dca ,_egdd :=_bcg .Value ,error (nil );if _egdd !=nil {return _egdd ;};_ag .LangAttr =&_dca ;continue ;};};_bee :for {_fag ,_fe :=d .Token ();if _fe !=nil {return _fe ;};switch _bb :=_fag .(type ){case _d .StartElement :switch _bb .Name {case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0076\u0061\u006cu\u0065"}:_cd :=NewCT_Keyword ();if _aeg :=d .DecodeElement (_cd ,&_bb );_aeg !=nil {return _aeg ;};_ag .Value =append (_ag .Value ,_cd );default:_b .Log .Debug ("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073\u0020\u0025\u0076",_bb .Name );if _bgb :=d .Skip ();_bgb !=nil {return _bgb ;};};case _d .EndElement :break _bee ;case _d .CharData :};};return nil ;};func (_fc *CT_CoreProperties )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {e .EncodeToken (start );if _fc .Category !=nil {_c :=_d .StartElement {Name :_d .Name {Local :"c\u0070\u003a\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}};_ae .AddPreserveSpaceAttr (&_c ,*_fc .Category );e .EncodeElement (_fc .Category ,_c );};if _fc .ContentStatus !=nil {_cf :=_d .StartElement {Name :_d .Name {Local :"\u0063\u0070:\u0063\u006f\u006et\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}};_ae .AddPreserveSpaceAttr (&_cf ,*_fc .ContentStatus );e .EncodeElement (_fc .ContentStatus ,_cf );};if _fc .Created !=nil {_e :=_d .StartElement {Name :_d .Name {Local :"\u0064c\u0074e\u0072\u006d\u0073\u003a\u0063\u0072\u0065\u0061\u0074\u0065\u0064"}};e .EncodeElement (_fc .Created ,_e );};if _fc .Creator !=nil {_ee :=_d .StartElement {Name :_d .Name {Local :"\u0064\u0063\u003a\u0063\u0072\u0065\u0061\u0074\u006f\u0072"}};e .EncodeElement (_fc .Creator ,_ee );};if _fc .Description !=nil {_ab :=_d .StartElement {Name :_d .Name {Local :"\u0064\u0063\u003a\u0064\u0065\u0073\u0063\u0072\u0069p\u0074\u0069\u006f\u006e"}};e .EncodeElement (_fc .Description ,_ab );};if _fc .Identifier !=nil {_ff :=_d .StartElement {Name :_d .Name {Local :"\u0064\u0063\u003a\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}};e .EncodeElement (_fc .Identifier ,_ff );};if _fc .Keywords !=nil {_cb :=_d .StartElement {Name :_d .Name {Local :"c\u0070\u003a\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}};e .EncodeElement (_fc .Keywords ,_cb );};if _fc .Language !=nil {_ef :=_d .StartElement {Name :_d .Name {Local :"d\u0063\u003a\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}};e .EncodeElement (_fc .Language ,_ef );};if _fc .LastModifiedBy !=nil {_bc :=_d .StartElement {Name :_d .Name {Local :"\u0063\u0070\u003a\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}};_ae .AddPreserveSpaceAttr (&_bc ,*_fc .LastModifiedBy );e .EncodeElement (_fc .LastModifiedBy ,_bc );};if _fc .LastPrinted !=nil {_cbf :=_d .StartElement {Name :_d .Name {Local :"\u0063\u0070\u003a\u006c\u0061\u0073\u0074\u0050\u0072i\u006e\u0074\u0065\u0064"}};e .EncodeElement (_fc .LastPrinted ,_cbf );};if _fc .Modified !=nil {_fg :=_d .StartElement {Name :_d .Name {Local :"\u0064\u0063t\u0065\u0072\u006ds\u003a\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}};e .EncodeElement (_fc .Modified ,_fg );};if _fc .Revision !=nil {_ac :=_d .StartElement {Name :_d .Name {Local :"c\u0070\u003a\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}};_ae .AddPreserveSpaceAttr (&_ac ,*_fc .Revision );e .EncodeElement (_fc .Revision ,_ac );};if _fc .Subject !=nil {_g :=_d .StartElement {Name :_d .Name {Local :"\u0064\u0063\u003a\u0073\u0075\u0062\u006a\u0065\u0063\u0074"}};e .EncodeElement (_fc .Subject ,_g );};if _fc .Title !=nil {_cff :=_d .StartElement {Name :_d .Name {Local :"\u0064\u0063\u003a\u0074\u0069\u0074\u006c\u0065"}};e .EncodeElement (_fc .Title ,_cff );};if _fc .Version !=nil {_fca :=_d .StartElement {Name :_d .Name {Local :"\u0063\u0070\u003a\u0076\u0065\u0072\u0073\u0069\u006f\u006e"}};_ae .AddPreserveSpaceAttr (&_fca ,*_fc .Version );e .EncodeElement (_fc .Version ,_fca );};e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func (_agf *CoreProperties )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073"});start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0063\u0070"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073"});start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0063"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f"});start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0063\u0074\u0065\u0072\u006d\u0073"},Value :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/"});start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u0063\u0070\u003a\u0063\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073";return _agf .CT_CoreProperties .MarshalXML (e ,start );};func NewCT_CoreProperties ()*CT_CoreProperties {_be :=&CT_CoreProperties {};return _be };func (_gg *CT_CoreProperties )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_bf :for {_gd ,_eg :=d .Token ();if _eg !=nil {return _eg ;};switch _fd :=_gd .(type ){case _d .StartElement :switch _fd .Name {case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}:_gg .Category =new (string );if _gb :=d .DecodeElement (_gg .Category ,&_fd );_gb !=nil {return _gb ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}:_gg .ContentStatus =new (string );if _ea :=d .DecodeElement (_gg .ContentStatus ,&_fd );_ea !=nil {return _ea ;};case _d .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u0063r\u0065\u0061\u0074\u0065\u0064"}:_gg .Created =new (_ae .XSDAny );if _cbd :=d .DecodeElement (_gg .Created ,&_fd );_cbd !=nil {return _cbd ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0063r\u0065\u0061\u0074\u006f\u0072"}:_gg .Creator =new (_ae .XSDAny );if _egf :=d .DecodeElement (_gg .Creator ,&_fd );_egf !=nil {return _egf ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"d\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e"}:_gg .Description =new (_ae .XSDAny );if _dc :=d .DecodeElement (_gg .Description ,&_fd );_dc !=nil {return _dc ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}:_gg .Identifier =new (_ae .XSDAny );if _eab :=d .DecodeElement (_gg .Identifier ,&_fd );_eab !=nil {return _eab ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}:_gg .Keywords =NewCT_Keywords ();if _df :=d .DecodeElement (_gg .Keywords ,&_fd );_df !=nil {return _df ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}:_gg .Language =new (_ae .XSDAny );if _ca :=d .DecodeElement (_gg .Language ,&_fd );_ca !=nil {return _ca ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}:_gg .LastModifiedBy =new (string );if _bef :=d .DecodeElement (_gg .LastModifiedBy ,&_fd );_bef !=nil {return _bef ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"l\u0061\u0073\u0074\u0050\u0072\u0069\u006e\u0074\u0065\u0064"}:_gg .LastPrinted =new (_a .Time );if _cac :=d .DecodeElement (_gg .LastPrinted ,&_fd );_cac !=nil {return _cac ;};case _d .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}:_gg .Modified =new (_ae .XSDAny );if _bfe :=d .DecodeElement (_gg .Modified ,&_fd );_bfe !=nil {return _bfe ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}:_gg .Revision =new (string );if _cbb :=d .DecodeElement (_gg .Revision ,&_fd );_cbb !=nil {return _cbb ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0073u\u0062\u006a\u0065\u0063\u0074"}:_gg .Subject =new (_ae .XSDAny );if _bg :=d .DecodeElement (_gg .Subject ,&_fd );_bg !=nil {return _bg ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0074\u0069\u0074l\u0065"}:_gg .Title =new (_ae .XSDAny );if _de :=d .DecodeElement (_gg .Title ,&_fd );_de !=nil {return _de ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0076e\u0072\u0073\u0069\u006f\u006e"}:_gg .Version =new (string );if _bec :=d .DecodeElement (_gg .Version ,&_fd );_bec !=nil {return _bec ;};default:_b .Log .Debug ("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u0020\u0025\u0076",_fd .Name );if _abc :=d .Skip ();_abc !=nil {return _abc ;};};case _d .EndElement :break _bf ;case _d .CharData :};};return nil ;};func (_fa *CT_Keyword )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for _ ,_ed :=range start .Attr {if _ed .Name .Space =="\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"&&_ed .Name .Local =="\u006c\u0061\u006e\u0067"{_abf ,_aa :=_ed .Value ,error (nil );if _aa !=nil {return _aa ;};_fa .LangAttr =&_abf ;continue ;};};for {_ce ,_bed :=d .Token ();if _bed !=nil {return _db .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u003a\u0020%\u0073",_bed );};if _ec ,_beb :=_ce .(_d .CharData );_beb {_fa .Content =string (_ec );};if _dbc ,_fga :=_ce .(_d .EndElement );_fga &&_dbc .Name ==start .Name {break ;};};return nil ;};func NewCT_Keyword ()*CT_Keyword {_efb :=&CT_Keyword {};return _efb };type CoreProperties struct{CT_CoreProperties };type CT_Keyword struct{LangAttr *string ;Content string ;};func (_ced *CoreProperties )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_ced .CT_CoreProperties =*NewCT_CoreProperties ();_eea :for {_deb ,_abb :=d .Token ();if _abb !=nil {return _abb ;};switch _bd :=_deb .(type ){case _d .StartElement :switch _bd .Name {case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}:_ced .Category =new (string );if _cae :=d .DecodeElement (_ced .Category ,&_bd );_cae !=nil {return _cae ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}:_ced .ContentStatus =new (string );if _geb :=d .DecodeElement (_ced .ContentStatus ,&_bd );_geb !=nil {return _geb ;};case _d .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u0063r\u0065\u0061\u0074\u0065\u0064"}:_ced .Created =new (_ae .XSDAny );if _cbdc :=d .DecodeElement (_ced .Created ,&_bd );_cbdc !=nil {return _cbdc ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0063r\u0065\u0061\u0074\u006f\u0072"}:_ced .Creator =new (_ae .XSDAny );if _dbg :=d .DecodeElement (_ced .Creator ,&_bd );_dbg !=nil {return _dbg ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"d\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e"}:_ced .Description =new (_ae .XSDAny );if _fgaf :=d .DecodeElement (_ced .Description ,&_bd );_fgaf !=nil {return _fgaf ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}:_ced .Identifier =new (_ae .XSDAny );if _fbe :=d .DecodeElement (_ced .Identifier ,&_bd );_fbe !=nil {return _fbe ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}:_ced .Keywords =NewCT_Keywords ();if _afg :=d .DecodeElement (_ced .Keywords ,&_bd );_afg !=nil {return _afg ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}:_ced .Language =new (_ae .XSDAny );if _geeb :=d .DecodeElement (_ced .Language ,&_bd );_geeb !=nil {return _geeb ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}:_ced .LastModifiedBy =new (string );if _fcaf :=d .DecodeElement (_ced .LastModifiedBy ,&_bd );_fcaf !=nil {return _fcaf ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"l\u0061\u0073\u0074\u0050\u0072\u0069\u006e\u0074\u0065\u0064"}:_ced .LastPrinted =new (_a .Time );if _eca :=d .DecodeElement (_ced .LastPrinted ,&_bd );_eca !=nil {return _eca ;};case _d .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}:_ced .Modified =new (_ae .XSDAny );if _bebad :=d .DecodeElement (_ced .Modified ,&_bd );_bebad !=nil {return _bebad ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}:_ced .Revision =new (string );if _bgg :=d .DecodeElement (_ced .Revision ,&_bd );_bgg !=nil {return _bgg ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0073u\u0062\u006a\u0065\u0063\u0074"}:_ced .Subject =new (_ae .XSDAny );if _abd :=d .DecodeElement (_ced .Subject ,&_bd );_abd !=nil {return _abd ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0074\u0069\u0074l\u0065"}:_ced .Title =new (_ae .XSDAny );if _gf :=d .DecodeElement (_ced .Title ,&_bd );_gf !=nil {return _gf ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0076e\u0072\u0073\u0069\u006f\u006e"}:_ced .Version =new (string );if _ece :=d .DecodeElement (_ced .Version ,&_bd );_ece !=nil {return _ece ;};default:_b .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065\u0072t\u0069e\u0073\u0020\u0025\u0076",_bd .Name );if _fad :=d .Skip ();_fad !=nil {return _fad ;};};case _d .EndElement :break _eea ;case _d .CharData :};};return nil ;};

// ValidateWithPath validates the CT_CoreProperties and its children, prefixing error messages with path
func (_ge *CT_CoreProperties )ValidateWithPath (path string )error {if _ge .Keywords !=nil {if _fb :=_ge .Keywords .ValidateWithPath (path +"\u002fK\u0065\u0079\u0077\u006f\u0072\u0064s");_fb !=nil {return _fb ;};};return nil ;};

// ValidateWithPath validates the CoreProperties and its children, prefixing error messages with path
func (_feg *CoreProperties )ValidateWithPath (path string )error {if _fea :=_feg .CT_CoreProperties .ValidateWithPath (path );_fea !=nil {return _fea ;};return nil ;};func NewCT_Keywords ()*CT_Keywords {_egd :=&CT_Keywords {};return _egd };

// Validate validates the CT_Keyword and its children
func (_aaa *CT_Keyword )Validate ()error {return _aaa .ValidateWithPath ("\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064");};

// Validate validates the CoreProperties and its children
func (_cda *CoreProperties )Validate ()error {return _cda .ValidateWithPath ("\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073");};type CT_Keywords struct{LangAttr *string ;Value []*CT_Keyword ;};func (_fdg *CT_Keyword )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {if _fdg .LangAttr !=nil {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078\u006d\u006c\u003a\u006c\u0061\u006e\u0067"},Value :_db .Sprintf ("\u0025\u0076",*_fdg .LangAttr )});};e .EncodeElement (_fdg .Content ,start );e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func NewCoreProperties ()*CoreProperties {_cc :=&CoreProperties {};_cc .CT_CoreProperties =*NewCT_CoreProperties ();return _cc ;};

// Validate validates the CT_CoreProperties and its children
func (_efd *CT_CoreProperties )Validate ()error {return _efd .ValidateWithPath ("\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073");};type CT_CoreProperties struct{Category *string ;ContentStatus *string ;Created *_ae .XSDAny ;Creator *_ae .XSDAny ;Description *_ae .XSDAny ;Identifier *_ae .XSDAny ;Keywords *CT_Keywords ;Language *_ae .XSDAny ;LastModifiedBy *string ;LastPrinted *_a .Time ;Modified *_ae .XSDAny ;Revision *string ;Subject *_ae .XSDAny ;Title *_ae .XSDAny ;Version *string ;};func (_aac *CT_Keywords )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {if _aac .LangAttr !=nil {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078\u006d\u006c\u003a\u006c\u0061\u006e\u0067"},Value :_db .Sprintf ("\u0025\u0076",*_aac .LangAttr )});};e .EncodeToken (start );if _aac .Value !=nil {_gee :=_d .StartElement {Name :_d .Name {Local :"\u0063\u0070\u003a\u0076\u0061\u006c\u0075\u0065"}};for _ ,_eag :=range _aac .Value {e .EncodeElement (_eag ,_gee );};};e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};

// Validate validates the CT_Keywords and its children
func (_ga *CT_Keywords )Validate ()error {return _ga .ValidateWithPath ("C\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073");};

// ValidateWithPath validates the CT_Keyword and its children, prefixing error messages with path
func (_eae *CT_Keyword )ValidateWithPath (path string )error {return nil };func init (){_ae .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073",NewCT_CoreProperties );_ae .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","C\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073",NewCT_Keywords );_ae .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064",NewCT_Keyword );_ae .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","\u0063\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073",NewCoreProperties );};