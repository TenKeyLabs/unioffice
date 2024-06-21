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

package drawing ;import (_f "github.com/unidoc/unioffice";_c "github.com/unidoc/unioffice/color";_d "github.com/unidoc/unioffice/measurement";_e "github.com/unidoc/unioffice/schema/soo/dml";);

// MakeRunProperties constructs a new RunProperties wrapper.
func MakeRunProperties (x *_e .CT_TextCharacterProperties )RunProperties {return RunProperties {x }};

// RunProperties controls the run properties.
type RunProperties struct{_egc *_e .CT_TextCharacterProperties ;};

// GetPosition gets the position of the shape in EMU.
func (_fab ShapeProperties )GetPosition ()(int64 ,int64 ){_fab .ensureXfrm ();if _fab ._bgf .Xfrm .Off ==nil {_fab ._bgf .Xfrm .Off =_e .NewCT_Point2D ();};return *_fab ._bgf .Xfrm .Off .XAttr .ST_CoordinateUnqualified ,*_fab ._bgf .Xfrm .Off .YAttr .ST_CoordinateUnqualified ;
};

// SetFont controls the font of a run.
func (_fac RunProperties )SetFont (s string ){_fac ._egc .Latin =_e .NewCT_TextFont ();_fac ._egc .Latin .TypefaceAttr =s ;};

// AddRun adds a new run to a paragraph.
func (_ga Paragraph )AddRun ()Run {_cc :=MakeRun (_e .NewEG_TextRun ());_ga ._bg .EG_TextRun =append (_ga ._bg .EG_TextRun ,_cc .X ());return _cc ;};func (_eg LineProperties )clearFill (){_eg ._cf .NoFill =nil ;_eg ._cf .GradFill =nil ;_eg ._cf .SolidFill =nil ;
_eg ._cf .PattFill =nil ;};

// LineJoin is the type of line join
type LineJoin byte ;type LineProperties struct{_cf *_e .CT_LineProperties };

// SetFlipVertical controls if the shape is flipped vertically.
func (_dg ShapeProperties )SetFlipVertical (b bool ){_dg .ensureXfrm ();if !b {_dg ._bgf .Xfrm .FlipVAttr =nil ;}else {_dg ._bgf .Xfrm .FlipVAttr =_f .Bool (true );};};

// Properties returns the run's properties.
func (_gc Run )Properties ()RunProperties {if _gc ._dbf .R ==nil {_gc ._dbf .R =_e .NewCT_RegularTextRun ();};if _gc ._dbf .R .RPr ==nil {_gc ._dbf .R .RPr =_e .NewCT_TextCharacterProperties ();};return RunProperties {_gc ._dbf .R .RPr };};

// SetBold controls the bolding of a run.
func (_ab RunProperties )SetBold (b bool ){_ab ._egc .BAttr =_f .Bool (b )};func (_dd ShapeProperties )LineProperties ()LineProperties {if _dd ._bgf .Ln ==nil {_dd ._bgf .Ln =_e .NewCT_LineProperties ();};return LineProperties {_dd ._bgf .Ln };};

// SetWidth sets the line width, MS products treat zero as the minimum width
// that can be displayed.
func (_b LineProperties )SetWidth (w _d .Distance ){_b ._cf .WAttr =_f .Int32 (int32 (w /_d .EMU ))};

// X returns the inner wrapped XML type.
func (_afe Run )X ()*_e .EG_TextRun {return _afe ._dbf };

// MakeParagraph constructs a new paragraph wrapper.
func MakeParagraph (x *_e .CT_TextParagraph )Paragraph {return Paragraph {x }};

// MakeRun constructs a new Run wrapper.
func MakeRun (x *_e .EG_TextRun )Run {return Run {x }};

// SetJoin sets the line join style.
func (_fb LineProperties )SetJoin (e LineJoin ){_fb ._cf .Round =nil ;_fb ._cf .Miter =nil ;_fb ._cf .Bevel =nil ;switch e {case LineJoinRound :_fb ._cf .Round =_e .NewCT_LineJoinRound ();case LineJoinBevel :_fb ._cf .Bevel =_e .NewCT_LineJoinBevel ();
case LineJoinMiter :_fb ._cf .Miter =_e .NewCT_LineJoinMiterProperties ();};};func (_cff LineProperties )SetSolidFill (c _c .Color ){_cff .clearFill ();_cff ._cf .SolidFill =_e .NewCT_SolidColorFillProperties ();_cff ._cf .SolidFill .SrgbClr =_e .NewCT_SRgbColor ();
_cff ._cf .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// Paragraph is a paragraph within a document.
type Paragraph struct{_bg *_e .CT_TextParagraph };func (_da ShapeProperties )clearFill (){_da ._bgf .NoFill =nil ;_da ._bgf .BlipFill =nil ;_da ._bgf .GradFill =nil ;_da ._bgf .GrpFill =nil ;_da ._bgf .SolidFill =nil ;_da ._bgf .PattFill =nil ;};func (_ec LineProperties )SetNoFill (){_ec .clearFill ();
_ec ._cf .NoFill =_e .NewCT_NoFillProperties ()};

// SetWidth sets the width of the shape.
func (_aag ShapeProperties )SetWidth (w _d .Distance ){_aag .ensureXfrm ();if _aag ._bgf .Xfrm .Ext ==nil {_aag ._bgf .Xfrm .Ext =_e .NewCT_PositiveSize2D ();};_aag ._bgf .Xfrm .Ext .CxAttr =int64 (w /_d .EMU );};

// SetSize sets the width and height of the shape.
func (_bb ShapeProperties )SetSize (w ,h _d .Distance ){_bb .SetWidth (w );_bb .SetHeight (h )};

// SetSize sets the font size of the run text
func (_gb RunProperties )SetSize (sz _d .Distance ){_gb ._egc .SzAttr =_f .Int32 (int32 (sz /_d .HundredthPoint ));};

// SetAlign controls the paragraph alignment
func (_db ParagraphProperties )SetAlign (a _e .ST_TextAlignType ){_db ._fbe .AlgnAttr =a };

// SetSolidFill controls the text color of a run.
func (_cb RunProperties )SetSolidFill (c _c .Color ){_cb ._egc .NoFill =nil ;_cb ._egc .BlipFill =nil ;_cb ._egc .GradFill =nil ;_cb ._egc .GrpFill =nil ;_cb ._egc .PattFill =nil ;_cb ._egc .SolidFill =_e .NewCT_SolidColorFillProperties ();_cb ._egc .SolidFill .SrgbClr =_e .NewCT_SRgbColor ();
_cb ._egc .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// ParagraphProperties allows controlling paragraph properties.
type ParagraphProperties struct{_fbe *_e .CT_TextParagraphProperties ;};

// SetBulletFont controls the font for the bullet character.
func (_fd ParagraphProperties )SetBulletFont (f string ){if f ==""{_fd ._fbe .BuFont =nil ;}else {_fd ._fbe .BuFont =_e .NewCT_TextFont ();_fd ._fbe .BuFont .TypefaceAttr =f ;};};const (LineJoinRound LineJoin =iota ;LineJoinBevel ;LineJoinMiter ;);

// SetLevel sets the level of indentation of a paragraph.
func (_gf ParagraphProperties )SetLevel (idx int32 ){_gf ._fbe .LvlAttr =_f .Int32 (idx )};func (_ef ShapeProperties )SetSolidFill (c _c .Color ){_ef .clearFill ();_ef ._bgf .SolidFill =_e .NewCT_SolidColorFillProperties ();_ef ._bgf .SolidFill .SrgbClr =_e .NewCT_SRgbColor ();
_ef ._bgf .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// SetFlipHorizontal controls if the shape is flipped horizontally.
func (_ac ShapeProperties )SetFlipHorizontal (b bool ){_ac .ensureXfrm ();if !b {_ac ._bgf .Xfrm .FlipHAttr =nil ;}else {_ac ._bgf .Xfrm .FlipHAttr =_f .Bool (true );};};

// Run is a run within a paragraph.
type Run struct{_dbf *_e .EG_TextRun };

// X returns the inner wrapped XML type.
func (_aa Paragraph )X ()*_e .CT_TextParagraph {return _aa ._bg };

// SetNumbered controls if bullets are numbered or not.
func (_fa ParagraphProperties )SetNumbered (scheme _e .ST_TextAutonumberScheme ){if scheme ==_e .ST_TextAutonumberSchemeUnset {_fa ._fbe .BuAutoNum =nil ;}else {_fa ._fbe .BuAutoNum =_e .NewCT_TextAutonumberBullet ();_fa ._fbe .BuAutoNum .TypeAttr =scheme ;
};};

// SetBulletChar sets the bullet character for the paragraph.
func (_ea ParagraphProperties )SetBulletChar (c string ){if c ==""{_ea ._fbe .BuChar =nil ;}else {_ea ._fbe .BuChar =_e .NewCT_TextCharBullet ();_ea ._fbe .BuChar .CharAttr =c ;};};type ShapeProperties struct{_bgf *_e .CT_ShapeProperties };

// SetHeight sets the height of the shape.
func (_df ShapeProperties )SetHeight (h _d .Distance ){_df .ensureXfrm ();if _df ._bgf .Xfrm .Ext ==nil {_df ._bgf .Xfrm .Ext =_e .NewCT_PositiveSize2D ();};_df ._bgf .Xfrm .Ext .CyAttr =int64 (h /_d .EMU );};

// AddBreak adds a new line break to a paragraph.
func (_bc Paragraph )AddBreak (){_gg :=_e .NewEG_TextRun ();_gg .Br =_e .NewCT_TextLineBreak ();_bc ._bg .EG_TextRun =append (_bc ._bg .EG_TextRun ,_gg );};func MakeShapeProperties (x *_e .CT_ShapeProperties )ShapeProperties {return ShapeProperties {x }};


// Properties returns the paragraph properties.
func (_cfe Paragraph )Properties ()ParagraphProperties {if _cfe ._bg .PPr ==nil {_cfe ._bg .PPr =_e .NewCT_TextParagraphProperties ();};return MakeParagraphProperties (_cfe ._bg .PPr );};

// MakeParagraphProperties constructs a new ParagraphProperties wrapper.
func MakeParagraphProperties (x *_e .CT_TextParagraphProperties )ParagraphProperties {return ParagraphProperties {x };};func (_gae ShapeProperties )SetNoFill (){_gae .clearFill ();_gae ._bgf .NoFill =_e .NewCT_NoFillProperties ();};func (_bga ShapeProperties )ensureXfrm (){if _bga ._bgf .Xfrm ==nil {_bga ._bgf .Xfrm =_e .NewCT_Transform2D ();
};};

// SetPosition sets the position of the shape.
func (_faa ShapeProperties )SetPosition (x ,y _d .Distance ){_faa .ensureXfrm ();if _faa ._bgf .Xfrm .Off ==nil {_faa ._bgf .Xfrm .Off =_e .NewCT_Point2D ();};_faa ._bgf .Xfrm .Off .XAttr .ST_CoordinateUnqualified =_f .Int64 (int64 (x /_d .EMU ));_faa ._bgf .Xfrm .Off .YAttr .ST_CoordinateUnqualified =_f .Int64 (int64 (y /_d .EMU ));
};

// X returns the inner wrapped XML type.
func (_cbf ShapeProperties )X ()*_e .CT_ShapeProperties {return _cbf ._bgf };

// SetGeometry sets the shape type of the shape
func (_ggf ShapeProperties )SetGeometry (g _e .ST_ShapeType ){if _ggf ._bgf .PrstGeom ==nil {_ggf ._bgf .PrstGeom =_e .NewCT_PresetGeometry2D ();};_ggf ._bgf .PrstGeom .PrstAttr =g ;};

// SetText sets the run's text contents.
func (_ce Run )SetText (s string ){_ce ._dbf .Br =nil ;_ce ._dbf .Fld =nil ;if _ce ._dbf .R ==nil {_ce ._dbf .R =_e .NewCT_RegularTextRun ();};_ce ._dbf .R .T =s ;};

// X returns the inner wrapped XML type.
func (_af ParagraphProperties )X ()*_e .CT_TextParagraphProperties {return _af ._fbe };

// X returns the inner wrapped XML type.
func (_a LineProperties )X ()*_e .CT_LineProperties {return _a ._cf };