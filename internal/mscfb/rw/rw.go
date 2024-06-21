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

package rw ;import (_e "bytes";_b "encoding/binary";_d "errors";_ee "fmt";_g "io";_cf "io/ioutil";_c "reflect";);func (_ed *Writer )WriteStringProperty (s string )error {_ed .align (4);_eec :=[]byte (s );if _cff :=_b .Write (_ed ,_b .LittleEndian ,&_eec );
_cff !=nil {return _cff ;};return nil ;};func (_gb *Writer )Write (p []byte )(_fde int ,_fad error ){_af ,_cgc :=_gb .tryGrowByReslice (len (p ));if !_cgc {var _dbe error ;_af ,_dbe =_gb .grow (len (p ));if _dbe !=nil {return 0,_dbe ;};};return copy (_gb ._db [_af :],p ),nil ;
};func (_edb *Writer )curPos ()int {return int (_edb .Cap ())-_edb .Len ()};type Reader struct{*_e .Reader };func (_feg *Reader )ReadPairProperty (p interface{})error {if _dfb :=_feg .align (4);_dfb !=nil {return _dfb ;};_de :=_c .ValueOf (p );for _de .Kind ()==_c .Ptr {_de =_de .Elem ();
};if !_de .IsValid (){return _ee .Errorf ("\u0076a\u006cu\u0065\u0020\u0069\u0073\u0020n\u006f\u0074 \u0076\u0061\u006c\u0069\u0064");};if _fdf :=_b .Read (_feg ,_b .LittleEndian ,p );_fdf !=nil {return _fdf ;};return nil ;};func (_da *Writer )reset (){_da ._db =_da ._db [:0];
_da ._ec =0};func NewWriter ()*Writer {return &Writer {_db :[]byte {}}};func (_fa *Writer )Cap ()int {return cap (_fa ._db )};func (_a *Writer )WriteByteAt (b byte ,off int )error {if off >=len (_a ._db ){return _d .New ("\u004f\u0075\u0074\u0020\u006f\u0066\u0020\u0062\u006f\u0075\u006e\u0064\u0073");
};_a ._db [off ]=b ;return nil ;};const _fg =int (^uint (0)>>1);var _bc =_d .New ("r\u0077.\u0057\u0072\u0069\u0074\u0065\u0072\u003a\u0020t\u006f\u006f\u0020\u006car\u0067\u0065");func (_gfd *Writer )WriteProperty (a interface{})error {if _baa :=_gfd .align (int (_c .TypeOf (a ).Size ()));
_baa !=nil {return _baa ;};return _gfd .WritePropertyNoAlign (a );};type Writer struct{_db []byte ;_ec int ;};func PushLeftUI32 (v uint32 ,flag bool )uint32 {v >>=1;if flag {v |=1<<31;};return v ;};func PopRightUI32 (v uint32 )(bool ,uint32 ){return (v &uint32 (1))==1,v >>1};
func _ggd (_bcd int )[]byte {defer func (){if recover ()!=nil {panic (_bc );};}();return make ([]byte ,_bcd );};func (_gac *Writer )WriteTo (wTo _g .Writer )(_bac int64 ,_dc error ){if _acc :=_gac .Len ();_acc > 0{_dbc ,_bb :=wTo .Write (_gac ._db [_gac ._ec :]);
if _dbc > _acc {return 0,_d .New ("\u0072\u0077\u002e\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0054\u006f\u003a\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0057\u0072\u0069t\u0065\u0020c\u006f\u0075\u006e\u0074");};
_gac ._ec +=_dbc ;_bac =int64 (_dbc );if _bb !=nil {return _bac ,_bb ;};if _dbc !=_acc {return _bac ,_g .ErrShortWrite ;};};_gac .reset ();return _bac ,nil ;};func (_ccd *Reader )ReadStringProperty (n uint32 )(string ,error ){if _bd :=_ccd .align (4);_bd !=nil {return "",_bd ;
};_fdd :=make ([]byte ,n );if _dg :=_b .Read (_ccd ,_b .LittleEndian ,&_fdd );_dg !=nil {return "",_dg ;};return string (_fdd ),nil ;};func (_ac *Writer )Bytes ()[]byte {return _ac ._db };func (_bg *Reader )align (_ga int )error {return _bg .skip ((_ga -_bg .curPos ()%_ga )%_ga )};
func PushLeftUI64 (v uint64 ,flag bool )uint64 {v >>=1;if flag {v |=1<<63;};return v ;};func (_ce *Writer )WritePropertyNoAlign (a interface{})error {if _ca :=_b .Write (_ce ,_b .LittleEndian ,a );_ca !=nil {return _ca ;};return nil ;};func PopRightUI64 (v uint64 )(bool ,uint64 ){return (v &uint64 (1))==1,v >>1};
func (_dga *Writer )AlignLength (alignTo int )error {_fc :=_dga .Len ()%alignTo ;if _fc > 0{_ ,_gg :=_dga .Write (make ([]byte ,alignTo -_fc ));if _gg !=nil {return _gg ;};};return nil ;};const _bf =64;func NewReader (b []byte )(*Reader ,error ){return &Reader {_e .NewReader (b )},nil };
func (_df *Reader )curPos ()int {return int (_df .Size ())-_df .Len ()};func (_eeg *Writer )Len ()int {return len (_eeg ._db )-_eeg ._ec };func (_beb *Writer )tryGrowByReslice (_fef int )(int ,bool ){if _dbg :=len (_beb ._db );_fef <=cap (_beb ._db )-_dbg {_beb ._db =_beb ._db [:_dbg +_fef ];
return _dbg ,true ;};return 0,false ;};func (_ba *Reader )ReadProperty (a interface{})error {_cc :=_c .ValueOf (a );for _cc .Kind ()==_c .Ptr {_cc =_cc .Elem ();};if !_cc .IsValid (){return _ee .Errorf ("\u0076a\u006cu\u0065\u0020\u0069\u0073\u0020n\u006f\u0074 \u0076\u0061\u006c\u0069\u0064");
};if _fe :=_ba .align (int (_cc .Type ().Size ()));_fe !=nil {return _fe ;};if _eeb :=_b .Read (_ba ,_b .LittleEndian ,a );_eeb !=nil {return _eeb ;};return nil ;};func (_gf *Reader )skip (_fd int )error {_ ,_eef :=_g .CopyN (_cf .Discard ,_gf ,int64 (_fd ));
if _eef !=nil {return _eef ;};return nil ;};func (_bge *Writer )grow (_dee int )(int ,error ){_bacf :=_bge .Len ();if _bacf ==0&&_bge ._ec !=0{_bge .reset ();};if _afb ,_bcc :=_bge .tryGrowByReslice (_dee );_bcc {return _afb ,nil ;};if _bge ._db ==nil &&_dee <=_bf {_bge ._db =make ([]byte ,_dee ,_bf );
return 0,nil ;};_ddd :=cap (_bge ._db );if _dee <=_ddd /2-_bacf {copy (_bge ._db ,_bge ._db [_bge ._ec :]);}else if _ddd > _fg -_ddd -_dee {return 0,_bc ;}else {_bag :=_ggd (2*_ddd +_dee );copy (_bag ,_bge ._db [_bge ._ec :]);_bge ._db =_bag ;};_bge ._ec =0;
_bge ._db =_bge ._db [:_bacf +_dee ];return _bacf ,nil ;};func (_ccdg *Writer )FillWithByte (fillSize int ,b byte )error {for _ge :=0;_ge < fillSize ;_ge ++{if _aag :=_ccdg .WritePropertyNoAlign (b );_aag !=nil {return _aag ;};};return nil ;};func (_cg *Writer )align (_ea int )error {return _cg .Skip ((_ea -(_cg .Len ())%_ea )%_ea )};
func (_bgf *Writer )Skip (n int )error {if n ==0{return nil ;};_ ,_be :=_bgf .Write (make ([]byte ,n ));return _be ;};