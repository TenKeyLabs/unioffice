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

// Package license helps manage commercial licenses and check if they
// are valid for the version of UniOffice used.
package license ;import _e "github.com/unidoc/unioffice/internal/license";

// SetLicenseKey sets and validates the license key.
func SetLicenseKey (content string ,customerName string )error {return _e .SetLicenseKey (content ,customerName );};

// MakeUnlicensedKey returns a default key.
func MakeUnlicensedKey ()*LicenseKey {return _e .MakeUnlicensedKey ()};const (LicenseTierUnlicensed =_e .LicenseTierUnlicensed ;LicenseTierCommunity =_e .LicenseTierCommunity ;LicenseTierIndividual =_e .LicenseTierIndividual ;LicenseTierBusiness =_e .LicenseTierBusiness ;
);

// GetLicenseKey returns the currently loaded license key.
func GetLicenseKey ()*LicenseKey {return _e .GetLicenseKey ()};

// LegacyLicense holds the old-style unioffice license information.
type LegacyLicense =_e .LegacyLicense ;

// SetMeteredKeyPersistentCache sets the metered License API Key persistent cache.
// Default value `true`, set to `false` will report the usage immediately to license server,
// this can be used when there's no access to persistent data storage.
func SetMeteredKeyPersistentCache (val bool ){_e .SetMeteredKeyPersistentCache (val )};

// LicenseKey represents a loaded license key.
type LicenseKey =_e .LicenseKey ;

// SetMeteredKey sets the metered License API key required for SaaS operation.
// Document usage is reported periodically for the product to function correctly.
func SetMeteredKey (apiKey string )error {return _e .SetMeteredKey (apiKey )};

// SetLegacyLicenseKey installs a legacy license code. License codes issued prior to June 2019.
// Will be removed at some point in a future major version.
func SetLegacyLicenseKey (s string )error {return _e .SetLegacyLicenseKey (s )};

// GetMeteredState checks the currently used metered document usage status,
// documents used and credits available.
func GetMeteredState ()(_e .MeteredStatus ,error ){return _e .GetMeteredState ()};

// LegacyLicenseType is the type of license
type LegacyLicenseType =_e .LegacyLicenseType ;