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

// Package memstore implements tempStorage interface
// by using memory as a storage
package memstore ;import (_cc "encoding/hex";_aa "errors";_ad "fmt";_d "github.com/unidoc/unioffice/common/tempstorage";_c "io";_ca "io/ioutil";_cd "math/rand";_e "sync";);

// Open returns tempstorage File object by name
func (_bc *memStorage )Open (path string )(_d .File ,error ){_cg ,_ddc :=_bc ._ga .Load (path );if !_ddc {return nil ,_aa .New (_ad .Sprintf ("\u0043\u0061\u006eno\u0074\u0020\u006f\u0070\u0065\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0025\u0073",path ));};return &memFile {_ab :_cg .(*memDataCell )},nil ;};

// TempFile creates a new empty file in the storage and returns it
func (_fc *memStorage )TempFile (dir ,pattern string )(_d .File ,error ){_cf :=dir +"\u002f"+_ac (pattern );_bd :=&memDataCell {_gd :_cf ,_dd :[]byte {}};_eg :=&memFile {_ab :_bd };_fc ._ga .Store (_cf ,_bd );return _eg ,nil ;};

// Name returns the filename of the underlying memDataCell
func (_adb *memFile )Name ()string {return _adb ._ab ._gd };func _ac (_ba string )string {_cge ,_ :=_dc (6);return _ba +_cge };type memDataCell struct{_gd string ;_dd []byte ;_gg int64 ;};

// ReadAt reads from the underlying memDataCell at an offset provided in order to implement ReaderAt interface.
// It does not affect f.readOffset.
func (_ag *memFile )ReadAt (p []byte ,readOffset int64 )(int ,error ){_af :=_ag ._ab ._gg ;_be :=int64 (len (p ));if _be > _af {_be =_af ;p =p [:_be ];};if readOffset >=_af {return 0,_c .EOF ;};_f :=readOffset +_be ;if _f >=_af {_f =_af ;};_dbg :=copy (p ,_ag ._ab ._dd [readOffset :_f ]);return _dbg ,nil ;};

// Add reads a file from a disk and adds it to the storage
func (_egd *memStorage )Add (path string )error {_ ,_ggc :=_egd ._ga .Load (path );if _ggc {return nil ;};_df ,_bb :=_ca .ReadFile (path );if _bb !=nil {return _bb ;};_egd ._ga .Store (path ,&memDataCell {_gd :path ,_dd :_df ,_gg :int64 (len (_df ))});return nil ;};func _dc (_fcf int )(string ,error ){_dcc :=make ([]byte ,_fcf );if _ ,_ec :=_cd .Read (_dcc );_ec !=nil {return "",_ec ;};return _cc .EncodeToString (_dcc ),nil ;};type memStorage struct{_ga _e .Map };

// RemoveAll removes all files according to the dir argument prefix
func (_ccd *memStorage )RemoveAll (dir string )error {_ccd ._ga .Range (func (_dbb ,_cgc interface{})bool {_ccd ._ga .Delete (_dbb );return true });return nil ;};type memFile struct{_ab *memDataCell ;_b int64 ;};

// Write writes to the end of the underlying memDataCell in order to implement Writer interface
func (_fd *memFile )Write (p []byte )(int ,error ){_fd ._ab ._dd =append (_fd ._ab ._dd ,p ...);_fd ._ab ._gg +=int64 (len (p ));return len (p ),nil ;};

// SetAsStorage sets temp storage as a memory storage
func SetAsStorage (){_cca :=memStorage {_ga :_e .Map {}};_d .SetAsStorage (&_cca )};

// Read reads from the underlying memDataCell in order to implement Reader interface
func (_g *memFile )Read (p []byte )(int ,error ){_gb :=_g ._b ;_da :=_g ._ab ._gg ;_db :=int64 (len (p ));if _db > _da {_db =_da ;p =p [:_db ];};if _gb >=_da {return 0,_c .EOF ;};_ee :=_gb +_db ;if _ee >=_da {_ee =_da ;};_bf :=copy (p ,_g ._ab ._dd [_gb :_ee ]);_g ._b =_ee ;return _bf ,nil ;};

// Close is not applicable in this implementation
func (_gc *memFile )Close ()error {return nil };

// TempDir creates a name for a new temp directory using a pattern argument
func (_ge *memStorage )TempDir (pattern string )(string ,error ){return _ac (pattern ),nil };