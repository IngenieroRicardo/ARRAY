# array

Librería en C para manipular arreglos.  
Para compilar la librería: `go build -o array.dll -buildmode=c-shared array.go`

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

### 🧪 Ejemplo 1

```C
#include <stdio.h>
#include "array.h"

int main() {
    // Concatenación
    String hello = Concat("Hola", " ", "mundo", "!", NULL);
    printf("Concat: %s\n", hello);
    
    // Conversiones
    int num = Atoi("42");
    double pi = Atof("3.14159");
    String num_str = Itoa(123);
    String pi_str = Ftoa(3.14159, 2);
    
    printf("Atoi: %d\n", num);
    printf("Atof: %f\n", pi);
    printf("Itoa: %s\n", num_str);
    printf("Ftoa: %s\n", pi_str);
    
    // Operaciones
    String upper = ToUpperCase("hola");
    String lower = ToLowerCase("MUNDO");
    String trimmed = Trim("  spaces  ");
    
    printf("Upper: %s\n", upper);
    printf("Lower: %s\n", lower);
    printf("Trim: '%s'\n", trimmed);
    
    // Liberar memoria
    FreeString(hello);
    FreeString(num_str);
    FreeString(pi_str);
    FreeString(upper);
    FreeString(lower);
    FreeString(trimmed);

    return 0;
}
```

### 🧪 Ejemplo 2

```C
#include <stdio.h>
#include "array.h"

int main() {
    // Crear y llenar array
    StringArray arr = NewStringArray(3);
    arr[0] = strdup("Primero");
    arr[1] = strdup("Segundo");
    arr[2] = strdup("Tercero");
    
    // Procesar array
    printf("Array size: %d\n", GetStringArraySize(arr));
    char* joined = ConcatAll(arr);
    printf("Joined: %s\n", joined);
    
    // Split de strings
    StringArray parts = Split("uno,dos,tres", ",");
    for (int i = 0; parts[i] != NULL; i++) {
        printf("Part %d: %s\n", i, parts[i]);
    }
    
    // Liberar memoria
    FreeString(joined);
    FreeStringArray(arr);
    FreeStringArray(parts);

    return 0;
}
```

### 🧪 Ejemplo 3

```C
#include <stdio.h>
#include "db.h"

int main() {
    char* diver = "mysql";
    char* conexion = "root:123456@tcp(127.0.0.1:3306)/test";
    
    // Ejemplo 1: Consulta INSERT con parámetros
    char* consulta_insert = "INSERT INTO chat.usuario(nickname, picture) VALUES (?, ?);";
    
    // Preparar los argumentos para el INSERT
    char* argumentos_inser1 = strdup("Ricardo");  // Parámetro de tipo cadena (nickname)
    // Parámetro de tipo blob (imagen codificada en base64)
    char* argumentos_inser2 = strdup("blob::iVBORw0KGgoAAAANSUhEUgAAAAgAAAAICAIAAABLbSncAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAArSURBVBhXY/iPA0AlGBgwGFAKlwQmAKrAIgcVRZODCsI5cAAVgVDo4P9/AHe4m2U/OJCWAAAAAElFTkSuQmCC");
    
    // Ejecutar la consulta INSERT
    SQLResult resultado_insert = SQLrun(diver, conexion, consulta_insert, argumentos_inser1, argumentos_inser2, NULL);
    
    // Mostrar los resultados
    printf("Resultado del INSERT:\n");
    printf("JSON: %s\n", resultado_insert.json);         // Respuesta en formato JSON
    printf("Es error: %d\n", resultado_insert.is_error); // 1 si hubo error, 0 si éxito
    printf("Está vacío: %d\n\n", resultado_insert.is_empty); // 1 para consultas que no retornan datos
    
    // Liberar los recursos utilizados
    FreeSQLResult(resultado_insert); // Liberar la memoria del resultado
    
    return 0;
}
```


---


## 📚 Documentación de la API

#### Definiciones de tipos de variables equivalentes
- ` char* = String `
- ` char** = StringArray `
- ` int* = IntArray `
- ` double* = DoubleArray `

#### Conversiones
- `int Atoi(char* str)`: Convierte string a entero.
- `int ParseBool(char* str)`: Convierte string a booleano.
- `double Atof(char* str)`: Convierte string a decimal.
- `char* Itoa(int n)`: Convierte entero a string.
- `char* Ftoa(double flo, int precision)`: Convierte decimal a string, especificadon decimales.

#### IntArray
- `int* NewIntArray(int size)`: Crea int array vacío.
- `void FreeIntArray(int* ints)`

#### DoubleArray
- `double* NewDoubleArray(int size)`: Crea double array vacío.
- `void FreeDoubleArray(double* flos)`

#### String
- `int Contains(char* str, char* substr)`: Busca coincidencias dentro del string.
- `int Equals(char* str1, char* str2)`: Compara dos string.
- `int GetStringSize(char* str)`: Obtiene longitud del string.
- `int IsNumeric(char* str)`: Valida si el string es numero.
- `char* SubString(char* str, int start, int end)`: Extrae string mediante índices.
- `char* ToLowerCase(char* str)`: Convierte el string a minúsculas.
- `char* ToUpperCase(char* str);`: Convierte el string a mayúsculas.
- `char* Trim(char* str)`: Elimina los espacios en blanco al inicio y al final del string.
- `char* ReplaceAll(char* str, char* old, char* new)`: Reemplaza un string dentro de otro.
- `char* NewString(int size);`: Crea string vacío.
- `void FreeString(char* str)`: Libera memoria

#### StringArray
- `int GetStringArraySize(char** strs)`: Obtiene longitud del string array.
- `char* ConcatAll(char** strs)`: Concatena todos los string del array.
- `char** Split(char* str, char* sep)`: Convierte string a string array por un separador.
- `char** NewStringArray(int size)`: Crea string array vacío.
- `void FreeStringArray(char** strs)`: Libera memoria

---
