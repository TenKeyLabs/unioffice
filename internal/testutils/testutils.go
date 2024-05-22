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

package testutils ;import (_df "crypto/md5";_af "encoding/hex";_e "encoding/json";_aa "errors";_acf "fmt";_eb "github.com/stretchr/testify/require";_bb "golang.org/x/image/font";_fg "golang.org/x/image/font/opentype";_gc "golang.org/x/image/math/fixed";
_ge "image";_dg "image/color";_bd "image/draw";_f "image/png";_baf "io";_bc "io/ioutil";_ba "log";_ac "math";_bdf "os";_g "os/exec";_dfe "path/filepath";_b "strings";_ab "sync";_a "testing";_db "time";);func (_bgg *ReferenceMap )Write (key string ,entry ReferenceEntry ){_bgg .Lock ();
defer _bgg .Unlock ();if _bgg ._fc ==nil {_bgg ._fc =map[string ]ReferenceEntry {};};_bgg ._fc [key ]=entry ;};func _dgb (_eda ,_ddf _ge .Rectangle )bool {return _eda .Min .X ==_ddf .Min .X &&_eda .Min .Y ==_ddf .Min .Y &&_eda .Max .X ==_ddf .Max .X &&_eda .Max .Y ==_ddf .Max .Y ;
};func (_aec *ReferenceMap )Read (key string )(ReferenceEntry ,bool ){_aec .Lock ();defer _aec .Unlock ();if _aec ._fc ==nil {return ReferenceEntry {},false ;};_eba ,_gee :=_aec ._fc [key ];return _eba ,_gee ;};func CompareImages (img1 ,img2 _ge .Image )(bool ,error ){_aaa :=img1 .Bounds ();
_abe :=0;for _bf :=0;_bf < _aaa .Size ().X ;_bf ++{for _ade :=0;_ade < _aaa .Size ().Y ;_ade ++{_bef ,_bfb ,_ceg ,_ :=img1 .At (_bf ,_ade ).RGBA ();_ecf ,_dcf ,_ef ,_ :=img2 .At (_bf ,_ade ).RGBA ();if _bef !=_ecf ||_bfb !=_dcf ||_ceg !=_ef {_abe ++;};
};};_ged :=float64 (_abe )/float64 (_aaa .Dx ()*_aaa .Dy ());if _ged > 0.0001{_acf .Printf ("\u0064\u0069\u0066f \u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076\u0020\u0028\u0025\u0064\u0029\u000a",_ged ,_abe );return false ,nil ;
};return true ,nil ;};func CopyFile (src ,dst string )error {_bdd ,_ea :=_bdf .Open (src );if _ea !=nil {return _ea ;};defer _bdd .Close ();_fd ,_ea :=_bdf .Create (dst );if _ea !=nil {return _ea ;};defer _fd .Close ();_ ,_ea =_baf .Copy (_fd ,_bdd );return _ea ;
};func RunRenderTest (t *_a .T ,pdfPath ,outputDir ,baselineRenderPath string ,saveBaseline bool ,currentHashMap *ReferenceMap ,refFile *ReferenceFile ){RunRenderOfficeTest (t ,pdfPath ,outputDir ,baselineRenderPath ,saveBaseline ,currentHashMap ,refFile ,"\u002em\u0073\u0077\u006f\u0072\u0064");
};func _cadb (_gfc ,_ddd _dg .Color )bool {_cc ,_fgg ,_dge ,_ffc :=_gfc .RGBA ();_bff ,_ebd ,_abc ,_bfa :=_ddd .RGBA ();return _cc ==_bff &&_fgg ==_ebd &&_dge ==_abc &&_ffc ==_bfa ;};func (_eg *ReferenceMap )Keys ()(_ae []string ){_ae =make ([]string ,len (_eg ._fc ));
var _bdff int ;for _bbc :=range _eg ._fc {_ae [_bdff ]=_bbc ;_bdff ++;};return _ae ;};type ReferenceMap struct{_ab .Mutex ;_fc map[string ]ReferenceEntry ;};func _gab (_dcg ,_egc string )error {_agab ,_deab :=_bdf .Open (_dcg );if _deab !=nil {return _deab ;
};defer _agab .Close ();_bge ,_ ,_deab :=_ge .DecodeConfig (_agab );if _deab !=nil {panic (_deab );};_eggb :=_ge .NewRGBA (_ge .Rect (0,0,_bge .Width ,_bge .Height ));_gge ,_deab :=_bdf .Create (_egc );if _deab !=nil {return _deab ;};defer _gge .Close ();
_deab =_f .Encode (_gge ,_eggb );if _deab !=nil {return _deab ;};return nil ;};type ReferenceEntry struct{Timestamp int64 `json:"timestamp"`;Value string `json:"value"`;ResultSize int64 `json:"resultSize,omitempty"`;DiffPercent float64 `json:"diffPercent,omitempty"`;
DiffTotal float64 `json:"diffValue,omitempty"`;Invalid bool `json:"markedInvalid,omitempty"`;};func _bece (_dff *_a .T ,_abf string )int64 {_fae ,_ccec :=_bdf .Stat (_abf );_eb .NoError (_dff ,_ccec );return _fae .Size ();};func (_cg *ReferenceFile )updateMap (_bg *ReferenceMap )int {var _geg int ;
if _cg .Map ._fc ==nil {_cg .Map ._fc =map[string ]ReferenceEntry {};};for _ec ,_be :=range _bg ._fc {_da ,_bcg :=_cg .Map ._fc [_ec ];if !_bcg {_cg .Map ._fc [_ec ]=_be ;_geg ++;continue ;};if string (_da .Value )!=string (_be .Value ){_cg .Map ._fc [_ec ]=_be ;
_geg ++;};};for _dfg :=range _cg .Map ._fc {if _ ,_dc :=_bg ._fc [_dfg ];!_dc {delete (_cg .Map ._fc ,_dfg );_geg ++;};};return _geg ;};type ReferenceFile struct{Timestamp _db .Time `json:"timestamp"`;Map *ReferenceMap `json:"map,omitempty"`;_c string ;
};func RunRenderOfficeTest (t *_a .T ,pdfPath ,outputDir ,baselineRenderPath string ,saveBaseline bool ,currentHashMap *ReferenceMap ,refFile *ReferenceFile ,postfix string ){_agf :=_b .TrimSuffix (_dfe .Base (pdfPath ),_dfe .Ext (pdfPath ));t .Run ("\u0072\u0065\u006e\u0064\u0065\u0072",func (_eee *_a .T ){_gdb :=_dfe .Join (outputDir ,_agf );
_ecgg :=_gdb +"\u002d%\u0064\u002e\u0070\u006e\u0067";if _bcgb :=RenderPDFToPNGs (pdfPath ,0,_ecgg );_bcgb !=nil {_eee .Skip (_bcgb );};_ega :=_agf +postfix ;_bfdc :=_dfe .Join (outputDir ,_ega );_eddd :=_bfdc +"\u002d%\u0064\u002e\u0070\u006e\u0067";_eag :=false ;
if saveBaseline {_fcd :=_dfe .Dir (pdfPath );_ebed :=_dfe .Join (_fcd ,_ega +"\u002e\u0070\u0064\u0066");if _ ,_ddda :=_bdf .Stat (_ebed );_ddda ==nil {_eee .Logf ("\u0052e\u006e\u0064\u0065\u0072\u0020\u004d\u0053\u0020\u004f\u0066\u0066i\u0063\u0065\u0020\u0050\u0044\u0046\u003a\u0020\u0025\u0076",_ebed );
if _caa :=RenderPDFToPNGs (_ebed ,0,_eddd );_caa !=nil {_eee .Skip (_caa );};_eag =true ;};};for _eef :=1;true ;_eef ++{_gdf :=_acf .Sprintf ("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg",_gdb ,_eef );_dgba :=_dfe .Join (baselineRenderPath ,_acf .Sprintf ("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067",_agf ,_eef ));
if _ ,_cac :=_bdf .Stat (_gdf );_cac !=nil {break ;};_eee .Logf ("\u0025\u0073",_dgba );if saveBaseline {_eee .Logf ("\u0043\u006fp\u0079\u0069\u006eg\u0020\u0025\u0073\u0020\u002d\u003e\u0020\u0025\u0073",_gdf ,_dgba );_fdf :=CopyFile (_gdf ,_dgba );if _fdf !=nil {_eee .Fatalf ("\u0045\u0052\u0052OR\u0020\u0063\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0025\u0073\u003a\u0020\u0025\u0076",_dgba ,_fdf );
};if _eag {_aff :=_acf .Sprintf ("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg",_bfdc ,_eef );_cge :=_dfe .Join (baselineRenderPath ,_acf .Sprintf ("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067",_ega ,_eef ));_eee .Logf ("\u0043\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u004d\u0053\u0020\u0057\u006f\u0072\u0064 \u0072e\u0073\u0075\u006c\u0074\u0073\u0020\u0025\u0073\u0020\u002d\u003e\u0020\u0025\u0073",_aff ,_cge );
_fa :=CopyFile (_aff ,_cge );if _fa !=nil {_eee .Logf ("\u0045\u0052\u0052\u004f\u0052\u0020\u0063\u006f\u0070\u0079\u0069\u006eg\u0020\u0074\u006f \u0025\u0073\u003a\u0020\u0025\u0076\u002c\u0020\u0068\u0061\u0076i\u006e\u0067\u0020d\u0069\u0066\u0066\u0065r\u0065\u006e\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0073\u0069\u007a\u0065\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0073\u002c\u0020\u0072\u0065\u0070\u006c\u0061\u0063\u0069\u006eg\u0020\u0069\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0062\u006ca\u006e\u006b\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0069\u006e\u0073\u0074\u0065\u0061\u0064",_cge ,_fa );
if _eca :=_gab (_dgba ,_cge );_eca !=nil {_eee .Fatalf ("\u0045\u0052\u0052\u004f\u0052\u0020c\u0072\u0065\u0061\u0074\u0069\u006e\u0067\u0020\u0065\u006d\u0070\u0074\u0079 \u0069\u006d\u0061\u0067\u0065\u0020\u0025s\u003a\u0020\u0025\u0076",_cge ,_eca );
};};_eee .Logf ("\u0043\u006fm\u0062\u0069\u006e\u0069\u006eg\u0020\u0055\u006e\u0069\u004ff\u0066\u0069\u0063\u0065\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0073\u0020\u0077\u0069\u0074\u0068\u0020\u004d\u0053\u0020\u004f\u0066\u0066\u0069\u0063\u0065\u0020\u0025\u0073\u0020\u0061\u006e\u0064\u0020\u0025\u0073",_dgba ,_cge );
_adef ,_fa :=CombinePNGFiles (_dgba ,_cge );if _bdf .IsNotExist (_fa ){_eee .Fatal ("\u0069m\u0061g\u0065\u0020\u0066\u0069\u006ce\u0020\u006di\u0073\u0073\u0069\u006e\u0067");}else if !_adef {_eee .Fatal ("\u0055n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0063\u006f\u006db\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0073");
};_eee .Logf ("\u0043\u0072\u0065\u0061\u0074\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0064\u0069\u0066f \u0055n\u0069\u004f\u0066\u0066\u0069\u0063\u0065\u0020\u0072\u0065\u0073\u0075l\u0074\u0073\u0020\u0077\u0069\u0074\u0068\u0020\u004d\u0053\u0020\u004f\u0066\u0066\u0069\u0063\u0065 \u0025\u0073\u0020\u0061\u006e\u0064\u0020\u0025\u0073",_dgba ,_cge );
_adef ,_cce ,_dafg ,_fgf ,_fa :=CreatePNGDiff (_dgba ,_cge );if _fa !=nil &&_aa .Is (_fa ,ErrImageSizeNotMatch ){_eee .Fatalf ("\u0045\u0072\u0072\u006fr\u0020\u006f\u006e\u0020\u0063\u0072\u0065\u0061\u0074\u0065 \u0050N\u0047\u0020\u0044\u0069\u0066\u0066\u003a \u0025\u0076",_fa );
};if _adef {_eee .Logf ("\u0049\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073\u000a",_cce );_eee .Logf ("D\u0069\u0066\u0066\u0020Pe\u0072c\u0065\u006e\u0074\u003a\u0020%\u0032\u002e\u0066\u0025\u0025\u000a",_dafg );_eee .Logf ("\u0044i\u0066f\u0020\u0054\u006f\u0074\u0061\u006c\u003a\u0020\u0025\u0066\u000a",_fgf );
_aed :=_dfe .Base (_cce );_abce ,_efe :=currentHashMap .Read (_aed );if _efe &&(_abce .DiffPercent < _dafg ||_abce .DiffTotal < _fgf ){_eee .Logf ("\u004e\u0065\u0077\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0073\u0020\u0068\u0061\u0076\u0069\u006e\u0067 h\u0069g\u0068\u0065\u0072\u0020\u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0070\u0065\u0072\u0063\u0065\u006e\u0074\u003a\u0020\u0025\u0066\u0020\u0061\u006e\u0064 \u0074\u006f\u0074\u0061\u006c\u0020\u0025\u0066\u000a",_dafg ,_fgf );
};_cgec ,_ffcb :=HashFile (_cce );_eb .NoError (_eee ,_ffcb );_gdce :=_bece (_eee ,_cce );if !_efe ||_abce .Value !=_cgec ||_abce .ResultSize !=_gdce ||_abce .DiffPercent !=_dafg ||_abce .DiffTotal !=_fgf {_fgcg :=ReferenceEntry {Timestamp :_db .Now ().UTC ().Unix (),Value :_cgec ,ResultSize :_gdce ,DiffPercent :_dafg ,DiffTotal :_fgf };
currentHashMap .Write (_aed ,_fgcg );if _ffcb =refFile .Update (currentHashMap );_ffcb !=nil {_eee .Logf ("\u0055\u0070\u0064\u0061\u0074\u0065\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063e\u0020f\u0069\u006c\u0065\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_ffcb );
};};};};continue ;};_eee .Run (_acf .Sprintf ("\u0070\u0061\u0067\u0065\u0025\u0064",_eef ),func (_gag *_a .T ){_gag .Logf ("\u0043o\u006dp\u0061\u0072\u0069\u006e\u0067 \u0025\u0073 \u0076\u0073\u0020\u0025\u0073",_gdf ,_dgba );_egb ,_bdg :=ComparePNGFiles (_gdf ,_dgba );
if _bdf .IsNotExist (_bdg ){_gag .Fatal ("\u0069m\u0061g\u0065\u0020\u0066\u0069\u006ce\u0020\u006di\u0073\u0073\u0069\u006e\u0067");}else if !_egb {_gag .Fatal ("\u0077\u0072\u006f\u006eg \u0070\u0061\u0067\u0065\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0065\u0064");
};});};});};func CombinePNGFiles (file1 ,file2 string )(bool ,error ){_gdc ,_cdb :=ReadPNG (file1 );if _cdb !=nil {return false ,_cdb ;};_fef ,_cdb :=ReadPNG (file2 );if _cdb !=nil {return false ,_cdb ;};_gcf :=_ge .Point {_gdc .Bounds ().Dx (),0};_dea :=_ge .Rectangle {_gcf ,_gcf .Add (_fef .Bounds ().Size ())};
_geed :=_ge .Rectangle {_ge .Point {0,0},_dea .Max };_gg :=_ge .NewRGBA (_geed );_bd .Draw (_gg ,_gdc .Bounds (),_gdc ,_ge .Point {0,0},_bd .Src );_bd .Draw (_gg ,_dea ,_fef ,_ge .Point {0,0},_bd .Src );_fed :=_dfe .Dir (file1 );_bec :=_b .TrimSuffix (_dfe .Base (file1 ),_dfe .Ext (file1 ));
_ead ,_cdb :=_bdf .Create (_fed +"\u002f"+_bec +"\u002d\u0063\u006f\u006d\u0062\u0069\u006e\u0065\u0064\u002e\u0070\u006e\u0067");if _cdb !=nil {return false ,_cdb ;};defer _ead .Close ();if _gfb :=_f .Encode (_ead ,_gg );_gfb !=nil {return false ,_gfb ;
};return true ,nil ;};func ReadPNG (file string )(_ge .Image ,error ){_ad ,_gd :=_bdf .Open (file );if _gd !=nil {return nil ,_gd ;};defer _ad .Close ();return _f .Decode (_ad );};func _dde (_ecc *_ge .RGBA ,_acd string ,_ece string ,_aeg ,_afc int )error {_gbe ,_edd :=_bc .ReadFile (_acd );
if _edd !=nil {return _edd ;};_bafa ,_edd :=_fg .Parse (_gbe );if _edd !=nil {return _edd ;};_afe ,_edd :=_fg .NewFace (_bafa ,&_fg .FaceOptions {Size :18,DPI :72,Hinting :_bb .HintingNone });if _edd !=nil {return _edd ;};_cgc :=&_bb .Drawer {Dst :_ecc ,Src :_ge .NewUniform (_dg .RGBA {200,100,0,255}),Face :_afe ,Dot :_gc .P (_aeg ,_afc )};
_cgc .DrawString (_ece );return nil ;};func (_dd *ReferenceMap )Len ()int {return len (_dd ._fc )};func RenderPDFToPNGs (pdfPath string ,dpi int ,outpathTpl string )error {if dpi <=0{dpi =100;};if _ ,_ced :=_g .LookPath ("\u0067\u0073");_ced !=nil {return ErrRenderNotSupported ;
};return _g .Command ("\u0067\u0073","\u002d\u0073\u0044\u0045\u0056\u0049\u0043\u0045\u003d\u0070\u006e\u0067a\u006c\u0070\u0068\u0061","\u002d\u006f",outpathTpl ,_acf .Sprintf ("\u002d\u0072\u0025\u0064",dpi ),pdfPath ).Run ();};func ComparePNGFiles (file1 ,file2 string )(bool ,error ){_ecg ,_ffb :=HashFile (file1 );
if _ffb !=nil {return false ,_ffb ;};_bebe ,_ffb :=HashFile (file2 );if _ffb !=nil {return false ,_ffb ;};if _ecg ==_bebe {return true ,nil ;};_ee ,_ffb :=ReadPNG (file1 );if _ffb !=nil {return false ,_ffb ;};_cd ,_ffb :=ReadPNG (file2 );if _ffb !=nil {return false ,_ffb ;
};if _ee .Bounds ()!=_cd .Bounds (){return false ,nil ;};return CompareImages (_ee ,_cd );};func (_ebe *ReferenceMap )MarshalJSON ()([]byte ,error ){return _e .Marshal (_ebe ._fc )};func (_bafe *ReferenceMap )UnmarshalJSON (data []byte )error {return _e .Unmarshal (data ,&_bafe ._fc )};
func (_fga *ReferenceMap )Copy ()*ReferenceMap {_de :=ReferenceMap {_fc :make (map[string ]ReferenceEntry ,len (_fga ._fc ))};for _daf ,_ff :=range _fga ._fc {_de ._fc [_daf ]=_ff ;};return &_de ;};func ReadFile (dirPath ,testName string ,createEmpty bool )(*ReferenceFile ,error ){if dirPath ==""&&createEmpty {return &ReferenceFile {Map :&ReferenceMap {}},nil ;
};if dirPath ==""{return nil ,_bdf .ErrNotExist ;};_fca :=_dfe .Join (dirPath ,testName +"\u005fr\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u002e\u006a\u0073\u006f\u006e");_aae :=&ReferenceFile {_c :_fca };_fgc ,_ca :=_bdf .Open (_fca );if _aa .Is (_ca ,_bdf .ErrNotExist )&&createEmpty {_aae .Timestamp =_db .Now ().UTC ();
_aae .Map =&ReferenceMap {};return _aae ,nil ;};if _ca !=nil {return nil ,_ca ;};defer _fgc .Close ();if _fb :=_e .NewDecoder (_fgc ).Decode (_aae );_fb !=nil {if _fb .Error ()=="i\u006c\u006c\u0065\u0067\u0061\u006c \u0062\u0061\u0073\u0065\u0036\u0034 \u0064\u0061\u0074\u0061\u0020\u0061\u0074 \u0069\u006e\u0070\u0075\u0074\u0020\u0062\u0079\u0074\u0065 \u0030"&&createEmpty {return _aae ,nil ;
};return nil ,_fb ;};return _aae ,nil ;};func CreatePNGDiff (src ,dst string )(bool ,string ,float64 ,float64 ,error ){_egf ,_bbb :=ReadPNG (src );if _bbb !=nil {return false ,"",0,0,_bbb ;};_egg ,_bbb :=ReadPNG (dst );if _bbb !=nil {return false ,"",0,0,_bbb ;
};_cef :=_egf .Bounds ();_aac :=_egg .Bounds ();if !_dgb (_cef ,_aac ){if _ac .Abs (float64 (_cef .Max .X )-float64 (_aac .Max .X ))> 1{_ba .Printf ("S\u006f\u0075\u0072\u0063\u0065\u0020b\u006f\u0075\u006e\u0064\u0073\u003a \u0025\u0076\u003b\u0020\u0044\u0065\u0073t\u0020\u0062\u006f\u0075\u006e\u0064\u0073\u003a\u0020\u0025v\u000a",_cef ,_aac );
return false ,"",0,0,ErrImageSizeNotMatch ;};};_dfc :=_ge .NewRGBA (_ge .Rect (0,0,_cef .Max .X ,_cef .Max .Y ));var (_eec float64 ;_aea float64 ;);for _aag :=_cef .Min .Y ;_aag < _cef .Max .Y ;_aag ++{for _abd :=_cef .Min .X ;_abd < _cef .Max .X ;_abd ++{_efg ,_fcaa ,_ggc ,_gb :=_egf .At (_abd ,_aag ).RGBA ();
_ddc ,_becc ,_aga ,_bfd :=_egg .At (_abd ,_aag ).RGBA ();_dfc .Set (_abd ,_aag ,_dg .RGBA {uint8 (_ddc ),uint8 (_becc ),uint8 (_aga ),64});_dac :=_gb ==0x00&&_efg ==0x00&&_fcaa ==0x00&&_ggc ==0x00&&_ddc ==0xFFFF&&_becc ==0xFFFF&&_aga ==0xFFFF;if !_dac &&!_cadb (_egf .At (_abd ,_aag ),_egg .At (_abd ,_aag )){_dfc .Set (_abd ,_aag ,_dg .RGBA {uint8 (_efg ),uint8 (_fcaa ),uint8 (_ggc ),uint8 (_gb )});
_efc :=float64 (_efg )+float64 (_fcaa )+float64 (_ggc )+float64 (_gb )-float64 (_ddc )+float64 (_becc )+float64 (_aga )+float64 (_bfd );_aeab :=_ac .Sqrt (_ac .Pow (_efc /float64 (_cef .Dx ()*_cef .Dy ()),2));_aea +=_aeab ;_eec ++;};};};_afb :=_eec /float64 (_cef .Dx ()*_cef .Dy ())*100;
_ffg :=_dfe .Dir (src );_cad :=_b .TrimSuffix (_dfe .Base (src ),_dfe .Ext (src ));_cdg ,_bbb :=_bdf .Create (_ffg +"\u002f"+_cad +"\u002dd\u0069\u0066\u0066\u002e\u0070\u006eg");if _bbb !=nil {return false ,"",0,0,_bbb ;};defer _cdg .Close ();_aded :=_b .Replace (_ffg ,"\u0072\u0065\u006e\u0064\u0065\u0072","\u0066\u006f\u006et\u0073",1)+"\u002f\u0043\u0061l\u0069\u0062\u0072\u0069\u002e\u0074\u0074\u0066";
_egd :=_acf .Sprintf ("\u0044\u0069f\u0066\u0065\u0072e\u006e\u0063\u0065\u003a\u0020\u0025\u0066\u0025\u0025",_afb );_bbb =_dde (_dfc ,_aded ,_egd ,15,22);if _bbb !=nil {return false ,"",0,0,_bbb ;};_egd =_acf .Sprintf ("T\u006ft\u0061\u006c\u0020\u0044\u0069\u0066\u0066\u0065r\u0065\u006e\u0063\u0065: \u0025\u0066",_aea );
_bbb =_dde (_dfc ,_aded ,_egd ,15,44);if _bbb !=nil {return false ,"",0,0,_bbb ;};if _fdb :=_f .Encode (_cdg ,_dfc );_fdb !=nil {return false ,"",0,0,_fdb ;};return true ,_cdg .Name (),_afb ,_aea ,nil ;};var (ErrRenderNotSupported =_aa .New ("\u0072\u0065\u006e\u0064\u0065r\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0073 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u006e\u0020\u0074\u0068\u0069\u0073\u0020\u0073\u0079\u0073\u0074\u0065m");
ErrImageSizeNotMatch =_aa .New ("\u0069\u006d\u0061ge\u0020\u0073\u0069\u007a\u0065\u0073\u0020\u0064\u006f\u006e\u0027\u0074\u0020\u006d\u0061\u0074\u0063\u0068"););func (_fe *ReferenceFile )Update (currentMap *ReferenceMap )error {if _fe ._c ==""{return nil ;
};_ce :=_fe .updateMap (currentMap );if _ce ==0{return nil ;};_bde ,_ag :=_bdf .OpenFile (_fe ._c ,_bdf .O_CREATE |_bdf .O_TRUNC |_bdf .O_WRONLY ,0664);if _ag !=nil {return _ag ;};defer _bde .Close ();_fe .Timestamp =_db .Now ().UTC ();_bce :=_e .NewEncoder (_bde );
_bce .SetIndent ("","\u0009");return _bce .Encode (_fe );};func HashFile (file string )(string ,error ){_beb ,_ed :=_bdf .Open (file );if _ed !=nil {return "",_ed ;};defer _beb .Close ();_gea :=_df .New ();if _ ,_ed =_baf .Copy (_gea ,_beb );_ed !=nil {return "",_ed ;
};return _af .EncodeToString (_gea .Sum (nil )),nil ;};