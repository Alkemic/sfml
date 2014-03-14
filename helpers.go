// Copyright (C) 2012-2014 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

/*
#include <SFML/System.h>
#include <string.h>

void incPtr(sfUint32** ptr)  { ++(*ptr); }
void copyData(void* source, void* dest, size_t size) { memcpy(dest,source,size); }
size_t sizeofInt16() { return sizeof(sfInt16); }
*/
import "C"

import "errors"

//As SFML does not provide useful errors we just return a generic error message
var genericError = errors.New("Error: See stderr for more details")

/////////////////////////////////////
///		WRAPPING HELPERS
/////////////////////////////////////

func sfBool2Go(b C.sfBool) bool {
	return b == 1
}

func goBool2C(b bool) C.sfBool {
	if b {
		return C.sfBool(1)
	}
	return C.sfBool(0)
}

func utf32CString2Go(cstr *C.sfUint32) string {
	var str string

	for ptr := cstr; *ptr != 0; C.incPtr(&ptr) {
		str += string(rune(uint32(*ptr)))
	}

	return str
}

func strToRunes(str string) []rune {
	return append([]rune(str), rune(0))
}
