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

package mscfb ;import (_b "bytes";_gb "encoding/binary";_ae "fmt";_ed "github.com/richardlehane/msoleps/types";_dd "github.com/unidoc/unioffice/internal/mscfb/rw";_c "io";_e "os";_a "strconv";_g "time";_f "unicode";_db "unicode/utf16";);type slicer interface{Slice (_gec int64 ,_cbbf int )([]byte ,error );
};func _cbdd (_bbf *directoryEntryFields )(*_b .Buffer ,error ){_cbde :=_b .NewBuffer ([]byte {});for _ ,_abb :=range _bbf ._dc {if _begd :=_gb .Write (_cbde ,_gb .LittleEndian ,&_abb );_begd !=nil {return nil ,_begd ;};};if _egc :=_gb .Write (_cbde ,_gb .LittleEndian ,&_bbf ._eb );
_egc !=nil {return nil ,_egc ;};if _eec :=_gb .Write (_cbde ,_gb .LittleEndian ,&_bbf ._fe );_eec !=nil {return nil ,_eec ;};if _ba :=_gb .Write (_cbde ,_gb .LittleEndian ,&_bbf ._edd );_ba !=nil {return nil ,_ba ;};if _ggg :=_gb .Write (_cbde ,_gb .LittleEndian ,&_bbf ._fb );
_ggg !=nil {return nil ,_ggg ;};if _ffe :=_gb .Write (_cbde ,_gb .LittleEndian ,&_bbf ._gg );_ffe !=nil {return nil ,_ffe ;};if _ccgg :=_gb .Write (_cbde ,_gb .LittleEndian ,&_bbf ._gbe );_ccgg !=nil {return nil ,_ccgg ;};if _fdf :=_gb .Write (_cbde ,_gb .LittleEndian ,&_bbf ._ged .DataA );
_fdf !=nil {return nil ,_fdf ;};if _dcde :=_gb .Write (_cbde ,_gb .LittleEndian ,&_bbf ._ged .DataB );_dcde !=nil {return nil ,_dcde ;};if _fbe :=_gb .Write (_cbde ,_gb .LittleEndian ,&_bbf ._ged .DataC );_fbe !=nil {return nil ,_fbe ;};if _ ,_gefb :=_cbde .Write (_bbf ._ged .DataD [:]);
_gefb !=nil {return nil ,_gefb ;};if _ ,_defa :=_cbde .Write (_bbf ._bf [:]);_defa !=nil {return nil ,_defa ;};if _fbeg :=_gb .Write (_cbde ,_gb .LittleEndian ,&_bbf ._bg .Low );_fbeg !=nil {return nil ,_fbeg ;};if _ebc :=_gb .Write (_cbde ,_gb .LittleEndian ,&_bbf ._bg .High );
_ebc !=nil {return nil ,_ebc ;};if _ddc :=_gb .Write (_cbde ,_gb .LittleEndian ,&_bbf ._aa .Low );_ddc !=nil {return nil ,_ddc ;};if _abf :=_gb .Write (_cbde ,_gb .LittleEndian ,&_bbf ._aa .High );_abf !=nil {return nil ,_abf ;};if _edg :=_gb .Write (_cbde ,_gb .LittleEndian ,&_bbf ._ce );
_edg !=nil {return nil ,_edg ;};if _ ,_ffd :=_cbde .Write (_bbf ._edb [:]);_ffd !=nil {return nil ,_ffd ;};return _cbde ,nil ;};func _bb (_bc []byte )*directoryEntryFields {_deb :=&directoryEntryFields {};for _gf :=range _deb ._dc {_deb ._dc [_gf ]=_gb .LittleEndian .Uint16 (_bc [_gf *2:_gf *2+2]);
};_deb ._eb =_gb .LittleEndian .Uint16 (_bc [64:66]);_deb ._fe =uint8 (_bc [66]);_deb ._edd =uint8 (_bc [67]);_deb ._fb =_gb .LittleEndian .Uint32 (_bc [68:72]);_deb ._gg =_gb .LittleEndian .Uint32 (_bc [72:76]);_deb ._gbe =_gb .LittleEndian .Uint32 (_bc [76:80]);
_deb ._ged =_ed .MustGuid (_bc [80:96]);copy (_deb ._bf [:],_bc [96:100]);_deb ._bg =_ed .MustFileTime (_bc [100:108]);_deb ._aa =_ed .MustFileTime (_bc [108:116]);_deb ._ce =_gb .LittleEndian .Uint32 (_bc [116:120]);copy (_deb ._edb [:],_bc [120:128]);
return _deb ;};func (_fge *Reader )exportHeader (_ceee *_dd .Writer )error {if _dgdgf :=_gb .Write (_ceee ,_gb .LittleEndian ,&_fge ._ffde ._gbf );_dgdgf !=nil {return _dgdgf ;};if _gged :=_ceee .Skip (16);_gged !=nil {return _gged ;};if _ggec :=_gb .Write (_ceee ,_gb .LittleEndian ,&_fge ._ffde ._bcg );
_ggec !=nil {return _ggec ;};if _gfce :=_gb .Write (_ceee ,_gb .LittleEndian ,&_fge ._ffde ._afe );_gfce !=nil {return _gfce ;};if _gba :=_gb .Write (_ceee ,_gb .LittleEndian ,uint16 (0xfffe));_gba !=nil {return _gba ;};if _gbae :=_gb .Write (_ceee ,_gb .LittleEndian ,&_fge ._ffde ._fbg );
_gbae !=nil {return _gbae ;};if _gead :=_gb .Write (_ceee ,_gb .LittleEndian ,uint16 (0x0006));_gead !=nil {return _gead ;};if _cbcc :=_ceee .Skip (6);_cbcc !=nil {return _cbcc ;};if _bfcf :=_gb .Write (_ceee ,_gb .LittleEndian ,&_fge ._ffde ._gag );_bfcf !=nil {return _bfcf ;
};if _dcdg :=_gb .Write (_ceee ,_gb .LittleEndian ,&_fge ._ffde ._bdba );_dcdg !=nil {return _dcdg ;};if _agb :=_gb .Write (_ceee ,_gb .LittleEndian ,&_fge ._ffde ._add );_agb !=nil {return _agb ;};if _cfeb :=_ceee .Skip (4);_cfeb !=nil {return _cfeb ;
};if _cdbf :=_gb .Write (_ceee ,_gb .LittleEndian ,uint32 (0x00001000));_cdbf !=nil {return _cdbf ;};if _cbfg :=_gb .Write (_ceee ,_gb .LittleEndian ,&_fge ._ffde ._cfb );_cbfg !=nil {return _cbfg ;};if _cag :=_gb .Write (_ceee ,_gb .LittleEndian ,&_fge ._ffde ._fea );
_cag !=nil {return _cag ;};if _cdf :=_gb .Write (_ceee ,_gb .LittleEndian ,&_fge ._ffde ._cce );_cdf !=nil {return _cdf ;};if _cbag :=_gb .Write (_ceee ,_gb .LittleEndian ,&_fge ._ffde ._ca );_cbag !=nil {return _cbag ;};for _dcfg :=0;_dcfg < 109;_dcfg ++{if _cdfc :=_gb .Write (_ceee ,_gb .LittleEndian ,&_fge ._ffde ._begc [_dcfg ]);
_cdfc !=nil {return _cdfc ;};};return nil ;};func (_gacc *File )FileInfo ()_e .FileInfo {return fileInfo {_gacc }};func (_cgcc *Reader )saveToFatLocs (_egb uint32 ,_cfg interface{},_bgd bool )error {_bebg :=_b .NewBuffer ([]byte {});if _bdbf :=_gb .Write (_bebg ,_gb .LittleEndian ,_cfg );
_bdbf !=nil {return _bdbf ;};_abfe :=_cgcc .findFatLocsOffset (_bgd )+int64 (_egb *4);_ ,_effg :=_cgcc ._cfcf .WriteAt (_bebg .Bytes (),_abfe );return _effg ;};func (_feb *File )stream (_fd int )([][2]int64 ,error ){var _fadd bool ;var _bccb int ;var _fbf int64 ;
if _feb .Size < _ccf {_fadd =true ;_bccb =_fd /int (_eab )+2;_fbf =int64 (_eab );}else {_bccb =_fd /int (_feb ._dea ._gagd )+2;_fbf =int64 (_feb ._dea ._gagd );};_fdd :=make ([][2]int64 ,0,_bccb );var _gab ,_aeag int ;if _feb ._bca > 0{_dcgg ,_afc :=_feb ._dea .getOffset (_feb ._cf ,_fadd );
if _afc !=nil {return nil ,_afc ;};if _fbf -_feb ._bca >=int64 (_fd ){_fdd =append (_fdd ,[2]int64 {_dcgg +_feb ._bca ,int64 (_fd )});}else {_fdd =append (_fdd ,[2]int64 {_dcgg +_feb ._bca ,_fbf -_feb ._bca });};if _fbf -_feb ._bca <=int64 (_fd ){_feb ._cf ,_afc =_feb ._dea .findNext (_feb ._cf ,_fadd );
if _afc !=nil {return nil ,_afc ;};_aeag +=int (_fbf -_feb ._bca );_feb ._bca =0;}else {_feb ._bca +=int64 (_fd );};if _fdd [0][1]==int64 (_fd ){return _fdd ,nil ;};if _feb ._cf ==_gcb {return nil ,Error {ErrRead ,"\u0075\u006ee\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0065\u0061\u0072\u006c\u0079\u0020\u0065\u006e\u0064\u0020\u006f\u0066\u0020\u0063ha\u0069\u006e",int64 (_feb ._cf )};
};_gab ++;};for {if _gab >=cap (_fdd ){return nil ,Error {ErrRead ,"\u0069\u006e\u0064\u0065x\u0020\u006f\u0076\u0065\u0072\u0072\u0075\u006e\u0073\u0020s\u0065c\u0074\u006f\u0072\u0020\u006c\u0065\u006eg\u0074\u0068",int64 (_gab )};};_acc ,_gga :=_feb ._dea .getOffset (_feb ._cf ,_fadd );
if _gga !=nil {return nil ,_gga ;};if _fd -_aeag < int (_fbf ){_fdd =append (_fdd ,[2]int64 {_acc ,int64 (_fd -_aeag )});_feb ._bca =int64 (_fd -_aeag );return _cbc (_fdd ),nil ;}else {_fdd =append (_fdd ,[2]int64 {_acc ,_fbf });_aeag +=int (_fbf );_feb ._cf ,_gga =_feb ._dea .findNext (_feb ._cf ,_fadd );
if _gga !=nil {return nil ,_gga ;};if _aeag ==_fd {return _cbc (_fdd ),nil ;};};_gab ++;};};func (_bcgf *Reader )GetHeader ()*header {return _bcgf ._ffde };func (_def fileInfo )IsDir ()bool {return _def .mode ().IsDir ()};const (_ge uint8 =0x0;_fg uint8 =0x1;
);type Reader struct{_acb bool ;_gagd uint32 ;_fff []byte ;_ffde *header ;File []*File ;_fbfb []*File ;_fbfd int ;_fbb _c .ReaderAt ;_cfcf _c .WriterAt ;};func (_adbe *Reader )exportFAT (_bgg *_dd .Writer ,_eac []uint32 )error {if _adbe ._ffde ._bdba ==0{return nil ;
};_bcad :=_b .NewBuffer ([]byte {});if _gbeec :=_gb .Write (_bcad ,_gb .LittleEndian ,_cbaa );_gbeec !=nil {return _gbeec ;};for _gaacg :=0;_gaacg < len (_eac )-1;_gaacg ++{for _cceg :=_eac [_gaacg ];_cceg < _eac [_gaacg +1]-1;_cceg ++{if _fdg :=_gb .Write (_bcad ,_gb .LittleEndian ,_cceg );
_fdg !=nil {return _fdg ;};};if _dae :=_gb .Write (_bcad ,_gb .LittleEndian ,_gcb );_dae !=nil {return _dae ;};};_cbbg :=512;for _ ,_dadb :=range _bcad .Bytes (){if _baec :=_bgg .WriteByteAt (_dadb ,_cbbg );_baec !=nil {return _baec ;};_cbbg ++;};return nil ;
};func (_bed *File )WriteAt (p []byte ,off int64 )(_eag int ,_aaf error ){_ded ,_cbd ,_bde :=_bed ._ccc ,_bed ._bca ,_bed ._cf ;_ ,_aaf =_bed .Seek (off ,0);if _aaf ==nil {_eag ,_aaf =_bed .Write (p );};_bed ._ccc ,_bed ._bca ,_bed ._cf =_ded ,_cbd ,_bde ;
return _eag ,_aaf ;};type fileInfo struct{*File };func (_age *File )ensureWriterAt ()error {if _age ._dea ._cfcf ==nil {_bec ,_ggb :=_age ._dea ._fbb .(_c .WriterAt );if !_ggb {return Error {ErrWrite ,"\u006d\u0073\u0063\u0066\u0062\u002e\u004ee\u0077\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0067\u0069\u0076\u0065n\u0020R\u0065\u0061\u0064\u0065\u0072\u0041t\u0020\u0063\u006f\u006e\u0076\u0065\u0072t\u0069\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0020\u0069\u006f\u002e\u0057\u0072\u0069\u0074\u0065\u0072\u0041\u0074\u0020\u0069n\u0020\u006f\u0072\u0064\u0065\u0072\u0020\u0074\u006f\u0020\u0077\u0072\u0069t\u0065",0};
};_age ._dea ._cfcf =_bec ;};return nil ;};func _cbc (_cgd [][2]int64 )[][2]int64 {_ad :=len (_cgd );for _dffc ,_beg :=0,0;_dffc < _ad &&_beg +1< len (_cgd );_dffc ++{if _cgd [_beg ][0]+_cgd [_beg ][1]==_cgd [_beg +1][0]{_cgd [_beg ][1]=_cgd [_beg ][1]+_cgd [_beg +1][1];
for _bgef :=range _cgd [_beg +1:len (_cgd )-1]{_cgd [_beg +1+_bgef ]=_cgd [_bgef +_beg +2];};_cgd =_cgd [:len (_cgd )-1];}else {_beg +=1;};};return _cgd ;};type Error struct{_faf int ;_bba string ;_edbd int64 ;};const (_bag uint64 =0xE11AB1A1E011CFD0;_eab uint32 =64;
_ccf int64 =4096;_eff uint32 =128;);func (_fbd Error )Typ ()int {return _fbd ._faf };func (_ff *File )Created ()_g .Time {return _ff ._bg .Time ()};func (_bbc *Reader )setHeader ()error {_bfba ,_eea :=_bbc .readAt (0,_aafc );if _eea !=nil {return _eea ;
};_bbc ._ffde =&header {headerFields :_cee (_bfba )};if _bbc ._ffde ._gbf !=_bag {return Error {ErrFormat ,"\u0062\u0061\u0064\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065",int64 (_bbc ._ffde ._gbf )};};if _bbc ._ffde ._fbg ==0x0009||_bbc ._ffde ._fbg ==0x000c{_bbc ._gagd =uint32 (1<<_bbc ._ffde ._fbg );
}else {return Error {ErrFormat ,"\u0069\u006c\u006c\u0065ga\u006c\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0073\u0069\u007a\u0065",int64 (_bbc ._ffde ._fbg )};};if _bbc ._ffde ._ca > 0{_cege :=(_bbc ._gagd /4)-1;if int (_bbc ._ffde ._ca *_cege +109)< 0{return Error {ErrFormat ,"\u0044I\u0046A\u0054\u0020\u0069\u006e\u0074 \u006f\u0076e\u0072\u0066\u006c\u006f\u0077",int64 (_bbc ._ffde ._ca )};
};if _bbc ._ffde ._ca *_cege +109> _bbc ._ffde ._bdba +_cege {return Error {ErrFormat ,"\u006e\u0075\u006d\u0020\u0044\u0049\u0046\u0041\u0054\u0073 \u0065\u0078\u0063\u0065\u0065\u0064\u0073 \u0046\u0041\u0054\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0073",int64 (_bbc ._ffde ._ca )};
};};if _bbc ._ffde ._fea > 0{if int (_bbc ._gagd /4*_bbc ._ffde ._fea )< 0{return Error {ErrFormat ,"m\u0069\u006e\u0069\u0020FA\u0054 \u0069\u006e\u0074\u0020\u006fv\u0065\u0072\u0066\u006c\u006f\u0077",int64 (_bbc ._ffde ._fea )};};if _bbc ._ffde ._fea > _bbc ._ffde ._bdba *(_bbc ._gagd /_eab ){return Error {ErrFormat ,"\u006e\u0075\u006d\u0020\u006d\u0069n\u0069\u0020\u0046\u0041\u0054\u0073\u0020\u0065\u0078\u0063\u0065\u0065\u0064s\u0020\u0046\u0041\u0054\u0020\u0073\u0065c\u0074\u006f\u0072\u0073",int64 (_bbc ._ffde ._bdba )};
};};return nil ;};func (_fcec *Reader )setDifats ()error {_fcec ._ffde ._gbd =_fcec ._ffde ._begc [:];if _fcec ._ffde ._ca ==0{return nil ;};_bab :=(_fcec ._gagd /4)-1;_bdc :=make ([]uint32 ,109,_fcec ._ffde ._ca *_bab +109);copy (_bdc ,_fcec ._ffde ._gbd );
_fcec ._ffde ._gbd =_bdc ;_dagf :=_fcec ._ffde ._cce ;for _aab :=0;_aab < int (_fcec ._ffde ._ca );_aab ++{_cff ,_bgb :=_fcec .readAt (_cbf (_fcec ._gagd ,_dagf ),int (_fcec ._gagd ));if _bgb !=nil {return Error {ErrFormat ,"e\u0072r\u006f\u0072\u0020\u0073\u0065\u0074\u0074\u0069n\u0067\u0020\u0044\u0049FA\u0054\u0028"+_bgb .Error ()+"\u0029",int64 (_dagf )};
};for _fgfd :=0;_fgfd < int (_bab );_fgfd ++{_fcec ._ffde ._gbd =append (_fcec ._ffde ._gbd ,_gb .LittleEndian .Uint32 (_cff [_fgfd *4:_fgfd *4+4]));};_dagf =_gb .LittleEndian .Uint32 (_cff [len (_cff )-4:]);};return nil ;};func (_ag *File )mode ()_e .FileMode {if _ag ._fe !=_cg {return _e .ModeDir |0777;
};return 0666;};func (_daga *Reader )exportDifats (_cdbc *_dd .Writer )error {if _daga ._ffde ._ca ==0{return nil ;};return nil ;};func (_fcb *Reader )findNextFreeSector (_dgf bool )(uint32 ,error ){_ade :=_fcb .findFatLocsOffset (_dgf );_bcbf :=uint32 (0);
_bae :=_fcb ._gagd /4;for {_cbea ,_cfc :=_fcb .readAt (_ade ,4);if _cfc !=nil {return 0,Error {ErrRead ,"\u0062\u0061\u0064\u0020\u0072\u0065\u0061\u0064\u0020\u0066i\u006e\u0064\u0069\u006e\u0067\u0020\u006ee\u0078\u0074\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0028"+_cfc .Error ()+"\u0029",_ade };
};_fda :=_gb .LittleEndian .Uint32 (_cbea );if _fda ==_aafg {break ;};if _bcbf >=_bae {return 0,Error {ErrRead ,"\u0065\u006e\u0064\u0020of\u0020\u006d\u0069\u006e\u0069\u0046\u0061\u0074\u0020\u0072\u0065\u0061\u0063\u0068e\u0064",_ade };};_bcbf ++;_ade +=4;
};return _bcbf ,nil ;};func (_agff *Reader )Created ()_g .Time {return _agff .File [0].Created ()};const _aafc int =8+16+10+6+12+8+16+109*4;func _cbf (_eeg ,_ead uint32 )int64 {return int64 ((_ead +1)*_eeg )};func (_aeb *Reader )exportDirEntries (_bbg *_dd .Writer )error {if int64 (_bbg .Len ())!=_cbf (_aeb ._gagd ,_aeb ._ffde ._add ){return Error {ErrWrite ,_ae .Sprintf ("I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0077\u0072\u0069\u0074\u0065\u0072\u0020l\u0065\u006e\u0067t\u0068:\u0020\u0025\u0076",_bbg .Len ()),0};
};for _ ,_ebb :=range _aeb ._fbfb {_bee ,_fgf :=_cbdd (_ebb .directoryEntryFields );if _fgf !=nil {return _fgf ;};if _ ,_fbc :=_c .Copy (_bbg ,_bee );_fbc !=nil {return _fbc ;};};return nil ;};func New (ra _c .ReaderAt )(*Reader ,error ){_bgfe :=&Reader {_fbb :ra };
if _ ,_aagd :=ra .(slicer );_aagd {_bgfe ._acb =true ;}else {_bgfe ._fff =make ([]byte ,_aafc );};if _baba :=_bgfe .setHeader ();_baba !=nil {return nil ,_baba ;};if !_bgfe ._acb &&int (_bgfe ._gagd )> len (_bgfe ._fff ){_bgfe ._fff =make ([]byte ,_bgfe ._gagd );
};if _fab :=_bgfe .setDifats ();_fab !=nil {return nil ,_fab ;};if _ggbd :=_bgfe .setDirEntries ();_ggbd !=nil {return nil ,_ggbd ;};if _gda :=_bgfe .setMiniStream ();_gda !=nil {return nil ,_gda ;};if _aebg :=_bgfe .traverse ();_aebg !=nil {return nil ,_aebg ;
};return _bgfe ,nil ;};func (_fgd *Reader )findFatLocsOffset (_gdb bool )int64 {var _aabe uint32 ;if _gdb {_aabe =_fgd ._ffde ._cfe [0];}else {_aabe =_fgd ._ffde ._gbd [0];};return _cbf (_fgd ._gagd ,_aabe );};func (_fdc *Reader )Modified ()_g .Time {return _fdc .File [0].Modified ()};
func _edc (_df uint16 ,_gac *File ){_ddd (_gac );if _gac ._fe !=_cg {return ;};if _df > 3{_gac .Size =int64 (_gb .LittleEndian .Uint64 (_gac ._edb [:]));}else {_gac .Size =int64 (_gb .LittleEndian .Uint32 (_gac ._edb [:4]));};};type header struct{*headerFields ;
_gbd []uint32 ;_cfe []uint32 ;_fgca []uint32 ;};type File struct{Name string ;Initial uint16 ;Path []string ;Size int64 ;_ccc int64 ;_cf uint32 ;_bca int64 ;*directoryEntryFields ;_dea *Reader ;};func (_gaf fileInfo )Sys ()interface{}{return nil };type headerFields struct{_gbf uint64 ;
_ [16]byte ;_bcg uint16 ;_afe uint16 ;_ [2]byte ;_fbg uint16 ;_ [2]byte ;_ [6]byte ;_gag uint32 ;_bdba uint32 ;_add uint32 ;_ [4]byte ;_ [4]byte ;_cfb uint32 ;_fea uint32 ;_cce uint32 ;_ca uint32 ;_begc [109]uint32 ;};func (_beec *Reader )readAt (_fgbe int64 ,_afb int )([]byte ,error ){if _beec ._acb {_gff ,_fag :=_beec ._fbb .(slicer ).Slice (_fgbe ,_afb );
if _fag !=nil {return nil ,Error {ErrRead ,"\u0073\u006c\u0069\u0063er\u0020\u0072\u0065\u0061\u0064\u0020\u0065\u0072\u0072\u006f\u0072\u0020\u0028"+_fag .Error ()+"\u0029",_fgbe };};return _gff ,nil ;};if _afb > len (_beec ._fff ){return nil ,Error {ErrRead ,"\u0072\u0065ad\u0020\u006c\u0065n\u0067\u0074\u0068\u0020gre\u0061te\u0072\u0020\u0074\u0068\u0061\u006e\u0020re\u0061\u0064\u0020\u0062\u0075\u0066\u0066e\u0072",int64 (_afb )};
};if _ ,_bfc :=_beec ._fbb .ReadAt (_beec ._fff [:_afb ],_fgbe );_bfc !=nil {return nil ,Error {ErrRead ,_bfc .Error (),_fgbe };};return _beec ._fff [:_afb ],nil ;};const _ee int =64+4*4+16+4+8*2+4+8;func (_dad *File )changeSize (_dbfe int64 )error {if _dbfe ==_dad .Size {return nil ;
};var _faa *File ;for _ ,_cgc :=range _dad ._dea ._fbfb {if _cgc .Name ==_dad .Name {_faa =_cgc ;break ;};};if _faa ==nil {return _ae .Errorf ("\u004e\u006f\u0020\u0064\u0069\u0072e\u0063\u0074\u006f\u0072\u0079\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066o\u0072\u0020\u0061\u0020\u0066\u0069\u006ce\u003a\u0020\u0025\u0073",_dad .Name );
};_eg :=_b .NewBuffer ([]byte {});if _cbe :=_gb .Write (_eg ,_gb .LittleEndian ,_dbfe );_cbe !=nil {return _cbe ;};for _aea ,_cec :=range _eg .Bytes (){_faa ._edb [_aea ]=_cec ;};var _cd int64 ;var _dfg bool ;if _dad .Size < _ccf {_dfg =true ;_cd =int64 (_eab );
}else {_cd =int64 (_dad ._dea ._gagd );};_aga :=int ((_dad .Size -1)/_cd )+1;_beb :=int ((_dbfe -1)/_cd )+1;if _aga < _beb {_aba ,_cdb :=_dad .findLast (_dfg );if _cdb !=nil {return _cdb ;};_dff ,_cdb :=_dad ._dea .findNextFreeSector (_dfg );if _cdb !=nil {return _cdb ;
};for _dga :=_beb -_aga ;_dga > 0;_dga --{if _aag :=_dad ._dea .saveToFatLocs (_aba ,_dff ,_dfg );_aag !=nil {return _aag ;};if _dga > 1{_aba =_dff ;_dff ++;}else if _eee :=_dad ._dea .saveToFatLocs (_dff ,_gcb ,_dfg );_eee !=nil {return _eee ;};};}else if _aga > _beb {_gcf :=_dad ._ce ;
var _dffd error ;for _gef :=0;_gef < _beb -1;_gef ++{_gcf ,_dffd =_dad ._dea .findNext (_gcf ,_dfg );if _dffd !=nil {return _dffd ;};};if _bcca :=_dad ._dea .saveToFatLocs (_gcf ,_gcb ,_dfg );_bcca !=nil {return _bcca ;};};_dad .Size =_dbfe ;return nil ;
};func (_ecgg *Reader )GetEntry (entryName string )(*File ,error ){for _bedd ,_ddg :=_ecgg .Next ();_ddg ==nil ;_bedd ,_ddg =_ecgg .Next (){if _bedd .Name ==entryName {return _bedd ,nil ;};};return nil ,Error {ErrTraverse ,"\u004e\u006f\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066o\u0075\u006e\u0064\u0020\u0066\u006f\u0072 \u0067\u0069\u0076\u0065\u006e\u0020\u006e\u0061\u006d\u0065\u002e",0};
};func _ddd (_dgd *File ){if _dgd ._eb < 4||_dgd ._eb > 64{return ;};_ec :=int (_dgd ._eb /2-1);_dgd .Initial =_dgd ._dc [0];var _dcf int ;if !_f .IsPrint (rune (_dgd .Initial )){_dcf =1;};_dgd .Name =string (_db .Decode (_dgd ._dc [_dcf :_ec ]));};type directoryEntryFields struct{_dc [32]uint16 ;
_eb uint16 ;_fe uint8 ;_edd uint8 ;_fb uint32 ;_gg uint32 ;_gbe uint32 ;_ged _ed .Guid ;_bf [4]byte ;_bg _ed .FileTime ;_aa _ed .FileTime ;_ce uint32 ;_edb [8]byte ;};func (_bbd fileInfo )Name ()string {return _bbd .File .Name };func (_acea *Reader )ID ()string {return _acea .File [0].ID ()};
func (_fgc *Reader )traverse ()error {_fgc .File =make ([]*File ,0,len (_fgc ._fbfb ));var (_cga func (int ,[]string );_fcc error ;_ggf int ;);_cga =func (_gaa int ,_dfc []string ){_ggf ++;if _ggf > len (_fgc ._fbfb ){_fcc =Error {ErrTraverse ,"\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c\u0020\u0063o\u0075\u006e\u0074\u0065\u0072\u0020\u006f\u0076\u0065\u0072f\u006c\u006f\u0077",int64 (_gaa )};
return ;};if _gaa < 0||_gaa >=len (_fgc ._fbfb ){_fcc =Error {ErrTraverse ,"\u0069\u006c\u006ceg\u0061\u006c\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c\u0020\u0069\u006e\u0064\u0065\u0078",int64 (_gaa )};return ;};_aef :=_fgc ._fbfb [_gaa ];
if _aef ._fb !=_cbb {_cga (int (_aef ._fb ),_dfc );};_fgc .File =append (_fgc .File ,_aef );_aef .Path =_dfc ;if _aef ._gbe !=_cbb {if _gaa > 0{_cga (int (_aef ._gbe ),append (_dfc ,_aef .Name ));}else {_cga (int (_aef ._gbe ),_dfc );};};if _aef ._gg !=_cbb {_cga (int (_aef ._gg ),_dfc );
};return ;};_cga (0,[]string {});return _fcc ;};func (_dde *File )SetEntryContent (b []byte )error {if _dfae :=_dde .ensureWriterAt ();_dfae !=nil {return _dfae ;};_dde .reset ();if _dec :=_dde .changeSize (int64 (len (b )));_dec !=nil {return _dec ;};
_ ,_fa :=_dde .Write (b );return _fa ;};func (_afeg *Reader )exportMiniStream ()(*_dd .Writer ,*_dd .Writer ,error ){_bbdgb ,_ebce :=_dd .NewWriter (),_dd .NewWriter ();_agfe :=uint32 (0);for _ ,_dbab :=range _afeg .File {if _dbab .Size ==0{continue ;};
_dbab .reset ();_dbab ._ce =_agfe ;_dbc :=int (_dbab .Size )/int (_eab );if int (_dbab .Size )%int (_eab )!=0{_dbc ++;};for _abd :=1;_abd < _dbc ;_abd ++{_agfe ++;if _cdd :=_gb .Write (_bbdgb ,_gb .LittleEndian ,_agfe );_cdd !=nil {return nil ,nil ,_cdd ;
};};if _dgb :=_gb .Write (_bbdgb ,_gb .LittleEndian ,_gcb );_dgb !=nil {return nil ,nil ,_dgb ;};_agfe ++;if _ ,_gbc :=_c .Copy (_ebce ,_dbab );_gbc !=nil {return nil ,nil ,_gbc ;};if _aeffd :=_ebce .AlignLength (64);_aeffd !=nil {return nil ,nil ,_aeffd ;
};};if _ffa :=_bbdgb .FillWithByte (int (_afeg ._gagd )-_bbdgb .Len (),_gabc );_ffa !=nil {return nil ,nil ,_ffa ;};if _ccb :=_ebce .AlignLength (int (_afeg ._gagd ));_ccb !=nil {return nil ,nil ,_ccb ;};return _bbdgb ,_ebce ,nil ;};func (_eef *File )Modified ()_g .Time {return _eef ._aa .Time ()};
func (_gea fileInfo )ModTime ()_g .Time {return _gea .Modified ()};func (_gaac fileInfo )Size ()int64 {if _gaac ._fe !=_cg {return 0;};return _gaac .File .Size ;};func (_fdaa *Reader )Next ()(*File ,error ){_fdaa ._fbfd ++;if _fdaa ._fbfd >=len (_fdaa .File ){return nil ,_c .EOF ;
};return _fdaa .File [_fdaa ._fbfd ],nil ;};func _cee (_ggae []byte )*headerFields {_bbb :=&headerFields {};_bbb ._gbf =_gb .LittleEndian .Uint64 (_ggae [:8]);_bbb ._bcg =_gb .LittleEndian .Uint16 (_ggae [24:26]);_bbb ._afe =_gb .LittleEndian .Uint16 (_ggae [26:28]);
_bbb ._fbg =_gb .LittleEndian .Uint16 (_ggae [30:32]);_bbb ._gag =_gb .LittleEndian .Uint32 (_ggae [40:44]);_bbb ._bdba =_gb .LittleEndian .Uint32 (_ggae [44:48]);_bbb ._add =_gb .LittleEndian .Uint32 (_ggae [48:52]);_bbb ._cfb =_gb .LittleEndian .Uint32 (_ggae [60:64]);
_bbb ._fea =_gb .LittleEndian .Uint32 (_ggae [64:68]);_bbb ._cce =_gb .LittleEndian .Uint32 (_ggae [68:72]);_bbb ._ca =_gb .LittleEndian .Uint32 (_ggae [72:76]);var _cad int ;for _aecd :=76;_aecd < 512;_aecd =_aecd +4{_bbb ._begc [_cad ]=_gb .LittleEndian .Uint32 (_ggae [_aecd :_aecd +4]);
_cad ++;};return _bbb ;};func (_deff *File )ReadAt (p []byte ,off int64 )(_dgc int ,_ddf error ){_af ,_bd ,_dcb :=_deff ._ccc ,_deff ._bca ,_deff ._cf ;_ ,_ddf =_deff .Seek (off ,0);if _ddf ==nil {_dgc ,_ddf =_deff .Read (p );};_deff ._ccc ,_deff ._bca ,_deff ._cf =_af ,_bd ,_dcb ;
return _dgc ,_ddf ;};func (_ccg *File )seek (_bdb int64 )error {var _agfc bool ;var _dag int64 ;if _ccg .Size < _ccf {_agfc =true ;_dag =64;}else {_dag =int64 (_ccg ._dea ._gagd );};var _bgeb int64 ;var _gaab error ;if _ccg ._bca > 0{if _dag -_ccg ._bca <=_bdb {_ccg ._cf ,_gaab =_ccg ._dea .findNext (_ccg ._cf ,_agfc );
if _gaab !=nil {return _gaab ;};_bgeb +=_dag -_ccg ._bca ;_ccg ._bca =0;if _bgeb ==_bdb {return nil ;};}else {_ccg ._bca +=_bdb ;return nil ;};if _ccg ._cf ==_gcb {return Error {ErrRead ,"\u0075\u006ee\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0065\u0061\u0072\u006c\u0079\u0020\u0065\u006e\u0064\u0020\u006f\u0066\u0020\u0063ha\u0069\u006e",int64 (_ccg ._cf )};
};};for {if _bdb -_bgeb < _dag {_ccg ._bca =_bdb -_bgeb ;return nil ;}else {_bgeb +=_dag ;_ccg ._cf ,_gaab =_ccg ._dea .findNext (_ccg ._cf ,_agfc );if _gaab !=nil {return _gaab ;};if _bgeb ==_bdb {return nil ;};};};};func (_dcd *File )Write (b []byte )(int ,error ){if _dcd .Size < 1||_dcd ._ccc >=_dcd .Size {return 0,_c .EOF ;
};if _gca :=_dcd .ensureWriterAt ();_gca !=nil {return 0,_gca ;};_bbdg :=len (b );if int64 (_bbdg )> _dcd .Size -_dcd ._ccc {_bbdg =int (_dcd .Size -_dcd ._ccc );};_efd ,_geda :=_dcd .stream (_bbdg );if _geda !=nil {return 0,_geda ;};var _dddd ,_ceg int ;
for _ ,_gfc :=range _efd {_cbgg :=_dddd +int (_gfc [1]);if _cbgg < _dddd ||_cbgg > _bbdg {return 0,Error {ErrWrite ,"\u0062\u0061d\u0020\u0077\u0072i\u0074\u0065\u0020\u006c\u0065\u006e\u0067\u0074\u0068",int64 (_cbgg )};};_da ,_gfb :=_dcd ._dea ._cfcf .WriteAt (b [_dddd :_cbgg ],_gfc [0]);
_ceg =_ceg +_da ;if _gfb !=nil {_dcd ._ccc +=int64 (_ceg );return _ceg ,Error {ErrWrite ,"\u0075n\u0064\u0065\u0072\u006c\u0079\u0069\u006e\u0067\u0020\u0077\u0072i\u0074\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0020\u0028"+_gfb .Error ()+"\u0029",int64 (_dddd )};
};_dddd =_cbgg ;};_dcd ._ccc +=int64 (_ceg );if _ceg !=_bbdg {_geda =Error {ErrWrite ,"\u0062\u0079t\u0065\u0073\u0020\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0064\u006f\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0077\u0072\u0069\u0074\u0065\u0020\u0073\u0069\u007a\u0065",int64 (_ceg )};
}else if _ceg < len (b ){_geda =_c .EOF ;};return _ceg ,_geda ;};func (_acg *Reader )Debug ()map[string ][]uint32 {_bgfd :=map[string ][]uint32 {"s\u0065\u0063\u0074\u006f\u0072\u0020\u0073\u0069\u007a\u0065":[]uint32 {_acg ._gagd },"\u006d\u0069\u006e\u0069\u0020\u0066\u0061\u0074\u0020\u006c\u006f\u0063\u0073":_acg ._ffde ._cfe ,"\u006d\u0069n\u0069\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u006c\u006f\u0063\u0073":_acg ._ffde ._fgca ,"\u0064\u0069r\u0065\u0063\u0074o\u0072\u0079\u0020\u0073\u0065\u0063\u0074\u006f\u0072":[]uint32 {_acg ._ffde ._add },"\u006d\u0069\u006e\u0069 s\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0074\u0061\u0072\u0074\u002f\u0073\u0069z\u0065":[]uint32 {_acg .File [0]._ce ,_gb .LittleEndian .Uint32 (_acg .File [0]._edb [:])}};
for _bgc ,_ggea :=_acg .Next ();_ggea ==nil ;_bgc ,_ggea =_acg .Next (){_bgfd [_bgc .Name +" \u0073\u0074\u0061\u0072\u0074\u002f\u0073\u0069\u007a\u0065"]=[]uint32 {_bgc ._ce ,_gb .LittleEndian .Uint32 (_bgc ._edb [:])};};return _bgfd ;};func (_gede *Reader )getOffset (_aca uint32 ,_fde bool )(int64 ,error ){if _fde {_gfg :=_gede ._gagd /64;
_ecg :=int (_aca /_gfg );if _ecg >=len (_gede ._ffde ._fgca ){return 0,Error {ErrRead ,"\u006di\u006e\u0069s\u0065\u0063\u0074o\u0072\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0069\u0073\u0020\u006f\u0075t\u0073\u0069\u0064\u0065\u0020\u006d\u0069\u006e\u0069\u0073\u0065c\u0074\u006f\u0072\u0020\u0072\u0061\u006e\u0067\u0065",int64 (_ecg )};
};_dfb :=_aca %_gfg ;return int64 ((_gede ._ffde ._fgca [_ecg ]+1)*_gede ._gagd +_dfb *64),nil ;};return _cbf (_gede ._gagd ,_aca ),nil ;};func (_dbd *Reader )Read (b []byte )(_bda int ,_cda error ){if _dbd ._fbfd >=len (_dbd .File ){return 0,_c .EOF ;
};return _dbd .File [_dbd ._fbfd ].Read (b );};func (_fad *File )findLast (_ced bool )(uint32 ,error ){_ac :=_fad ._ce ;for {_aeff ,_dcg :=_fad ._dea .findNext (_ac ,_ced );if _dcg !=nil {return 0,Error {ErrRead ,"\u0062\u0061\u0064\u0020\u0072\u0065\u0061\u0064\u0020\u0066i\u006e\u0064\u0069\u006e\u0067\u0020\u006ee\u0078\u0074\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0028"+_dcg .Error ()+"\u0029",0};
};if _aeff ==_gcb {break ;};_ac =_aeff ;};return _ac ,nil ;};const (_egd uint32 =0xFFFFFFFA;_fdfg uint32 =0xFFFFFFFC;_cbaa uint32 =0xFFFFFFFD;_gcb uint32 =0xFFFFFFFE;_aafg uint32 =0xFFFFFFFF;_gabc byte =0xFF;_dda uint32 =0xFFFFFFFA;_cbb uint32 =0xFFFFFFFF;
);func (_gee *Reader )setDirEntries ()error {_dg :=20;if _gee ._ffde ._gag > 0{_dg =int (_gee ._ffde ._gag );};_bcb :=make ([]*File ,0,_dg );_ga :=make (map[uint32 ]bool );_gad :=int (_gee ._gagd /_eff );_cb :=_gee ._ffde ._add ;for _cb !=_gcb {_ab ,_ef :=_gee .readAt (_cbf (_gee ._gagd ,_cb ),int (_gee ._gagd ));
if _ef !=nil {return Error {ErrRead ,"\u0064\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u0020e\u006e\u0074\u0072\u0069\u0065\u0073\u0020r\u0065\u0061\u0064\u0020\u0065\u0072\u0072\u006f\u0072\u0020\u0028"+_ef .Error ()+"\u0029",_cbf (_gee ._gagd ,_cb )};
};for _fc :=0;_fc < _gad ;_fc ++{_eed :=&File {_dea :_gee };_eed .directoryEntryFields =_bb (_ab [_fc *int (_eff ):]);_edc (_gee ._ffde ._afe ,_eed );_eed ._cf =_eed ._ce ;_bcb =append (_bcb ,_eed );};_aec ,_ef :=_gee .findNext (_cb ,false );if _ef !=nil {return Error {ErrRead ,"d\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0065r\u0072\u006f\u0072\u0020\u0066\u0069\u006e\u0064\u0069\u006eg \u0073\u0065\u0063t\u006fr\u0020\u0028"+_ef .Error ()+"\u0029",int64 (_aec )};
};if _aec <=_cb {if _aec ==_cb ||_ga [_aec ]{return Error {ErrRead ,"\u0064\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u0020e\u006e\u0074\u0072\u0069\u0065\u0073\u0020s\u0065\u0063\u0074\u006f\u0072\u0020\u0063\u0079\u0063\u006c\u0065",int64 (_aec )};
};_ga [_aec ]=true ;};_cb =_aec ;};_gee ._fbfb =_bcb ;return nil ;};func (_bagc *Reader )Export ()([]byte ,error ){_bccbf :=_dd .NewWriter ();if _gdbc :=_bagc .exportHeader (_bccbf );_gdbc !=nil {return nil ,_gdbc ;};if _ebag :=_bccbf .FillWithByte (512,_gabc );
_ebag !=nil {return nil ,_ebag ;};_ece :=[]uint32 {};if _babd :=_bagc .exportDifats (_bccbf );_babd !=nil {return nil ,_babd ;};_fafe ,_gcef ,_ffdc :=_bagc .exportMiniStream ();if _ffdc !=nil {return nil ,_ffdc ;};_ece =append (_ece ,uint32 (_bccbf .Len ())/_bagc ._gagd );
if _gbee :=_bagc .exportDirEntries (_bccbf );_gbee !=nil {return nil ,_gbee ;};_ece =append (_ece ,uint32 (_bccbf .Len ())/_bagc ._gagd );if _ ,_bef :=_fafe .WriteTo (_bccbf );_bef !=nil {return nil ,_bef ;};_ece =append (_ece ,uint32 (_bccbf .Len ())/_bagc ._gagd );
if _ ,_eegg :=_gcef .WriteTo (_bccbf );_eegg !=nil {return nil ,_eegg ;};_ece =append (_ece ,uint32 (_bccbf .Len ())/_bagc ._gagd );if _agec :=_bagc .exportFAT (_bccbf ,_ece );_agec !=nil {return nil ,_agec ;};return _bccbf .Bytes (),nil ;};func (_ea *File )reset (){_ea ._ccc =0;
_ea ._bca =0;_ea ._cf =_ea ._ce };const (_dbf uint8 =0x0;_de uint8 =0x1;_cg uint8 =0x2;_cc uint8 =0x5;);func (_bgf fileInfo )Mode ()_e .FileMode {return _bgf .File .mode ()};func (_gce Error )Error ()string {return "\u006ds\u0063\u0066\u0062\u003a\u0020"+_gce ._bba +"\u003b\u0020"+_a .FormatInt (_gce ._edbd ,10);
};const (ErrFormat =iota ;ErrRead ;ErrSeek ;ErrWrite ;ErrTraverse ;);func (_dfd *Reader )setMiniStream ()error {if _dfd ._fbfb [0]._ce ==_gcb ||_dfd ._ffde ._cfb ==_gcb ||_dfd ._ffde ._fea ==0{return nil ;};_ebd :=int (_dfd ._ffde ._fea );_dfd ._ffde ._cfe =make ([]uint32 ,_ebd );
_dfd ._ffde ._cfe [0]=_dfd ._ffde ._cfb ;for _bfe :=1;_bfe < _ebd ;_bfe ++{_ebf ,_dba :=_dfd .findNext (_dfd ._ffde ._cfe [_bfe -1],false );if _dba !=nil {return Error {ErrFormat ,"s\u0065\u0074\u0074\u0069ng\u0020m\u0069\u006e\u0069\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u0028"+_dba .Error ()+"\u0029",int64 (_dfd ._ffde ._cfe [_bfe -1])};
};_dfd ._ffde ._cfe [_bfe ]=_ebf ;};_ebd =int (_dfd ._gagd /4*_dfd ._ffde ._fea );_dfd ._ffde ._fgca =make ([]uint32 ,0,_ebd );_ddb :=_dfd ._fbfb [0]._ce ;var _gabce error ;for _ddb !=_gcb {_dfd ._ffde ._fgca =append (_dfd ._ffde ._fgca ,_ddb );_ddb ,_gabce =_dfd .findNext (_ddb ,false );
if _gabce !=nil {return Error {ErrFormat ,"s\u0065\u0074\u0074\u0069ng\u0020m\u0069\u006e\u0069\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u0028"+_gabce .Error ()+"\u0029",int64 (_ddb )};};};return nil ;};func (_aee *File )ID ()string {return _aee ._ged .String ()};
func (_eba *File )Read (b []byte )(int ,error ){if _eba .Size < 1||_eba ._ccc >=_eba .Size {return 0,_c .EOF ;};_fce :=len (b );if int64 (_fce )> _eba .Size -_eba ._ccc {_fce =int (_eba .Size -_eba ._ccc );};_gc ,_agf :=_eba .stream (_fce );if _agf !=nil {return 0,_agf ;
};var _dfa ,_bbdc int ;for _ ,_cbg :=range _gc {_be :=_dfa +int (_cbg [1]);if _be < _dfa ||_be > _fce {return 0,Error {ErrRead ,"\u0062a\u0064 \u0072\u0065\u0061\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068",int64 (_be )};};_geb ,_bcc :=_eba ._dea ._fbb .ReadAt (b [_dfa :_be ],_cbg [0]);
_bbdc =_bbdc +_geb ;if _bcc !=nil {_eba ._ccc +=int64 (_bbdc );return _bbdc ,Error {ErrRead ,"\u0075n\u0064\u0065\u0072\u006c\u0079\u0069\u006e\u0067\u0020\u0072\u0065a\u0064\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0020\u0028"+_bcc .Error ()+"\u0029",int64 (_dfa )};
};_dfa =_be ;};_eba ._ccc +=int64 (_bbdc );if _bbdc !=_fce {_agf =Error {ErrRead ,"\u0062\u0079\u0074e\u0073\u0020\u0072\u0065\u0061\u0064\u0020\u0064\u006f\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020r\u0065\u0061\u0064\u0020\u0073\u0069\u007a\u0065",int64 (_bbdc )};
}else if _bbdc < len (b ){_agf =_c .EOF ;};return _bbdc ,_agf ;};func (_bdf *Reader )findNext (_fae uint32 ,_bcag bool )(uint32 ,error ){_dcca :=_bdf ._gagd /4;_aebd :=int (_fae /_dcca );var _adc uint32 ;if _bcag {if _aebd < 0||_aebd >=len (_bdf ._ffde ._cfe ){return 0,Error {ErrRead ,"\u006d\u0069\u006e\u0069\u0073e\u0063\u0074\u006f\u0072\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u0069\u0073 \u006f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u006d\u0069\u006e\u0069\u0046\u0041\u0054\u0020\u0072\u0061\u006e\u0067\u0065",int64 (_aebd )};
};_adc =_bdf ._ffde ._cfe [_aebd ];}else {if _aebd < 0||_aebd >=len (_bdf ._ffde ._gbd ){return 0,Error {ErrRead ,"\u0046\u0041\u0054\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u0069\u0073\u0020\u006f\u0075t\u0073i\u0064\u0065\u0020\u0044\u0049\u0046\u0041\u0054\u0020\u0072\u0061\u006e\u0067\u0065",int64 (_aebd )};
};_adc =_bdf ._ffde ._gbd [_aebd ];};_ace :=_fae %_dcca ;_efe :=_cbf (_bdf ._gagd ,_adc )+int64 (_ace *4);_dgce ,_baa :=_bdf .readAt (_efe ,4);if _baa !=nil {return 0,Error {ErrRead ,"\u0062\u0061\u0064\u0020\u0072\u0065\u0061\u0064\u0020\u0066i\u006e\u0064\u0069\u006e\u0067\u0020\u006ee\u0078\u0074\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0028"+_baa .Error ()+"\u0029",_efe };
};_cadb :=_gb .LittleEndian .Uint32 (_dgce );return _cadb ,nil ;};func (_dbe *File )Seek (offset int64 ,whence int )(int64 ,error ){var _cccg int64 ;switch whence {default:return 0,Error {ErrSeek ,"\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077h\u0065\u006e\u0063\u0065",int64 (whence )};
case 0:_cccg =offset ;case 1:_cccg =_dbe ._ccc +offset ;case 2:_cccg =_dbe .Size -offset ;};switch {case _cccg < 0:return _dbe ._ccc ,Error {ErrSeek ,"\u0063\u0061\u006e'\u0074\u0020\u0073\u0065e\u006b\u0020\u0062\u0065\u0066\u006f\u0072e\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u006f\u0066\u0020\u0046\u0069\u006c\u0065",_cccg };
case _cccg >=_dbe .Size :return _dbe ._ccc ,Error {ErrSeek ,"c\u0061\u006e\u0027\u0074\u0020\u0073e\u0065\u006b\u0020\u0070\u0061\u0073\u0074\u0020\u0046i\u006c\u0065\u0020l\u0065n\u0067\u0074\u0068",_cccg };case _cccg ==_dbe ._ccc :return _cccg ,nil ;
case _cccg > _dbe ._ccc :_dgg :=_dbe ._ccc ;_dbe ._ccc =_cccg ;return _dbe ._ccc ,_dbe .seek (_cccg -_dgg );};if _dbe ._bca >=_dbe ._ccc -_cccg {_dbe ._bca =_dbe ._bca -(_dbe ._ccc -_cccg );_dbe ._ccc =_cccg ;return _dbe ._ccc ,nil ;};_dbe ._bca =0;_dbe ._cf =_dbe ._ce ;
_dbe ._ccc =_cccg ;return _dbe ._ccc ,_dbe .seek (_cccg );};