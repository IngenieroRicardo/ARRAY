package main

/*
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdarg.h>

typedef char* String; // Definimos String como alias de char*
typedef char** StringArray; // Definimos String como alias de char*

static String Concat(String first, ...) {
    va_list args;
    String token;
    size_t total_len = 0;
    
    // 1. Calcular longitud total necesaria
    va_start(args, first);
    token = first;
    while(token != NULL) {
        total_len += strlen(token);
        token = va_arg(args, String);
    }
    va_end(args);
    
    // 2. Reservar memoria (incluyendo espacio para \0)
    String result = malloc(total_len + 1);
    if(result == NULL) return NULL;
    result[0] = '\0';
    
    // 3. Concatenar todos los strings
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


// ========== Funciones de Conversión ==========

//export Atoi
func Atoi(s C.String) C.int {
	goStr := C.GoString(s)
	val, err := strconv.Atoi(goStr)
	if err != nil {
		return 0
	}
	return C.int(val)
}

//export Atof
func Atof(s C.String) C.double {
	goStr := C.GoString(s)
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
func Ftoa(f C.double, precision C.int) C.String {
	return C.CString(strconv.FormatFloat(float64(f), 'f', int(precision), 64))
}

//export ParseBool
func ParseBool(s C.String) C.int {
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
func StrLen(s C.String) C.int {
	return C.int(len(C.GoString(s)))
}

//export Substring
func Substring(s C.String, start, end C.int) C.String {
	goStr := C.GoString(s)
	goStart := int(start)
	goEnd := int(end)
	
	if goStart < 0 || goEnd > len(goStr) || goStart > goEnd {
		return C.CString("")
	}
	
	return C.CString(goStr[goStart:goEnd])
}

//export IsNumeric
func IsNumeric(s C.String) C.int {
	goStr := C.GoString(s)
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

    // Convert C string array to Go []string
    var goStrings []string
    for i := 0; ; i++ {
        // Get pointer to the i-th string
        ptr := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(strs)) + uintptr(i)*unsafe.Sizeof(strs)))
        if *ptr == nil {
            break
        }
        goStrings = append(goStrings, C.GoString(*ptr))
    }

    // Join all strings
    result := strings.Join(goStrings, "")

    // Return as C string (caller must free this)
    return C.CString(result)
}

//export ToUpperCase
func ToUpperCase(s C.String) C.String {
	goStr := C.GoString(s)
	return C.CString(strings.ToUpper(goStr))
}

//export ToLowerCase
func ToLowerCase(s C.String) C.String {
	goStr := C.GoString(s)
	return C.CString(strings.ToLower(goStr))
}

//export Trim
func Trim(s C.String) C.String {
	goStr := C.GoString(s)
	return C.CString(strings.TrimSpace(goStr))
}

//export ReplaceAll
func ReplaceAll(s, old, new C.String) C.String {
	goStr := C.GoString(s)
	goOld := C.GoString(old)
	goNew := C.GoString(new)
	return C.CString(strings.ReplaceAll(goStr, goOld, goNew))
}






//export Equals
func Equals(s1 *C.char, s2 *C.char) C.int {
    // Convert C strings to Go strings
    goStr1 := C.GoString(s1)
    goStr2 := C.GoString(s2)
    
    // Compare the strings
    if goStr1 == goStr2 {
        return C.int(1)  // true
    }
    return C.int(0)       // false
}


//export Contains
func Contains(s *C.char, substr *C.char) C.int {
    // Convert C strings to Go strings
    goStr := C.GoString(s)
    goSubstr := C.GoString(substr)
    
    // Check if string contains substring
    if strings.Contains(goStr, goSubstr) {
        return C.int(1)  // true
    }
    return C.int(0)      // false
}


// ========== Funciones StringArray ==========




//export NewStringArray
func NewStringArray(size C.int) **C.char {
    if size <= 0 {
        return nil
    }

    // Calculate total memory needed (size elements + NULL terminator)
    elementSize := unsafe.Sizeof((*C.char)(nil))
    totalSize := uintptr(size+1) * elementSize

    // Allocate memory
    ptr := C.malloc(C.size_t(totalSize))
    if ptr == nil {
        return nil
    }

    // Zero out the memory (sets all pointers to NULL)
    C.memset(ptr, 0, C.size_t(totalSize))

    return (**C.char)(ptr)
}



//export Split
func Split(s *C.char, sep *C.char) **C.char {
    goStr := C.GoString(s)
    goSep := C.GoString(sep)
    
    parts := strings.Split(goStr, goSep)
    
    // Allocate the array (size + 1 for NULL terminator)
    arr := (**C.char)(C.malloc(C.size_t(len(parts)+1) * C.size_t(unsafe.Sizeof((*C.char)(nil)))))
    if arr == nil {
        return nil
    }
    
    // Initialize with NULLs (especially the last element)
    C.memset(unsafe.Pointer(arr), 0, C.size_t(len(parts)+1)*C.size_t(unsafe.Sizeof((*C.char)(nil))))
    
    // Convert each part to C string and store in the array
    for i, part := range parts {
        cStr := C.CString(part)
        ptr := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(arr)) + uintptr(i)*unsafe.Sizeof((*C.char)(nil))))
        *ptr = cStr
    }
    
    return arr
}

//export GetStringArraySize
func GetStringArraySize(arr **C.char) C.int {
    var count C.int = 0

    // Convertimos arr en un puntero a punteros y lo recorremos hasta NULL
    for {
        // offset para acceder al elemento arr[i]
        ptr := *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(arr)) + uintptr(count)*unsafe.Sizeof(*arr)))
        if ptr == nil {
            break
        }
        count++
    }

    return count
}

//export FreeStringArray
func FreeStringArray(arr **C.char) {
    if arr == nil {
        return
    }

    // Recorremos hasta NULL
    for i := 0; ; i++ {
        // Obtener puntero al string en la posición i
        ptr := *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(arr)) + uintptr(i)*unsafe.Sizeof(*arr)))
        if ptr == nil {
            break
        }
        // Liberar cada string
        C.free(unsafe.Pointer(ptr))
    }

    // Liberar el arreglo de punteros
    C.free(unsafe.Pointer(arr))
}












//export FreeString
func FreeString(s C.String) {
    C.free(unsafe.Pointer(s))
}

func main() {} // Necesario para buildear como plugin