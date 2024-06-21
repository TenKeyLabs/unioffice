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

package mscfb ;import (_a "bytes";_fa "encoding/binary";_f "fmt";_ga "github.com/richardlehane/msoleps/types";_gg "github.com/unidoc/unioffice/internal/mscfb/rw";_afd "io";_ef "os";_b "strconv";_af "time";_c "unicode";_g "unicode/utf16";);const (_d uint8 =0x0;
_bb uint8 =0x1;_efb uint8 =0x2;_efd uint8 =0x5;);func (_aga *File )FileInfo ()_ef .FileInfo {return fileInfo {_aga }};func (_ad *File )Read (b []byte )(int ,error ){if _ad .Size < 1||_ad ._cbc >=_ad .Size {return 0,_afd .EOF ;};_ede :=len (b );if int64 (_ede )> _ad .Size -_ad ._cbc {_ede =int (_ad .Size -_ad ._cbc );
};_adg ,_feg :=_ad .stream (_ede );if _feg !=nil {return 0,_feg ;};var _dad ,_dfd int ;for _ ,_aac :=range _adg {_bge :=_dad +int (_aac [1]);if _bge < _dad ||_bge > _ede {return 0,Error {ErrRead ,"\u0062a\u0064 \u0072\u0065\u0061\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068",int64 (_bge )};
};_gdf ,_fcde :=_ad ._dff ._aeba .ReadAt (b [_dad :_bge ],_aac [0]);_dfd =_dfd +_gdf ;if _fcde !=nil {_ad ._cbc +=int64 (_dfd );return _dfd ,Error {ErrRead ,"\u0075n\u0064\u0065\u0072\u006c\u0079\u0069\u006e\u0067\u0020\u0072\u0065a\u0064\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0020\u0028"+_fcde .Error ()+"\u0029",int64 (_dad )};
};_dad =_bge ;};_ad ._cbc +=int64 (_dfd );if _dfd !=_ede {_feg =Error {ErrRead ,"\u0062\u0079\u0074e\u0073\u0020\u0072\u0065\u0061\u0064\u0020\u0064\u006f\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020r\u0065\u0061\u0064\u0020\u0073\u0069\u007a\u0065",int64 (_dfd )};
}else if _dfd < len (b ){_feg =_afd .EOF ;};return _dfd ,_feg ;};const (_dcgg uint32 =0xFFFFFFFA;_cge uint32 =0xFFFFFFFC;_fdc uint32 =0xFFFFFFFD;_abf uint32 =0xFFFFFFFE;_edd uint32 =0xFFFFFFFF;_fbgf byte =0xFF;_deb uint32 =0xFFFFFFFA;_gaed uint32 =0xFFFFFFFF;
);func _aabc (_bae [][2]int64 )[][2]int64 {_dfa :=len (_bae );for _agaf ,_bac :=0,0;_agaf < _dfa &&_bac +1< len (_bae );_agaf ++{if _bae [_bac ][0]+_bae [_bac ][1]==_bae [_bac +1][0]{_bae [_bac ][1]=_bae [_bac ][1]+_bae [_bac +1][1];for _fcfd :=range _bae [_bac +1:len (_bae )-1]{_bae [_bac +1+_fcfd ]=_bae [_fcfd +_bac +2];
};_bae =_bae [:len (_bae )-1];}else {_bac +=1;};};return _bae ;};func (_dce *Reader )traverse ()error {_dce .File =make ([]*File ,0,len (_dce ._aebd ));var (_fc func (int ,[]string );_fg error ;_ec int ;);_fc =func (_ae int ,_bbe []string ){_ec ++;if _ec > len (_dce ._aebd ){_fg =Error {ErrTraverse ,"\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c\u0020\u0063o\u0075\u006e\u0074\u0065\u0072\u0020\u006f\u0076\u0065\u0072f\u006c\u006f\u0077",int64 (_ae )};
return ;};if _ae < 0||_ae >=len (_dce ._aebd ){_fg =Error {ErrTraverse ,"\u0069\u006c\u006ceg\u0061\u006c\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c\u0020\u0069\u006e\u0064\u0065\u0078",int64 (_ae )};return ;};_ffc :=_dce ._aebd [_ae ];
if _ffc ._fe !=_gaed {_fc (int (_ffc ._fe ),_bbe );};_dce .File =append (_dce .File ,_ffc );_ffc .Path =_bbe ;if _ffc ._bbg !=_gaed {if _ae > 0{_fc (int (_ffc ._bbg ),append (_bbe ,_ffc .Name ));}else {_fc (int (_ffc ._bbg ),_bbe );};};if _ffc ._de !=_gaed {_fc (int (_ffc ._de ),_bbe );
};return ;};_fc (0,[]string {});return _fg ;};func (_agec *Reader )getOffset (_cecg uint32 ,_bee bool )(int64 ,error ){if _bee {_ecaca :=_agec ._beaa /64;_cged :=int (_cecg /_ecaca );if _cged >=len (_agec ._cbf ._acefd ){return 0,Error {ErrRead ,"\u006di\u006e\u0069s\u0065\u0063\u0074o\u0072\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0069\u0073\u0020\u006f\u0075t\u0073\u0069\u0064\u0065\u0020\u006d\u0069\u006e\u0069\u0073\u0065c\u0074\u006f\u0072\u0020\u0072\u0061\u006e\u0067\u0065",int64 (_cged )};
};_aeb :=_cecg %_ecaca ;return int64 ((_agec ._cbf ._acefd [_cged ]+1)*_agec ._beaa +_aeb *64),nil ;};return _aeg (_agec ._beaa ,_cecg ),nil ;};func (_gc *File )Modified ()_af .Time {return _gc ._efe .Time ()};func (_dcg *File )WriteAt (p []byte ,off int64 )(_fbc int ,_ddc error ){_cda ,_daf ,_gga :=_dcg ._cbc ,_dcg ._ead ,_dcg ._fcd ;
_ ,_ddc =_dcg .Seek (off ,0);if _ddc ==nil {_fbc ,_ddc =_dcg .Write (p );};_dcg ._cbc ,_dcg ._ead ,_dcg ._fcd =_cda ,_daf ,_gga ;return _fbc ,_ddc ;};func (_gaf *Reader )GetEntry (entryName string )(*File ,error ){for _aba ,_fba :=_gaf .Next ();_fba ==nil ;
_aba ,_fba =_gaf .Next (){if _aba .Name ==entryName {return _aba ,nil ;};};return nil ,Error {ErrTraverse ,"\u004e\u006f\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066o\u0075\u006e\u0064\u0020\u0066\u006f\u0072 \u0067\u0069\u0076\u0065\u006e\u0020\u006e\u0061\u006d\u0065\u002e",0};
};func (_feab Error )Error ()string {return "\u006ds\u0063\u0066\u0062\u003a\u0020"+_feab ._ecab +"\u003b\u0020"+_b .FormatInt (_feab ._afbd ,10);};func (_gba Error )Typ ()int {return _gba ._daca };func (_cba fileInfo )ModTime ()_af .Time {return _cba .Modified ()};
func _dg (_fabd *File ){if _fabd ._da < 4||_fabd ._da > 64{return ;};_cb :=int (_fabd ._da /2-1);_fabd .Initial =_fabd ._ed [0];var _ag int ;if !_c .IsPrint (rune (_fabd .Initial )){_ag =1;};_fabd .Name =string (_g .Decode (_fabd ._ed [_ag :_cb ]));};func (_fag *File )mode ()_ef .FileMode {if _fag ._ac !=_efb {return _ef .ModeDir |0777;
};return 0666;};type Error struct{_daca int ;_ecab string ;_afbd int64 ;};func (_gbe fileInfo )Sys ()interface{}{return nil };type directoryEntryFields struct{_ed [32]uint16 ;_da uint16 ;_ac uint8 ;_aa uint8 ;_fe uint32 ;_de uint32 ;_bbg uint32 ;_bd _ga .Guid ;
_ea [4]byte ;_eg _ga .FileTime ;_efe _ga .FileTime ;_dd uint32 ;_fb [8]byte ;};func (_eadg *Reader )exportMiniStream ()(*_gg .Writer ,*_gg .Writer ,error ){_bedb ,_ddb :=_gg .NewWriter (),_gg .NewWriter ();_cbec :=uint32 (0);for _ ,_bbbg :=range _eadg .File {if _bbbg .Size ==0{continue ;
};_bbbg .reset ();_bbbg ._dd =_cbec ;_dcge :=int (_bbbg .Size )/int (_ageg );if int (_bbbg .Size )%int (_ageg )!=0{_dcge ++;};for _fbab :=1;_fbab < _dcge ;_fbab ++{_cbec ++;if _eae :=_fa .Write (_bedb ,_fa .LittleEndian ,_cbec );_eae !=nil {return nil ,nil ,_eae ;
};};if _cbd :=_fa .Write (_bedb ,_fa .LittleEndian ,_abf );_cbd !=nil {return nil ,nil ,_cbd ;};_cbec ++;if _ ,_dag :=_afd .Copy (_ddb ,_bbbg );_dag !=nil {return nil ,nil ,_dag ;};if _fecd :=_ddb .AlignLength (64);_fecd !=nil {return nil ,nil ,_fecd ;
};};if _cef :=_bedb .FillWithByte (int (_eadg ._beaa )-_bedb .Len (),_fbgf );_cef !=nil {return nil ,nil ,_cef ;};if _bagcb :=_ddb .AlignLength (int (_eadg ._beaa ));_bagcb !=nil {return nil ,nil ,_bagcb ;};return _bedb ,_ddb ,nil ;};func (_cca *Reader )setDifats ()error {_cca ._cbf ._cfd =_cca ._cbf ._bfc [:];
if _cca ._cbf ._dedd ==0{return nil ;};_bba :=(_cca ._beaa /4)-1;_gbed :=make ([]uint32 ,109,_cca ._cbf ._dedd *_bba +109);copy (_gbed ,_cca ._cbf ._cfd );_cca ._cbf ._cfd =_gbed ;_dgeb :=_cca ._cbf ._bag ;for _fdg :=0;_fdg < int (_cca ._cbf ._dedd );_fdg ++{_fgg ,_cgg :=_cca .readAt (_aeg (_cca ._beaa ,_dgeb ),int (_cca ._beaa ));
if _cgg !=nil {return Error {ErrFormat ,"e\u0072r\u006f\u0072\u0020\u0073\u0065\u0074\u0074\u0069n\u0067\u0020\u0044\u0049FA\u0054\u0028"+_cgg .Error ()+"\u0029",int64 (_dgeb )};};for _debc :=0;_debc < int (_bba );_debc ++{_cca ._cbf ._cfd =append (_cca ._cbf ._cfd ,_fa .LittleEndian .Uint32 (_fgg [_debc *4:_debc *4+4]));
};_dgeb =_fa .LittleEndian .Uint32 (_fgg [len (_fgg )-4:]);};return nil ;};func (_eef *File )ensureWriterAt ()error {if _eef ._dff ._ega ==nil {_eee ,_bgec :=_eef ._dff ._aeba .(_afd .WriterAt );if !_bgec {return Error {ErrWrite ,"\u006d\u0073\u0063\u0066\u0062\u002e\u004ee\u0077\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0067\u0069\u0076\u0065n\u0020R\u0065\u0061\u0064\u0065\u0072\u0041t\u0020\u0063\u006f\u006e\u0076\u0065\u0072t\u0069\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0020\u0069\u006f\u002e\u0057\u0072\u0069\u0074\u0065\u0072\u0041\u0074\u0020\u0069n\u0020\u006f\u0072\u0064\u0065\u0072\u0020\u0074\u006f\u0020\u0077\u0072\u0069t\u0065",0};
};_eef ._dff ._ega =_eee ;};return nil ;};func (_dea *Reader )setHeader ()error {_dadd ,_dacb :=_dea .readAt (0,_fdae );if _dacb !=nil {return _dacb ;};_dea ._cbf =&header {headerFields :_faa (_dadd )};if _dea ._cbf ._ecac !=_edea {return Error {ErrFormat ,"\u0062\u0061\u0064\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065",int64 (_dea ._cbf ._ecac )};
};if _dea ._cbf ._bcc ==0x0009||_dea ._cbf ._bcc ==0x000c{_dea ._beaa =uint32 (1<<_dea ._cbf ._bcc );}else {return Error {ErrFormat ,"\u0069\u006c\u006c\u0065ga\u006c\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0073\u0069\u007a\u0065",int64 (_dea ._cbf ._bcc )};
};if _dea ._cbf ._dedd > 0{_gecb :=(_dea ._beaa /4)-1;if int (_dea ._cbf ._dedd *_gecb +109)< 0{return Error {ErrFormat ,"\u0044I\u0046A\u0054\u0020\u0069\u006e\u0074 \u006f\u0076e\u0072\u0066\u006c\u006f\u0077",int64 (_dea ._cbf ._dedd )};};if _dea ._cbf ._dedd *_gecb +109> _dea ._cbf ._bgf +_gecb {return Error {ErrFormat ,"\u006e\u0075\u006d\u0020\u0044\u0049\u0046\u0041\u0054\u0073 \u0065\u0078\u0063\u0065\u0065\u0064\u0073 \u0046\u0041\u0054\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0073",int64 (_dea ._cbf ._dedd )};
};};if _dea ._cbf ._fgd > 0{if int (_dea ._beaa /4*_dea ._cbf ._fgd )< 0{return Error {ErrFormat ,"m\u0069\u006e\u0069\u0020FA\u0054 \u0069\u006e\u0074\u0020\u006fv\u0065\u0072\u0066\u006c\u006f\u0077",int64 (_dea ._cbf ._fgd )};};if _dea ._cbf ._fgd > _dea ._cbf ._bgf *(_dea ._beaa /_ageg ){return Error {ErrFormat ,"\u006e\u0075\u006d\u0020\u006d\u0069n\u0069\u0020\u0046\u0041\u0054\u0073\u0020\u0065\u0078\u0063\u0065\u0065\u0064s\u0020\u0046\u0041\u0054\u0020\u0073\u0065c\u0074\u006f\u0072\u0073",int64 (_dea ._cbf ._bgf )};
};};return nil ;};func _aeg (_ecad ,_dcf uint32 )int64 {return int64 ((_dcf +1)*_ecad )};func (_fae *Reader )setDirEntries ()error {_gdc :=20;if _fae ._cbf ._cbe > 0{_gdc =int (_fae ._cbf ._cbe );};_dc :=make ([]*File ,0,_gdc );_dab :=make (map[uint32 ]bool );
_dae :=int (_fae ._beaa /_dde );_gb :=_fae ._cbf ._befg ;for _gb !=_abf {_gf ,_ff :=_fae .readAt (_aeg (_fae ._beaa ,_gb ),int (_fae ._beaa ));if _ff !=nil {return Error {ErrRead ,"\u0064\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u0020e\u006e\u0074\u0072\u0069\u0065\u0073\u0020r\u0065\u0061\u0064\u0020\u0065\u0072\u0072\u006f\u0072\u0020\u0028"+_ff .Error ()+"\u0029",_aeg (_fae ._beaa ,_gb )};
};for _bg :=0;_bg < _dae ;_bg ++{_ge :=&File {_dff :_fae };_ge .directoryEntryFields =_daa (_gf [_bg *int (_dde ):]);_gfa (_fae ._cbf ._dcb ,_ge );_ge ._fcd =_ge ._dd ;_dc =append (_dc ,_ge );};_eea ,_ff :=_fae .findNext (_gb ,false );if _ff !=nil {return Error {ErrRead ,"d\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0065r\u0072\u006f\u0072\u0020\u0066\u0069\u006e\u0064\u0069\u006eg \u0073\u0065\u0063t\u006fr\u0020\u0028"+_ff .Error ()+"\u0029",int64 (_eea )};
};if _eea <=_gb {if _eea ==_gb ||_dab [_eea ]{return Error {ErrRead ,"\u0064\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u0020e\u006e\u0074\u0072\u0069\u0065\u0073\u0020s\u0065\u0063\u0074\u006f\u0072\u0020\u0063\u0079\u0063\u006c\u0065",int64 (_eea )};
};_dab [_eea ]=true ;};_gb =_eea ;};_fae ._aebd =_dc ;return nil ;};type slicer interface{Slice (_afde int64 ,_ged int )([]byte ,error );};func (_gcab *Reader )exportHeader (_dbbf *_gg .Writer )error {if _dcaa :=_fa .Write (_dbbf ,_fa .LittleEndian ,&_gcab ._cbf ._ecac );
_dcaa !=nil {return _dcaa ;};if _feaf :=_dbbf .Skip (16);_feaf !=nil {return _feaf ;};if _bedg :=_fa .Write (_dbbf ,_fa .LittleEndian ,&_gcab ._cbf ._bad );_bedg !=nil {return _bedg ;};if _gac :=_fa .Write (_dbbf ,_fa .LittleEndian ,&_gcab ._cbf ._dcb );
_gac !=nil {return _gac ;};if _bcgg :=_fa .Write (_dbbf ,_fa .LittleEndian ,uint16 (0xfffe));_bcgg !=nil {return _bcgg ;};if _afdb :=_fa .Write (_dbbf ,_fa .LittleEndian ,&_gcab ._cbf ._bcc );_afdb !=nil {return _afdb ;};if _dfg :=_fa .Write (_dbbf ,_fa .LittleEndian ,uint16 (0x0006));
_dfg !=nil {return _dfg ;};if _bgae :=_dbbf .Skip (6);_bgae !=nil {return _bgae ;};if _bgfb :=_fa .Write (_dbbf ,_fa .LittleEndian ,&_gcab ._cbf ._cbe );_bgfb !=nil {return _bgfb ;};if _egab :=_fa .Write (_dbbf ,_fa .LittleEndian ,&_gcab ._cbf ._bgf );
_egab !=nil {return _egab ;};if _adc :=_fa .Write (_dbbf ,_fa .LittleEndian ,&_gcab ._cbf ._befg );_adc !=nil {return _adc ;};if _cdgb :=_dbbf .Skip (4);_cdgb !=nil {return _cdgb ;};if _dfe :=_fa .Write (_dbbf ,_fa .LittleEndian ,uint32 (0x00001000));_dfe !=nil {return _dfe ;
};if _gdfa :=_fa .Write (_dbbf ,_fa .LittleEndian ,&_gcab ._cbf ._egff );_gdfa !=nil {return _gdfa ;};if _efde :=_fa .Write (_dbbf ,_fa .LittleEndian ,&_gcab ._cbf ._fgd );_efde !=nil {return _efde ;};if _bace :=_fa .Write (_dbbf ,_fa .LittleEndian ,&_gcab ._cbf ._bag );
_bace !=nil {return _bace ;};if _caa :=_fa .Write (_dbbf ,_fa .LittleEndian ,&_gcab ._cbf ._dedd );_caa !=nil {return _caa ;};for _gcea :=0;_gcea < 109;_gcea ++{if _bagc :=_fa .Write (_dbbf ,_fa .LittleEndian ,&_gcab ._cbf ._bfc [_gcea ]);_bagc !=nil {return _bagc ;
};};return nil ;};func (_aacd *File )Seek (offset int64 ,whence int )(int64 ,error ){var _ffda int64 ;switch whence {default:return 0,Error {ErrSeek ,"\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077h\u0065\u006e\u0063\u0065",int64 (whence )};case 0:_ffda =offset ;
case 1:_ffda =_aacd ._cbc +offset ;case 2:_ffda =_aacd .Size -offset ;};switch {case _ffda < 0:return _aacd ._cbc ,Error {ErrSeek ,"\u0063\u0061\u006e'\u0074\u0020\u0073\u0065e\u006b\u0020\u0062\u0065\u0066\u006f\u0072e\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u006f\u0066\u0020\u0046\u0069\u006c\u0065",_ffda };
case _ffda >=_aacd .Size :return _aacd ._cbc ,Error {ErrSeek ,"c\u0061\u006e\u0027\u0074\u0020\u0073e\u0065\u006b\u0020\u0070\u0061\u0073\u0074\u0020\u0046i\u006c\u0065\u0020l\u0065n\u0067\u0074\u0068",_ffda };case _ffda ==_aacd ._cbc :return _ffda ,nil ;
case _ffda > _aacd ._cbc :_dda :=_aacd ._cbc ;_aacd ._cbc =_ffda ;return _aacd ._cbc ,_aacd .seek (_ffda -_dda );};if _aacd ._ead >=_aacd ._cbc -_ffda {_aacd ._ead =_aacd ._ead -(_aacd ._cbc -_ffda );_aacd ._cbc =_ffda ;return _aacd ._cbc ,nil ;};_aacd ._ead =0;
_aacd ._fcd =_aacd ._dd ;_aacd ._cbc =_ffda ;return _aacd ._cbc ,_aacd .seek (_ffda );};func (_bga *File )ReadAt (p []byte ,off int64 )(_ddg int ,_add error ){_bbf ,_ddgf ,_aaa :=_bga ._cbc ,_bga ._ead ,_bga ._fcd ;_ ,_add =_bga .Seek (off ,0);if _add ==nil {_ddg ,_add =_bga .Read (p );
};_bga ._cbc ,_bga ._ead ,_bga ._fcd =_bbf ,_ddgf ,_aaa ;return _ddg ,_add ;};func (_dfb *Reader )findFatLocsOffset (_ggea bool )int64 {var _fgca uint32 ;if _ggea {_fgca =_dfb ._cbf ._gcd [0];}else {_fgca =_dfb ._cbf ._cfd [0];};return _aeg (_dfb ._beaa ,_fgca );
};func (_ecb *File )Created ()_af .Time {return _ecb ._eg .Time ()};const _cd int =64+4*4+16+4+8*2+4+8;func (_eead *Reader )exportDirEntries (_afb *_gg .Writer )error {if int64 (_afb .Len ())!=_aeg (_eead ._beaa ,_eead ._cbf ._befg ){return Error {ErrWrite ,_f .Sprintf ("I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0077\u0072\u0069\u0074\u0065\u0072\u0020l\u0065\u006e\u0067t\u0068:\u0020\u0025\u0076",_afb .Len ()),0};
};for _ ,_fgc :=range _eead ._aebd {_dgc ,_edfa :=_eaab (_fgc .directoryEntryFields );if _edfa !=nil {return _edfa ;};if _ ,_ddce :=_afd .Copy (_afb ,_dgc );_ddce !=nil {return _ddce ;};};return nil ;};func (_dbb *Reader )GetHeader ()*header {return _dbb ._cbf };
type Reader struct{_ggd bool ;_beaa uint32 ;_acf []byte ;_cbf *header ;File []*File ;_aebd []*File ;_fea int ;_aeba _afd .ReaderAt ;_ega _afd .WriterAt ;};func (_ebc *Reader )findNextFreeSector (_bab bool )(uint32 ,error ){_fbb :=_ebc .findFatLocsOffset (_bab );
_bcg :=uint32 (0);_cfdf :=_ebc ._beaa /4;for {_gad ,_ggf :=_ebc .readAt (_fbb ,4);if _ggf !=nil {return 0,Error {ErrRead ,"\u0062\u0061\u0064\u0020\u0072\u0065\u0061\u0064\u0020\u0066i\u006e\u0064\u0069\u006e\u0067\u0020\u006ee\u0078\u0074\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0028"+_ggf .Error ()+"\u0029",_fbb };
};_cdf :=_fa .LittleEndian .Uint32 (_gad );if _cdf ==_edd {break ;};if _bcg >=_cfdf {return 0,Error {ErrRead ,"\u0065\u006e\u0064\u0020of\u0020\u006d\u0069\u006e\u0069\u0046\u0061\u0074\u0020\u0072\u0065\u0061\u0063\u0068e\u0064",_fbb };};_bcg ++;_fbb +=4;
};return _bcg ,nil ;};func (_gec *File )reset (){_gec ._cbc =0;_gec ._ead =0;_gec ._fcd =_gec ._dd };type header struct{*headerFields ;_cfd []uint32 ;_gcd []uint32 ;_acefd []uint32 ;};const (_edea uint64 =0xE11AB1A1E011CFD0;_ageg uint32 =64;_gcb int64 =4096;
_dde uint32 =128;);func (_dbdf *Reader )saveToFatLocs (_fcdc uint32 ,_gcg interface{},_ddbg bool )error {_febf :=_a .NewBuffer ([]byte {});if _gbg :=_fa .Write (_febf ,_fa .LittleEndian ,_gcg );_gbg !=nil {return _gbg ;};_cde :=_dbdf .findFatLocsOffset (_ddbg )+int64 (_fcdc *4);
_ ,_cdb :=_dbdf ._ega .WriteAt (_febf .Bytes (),_cde );return _cdb ;};func (_cdg *File )seek (_agef int64 )error {var _deg bool ;var _efdd int64 ;if _cdg .Size < _gcb {_deg =true ;_efdd =64;}else {_efdd =int64 (_cdg ._dff ._beaa );};var _acef int64 ;var _ceg error ;
if _cdg ._ead > 0{if _efdd -_cdg ._ead <=_agef {_cdg ._fcd ,_ceg =_cdg ._dff .findNext (_cdg ._fcd ,_deg );if _ceg !=nil {return _ceg ;};_acef +=_efdd -_cdg ._ead ;_cdg ._ead =0;if _acef ==_agef {return nil ;};}else {_cdg ._ead +=_agef ;return nil ;};if _cdg ._fcd ==_abf {return Error {ErrRead ,"\u0075\u006ee\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0065\u0061\u0072\u006c\u0079\u0020\u0065\u006e\u0064\u0020\u006f\u0066\u0020\u0063ha\u0069\u006e",int64 (_cdg ._fcd )};
};};for {if _agef -_acef < _efdd {_cdg ._ead =_agef -_acef ;return nil ;}else {_acef +=_efdd ;_cdg ._fcd ,_ceg =_cdg ._dff .findNext (_cdg ._fcd ,_deg );if _ceg !=nil {return _ceg ;};if _acef ==_agef {return nil ;};};};};func (_bda *File )stream (_dac int )([][2]int64 ,error ){var _aee bool ;
var _aca int ;var _edg int64 ;if _bda .Size < _gcb {_aee =true ;_aca =_dac /int (_ageg )+2;_edg =int64 (_ageg );}else {_aca =_dac /int (_bda ._dff ._beaa )+2;_edg =int64 (_bda ._dff ._beaa );};_ceb :=make ([][2]int64 ,0,_aca );var _bbbc ,_adf int ;if _bda ._ead > 0{_dec ,_fed :=_bda ._dff .getOffset (_bda ._fcd ,_aee );
if _fed !=nil {return nil ,_fed ;};if _edg -_bda ._ead >=int64 (_dac ){_ceb =append (_ceb ,[2]int64 {_dec +_bda ._ead ,int64 (_dac )});}else {_ceb =append (_ceb ,[2]int64 {_dec +_bda ._ead ,_edg -_bda ._ead });};if _edg -_bda ._ead <=int64 (_dac ){_bda ._fcd ,_fed =_bda ._dff .findNext (_bda ._fcd ,_aee );
if _fed !=nil {return nil ,_fed ;};_adf +=int (_edg -_bda ._ead );_bda ._ead =0;}else {_bda ._ead +=int64 (_dac );};if _ceb [0][1]==int64 (_dac ){return _ceb ,nil ;};if _bda ._fcd ==_abf {return nil ,Error {ErrRead ,"\u0075\u006ee\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0065\u0061\u0072\u006c\u0079\u0020\u0065\u006e\u0064\u0020\u006f\u0066\u0020\u0063ha\u0069\u006e",int64 (_bda ._fcd )};
};_bbbc ++;};for {if _bbbc >=cap (_ceb ){return nil ,Error {ErrRead ,"\u0069\u006e\u0064\u0065x\u0020\u006f\u0076\u0065\u0072\u0072\u0075\u006e\u0073\u0020s\u0065c\u0074\u006f\u0072\u0020\u006c\u0065\u006eg\u0074\u0068",int64 (_bbbc )};};_ddga ,_eaa :=_bda ._dff .getOffset (_bda ._fcd ,_aee );
if _eaa !=nil {return nil ,_eaa ;};if _dac -_adf < int (_edg ){_ceb =append (_ceb ,[2]int64 {_ddga ,int64 (_dac -_adf )});_bda ._ead =int64 (_dac -_adf );return _aabc (_ceb ),nil ;}else {_ceb =append (_ceb ,[2]int64 {_ddga ,_edg });_adf +=int (_edg );_bda ._fcd ,_eaa =_bda ._dff .findNext (_bda ._fcd ,_aee );
if _eaa !=nil {return nil ,_eaa ;};if _adf ==_dac {return _aabc (_ceb ),nil ;};};_bbbc ++;};};func (_afg *Reader )setMiniStream ()error {if _afg ._aebd [0]._dd ==_abf ||_afg ._cbf ._egff ==_abf ||_afg ._cbf ._fgd ==0{return nil ;};_aaaa :=int (_afg ._cbf ._fgd );
_afg ._cbf ._gcd =make ([]uint32 ,_aaaa );_afg ._cbf ._gcd [0]=_afg ._cbf ._egff ;for _dbd :=1;_dbd < _aaaa ;_dbd ++{_efba ,_fcb :=_afg .findNext (_afg ._cbf ._gcd [_dbd -1],false );if _fcb !=nil {return Error {ErrFormat ,"s\u0065\u0074\u0074\u0069ng\u0020m\u0069\u006e\u0069\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u0028"+_fcb .Error ()+"\u0029",int64 (_afg ._cbf ._gcd [_dbd -1])};
};_afg ._cbf ._gcd [_dbd ]=_efba ;};_aaaa =int (_afg ._beaa /4*_afg ._cbf ._fgd );_afg ._cbf ._acefd =make ([]uint32 ,0,_aaaa );_eeef :=_afg ._aebd [0]._dd ;var _abgg error ;for _eeef !=_abf {_afg ._cbf ._acefd =append (_afg ._cbf ._acefd ,_eeef );_eeef ,_abgg =_afg .findNext (_eeef ,false );
if _abgg !=nil {return Error {ErrFormat ,"s\u0065\u0074\u0074\u0069ng\u0020m\u0069\u006e\u0069\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u0028"+_abgg .Error ()+"\u0029",int64 (_eeef )};};};return nil ;};func (_cc *File )ID ()string {return _cc ._bd .String ()};
func (_cg fileInfo )Name ()string {return _cg .File .Name };func (_eaf fileInfo )IsDir ()bool {return _eaf .mode ().IsDir ()};const _fdae int =8+16+10+6+12+8+16+109*4;type headerFields struct{_ecac uint64 ;_ [16]byte ;_bad uint16 ;_dcb uint16 ;_ [2]byte ;
_bcc uint16 ;_ [2]byte ;_ [6]byte ;_cbe uint32 ;_bgf uint32 ;_befg uint32 ;_ [4]byte ;_ [4]byte ;_egff uint32 ;_fgd uint32 ;_bag uint32 ;_dedd uint32 ;_bfc [109]uint32 ;};func New (ra _afd .ReaderAt )(*Reader ,error ){_edc :=&Reader {_aeba :ra };if _ ,_dbc :=ra .(slicer );
_dbc {_edc ._ggd =true ;}else {_edc ._acf =make ([]byte ,_fdae );};if _cdc :=_edc .setHeader ();_cdc !=nil {return nil ,_cdc ;};if !_edc ._ggd &&int (_edc ._beaa )> len (_edc ._acf ){_edc ._acf =make ([]byte ,_edc ._beaa );};if _ecg :=_edc .setDifats ();
_ecg !=nil {return nil ,_ecg ;};if _daad :=_edc .setDirEntries ();_daad !=nil {return nil ,_daad ;};if _eba :=_edc .setMiniStream ();_eba !=nil {return nil ,_eba ;};if _ebb :=_edc .traverse ();_ebb !=nil {return nil ,_ebb ;};return _edc ,nil ;};func _gfa (_ffa uint16 ,_fab *File ){_dg (_fab );
if _fab ._ac !=_efb {return ;};if _ffa > 3{_fab .Size =int64 (_fa .LittleEndian .Uint64 (_fab ._fb [:]));}else {_fab .Size =int64 (_fa .LittleEndian .Uint32 (_fab ._fb [:4]));};};func (_gde *Reader )exportDifats (_feb *_gg .Writer )error {if _gde ._cbf ._dedd ==0{return nil ;
};return nil ;};func (_beac *Reader )readAt (_eec int64 ,_fabc int )([]byte ,error ){if _beac ._ggd {_aea ,_dafd :=_beac ._aeba .(slicer ).Slice (_eec ,_fabc );if _dafd !=nil {return nil ,Error {ErrRead ,"\u0073\u006c\u0069\u0063er\u0020\u0072\u0065\u0061\u0064\u0020\u0065\u0072\u0072\u006f\u0072\u0020\u0028"+_dafd .Error ()+"\u0029",_eec };
};return _aea ,nil ;};if _fabc > len (_beac ._acf ){return nil ,Error {ErrRead ,"\u0072\u0065ad\u0020\u006c\u0065n\u0067\u0074\u0068\u0020gre\u0061te\u0072\u0020\u0074\u0068\u0061\u006e\u0020re\u0061\u0064\u0020\u0062\u0075\u0066\u0066e\u0072",int64 (_fabc )};
};if _ ,_dfdg :=_beac ._aeba .ReadAt (_beac ._acf [:_fabc ],_eec );_dfdg !=nil {return nil ,Error {ErrRead ,_dfdg .Error (),_eec };};return _beac ._acf [:_fabc ],nil ;};func (_dace *Reader )Modified ()_af .Time {return _dace .File [0].Modified ()};type File struct{Name string ;
Initial uint16 ;Path []string ;Size int64 ;_cbc int64 ;_fcd uint32 ;_ead int64 ;*directoryEntryFields ;_dff *Reader ;};func (_egc *Reader )ID ()string {return _egc .File [0].ID ()};func (_ab fileInfo )Size ()int64 {if _ab ._ac !=_efb {return 0;};return _ab .File .Size ;
};func (_cgd *Reader )Export ()([]byte ,error ){_ddab :=_gg .NewWriter ();if _ebe :=_cgd .exportHeader (_ddab );_ebe !=nil {return nil ,_ebe ;};if _bggd :=_ddab .FillWithByte (512,_fbgf );_bggd !=nil {return nil ,_bggd ;};_cecc :=[]uint32 {};if _gda :=_cgd .exportDifats (_ddab );
_gda !=nil {return nil ,_gda ;};_efc ,_ccaa ,_dfc :=_cgd .exportMiniStream ();if _dfc !=nil {return nil ,_dfc ;};_cecc =append (_cecc ,uint32 (_ddab .Len ())/_cgd ._beaa );if _edba :=_cgd .exportDirEntries (_ddab );_edba !=nil {return nil ,_edba ;};_cecc =append (_cecc ,uint32 (_ddab .Len ())/_cgd ._beaa );
if _ ,_dgg :=_efc .WriteTo (_ddab );_dgg !=nil {return nil ,_dgg ;};_cecc =append (_cecc ,uint32 (_ddab .Len ())/_cgd ._beaa );if _ ,_dade :=_ccaa .WriteTo (_ddab );_dade !=nil {return nil ,_dade ;};_cecc =append (_cecc ,uint32 (_ddab .Len ())/_cgd ._beaa );
if _gadg :=_cgd .exportFAT (_ddab ,_cecc );_gadg !=nil {return nil ,_gadg ;};return _ddab .Bytes (),nil ;};func (_geg *Reader )findNext (_baf uint32 ,_ffac bool )(uint32 ,error ){_fcae :=_geg ._beaa /4;_abd :=int (_baf /_fcae );var _bgb uint32 ;if _ffac {if _abd < 0||_abd >=len (_geg ._cbf ._gcd ){return 0,Error {ErrRead ,"\u006d\u0069\u006e\u0069\u0073e\u0063\u0074\u006f\u0072\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u0069\u0073 \u006f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u006d\u0069\u006e\u0069\u0046\u0041\u0054\u0020\u0072\u0061\u006e\u0067\u0065",int64 (_abd )};
};_bgb =_geg ._cbf ._gcd [_abd ];}else {if _abd < 0||_abd >=len (_geg ._cbf ._cfd ){return 0,Error {ErrRead ,"\u0046\u0041\u0054\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u0069\u0073\u0020\u006f\u0075t\u0073i\u0064\u0065\u0020\u0044\u0049\u0046\u0041\u0054\u0020\u0072\u0061\u006e\u0067\u0065",int64 (_abd )};
};_bgb =_geg ._cbf ._cfd [_abd ];};_cedd :=_baf %_fcae ;_acea :=_aeg (_geg ._beaa ,_bgb )+int64 (_cedd *4);_eac ,_gge :=_geg .readAt (_acea ,4);if _gge !=nil {return 0,Error {ErrRead ,"\u0062\u0061\u0064\u0020\u0072\u0065\u0061\u0064\u0020\u0066i\u006e\u0064\u0069\u006e\u0067\u0020\u006ee\u0078\u0074\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0028"+_gge .Error ()+"\u0029",_acea };
};_aag :=_fa .LittleEndian .Uint32 (_eac );return _aag ,nil ;};func (_bed *File )findLast (_ffb bool )(uint32 ,error ){_ba :=_bed ._dd ;for {_bbde ,_gfe :=_bed ._dff .findNext (_ba ,_ffb );if _gfe !=nil {return 0,Error {ErrRead ,"\u0062\u0061\u0064\u0020\u0072\u0065\u0061\u0064\u0020\u0066i\u006e\u0064\u0069\u006e\u0067\u0020\u006ee\u0078\u0074\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0028"+_gfe .Error ()+"\u0029",0};
};if _bbde ==_abf {break ;};_ba =_bbde ;};return _ba ,nil ;};func (_ddcd *Reader )Created ()_af .Time {return _ddcd .File [0].Created ()};func (_dffd *File )Write (b []byte )(int ,error ){if _dffd .Size < 1||_dffd ._cbc >=_dffd .Size {return 0,_afd .EOF ;
};if _ca :=_dffd .ensureWriterAt ();_ca !=nil {return 0,_ca ;};_bef :=len (b );if int64 (_bef )> _dffd .Size -_dffd ._cbc {_bef =int (_dffd .Size -_dffd ._cbc );};_gfg ,_ece :=_dffd .stream (_bef );if _ece !=nil {return 0,_ece ;};var _eff ,_edf int ;for _ ,_gbf :=range _gfg {_adb :=_eff +int (_gbf [1]);
if _adb < _eff ||_adb > _bef {return 0,Error {ErrWrite ,"\u0062\u0061d\u0020\u0077\u0072i\u0074\u0065\u0020\u006c\u0065\u006e\u0067\u0074\u0068",int64 (_adb )};};_gce ,_ggb :=_dffd ._dff ._ega .WriteAt (b [_eff :_adb ],_gbf [0]);_edf =_edf +_gce ;if _ggb !=nil {_dffd ._cbc +=int64 (_edf );
return _edf ,Error {ErrWrite ,"\u0075n\u0064\u0065\u0072\u006c\u0079\u0069\u006e\u0067\u0020\u0077\u0072i\u0074\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0020\u0028"+_ggb .Error ()+"\u0029",int64 (_eff )};};_eff =_adb ;};_dffd ._cbc +=int64 (_edf );if _edf !=_bef {_ece =Error {ErrWrite ,"\u0062\u0079t\u0065\u0073\u0020\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0064\u006f\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0077\u0072\u0069\u0074\u0065\u0020\u0073\u0069\u007a\u0065",int64 (_edf )};
}else if _edf < len (b ){_ece =_afd .EOF ;};return _edf ,_ece ;};func (_bdag *Reader )Read (b []byte )(_fggb int ,_agf error ){if _bdag ._fea >=len (_bdag .File ){return 0,_afd .EOF ;};return _bdag .File [_bdag ._fea ].Read (b );};func (_caec *Reader )Next ()(*File ,error ){_caec ._fea ++;
if _caec ._fea >=len (_caec .File ){return nil ,_afd .EOF ;};return _caec .File [_caec ._fea ],nil ;};func _eaab (_aef *directoryEntryFields )(*_a .Buffer ,error ){_eca :=_a .NewBuffer ([]byte {});for _ ,_aad :=range _aef ._ed {if _dede :=_fa .Write (_eca ,_fa .LittleEndian ,&_aad );
_dede !=nil {return nil ,_dede ;};};if _abg :=_fa .Write (_eca ,_fa .LittleEndian ,&_aef ._da );_abg !=nil {return nil ,_abg ;};if _ebg :=_fa .Write (_eca ,_fa .LittleEndian ,&_aef ._ac );_ebg !=nil {return nil ,_ebg ;};if _gff :=_fa .Write (_eca ,_fa .LittleEndian ,&_aef ._aa );
_gff !=nil {return nil ,_gff ;};if _fce :=_fa .Write (_eca ,_fa .LittleEndian ,&_aef ._fe );_fce !=nil {return nil ,_fce ;};if _fca :=_fa .Write (_eca ,_fa .LittleEndian ,&_aef ._de );_fca !=nil {return nil ,_fca ;};if _baa :=_fa .Write (_eca ,_fa .LittleEndian ,&_aef ._bbg );
_baa !=nil {return nil ,_baa ;};if _gfgb :=_fa .Write (_eca ,_fa .LittleEndian ,&_aef ._bd .DataA );_gfgb !=nil {return nil ,_gfgb ;};if _fbe :=_fa .Write (_eca ,_fa .LittleEndian ,&_aef ._bd .DataB );_fbe !=nil {return nil ,_fbe ;};if _ddgfe :=_fa .Write (_eca ,_fa .LittleEndian ,&_aef ._bd .DataC );
_ddgfe !=nil {return nil ,_ddgfe ;};if _ ,_dge :=_eca .Write (_aef ._bd .DataD [:]);_dge !=nil {return nil ,_dge ;};if _ ,_bf :=_eca .Write (_aef ._ea [:]);_bf !=nil {return nil ,_bf ;};if _fda :=_fa .Write (_eca ,_fa .LittleEndian ,&_aef ._eg .Low );_fda !=nil {return nil ,_fda ;
};if _bea :=_fa .Write (_eca ,_fa .LittleEndian ,&_aef ._eg .High );_bea !=nil {return nil ,_bea ;};if _ecc :=_fa .Write (_eca ,_fa .LittleEndian ,&_aef ._efe .Low );_ecc !=nil {return nil ,_ecc ;};if _fef :=_fa .Write (_eca ,_fa .LittleEndian ,&_aef ._efe .High );
_fef !=nil {return nil ,_fef ;};if _fcfc :=_fa .Write (_eca ,_fa .LittleEndian ,&_aef ._dd );_fcfc !=nil {return nil ,_fcfc ;};if _ ,_gfag :=_eca .Write (_aef ._fb [:]);_gfag !=nil {return nil ,_gfag ;};return _eca ,nil ;};const (ErrFormat =iota ;ErrRead ;
ErrSeek ;ErrWrite ;ErrTraverse ;);func _faa (_acbe []byte )*headerFields {_dca :=&headerFields {};_dca ._ecac =_fa .LittleEndian .Uint64 (_acbe [:8]);_dca ._bad =_fa .LittleEndian .Uint16 (_acbe [24:26]);_dca ._dcb =_fa .LittleEndian .Uint16 (_acbe [26:28]);
_dca ._bcc =_fa .LittleEndian .Uint16 (_acbe [30:32]);_dca ._cbe =_fa .LittleEndian .Uint32 (_acbe [40:44]);_dca ._bgf =_fa .LittleEndian .Uint32 (_acbe [44:48]);_dca ._befg =_fa .LittleEndian .Uint32 (_acbe [48:52]);_dca ._egff =_fa .LittleEndian .Uint32 (_acbe [60:64]);
_dca ._fgd =_fa .LittleEndian .Uint32 (_acbe [64:68]);_dca ._bag =_fa .LittleEndian .Uint32 (_acbe [68:72]);_dca ._dedd =_fa .LittleEndian .Uint32 (_acbe [72:76]);var _cae int ;for _bgga :=76;_bgga < 512;_bgga =_bgga +4{_dca ._bfc [_cae ]=_fa .LittleEndian .Uint32 (_acbe [_bgga :_bgga +4]);
_cae ++;};return _dca ;};const (_ee uint8 =0x0;_be uint8 =0x1;);func (_ace *File )changeSize (_egg int64 )error {if _egg ==_ace .Size {return nil ;};var _caf *File ;for _ ,_cgb :=range _ace ._dff ._aebd {if _cgb .Name ==_ace .Name {_caf =_cgb ;break ;};
};if _caf ==nil {return _f .Errorf ("\u004e\u006f\u0020\u0064\u0069\u0072e\u0063\u0074\u006f\u0072\u0079\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066o\u0072\u0020\u0061\u0020\u0066\u0069\u006ce\u003a\u0020\u0025\u0073",_ace .Name );};_ced :=_a .NewBuffer ([]byte {});
if _dee :=_fa .Write (_ced ,_fa .LittleEndian ,_egg );_dee !=nil {return _dee ;};for _daeb ,_ded :=range _ced .Bytes (){_caf ._fb [_daeb ]=_ded ;};var _eb int64 ;var _abe bool ;if _ace .Size < _gcb {_abe =true ;_eb =int64 (_ageg );}else {_eb =int64 (_ace ._dff ._beaa );
};_bbd :=int ((_ace .Size -1)/_eb )+1;_fd :=int ((_egg -1)/_eb )+1;if _bbd < _fd {_gca ,_dga :=_ace .findLast (_abe );if _dga !=nil {return _dga ;};_bbb ,_dga :=_ace ._dff .findNextFreeSector (_abe );if _dga !=nil {return _dga ;};for _edb :=_fd -_bbd ;
_edb > 0;_edb --{if _bc :=_ace ._dff .saveToFatLocs (_gca ,_bbb ,_abe );_bc !=nil {return _bc ;};if _edb > 1{_gca =_bbb ;_bbb ++;}else if _dgd :=_ace ._dff .saveToFatLocs (_bbb ,_abf ,_abe );_dgd !=nil {return _dgd ;};};}else if _bbd > _fd {_ccb :=_ace ._dd ;
var _agb error ;for _fcf :=0;_fcf < _fd -1;_fcf ++{_ccb ,_agb =_ace ._dff .findNext (_ccb ,_abe );if _agb !=nil {return _agb ;};};if _db :=_ace ._dff .saveToFatLocs (_ccb ,_abf ,_abe );_db !=nil {return _db ;};};_ace .Size =_egg ;return nil ;};func (_bdc fileInfo )Mode ()_ef .FileMode {return _bdc .File .mode ()};
func (_cad *File )SetEntryContent (b []byte )error {if _age :=_cad .ensureWriterAt ();_age !=nil {return _age ;};_cad .reset ();if _cf :=_cad .changeSize (int64 (len (b )));_cf !=nil {return _cf ;};_ ,_ffd :=_cad .Write (b );return _ffd ;};func _daa (_fbg []byte )*directoryEntryFields {_gd :=&directoryEntryFields {};
for _df :=range _gd ._ed {_gd ._ed [_df ]=_fa .LittleEndian .Uint16 (_fbg [_df *2:_df *2+2]);};_gd ._da =_fa .LittleEndian .Uint16 (_fbg [64:66]);_gd ._ac =uint8 (_fbg [66]);_gd ._aa =uint8 (_fbg [67]);_gd ._fe =_fa .LittleEndian .Uint32 (_fbg [68:72]);
_gd ._de =_fa .LittleEndian .Uint32 (_fbg [72:76]);_gd ._bbg =_fa .LittleEndian .Uint32 (_fbg [76:80]);_gd ._bd =_ga .MustGuid (_fbg [80:96]);copy (_gd ._ea [:],_fbg [96:100]);_gd ._eg =_ga .MustFileTime (_fbg [100:108]);_gd ._efe =_ga .MustFileTime (_fbg [108:116]);
_gd ._dd =_fa .LittleEndian .Uint32 (_fbg [116:120]);copy (_gd ._fb [:],_fbg [120:128]);return _gd ;};func (_gfc *Reader )Debug ()map[string ][]uint32 {_aace :=map[string ][]uint32 {"s\u0065\u0063\u0074\u006f\u0072\u0020\u0073\u0069\u007a\u0065":[]uint32 {_gfc ._beaa },"\u006d\u0069\u006e\u0069\u0020\u0066\u0061\u0074\u0020\u006c\u006f\u0063\u0073":_gfc ._cbf ._gcd ,"\u006d\u0069n\u0069\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u006c\u006f\u0063\u0073":_gfc ._cbf ._acefd ,"\u0064\u0069r\u0065\u0063\u0074o\u0072\u0079\u0020\u0073\u0065\u0063\u0074\u006f\u0072":[]uint32 {_gfc ._cbf ._befg },"\u006d\u0069\u006e\u0069 s\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0074\u0061\u0072\u0074\u002f\u0073\u0069z\u0065":[]uint32 {_gfc .File [0]._dd ,_fa .LittleEndian .Uint32 (_gfc .File [0]._fb [:])}};
for _acbb ,_eacg :=_gfc .Next ();_eacg ==nil ;_acbb ,_eacg =_gfc .Next (){_aace [_acbb .Name +" \u0073\u0074\u0061\u0072\u0074\u002f\u0073\u0069\u007a\u0065"]=[]uint32 {_acbb ._dd ,_fa .LittleEndian .Uint32 (_acbb ._fb [:])};};return _aace ;};func (_adgd *Reader )exportFAT (_eeeb *_gg .Writer ,_bfa []uint32 )error {if _adgd ._cbf ._bgf ==0{return nil ;
};_cdae :=_a .NewBuffer ([]byte {});if _fec :=_fa .Write (_cdae ,_fa .LittleEndian ,_fdc );_fec !=nil {return _fec ;};for _aaaf :=0;_aaaf < len (_bfa )-1;_aaaf ++{for _egb :=_bfa [_aaaf ];_egb < _bfa [_aaaf +1]-1;_egb ++{if _cgbf :=_fa .Write (_cdae ,_fa .LittleEndian ,_egb );
_cgbf !=nil {return _cgbf ;};};if _fcdb :=_fa .Write (_cdae ,_fa .LittleEndian ,_abf );_fcdb !=nil {return _fcdb ;};};_ecf :=512;for _ ,_eag :=range _cdae .Bytes (){if _bccg :=_eeeb .WriteByteAt (_eag ,_ecf );_bccg !=nil {return _bccg ;};_ecf ++;};return nil ;
};type fileInfo struct{*File };