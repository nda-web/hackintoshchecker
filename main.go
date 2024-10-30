package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// Códigos de color ANSI
const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

// GPUInfo almacena la información de la tarjeta gráfica
type GPUInfo struct {
	Name     string
	Vendor   string
	DeviceID string
	MemoryGB int // Almacena la memoria de la GPU en GB
}

func main() {
	// Mostrar el arte ASCII en el inicio
	fmt.Println(`    __  __           __   _       __             __  
   / / / /___ ______/ /__(_)___  / /_____  _____/ /_ 
  / /_/ / __ '/ ___/ //_/ / __ \/ __/ __ \/ ___/ __ \
 / __  / /_/ / /__/ ,< / / / / / /_/ /_/ (__  ) / / /
/_/ /_/\__,_/\___/_/|_/_/_/ /_/\__/\____/____/_/ /_/ 
   ________              __                          
  / ____/ /_  ___  _____/ /_____  _____              
 / /   / __ \/ _ \/ ___/ //_/ _ \/ ___/              
/ /___/ / / /  __/ /__/ ,< /  __/ /                  
\____/_/ /_/\___/\___/_/|_|\___/_/                   
`)

	fmt.Println("Creadores: Martin Oviedo & Daedalus")
	fmt.Println("----------------------------------------\n")
	fmt.Println("\n### Informacion de la PC ###")

	var veredictoCPU, veredictoRAM, veredictoDisk, veredictoUEFI, veredictoSecureBoot, veredictoGPU string
	var secureBoot bool
	var compatibleGPU bool

	// Información del procesador
	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range cpuInfo {
		fmt.Println("----------------------------------------")
		fmt.Printf("CPU: %s\nNúcleos: %d, Frecuencia: %.2f GHz\n", c.ModelName, c.Cores, c.Mhz/1000)
		fmt.Println("----------------------------------------")

		if strings.Contains(c.ModelName, "Intel") {
			veredictoCPU = fmt.Sprintf("%sTu CPU es Intel, compatible con Hackintosh.%s", Green, Reset)
		} else if strings.Contains(c.ModelName, "AMD") {
			veredictoCPU = fmt.Sprintf("%sTu CPU es AMD, es posible instalar Hackintosh, pero requerirá parches adicionales.%s", Yellow, Reset)
		} else {
			veredictoCPU = fmt.Sprintf("%sNo se pudo identificar el tipo de CPU. Verifica si tu procesador es Intel o AMD.%s", Red, Reset)
		}
	}

	// Detectar GPU
	gpus := getGPUInfo() // Captura el valor de retorno
	fmt.Println("\n### Información de GPU ###")
	for _, gpu := range gpus {
		fmt.Printf("GPU: %s\nFabricante: %s\nDevice ID: %s\nMemoria de GPU: %d GB\n", gpu.Name, gpu.Vendor, gpu.DeviceID, gpu.MemoryGB)

		compatibleGPU = isGPUCompatible(gpu)
		if compatibleGPU {
			veredictoGPU = fmt.Sprintf("%sGPU compatible con Hackintosh: %s%s", Green, gpu.Name, Reset)
		} else {
			veredictoGPU = fmt.Sprintf("%sGPU no compatible o requiere configuración especial: %s%s", Yellow, gpu.Name, Reset)
		}
	}

	// Información de la memoria
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nMemoria total: %v MB\n", memInfo.Total/1024/1024)
	if memInfo.Total >= 8*1024*1024*1024 {
		veredictoRAM = fmt.Sprintf("%sTienes suficiente memoria RAM para instalar macOS (mínimo 8GB).%s", Green, Reset)
	} else if memInfo.Total >= 4*1024*1024*1024 {
		veredictoRAM = fmt.Sprintf("%sTienes suficiente memoria RAM para instalar macOS, pero se recomienda tener al menos 8GB.%s", Yellow, Reset)
	} else {
		veredictoRAM = fmt.Sprintf("%sNo tienes suficiente memoria RAM para instalar macOS (mínimo 4GB).%s", Red, Reset)
	}

	// Información del disco
	diskInfo, err := disk.Usage("/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Disco total: %v GB, Usado: %.2f%%\n", diskInfo.Total/1024/1024/1024, diskInfo.UsedPercent)
	if diskInfo.Free > 50*1024*1024*1024 {
		veredictoDisk = fmt.Sprintf("%sTienes suficiente espacio libre en el disco (mínimo 50 GB).%s", Green, Reset)
	} else {
		veredictoDisk = fmt.Sprintf("%sNo tienes suficiente espacio libre en el disco. Se recomienda al menos 50 GB libres.%s", Red, Reset)
	}

	// Verificaciones de UEFI y Secure Boot
	isUEFI, err := checkUEFIMode()
	if err != nil {
		veredictoUEFI = fmt.Sprintf("No se pudo determinar el modo UEFI: %v", err)
	} else {
		if isUEFI {
			veredictoUEFI = fmt.Sprintf("%sEl sistema está en modo UEFI, listo para Hackintosh.%s", Green, Reset)
		} else {
			veredictoUEFI = fmt.Sprintf("%sEl sistema está en modo Legacy (BIOS). Necesitarás cambiar a UEFI para instalar Hackintosh.%s", Yellow, Reset)
		}
	}

	if runtime.GOOS == "windows" {
		secureBoot, err = checkSecureBootWindows()
		if err != nil {
			veredictoSecureBoot = fmt.Sprintf("No se pudo determinar el estado de Secure Boot: %v", err)
		} else {
			if secureBoot {
				veredictoSecureBoot = fmt.Sprintf("%sSecure Boot está habilitado. Deberás deshabilitarlo en la BIOS para instalar Hackintosh.%s", Yellow, Reset)
			} else {
				veredictoSecureBoot = fmt.Sprintf("%sSecure Boot está deshabilitado, listo para Hackintosh.%s", Green, Reset)
			}
		}
	}

	// Mostrar el veredicto final
	fmt.Println("\n### Veredicto Final ###")
	fmt.Println(veredictoCPU)
	fmt.Println(veredictoGPU)
	fmt.Println(veredictoRAM)
	fmt.Println(veredictoDisk)
	fmt.Println(veredictoUEFI)
	fmt.Println(veredictoSecureBoot)

	// Veredicto general de compatibilidad
	fmt.Println("\n### Compatibilidad General ###")
	if strings.Contains(veredictoCPU, "compatible") && compatibleGPU &&
		strings.Contains(veredictoRAM, "suficiente") &&
		strings.Contains(veredictoDisk, "suficiente") &&
		isUEFI && !secureBoot {
		fmt.Printf("%sTu sistema es COMPATIBLE con Hackintosh!%s\n", Green, Reset)
	} else {
		fmt.Printf("%sTu sistema requiere ajustes antes de instalar Hackintosh. Revisa los detalles anteriores.%s\n", Yellow, Reset)
	}

	fmt.Println("\n### Consejos para Hackintosh ###")
	if isUEFI && (runtime.GOOS == "windows" && !secureBoot) {
		fmt.Println("1. Asegúrate de tener una copia de seguridad de todos tus datos importantes")
		fmt.Println("2. Descarga los kexts necesarios para tu hardware específico")
		fmt.Println("3. Prepara un USB booteable con OpenCore o Clover")
		if !compatibleGPU {
			fmt.Println("4. Tu GPU puede requerir configuración especial o no ser compatible")
		}
	} else {
		fmt.Println("Debes resolver los problemas mencionados en el veredicto antes de proceder")
	}

	fmt.Println("\nGracias por usar HackintoshChecker")
	fmt.Println("\nPresiona 'Enter' para salir...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}

// Función para capturar la información de la GPU, incluyendo la memoria en GB desde el registro de Windows
func getGPUInfo() []GPUInfo {
	var gpus []GPUInfo

	if runtime.GOOS == "windows" {
		cmd := exec.Command("powershell", "-Command", "$qwMemorySize = (Get-ItemProperty -Path \"HKLM:\\SYSTEM\\ControlSet001\\Control\\Class\\{4d36e968-e325-11ce-bfc1-08002be10318}\\0*\" -Name HardwareInformation.qwMemorySize -ErrorAction SilentlyContinue).\"HardwareInformation.qwMemorySize\"; [math]::round($qwMemorySize / 1GB)")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(Red + "Error ejecutando PowerShell: " + err.Error() + Reset)
			fmt.Println("Salida de PowerShell:", string(output))
			return gpus
		}

		memoryGBStr := strings.TrimSpace(string(output))
		memoryGB, err := strconv.Atoi(memoryGBStr)
		if err != nil {
			fmt.Println(Red + "Error convirtiendo la memoria de VRAM a GB: " + err.Error() + Reset)
			memoryGB = 0 // Default si falla la conversión
		}

		// Aquí asignamos los valores a la GPU (supongamos que ya tienes el nombre y demás)
		gpu := GPUInfo{
			Name:     "Radeon RX 570", // ejemplo
			Vendor:   "AMD",           // ejemplo
			DeviceID: "VideoController1",
			MemoryGB: memoryGB,
		}

		gpus = append(gpus, gpu)
	}

	return gpus
}

// Función para verificar la compatibilidad de la GPU
func isGPUCompatible(gpu GPUInfo) bool {
	gpuName := strings.ToLower(gpu.Name)
	gpuVendor := strings.ToLower(gpu.Vendor)

	// Lista de GPUs compatibles
	if strings.Contains(gpuName, "intel") || strings.Contains(gpuVendor, "intel") {
		return true
	}

	if strings.Contains(gpuName, "amd") || strings.Contains(gpuVendor, "amd") ||
		strings.Contains(gpuName, "radeon") {
		compatibleSeries := []string{"rx", "vega", "polaris"}
		for _, series := range compatibleSeries {
			if strings.Contains(gpuName, series) {
				return true
			}
		}
		return false
	}

	if strings.Contains(gpuName, "nvidia") || strings.Contains(gpuVendor, "nvidia") {
		return false
	}

	return false
}

// checkUEFIMode verifica el modo UEFI
func checkUEFIMode() (bool, error) {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("mountvol")
		output, err := cmd.Output()
		if err != nil {
			return false, err
		}
		return bytes.Contains(output, []byte("EFI")), nil
	case "linux":
		if _, err := os.Stat("/sys/firmware/efi"); err == nil {
			return true, nil
		}
		return false, nil
	default:
		return false, fmt.Errorf("verificación de modo UEFI no soportada para %s", runtime.GOOS)
	}
}

// checkSecureBootWindows verifica el estado de Secure Boot en Windows
func checkSecureBootWindows() (bool, error) {
	cmd := exec.Command("powershell", "-Command", "Confirm-SecureBootUEFI")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return false, fmt.Errorf("error ejecutando PowerShell: %v, %s", err, stderr.String())
	}

	result := strings.TrimSpace(out.String())
	if result == "True" {
		return true, nil
	} else if result == "False" {
		return false, nil
	} else {
		return false, fmt.Errorf("resultado inesperado: %s", result)
	}
}
