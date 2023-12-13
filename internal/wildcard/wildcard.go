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

package wildcard ;func MatchSimple (pattern ,name string )bool {if pattern ==""{return name ==pattern ;};if pattern =="\u002a"{return true ;};_f :=make ([]rune ,0,len (name ));_fd :=make ([]rune ,0,len (pattern ));for _ ,_e :=range name {_f =append (_f ,_e );
};for _ ,_c :=range pattern {_fd =append (_fd ,_c );};_cf :=true ;return _dd (_f ,_fd ,_cf );};func _dd (_bg ,_bb []rune ,_dde bool )bool {for len (_bb )> 0{switch _bb [0]{default:if len (_bg )==0||_bg [0]!=_bb [0]{return false ;};case '?':if len (_bg )==0&&!_dde {return false ;
};case '*':return _dd (_bg ,_bb [1:],_dde )||(len (_bg )> 0&&_dd (_bg [1:],_bb ,_dde ));};_bg =_bg [1:];_bb =_bb [1:];};return len (_bg )==0&&len (_bb )==0;};func Match (pattern ,name string )(_a bool ){if pattern ==""{return name ==pattern ;};if pattern =="\u002a"{return true ;
};_gf :=make ([]rune ,0,len (name ));_cg :=make ([]rune ,0,len (pattern ));for _ ,_fb :=range name {_gf =append (_gf ,_fb );};for _ ,_bf :=range pattern {_cg =append (_cg ,_bf );};_ga :=false ;return _dd (_gf ,_cg ,_ga );};func Index (pattern ,name string )(_af int ){if pattern ==""||pattern =="\u002a"{return 0;
};_cd :=make ([]rune ,0,len (name ));_bbf :=make ([]rune ,0,len (pattern ));for _ ,_ab :=range name {_cd =append (_cd ,_ab );};for _ ,_fa :=range pattern {_bbf =append (_bbf ,_fa );};return _cde (_cd ,_bbf ,0);};func _cde (_gfa ,_aa []rune ,_bd int )int {for len (_aa )> 0{switch _aa [0]{default:if len (_gfa )==0{return -1;
};if _gfa [0]!=_aa [0]{return _cde (_gfa [1:],_aa ,_bd +1);};case '?':if len (_gfa )==0{return -1;};case '*':if len (_gfa )==0{return -1;};_ce :=_cde (_gfa ,_aa [1:],_bd );if _ce !=-1{return _bd ;}else {_ce =_cde (_gfa [1:],_aa ,_bd );if _ce !=-1{return _bd ;
}else {return -1;};};};_gfa =_gfa [1:];_aa =_aa [1:];};return _bd ;};