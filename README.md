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
    // Ejemplo de conexión y consulta
    char* conexion = "root:123456@tcp(127.0.0.1:3306)/test";
    char* query = "SELECT now();"; //Construcción de JSON desde Result
    //char* query = "SELECT '{\"status\": \"OK\"}' AS JSON"; //Construcción de JSON desde Query
    
    SQLResult resultado = SQLrun(conexion, query, NULL, 0);
    
    if (resultado.is_error) {
        printf("Error: %s\n", resultado.json);
    } else if (resultado.is_empty) {
        printf("Consulta ejecutada pero no retornó datos\n");
        printf("JSON: %s\n", resultado.json); // Mostrará {"status":"OK"} o []
    } else {
        printf("Datos obtenidos:\n%s\n", resultado.json);
    }
    
    // Liberar memoria
    FreeSQLResult(resultado);
    
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
    char* text = "one,two,three,four";
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


---

### 🧪 Ejemplo 3: Concatenación y reemplazo avanzado

```C
#include <stdio.h>
#include "STRING.h"

int main() {
    // Crear un array de strings
    StringArray arr = NewStringArray(4);
    SetStringArrayValue(arr, 0, "Apple");
    SetStringArrayValue(arr, 1, "Banana");
    SetStringArrayValue(arr, 2, "Cherry");
    SetStringArrayValue(arr, 3, "Date");
    
    // Obtener valores individuales
    printf("Fruits:\n");
    for (int i = 0; i < GetStringArraySize(arr); i++) {
        printf("- %s\n", GetStringArrayValue(arr, i));
    }
    
    // Unir strings con un delimitador
    char* joined = JoinStringArray(arr, ", ");
    printf("Joined: %s\n", joined);
    
    // Dividir un string
    char* text = "one,two,three,four";
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
