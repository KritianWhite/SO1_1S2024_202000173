package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

// ================= STRUCTS PARA PROCESOS =================
// childProcess is the model for the child process
type childTasks struct {
	PID       int    `json:"pid"`
	Name      string `json:"name"`
	State     int    `json:"state"`
	PIDPadre  int    `json:"pidPadre"`
	RssChild  int    `json:"rssChild"`
	UserChild int    `json:"childUID"`
}

// ProcessParent is the model for the parent process
type TasksParent struct {
	IdTask int          `json:"idTask"`
	PID    int          `json:"pid"`
	Name   string       `json:"name"`
	User   int          `json:"user"`
	State  int          `json:"state"`
	RAM    int          `json:"ram"`
	Rss    int          `json:"rss"`
	Child  []childTasks `json:"child"`
}

type CpuData struct {
	Cpu_total      float64       `json:"Cpu_total"`
	Cpu_porcentaje float64       `json:"Cpu_porcentaje"`
	Cpu_percentage float64       `json:"Cpu_percentage"`
	Processes      []TasksParent `json:"Processes"`
	Running        int           `json:"running"`
	Sleeping       int           `json:"sleeping"`
	Zombie         int           `json:"zombie"`
	Stopped        int           `json:"stopped"`
	Total          int           `json:"total"`
}

type CPUDataHistoric struct {
	ID_CPU         int     `json:"id"`
	Cpu_porcentaje float64 `json:"porcentaje_utilizado"`
	Time           string  `json:"time"`
}

// STRUCTS PARA RAM
// RawRAM es una estructura intermedia para deserializar datos de RAM desde JSON
type RawRAM struct {
	TotalRam     int     `json:"TotalRam"`
	Libre        int     `json:"Libre"`
	MemoriaEnUso int     `json:"MemoriaEnUso"`
	Porcentaje   float64 `json:"Porcentaje"`
}

// RAM es la estructura de información de RAM
type RAM struct {
	TotalRam     int     `json:"total_memoria"`
	Libre        int     `json:"memoria_libre"`
	MemoriaEnUso int     `json:"memoria_utilizada"`
	Porcentaje   float64 `json:"porcentaje_utilizado"`
}

type RamDataHistoric struct {
	ID_RAM             int     `json:"id"`
	PorcentajeUtiizado float64 `json:"porcentaje_utilizado"`
	Time               string  `json:"time"`
}

// Función para convertir de RawRAM a RAM
func ConvertToRAM(rawRAM RawRAM) RAM {
	return RAM{
		TotalRam:     rawRAM.TotalRam,
		Libre:        rawRAM.Libre,
		MemoriaEnUso: rawRAM.MemoriaEnUso,
		Porcentaje:   rawRAM.Porcentaje,
	}
}

// ================= FUNCION MAIN =================
func main() {
	fmt.Println("¡Hola, mundo!")

	app := fiber.New()

	// Middleware de log
	app.Use(logger.New())

	// Middleware de CORS
	app.Use(cors.New())

	// Conectar a la base de datos
	if err := Connect(); err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Defer para cerrar la conexión a la base de datos al salir de main
	defer CloseDB()

	// Iniciar la rutina de actualización de datos de CPU
	go UpdateCPU()

	// Iniciar la rutina de actualización de datos de RAM
	go UpdateRAM()

	// Configurar rutas
	setupRoutes(app)

	port := 5000
	fmt.Printf("Escuchando en http://localhost:%d\n", port)

	// Iniciar el servidor
	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}

func setupRoutes(app *fiber.App) {

	// Rutas relacionadas con procesos
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Rutas relacionadas con CPU
	app.Get("/api/cpu", HandleCPUData)

	app.Get("/api/cpu_historic", GetCpuDataHistoric)

	// Rutas relacionadas con RAM
	app.Get("/api/ram", HandleRAMData)

	app.Get("/api/ram_historic", GetRamDataHistoric)

	app.Get("/api/tasks_ids", GetListTaskPid)

	app.Get("/api/generateTreeTasks/:pid", CreateTreeTask)

	app.Get("/api/startTasks", ControllerStartTask)

	app.Get("/api/stopTasks", ControllerStopTask)

}

// ================= VARIABLES GLOBALES =================
var db *sql.DB
var DATABASE_URI string

const (
	cpuFilePath = "/proc/cpu_so1_1s2024"
	ramFilePath = "/proc/ram_so1_1s2024"
)

var cpuDataChan = make(chan string) // Channel to send the CPU data
var ramDataChan = make(chan string) // Channel to send the RAM data

var tasks *exec.Cmd

// ================= FUNCIONES PARA LA BASE DE DATOS =================

func Connect() error {
	loadEnvVariables()

	var err error

	db, err = sql.Open("mysql", DATABASE_URI)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	/*
		deleteChildTaskTable()
		deleteTaskTable()
		deleteRamTable()
		deleteCpuTable()
	*/
	createTaskTable()
	createChildTaskTable()
	createRamTable()
	createCpuTable()
	createStateTable()

	return nil
}

func loadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBUser := os.Getenv("DB_USER")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")

	// Si DBPassword está vacío, no incluirlo en la cadena de conexión
	DBPassword := os.Getenv("DB_PASSWORD")
	if DBPassword != "" {
		DBPassword = ":" + DBPassword
	}

	DATABASE_URI = DBUser + DBPassword + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	//fmt.Println(DATABASE_URI)
}

func createTaskTable() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS TASK (
			ID_TASK INT PRIMARY KEY AUTO_INCREMENT, 
			PID INT,
			NOMBRE VARCHAR(255),
			ESTADO INT,
			RSS INT,
			UID INT,
			RAM INT,
			TIME DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		fmt.Println("Error creating task table:", err)
	}
}

func createChildTaskTable() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS CHILD_TASK (
			ID_CHILD_TASK INT PRIMARY KEY AUTO_INCREMENT,
			ID_TASK INT,
			PID INT,
			NOMBRE VARCHAR(255),
			ESTADO INT,
			RSS INT,
			UID INT,
			TIME DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (ID_TASK) REFERENCES TASK (ID_TASK)
		);
	`)
	if err != nil {
		fmt.Println("Error creating child Task table:", err)
	}
}

func createRamTable() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS RAM (
			ID_RAM INT PRIMARY KEY AUTO_INCREMENT,
			TOTAL_MEMORIA FLOAT,
			MEMORIA_LIBRE FLOAT,
			MEMORIA_UTILIZADA FLOAT,
			PORCENTAJE_UTILIZADO FLOAT,
			TIME DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		fmt.Println("Error creating child Ram table:", err)
	}
}

func createCpuTable() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS CPU (
			ID_CPU INT PRIMARY KEY AUTO_INCREMENT,
			CPU_TOTAL FLOAT,
			CPU_PORCENTAJE FLOAT,
			CPU_PERCENTAGE FLOAT,
			RUNNING INT,
			SLEEPING INT,
			ZOMBIE INT,
			STOPPED INT,
			TOTAL INT,
			TIME DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		fmt.Println("Error creating CPU table:", err)
	}

}

func createStateTable() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS STATE (
			ID_STATE INT PRIMARY KEY AUTO_INCREMENT,
			ID_TASK INT,
			NAME VARCHAR (200),
			TIME DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		fmt.Println("Error creating state table:", err)
	}
}

func deleteTaskTable() {
	_, err := db.Exec("DROP TABLE IF EXISTS TASK;")
	if err != nil {
		fmt.Println("Error deleting Task table:", err)
	}
}

func deleteChildTaskTable() {
	_, err := db.Exec("DROP TABLE IF EXISTS CHILD_TASK;")
	if err != nil {
		fmt.Println("Error deleting child Task table:", err)
	}
}

func deleteRamTable() {
	_, err := db.Exec("DROP TABLE IF EXISTS RAM;")
	if err != nil {
		fmt.Println("Error deleting ram table:", err)
	}
}

func deleteCpuTable() {
	_, err := db.Exec("DROP TABLE IF EXISTS CPU;")
	if err != nil {
		fmt.Println("Error deleting cpu table:", err)
	}
}

func deleteStateTable() {
	_, err := db.Exec("DROP TABLE IF EXISTS STATE;")
	if err != nil {
		fmt.Println("Error deleting state table:", err)
	}
}

func InsertRam(totalMemory, freeMemory, usedMemory int, usedPercentage float64) error {
	_, err := db.Exec(`
		INSERT INTO RAM (TOTAL_MEMORIA, MEMORIA_LIBRE, MEMORIA_UTILIZADA, PORCENTAJE_UTILIZADO)
		VALUES (?, ?, ?, ?);
	`, totalMemory, freeMemory, usedMemory, usedPercentage)
	return err
}

func InsertCpu(cpuTotal float64, cpuPercentage float64, cpuPorcentaje float64, running int, sleeping int, zombie int, stopped int, total int) error {
	_, err := db.Exec(`
		INSERT INTO CPU (CPU_TOTAL, CPU_PERCENTAGE, CPU_PORCENTAJE, RUNNING, SLEEPING, ZOMBIE, STOPPED, TOTAL)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?);
	`, cpuTotal, cpuPercentage, cpuPorcentaje, running, sleeping, zombie, stopped, total)
	return err
}

func InsertTask(pid int, nombre string, estado int, rss int, uid int, RAM int) (int, error) {
	result, err := db.Exec(`
		INSERT INTO TASK (PID, NOMBRE, ESTADO, RSS, UID, RAM)
		VALUES (?, ?, ?, ?, ?, ?);
	`, pid, nombre, estado, rss, uid, RAM)
	if err != nil {
		return 0, err
	}
	taskId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(taskId), nil
}

func InsertChildTask(idTask int, pid int, nombre string, estado int, rss int, uidChild int) error {
	_, err := db.Exec(`
		INSERT INTO CHILD_TASK (ID_TASK, PID, NOMBRE, ESTADO, RSS, UID)
		VALUES (?, ?, ?, ?, ?, ?);
	`, idTask, pid, nombre, estado, rss, uidChild)
	return err
}

func GetListTasks() ([]RamDataHistoric, error) {
	// limitTime := time.Now().Add(-10 * time.Minute)
	rows, err := db.Query(`
		SELECT ID_RAM, PORCENTAJE_UTILIZADO, TIME
		FROM RAM
		ORDER BY TIME DESC
		LIMIT 25;
	`)
	if err != nil {
		log.Println("Error querying RAM table:", err)
		return nil, err
	}
	defer rows.Close()

	var ramDataHistoric []RamDataHistoric
	for rows.Next() {
		var ramData RamDataHistoric
		var time8 []uint8
		err := rows.Scan(&ramData.ID_RAM, &ramData.PorcentajeUtiizado, &time8)
		if err != nil {
			log.Println("Error scanning RAM table:", err)
			return nil, err
		}
		ramData.Time = string(time8)
		ramDataHistoric = append(ramDataHistoric, ramData)
	}
	return ramDataHistoric, nil
}

func GetListCpuHistoric() ([]CPUDataHistoric, error) {
	rows, err := db.Query(`
		SELECT ID_CPU, CPU_PORCENTAJE, TIME
		FROM CPU
		ORDER BY TIME DESC
		LIMIT 25;
	`)
	if err != nil {
		log.Println("Error querying CPU table:", err)
		return nil, err
	}
	defer rows.Close()

	var cpuDataHistoric []CPUDataHistoric
	for rows.Next() {
		var cpuData CPUDataHistoric
		var time8 []uint8
		err := rows.Scan(&cpuData.ID_CPU, &cpuData.Cpu_porcentaje, &time8)
		if err != nil {
			log.Println("Error scanning CPU table:", err)
			return nil, err
		}
		cpuData.Time = string(time8)
		cpuDataHistoric = append(cpuDataHistoric, cpuData)
	}
	return cpuDataHistoric, nil
}

func InsertStateTask(idTask int, name string) error {
	_, err := db.Exec(`
		INSERT INTO STATE (ID_TASK, NAME)
		VALUES (?, ?);
	`, idTask, name)
	return err
}

// Otros métodos de base de datos aquí...

// No olvides cerrar la conexión a la base de datos al cerrar la aplicación
func CloseDB() {
	if db != nil {
		db.Close()
	}
}

// ================= FUNCIONES PARA PROCESOS =================

// Read the RAM data from the file
func GetDataFile(file string) (string, error) {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("cat %s", file))
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	// Removes whitespace at the beginning and at the end of the string
	result := strings.TrimSpace(string(out))
	return result, nil
}

// Read use of the CPU
func GetCPUUse() (float64, error) {
	cmd := exec.Command("sh", "-c", "top -bn1 | awk '/Cpu/ { cpu = 100 - $8 }; END { print cpu }'")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return 0, err
	}

	// Removes whitespace and newline at the beginning and at the end of the string
	result := strings.TrimSpace(string(out))

	// Convert the string to a float64
	cpuUsage, err := strconv.ParseFloat(result, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting CPU usage to float: %v", err)
	}

	return cpuUsage, nil
}

// ================= FUNCIONES DE CONTROL  =================
// ESTAMOS HABLANDO DE LA CONFIGURACION DE RUTAS Y DE LOS HANDLERS

// Read the CPU data from the file
func UpdateCPU() {
	for {
		cpuData, err := GetDataFile(cpuFilePath)
		if err != nil {
			fmt.Println(err)
			continue
		}
		cpuDataChan <- cpuData

		//update every 500 milliseconds
		time.Sleep(time.Second / 2)
	}
}

// Read the RAM data from the file
func UpdateRAM() {
	for {
		ramData, err := GetDataFile(ramFilePath)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ramDataChan <- ramData

		//update every 500 milliseconds
		time.Sleep(time.Second / 2)
	}
}

// HandleCPUData retorna los datos de CPU al endpoint correspondiente
func HandleCPUData(c *fiber.Ctx) error {
	dataCPU := <-cpuDataChan
	//fmt.Println(dataCPU)

	// Deserializar datos de CPU
	var rawCPU CpuData
	if err := json.Unmarshal([]byte(dataCPU), &rawCPU); err != nil {
		log.Println("Error al deserializar datos de CPU:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
	}
	//fmt.Println(rawCPU)

	cpuUsage, err := GetCPUUse()
	if err != nil {
		log.Println("Error al obtener el uso de CPU:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
	}
	rawCPU.Cpu_percentage = cpuUsage

	err = InsertCpu(rawCPU.Cpu_total, cpuUsage, rawCPU.Cpu_porcentaje, rawCPU.Running, rawCPU.Sleeping, rawCPU.Zombie, rawCPU.Stopped, rawCPU.Total)
	if err != nil {
		log.Println("Error al guardar datos de CPU en la base de datos:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
	}

	tasks := rawCPU.Processes
	//fmt.Println("\ntask: ", tasks)
	for _, task := range tasks {
		//Insert task in the database
		taskId, err := InsertTask(task.PID, task.Name, task.State, task.Rss, task.User, task.RAM)
		if err != nil {
			// Handle the error here
			log.Println("Error al guardar datos de procesos en la base de datos:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
		}
		//update id struct task
		task.IdTask = taskId

		taskChild := task.Child
		//Insert child task in the database
		for _, child := range taskChild {
			if err := InsertChildTask(taskId, child.PID, child.Name, child.State, child.RssChild, child.UserChild); err != nil {
				// Handle the error here
				log.Println("Error al guardar datos de procesos hijos en la base de datos:", err)
				return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
			}
		}
	}
	return c.JSON(map[string]interface{}{"informacion_cpu": rawCPU})
}

// HandleRAMData retorna los datos de RAM al endpoint correspondiente
func HandleRAMData(c *fiber.Ctx) error {
	dataRAM := <-ramDataChan
	//fmt.Println(dataRAM)
	// Deserializar datos de RAM
	var rawRAM RawRAM
	if err := json.Unmarshal([]byte(dataRAM), &rawRAM); err != nil {
		log.Println("Error al deserializar datos de RAM:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
	}
	// Convertir a la estructura final
	ram := ConvertToRAM(rawRAM)
	//fmt.Println(ram)

	//Guardar datos de RAM en la base de datos
	if err := InsertRam(ram.TotalRam, ram.Libre,
		ram.MemoriaEnUso, ram.Porcentaje); err != nil {
		log.Println("Error al guardar datos de RAM en la base de datos:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
	}

	return c.JSON(map[string]interface{}{"informacion_ram": ram})
}

func GetRamDataHistoric(c *fiber.Ctx) error {
	ramData, err := GetListTasks()
	if err != nil {
		log.Println("Error al obtener datos de RAM de la base de datos:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
	}
	return c.JSON(map[string]interface{}{"ram_historic": ramData})
}

func GetCpuDataHistoric(c *fiber.Ctx) error {
	cpuData, err := GetListCpuHistoric()
	if err != nil {
		log.Println("Error al obtener datos de CPU de la base de datos:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
	}
	return c.JSON(map[string]interface{}{"cpu_historic": cpuData})
}

func GetListTaskPid(c *fiber.Ctx) error {
	dataCPU := <-cpuDataChan

	// Deserializar datos de CPU
	var dataTasks CpuData
	if err := json.Unmarshal([]byte(dataCPU), &dataTasks); err != nil {
		log.Println("Error al deserializar datos de CPU:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
	}
	fmt.Println("dataTasks: ", dataTasks)

	//map
	mapTasks := make(map[int]bool)
	var listPid []int

	for _, task := range dataTasks.Processes {
		if _, exists := mapTasks[task.PID]; !exists {
			mapTasks[task.PID] = true
			listPid = append(listPid, task.PID)
		}
	}
	fmt.Println("listPid: ", listPid)

	return c.JSON(map[string]interface{}{"list_pid": listPid})

}

func CreateTreeTask(c *fiber.Ctx) error {
	// Log para verificar que la solicitud llega correctamente
	fmt.Println("Solicitud recibida en /generateTreeTasks")

	// Obtén el parámetro de la ruta ":pid"
	pidParam := c.Params("pid")

	// Convierte el parámetro a un número (int)
	pid, err := strconv.Atoi(pidParam)
	if err != nil {
		// Manejo de error si el parámetro no es un número
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "El parámetro 'pid' debe ser un número entero",
		})
	}

	// Haz algo con el número (pid)
	// Aquí puedes utilizar el valor de 'pid' en tu lógica
	// Convierte el número a un string
	pidString := strconv.Itoa(pid)
	treeDot, err := generatedTreeDot(pidString)
	if err != nil {
		// Manejo de error si no se pudo generar el árbol
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo generar el árbol de procesos",
		})
	}

	// Respuesta de ejemplo
	return c.JSON(map[string]interface{}{"treeDot": treeDot})
}

func generatedTreeDot(pid string) (string, error) {
	dataCPU := <-cpuDataChan

	var dataTasks CpuData
	if err := json.Unmarshal([]byte(dataCPU), &dataTasks); err != nil {
		log.Println("Error al deserializar datos de CPU:", err)
		return "", err
	}

	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		log.Println("Error al convertir el PID a entero:", err)
		return "", err
	}

	var nameTaskParent string
	var taskSelected *TasksParent
	for _, task := range dataTasks.Processes {
		if task.PID == pidInt {
			taskSelected = &task
			nameTaskParent = task.Name
			break
		}
	}

	if taskSelected == nil {
		return "", fmt.Errorf("No se encontró el proceso con PID %s", pid)
	}
	treeDot := generatedGraphvizDot(taskSelected)

	treeDot = fmt.Sprintf(`
	digraph G {
		label="process parent: %s";
		rankdir=TB;
		node [shape = record, color=blue ,style="rounded,filled", fillcolor=gray93];
		%s
	}
	`, nameTaskParent, treeDot)
	//treeDot = "digraph G {\n" + treeDot + "}"

	return treeDot, nil
}

func generatedGraphvizDot(task *TasksParent) string {
	// Create the root node
	dot := fmt.Sprintf("  %d [label=\"%s \n pid= %d \"];\n", task.PID, task.Name, task.PID)

	// Create the child nodes
	for _, child := range task.Child {
		dot += fmt.Sprintf("  %d [label=\"%s \n pid= %d\"];\n", child.PID, child.Name, child.PID)
	}

	// Create the edges
	for _, child := range task.Child {
		dot += fmt.Sprintf("  %d -> %d;\n", task.PID, child.PID)
	}

	return dot
}

func ControllerStartTask(c *fiber.Ctx) error {
	cmd := exec.Command("sleep", "infinity")
	err := cmd.Start()
	if err != nil {
		log.Println("Error al iniciar el proceso:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
	}

	tasks = cmd
	pid := tasks.Process.Pid
	err = InsertStateTask(pid, "NEW")
	if err != nil {
		log.Println("Error al guardar datos de procesos (new) en la base de datos:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
	}

	err = InsertStateTask(pid, "READY")
	if err != nil {
		log.Println("Error al guardar datos de procesos (ready) en la base de datos:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
	}

	err = InsertStateTask(pid, "RUNNING")
	if err != nil {
		log.Println("Error al guardar datos de procesos (running) en la base de datos:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
	}

	return c.SendString("Proceso iniciado correctamente con PID: " + strconv.Itoa(pid) + "\n")
}

func ControllerStopTask(c *fiber.Ctx) error {
	pidParam := c.Query("pid")
	if pidParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "El parámetro 'pid' es requerido en la query",
		})
	}

	pid, err := strconv.Atoi(pidParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "El parámetro 'pid' debe ser un número entero",
		})
	}

	cmd := exec.Command("kill", "-SIGSTOP", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		log.Println("Error al detener el proceso:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
	}

	err = InsertStateTask(pid, "READY")
	if err != nil {
		log.Println("Error al guardar datos de procesos (ready) en la base de datos:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
	}

	return c.SendString("Proceso detenido correctamente con PID: " + strconv.Itoa(pid) + "\n")
}

func ControllerReadyTask(c *fiber.Ctx) {
	pidParam := c.Query("pid")
	if pidParam == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "El parámetro 'pid' es requerido en la query",
		})
		return
	}

	pid, err := strconv.Atoi(pidParam)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "El parámetro 'pid' debe ser un número entero",
		})
		return
	}

	cmd := exec.Command("kill", "-SIGCONT", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		log.Println("Error al continuar el proceso:", err)
		c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
		return
	}

	err = InsertStateTask(pid, "RUNNING")
	if err != nil {
		log.Println("Error al guardar datos de procesos (running) en la base de datos:", err)
		c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
		return
	}

	c.SendString("Proceso continuado correctamente con PID: " + strconv.Itoa(pid) + "\n")
}

func ControllerKillTask(c *fiber.Ctx) {
	pidParam := c.Query("pid")
	if pidParam == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "El parámetro 'pid' es requerido en la query",
		})
		return
	}

	pid, err := strconv.Atoi(pidParam)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "El parámetro 'pid' debe ser un número entero",
		})
		return
	}

	cmd := exec.Command("kill", "-9", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		log.Println("Error al matar el proceso:", err)
		c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
		return
	}

	err = InsertStateTask(pid, "TERMINATED")
	if err != nil {
		log.Println("Error al guardar datos de procesos (terminated) en la base de datos:", err)
		c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
		return
	}

	c.SendString("Proceso matado correctamente con PID: " + strconv.Itoa(pid) + "\n")
}
