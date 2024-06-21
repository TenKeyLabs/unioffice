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

package tempstorage ;import _e "io";

// TempFile creates new empty file in the storage and returns it.
func TempFile (dir ,pattern string )(File ,error ){return _cc .TempFile (dir ,pattern )};

// Open returns tempstorage File object by name.
func Open (path string )(File ,error ){return _cc .Open (path )};var _cc storage ;type storage interface{Open (_g string )(File ,error );TempFile (_c ,_b string )(File ,error );TempDir (_d string )(string ,error );RemoveAll (_ec string )error ;Add (_eb string )error ;
};

// RemoveAll removes all files according to the dir argument prefix.
func RemoveAll (dir string )error {return _cc .RemoveAll (dir )};

// Add reads a file from a disk and adds it to the storage.
func Add (path string )error {return _cc .Add (path )};

// SetAsStorage changes temporary storage to newStorage.
func SetAsStorage (newStorage storage ){_cc =newStorage };

// File is a representation of a storage file
// with Read, Write, Close and Name methods identical to os.File.
type File interface{_e .Reader ;_e .ReaderAt ;_e .Writer ;_e .Closer ;Name ()string ;};

// TempDir creates a name for a new temp directory using a pattern argument.
func TempDir (pattern string )(string ,error ){return _cc .TempDir (pattern )};