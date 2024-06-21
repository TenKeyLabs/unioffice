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

package wildcard ;func _gag (_bg ,_dc []rune ,_ag bool )bool {for len (_dc )> 0{switch _dc [0]{default:if len (_bg )==0||_bg [0]!=_dc [0]{return false ;};case '?':if len (_bg )==0&&!_ag {return false ;};case '*':return _gag (_bg ,_dc [1:],_ag )||(len (_bg )> 0&&_gag (_bg [1:],_dc ,_ag ));
};_bg =_bg [1:];_dc =_dc [1:];};return len (_bg )==0&&len (_dc )==0;};func Match (pattern ,name string )(_dg bool ){if pattern ==""{return name ==pattern ;};if pattern =="\u002a"{return true ;};_c :=make ([]rune ,0,len (name ));_gc :=make ([]rune ,0,len (pattern ));
for _ ,_ga :=range name {_c =append (_c ,_ga );};for _ ,_dad :=range pattern {_gc =append (_gc ,_dad );};_gaa :=false ;return _gag (_c ,_gc ,_gaa );};func Index (pattern ,name string )(_cg int ){if pattern ==""||pattern =="\u002a"{return 0;};_eb :=make ([]rune ,0,len (name ));
_cf :=make ([]rune ,0,len (pattern ));for _ ,_cd :=range name {_eb =append (_eb ,_cd );};for _ ,_bd :=range pattern {_cf =append (_cf ,_bd );};return _gg (_eb ,_cf ,0);};func MatchSimple (pattern ,name string )bool {if pattern ==""{return name ==pattern ;
};if pattern =="\u002a"{return true ;};_d :=make ([]rune ,0,len (name ));_e :=make ([]rune ,0,len (pattern ));for _ ,_da :=range name {_d =append (_d ,_da );};for _ ,_bc :=range pattern {_e =append (_e ,_bc );};_f :=true ;return _gag (_d ,_e ,_f );};func _gg (_ge ,_bdd []rune ,_ef int )int {for len (_bdd )> 0{switch _bdd [0]{default:if len (_ge )==0{return -1;
};if _ge [0]!=_bdd [0]{return _gg (_ge [1:],_bdd ,_ef +1);};case '?':if len (_ge )==0{return -1;};case '*':if len (_ge )==0{return -1;};_fd :=_gg (_ge ,_bdd [1:],_ef );if _fd !=-1{return _ef ;}else {_fd =_gg (_ge [1:],_bdd ,_ef );if _fd !=-1{return _ef ;
}else {return -1;};};};_ge =_ge [1:];_bdd =_bdd [1:];};return _ef ;};