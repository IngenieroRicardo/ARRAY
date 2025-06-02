# STRING

Librería en C para manipular Arreglos de Caracteres

---

### 📥 Descargar la librería

| Linux | Windows |
| --- | --- |
| `wget https://raw.githubusercontent.com/IngenieroRicardo/STRING/refs/heads/main/STRING.so` | `Invoke-WebRequest https://raw.githubusercontent.com/IngenieroRicardo/STRING/refs/heads/main/STRING.dll -OutFile ./STRING.dll` |
| `wget https://raw.githubusercontent.com/IngenieroRicardo/STRING/refs/heads/main/STRING.h` | `Invoke-WebRequest https://raw.githubusercontent.com/IngenieroRicardo/STRING/refs/heads/main/STRING.h -OutFile ./STRING.h` |

---

### 🛠️ Compilar

| Linux | Windows |
| --- | --- |
| `gcc -o main.bin main.c ./STRING.so` | `gcc -o main.exe main.c ./STRING.dll` |
| `x86_64-w64-mingw32-gcc -o main.exe main.c ./STRING.dll` |  |

---

### 🧪 Ejemplo 1: Manipulación básica de strings

```C
#include <stdio.h>
#include "STRING.h"

int main() {
    // Conversión de tipos
    char* numStr = "123";
    int num = Atoi(numStr);
    printf("Atoi: %s -> %d\n", numStr, num);
    
    char* floatStr = "3.14159";
    double pi = Atof(floatStr);
    printf("Atof: %s -> %f\n", floatStr, pi);
    
    // Creación de strings
    char* intStr = Itoa(42);
    printf("Itoa: 42 -> %s\n", intStr);
    
    char* floatStr2 = Ftoa(3.14159, 2);
    printf("Ftoa: 3.14159 (prec 2) -> %s\n", floatStr2);
    
    // Modificación de strings
    char* original = "   Hola Mundo!   ";
    char* trimmed = Trim(original);
    printf("Trim: '%s' -> '%s'\n", original, trimmed);
    
    char* upper = ToUpperCase(trimmed);
    char* lower = ToLowerCase(trimmed);
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
#include "STRING.h"

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
    char* joined = JoinStringArray(arr, ", ");
    printf("Joined: %s\n", joined);
    
    // Dividir un string
    char* text = "uno,dos,tres,cuatro";
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
#include "STRING.h"

int main() {
    // --------------------------------------------------
    // 1. Concatenación básica con Concat
    // (para unir pocos strings conocidos)
    // --------------------------------------------------
    char* saludo = Concat("Hola", " ", "Mundo", NULL);
    printf("Saludo combinado: %s\n", saludo);
    
    // --------------------------------------------------
    // 2. Concatenación avanzada con ConcatAll 
    // (para unir arrays de strings)
    // --------------------------------------------------
    char* partes[] = {
        "El rápido ", 
        "zorro marrón ", 
        "salta sobre ", 
        "el perro perezoso.", 
        NULL
    };
    char* fraseCompleta = ConcatAll(partes, 4); // 4 elementos (sin contar NULL)
    printf("Frase completa: %s\n", fraseCompleta);
    
    // --------------------------------------------------
    // 3. Reemplazo de texto en un string
    // --------------------------------------------------
    char* textoOriginal = "Me gustan las manzanas, las manzanas son mis favoritas.";
    char* textoModificado = ReplaceAll(textoOriginal, "manzanas", "naranjas");
    printf("Texto original: %s\n", textoOriginal);
    printf("Texto modificado: %s\n", textoModificado);
    
    // --------------------------------------------------
    // 4. Operaciones con substrings y longitud
    // --------------------------------------------------
    char* textoLargo = "Este es un texto muy largo para el ejemplo";
    int longitud = StrLen(textoLargo);
    char* fragmento = Substring(textoLargo, 8, 22);
    printf("Longitud del texto: %d caracteres\n", longitud);
    printf("Fragmento (8-22): '%s'\n", fragmento);
    
    // --------------------------------------------------
    // LIMPIEZA - Liberar memoria asignada
    // --------------------------------------------------
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

#### Conversiones
- `int Atoi(char* s)`: Conviere una cadena de caracteres a entero
- `double Atof(char* s)`: Conviere una cadena de caracteres a decimal
- `char* Itoa(int n)`: Conviere un entero a cadena de caracteres
- `char* Ftoa(double f, int precision)`: Conviere una decimal a cadena de caracteres, especificadon decimales
- `int ParseBool(char* s)`: Conviere una cadena de caracteres booleana a entero

#### String
- `int StrLen(char* s)`: Obtiene longitud del string
- `char* Substring(char* s, int start, int end)`: Extrae un string mediante indices
- `IsNumeric(char* s)`: Valida si es un valor numerico
- `char* ConcatAll(char** strs, int count)`: Concatena todos los string del array
- `char* ToUpperCase(char* s);`: Convierte en mayusculas todos los caracteres
- `char* ToLowerCase(char* s)`: Convierte en minusculas todos los caracteres
- `char* Trim(char* s)`: Quita los espacios en blanco al inicio y final de un string
- `char* ReplaceAll(char* s, char* old, char* new)`: Reemplaza una string dentro de otro string

#### StringArray
- `StringArray NewStringArray(int size)`: Crea objeto StringArray vacío
- `void SetStringArrayValue(StringArray arr, int index, char* value)`
- `char* GetStringArrayValue(StringArray arr, int index)`
- `int GetStringArraySize(StringArray arr)`
- `char* JoinStringArray(StringArray arr, char* delimiter)`
- `StringArray Split(char* s, char* sep)`: Convierte un string en StringArray por un separador
- `JsonResult MergeJSON(const char* json1, const char* json2)`: Combina dos JSONs

#### Utilidades
- `void FreeString(char* s)`: Libera memoria
- `void FreeStringArray(StringArray arr)`: Libera memoria de arrays

### Estructuras
```c
typedef struct {
    char** data;     // Array de elementos
    int count;        // Número de elementos
} StringArray;
```

---
