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

package mscfb ;import (_e "bytes";_ac "encoding/binary";_d "fmt";_ce "github.com/richardlehane/msoleps/types";_dd "github.com/unidoc/unioffice/internal/mscfb/rw";_gf "io";_g "os";_c "strconv";_cb "time";_cd "unicode";_ee "unicode/utf16";);const _ef int =64+4*4+16+4+8*2+4+8;
const (_b uint8 =0x0;_dc uint8 =0x1;_f uint8 =0x2;_gc uint8 =0x5;);func (_gbe *File )Modified ()_cb .Time {return _gbe ._af .Time ()};func (_fgg *File )stream (_feaf int )([][2]int64 ,error ){var _gcee bool ;var _bbg int ;var _fccb int64 ;if _fgg .Size < _cgdc {_gcee =true ;
_bbg =_feaf /int (_gcbd )+2;_fccb =int64 (_gcbd );}else {_bbg =_feaf /int (_fgg ._adg ._cad )+2;_fccb =int64 (_fgg ._adg ._cad );};_ffg :=make ([][2]int64 ,0,_bbg );var _fbd ,_bdba int ;if _fgg ._aba > 0{_bfa ,_edd :=_fgg ._adg .getOffset (_fgg ._cc ,_gcee );
if _edd !=nil {return nil ,_edd ;};if _fccb -_fgg ._aba >=int64 (_feaf ){_ffg =append (_ffg ,[2]int64 {_bfa +_fgg ._aba ,int64 (_feaf )});}else {_ffg =append (_ffg ,[2]int64 {_bfa +_fgg ._aba ,_fccb -_fgg ._aba });};if _fccb -_fgg ._aba <=int64 (_feaf ){_fgg ._cc ,_edd =_fgg ._adg .findNext (_fgg ._cc ,_gcee );
if _edd !=nil {return nil ,_edd ;};_bdba +=int (_fccb -_fgg ._aba );_fgg ._aba =0;}else {_fgg ._aba +=int64 (_feaf );};if _ffg [0][1]==int64 (_feaf ){return _ffg ,nil ;};if _fgg ._cc ==_ffc {return nil ,Error {ErrRead ,"\u0075\u006ee\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0065\u0061\u0072\u006c\u0079\u0020\u0065\u006e\u0064\u0020\u006f\u0066\u0020\u0063ha\u0069\u006e",int64 (_fgg ._cc )};
};_fbd ++;};for {if _fbd >=cap (_ffg ){return nil ,Error {ErrRead ,"\u0069\u006e\u0064\u0065x\u0020\u006f\u0076\u0065\u0072\u0072\u0075\u006e\u0073\u0020s\u0065c\u0074\u006f\u0072\u0020\u006c\u0065\u006eg\u0074\u0068",int64 (_fbd )};};_bgb ,_ccb :=_fgg ._adg .getOffset (_fgg ._cc ,_gcee );
if _ccb !=nil {return nil ,_ccb ;};if _feaf -_bdba < int (_fccb ){_ffg =append (_ffg ,[2]int64 {_bgb ,int64 (_feaf -_bdba )});_fgg ._aba =int64 (_feaf -_bdba );return _dgec (_ffg ),nil ;}else {_ffg =append (_ffg ,[2]int64 {_bgb ,_fccb });_bdba +=int (_fccb );
_fgg ._cc ,_ccb =_fgg ._adg .findNext (_fgg ._cc ,_gcee );if _ccb !=nil {return nil ,_ccb ;};if _bdba ==_feaf {return _dgec (_ffg ),nil ;};};_fbd ++;};};func (_bgfb *File )reset (){_bgfb ._acg =0;_bgfb ._aba =0;_bgfb ._cc =_bgfb ._ec };func (_bbd fileInfo )ModTime ()_cb .Time {return _bbd .Modified ()};
func (_bee *File )Created ()_cb .Time {return _bee ._gg .Time ()};func (_gge *File )SetEntryContent (b []byte )error {if _cfa :=_gge .ensureWriterAt ();_cfa !=nil {return _cfa ;};_gge .reset ();if _bbb :=_gge .changeSize (int64 (len (b )));_bbb !=nil {return _bbb ;
};_ ,_ecc :=_gge .Write (b );return _ecc ;};func (_cagg *File )WriteAt (p []byte ,off int64 )(_gfcd int ,_ggce error ){_fgb ,_abc ,_ged :=_cagg ._acg ,_cagg ._aba ,_cagg ._cc ;_ ,_ggce =_cagg .Seek (off ,0);if _ggce ==nil {_gfcd ,_ggce =_cagg .Write (p );
};_cagg ._acg ,_cagg ._aba ,_cagg ._cc =_fgb ,_abc ,_ged ;return _gfcd ,_ggce ;};const (ErrFormat =iota ;ErrRead ;ErrSeek ;ErrWrite ;ErrTraverse ;);func (_ffb *Reader )getOffset (_faba uint32 ,_gcbe bool )(int64 ,error ){if _gcbe {_fcac :=_ffb ._cad /64;
_ddc :=int (_faba /_fcac );if _ddc >=len (_ffb ._adb ._bbdf ){return 0,Error {ErrRead ,"\u006di\u006e\u0069s\u0065\u0063\u0074o\u0072\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0069\u0073\u0020\u006f\u0075t\u0073\u0069\u0064\u0065\u0020\u006d\u0069\u006e\u0069\u0073\u0065c\u0074\u006f\u0072\u0020\u0072\u0061\u006e\u0067\u0065",int64 (_ddc )};
};_becc :=_faba %_fcac ;return int64 ((_ffb ._adb ._bbdf [_ddc ]+1)*_ffb ._cad +_becc *64),nil ;};return _eae (_ffb ._cad ,_faba ),nil ;};func (_ecg *Reader )Next ()(*File ,error ){_ecg ._dbd ++;if _ecg ._dbd >=len (_ecg .File ){return nil ,_gf .EOF ;};
return _ecg .File [_ecg ._dbd ],nil ;};func (_ffeb *Reader )Debug ()map[string ][]uint32 {_aagb :=map[string ][]uint32 {"s\u0065\u0063\u0074\u006f\u0072\u0020\u0073\u0069\u007a\u0065":[]uint32 {_ffeb ._cad },"\u006d\u0069\u006e\u0069\u0020\u0066\u0061\u0074\u0020\u006c\u006f\u0063\u0073":_ffeb ._adb ._fef ,"\u006d\u0069n\u0069\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u006c\u006f\u0063\u0073":_ffeb ._adb ._bbdf ,"\u0064\u0069r\u0065\u0063\u0074o\u0072\u0079\u0020\u0073\u0065\u0063\u0074\u006f\u0072":[]uint32 {_ffeb ._adb ._gbb },"\u006d\u0069\u006e\u0069 s\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0074\u0061\u0072\u0074\u002f\u0073\u0069z\u0065":[]uint32 {_ffeb .File [0]._ec ,_ac .LittleEndian .Uint32 (_ffeb .File [0]._ca [:])}};
for _dcba ,_gbg :=_ffeb .Next ();_gbg ==nil ;_dcba ,_gbg =_ffeb .Next (){_aagb [_dcba .Name +" \u0073\u0074\u0061\u0072\u0074\u002f\u0073\u0069\u007a\u0065"]=[]uint32 {_dcba ._ec ,_ac .LittleEndian .Uint32 (_dcba ._ca [:])};};return _aagb ;};func (_ffcb *Reader )findFatLocsOffset (_ggf bool )int64 {var _dddf uint32 ;
if _ggf {_dddf =_ffcb ._adb ._fef [0];}else {_dddf =_ffcb ._adb ._fae [0];};return _eae (_ffcb ._cad ,_dddf );};func (_gfd *Reader )readAt (_ced int64 ,_abe int )([]byte ,error ){if _gfd ._abb {_ddec ,_bbde :=_gfd ._dcac .(slicer ).Slice (_ced ,_abe );
if _bbde !=nil {return nil ,Error {ErrRead ,"\u0073\u006c\u0069\u0063er\u0020\u0072\u0065\u0061\u0064\u0020\u0065\u0072\u0072\u006f\u0072\u0020\u0028"+_bbde .Error ()+"\u0029",_ced };};return _ddec ,nil ;};if _abe > len (_gfd ._beb ){return nil ,Error {ErrRead ,"\u0072\u0065ad\u0020\u006c\u0065n\u0067\u0074\u0068\u0020gre\u0061te\u0072\u0020\u0074\u0068\u0061\u006e\u0020re\u0061\u0064\u0020\u0062\u0075\u0066\u0066e\u0072",int64 (_abe )};
};if _ ,_edf :=_gfd ._dcac .ReadAt (_gfd ._beb [:_abe ],_ced );_edf !=nil {return nil ,Error {ErrRead ,_edf .Error (),_ced };};return _gfd ._beb [:_abe ],nil ;};func (_fcb *File )seek (_bca int64 )error {var _gdd bool ;var _deg int64 ;if _fcb .Size < _cgdc {_gdd =true ;
_deg =64;}else {_deg =int64 (_fcb ._adg ._cad );};var _db int64 ;var _edc error ;if _fcb ._aba > 0{if _deg -_fcb ._aba <=_bca {_fcb ._cc ,_edc =_fcb ._adg .findNext (_fcb ._cc ,_gdd );if _edc !=nil {return _edc ;};_db +=_deg -_fcb ._aba ;_fcb ._aba =0;
if _db ==_bca {return nil ;};}else {_fcb ._aba +=_bca ;return nil ;};if _fcb ._cc ==_ffc {return Error {ErrRead ,"\u0075\u006ee\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0065\u0061\u0072\u006c\u0079\u0020\u0065\u006e\u0064\u0020\u006f\u0066\u0020\u0063ha\u0069\u006e",int64 (_fcb ._cc )};
};};for {if _bca -_db < _deg {_fcb ._aba =_bca -_db ;return nil ;}else {_db +=_deg ;_fcb ._cc ,_edc =_fcb ._adg .findNext (_fcb ._cc ,_gdd );if _edc !=nil {return _edc ;};if _db ==_bca {return nil ;};};};};type header struct{*headerFields ;_fae []uint32 ;
_fef []uint32 ;_bbdf []uint32 ;};func (_cdbc *File )FileInfo ()_g .FileInfo {return fileInfo {_cdbc }};func (_dga *Reader )saveToFatLocs (_efc uint32 ,_bdec interface{},_fbfa bool )error {_edea :=_e .NewBuffer ([]byte {});if _fece :=_ac .Write (_edea ,_ac .LittleEndian ,_bdec );
_fece !=nil {return _fece ;};_efbf :=_dga .findFatLocsOffset (_fbfa )+int64 (_efc *4);_ ,_bbc :=_dga ._cddg .WriteAt (_edea .Bytes (),_efbf );return _bbc ;};func (_gfc fileInfo )Size ()int64 {if _gfc ._ed !=_f {return 0;};return _gfc .File .Size ;};func (_fca *File )ensureWriterAt ()error {if _fca ._adg ._cddg ==nil {_ece ,_cg :=_fca ._adg ._dcac .(_gf .WriterAt );
if !_cg {return Error {ErrWrite ,"\u006d\u0073\u0063\u0066\u0062\u002e\u004ee\u0077\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0067\u0069\u0076\u0065n\u0020R\u0065\u0061\u0064\u0065\u0072\u0041t\u0020\u0063\u006f\u006e\u0076\u0065\u0072t\u0069\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0020\u0069\u006f\u002e\u0057\u0072\u0069\u0074\u0065\u0072\u0041\u0074\u0020\u0069n\u0020\u006f\u0072\u0064\u0065\u0072\u0020\u0074\u006f\u0020\u0077\u0072\u0069t\u0065",0};
};_fca ._adg ._cddg =_ece ;};return nil ;};func (_dg *File )Read (b []byte )(int ,error ){if _dg .Size < 1||_dg ._acg >=_dg .Size {return 0,_gf .EOF ;};_eab :=len (b );if int64 (_eab )> _dg .Size -_dg ._acg {_eab =int (_dg .Size -_dg ._acg );};_abad ,_dge :=_dg .stream (_eab );
if _dge !=nil {return 0,_dge ;};var _fga ,_gce int ;for _ ,_fa :=range _abad {_fec :=_fga +int (_fa [1]);if _fec < _fga ||_fec > _eab {return 0,Error {ErrRead ,"\u0062a\u0064 \u0072\u0065\u0061\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068",int64 (_fec )};
};_ccf ,_aag :=_dg ._adg ._dcac .ReadAt (b [_fga :_fec ],_fa [0]);_gce =_gce +_ccf ;if _aag !=nil {_dg ._acg +=int64 (_gce );return _gce ,Error {ErrRead ,"\u0075n\u0064\u0065\u0072\u006c\u0079\u0069\u006e\u0067\u0020\u0072\u0065a\u0064\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0020\u0028"+_aag .Error ()+"\u0029",int64 (_fga )};
};_fga =_fec ;};_dg ._acg +=int64 (_gce );if _gce !=_eab {_dge =Error {ErrRead ,"\u0062\u0079\u0074e\u0073\u0020\u0072\u0065\u0061\u0064\u0020\u0064\u006f\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020r\u0065\u0061\u0064\u0020\u0073\u0069\u007a\u0065",int64 (_gce )};
}else if _gce < len (b ){_dge =_gf .EOF ;};return _gce ,_dge ;};func (_gedd *Reader )GetHeader ()*header {return _gedd ._adb };func (_bbf *Reader )Modified ()_cb .Time {return _bbf .File [0].Modified ()};func (_ecef *Reader )Read (b []byte )(_dbb int ,_cbee error ){if _ecef ._dbd >=len (_ecef .File ){return 0,_gf .EOF ;
};return _ecef .File [_ecef ._dbd ].Read (b );};func _eaa (_ead []byte )*headerFields {_efa :=&headerFields {};_efa ._ba =_ac .LittleEndian .Uint64 (_ead [:8]);_efa ._aefc =_ac .LittleEndian .Uint16 (_ead [24:26]);_efa ._efg =_ac .LittleEndian .Uint16 (_ead [26:28]);
_efa ._dff =_ac .LittleEndian .Uint16 (_ead [30:32]);_efa ._cced =_ac .LittleEndian .Uint32 (_ead [40:44]);_efa ._afa =_ac .LittleEndian .Uint32 (_ead [44:48]);_efa ._gbb =_ac .LittleEndian .Uint32 (_ead [48:52]);_efa ._cdd =_ac .LittleEndian .Uint32 (_ead [60:64]);
_efa ._ccef =_ac .LittleEndian .Uint32 (_ead [64:68]);_efa ._cfee =_ac .LittleEndian .Uint32 (_ead [68:72]);_efa ._feca =_ac .LittleEndian .Uint32 (_ead [72:76]);var _geff int ;for _ccg :=76;_ccg < 512;_ccg =_ccg +4{_efa ._gegc [_geff ]=_ac .LittleEndian .Uint32 (_ead [_ccg :_ccg +4]);
_geff ++;};return _efa ;};func _gdc (_dag uint16 ,_eg *File ){_egb (_eg );if _eg ._ed !=_f {return ;};if _dag > 3{_eg .Size =int64 (_ac .LittleEndian .Uint64 (_eg ._ca [:]));}else {_eg .Size =int64 (_ac .LittleEndian .Uint32 (_eg ._ca [:4]));};};func (_aae Error )Typ ()int {return _aae ._ecbb };
func (_dde fileInfo )Sys ()interface{}{return nil };func (_dgg *Reader )exportFAT (_cfc *_dd .Writer ,_ccge []uint32 )error {if _dgg ._adb ._afa ==0{return nil ;};_ggea :=_e .NewBuffer ([]byte {});if _ggd :=_ac .Write (_ggea ,_ac .LittleEndian ,_efb );
_ggd !=nil {return _ggd ;};for _debd :=0;_debd < len (_ccge )-1;_debd ++{for _fgd :=_ccge [_debd ];_fgd < _ccge [_debd +1]-1;_fgd ++{if _abadd :=_ac .Write (_ggea ,_ac .LittleEndian ,_fgd );_abadd !=nil {return _abadd ;};};if _cee :=_ac .Write (_ggea ,_ac .LittleEndian ,_ffc );
_cee !=nil {return _cee ;};};_eed :=512;for _ ,_ega :=range _ggea .Bytes (){if _ddbc :=_cfc .WriteByteAt (_ega ,_eed );_ddbc !=nil {return _ddbc ;};_eed ++;};return nil ;};func (_cfed *Reader )setMiniStream ()error {if _cfed ._gcd [0]._ec ==_ffc ||_cfed ._adb ._cdd ==_ffc ||_cfed ._adb ._ccef ==0{return nil ;
};_bef :=int (_cfed ._adb ._ccef );_cfed ._adb ._fef =make ([]uint32 ,_bef );_cfed ._adb ._fef [0]=_cfed ._adb ._cdd ;for _adgg :=1;_adgg < _bef ;_adgg ++{_ecb ,_dfg :=_cfed .findNext (_cfed ._adb ._fef [_adgg -1],false );if _dfg !=nil {return Error {ErrFormat ,"s\u0065\u0074\u0074\u0069ng\u0020m\u0069\u006e\u0069\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u0028"+_dfg .Error ()+"\u0029",int64 (_cfed ._adb ._fef [_adgg -1])};
};_cfed ._adb ._fef [_adgg ]=_ecb ;};_bef =int (_cfed ._cad /4*_cfed ._adb ._ccef );_cfed ._adb ._bbdf =make ([]uint32 ,0,_bef );_fdc :=_cfed ._gcd [0]._ec ;var _eggd error ;for _fdc !=_ffc {_cfed ._adb ._bbdf =append (_cfed ._adb ._bbdf ,_fdc );_fdc ,_eggd =_cfed .findNext (_fdc ,false );
if _eggd !=nil {return Error {ErrFormat ,"s\u0065\u0074\u0074\u0069ng\u0020m\u0069\u006e\u0069\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u0028"+_eggd .Error ()+"\u0029",int64 (_fdc )};};};return nil ;};func (_add fileInfo )Mode ()_g .FileMode {return _add .File .mode ()};
func (_be fileInfo )IsDir ()bool {return _be .mode ().IsDir ()};const (_fab uint32 =0xFFFFFFFA;_fdd uint32 =0xFFFFFFFC;_efb uint32 =0xFFFFFFFD;_ffc uint32 =0xFFFFFFFE;_fbfd uint32 =0xFFFFFFFF;_edb byte =0xFF;_aeab uint32 =0xFFFFFFFA;_dcfeg uint32 =0xFFFFFFFF;
);func _eae (_ffe ,_bdgb uint32 )int64 {return int64 ((_bdgb +1)*_ffe )};func (_ggbf *Reader )Export ()([]byte ,error ){_cadc :=_dd .NewWriter ();if _cdg :=_ggbf .exportHeader (_cadc );_cdg !=nil {return nil ,_cdg ;};if _afdf :=_cadc .FillWithByte (512,_edb );
_afdf !=nil {return nil ,_afdf ;};_fbgf :=[]uint32 {};if _gaf :=_ggbf .exportDifats (_cadc );_gaf !=nil {return nil ,_gaf ;};_dfe ,_gggc ,_ebcb :=_ggbf .exportMiniStream ();if _ebcb !=nil {return nil ,_ebcb ;};_fbgf =append (_fbgf ,uint32 (_cadc .Len ())/_ggbf ._cad );
if _adcb :=_ggbf .exportDirEntries (_cadc );_adcb !=nil {return nil ,_adcb ;};_fbgf =append (_fbgf ,uint32 (_cadc .Len ())/_ggbf ._cad );if _ ,_dcfd :=_dfe .WriteTo (_cadc );_dcfd !=nil {return nil ,_dcfd ;};_fbgf =append (_fbgf ,uint32 (_cadc .Len ())/_ggbf ._cad );
if _ ,_gfca :=_gggc .WriteTo (_cadc );_gfca !=nil {return nil ,_gfca ;};_fbgf =append (_fbgf ,uint32 (_cadc .Len ())/_ggbf ._cad );if _bea :=_ggbf .exportFAT (_cadc ,_fbgf );_bea !=nil {return nil ,_bea ;};return _cadc .Bytes (),nil ;};func New (ra _gf .ReaderAt )(*Reader ,error ){_cef :=&Reader {_dcac :ra };
if _ ,_beca :=ra .(slicer );_beca {_cef ._abb =true ;}else {_cef ._beb =make ([]byte ,_gfb );};if _gfeb :=_cef .setHeader ();_gfeb !=nil {return nil ,_gfeb ;};if !_cef ._abb &&int (_cef ._cad )> len (_cef ._beb ){_cef ._beb =make ([]byte ,_cef ._cad );
};if _dcbb :=_cef .setDifats ();_dcbb !=nil {return nil ,_dcbb ;};if _cceg :=_cef .setDirEntries ();_cceg !=nil {return nil ,_cceg ;};if _dad :=_cef .setMiniStream ();_dad !=nil {return nil ,_dad ;};if _bde :=_cef .traverse ();_bde !=nil {return nil ,_bde ;
};return _cef ,nil ;};func (_facc *Reader )GetEntry (entryName string )(*File ,error ){for _cbe ,_fefb :=_facc .Next ();_fefb ==nil ;_cbe ,_fefb =_facc .Next (){if _cbe .Name ==entryName {return _cbe ,nil ;};};return nil ,Error {ErrTraverse ,"\u004e\u006f\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066o\u0075\u006e\u0064\u0020\u0066\u006f\u0072 \u0067\u0069\u0076\u0065\u006e\u0020\u006e\u0061\u006d\u0065\u002e",0};
};type File struct{Name string ;Initial uint16 ;Path []string ;Size int64 ;_acg int64 ;_cc uint32 ;_aba int64 ;*directoryEntryFields ;_adg *Reader ;};func (_ade *Reader )setHeader ()error {_gbd ,_ag :=_ade .readAt (0,_gfb );if _ag !=nil {return _ag ;};
_ade ._adb =&header {headerFields :_eaa (_gbd )};if _ade ._adb ._ba !=_dgf {return Error {ErrFormat ,"\u0062\u0061\u0064\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065",int64 (_ade ._adb ._ba )};};if _ade ._adb ._dff ==0x0009||_ade ._adb ._dff ==0x000c{_ade ._cad =uint32 (1<<_ade ._adb ._dff );
}else {return Error {ErrFormat ,"\u0069\u006c\u006c\u0065ga\u006c\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0073\u0069\u007a\u0065",int64 (_ade ._adb ._dff )};};if _ade ._adb ._feca > 0{_dce :=(_ade ._cad /4)-1;if int (_ade ._adb ._feca *_dce +109)< 0{return Error {ErrFormat ,"\u0044I\u0046A\u0054\u0020\u0069\u006e\u0074 \u006f\u0076e\u0072\u0066\u006c\u006f\u0077",int64 (_ade ._adb ._feca )};
};if _ade ._adb ._feca *_dce +109> _ade ._adb ._afa +_dce {return Error {ErrFormat ,"\u006e\u0075\u006d\u0020\u0044\u0049\u0046\u0041\u0054\u0073 \u0065\u0078\u0063\u0065\u0065\u0064\u0073 \u0046\u0041\u0054\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0073",int64 (_ade ._adb ._feca )};
};};if _ade ._adb ._ccef > 0{if int (_ade ._cad /4*_ade ._adb ._ccef )< 0{return Error {ErrFormat ,"m\u0069\u006e\u0069\u0020FA\u0054 \u0069\u006e\u0074\u0020\u006fv\u0065\u0072\u0066\u006c\u006f\u0077",int64 (_ade ._adb ._ccef )};};if _ade ._adb ._ccef > _ade ._adb ._afa *(_ade ._cad /_gcbd ){return Error {ErrFormat ,"\u006e\u0075\u006d\u0020\u006d\u0069n\u0069\u0020\u0046\u0041\u0054\u0073\u0020\u0065\u0078\u0063\u0065\u0065\u0064s\u0020\u0046\u0041\u0054\u0020\u0073\u0065c\u0074\u006f\u0072\u0073",int64 (_ade ._adb ._afa )};
};};return nil ;};func (_fcde *Reader )Created ()_cb .Time {return _fcde .File [0].Created ()};func _dgec (_dddd [][2]int64 )[][2]int64 {_gbf :=len (_dddd );for _bbge ,_fbg :=0,0;_bbge < _gbf &&_fbg +1< len (_dddd );_bbge ++{if _dddd [_fbg ][0]+_dddd [_fbg ][1]==_dddd [_fbg +1][0]{_dddd [_fbg ][1]=_dddd [_fbg ][1]+_dddd [_fbg +1][1];
for _edcf :=range _dddd [_fbg +1:len (_dddd )-1]{_dddd [_fbg +1+_edcf ]=_dddd [_edcf +_fbg +2];};_dddd =_dddd [:len (_dddd )-1];}else {_fbg +=1;};};return _dddd ;};func (_fdf Error )Error ()string {return "\u006ds\u0063\u0066\u0062\u003a\u0020"+_fdf ._addb +"\u003b\u0020"+_c .FormatInt (_fdf ._ggbe ,10);
};func (_fadd *Reader )findNextFreeSector (_bba bool )(uint32 ,error ){_gag :=_fadd .findFatLocsOffset (_bba );_bfb :=uint32 (0);_fbdb :=_fadd ._cad /4;for {_bbbd ,_ccgc :=_fadd .readAt (_gag ,4);if _ccgc !=nil {return 0,Error {ErrRead ,"\u0062\u0061\u0064\u0020\u0072\u0065\u0061\u0064\u0020\u0066i\u006e\u0064\u0069\u006e\u0067\u0020\u006ee\u0078\u0074\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0028"+_ccgc .Error ()+"\u0029",_gag };
};_ecbc :=_ac .LittleEndian .Uint32 (_bbbd );if _ecbc ==_fbfd {break ;};if _bfb >=_fbdb {return 0,Error {ErrRead ,"\u0065\u006e\u0064\u0020of\u0020\u006d\u0069\u006e\u0069\u0046\u0061\u0074\u0020\u0072\u0065\u0061\u0063\u0068e\u0064",_gag };};_bfb ++;_gag +=4;
};return _bfb ,nil ;};func _egb (_dcc *File ){if _dcc ._cdb < 4||_dcc ._cdb > 64{return ;};_ae :=int (_dcc ._cdb /2-1);_dcc .Initial =_dcc ._bd [0];var _gcf int ;if !_cd .IsPrint (rune (_dcc .Initial )){_gcf =1;};_dcc .Name =string (_ee .Decode (_dcc ._bd [_gcf :_ae ]));
};func _fd (_dfc []byte )*directoryEntryFields {_fc :=&directoryEntryFields {};for _gb :=range _fc ._bd {_fc ._bd [_gb ]=_ac .LittleEndian .Uint16 (_dfc [_gb *2:_gb *2+2]);};_fc ._cdb =_ac .LittleEndian .Uint16 (_dfc [64:66]);_fc ._ed =uint8 (_dfc [66]);
_fc ._dcf =uint8 (_dfc [67]);_fc ._ge =_ac .LittleEndian .Uint32 (_dfc [68:72]);_fc ._fe =_ac .LittleEndian .Uint32 (_dfc [72:76]);_fc ._df =_ac .LittleEndian .Uint32 (_dfc [76:80]);_fc ._bc =_ce .MustGuid (_dfc [80:96]);copy (_fc ._ad [:],_dfc [96:100]);
_fc ._gg =_ce .MustFileTime (_dfc [100:108]);_fc ._af =_ce .MustFileTime (_dfc [108:116]);_fc ._ec =_ac .LittleEndian .Uint32 (_dfc [116:120]);copy (_fc ._ca [:],_dfc [120:128]);return _fc ;};const (_ab uint8 =0x0;_ea uint8 =0x1;);func (_eb *File )Write (b []byte )(int ,error ){if _eb .Size < 1||_eb ._acg >=_eb .Size {return 0,_gf .EOF ;
};if _geg :=_eb .ensureWriterAt ();_geg !=nil {return 0,_geg ;};_fcd :=len (b );if int64 (_fcd )> _eb .Size -_eb ._acg {_fcd =int (_eb .Size -_eb ._acg );};_dage ,_dea :=_eb .stream (_fcd );if _dea !=nil {return 0,_dea ;};var _gef ,_cdc int ;for _ ,_cf :=range _dage {_dcfe :=_gef +int (_cf [1]);
if _dcfe < _gef ||_dcfe > _fcd {return 0,Error {ErrWrite ,"\u0062\u0061d\u0020\u0077\u0072i\u0074\u0065\u0020\u006c\u0065\u006e\u0067\u0074\u0068",int64 (_dcfe )};};_adc ,_bdb :=_eb ._adg ._cddg .WriteAt (b [_gef :_dcfe ],_cf [0]);_cdc =_cdc +_adc ;if _bdb !=nil {_eb ._acg +=int64 (_cdc );
return _cdc ,Error {ErrWrite ,"\u0075n\u0064\u0065\u0072\u006c\u0079\u0069\u006e\u0067\u0020\u0077\u0072i\u0074\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0020\u0028"+_bdb .Error ()+"\u0029",int64 (_gef )};};_gef =_dcfe ;};_eb ._acg +=int64 (_cdc );if _cdc !=_fcd {_dea =Error {ErrWrite ,"\u0062\u0079t\u0065\u0073\u0020\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0064\u006f\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0077\u0072\u0069\u0074\u0065\u0020\u0073\u0069\u007a\u0065",int64 (_cdc )};
}else if _cdc < len (b ){_dea =_gf .EOF ;};return _cdc ,_dea ;};func (_efe *File )Seek (offset int64 ,whence int )(int64 ,error ){var _ccd int64 ;switch whence {default:return 0,Error {ErrSeek ,"\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077h\u0065\u006e\u0063\u0065",int64 (whence )};
case 0:_ccd =offset ;case 1:_ccd =_efe ._acg +offset ;case 2:_ccd =_efe .Size -offset ;};switch {case _ccd < 0:return _efe ._acg ,Error {ErrSeek ,"\u0063\u0061\u006e'\u0074\u0020\u0073\u0065e\u006b\u0020\u0062\u0065\u0066\u006f\u0072e\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u006f\u0066\u0020\u0046\u0069\u006c\u0065",_ccd };
case _ccd >=_efe .Size :return _efe ._acg ,Error {ErrSeek ,"c\u0061\u006e\u0027\u0074\u0020\u0073e\u0065\u006b\u0020\u0070\u0061\u0073\u0074\u0020\u0046i\u006c\u0065\u0020l\u0065n\u0067\u0074\u0068",_ccd };case _ccd ==_efe ._acg :return _ccd ,nil ;case _ccd > _efe ._acg :_ggb :=_efe ._acg ;
_efe ._acg =_ccd ;return _efe ._acg ,_efe .seek (_ccd -_ggb );};if _efe ._aba >=_efe ._acg -_ccd {_efe ._aba =_efe ._aba -(_efe ._acg -_ccd );_efe ._acg =_ccd ;return _efe ._acg ,nil ;};_efe ._aba =0;_efe ._cc =_efe ._ec ;_efe ._acg =_ccd ;return _efe ._acg ,_efe .seek (_ccd );
};func (_bfe *File )findLast (_fcc bool )(uint32 ,error ){_eba :=_bfe ._ec ;for {_dac ,_gfg :=_bfe ._adg .findNext (_eba ,_fcc );if _gfg !=nil {return 0,Error {ErrRead ,"\u0062\u0061\u0064\u0020\u0072\u0065\u0061\u0064\u0020\u0066i\u006e\u0064\u0069\u006e\u0067\u0020\u006ee\u0078\u0074\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0028"+_gfg .Error ()+"\u0029",0};
};if _dac ==_ffc {break ;};_eba =_dac ;};return _eba ,nil ;};const (_dgf uint64 =0xE11AB1A1E011CFD0;_gcbd uint32 =64;_cgdc int64 =4096;_accf uint32 =128;);func (_ace *File )ID ()string {return _ace ._bc .String ()};func (_gcb *File )mode ()_g .FileMode {if _gcb ._ed !=_f {return _g .ModeDir |0777;
};return 0666;};func (_dged *Reader )setDifats ()error {_dged ._adb ._fae =_dged ._adb ._gegc [:];if _dged ._adb ._feca ==0{return nil ;};_cddb :=(_dged ._cad /4)-1;_fada :=make ([]uint32 ,109,_dged ._adb ._feca *_cddb +109);copy (_fada ,_dged ._adb ._fae );
_dged ._adb ._fae =_fada ;_gaaf :=_dged ._adb ._cfee ;for _eag :=0;_eag < int (_dged ._adb ._feca );_eag ++{_acaa ,_afd :=_dged .readAt (_eae (_dged ._cad ,_gaaf ),int (_dged ._cad ));if _afd !=nil {return Error {ErrFormat ,"e\u0072r\u006f\u0072\u0020\u0073\u0065\u0074\u0074\u0069n\u0067\u0020\u0044\u0049FA\u0054\u0028"+_afd .Error ()+"\u0029",int64 (_gaaf )};
};for _fcg :=0;_fcg < int (_cddb );_fcg ++{_dged ._adb ._fae =append (_dged ._adb ._fae ,_ac .LittleEndian .Uint32 (_acaa [_fcg *4:_fcg *4+4]));};_gaaf =_ac .LittleEndian .Uint32 (_acaa [len (_acaa )-4:]);};return nil ;};const _gfb int =8+16+10+6+12+8+16+109*4;
type Reader struct{_abb bool ;_cad uint32 ;_beb []byte ;_adb *header ;File []*File ;_gcd []*File ;_dbd int ;_dcac _gf .ReaderAt ;_cddg _gf .WriterAt ;};type directoryEntryFields struct{_bd [32]uint16 ;_cdb uint16 ;_ed uint8 ;_dcf uint8 ;_ge uint32 ;_fe uint32 ;
_df uint32 ;_bc _ce .Guid ;_ad [4]byte ;_gg _ce .FileTime ;_af _ce .FileTime ;_ec uint32 ;_ca [8]byte ;};func (_ffbd *Reader )findNext (_ccec uint32 ,_fbc bool )(uint32 ,error ){_efed :=_ffbd ._cad /4;_ffgc :=int (_ccec /_efed );var _gbcc uint32 ;if _fbc {if _ffgc < 0||_ffgc >=len (_ffbd ._adb ._fef ){return 0,Error {ErrRead ,"\u006d\u0069\u006e\u0069\u0073e\u0063\u0074\u006f\u0072\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u0069\u0073 \u006f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u006d\u0069\u006e\u0069\u0046\u0041\u0054\u0020\u0072\u0061\u006e\u0067\u0065",int64 (_ffgc )};
};_gbcc =_ffbd ._adb ._fef [_ffgc ];}else {if _ffgc < 0||_ffgc >=len (_ffbd ._adb ._fae ){return 0,Error {ErrRead ,"\u0046\u0041\u0054\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u0069\u0073\u0020\u006f\u0075t\u0073i\u0064\u0065\u0020\u0044\u0049\u0046\u0041\u0054\u0020\u0072\u0061\u006e\u0067\u0065",int64 (_ffgc )};
};_gbcc =_ffbd ._adb ._fae [_ffgc ];};_age :=_ccec %_efed ;_agd :=_eae (_ffbd ._cad ,_gbcc )+int64 (_age *4);_adca ,_dbf :=_ffbd .readAt (_agd ,4);if _dbf !=nil {return 0,Error {ErrRead ,"\u0062\u0061\u0064\u0020\u0072\u0065\u0061\u0064\u0020\u0066i\u006e\u0064\u0069\u006e\u0067\u0020\u006ee\u0078\u0074\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0028"+_dbf .Error ()+"\u0029",_agd };
};_cfac :=_ac .LittleEndian .Uint32 (_adca );return _cfac ,nil ;};type fileInfo struct{*File };func (_gga *File )changeSize (_fea int64 )error {if _fea ==_gga .Size {return nil ;};var _ebe *File ;for _ ,_bdc :=range _gga ._adg ._gcd {if _bdc .Name ==_gga .Name {_ebe =_bdc ;
break ;};};if _ebe ==nil {return _d .Errorf ("\u004e\u006f\u0020\u0064\u0069\u0072e\u0063\u0074\u006f\u0072\u0079\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066o\u0072\u0020\u0061\u0020\u0066\u0069\u006ce\u003a\u0020\u0025\u0073",_gga .Name );};_gbc :=_e .NewBuffer ([]byte {});
if _cbgg :=_ac .Write (_gbc ,_ac .LittleEndian ,_fea );_cbgg !=nil {return _cbgg ;};for _cgd ,_def :=range _gbc .Bytes (){_ebe ._ca [_cgd ]=_def ;};var _fed int64 ;var _dgd bool ;if _gga .Size < _cgdc {_dgd =true ;_fed =int64 (_gcbd );}else {_fed =int64 (_gga ._adg ._cad );
};_caa :=int ((_gga .Size -1)/_fed )+1;_cfe :=int ((_fea -1)/_fed )+1;if _caa < _cfe {_dagf ,_fb :=_gga .findLast (_dgd );if _fb !=nil {return _fb ;};_gfe ,_fb :=_gga ._adg .findNextFreeSector (_dgd );if _fb !=nil {return _fb ;};for _fee :=_cfe -_caa ;
_fee > 0;_fee --{if _fac :=_gga ._adg .saveToFatLocs (_dagf ,_gfe ,_dgd );_fac !=nil {return _fac ;};if _fee > 1{_dagf =_gfe ;_gfe ++;}else if _ga :=_gga ._adg .saveToFatLocs (_gfe ,_ffc ,_dgd );_ga !=nil {return _ga ;};};}else if _caa > _cfe {_bce :=_gga ._ec ;
var _ff error ;for _acaf :=0;_acaf < _cfe -1;_acaf ++{_bce ,_ff =_gga ._adg .findNext (_bce ,_dgd );if _ff !=nil {return _ff ;};};if _face :=_gga ._adg .saveToFatLocs (_bce ,_ffc ,_dgd );_face !=nil {return _face ;};};_gga .Size =_fea ;return nil ;};func (_ceca *Reader )exportMiniStream ()(*_dd .Writer ,*_dd .Writer ,error ){_egd ,_egda :=_dd .NewWriter (),_dd .NewWriter ();
_fdg :=uint32 (0);for _ ,_adda :=range _ceca .File {if _adda .Size ==0{continue ;};_adda .reset ();_adda ._ec =_fdg ;_gcbdc :=int (_adda .Size )/int (_gcbd );if int (_adda .Size )%int (_gcbd )!=0{_gcbdc ++;};for _ebd :=1;_ebd < _gcbdc ;_ebd ++{_fdg ++;
if _bcf :=_ac .Write (_egd ,_ac .LittleEndian ,_fdg );_bcf !=nil {return nil ,nil ,_bcf ;};};if _acec :=_ac .Write (_egd ,_ac .LittleEndian ,_ffc );_acec !=nil {return nil ,nil ,_acec ;};_fdg ++;if _ ,_eff :=_gf .Copy (_egda ,_adda );_eff !=nil {return nil ,nil ,_eff ;
};if _gbdf :=_egda .AlignLength (64);_gbdf !=nil {return nil ,nil ,_gbdf ;};};if _gbce :=_egd .FillWithByte (int (_ceca ._cad )-_egd .Len (),_edb );_gbce !=nil {return nil ,nil ,_gbce ;};if _ccefd :=_egda .AlignLength (int (_ceca ._cad ));_ccefd !=nil {return nil ,nil ,_ccefd ;
};return _egd ,_egda ,nil ;};func (_gaad *Reader )exportHeader (_fcfe *_dd .Writer )error {if _ceb :=_ac .Write (_fcfe ,_ac .LittleEndian ,&_gaad ._adb ._ba );_ceb !=nil {return _ceb ;};if _faf :=_fcfe .Skip (16);_faf !=nil {return _faf ;};if _edg :=_ac .Write (_fcfe ,_ac .LittleEndian ,&_gaad ._adb ._aefc );
_edg !=nil {return _edg ;};if _eabg :=_ac .Write (_fcfe ,_ac .LittleEndian ,&_gaad ._adb ._efg );_eabg !=nil {return _eabg ;};if _cade :=_ac .Write (_fcfe ,_ac .LittleEndian ,uint16 (0xfffe));_cade !=nil {return _cade ;};if _gff :=_ac .Write (_fcfe ,_ac .LittleEndian ,&_gaad ._adb ._dff );
_gff !=nil {return _gff ;};if _gdec :=_ac .Write (_fcfe ,_ac .LittleEndian ,uint16 (0x0006));_gdec !=nil {return _gdec ;};if _efeg :=_fcfe .Skip (6);_efeg !=nil {return _efeg ;};if _dba :=_ac .Write (_fcfe ,_ac .LittleEndian ,&_gaad ._adb ._cced );_dba !=nil {return _dba ;
};if _bgce :=_ac .Write (_fcfe ,_ac .LittleEndian ,&_gaad ._adb ._afa );_bgce !=nil {return _bgce ;};if _dgc :=_ac .Write (_fcfe ,_ac .LittleEndian ,&_gaad ._adb ._gbb );_dgc !=nil {return _dgc ;};if _beg :=_fcfe .Skip (4);_beg !=nil {return _beg ;};if _bdbg :=_ac .Write (_fcfe ,_ac .LittleEndian ,uint32 (0x00001000));
_bdbg !=nil {return _bdbg ;};if _gcgd :=_ac .Write (_fcfe ,_ac .LittleEndian ,&_gaad ._adb ._cdd );_gcgd !=nil {return _gcgd ;};if _bbfc :=_ac .Write (_fcfe ,_ac .LittleEndian ,&_gaad ._adb ._ccef );_bbfc !=nil {return _bbfc ;};if _afb :=_ac .Write (_fcfe ,_ac .LittleEndian ,&_gaad ._adb ._cfee );
_afb !=nil {return _afb ;};if _gdcd :=_ac .Write (_fcfe ,_ac .LittleEndian ,&_gaad ._adb ._feca );_gdcd !=nil {return _gdcd ;};for _ceg :=0;_ceg < 109;_ceg ++{if _bbe :=_ac .Write (_fcfe ,_ac .LittleEndian ,&_gaad ._adb ._gegc [_ceg ]);_bbe !=nil {return _bbe ;
};};return nil ;};func (_bec *File )ReadAt (p []byte ,off int64 )(_aac int ,_abd error ){_gde ,_cba ,_ebc :=_bec ._acg ,_bec ._aba ,_bec ._cc ;_ ,_abd =_bec .Seek (off ,0);if _abd ==nil {_aac ,_abd =_bec .Read (p );};_bec ._acg ,_bec ._aba ,_bec ._cc =_gde ,_cba ,_ebc ;
return _aac ,_abd ;};func (_dcfb *Reader )traverse ()error {_dcfb .File =make ([]*File ,0,len (_dcfb ._gcd ));var (_aed func (int ,[]string );_bgc error ;_bgf int ;);_aed =func (_aef int ,_cec []string ){_bgf ++;if _bgf > len (_dcfb ._gcd ){_bgc =Error {ErrTraverse ,"\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c\u0020\u0063o\u0075\u006e\u0074\u0065\u0072\u0020\u006f\u0076\u0065\u0072f\u006c\u006f\u0077",int64 (_aef )};
return ;};if _aef < 0||_aef >=len (_dcfb ._gcd ){_bgc =Error {ErrTraverse ,"\u0069\u006c\u006ceg\u0061\u006c\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c\u0020\u0069\u006e\u0064\u0065\u0078",int64 (_aef )};return ;};_bf :=_dcfb ._gcd [_aef ];
if _bf ._ge !=_dcfeg {_aed (int (_bf ._ge ),_cec );};_dcfb .File =append (_dcfb .File ,_bf );_bf .Path =_cec ;if _bf ._df !=_dcfeg {if _aef > 0{_aed (int (_bf ._df ),append (_cec ,_bf .Name ));}else {_aed (int (_bf ._df ),_cec );};};if _bf ._fe !=_dcfeg {_aed (int (_bf ._fe ),_cec );
};return ;};_aed (0,[]string {});return _bgc ;};type headerFields struct{_ba uint64 ;_ [16]byte ;_aefc uint16 ;_efg uint16 ;_ [2]byte ;_dff uint16 ;_ [2]byte ;_ [6]byte ;_cced uint32 ;_afa uint32 ;_gbb uint32 ;_ [4]byte ;_ [4]byte ;_cdd uint32 ;_ccef uint32 ;
_cfee uint32 ;_feca uint32 ;_gegc [109]uint32 ;};func (_gd *Reader )setDirEntries ()error {_aa :=20;if _gd ._adb ._cced > 0{_aa =int (_gd ._adb ._cced );};_de :=make ([]*File ,0,_aa );_cag :=make (map[uint32 ]bool );_bg :=int (_gd ._cad /_accf );_cbg :=_gd ._adb ._gbb ;
for _cbg !=_ffc {_cbb ,_ggc :=_gd .readAt (_eae (_gd ._cad ,_cbg ),int (_gd ._cad ));if _ggc !=nil {return Error {ErrRead ,"\u0064\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u0020e\u006e\u0074\u0072\u0069\u0065\u0073\u0020r\u0065\u0061\u0064\u0020\u0065\u0072\u0072\u006f\u0072\u0020\u0028"+_ggc .Error ()+"\u0029",_eae (_gd ._cad ,_cbg )};
};for _bb :=0;_bb < _bg ;_bb ++{_aca :=&File {_adg :_gd };_aca .directoryEntryFields =_fd (_cbb [_bb *int (_accf ):]);_gdc (_gd ._adb ._efg ,_aca );_aca ._cc =_aca ._ec ;_de =append (_de ,_aca );};_da ,_ggc :=_gd .findNext (_cbg ,false );if _ggc !=nil {return Error {ErrRead ,"d\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0065r\u0072\u006f\u0072\u0020\u0066\u0069\u006e\u0064\u0069\u006eg \u0073\u0065\u0063t\u006fr\u0020\u0028"+_ggc .Error ()+"\u0029",int64 (_da )};
};if _da <=_cbg {if _da ==_cbg ||_cag [_da ]{return Error {ErrRead ,"\u0064\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u0020e\u006e\u0074\u0072\u0069\u0065\u0073\u0020s\u0065\u0063\u0074\u006f\u0072\u0020\u0063\u0079\u0063\u006c\u0065",int64 (_da )};
};_cag [_da ]=true ;};_cbg =_da ;};_gd ._gcd =_de ;return nil ;};func (_feg *Reader )exportDifats (_gec *_dd .Writer )error {if _feg ._adb ._feca ==0{return nil ;};return nil ;};func (_bbdb *Reader )exportDirEntries (_daf *_dd .Writer )error {if int64 (_daf .Len ())!=_eae (_bbdb ._cad ,_bbdb ._adb ._gbb ){return Error {ErrWrite ,_d .Sprintf ("I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0077\u0072\u0069\u0074\u0065\u0072\u0020l\u0065\u006e\u0067t\u0068:\u0020\u0025\u0076",_daf .Len ()),0};
};for _ ,_bdg :=range _bbdb ._gcd {_dfa ,_cfeg :=_cgb (_bdg .directoryEntryFields );if _cfeg !=nil {return _cfeg ;};if _ ,_dagd :=_gf .Copy (_daf ,_dfa );_dagd !=nil {return _dagd ;};};return nil ;};type slicer interface{Slice (_bdf int64 ,_aceg int )([]byte ,error );
};type Error struct{_ecbb int ;_addb string ;_ggbe int64 ;};func _cgb (_gdf *directoryEntryFields )(*_e .Buffer ,error ){_gcg :=_e .NewBuffer ([]byte {});for _ ,_ddef :=range _gdf ._bd {if _fdb :=_ac .Write (_gcg ,_ac .LittleEndian ,&_ddef );_fdb !=nil {return nil ,_fdb ;
};};if _cce :=_ac .Write (_gcg ,_ac .LittleEndian ,&_gdf ._cdb );_cce !=nil {return nil ,_cce ;};if _cca :=_ac .Write (_gcg ,_ac .LittleEndian ,&_gdf ._ed );_cca !=nil {return nil ,_cca ;};if _ede :=_ac .Write (_gcg ,_ac .LittleEndian ,&_gdf ._dcf );_ede !=nil {return nil ,_ede ;
};if _aea :=_ac .Write (_gcg ,_ac .LittleEndian ,&_gdf ._ge );_aea !=nil {return nil ,_aea ;};if _cbd :=_ac .Write (_gcg ,_ac .LittleEndian ,&_gdf ._fe );_cbd !=nil {return nil ,_cbd ;};if _aacg :=_ac .Write (_gcg ,_ac .LittleEndian ,&_gdf ._df );_aacg !=nil {return nil ,_aacg ;
};if _dca :=_ac .Write (_gcg ,_ac .LittleEndian ,&_gdf ._bc .DataA );_dca !=nil {return nil ,_dca ;};if _aefe :=_ac .Write (_gcg ,_ac .LittleEndian ,&_gdf ._bc .DataB );_aefe !=nil {return nil ,_aefe ;};if _bff :=_ac .Write (_gcg ,_ac .LittleEndian ,&_gdf ._bc .DataC );
_bff !=nil {return nil ,_bff ;};if _ ,_ege :=_gcg .Write (_gdf ._bc .DataD [:]);_ege !=nil {return nil ,_ege ;};if _ ,_gdg :=_gcg .Write (_gdf ._ad [:]);_gdg !=nil {return nil ,_gdg ;};if _gad :=_ac .Write (_gcg ,_ac .LittleEndian ,&_gdf ._gg .Low );_gad !=nil {return nil ,_gad ;
};if _ggg :=_ac .Write (_gcg ,_ac .LittleEndian ,&_gdf ._gg .High );_ggg !=nil {return nil ,_ggg ;};if _aff :=_ac .Write (_gcg ,_ac .LittleEndian ,&_gdf ._af .Low );_aff !=nil {return nil ,_aff ;};if _deb :=_ac .Write (_gcg ,_ac .LittleEndian ,&_gdf ._af .High );
_deb !=nil {return nil ,_deb ;};if _ddb :=_ac .Write (_gcg ,_ac .LittleEndian ,&_gdf ._ec );_ddb !=nil {return nil ,_ddb ;};if _ ,_dcfc :=_gcg .Write (_gdf ._ca [:]);_dcfc !=nil {return nil ,_dcfc ;};return _gcg ,nil ;};func (_ageg *Reader )ID ()string {return _ageg .File [0].ID ()};
func (_fg fileInfo )Name ()string {return _fg .File .Name };