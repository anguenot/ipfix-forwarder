package main

import (
	"github.com/calmh/ipfix"
)

// Nokia constants
const (
	VendorNokia = "vendor-nokia"
)

// add Nokia NAT vendor fields to the dictionary so that it will be resolved by
// the interpreter
func includeNokiaFields(i *ipfix.Interpreter) {
	nokia91 := ipfix.DictionaryEntry{
		Name:         "natInsideSvcid",
		FieldID:      91,
		EnterpriseID: 637,
		Type:         ipfix.Uint16,
	}
	i.AddDictionaryEntry(nokia91)

	nokia92 := ipfix.DictionaryEntry{
		Name:         "natOutsideSvcid",
		FieldID:      92,
		EnterpriseID: 637,
		Type:         ipfix.Uint16,
	}
	i.AddDictionaryEntry(nokia92)

	nokia93 := ipfix.DictionaryEntry{
		Name:         "natSubString",
		FieldID:      93,
		EnterpriseID: 637,
		Type:         ipfix.String,
	}
	i.AddDictionaryEntry(nokia93)
}
