package main

/*
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdarg.h>

typedef struct {
    char** data;
    int count;
} StringArray;

static char* Concat(char *first, ...) {
    va_list args;
    char *token;
    size_t total_len = 0;
    
    // 1. Calcular longitud total necesaria
    va_start(args, first);
    token = first;
    while(token != NULL) {
        total_len += strlen(token);
        token = va_arg(args, char*);
    }
    va_end(args);
    
    // 2. Reservar memoria (incluyendo espacio para \0)
    char *result = malloc(total_len + 1);
    if(result == NULL) return NULL;
    result[0] = '\0';
    
    // 3. Concatenar todos los strings
    va_start(args, first);
    token = first;
    while(token != NULL) {
        strcat(result, token);
        token = va_arg(args, char*);
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


// ========== Funciones de Conversión ==========

//export Atoi
func Atoi(s *C.char) C.int {
	goStr := C.GoString(s)
	val, err := strconv.Atoi(goStr)
	if err != nil {
		return 0
	}
	return C.int(val)
}

//export Atof
func Atof(s *C.char) C.double {
	goStr := C.GoString(s)
	val, err := strconv.ParseFloat(goStr, 64)
	if err != nil {
		return 0.0
	}
	return C.double(val)
}

//export Itoa
func Itoa(n C.int) *C.char {
	return C.CString(strconv.Itoa(int(n)))
}

//export Ftoa
func Ftoa(f C.double, precision C.int) *C.char {
	return C.CString(strconv.FormatFloat(float64(f), 'f', int(precision), 64))
}

//export ParseBool
func ParseBool(s *C.char) C.int {
	goStr := C.GoString(s)
	val, err := strconv.ParseBool(goStr)
	if err != nil {
		return 0
	}
	if val {
		return 1
	}
	return 0
}

// ========== Funciones String ==========

//export StrLen
func StrLen(s *C.char) C.int {
	return C.int(len(C.GoString(s)))
}

//export Substring
func Substring(s *C.char, start, end C.int) *C.char {
	goStr := C.GoString(s)
	goStart := int(start)
	goEnd := int(end)
	
	if goStart < 0 || goEnd > len(goStr) || goStart > goEnd {
		return C.CString("")
	}
	
	return C.CString(goStr[goStart:goEnd])
}

//export IsNumeric
func IsNumeric(s *C.char) C.int {
	goStr := C.GoString(s)
	_, err := strconv.ParseFloat(goStr, 64)
	if err != nil {
		return 0
	}
	return 1
}

//export ConcatAll
func ConcatAll(strs **C.char, count C.int) *C.char {
	// Convertir el array de C a slice de Go
	length := int(count)
	tmpslice := (*[1 << 30]*C.char)(unsafe.Pointer(strs))[:length:length]
	
	goStrs := make([]string, length)
	for i, s := range tmpslice {
		goStrs[i] = C.GoString(s)
	}
	
	result := strings.Join(goStrs, "")
	return C.CString(result)
}

//export ToUpperCase
func ToUpperCase(s *C.char) *C.char {
	goStr := C.GoString(s)
	return C.CString(strings.ToUpper(goStr))
}

//export ToLowerCase
func ToLowerCase(s *C.char) *C.char {
	goStr := C.GoString(s)
	return C.CString(strings.ToLower(goStr))
}

//export Trim
func Trim(s *C.char) *C.char {
	goStr := C.GoString(s)
	return C.CString(strings.TrimSpace(goStr))
}

//export ReplaceAll
func ReplaceAll(s, old, new *C.char) *C.char {
	goStr := C.GoString(s)
	goOld := C.GoString(old)
	goNew := C.GoString(new)
	return C.CString(strings.ReplaceAll(goStr, goOld, goNew))
}

// ========== Funciones StringArray ==========

//export NewStringArray
func NewStringArray(size C.int) C.StringArray {
    goSize := int(size)
    
    // Allocate memory for the char* array
    cArray := make([]*C.char, goSize)
    
    // Convert Go slice to C array
    cArrayPtr := (**C.char)(C.malloc(C.size_t(goSize) * C.size_t(unsafe.Sizeof((*C.char)(nil)))))
    for i := range cArray {
        *(*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(cArrayPtr)) + uintptr(i)*unsafe.Sizeof((*C.char)(nil)))) = 
            unsafe.Pointer(cArray[i])
    }
    
    // Create and return the StringArray structure by value
    return C.StringArray{
        data:  cArrayPtr,
        count: C.int(goSize),
    }
}

//export SetStringArrayValue
func SetStringArrayValue(arr C.StringArray, index C.int, value *C.char) {
    goIndex := int(index)
    
    if goIndex < 0 || goIndex >= int(arr.count) {
        return // Índice fuera de rango
    }
    
    // Liberamos el string anterior si existía
    ptr := *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(arr.data)) + uintptr(goIndex)*unsafe.Sizeof((*C.char)(nil))))
    if ptr != nil {
        C.free(unsafe.Pointer(ptr))
    }
    
    // Asignamos el nuevo valor
    *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(arr.data)) + uintptr(goIndex)*unsafe.Sizeof((*C.char)(nil)))) = C.CString(C.GoString(value))
}

//export GetStringArrayValue
func GetStringArrayValue(arr C.StringArray, index C.int) *C.char {
    goIndex := int(index)
    
    if goIndex < 0 || goIndex >= int(arr.count) {
        return nil // Índice fuera de rango
    }
    
    return *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(arr.data)) + uintptr(goIndex)*unsafe.Sizeof((*C.char)(nil))))
}

//export GetStringArraySize
func GetStringArraySize(arr C.StringArray) C.int {
    return arr.count
}

//export JoinStringArray
func JoinStringArray(arr C.StringArray, delimiter *C.char) *C.char {
    goDelimiter := C.GoString(delimiter)
    var builder strings.Builder
    
    for i := 0; i < int(arr.count); i++ {
        if i > 0 {
            builder.WriteString(goDelimiter)
        }
        ptr := *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(arr.data)) + uintptr(i)*unsafe.Sizeof((*C.char)(nil))))
        if ptr != nil {
            builder.WriteString(C.GoString(ptr))
        }
    }
    
    return C.CString(builder.String())
}

//export Split
func Split(s *C.char, sep *C.char) C.StringArray {
    goStr := C.GoString(s)
    goSep := C.GoString(sep)
    
    // Dividir el string en partes
    parts := strings.Split(goStr, goSep)
    
    // Crear un StringArray en C
    arr := C.StringArray{
        data:  (**C.char)(C.malloc(C.size_t(len(parts)) * C.size_t(unsafe.Sizeof((*C.char)(nil))))),
        count: C.int(len(parts)),
    }
    
    // Convertir cada parte a C string y almacenar en el array
    for i, part := range parts {
        cStr := C.CString(part)
        ptr := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(arr.data)) + uintptr(i)*unsafe.Sizeof((*C.char)(nil))))
        *ptr = cStr
    }
    
    return arr
}

//export FreeStringArray
func FreeStringArray(arr C.StringArray) {
    // Liberamos cada string individual
    for i := 0; i < int(arr.count); i++ {
        ptr := *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(arr.data)) + uintptr(i)*unsafe.Sizeof((*C.char)(nil))))
        if ptr != nil {
            C.free(unsafe.Pointer(ptr))
        }
    }
    
    // Liberamos el array de punteros
    C.free(unsafe.Pointer(arr.data))
}

//export FreeString
func FreeString(s *C.char) {
    C.free(unsafe.Pointer(s))
}

func main() {} // Necesario para buildear como plugin