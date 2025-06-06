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

```

### 🧪 Ejemplo 2

```C

```

### 🧪 Ejemplo 3

```C

```


---


## 📚 Documentación de la API

#### Definiciones de tipos de variables equivalentes
- ` char* = String `
- ` char** = StringArray `

FALTA: - ` int* = IntArray ` 

#### Conversiones
- `int Atoi(char* str)`: Convierte string a entero.
- `int ParseBool(char* str)`: Convierte string a booleano.
- `double Atof(char* str)`: Convierte string a decimal.
- `char* Itoa(int n)`: Convierte entero a string.
- `char* Ftoa(double flo, int precision)`: Convierte decimal a string, especificadon decimales.

#### IntArray

FALTA: - `int GetIntArraySize(int* ints)`: Obtiene longitud del int array.

FALTA: - `int SumaAll(int* ints)`: Suma todos los int del array.

FALTA: - `char** NewIntArray(int size)`: Crea int array vacío.

#### String
- `int Contains(char* str, char* substr)`: Busca coincidencias dentro del string.
- `int Equals(char* str1, char* str2)`: Compara dos string.

MODIFICAR: - `int GetStringSize(char* str)`: Obtiene longitud del string.

- `int IsNumeric(char* str)`: Valida si el string es numero.
- `char* SubString(char* str, int start, int end)`: Extrae string mediante índices.
- `char* ToLowerCase(char* str)`: Convierte el string a minúsculas.
- `char* ToUpperCase(char* str);`: Convierte el string a mayúsculas.
- `char* Trim(char* str)`: Elimina los espacios en blanco al inicio y al final del string.
- `char* ReplaceAll(char* str, char* old, char* new)`: Reemplaza un string dentro de otro.

#### StringArray
- `int GetStringArraySize(char** strs)`: Obtiene longitud del string array.
- `char* ConcatAll(char** strs)`: Concatena todos los string del array.
- `char** Split(char* str, char* sep)`: Convierte string a string array por un separador.
- `char** NewStringArray(int size)`: Crea string array vacío.

#### Liberar memoria
- `void FreeString(char* str)`
- `void FreeStringArray(char** strs)`

---
