package main

/*
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdarg.h>

typedef char* String;
typedef char** StringArray;
typedef int* IntArray;
typedef double* DoubleArray;

static String Concat(String first, ...) {
    va_list args;
    String token;
    size_t total_len = 0;
    va_start(args, first);
    token = first;
    while(token != NULL) {
        total_len += strlen(token);
        token = va_arg(args, String);
    }
    va_end(args);
    String result = malloc(total_len + 1);
    if(result == NULL) return NULL;
    result[0] = '\0';
    va_start(args, first);
    token = first;
    while(token != NULL) {
        strcat(result, token);
        token = va_arg(args, String);
    }
    va_end(args);    
    return result;
}*/
import "C"
import (
	"strconv"
	"strings"
	"unsafe"
)

//export Atoi
func Atoi(str C.String) C.int {
	goStr := C.GoString(str)
	val, err := strconv.Atoi(goStr)
	if err != nil {
		return 0
	}
	return C.int(val)
}

//export Atof
func Atof(str C.String) C.double {
	goStr := C.GoString(str)
	val, err := strconv.ParseFloat(goStr, 64)
	if err != nil {
		return 0.0
	}
	return C.double(val)
}

//export Itoa
func Itoa(n C.int) C.String {
	return C.CString(strconv.Itoa(int(n)))
}

//export Ftoa
func Ftoa(flo C.double, precision C.int) C.String {
	return C.CString(strconv.FormatFloat(float64(flo), 'f', int(precision), 64))
}

//export ParseBool
func ParseBool(str C.String) C.int {
	goStr := C.GoString(str)
	val, err := strconv.ParseBool(goStr)
	if err != nil {
		return 0
	}
	if val {
		return 1
	}
	return 0
}

//export NewString
func NewString(size C.int) *C.char {
    if size <= 0 {
        return nil
    }
    ptr := C.malloc(C.size_t(size + 1))
    if ptr == nil {
        return nil
    }
    C.memset(ptr, 0, C.size_t(size + 1))
    return (*C.char)(ptr)
}

//export GetStringSize
func GetStringSize(str C.String) C.int {
	return C.int(len(C.GoString(str)))
}

//export Substring
func Substring(str C.String, start, end C.int) C.String {
	goStr := C.GoString(str)
	goStart := int(start)
	goEnd := int(end)
	
	if goStart < 0 || goEnd > len(goStr) || goStart > goEnd {
		return C.CString("")
	}
	
	return C.CString(goStr[goStart:goEnd])
}

//export IsNumeric
func IsNumeric(str C.String) C.int {
	goStr := C.GoString(str)
	_, err := strconv.ParseFloat(goStr, 64)
	if err != nil {
		return 0
	}
	return 1
}

//export ConcatAll
func ConcatAll(strs **C.char) *C.char {
    if strs == nil {
        return C.CString("")
    }
    var goStrings []string
    for i := 0; ; i++ {
        ptr := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(strs)) + uintptr(i)*unsafe.Sizeof(strs)))
        if *ptr == nil {
            break
        }
        goStrings = append(goStrings, C.GoString(*ptr))
    }
    result := strings.Join(goStrings, "")
    return C.CString(result)
}

//export ToUpperCase
func ToUpperCase(str C.String) C.String {
	goStr := C.GoString(str)
	return C.CString(strings.ToUpper(goStr))
}

//export ToLowerCase
func ToLowerCase(str C.String) C.String {
	goStr := C.GoString(str)
	return C.CString(strings.ToLower(goStr))
}

//export Trim
func Trim(str C.String) C.String {
	goStr := C.GoString(str)
	return C.CString(strings.TrimSpace(goStr))
}

//export ReplaceAll
func ReplaceAll(str, old, new C.String) C.String {
	goStr := C.GoString(str)
	goOld := C.GoString(old)
	goNew := C.GoString(new)
	return C.CString(strings.ReplaceAll(goStr, goOld, goNew))
}

//export Equals
func Equals(str1 *C.char, str2 *C.char) C.int {
    goStr1 := C.GoString(str1)
    goStr2 := C.GoString(str2)
    if goStr1 == goStr2 {
        return C.int(1)  // true
    }
    return C.int(0)       // false
}

//export Contains
func Contains(str *C.char, substr *C.char) C.int {
    goStr := C.GoString(str)
    goSubstr := C.GoString(substr)
    if strings.Contains(goStr, goSubstr) {
        return C.int(1)  // true
    }
    return C.int(0)      // false
}

//export NewStringArray
func NewStringArray(size C.int) **C.char {
    if size <= 0 {
        return nil
    }
    elementSize := unsafe.Sizeof((*C.char)(nil))
    totalSize := uintptr(size+1) * elementSize
    ptr := C.malloc(C.size_t(totalSize))
    if ptr == nil {
        return nil
    }
    C.memset(ptr, 0, C.size_t(totalSize))
    return (**C.char)(ptr)
}

//export Split
func Split(str *C.char, sep *C.char) **C.char {
    goStr := C.GoString(str)
    goSep := C.GoString(sep)
    parts := strings.Split(goStr, goSep)
    arr := (**C.char)(C.malloc(C.size_t(len(parts)+1) * C.size_t(unsafe.Sizeof((*C.char)(nil)))))
    if arr == nil {
        return nil
    }
    C.memset(unsafe.Pointer(arr), 0, C.size_t(len(parts)+1)*C.size_t(unsafe.Sizeof((*C.char)(nil))))
    for i, part := range parts {
        cStr := C.CString(part)
        ptr := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(arr)) + uintptr(i)*unsafe.Sizeof((*C.char)(nil))))
        *ptr = cStr
    }    
    return arr
}

//export GetStringArraySize
func GetStringArraySize(strs **C.char) C.int {
    var count C.int = 0
    for {
        ptr := *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(strs)) + uintptr(count)*unsafe.Sizeof(*strs)))
        if ptr == nil {
            break
        }
        count++
    }
    return count
}

//export NewIntArray
func NewIntArray(size C.int) *C.int {
    if size <= 0 {
        return nil
    }
    ptr := C.malloc(C.size_t(size) * C.size_t(unsafe.Sizeof(C.int(0))))
    if ptr == nil {
        return nil
    }
    C.memset(ptr, 0, C.size_t(size)*C.size_t(unsafe.Sizeof(C.int(0))))    
    return (*C.int)(ptr)
}

//export NewDoubleArray
func NewDoubleArray(size C.int) *C.double {
    if size <= 0 {
        return nil
    }
    ptr := C.malloc(C.size_t(size) * C.size_t(unsafe.Sizeof(C.double(0))))
    if ptr == nil {
        return nil
    }
    C.memset(ptr, 0, C.size_t(size)*C.size_t(unsafe.Sizeof(C.double(0))))
    
    return (*C.double)(ptr)
}

//export FreeIntArray
func FreeIntArray(ints *C.int) {
    if ints != nil {
        C.free(unsafe.Pointer(ints))
    }
}

//export FreeDoubleArray
func FreeDoubleArray(flos *C.double) {
    if flos != nil {
        C.free(unsafe.Pointer(flos))
    }
}

//export FreeStringArray
func FreeStringArray(strs **C.char) {
    if strs == nil {
        return
    }
    for i := 0; ; i++ {
        ptr := *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(strs)) + uintptr(i)*unsafe.Sizeof(*strs)))
        if ptr == nil {
            break
        }
        C.free(unsafe.Pointer(ptr))
    }
    C.free(unsafe.Pointer(strs))
}

//export FreeString
func FreeString(str C.String) {
    C.free(unsafe.Pointer(str))
}

func main() {}
