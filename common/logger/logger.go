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

package logger ;import (_b "fmt";_be "io";_d "os";_a "path/filepath";_c "runtime";);

// Warning logs warning message.
func (_ce ConsoleLogger )Warning (format string ,args ...interface{}){if _ce .LogLevel >=LogLevelWarning {_cgb :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_ce .output (_d .Stdout ,_cgb ,format ,args ...);};};

// Debug logs debug message.
func (_fd ConsoleLogger )Debug (format string ,args ...interface{}){if _fd .LogLevel >=LogLevelDebug {_cea :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_fd .output (_d .Stdout ,_cea ,format ,args ...);};};

// Info does nothing for dummy logger.
func (DummyLogger )Info (format string ,args ...interface{}){};

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger (logLevel LogLevel ,writer _be .Writer )*WriterLogger {logger :=WriterLogger {Output :writer ,LogLevel :logLevel };return &logger ;};

// IsLogLevel returns true from dummy logger.
func (DummyLogger )IsLogLevel (level LogLevel )bool {return true };

// Logger is the interface used for logging in the unipdf package.
type Logger interface{Error (_dg string ,_g ...interface{});Warning (_f string ,_ec ...interface{});Notice (_gd string ,_beg ...interface{});Info (_cc string ,_df ...interface{});Debug (_af string ,_ab ...interface{});Trace (_bf string ,_bfb ...interface{});
IsLogLevel (_bg LogLevel )bool ;};

// Error logs error message.
func (_gb WriterLogger )Error (format string ,args ...interface{}){if _gb .LogLevel >=LogLevelError {_acb :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_gb .logToWriter (_gb .Output ,_acb ,format ,args ...);};};

// DummyLogger does nothing.
type DummyLogger struct{};

// Info logs info message.
func (_abg WriterLogger )Info (format string ,args ...interface{}){if _abg .LogLevel >=LogLevelInfo {_bgg :="\u005bI\u004e\u0046\u004f\u005d\u0020";_abg .logToWriter (_abg .Output ,_bgg ,format ,args ...);};};

// Info logs info message.
func (_ac ConsoleLogger )Info (format string ,args ...interface{}){if _ac .LogLevel >=LogLevelInfo {_cgc :="\u005bI\u004e\u0046\u004f\u005d\u0020";_ac .output (_d .Stdout ,_cgc ,format ,args ...);};};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_eb ConsoleLogger )IsLogLevel (level LogLevel )bool {return _eb .LogLevel >=level };

// Notice logs notice message.
func (_ed ConsoleLogger )Notice (format string ,args ...interface{}){if _ed .LogLevel >=LogLevelNotice {_ad :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_ed .output (_d .Stdout ,_ad ,format ,args ...);};};

// Debug logs debug message.
func (_acfe WriterLogger )Debug (format string ,args ...interface{}){if _acfe .LogLevel >=LogLevelDebug {_bgf :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_acfe .logToWriter (_acfe .Output ,_bgf ,format ,args ...);};};

// Debug does nothing for dummy logger.
func (DummyLogger )Debug (format string ,args ...interface{}){};

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{LogLevel LogLevel ;};

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger (logger Logger ){Log =logger };

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct{LogLevel LogLevel ;Output _be .Writer ;};

// Error logs error message.
func (_gg ConsoleLogger )Error (format string ,args ...interface{}){if _gg .LogLevel >=LogLevelError {_cg :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_gg .output (_d .Stdout ,_cg ,format ,args ...);};};

// Notice does nothing for dummy logger.
func (DummyLogger )Notice (format string ,args ...interface{}){};

// NewConsoleLogger creates new console logger.
func NewConsoleLogger (logLevel LogLevel )*ConsoleLogger {return &ConsoleLogger {LogLevel :logLevel }};func _bca (_gf _be .Writer ,_gge string ,_cb string ,_aa ...interface{}){_ ,_da ,_de ,_gae :=_c .Caller (3);if !_gae {_da ="\u003f\u003f\u003f";_de =0;
}else {_da =_a .Base (_da );};_bgfb :=_b .Sprintf ("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ",_gge ,_da ,_de )+_cb +"\u000a";_b .Fprintf (_gf ,_bgfb ,_aa ...);};var Log Logger =DummyLogger {};

// Trace does nothing for dummy logger.
func (DummyLogger )Trace (format string ,args ...interface{}){};

// Notice logs notice message.
func (_ede WriterLogger )Notice (format string ,args ...interface{}){if _ede .LogLevel >=LogLevelNotice {_dfb :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_ede .logToWriter (_ede .Output ,_dfb ,format ,args ...);};};

// LogLevel is the verbosity level for logging.
type LogLevel int ;

// Trace logs trace message.
func (_gba WriterLogger )Trace (format string ,args ...interface{}){if _gba .LogLevel >=LogLevelTrace {_caf :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_gba .logToWriter (_gba .Output ,_caf ,format ,args ...);};};

// Warning does nothing for dummy logger.
func (DummyLogger )Warning (format string ,args ...interface{}){};

// Error does nothing for dummy logger.
func (DummyLogger )Error (format string ,args ...interface{}){};

// Trace logs trace message.
func (_ge ConsoleLogger )Trace (format string ,args ...interface{}){if _ge .LogLevel >=LogLevelTrace {_bc :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_ge .output (_d .Stdout ,_bc ,format ,args ...);};};func (_ebg ConsoleLogger )output (_acf _be .Writer ,_ebd string ,_cca string ,_acg ...interface{}){_bca (_acf ,_ebd ,_cca ,_acg ...);
};func (_edf WriterLogger )logToWriter (_caa _be .Writer ,_dc string ,_dd string ,_ga ...interface{}){_bca (_caa ,_dc ,_dd ,_ga );};const (LogLevelTrace LogLevel =5;LogLevelDebug LogLevel =4;LogLevelInfo LogLevel =3;LogLevelNotice LogLevel =2;LogLevelWarning LogLevel =1;
LogLevelError LogLevel =0;);

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_ca WriterLogger )IsLogLevel (level LogLevel )bool {return _ca .LogLevel >=level };

// Warning logs warning message.
func (_bfc WriterLogger )Warning (format string ,args ...interface{}){if _bfc .LogLevel >=LogLevelWarning {_geg :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_bfc .logToWriter (_bfc .Output ,_geg ,format ,args ...);};};