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

package vmldrawing ;import (_dg "encoding/xml";_d "fmt";_dd "github.com/unidoc/unioffice";_ca "github.com/unidoc/unioffice/common/logger";_ab "github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes";_c "github.com/unidoc/unioffice/schema/urn/schemas_microsoft_com/office/excel";_g "github.com/unidoc/unioffice/schema/urn/schemas_microsoft_com/vml";_f "strconv";_a "strings";);

// FontSize returns fontSize of the text.
func (_agb *TextpathStyle )FontSize ()int64 {return _agb ._fgd };

// MSOPositionVerticalRelative get `mso-position-vertical-relative` attribute of shape style.
func (_ddb *ShapeStyle )MSOPositionVerticalRelative ()string {return _ddb ._gee };type Container struct{Layout *_g .OfcShapelayout ;ShapeType *_g .Shapetype ;Shape []*_g .Shape ;};

// FontFamily returns fontFamily of the text.
func (_ae *TextpathStyle )FontFamily ()string {return _ae ._ef };func (_fc *Container )UnmarshalXML (d *_dg .Decoder ,start _dg .StartElement )error {_fc .Shape =nil ;_fd :for {_df ,_gb :=d .Token ();if _gb !=nil {return _gb ;};switch _dde :=_df .(type ){case _dg .StartElement :switch _dde .Name .Local {case "s\u0068\u0061\u0070\u0065\u006c\u0061\u0079\u006f\u0075\u0074":_fc .Layout =_g .NewOfcShapelayout ();if _dfc :=d .DecodeElement (_fc .Layout ,&_dde );_dfc !=nil {return _dfc ;};case "\u0073h\u0061\u0070\u0065\u0074\u0079\u0070e":_fc .ShapeType =_g .NewShapetype ();if _dfd :=d .DecodeElement (_fc .ShapeType ,&_dde );_dfd !=nil {return _dfd ;};case "\u0073\u0068\u0061p\u0065":_gcb :=_g .NewShape ();if _cd :=d .DecodeElement (_gcb ,&_dde );_cd !=nil {return _cd ;};_fc .Shape =append (_fc .Shape ,_gcb );};case _dg .EndElement :break _fd ;};};return nil ;};const (ShapeStylePositionAbsolute ="\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065";ShapeStylePositionRelative ="\u0072\u0065\u006c\u0061\u0074\u0069\u0076\u0065";);

// MSOPositionHorizontalRelative get `mso-position-horizontal-relative` attribute of shape style.
func (_bae *ShapeStyle )MSOPositionHorizontalRelative ()string {return _bae ._cf };

// SetBold sets text to bold.
func (_ffd *TextpathStyle )SetBold (bold bool ){_ffd ._gff =bold };

// NewCommentShape creates a new comment shape for a given cell index.  The
// indices here are zero based.
func NewCommentShape (col ,row int64 )*_g .Shape {_dge :=_g .NewShape ();_dge .IdAttr =_dd .String (_d .Sprintf ("\u0063\u0073\u005f\u0025\u0064\u005f\u0025\u0064",col ,row ));_dge .TypeAttr =_dd .String ("\u0023\u005f\u00780\u0030\u0030\u0030\u005f\u0074\u0032\u0030\u0032");_dge .StyleAttr =_dd .String ("\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u003a\u0061\u0062\u0073\u006f\u006cu\u0074\u0065\u003b\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u006c\u0065\u0066\u0074:\u0038\u0030\u0070\u0074;\u006d\u0061\u0072\u0067\u0069n-\u0074o\u0070\u003a\u0032pt\u003b\u0077\u0069\u0064\u0074\u0068\u003a1\u0030\u0034\u0070\u0074\u003b\u0068\u0065\u0069\u0067\u0068\u0074\u003a\u0037\u0036\u0070\u0074\u003b\u007a\u002d\u0069\u006e\u0064\u0065x\u003a\u0031\u003bv\u0069\u0073\u0069\u0062\u0069\u006c\u0069t\u0079\u003a\u0068\u0069\u0064\u0064\u0065\u006e");_dge .FillcolorAttr =_dd .String ("\u0023f\u0062\u0066\u0036\u0064\u0036");_dge .StrokecolorAttr =_dd .String ("\u0023e\u0064\u0065\u0061\u0061\u0031");_ge :=_g .NewEG_ShapeElements ();_ge .Fill =_g .NewFill ();_ge .Fill .Color2Attr =_dd .String ("\u0023f\u0062\u0066\u0065\u0038\u0032");_ge .Fill .AngleAttr =_dd .Float64 (-180);_ge .Fill .TypeAttr =_g .ST_FillTypeGradient ;_ge .Fill .Fill =_g .NewOfcFill ();_ge .Fill .Fill .ExtAttr =_g .ST_ExtView ;_ge .Fill .Fill .TypeAttr =_g .OfcST_FillTypeGradientUnscaled ;_dge .EG_ShapeElements =append (_dge .EG_ShapeElements ,_ge );_gf :=_g .NewEG_ShapeElements ();_gf .Shadow =_g .NewShadow ();_gf .Shadow .OnAttr =_ab .ST_TrueFalseT ;_gf .Shadow .ObscuredAttr =_ab .ST_TrueFalseT ;_dge .EG_ShapeElements =append (_dge .EG_ShapeElements ,_gf );_bb :=_g .NewEG_ShapeElements ();_bb .Path =_g .NewPath ();_bb .Path .ConnecttypeAttr =_g .OfcST_ConnectTypeNone ;_dge .EG_ShapeElements =append (_dge .EG_ShapeElements ,_bb );_ff :=_g .NewEG_ShapeElements ();_ff .Textbox =_g .NewTextbox ();_ff .Textbox .StyleAttr =_dd .String ("\u006d\u0073\u006f\u002ddi\u0072\u0065\u0063\u0074\u0069\u006f\u006e\u002d\u0061\u006c\u0074\u003a\u0061\u0075t\u006f");_dge .EG_ShapeElements =append (_dge .EG_ShapeElements ,_ff );_gc :=_g .NewEG_ShapeElements ();_gc .ClientData =_c .NewClientData ();_gc .ClientData .ObjectTypeAttr =_c .ST_ObjectTypeNote ;_gc .ClientData .MoveWithCells =_ab .ST_TrueFalseBlankT ;_gc .ClientData .SizeWithCells =_ab .ST_TrueFalseBlankT ;_gc .ClientData .Anchor =_dd .String ("\u0031,\u0020\u0031\u0035\u002c\u0020\u0030\u002c\u0020\u0032\u002c\u00202\u002c\u0020\u0035\u0034\u002c\u0020\u0035\u002c\u0020\u0033");_gc .ClientData .AutoFill =_ab .ST_TrueFalseBlankFalse ;_gc .ClientData .Row =_dd .Int64 (row );_gc .ClientData .Column =_dd .Int64 (col );_dge .EG_ShapeElements =append (_dge .EG_ShapeElements ,_gc );return _dge ;};

// NewShapeStyle accept value of string style attribute in v:shape and format it to generate ShapeStyle.
func NewShapeStyle (style string )ShapeStyle {_cgec :=ShapeStyle {_bfe :0,_bac :0};_gd :=_a .Split (style ,"\u003b");for _ ,_cgb :=range _gd {_da :=_a .Split (_cgb ,"\u003a");if len (_da )!=2{continue ;};var _bcc error ;switch _da [0]{case "\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e":_cgec ._bf =_da [1];break ;case "\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u0074\u006f\u0070":_cgec ._fe ,_bcc =_f .ParseFloat (_a .ReplaceAll (_da [1],"\u0070\u0074",""),64);break ;case "m\u0061\u0072\u0067\u0069\u006e\u002d\u006c\u0065\u0066\u0074":_cgec ._ba ,_bcc =_f .ParseFloat (_a .ReplaceAll (_da [1],"\u0070\u0074",""),64);break ;case "\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u0062\u006f\u0074\u0074\u006f\u006d":_cgec ._fdg ,_bcc =_f .ParseFloat (_a .ReplaceAll (_da [1],"\u0070\u0074",""),64);break ;case "\u006d\u0061\u0072g\u0069\u006e\u002d\u0072\u0069\u0067\u0068\u0074":_cgec ._gcc ,_bcc =_f .ParseFloat (_a .ReplaceAll (_da [1],"\u0070\u0074",""),64);break ;case "\u0074\u006f\u0070":_cgec ._gfd ,_bcc =_f .ParseFloat (_a .ReplaceAll (_da [1],"\u0070\u0074",""),64);break ;case "\u006c\u0065\u0066\u0074":_cgec ._gbc ,_bcc =_f .ParseFloat (_a .ReplaceAll (_da [1],"\u0070\u0074",""),64);break ;case "\u0062\u006f\u0074\u0074\u006f\u006d":_cgec ._e ,_bcc =_f .ParseFloat (_a .ReplaceAll (_da [1],"\u0070\u0074",""),64);break ;case "\u0072\u0069\u0067h\u0074":_cgec ._fa ,_bcc =_f .ParseFloat (_a .ReplaceAll (_da [1],"\u0070\u0074",""),64);break ;case "\u0077\u0069\u0064t\u0068":_cgec ._bfe ,_bcc =_f .ParseFloat (_a .ReplaceAll (_da [1],"\u0070\u0074",""),64);break ;case "\u0068\u0065\u0069\u0067\u0068\u0074":_cgec ._bac ,_bcc =_f .ParseFloat (_a .ReplaceAll (_da [1],"\u0070\u0074",""),64);break ;case "\u007a-\u0069\u006e\u0064\u0065\u0078":_cgec ._fce ,_bcc =_f .ParseInt (_da [1],10,64);break ;case "\u006d\u0073\u006f-p\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c":_cgec ._abc =_da [1];break ;case "\u006d\u0073\u006f\u002d\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0068\u006fr\u0069z\u006f\u006e\u0074\u0061\u006c\u002d\u0072\u0065\u006c\u0061\u0074\u0069\u0076\u0065":_cgec ._cf =_da [1];break ;case "m\u0073\u006f\u002d\u0070os\u0069t\u0069\u006f\u006e\u002d\u0076e\u0072\u0074\u0069\u0063\u0061\u006c":_cgec ._db =_da [1];break ;case "\u006d\u0073\u006f\u002d\u0070\u006f\u0073\u0069\u0074\u0069o\u006e\u002d\u0076\u0065\u0072\u0074\u0069c\u0061\u006c\u002d\u0072\u0065\u006c\u0061\u0074\u0069\u0076\u0065":_cgec ._gee =_da [1];break ;};if _bcc !=nil {_ca .Log .Debug ("\u0055n\u0061\u0062l\u0065\u0020\u0074o\u0020\u0070\u0061\u0072\u0073\u0065\u0020s\u0074\u0079\u006c\u0065\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0025\u0073 \u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_da [0],_da [1]);};};return _cgec ;};

// SetFontSize sets text's fontSize.
func (_be *TextpathStyle )SetFontSize (fontSize int64 ){_be ._fgd =fontSize };

// Top get top attribute of shape style.
func (_ee *ShapeStyle )Top ()float64 {return _ee ._gfd };

// IsItalic returns true if text is italic.
func (_gcd *TextpathStyle )IsItalic ()bool {return _gcd ._dba };

// Margins get margin top, left, bottom, and right of shape style.
func (_bbc *ShapeStyle )Margins ()(float64 ,float64 ,float64 ,float64 ){return _bbc ._fe ,_bbc ._ba ,_bbc ._fdg ,_bbc ._gcc ;};

// Width return width of shape.
func (_eb *ShapeStyle )Width ()float64 {return _eb ._bfe };

// IsBold returns true if text is bold.
func (_agab *TextpathStyle )IsBold ()bool {return _agab ._gff };

// Right get right attribute of shape style.
func (_eec *ShapeStyle )Right ()float64 {return _eec ._fa };func (_dgd *Container )MarshalXML (e *_dg .Encoder ,start _dg .StartElement )error {start .Attr =append (start .Attr ,_dg .Attr {Name :_dg .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0076"},Value :"\u0075\u0072n\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d:v\u006d\u006c"});start .Attr =append (start .Attr ,_dg .Attr {Name :_dg .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u006f"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006di\u0063\u0072\u006f\u0073\u006f\u0066t\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u006ff\u0066\u0069\u0063\u0065"});start .Attr =append (start .Attr ,_dg .Attr {Name :_dg .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c"});start .Name .Local ="\u0078\u006d\u006c";e .EncodeToken (start );if _dgd .Layout !=nil {_ag :=_dg .StartElement {Name :_dg .Name {Local :"\u006f\u003a\u0073\u0068\u0061\u0070\u0065\u006c\u0061\u0079\u006f\u0075\u0074"}};e .EncodeElement (_dgd .Layout ,_ag );};if _dgd .ShapeType !=nil {_ad :=_dg .StartElement {Name :_dg .Name {Local :"v\u003a\u0073\u0068\u0061\u0070\u0065\u0074\u0079\u0070\u0065"}};e .EncodeElement (_dgd .ShapeType ,_ad );};for _ ,_cge :=range _dgd .Shape {_fg :=_dg .StartElement {Name :_dg .Name {Local :"\u0076:\u0073\u0068\u0061\u0070\u0065"}};e .EncodeElement (_cge ,_fg );};return e .EncodeToken (_dg .EndElement {Name :start .Name });};

// ToString formatting ShapeStyle to string.
func (_fdf *ShapeStyle )String ()string {_geed :="";_geed +=_d .Sprintf ("\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u003a\u0025\u0073\u003b",_fdf ._bf );_geed +=_d .Sprintf ("\u006da\u0072g\u0069\u006e\u002d\u006c\u0065\u0066\u0074\u003a\u0025\u0064\u003b",int64 (_fdf ._ba ));_geed +=_d .Sprintf ("\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u0074\u006fp\u003a\u0025\u0064\u003b",int64 (_fdf ._fe ));_geed +=_d .Sprintf ("w\u0069\u0064\u0074\u0068\u003a\u0025\u0064\u0070\u0074\u003b",int64 (_fdf ._bfe ));_geed +=_d .Sprintf ("\u0068\u0065\u0069g\u0068\u0074\u003a\u0025\u0064\u0070\u0074\u003b",int64 (_fdf ._bac ));_geed +=_d .Sprintf ("z\u002d\u0069\u006e\u0064\u0065\u0078\u003a\u0025\u0064\u003b",_fdf ._fce );_geed +=_d .Sprintf ("m\u0073\u006f\u002d\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u002d\u0068\u006f\u0072\u0069z\u006f\u006e\u0074a\u006c:\u0025\u0073\u003b",_fdf ._abc );_geed +=_d .Sprintf ("\u006d\u0073o-\u0070\u006f\u0073i\u0074\u0069\u006f\u006e-ho\u0072iz\u006f\u006e\u0074\u0061\u006c\u002d\u0072el\u0061\u0074\u0069\u0076\u0065\u003a\u0025s\u003b",_fdf ._cf );_geed +=_d .Sprintf ("\u006ds\u006f\u002d\u0070\u006fs\u0069\u0074\u0069\u006f\u006e-\u0076e\u0072t\u0069\u0063\u0061\u006c\u003a\u0025\u0073;",_fdf ._db );_geed +=_d .Sprintf ("\u006d\u0073\u006f-p\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0076e\u0072t\u0069c\u0061l\u002d\u0072\u0065\u006c\u0061\u0074\u0069\u0076\u0065\u003a\u0025\u0073\u003b",_fdf ._gee );return _geed ;};

// Bottom get bottom attribute of shape style.
func (_gca *ShapeStyle )Bottom ()float64 {return _gca ._e };

// SetItalic sets text to italic.
func (_geef *TextpathStyle )SetItalic (italic bool ){_geef ._dba =italic };

// NewTextpathStyle accept value of string style attribute of element v:textpath and format it to generate TextpathStyle.
func NewTextpathStyle (style string )TextpathStyle {_bbb :=TextpathStyle {_ef :"\u0022C\u0061\u006c\u0069\u0062\u0072\u0069\"",_fgd :44,_gff :false ,_dba :false };_fef :=_a .Split (style ,"\u003b");for _ ,_eecg :=range _fef {_dfg :=_a .Split (_eecg ,"\u003a");if len (_dfg )!=2{continue ;};switch _dfg [0]{case "f\u006f\u006e\u0074\u002d\u0066\u0061\u006d\u0069\u006c\u0079":_bbb ._ef =_dfg [1];break ;case "\u0066o\u006e\u0074\u002d\u0073\u0069\u007ae":_bbb ._fgd ,_ =_f .ParseInt (_a .ReplaceAll (_dfg [1],"\u0070\u0074",""),10,64);break ;case "f\u006f\u006e\u0074\u002d\u0077\u0065\u0069\u0067\u0068\u0074":_bbb ._gff =_dfg [1]=="\u0062\u006f\u006c\u0064";break ;case "\u0066\u006f\u006e\u0074\u002d\u0073\u0074\u0079\u006c\u0065":_bbb ._dba =_dfg [1]=="\u0069\u0074\u0061\u006c\u0069\u0063";break ;};};return _bbb ;};

// CreateFormula creates F element for typeFormulas.
func CreateFormula (s string )*_g .CT_F {_gdf :=_g .NewCT_F ();_gdf .EqnAttr =&s ;return _gdf };func NewContainer ()*Container {return &Container {}};

// SetWidth set width of shape.
func (_gg *ShapeStyle )SetWidth (width float64 ){_gg ._bfe =width };

// SetFontFamily sets text's fontFamily.
func (_fdgf *TextpathStyle )SetFontFamily (fontFamily string ){_fdgf ._ef =fontFamily };

// NewCommentDrawing constructs a new comment drawing.
func NewCommentDrawing ()*Container {_ce :=NewContainer ();_ce .Layout =_g .NewOfcShapelayout ();_ce .Layout .ExtAttr =_g .ST_ExtEdit ;_ce .Layout .Idmap =_g .NewOfcCT_IdMap ();_ce .Layout .Idmap .DataAttr =_dd .String ("\u0031");_ce .Layout .Idmap .ExtAttr =_g .ST_ExtEdit ;_ce .ShapeType =_g .NewShapetype ();_ce .ShapeType .IdAttr =_dd .String ("_\u0078\u0030\u0030\u0030\u0030\u005f\u0074\u0032\u0030\u0032");_ce .ShapeType .CoordsizeAttr =_dd .String ("2\u0031\u0036\u0030\u0030\u002c\u0032\u0031\u0036\u0030\u0030");_ce .ShapeType .SptAttr =_dd .Float32 (202);_ce .ShapeType .PathAttr =_dd .String ("\u006d\u0030\u002c0l\u0030\u002c\u0032\u0031\u0036\u0030\u0030\u002c\u00321\u00360\u0030,\u00321\u0036\u0030\u0030\u002c\u0032\u0031\u0036\u0030\u0030\u002c\u0030\u0078\u0065");_bc :=_g .NewEG_ShapeElements ();_ce .ShapeType .EG_ShapeElements =append (_ce .ShapeType .EG_ShapeElements ,_bc );_bc .Path =_g .NewPath ();_bc .Path .GradientshapeokAttr =_ab .ST_TrueFalseT ;_bc .Path .ConnecttypeAttr =_g .OfcST_ConnectTypeRect ;return _ce ;};

// ShapeStyle is style attribute of v:shape element.
type ShapeStyle struct{_bf string ;_fe float64 ;_ba float64 ;_fdg float64 ;_gcc float64 ;_gfd float64 ;_gbc float64 ;_e float64 ;_fa float64 ;_bfe float64 ;_bac float64 ;_fce int64 ;_abc string ;_cf string ;_db string ;_gee string ;};

// SetHeight set height of shape.
func (_bfa *ShapeStyle )SetHeight (height float64 ){_bfa ._bac =height };

// Position get position attribute of shape style.
func (_aga *ShapeStyle )Position ()string {return _aga ._bf };

// Height return height of shape.
func (_ga *ShapeStyle )Height ()float64 {return _ga ._bac };

// Left get left attribute of shape style.
func (_ac *ShapeStyle )Left ()float64 {return _ac ._gbc };

// ToString generate string of TextpathStyle.
func (_bacc *TextpathStyle )String ()string {_fb :="";_fb +=_d .Sprintf ("\u0066o\u006et\u002d\u0066\u0061\u006d\u0069\u006c\u0079\u003a\u0025\u0073\u003b",_bacc ._ef );_fb +=_d .Sprintf ("\u0066o\u006et\u002d\u0073\u0069\u007a\u0065\u003a\u0025\u0064\u0070\u0074\u003b",_bacc ._fgd );if _bacc ._dba {_fb +=_d .Sprintf ("\u0066o\u006et\u002d\u0073\u0074\u0079\u006ce\u003a\u0069t\u0061\u006c\u0069\u0063\u003b");};if _bacc ._gff {_fb +=_d .Sprintf ("\u0066\u006f\u006e\u0074\u002d\u0077\u0065\u0069\u0067\u0068\u0074\u003ab\u006f\u006c\u0064\u003b");};return _fb ;};

// TextpathStyle is style attribute of element v:textpath.
type TextpathStyle struct{_ef string ;_fgd int64 ;_gff bool ;_dba bool ;};