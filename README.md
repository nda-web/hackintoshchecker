![logo](https://i3.wp.com/raw.githubusercontent.com/Quamagi/hackintoshchecker/main/logo2.png)

```markdown
¡Por supuesto! He integrado la funcionalidad para detectar correctamente la memoria VRAM de la GPU, incluso cuando sea mayor de 4GB, en tu aplicación de Go. Ahora el programa extrae la información directamente desde el registro de Windows, lo cual asegura una lectura precisa de la memoria de VRAM, incluyendo las GPUs que superan los 4GB de memoria. 

Además, he actualizado la lógica del `README.md` para que refleje estas nuevas capacidades y ofrezca una mejor explicación de las características del programa. Aquí tienes la versión actualizada del `README.md` para GitHub:

```markdown
# HackintoshChecker

HackintoshChecker es una herramienta escrita en **Go** que verifica la compatibilidad de tu hardware para la instalación de **Hackintosh**. El programa recopila información clave del sistema como CPU, RAM, GPU, disco, UEFI y Secure Boot, y muestra un veredicto con colores (verde para OK, amarillo para advertencias y rojo para problemas).

## Características

- **Verificación del procesador (CPU)**: Detecta si es Intel o AMD y evalúa su compatibilidad con Hackintosh.
- **Verificación de la memoria RAM**: Evalúa si tienes suficiente RAM para instalar macOS.
- **Detección de la GPU**: Extrae el nombre, fabricante, y cantidad de VRAM incluso si es mayor de 4GB.
- **Verificación del disco**: Verifica si hay suficiente espacio en el disco para la instalación.
- **Modo UEFI**: Detecta si el sistema está en modo UEFI o Legacy (BIOS).
- **Secure Boot**: Verifica si Secure Boot está habilitado o deshabilitado (solo en Windows).
- **Salida con colores**: Verde para OK, Amarillo para advertencias y Rojo para problemas críticos.

## Requisitos

- **Go**: Necesitarás tener Go instalado para ejecutar este proyecto. Puedes descargarlo desde [golang.org](https://golang.org/dl/).
- **Windows o Linux**: El programa está diseñado para ambos sistemas operativos.

## Instalación

1. **Clona este repositorio**:
   ```bash
   git clone https://github.com/tuusuario/HackintoshChecker.git
   ```

2. **Instala las dependencias** del proyecto:
   ```bash
   go get github.com/shirou/gopsutil/cpu
   go get github.com/shirou/gopsutil/disk
   go get github.com/shirou/gopsutil/mem
   ```

## Compilación

Para **PowerShell**:
```bash
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o HackintoshChecker.exe
```

Para **CMD** (símbolo del sistema de Windows):
```bash
set GOOS=windows && set GOARCH=amd64 && go build -o HackintoshChecker.exe
```

## Uso

1. **Ejecuta el programa**:
   ```bash
   ./HackintoshChecker
   ```

2. El programa te mostrará los resultados con colores para que puedas determinar si tu hardware es compatible con Hackintosh.

## Ejemplo de salida

```plaintext
HackintoshChecker - Versión: 1.0
Creadores: Martin Oviedo & Daedalus

----------------------------------------

### Informacion de la PC ###
----------------------------------------
CPU: AMD Ryzen 5 2600 Six-Core Processor
Núcleos: 12, Frecuencia: 3.40 GHz
----------------------------------------
### Información de GPU ###
GPU: Radeon RX 570
Fabricante: AMD
Device ID: VideoController1
Memoria de GPU: 8 GB
Memoria total: 32719 MB
Disco total: 931 GB, Usado: 17.86%

### Veredicto Final ###
Tu CPU es AMD, es posible instalar Hackintosh, pero requerirá parches adicionales.
GPU compatible con Hackintosh: Radeon RX 570
Tienes suficiente memoria RAM para instalar macOS (mínimo 8GB).
Tienes suficiente espacio libre en el disco (mínimo 50 GB).
El sistema está en modo UEFI, listo para Hackintosh.
Secure Boot está deshabilitado, listo para Hackintosh.
```

## Contribuir

¡Las contribuciones son bienvenidas! Si tienes ideas para mejorar el proyecto o encuentras algún problema, abre un **pull request** o un **issue** en el repositorio.

## Créditos

- **Martin Oviedo**: Desarrollador principal.
- **Daedalus**: Asistente extraordinario.

## Licencia

Este proyecto está bajo la licencia MIT. Consulta el archivo `LICENSE` para más detalles.

## Arte ASCII

El programa ahora incluye un genial arte ASCII al inicio:

```plaintext
    __  __           __   _       __             __  
   / / / /___ ______/ /__(_)___  / /_____  _____/ /_ 
  / /_/ / __ `/ ___/ //_/ / __ \/ __/ __ \/ ___/ __ \
 / __  / /_/ / /__/ ,< / / / / / /_/ /_/ (__  ) / / /
/_/ /_/\__,_/\___/_/|_/_/_/ /_/\__/\____/____/_/ /_/ 
   ________              __                          
  / ____/ /_  ___  _____/ /_____  _____              
 / /   / __ \/ _ \/ ___/ //_/ _ \/ ___/              
/ /___/ / / /  __/ /__/ ,< /  __/ /                  
\____/_/ /_/\___/\___/_/|_|\___/_/                   
```

¡Gracias por usar HackintoshChecker! 😊
```

### **Cambios incluidos en el README:**
1. **Nueva funcionalidad para detectar memoria VRAM**: Ahora la herramienta obtiene la memoria VRAM directamente desde el registro de Windows, permitiendo leer valores superiores a 4 GB.
2. **Salida mejorada**: Incluye información detallada de la GPU (nombre, fabricante, ID del dispositivo, memoria).
3. **Créditos actualizados**: Añadí reconocimiento a **Martin Oviedo** y **Daedalus**.

Este `README.md` refleja las mejoras recientes y brinda una explicación clara de las capacidades del programa. Déjame saber si necesitas más cambios o ajustes. 😊
