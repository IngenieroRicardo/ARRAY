# array

Librería en C para manipular Arreglos  
Compilar la libreria: `go build -o array.dll -buildmode=c-shared array.go`

---

### 📥 Descargar la librería

| Linux | Windows |
| --- | --- |
| `wget https://github.com/IngenieroRicardo/array/releases/download/1.0/array.so` | `Invoke-WebRequest https://github.com/IngenieroRicardo/array/releases/download/1.0/array.dll -OutFile ./array.dll` |
| `wget https://github.com/IngenieroRicardo/array/releases/download/1.0/array.h` | `Invoke-WebRequest https://github.com/IngenieroRicardo/array/releases/download/1.0/array.h -OutFile ./array.h` |

---

### 🛠️ Compilar

| Linux | Windows |
| --- | --- |
| `gcc -o main.bin main.c ./array.so` | `gcc -o main.exe main.c ./array.dll` |
| `x86_64-w64-mingw32-gcc -o main.exe main.c ./array.dll` |  |

---

### 🧪 Ejemplo 1: Manipulación básica de strings

```C
#include <stdio.h>
#include "array.h"

int main() {
    // Conversión de tipos
    String numStr = "123";
    int num = Atoi(numStr);
    printf("Atoi: %s -> %d\n", numStr, num);
    
    String floatStr = "3.14159";
    double pi = Atof(floatStr);
    printf("Atof: %s -> %f\n", floatStr, pi);
    
    // Creación de strings
    String intStr = Itoa(42);
    printf("Itoa: 42 -> %s\n", intStr);
    
    String floatStr2 = Ftoa(3.14159, 2);
    printf("Ftoa: 3.14159 (prec 2) -> %s\n", floatStr2);
    
    // Modificación de strings
    String original = "   Hola Mundo!   ";
    String trimmed = Trim(original);
    printf("Trim: '%s' -> '%s'\n", original, trimmed);
    
    String upper = ToUpperCase(trimmed);
    String lower = ToLowerCase(trimmed);
    printf("ToUpperCase: '%s' -> '%s'\n", trimmed, upper);
    printf("ToLowerCase: '%s' -> '%s'\n", trimmed, lower);
    
    // Limpieza de memoria
    FreeString(intStr);
    FreeString(floatStr2);
    FreeString(trimmed);
    FreeString(upper);
    FreeString(lower);
    
    return 0;
}
```

---

### 🧪 Ejemplo 2: Operaciones con arrays de strings

```C
#include <stdio.h>
#include "array.h"

int main() {
    // Crear un array de strings
    StringArray arr = NewStringArray(4);
    SetStringArrayValue(arr, 0, "Manzana");
    SetStringArrayValue(arr, 1, "Banana");
    SetStringArrayValue(arr, 2, "Uvas");
    SetStringArrayValue(arr, 3, "Naranjas");
    
    // Obtener valores individuales
    printf("Fruits:\n");
    for (int i = 0; i < GetStringArraySize(arr); i++) {
        printf("- %s\n", GetStringArrayValue(arr, i));
    }
    
    // Unir strings con un delimitador
    String joined = JoinStringArray(arr, ", ");
    printf("Joined: %s\n", joined);
    
    // Dividir un string
    String text = "uno,dos,tres,cuatro";
    StringArray splitArr = Split(text, ",");
    printf("Split '%s':\n", text);
    for (int i = 0; i < GetStringArraySize(splitArr); i++) {
        printf("- %s\n", GetStringArrayValue(splitArr, i));
    }
    
    // Limpieza de memoria
    FreeString(joined);
    FreeStringArray(arr);
    FreeStringArray(splitArr);
    
    return 0;
}
```

### 🧪 Ejemplo 3: Concatenación y reemplazo avanzado

```C
#include <stdio.h>
#include "array.h"

int main() {
    // Concatenación básica con Concat
    String saludo = Concat("Hola", " ", "Mundo", NULL);
    printf("Saludo combinado: %s\n", saludo);
    
    // Concatenación avanzada con ConcatAll 
    String partes[] = {
        "El rápido ", 
        "zorro marrón ", 
        "salta sobre ", 
        "el perro perezoso.", 
        NULL
    };
    String fraseCompleta = ConcatAll(partes, 4); // 4 elementos (sin contar NULL)
    printf("Frase completa: %s\n", fraseCompleta);
    
    // Reemplazo de texto en un string
    String textoOriginal = "Me gustan las manzanas, las manzanas son mis favoritas.";
    String textoModificado = ReplaceAll(textoOriginal, "manzanas", "naranjas");
    printf("Texto original: %s\n", textoOriginal);
    printf("Texto modificado: %s\n", textoModificado);
    
    // Operaciones con substrings y longitud
    String textoLargo = "Este es un texto muy largo para el ejemplo";
    int longitud = StrLen(textoLargo);
    String fragmento = Substring(textoLargo, 8, 22);
    printf("Longitud del texto: %d caracteres\n", longitud);
    printf("Fragmento (8-22): '%s'\n", fragmento);
    
    // LIMPIEZA - Liberar memoria asignada
    FreeString(saludo);
    FreeString(fraseCompleta);
    FreeString(textoModificado);
    FreeString(fragmento);
    
    return 0;
}
```


---


## 📚 Documentación de la API
### Funciones Principales

#### Definiciones de tipos de variables equivalentes
- ` char* = String `
- ` char** = StringArray `
- ` int* = IntArray `

#### Conversiones
- `int Atoi(char* str)`: Conviere string a entero
- `int ParseBool(char* str)`: Conviere string a booleano
- `double Atof(char* str)`: Conviere string a decimal
- `char* Itoa(int n)`: Conviere entero a string
- `char* Ftoa(double flo, int precision)`: Conviere decimal a string, especificadon decimales

#### IntArray
- `int GetIntArraySize(int* ints)`: Obtiene longitud del int array
- `int SumaAll(int* ints)`: Suma todos los int del array
- `char** NewIntArray(int size)`: Crea int array vacío

#### String
- `int Contains(char* str, char* substr)`: Busca coincidencias dentro del string
- `int Equals(char* str1, char* str2)`: Compara dos string
- `int GetStringSize(char* str)`: Obtiene longitud del string
- `int IsNumeric(char* str)`: Valida si el string es numero
- `char* SubString(char* str, int start, int end)`: Extrae string mediante indices
- `char* ToLowerCase(char* str)`: Convierte en minusculas el string
- `char* ToUpperCase(char* str);`: Convierte en mayusculas el string
- `char* Trim(char* str)`: Quita los espacios en blanco al inicio y final del string
- `char* ReplaceAll(char* str, char* old, char* new)`: Reemplaza una string dentro de otro string

#### StringArray
- `int GetStringArraySize(char** strs)`: Obtiene longitud del string array
- `char* ConcatAll(char** strs)`: Concatena todos los string del array
- `char** Split(char* str, char* sep)`: Convierte string a string array por un separador
- `char** NewStringArray(int size)`: Crea string array vacío

#### Utilidades
- `void FreeString(char* str)`: Libera memoria de string
- `void FreeStringArray(char** strs)`: Libera memoria de string array

---
