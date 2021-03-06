/*
Copyright 2018 The pdfcpu Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pdfcpu

import (
	"github.com/hhrutter/pdfcpu/pkg/log"
	"github.com/pkg/errors"
)

func validateAAPLAKExtrasDictEntry(xRefTable *XRefTable, dict *PDFDict, dictName, entryName string, required bool, sinceVersion PDFVersion) error {

	// No documentation for this PDF-Extension - purely speculative implementation.

	d, err := validateDictEntry(xRefTable, dict, dictName, entryName, required, sinceVersion, nil)
	if err != nil || d == nil {
		return err
	}

	dictName = "AAPLAKExtrasDict"

	// AAPL:AKAnnotationObject, string
	_, err = validateStringEntry(xRefTable, d, dictName, "AAPL:AKAnnotationObject", OPTIONAL, sinceVersion, nil)
	if err != nil {
		return err
	}

	// AAPL:AKPDFAnnotationDictionary, annotationDict
	ad, err := validateDictEntry(xRefTable, d, dictName, "AAPL:AKPDFAnnotationDictionary", OPTIONAL, sinceVersion, nil)
	if err != nil {
		return err
	}

	_, err = validateAnnotationDict(xRefTable, ad)
	if err != nil {
		return err
	}

	return nil
}

func validateBorderEffectDictEntry(xRefTable *XRefTable, dict *PDFDict, dictName, entryName string, required bool, sinceVersion PDFVersion) error {

	// see 12.5.4

	d, err := validateDictEntry(xRefTable, dict, dictName, entryName, required, sinceVersion, nil)
	if err != nil || d == nil {
		return err
	}

	dictName = "borderEffectDict"

	// S, optional, name, S or C
	_, err = validateNameEntry(xRefTable, d, dictName, "S", OPTIONAL, V10, func(s string) bool { return s == "S" || s == "C" })
	if err != nil {
		return err
	}

	// I, optional, number in the range 0 to 2
	_, err = validateNumberEntry(xRefTable, d, dictName, "I", OPTIONAL, V10, func(f float64) bool { return 0 <= f && f <= 2 }) // validation missing
	if err != nil {
		return err
	}

	return nil
}

func validateBorderStyleDict(xRefTable *XRefTable, dict *PDFDict, dictName, entryName string, required bool, sinceVersion PDFVersion) error {

	// see 12.5.4

	d, err := validateDictEntry(xRefTable, dict, dictName, entryName, required, sinceVersion, nil)
	if err != nil || d == nil {
		return err
	}

	dictName = "borderStyleDict"

	// Type, optional, name, "Border"
	_, err = validateNameEntry(xRefTable, d, dictName, "Type", OPTIONAL, V10, func(s string) bool { return s == "Border" })
	if err != nil {
		return err
	}

	// W, optional, number, border width in points
	_, err = validateNumberEntry(xRefTable, d, dictName, "W", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// S, optional, name, border style
	validate := func(s string) bool { return memberOf(s, []string{"S", "D", "B", "I", "U", "A"}) }
	_, err = validateNameEntry(xRefTable, d, dictName, "S", OPTIONAL, V10, validate)
	if err != nil {
		return err
	}

	// D, optional, dash array
	_, err = validateNumberArrayEntry(xRefTable, d, dictName, "D", OPTIONAL, V10, func(a PDFArray) bool { return len(a) <= 2 })

	return err
}

func validateIconFitDictEntry(xRefTable *XRefTable, dict *PDFDict, dictName, entryName string, required bool, sinceVersion PDFVersion) error {

	// see table 247

	d, err := validateDictEntry(xRefTable, dict, dictName, entryName, required, sinceVersion, nil)
	if err != nil || d == nil {
		return err
	}

	dictName = "iconFitDict"

	// SW, optional, name, A,B,S,N
	validate := func(s string) bool { return memberOf(s, []string{"A", "B", "S", "N"}) }
	_, err = validateNameEntry(xRefTable, d, dictName, "SW", OPTIONAL, V10, validate)
	if err != nil {
		return err
	}

	// S, optional, name, A,P
	_, err = validateNameEntry(xRefTable, d, dictName, "S", OPTIONAL, V10, func(s string) bool { return s == "A" || s == "P" })
	if err != nil {
		return err
	}

	// A,optional, array of 2 numbers between 0.0 and 1.0
	_, err = validateNumberArrayEntry(xRefTable, d, dictName, "A", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// FB, optional, bool, since V1.5
	_, err = validateBooleanEntry(xRefTable, d, dictName, "FB", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	return nil
}

func validateAppearanceCharacteristicsDictEntry(xRefTable *XRefTable, dict *PDFDict, dictName, entryName string, required bool, sinceVersion PDFVersion) error {

	// see 12.5.6.19

	d, err := validateDictEntry(xRefTable, dict, dictName, entryName, required, sinceVersion, nil)
	if err != nil || d == nil {
		return err
	}

	dictName = "appCharDict"

	// R, optional, integer
	_, err = validateIntegerEntry(xRefTable, d, dictName, "R", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// BC, optional, array of numbers, len=0,1,3,4
	_, err = validateNumberArrayEntry(xRefTable, d, dictName, "BC", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// BG, optional, array of numbers between 0.0 and 0.1, len=0,1,3,4
	_, err = validateNumberArrayEntry(xRefTable, d, dictName, "BG", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// CA, optional, text string
	_, err = validateStringEntry(xRefTable, d, dictName, "CA", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// RC, optional, text string
	_, err = validateStringEntry(xRefTable, d, dictName, "RC", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// AC, optional, text string
	_, err = validateStringEntry(xRefTable, d, dictName, "AC", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// I, optional, stream dict
	_, err = validateStreamDictEntry(xRefTable, d, dictName, "I", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// RI, optional, stream dict
	_, err = validateStreamDictEntry(xRefTable, d, dictName, "RI", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// IX, optional, stream dict
	_, err = validateStreamDictEntry(xRefTable, d, dictName, "IX", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// IF, optional, icon fit dict,
	err = validateIconFitDictEntry(xRefTable, d, dictName, "IF", OPTIONAL, V10)
	if err != nil {
		return err
	}

	// TP, optional, integer 0..6
	_, err = validateIntegerEntry(xRefTable, d, dictName, "TP", OPTIONAL, V10, func(i int) bool { return 0 <= i && i <= 6 })

	return err
}

func validateAnnotationDictText(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.4

	// Open, optional, boolean
	_, err := validateBooleanEntry(xRefTable, dict, dictName, "Open", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// Name, optional, name
	_, err = validateNameEntry(xRefTable, dict, dictName, "Name", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// State, optional, text string, since V1.5
	validate := func(s string) bool { return memberOf(s, []string{"None", "Unmarked"}) }
	state, err := validateStringEntry(xRefTable, dict, dictName, "State", OPTIONAL, V15, validate)
	if err != nil {
		return err
	}

	// StateModel, text string, since V1.5
	validate = func(s string) bool { return memberOf(s, []string{"Marked", "Review"}) }
	_, err = validateStringEntry(xRefTable, dict, dictName, "StateModel", state != nil, V15, validate)

	return err
}

func validateActionOrDestination(xRefTable *XRefTable, dict *PDFDict, dictName string, sinceVersion PDFVersion) error {

	// The action that shall be performed when this item is activated.
	d, err := validateDictEntry(xRefTable, dict, dictName, "A", OPTIONAL, sinceVersion, nil)
	if err != nil {
		return err
	}
	if d != nil {
		return validateActionDict(xRefTable, d)
	}

	// A destination that shall be displayed when this item is activated.
	obj, err := validateEntry(xRefTable, dict, dictName, "Dest", OPTIONAL, sinceVersion)
	if err != nil || obj == nil {
		return err
	}

	return validateDestination(xRefTable, obj)
}

func validateURIActionDictEntry(xRefTable *XRefTable, dict *PDFDict, dictName, entryName string, required bool, sinceVersion PDFVersion) error {

	d, err := validateDictEntry(xRefTable, dict, dictName, entryName, required, sinceVersion, nil)
	if err != nil || d == nil {
		return err
	}

	dictName = "URIActionDict"

	// Type, optional, name
	_, err = validateNameEntry(xRefTable, d, dictName, "Type", OPTIONAL, V10, func(s string) bool { return s == "Action" })
	if err != nil {
		return err
	}

	// S, required, name, action Type
	_, err = validateNameEntry(xRefTable, d, dictName, "S", REQUIRED, V10, func(s string) bool { return s == "URI" })
	if err != nil {
		return err
	}

	return validateURIActionDict(xRefTable, d, dictName)
}

func validateAnnotationDictLink(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.5

	// A or Dest, required either or
	err := validateActionOrDestination(xRefTable, dict, dictName, V11)
	if err != nil {
		return err
	}

	// H, optional, name, since V1.2
	_, err = validateNameEntry(xRefTable, dict, dictName, "H", OPTIONAL, V12, nil)
	if err != nil {
		return err
	}

	// PA, optional, URI action dict, since V1.3
	err = validateURIActionDictEntry(xRefTable, dict, dictName, "PA", OPTIONAL, V13)
	if err != nil {
		return err
	}

	// QuadPoints, optional, number array, len=8, since V1.6
	_, err = validateNumberArrayEntry(xRefTable, dict, dictName, "QuadPoints", OPTIONAL, V16, func(a PDFArray) bool { return len(a) == 8 })
	if err != nil {
		return err
	}

	// BS, optional, border style dict, since V1.6
	sinceVersion := V16
	if xRefTable.ValidationMode == ValidationRelaxed {
		sinceVersion = V13
	}

	return validateBorderStyleDict(xRefTable, dict, dictName, "BS", OPTIONAL, sinceVersion)
}

func validateAnnotationDictFreeTextPart1(xRefTable *XRefTable, dict *PDFDict, dictName string, sinceVersion PDFVersion) error {

	// DA, required, string
	_, err := validateStringEntry(xRefTable, dict, dictName, "DA", REQUIRED, V10, nil)
	if err != nil {
		return err
	}

	// Q, optional, integer, since V1.4, 0,1,2
	sinceVersion = V14
	if xRefTable.ValidationMode == ValidationRelaxed {
		sinceVersion = V13
	}
	_, err = validateIntegerEntry(xRefTable, dict, dictName, "Q", OPTIONAL, sinceVersion, func(i int) bool { return 0 <= i && i <= 2 })
	if err != nil {
		return err
	}

	// RC, optional, text string or text stream, since V1.5
	sinceVersion = V15
	if xRefTable.ValidationMode == ValidationRelaxed {
		sinceVersion = V14
	}
	err = validateStringOrStreamEntry(xRefTable, dict, dictName, "RC", OPTIONAL, sinceVersion)
	if err != nil {
		return err
	}

	// DS, optional, text string, since V1.5
	_, err = validateStringEntry(xRefTable, dict, dictName, "DS", OPTIONAL, V15, nil)
	if err != nil {
		return err
	}

	// CL, optional, number array, since V1.6, len: 4 or 6
	sinceVersion = V16
	if xRefTable.ValidationMode == ValidationRelaxed {
		sinceVersion = V14
	}

	_, err = validateNumberArrayEntry(xRefTable, dict, dictName, "CL", OPTIONAL, sinceVersion, func(a PDFArray) bool { return len(a) == 4 || len(a) == 6 })

	return err
}

func validateAnnotationDictFreeTextPart2(xRefTable *XRefTable, dict *PDFDict, dictName string, sinceVersion PDFVersion) error {

	// IT, optional, name, since V1.6
	sinceVersion = V16
	if xRefTable.ValidationMode == ValidationRelaxed {
		sinceVersion = V14
	}
	validate := func(s string) bool {
		return memberOf(s, []string{"FreeText", "FreeTextCallout", "FreeTextTypeWriter", "FreeTextTypewriter"})
	}
	_, err := validateNameEntry(xRefTable, dict, dictName, "IT", OPTIONAL, sinceVersion, validate)
	if err != nil {
		return err
	}

	// BE, optional, border effect dict, since V1.6
	err = validateBorderEffectDictEntry(xRefTable, dict, dictName, "BE", OPTIONAL, V15)
	if err != nil {
		return err
	}

	// RD, optional, rectangle, since V1.6
	sinceVersion = V16
	if xRefTable.ValidationMode == ValidationRelaxed {
		sinceVersion = V14
	}
	_, err = validateRectangleEntry(xRefTable, dict, dictName, "RD", OPTIONAL, sinceVersion, nil)
	if err != nil {
		return err
	}

	// BS, optional, border style dict, since V1.6
	sinceVersion = V16
	if xRefTable.ValidationMode == ValidationRelaxed {
		sinceVersion = V13
	}
	err = validateBorderStyleDict(xRefTable, dict, dictName, "BS", OPTIONAL, sinceVersion)
	if err != nil {
		return err
	}

	// LE, optional, name, since V1.6
	sinceVersion = V16
	if xRefTable.ValidationMode == ValidationRelaxed {
		sinceVersion = V14
	}
	_, err = validateNameEntry(xRefTable, dict, dictName, "LE", OPTIONAL, sinceVersion, nil)

	return err
}

func validateAnnotationDictFreeText(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.6

	err := validateAnnotationDictFreeTextPart1(xRefTable, dict, dictName, V13)
	if err != nil {
		return err
	}

	return validateAnnotationDictFreeTextPart2(xRefTable, dict, dictName, V13)
}

func validateEntryMeasure(xRefTable *XRefTable, dict *PDFDict, dictName string, required bool, sinceVersion PDFVersion) error {

	d, err := validateDictEntry(xRefTable, dict, dictName, "Measure", required, sinceVersion, nil)
	if err != nil {
		return err
	}

	if d != nil {
		err = validateMeasureDict(xRefTable, d, sinceVersion)
		if err != nil {
			return err
		}
	}

	return nil
}

func validateCP(s string) bool { return s == "Inline" || s == "Top" }

func validateAnnotationDictLine(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.7

	// L, required, array of numbers, len:4
	_, err := validateNumberArrayEntry(xRefTable, dict, dictName, "L", REQUIRED, V10, func(a PDFArray) bool { return len(a) == 4 })
	if err != nil {
		return err
	}

	// BS, optional, border style dict
	err = validateBorderStyleDict(xRefTable, dict, dictName, "BS", OPTIONAL, V10)
	if err != nil {
		return err
	}

	// LE, optional, name array, since V1.4, len:2
	sinceVersion := V14
	if xRefTable.ValidationMode == ValidationRelaxed {
		sinceVersion = V13
	}
	_, err = validateNameArrayEntry(xRefTable, dict, dictName, "LE", OPTIONAL, sinceVersion, func(a PDFArray) bool { return len(a) == 2 })
	if err != nil {
		return err
	}

	// IC, optional, number array, since V1.4, len:0,1,3,4
	_, err = validateNumberArrayEntry(xRefTable, dict, dictName, "IC", OPTIONAL, sinceVersion, nil)
	if err != nil {
		return err
	}

	// LLE, optional, number, since V1.6, >0
	lle, err := validateNumberEntry(xRefTable, dict, dictName, "LLE", OPTIONAL, V16, func(f float64) bool { return f > 0 })
	if err != nil {
		return err
	}

	// LL, required if LLE present, number, since V1.6
	_, err = validateNumberEntry(xRefTable, dict, dictName, "LL", lle != nil, V16, nil)
	if err != nil {
		return err
	}

	// Cap, optional, bool, since V1.6
	_, err = validateBooleanEntry(xRefTable, dict, dictName, "Cap", OPTIONAL, V16, nil)
	if err != nil {
		return err
	}

	// IT, optional, name, since V1.6
	_, err = validateNameEntry(xRefTable, dict, dictName, "IT", OPTIONAL, V16, nil)
	if err != nil {
		return err
	}

	// LLO, optionl, number, since V1.7, >0
	_, err = validateNumberEntry(xRefTable, dict, dictName, "LLO", OPTIONAL, V17, func(f float64) bool { return f > 0 })
	if err != nil {
		return err
	}

	// CP, optional, name, since V1.7
	_, err = validateNameEntry(xRefTable, dict, dictName, "CP", OPTIONAL, V17, validateCP)
	if err != nil {
		return err
	}

	// Measure, optional, measure dict, since V1.7
	err = validateEntryMeasure(xRefTable, dict, dictName, OPTIONAL, V17)
	if err != nil {
		return err
	}

	// CO, optional, number array, since V1.7, len=2
	_, err = validateNumberArrayEntry(xRefTable, dict, dictName, "CO", OPTIONAL, V17, func(a PDFArray) bool { return len(a) == 2 })

	return err
}

func validateAnnotationDictCircleOrSquare(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.8

	// BS, optional, border style dict
	err := validateBorderStyleDict(xRefTable, dict, dictName, "BS", OPTIONAL, V10)
	if err != nil {
		return err
	}

	// IC, optional, array, since V1.4
	sinceVersion := V14
	if xRefTable.ValidationMode == ValidationRelaxed {
		sinceVersion = V13
	}
	_, err = validateNumberArrayEntry(xRefTable, dict, dictName, "IC", OPTIONAL, sinceVersion, nil)
	if err != nil {
		return err
	}

	// BE, optional, border effect dict, since V1.5
	err = validateBorderEffectDictEntry(xRefTable, dict, dictName, "BE", OPTIONAL, V15)
	if err != nil {
		return err
	}

	// RD, optional, rectangle, since V1.5
	_, err = validateRectangleEntry(xRefTable, dict, dictName, "RD", OPTIONAL, V15, nil)

	return err
}

func validateEntryIT(xRefTable *XRefTable, dict *PDFDict, dictName string, required bool, sinceVersion PDFVersion) error {

	// IT, optional, name, since V1.6
	validateIntent := func(s string) bool {

		if xRefTable.Version() == V16 {
			return s == "PolygonCloud"
		}

		if xRefTable.Version() == V17 {
			if memberOf(s, []string{"PolygonCloud", "PolyLineDimension", "PolygonDimension"}) {
				return true
			}
		}

		return false

	}

	_, err := validateNameEntry(xRefTable, dict, dictName, "IT", required, sinceVersion, validateIntent)

	return err
}

func validateAnnotationDictPolyLine(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.9

	// Vertices, required, array of numbers
	_, err := validateNumberArrayEntry(xRefTable, dict, dictName, "Vertices", REQUIRED, V10, nil)
	if err != nil {
		return err
	}

	// LE, optional, array of 2 names, meaningful only for polyline annotations.
	if dictName == "PolyLine" {
		_, err = validateNameArrayEntry(xRefTable, dict, dictName, "LE", OPTIONAL, V10, func(a PDFArray) bool { return len(a) == 2 })
		if err != nil {
			return err
		}
	}

	// BS, optional, border style dict
	err = validateBorderStyleDict(xRefTable, dict, dictName, "BS", OPTIONAL, V10)
	if err != nil {
		return err
	}

	// IC, optional, array of numbers [0.0 .. 1.0], len:1,3,4
	ensureArrayLength := func(a PDFArray, lengths ...int) bool {
		for _, length := range lengths {
			if len(a) == length {
				return true
			}
		}
		return false
	}
	_, err = validateNumberArrayEntry(xRefTable, dict, dictName, "IC", OPTIONAL, V14, func(a PDFArray) bool { return ensureArrayLength(a, 1, 3, 4) })
	if err != nil {
		return err
	}

	// BE, optional, border effect dict, meaningful only for polygon annotations
	if dictName == "Polygon" {
		err = validateBorderEffectDictEntry(xRefTable, dict, dictName, "BE", OPTIONAL, V10)
		if err != nil {
			return err
		}
	}

	return validateEntryIT(xRefTable, dict, dictName, OPTIONAL, V16)
}

func validateTextMarkupAnnotation(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.10

	// QuadPoints, required, number array, len:8
	_, err := validateNumberArrayEntry(xRefTable, dict, dictName, "QuadPoints", REQUIRED, V10, func(a PDFArray) bool { return len(a) == 8 })

	return err
}

func validateAnnotationDictStamp(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.12

	// Name, optional, name
	_, err := validateNameEntry(xRefTable, dict, dictName, "Name", OPTIONAL, V10, nil)

	return err
}

func validateAnnotationDictCaret(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.11

	// RD, optional, rectangle, since V1.5
	_, err := validateRectangleEntry(xRefTable, dict, dictName, "RD", OPTIONAL, V15, nil)
	if err != nil {
		return err
	}

	// Sy, optional, name
	_, err = validateNameEntry(xRefTable, dict, dictName, "Sy", OPTIONAL, V10, func(s string) bool { return s == "P" || s == "None" })

	return err
}

func validateAnnotationDictInk(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.13

	// InkList, required, array of stroked path arrays
	_, err := validateArrayArrayEntry(xRefTable, dict, dictName, "InkList", REQUIRED, V10, nil)
	if err != nil {
		return err
	}

	// BS, optional, border style dict
	err = validateBorderStyleDict(xRefTable, dict, dictName, "BS", OPTIONAL, V10)

	return err
}

func validateAnnotationDictPopup(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.14

	// Parent, optional, dict indRef
	indRef, err := validateIndRefEntry(xRefTable, dict, dictName, "Parent", OPTIONAL, V10)
	if err != nil {
		return err
	}
	if indRef != nil {
		d, err := xRefTable.DereferenceDict(*indRef)
		if err != nil || d == nil {
			return err
		}
	}

	// Open, optional, boolean
	_, err = validateBooleanEntry(xRefTable, dict, dictName, "Open", OPTIONAL, V10, nil)

	return err
}

func validateAnnotationDictFileAttachment(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.15

	// FS, required, file specification
	_, err := validateFileSpecEntry(xRefTable, dict, dictName, "FS", REQUIRED, V10)
	if err != nil {
		return err
	}

	// Name, optional, name
	_, err = validateNameEntry(xRefTable, dict, dictName, "Name", OPTIONAL, V10, nil)

	return err
}

func validateAnnotationDictSound(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.16

	// Sound, required, stream dict
	err := validateSoundDictEntry(xRefTable, dict, dictName, "Sound", REQUIRED, V10)
	if err != nil {
		return err
	}

	// Name, optional, name
	_, err = validateNameEntry(xRefTable, dict, dictName, "Name", OPTIONAL, V10, nil)

	return err
}

func validateMovieDict(xRefTable *XRefTable, dict *PDFDict) error {

	dictName := "movieDict"

	// F, required, file specification
	_, err := validateFileSpecEntry(xRefTable, dict, dictName, "F", REQUIRED, V10)
	if err != nil {
		return err
	}

	// Aspect, optional, integer array, length 2
	_, err = validateIntegerArrayEntry(xRefTable, dict, dictName, "Ascpect", OPTIONAL, V10, func(a PDFArray) bool { return len(a) == 2 })
	if err != nil {
		return err
	}

	// Rotate, optional, integer
	_, err = validateIntegerEntry(xRefTable, dict, dictName, "Rotate", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// Poster, optional boolean or stream
	return validateBooleanOrStreamEntry(xRefTable, dict, dictName, "Poster", OPTIONAL, V10)
}

func validateAnnotationDictMovie(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.17 Movie Annotations
	// 13.4 Movies
	// The features described in this sub-clause are obsolescent and their use is no longer recommended.
	// They are superseded by the general multimedia framework described in 13.2, “Multimedia.”

	// T, optional, text string
	_, err := validateStringEntry(xRefTable, dict, dictName, "T", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// Movie, required, movie dict
	d, err := validateDictEntry(xRefTable, dict, dictName, "Movie", REQUIRED, V10, nil)
	if err != nil {
		return err
	}

	err = validateMovieDict(xRefTable, d)
	if err != nil {
		return err
	}

	// A, optional, boolean or movie activation dict
	obj, found := dict.Find("A")

	if found {

		obj, err = xRefTable.Dereference(obj)
		if err != nil {
			return err
		}

		if obj != nil {
			switch obj := obj.(type) {
			case PDFBoolean:
				// no further processing

			case PDFDict:
				err = validateMovieActivationDict(xRefTable, &obj)
				if err != nil {
					return err
				}
			}
		}

	}

	return nil
}

func validateAnnotationDictWidget(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.19

	// H, optional, name
	validate := func(s string) bool { return memberOf(s, []string{"N", "I", "O", "P", "T", "A"}) }
	_, err := validateNameEntry(xRefTable, dict, dictName, "H", OPTIONAL, V10, validate)
	if err != nil {
		return err
	}

	// MK, optional, dict
	// An appearance characteristics dictionary that shall be used in constructing
	// a dynamic appearance stream specifying the annotation’s visual presentation on the page.dict
	err = validateAppearanceCharacteristicsDictEntry(xRefTable, dict, dictName, "MK", OPTIONAL, V10)
	if err != nil {
		return err
	}

	// A, optional, dict, since V1.1
	// An action that shall be performed when the annotation is activated.
	d, err := validateDictEntry(xRefTable, dict, dictName, "A", OPTIONAL, V11, nil)
	if err != nil {
		return err
	}
	if d != nil {
		err = validateActionDict(xRefTable, d)
		if err != nil {
			return err
		}
	}

	// AA, optional, dict, since V1.2
	// An additional-actions dictionary defining the annotation’s behaviour in response to various trigger events.
	err = validateAdditionalActions(xRefTable, dict, dictName, "AA", OPTIONAL, V12, "fieldOrAnnot")
	if err != nil {
		return err
	}

	// BS, optional, border style dict, since V1.2
	// A border style dictionary specifying the width and dash pattern
	// that shall be used in drawing the annotation’s border.
	validateBorderStyleDict(xRefTable, dict, dictName, "BS", OPTIONAL, V12)
	if err != nil {
		return err
	}

	// Parent, dict, required if one of multiple children in a field.
	// An indirect reference to the widget annotation’s parent field.
	_, err = validateIndRefEntry(xRefTable, dict, dictName, "Parent", OPTIONAL, V10)

	return err
}

func validateAnnotationDictScreen(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.18

	// T, optional, name
	_, err := validateNameEntry(xRefTable, dict, dictName, "T", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// MK, optional, appearance characteristics dict
	err = validateAppearanceCharacteristicsDictEntry(xRefTable, dict, dictName, "MK", OPTIONAL, V10)
	if err != nil {
		return err
	}

	// A, optional, action dict, since V1.0
	d, err := validateDictEntry(xRefTable, dict, dictName, "A", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}
	if d != nil {
		err = validateActionDict(xRefTable, d)
		if err != nil {
			return err
		}
	}

	// AA, optional, additional-actions dict, since V1.2
	return validateAdditionalActions(xRefTable, dict, dictName, "AA", OPTIONAL, V12, "fieldOrAnnot")
}

func validateAnnotationDictPrinterMark(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.20

	// MN, optional, name
	_, err := validateNameEntry(xRefTable, dict, dictName, "MN", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// F, required integer, since V1.1, annotation flags
	_, err = validateIntegerEntry(xRefTable, dict, dictName, "F", REQUIRED, V11, nil)
	if err != nil {
		return err
	}

	// AP, required, appearance dict, since V1.2
	return validateAppearDictEntry(xRefTable, dict, dictName, REQUIRED, V12)
}

func validateAnnotationDictTrapNet(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.21

	// LastModified, optional, date
	_, err := validateDateEntry(xRefTable, dict, dictName, "LastModified", OPTIONAL, V10)
	if err != nil {
		return err
	}

	// Version, optional, array
	_, err = validateArrayEntry(xRefTable, dict, dictName, "Version", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// AnnotStates, optional, array of names
	_, err = validateNameArrayEntry(xRefTable, dict, dictName, "AnnotStates", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// FontFauxing, optional, font dict array
	validateFontDictArray := func(a PDFArray) bool {

		var retValue bool

		for _, v := range a {

			if v == nil {
				continue
			}

			dict, err := xRefTable.DereferenceDict(v)
			if err != nil {
				return false
			}

			if dict == nil {
				continue
			}

			if dict.Type() == nil || *dict.Type() != "Font" {
				return false
			}

			retValue = true

		}

		return retValue
	}

	_, err = validateArrayEntry(xRefTable, dict, dictName, "FontFauxing", OPTIONAL, V10, validateFontDictArray)
	if err != nil {
		return err
	}

	_, err = validateIntegerEntry(xRefTable, dict, dictName, "F", REQUIRED, V11, nil)

	return err
}

func validateAnnotationDictWatermark(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.22

	// FixedPrint, optional, dict

	validateFixedPrintDict := func(dict PDFDict) bool {

		dictName := "fixedPrintDict"

		// Type, required, name
		_, err := validateNameEntry(xRefTable, &dict, dictName, "Type", REQUIRED, V10, func(s string) bool { return s == "FixedPrint" })
		if err != nil {
			return false
		}

		// Matrix, optional, integer array, length = 6
		_, err = validateIntegerArrayEntry(xRefTable, &dict, dictName, "Matrix", OPTIONAL, V10, func(a PDFArray) bool { return len(a) == 6 })
		if err != nil {
			return false
		}

		// H, optional, number
		_, err = validateNumberEntry(xRefTable, &dict, dictName, "H", OPTIONAL, V10, nil)
		if err != nil {
			return false
		}

		// V, optional, number
		_, err = validateNumberEntry(xRefTable, &dict, dictName, "V", OPTIONAL, V10, nil)
		if err != nil {
			return false
		}

		return true
	}

	_, err := validateDictEntry(xRefTable, dict, dictName, "FixedPrint", OPTIONAL, V10, validateFixedPrintDict)

	return err
}

func validateAnnotationDict3D(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 13.6.2

	// AP with entry N, required

	// 3DD, required, 3D stream or 3D reference dict
	err := validateStreamDictOrDictEntry(xRefTable, dict, dictName, "3DD", REQUIRED, V16)
	if err != nil {
		return err
	}

	// 3DV, optional, various
	_, err = validateEntry(xRefTable, dict, dictName, "3DV", OPTIONAL, V16)
	if err != nil {
		return err
	}

	// 3DA, optional, activation dict
	_, err = validateDictEntry(xRefTable, dict, dictName, "3DA", OPTIONAL, V16, nil)
	if err != nil {
		return err
	}

	// 3DI, optional, boolean
	_, err = validateBooleanEntry(xRefTable, dict, dictName, "3DI", OPTIONAL, V16, nil)

	return err
}

func validateEntryIC(xRefTable *XRefTable, dict *PDFDict, dictName string, required bool, sinceVersion PDFVersion) error {

	// IC, optional, number array, length:3 [0.0 .. 1.0]
	validateICArray := func(arr PDFArray) bool {

		if len(arr) != 3 {
			return false
		}

		for _, v := range arr {

			n, err := xRefTable.Dereference(v)
			if err != nil {
				return false
			}

			switch n := n.(type) {
			case PDFInteger:
				if n < 0 || n > 1 {
					return false
				}

			case PDFFloat:
				if n < 0.0 || n > 1.0 {
					return false
				}
			}
		}

		return true
	}

	_, err := validateNumberArrayEntry(xRefTable, dict, dictName, "IC", required, sinceVersion, validateICArray)

	return err
}

func validateAnnotationDictRedact(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// see 12.5.6.23

	// QuadPoints, optional, number array
	_, err := validateNumberArrayEntry(xRefTable, dict, dictName, "QuadPoints", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// IC, optional, number array, length:3 [0.0 .. 1.0]
	err = validateEntryIC(xRefTable, dict, dictName, OPTIONAL, V10)
	if err != nil {
		return err
	}

	// RO, optional, stream
	_, err = validateStreamDictEntry(xRefTable, dict, dictName, "RO", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// OverlayText, optional, text string
	_, err = validateStringEntry(xRefTable, dict, dictName, "OverlayText", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// Repeat, optional, boolean
	_, err = validateBooleanEntry(xRefTable, dict, dictName, "Repeat", OPTIONAL, V10, nil)
	if err != nil {
		return err
	}

	// DA, required, byte string
	_, err = validateStringEntry(xRefTable, dict, dictName, "DA", REQUIRED, V10, nil)
	if err != nil {
		return err
	}

	// Q, optional, integer
	_, err = validateIntegerEntry(xRefTable, dict, dictName, "Q", OPTIONAL, V10, nil)

	return err
}

func validateExDataDict(xRefTable *XRefTable, dict *PDFDict) error {

	dictName := "ExData"

	_, err := validateNameEntry(xRefTable, dict, dictName, "Type", OPTIONAL, V10, func(s string) bool { return s == "ExData" })
	if err != nil {
		return err
	}

	_, err = validateNameEntry(xRefTable, dict, dictName, "Subtype", REQUIRED, V10, func(s string) bool { return s == "Markup3D" })

	return err
}

func validatePopupEntry(xRefTable *XRefTable, dict *PDFDict, dictName, entryName string, required bool, sinceVersion PDFVersion) error {

	d, err := validateDictEntry(xRefTable, dict, dictName, entryName, required, sinceVersion, nil)
	if err != nil {
		return err
	}

	if d != nil {

		_, err = validateNameEntry(xRefTable, d, dictName, "Subtype", REQUIRED, V10, func(s string) bool { return s == "Popup" })
		if err != nil {
			return err
		}

		_, err = validateAnnotationDict(xRefTable, d)
		if err != nil {
			return err
		}

	}

	return nil
}

func validateIRTEntry(xRefTable *XRefTable, dict *PDFDict, dictName, entryName string, required bool, sinceVersion PDFVersion) error {

	d, err := validateDictEntry(xRefTable, dict, dictName, entryName, required, sinceVersion, nil)
	if err != nil {
		return err
	}

	if d != nil {
		_, err = validateAnnotationDict(xRefTable, d)
		if err != nil {
			return err
		}
	}

	return nil
}

func validateMarkupAnnotation(xRefTable *XRefTable, dict *PDFDict) error {

	dictName := "markupAnnot"

	// T, optional, text string, since V1.1
	_, err := validateStringEntry(xRefTable, dict, dictName, "T", OPTIONAL, V11, nil)
	if err != nil {
		return err
	}

	// Popup, optional, dict, since V1.3
	err = validatePopupEntry(xRefTable, dict, dictName, "Popup", OPTIONAL, V13)
	if err != nil {
		return err
	}

	// CA, optional, number, since V1.4
	_, err = validateNumberEntry(xRefTable, dict, dictName, "CA", OPTIONAL, V14, nil)
	if err != nil {
		return err
	}

	// RC, optional, text string or stream, since V1.5
	err = validateStringOrStreamEntry(xRefTable, dict, dictName, "RC", OPTIONAL, V15)
	if err != nil {
		return err
	}

	// CreationDate, optional, date, since V1.5
	_, err = validateDateEntry(xRefTable, dict, dictName, "CreationDate", OPTIONAL, V15)
	if err != nil {
		return err
	}

	// IRT, optional, (in reply to) dict, since V1.5
	err = validateIRTEntry(xRefTable, dict, dictName, "IRT", OPTIONAL, V15)
	if err != nil {
		return err
	}

	// Subj, optional, text string, since V1.5
	sinceVersion := V15
	if xRefTable.ValidationMode == ValidationRelaxed {
		sinceVersion = V14
	}
	_, err = validateStringEntry(xRefTable, dict, dictName, "Subj", OPTIONAL, sinceVersion, nil)
	if err != nil {
		return err
	}

	// RT, optional, name, since V1.6
	validate := func(s string) bool { return s == "R" || s == "Group" }
	_, err = validateNameEntry(xRefTable, dict, dictName, "RT", OPTIONAL, V16, validate)
	if err != nil {
		return err
	}

	// IT, optional, name, since V1.6
	_, err = validateNameEntry(xRefTable, dict, dictName, "IT", OPTIONAL, V16, nil)
	if err != nil {
		return err
	}

	// ExData, optional, dict, since V1.7
	d, err := validateDictEntry(xRefTable, dict, dictName, "ExData", OPTIONAL, V17, nil)
	if err != nil {
		return err
	}
	if d != nil {
		err = validateExDataDict(xRefTable, d)
		if err != nil {
			return err
		}
	}

	return nil
}

func validateEntryP(xRefTable *XRefTable, dict *PDFDict, dictName string, required bool, sinceVersion PDFVersion) error {

	indRef, err := validateIndRefEntry(xRefTable, dict, dictName, "P", required, sinceVersion)
	if err != nil || indRef == nil {
		return err
	}

	// check if this indRef points to a pageDict.

	d, err := xRefTable.DereferenceDict(*indRef)
	if err != nil {
		return err
	}

	if d == nil {
		return errors.Errorf("validateEntryP: entry \"P\" (obj#%d) is nil", indRef.ObjectNumber)
	}

	_, err = validateNameEntry(xRefTable, d, "pageDict", "Type", REQUIRED, V10, func(s string) bool { return s == "Page" })

	return err
}

func validateAppearDictEntry(xRefTable *XRefTable, dict *PDFDict, dictName string, required bool, sinceVersion PDFVersion) error {

	d, err := validateDictEntry(xRefTable, dict, dictName, "AP", required, sinceVersion, nil)
	if err != nil {
		return err
	}

	if d != nil {
		err = validateAppearanceDict(xRefTable, *d)
	}

	return err
}

func validateBorderArrayLength(a PDFArray) bool {
	return len(a) == 3 || len(a) == 4
}

func validateAnnotationDictGeneral(xRefTable *XRefTable, dict *PDFDict, dictName string) (*PDFName, error) {

	// Type, optional, name
	_, err := validateNameEntry(xRefTable, dict, dictName, "Type", OPTIONAL, V10, func(s string) bool { return s == "Annot" })
	if err != nil {
		return nil, err
	}

	// Subtype, required, name
	subtype, err := validateNameEntry(xRefTable, dict, dictName, "Subtype", REQUIRED, V10, nil)
	if err != nil {
		return nil, err
	}

	// Rect, required, rectangle
	_, err = validateRectangleEntry(xRefTable, dict, dictName, "Rect", REQUIRED, V10, nil)
	if err != nil {
		return nil, err
	}

	// Contents, optional, text string
	_, err = validateStringEntry(xRefTable, dict, dictName, "Contents", OPTIONAL, V10, nil)
	if err != nil {
		return nil, err
	}

	// P, optional, indRef of page dict
	err = validateEntryP(xRefTable, dict, dictName, OPTIONAL, V10)
	if err != nil {
		return nil, err
	}

	// NM, optional, text string, since V1.4
	_, err = validateStringEntry(xRefTable, dict, dictName, "NM", OPTIONAL, V14, nil)
	if err != nil {
		return nil, err
	}

	// M, optional, date string in any format, since V1.1
	_, err = validateStringEntry(xRefTable, dict, dictName, "M", OPTIONAL, V11, nil)
	if err != nil {
		return nil, err
	}

	// F, optional integer, since V1.1, annotation flags
	_, err = validateIntegerEntry(xRefTable, dict, dictName, "F", OPTIONAL, V11, nil)
	if err != nil {
		return nil, err
	}

	// AP, optional, appearance dict, since V1.2
	err = validateAppearDictEntry(xRefTable, dict, dictName, OPTIONAL, V12)
	if err != nil {
		return nil, err
	}

	// AS, optional, name, since V1.2
	_, err = validateNameEntry(xRefTable, dict, dictName, "AS", OPTIONAL, V11, nil)
	if err != nil {
		return nil, err
	}

	// Border, optional, array of numbers
	//v := func(a PDFArray) bool { return len(a) == 3 || len(a) == 4 }
	_, err = validateNumberArrayEntry(xRefTable, dict, dictName, "Border", OPTIONAL, V10, validateBorderArrayLength)
	if err != nil {
		return nil, err
	}

	// C, optional array, of numbers, since V1.1
	_, err = validateNumberArrayEntry(xRefTable, dict, dictName, "C", OPTIONAL, V11, nil)
	if err != nil {
		return nil, err
	}

	// StructParent, optional, integer, since V1.3
	_, err = validateIntegerEntry(xRefTable, dict, dictName, "StructParent", OPTIONAL, V13, nil)
	if err != nil {
		return nil, err
	}

	// OC, optional, content group dict or content membership dict, since V1.5
	// Specifying the optional content properties for the annotation.
	err = validateOptionalContent(xRefTable, dict, dictName, "OC", OPTIONAL, V15)
	if err != nil {
		return nil, err
	}

	return subtype, nil
}

func validateAnnotationDictConcrete(xRefTable *XRefTable, dict *PDFDict, subtype PDFName) error {

	// see table 169

	for k, v := range map[string]struct {
		validate     func(xRefTable *XRefTable, dict *PDFDict, dictName string) error
		sinceVersion PDFVersion
		markup       bool
	}{
		"Text":           {validateAnnotationDictText, V10, true},
		"Link":           {validateAnnotationDictLink, V10, false},
		"FreeText":       {validateAnnotationDictFreeText, V13, true},
		"Line":           {validateAnnotationDictLine, V13, true},
		"Polygon":        {validateAnnotationDictPolyLine, V15, true},
		"PolyLine":       {validateAnnotationDictPolyLine, V15, true},
		"Highlight":      {validateTextMarkupAnnotation, V13, true},
		"Underline":      {validateTextMarkupAnnotation, V13, true},
		"Squiggly":       {validateTextMarkupAnnotation, V14, true},
		"StrikeOut":      {validateTextMarkupAnnotation, V13, true},
		"Square":         {validateAnnotationDictCircleOrSquare, V13, true},
		"Circle":         {validateAnnotationDictCircleOrSquare, V13, true},
		"Stamp":          {validateAnnotationDictStamp, V13, true},
		"Caret":          {validateAnnotationDictCaret, V15, true},
		"Ink":            {validateAnnotationDictInk, V13, true},
		"Popup":          {validateAnnotationDictPopup, V13, false},
		"FileAttachment": {validateAnnotationDictFileAttachment, V13, true},
		"Sound":          {validateAnnotationDictSound, V12, true},
		"Movie":          {validateAnnotationDictMovie, V12, false},
		"Widget":         {validateAnnotationDictWidget, V12, false},
		"Screen":         {validateAnnotationDictScreen, V15, false},
		"PrinterMark":    {validateAnnotationDictPrinterMark, V14, false},
		"TrapNet":        {validateAnnotationDictTrapNet, V13, false},
		"Watermark":      {validateAnnotationDictWatermark, V16, false},
		"3D":             {validateAnnotationDict3D, V16, false},
		"Redact":         {validateAnnotationDictRedact, V17, true},
	} {
		if subtype.Value() == k {

			err := xRefTable.ValidateVersion(k, v.sinceVersion)
			if err != nil {
				return err
			}

			if v.markup {
				err := validateMarkupAnnotation(xRefTable, dict)
				if err != nil {
					return err
				}
			}

			return v.validate(xRefTable, dict, k)
		}
	}

	return errors.Errorf("validateAnnotationDictConcrete: unsupported annotation subtype:%s\n", subtype)
}

func validateAnnotationDictSpecial(xRefTable *XRefTable, dict *PDFDict, dictName string) error {

	// AAPL:AKExtras
	// No documentation for this PDF-Extension - this is a speculative implementation.
	return validateAAPLAKExtrasDictEntry(xRefTable, dict, dictName, "AAPL:AKExtras", OPTIONAL, V10)
}

func validateAnnotationDict(xRefTable *XRefTable, dict *PDFDict) (isTrapNet bool, err error) {

	dictName := "annotDict"

	subtype, err := validateAnnotationDictGeneral(xRefTable, dict, dictName)
	if err != nil {
		return false, err
	}

	err = validateAnnotationDictConcrete(xRefTable, dict, *subtype)
	if err != nil {
		return false, err
	}

	err = validateAnnotationDictSpecial(xRefTable, dict, dictName)
	if err != nil {
		return false, err
	}

	return *subtype == "TrapNet", nil
}

func validatePageAnnotations(xRefTable *XRefTable, dict *PDFDict) error {

	arr, err := validateArrayEntry(xRefTable, dict, "pageDict", "Annots", OPTIONAL, V10, nil)
	if err != nil || arr == nil {
		return err
	}

	// array of indrefs to annotation dicts.
	var annotsDict PDFDict

	// an optional TrapNetAnnotation has to be the final entry in this list.
	hasTrapNet := false

	for _, v := range *arr {

		if hasTrapNet {
			return errors.New("validatePageAnnotations: corrupted page annotation list, \"TrapNet\" has to be the last entry")
		}

		if indRef, ok := v.(PDFIndirectRef); ok {

			log.Debug.Printf("processing annotDict %d\n", indRef.ObjectNumber)

			annotsDictp, err := xRefTable.DereferenceDict(indRef)
			if err != nil || annotsDictp == nil {
				return errors.New("validatePageAnnotations: corrupted annotation dict")
			}

			annotsDict = *annotsDictp

		} else if annotsDict, ok = v.(PDFDict); !ok {
			return errors.New("validatePageAnnotations: corrupted array of indrefs")
		}

		hasTrapNet, err = validateAnnotationDict(xRefTable, &annotsDict)
		if err != nil {
			return err
		}

	}

	return nil
}

func validatePagesAnnotations(xRefTable *XRefTable, dict *PDFDict) error {

	// Get number of pages of this PDF file.
	pageCount := dict.IntEntry("Count")
	if pageCount == nil {
		return errors.New("validatePagesAnnotations: missing \"Count\"")
	}

	log.Debug.Printf("validatePagesAnnotations: This page node has %d pages\n", *pageCount)

	// Iterate over page tree.
	kidsArray := dict.PDFArrayEntry("Kids")

	for _, v := range *kidsArray {

		if v == nil {
			log.Debug.Println("validatePagesAnnotations: kid is nil")
			continue
		}

		d, err := xRefTable.DereferenceDict(v)
		if err != nil {
			return err
		}
		if d == nil {
			return errors.New("validatePagesAnnotations: pageNodeDict is null")
		}

		dictType := d.Type()
		if dictType == nil {
			return errors.New("validatePagesAnnotations: missing pageNodeDict type")
		}

		switch *dictType {

		case "Pages":
			// Recurse over pagetree
			err = validatePagesAnnotations(xRefTable, d)
			if err != nil {
				return err
			}

		case "Page":
			err = validatePageAnnotations(xRefTable, d)
			if err != nil {
				return err
			}

		default:
			return errors.Errorf("validatePagesAnnotations: expected dict type: %s\n", *dictType)

		}

	}

	return nil
}
