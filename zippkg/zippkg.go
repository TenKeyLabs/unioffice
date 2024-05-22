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

package zippkg ;import (_b "archive/zip";_f "bytes";_gb "encoding/xml";_g "fmt";_ag "github.com/unidoc/unioffice";_cg "github.com/unidoc/unioffice/algo";_ga "github.com/unidoc/unioffice/common/tempstorage";_db "github.com/unidoc/unioffice/schema/soo/pkg/relationships";
_cf "io";_d "path";_fb "path/filepath";_ed "sort";_a "strings";_e "time";);func (_agg SelfClosingWriter )Write (b []byte )(int ,error ){_ged :=0;_agbb :=0;for _dbf :=0;_dbf < len (b )-2;_dbf ++{if b [_dbf ]=='>'&&b [_dbf +1]=='<'&&b [_dbf +2]=='/'{_aca :=[]byte {};
_cdc :=_dbf ;for _fg :=_dbf ;_fg >=0;_fg --{if b [_fg ]==' '{_cdc =_fg ;}else if b [_fg ]=='<'{_aca =b [_fg +1:_cdc ];break ;};};_bd :=[]byte {};for _gba :=_dbf +3;_gba < len (b );_gba ++{if b [_gba ]=='>'{_bd =b [_dbf +3:_gba ];break ;};};if !_f .Equal (_aca ,_bd ){continue ;
};_dgb ,_aae :=_agg .W .Write (b [_ged :_dbf ]);if _aae !=nil {return _agbb +_dgb ,_aae ;};_agbb +=_dgb ;_ ,_aae =_agg .W .Write (_gad );if _aae !=nil {return _agbb ,_aae ;};_agbb +=3;for _eda :=_dbf +2;_eda < len (b )&&b [_eda ]!='>';_eda ++{_agbb ++;
_ged =_eda +2;_dbf =_ged ;};};};_aga ,_gfa :=_agg .W .Write (b [_ged :]);return _aga +_agbb ,_gfa ;};

// ExtractToDiskTmp extracts a zip file to a temporary file in a given path,
// returning the name of the file.
func ExtractToDiskTmp (f *_b .File ,path string )(string ,error ){_ega ,_da :=_ga .TempFile (path ,"\u007a\u007a");if _da !=nil {return "",_da ;};defer _ega .Close ();_dcc ,_da :=f .Open ();if _da !=nil {return "",_da ;};defer _dcc .Close ();_ ,_da =_cf .Copy (_ega ,_dcc );
if _da !=nil {return "",_da ;};return _ega .Name (),nil ;};type Target struct{Path string ;Typ string ;Ifc interface{};Index uint32 ;};

// SetOnNewRelationshipFunc sets the function to be called when a new
// relationship has been discovered.
func (_fe *DecodeMap )SetOnNewRelationshipFunc (fn OnNewRelationshipFunc ){_fe ._fbg =fn };func (_dbb *DecodeMap )RecordIndex (path string ,idx int ){_dbb ._aad [path ]=idx };func MarshalXMLByType (z *_b .Writer ,dt _ag .DocType ,typ string ,v interface{})error {_df :=_ag .AbsoluteFilename (dt ,typ ,0);
return MarshalXML (z ,_df ,v );};var _gad =[]byte {'/','>'};

// AddFileFromBytes takes a byte array and adds it at a given path to a zip file.
func AddFileFromBytes (z *_b .Writer ,zipPath string ,data []byte )error {_ff ,_afbg :=z .Create (zipPath );if _afbg !=nil {return _g .Errorf ("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",zipPath ,_afbg );
};_ ,_afbg =_cf .Copy (_ff ,_f .NewReader (data ));return _afbg ;};

// MarshalXML creates a file inside of a zip and marshals an object as xml, prefixing it
// with a standard XML header.
func MarshalXML (z *_b .Writer ,filename string ,v interface{})error {_age :=&_b .FileHeader {};_age .Method =_b .Deflate ;_age .Name =filename ;_age .SetModTime (_e .Now ());_fff ,_ggd :=z .CreateHeader (_age );if _ggd !=nil {return _g .Errorf ("\u0063\u0072\u0065\u0061ti\u006e\u0067\u0020\u0025\u0073\u0020\u0069\u006e\u0020\u007a\u0069\u0070\u003a\u0020%\u0073",filename ,_ggd );
};_ ,_ggd =_fff .Write ([]byte (XMLHeader ));if _ggd !=nil {return _g .Errorf ("\u0063\u0072e\u0061\u0074\u0069\u006e\u0067\u0020\u0078\u006d\u006c\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0074\u006f\u0020\u0025\u0073: \u0025\u0073",filename ,_ggd );
};if _ggd =_gb .NewEncoder (SelfClosingWriter {_fff }).Encode (v );_ggd !=nil {return _g .Errorf ("\u006d\u0061\u0072\u0073\u0068\u0061\u006c\u0069\u006e\u0067\u0020\u0025s\u003a\u0020\u0025\u0073",filename ,_ggd );};_ ,_ggd =_fff .Write (_aaf );return _ggd ;
};

// OnNewRelationshipFunc is called when a new relationship has been discovered.
//
// target is a resolved path that takes into account the location of the
// relationships file source and should be the path in the zip file.
//
// files are passed so non-XML files that can't be handled by AddTarget can be
// decoded directly (e.g. images)
//
// rel is the actual relationship so its target can be modified if the source
// target doesn't match where unioffice will write the file (e.g. read in
// 'xl/worksheets/MyWorksheet.xml' and we'll write out
// 'xl/worksheets/sheet1.xml')
type OnNewRelationshipFunc func (_ec *DecodeMap ,_bf ,_eg string ,_gaa []*_b .File ,_fc *_db .Relationship ,_bb Target )error ;var _aaf =[]byte {'\r','\n'};

// RelationsPathFor returns the relations path for a given filename.
func RelationsPathFor (path string )string {_bcf :=_a .Split (path ,"\u002f");_abe :=_a .Join (_bcf [0:len (_bcf )-1],"\u002f");_abg :=_bcf [len (_bcf )-1];_abe +="\u002f_\u0072\u0065\u006c\u0073\u002f";_abg +="\u002e\u0072\u0065l\u0073";return _abe +_abg ;
};func (_cfc *DecodeMap )IndexFor (path string )int {return _cfc ._aad [path ]};

// AddTarget allows documents to register decode targets. Path is a path that
// will be found in the zip file and ifc is an XML element that the file will be
// unmarshaled to.  filePath is the absolute path to the target, ifc is the
// object to decode into, sourceFileType is the type of file that the reference
// was discovered in, and index is the index of the source file type.
func (_ef *DecodeMap )AddTarget (filePath string ,ifc interface{},sourceFileType string ,idx uint32 )bool {if _ef ._af ==nil {_ef ._af =make (map[string ]Target );_ef ._dg =make (map[*_db .Relationships ]string );_ef ._de =make (map[string ]struct{});_ef ._aad =make (map[string ]int );
};if _d .IsAbs (filePath ){filePath =_a .TrimPrefix (filePath ,"\u002f");};_dbe :=_d .Clean (filePath );if _ ,_bc :=_ef ._de [_dbe ];_bc {return false ;};_ef ._de [_dbe ]=struct{}{};_ef ._af [_dbe ]=Target {Path :_dbe ,Typ :sourceFileType ,Ifc :ifc ,Index :idx };
return true ;};

// Decode loops decoding targets registered with AddTarget and calling th
func (_dgg *DecodeMap )Decode (files []*_b .File )error {_dee :=1;for _dee > 0{for len (_dgg ._aa )> 0{_gg :=_dgg ._aa [0];_dgg ._aa =_dgg ._aa [1:];_ea :=_gg .Ifc .(*_db .Relationships );for _ ,_fa :=range _ea .Relationship {_ab :=_dgg ._dg [_ea ];if _fb .IsAbs (_fa .TargetAttr ){_fa .TargetAttr =_a .TrimPrefix (_fa .TargetAttr ,"\u002f");
if _a .HasPrefix (_fa .TargetAttr ,_ab ){_ab ="";};};_dgg ._fbg (_dgg ,_ab +_fa .TargetAttr ,_fa .TypeAttr ,files ,_fa ,_gg );};};for _afb ,_cd :=range files {if _cd ==nil {continue ;};if _bfc ,_fcd :=_dgg ._af [_cd .Name ];_fcd {delete (_dgg ._af ,_cd .Name );
if _eff :=Decode (_cd ,_bfc .Ifc );_eff !=nil {return _eff ;};files [_afb ]=nil ;if _ecb ,_dc :=_bfc .Ifc .(*_db .Relationships );_dc {_dgg ._aa =append (_dgg ._aa ,_bfc );_aab ,_ :=_d .Split (_d .Clean (_cd .Name +"\u002f\u002e\u002e\u002f"));_dgg ._dg [_ecb ]=_aab ;
_dee ++;};};};_dee --;};return nil ;};

// Decode unmarshals the content of a *zip.File as XML to a given destination.
func Decode (f *_b .File ,dest interface{})error {_fd ,_cb :=f .Open ();if _cb !=nil {return _g .Errorf ("e\u0072r\u006f\u0072\u0020\u0072\u0065\u0061\u0064\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073",f .Name ,_cb );};defer _fd .Close ();_abeb :=_gb .NewDecoder (_fd );
if _bcd :=_abeb .Decode (dest );_bcd !=nil {return _g .Errorf ("e\u0072\u0072\u006f\u0072 d\u0065c\u006f\u0064\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",f .Name ,_bcd );};if _cdg ,_aaa :=dest .(*_db .Relationships );_aaa {for _ge ,_ac :=range _cdg .Relationship {switch _ac .TypeAttr {case _ag .OfficeDocumentTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .OfficeDocumentType ;
case _ag .StylesTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .StylesType ;case _ag .ThemeTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .ThemeType ;case _ag .ControlTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .ControlType ;case _ag .SettingsTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .SettingsType ;
case _ag .ImageTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .ImageType ;case _ag .CommentsTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .CommentsType ;case _ag .ThumbnailTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .ThumbnailType ;case _ag .DrawingTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .DrawingType ;
case _ag .ChartTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .ChartType ;case _ag .ExtendedPropertiesTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .ExtendedPropertiesType ;case _ag .CustomXMLTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .CustomXMLType ;
case _ag .WorksheetTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .WorksheetType ;case _ag .SharedStringsTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .SharedStringsType ;case _ag .TableTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .TableType ;
case _ag .HeaderTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .HeaderType ;case _ag .FooterTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .FooterType ;case _ag .NumberingTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .NumberingType ;case _ag .FontTableTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .FontTableType ;
case _ag .WebSettingsTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .WebSettingsType ;case _ag .FootNotesTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .FootNotesType ;case _ag .EndNotesTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .EndNotesType ;
case _ag .SlideTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .SlideType ;case _ag .VMLDrawingTypeStrict :_cdg .Relationship [_ge ].TypeAttr =_ag .VMLDrawingType ;};};_ed .Slice (_cdg .Relationship ,func (_fcc ,_be int )bool {_gc :=_cdg .Relationship [_fcc ];
_fba :=_cdg .Relationship [_be ];return _cg .NaturalLess (_gc .IdAttr ,_fba .IdAttr );});};return nil ;};const XMLHeader ="\u003c\u003f\u0078\u006d\u006c\u0020\u0076e\u0072\u0073\u0069o\u006e\u003d\u00221\u002e\u0030\"\u0020\u0065\u006e\u0063\u006f\u0064i\u006eg=\u0022\u0055\u0054\u0046\u002d\u0038\u0022\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u006c\u006f\u006e\u0065\u003d\u0022\u0079\u0065\u0073\u0022\u003f\u003e"+"\u000a";


// DecodeMap is used to walk a tree of relationships, decoding files and passing
// control back to the document.
type DecodeMap struct{_af map[string ]Target ;_dg map[*_db .Relationships ]string ;_aa []Target ;_fbg OnNewRelationshipFunc ;_de map[string ]struct{};_aad map[string ]int ;};

// AddFileFromDisk reads a file from internal storage and adds it at a given path to a zip file.
// TODO: Rename to AddFileFromStorage in next major version release (v2).
// NOTE: If disk storage cannot be used, memory storage can be used instead by calling memstore.SetAsStorage().
func AddFileFromDisk (z *_b .Writer ,zipPath ,storagePath string )error {_agb ,_gf :=z .Create (zipPath );if _gf !=nil {return _g .Errorf ("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",zipPath ,_gf );
};_bca ,_gf :=_ga .Open (storagePath );if _gf !=nil {return _g .Errorf ("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073",storagePath ,_gf );};defer _bca .Close ();_ ,_gf =_cf .Copy (_agb ,_bca );return _gf ;
};

// SelfClosingWriter wraps a writer and replaces XML tags of the
// type <foo></foo> with <foo/>
type SelfClosingWriter struct{W _cf .Writer ;};func MarshalXMLByTypeIndex (z *_b .Writer ,dt _ag .DocType ,typ string ,idx int ,v interface{})error {_gfb :=_ag .AbsoluteFilename (dt ,typ ,idx );return MarshalXML (z ,_gfb ,v );};