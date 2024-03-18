# PROYECTO 1 - SISTEMAS OPERATIVOS 1
Christian Alessander Blanco González - 202000173


## 1. MODULOS DE KERNEL

### MODULO PARA EL CPU

1. Inclusión de encabezados:
   - `linux/module.h`: Proporciona funciones y macros para trabajar con módulos del kernel.
   - `linux/init.h`: Define las macros para las funciones de inicialización y limpieza del módulo.
   - `linux/proc_fs.h`: Permite crear y administrar archivos en el sistema de archivos `/proc`.
   - `linux/sched/signal.h`: Proporciona funciones y estructuras para trabajar con señales y procesos.
   - `linux/seq_file.h`: Proporciona funciones para escribir en archivos de tipo `seq_file`.
   - `linux/fs.h`, `linux/sched.h`, `linux/mm.h`: Encabezados para funciones y estructuras relacionadas con el sistema de archivos, la planificación de procesos y la gestión de memoria.

2. Declaración del módulo:
   - `MODULE_LICENSE("GPL")`: Especifica la licencia del módulo.
   - `MODULE_AUTHOR`, `MODULE_DESCRIPTION`, `MODULE_VERSION`: Metadatos del módulo.

3. Función `calcularPorcentajeCpu`:
   - Abre y lee el archivo `/proc/stat` para obtener información sobre el uso de la CPU y calcular el porcentaje de uso.

4. Función `escribir_a_proc`:
   - Implementa la escritura de información en el archivo `/proc/cpu_so1_1s2024`.
   - Calcula el tiempo total de CPU, el porcentaje de uso de CPU, y recopila información sobre los procesos en el sistema.

5. Funciones `abrir_aproc`, `archivo_operaciones`, `modulo_init`, `modulo_cleanup`:
   - `abrir_aproc`: Abre el archivo `/proc/cpu_so1_1s2024`.
   - `archivo_operaciones`: Define las operaciones de archivo para el módulo.
   - `modulo_init`: Función de inicialización del módulo, crea el archivo `/proc/cpu_so1_1s2024`.
   - `modulo_cleanup`: Función de limpieza del módulo, elimina el archivo `/proc/cpu_so1_1s2024`.

### MODULO PARA LA RAM

1. Inclusión de encabezados:
   - `linux/module.h`: Proporciona funciones y macros para trabajar con módulos del kernel.
   - `linux/proc_fs.h`: Permite crear y administrar archivos en el sistema de archivos `/proc`.
   - `linux/sysinfo.h`: Proporciona funciones para obtener información del sistema, como la cantidad de memoria RAM.
   - `linux/seq_file.h`: Proporciona funciones para escribir en archivos de tipo `seq_file`.
   - `linux/mm.h`: Encabezados para funciones y estructuras relacionadas con la gestión de memoria del kernel.

2. Declaración del módulo:
   - `MODULE_LICENSE("GPL")`: Especifica la licencia del módulo.
   - `MODULE_AUTHOR`, `MODULE_DESCRIPTION`, `MODULE_VERSION`: Metadatos del módulo.

3. Declaración de estructuras:
   - `struct sysinfo inf;`: Estructura que contendrá la información del sistema, incluida la información sobre la memoria RAM.

4. Función `escribir_a_proc`:
   - Utiliza la función `si_meminfo(&inf)` para obtener información sobre la memoria del sistema y calcular el uso de la memoria.
   - Calcula el porcentaje de memoria utilizada, la memoria total, la memoria en uso y la memoria libre.
   - Escribe esta información en el archivo `/proc/ram_so1_1s2024` utilizando `seq_printf`.

5. Función `abrir_aproc`:
   - Abre el archivo `/proc/ram_so1_1s2024` para escritura.

6. Operaciones de archivo `archivo_operaciones`:
   - Define las operaciones de apertura y lectura del archivo `/proc/ram_so1_1s2024`.

7. Funciones de inicialización y limpieza del módulo:
   - `modulo_init`: Función de inicialización del módulo, crea el archivo `/proc/ram_so1_1s2024`.
   - `modulo_cleanup`: Función de limpieza del módulo, elimina el archivo `/proc/ram_so1_1s2024`.

### Makefile

1. `obj-m += ram_so1_1s2024.o`: Esta línea indica al Makefile que incluya el archivo `ram_so1_1s2024.o` como un objeto a compilar. Este archivo probablemente contiene el código fuente del módulo del kernel que estás compilando.

2. `all:`: Esta es una regla en el Makefile que define cómo compilar todos los objetivos. En este caso, ejecuta el comando `make` en el directorio `/lib/modules/$(shell uname -r)/build` con la opción `M=$(PWD)` para compilar los módulos del kernel en el directorio actual (`$(PWD)` es una variable que representa el directorio actual).

3. `clean:`: Esta es una regla en el Makefile que define cómo limpiar los archivos generados durante la compilación. Ejecuta el comando `make` en el directorio `/lib/modules/$(shell uname -r)/build` con la opción `M=$(PWD)` y el argumento `clean` para limpiar los archivos generados durante la compilación.


## 2. ENDPOINTS (BACKEND)

1. **Función `main()`**:
   - La función `main()` es el punto de entrada de la aplicación. En este caso, comienza imprimiendo "¡Hola, mundo!" en la consola.
   - Se crea una instancia de la aplicación Fiber usando `fiber.New()`.
   - Se utilizan middleware para la aplicación, como `logger.New()` para el registro de solicitudes y `cors.New()` para el manejo de solicitudes de recursos cruzados (CORS).
   - La función `Connect()` se llama para conectar a la base de datos. Si hay un error al conectarse, se muestra un mensaje de error y se detiene la ejecución con `log.Fatal()`.
   - Se inician dos rutinas (`goroutines`) con `go UpdateCPU()` y `go UpdateRAM()` para actualizar los datos de CPU y RAM de forma asíncrona.
   - Se llama a la función `setupRoutes(app)` para configurar las rutas de la aplicación.
   - Se define el puerto en el que se escucharán las solicitudes (puerto 5000) y se inicia el servidor con `app.Listen()`.

2. **Función `setupRoutes(app *fiber.App)`**:
   - Esta función configura las rutas de la aplicación utilizando el objeto `app` de Fiber que se pasa como argumento.
   - Se definen rutas para diferentes endpoints relacionados con procesos, CPU, RAM y otras funcionalidades de la aplicación.
   - Por ejemplo, `/api/cpu` maneja las solicitudes relacionadas con datos de CPU usando la función `HandleCPUData`.
   - También hay rutas para obtener datos históricos de CPU y RAM, listar IDs de tareas, generar árboles de tareas, y controlar el inicio y la detención de tareas.


## 3. FRONTEND

### DASHBOARD

1. **Definición de Estados:**
   - Se definen varios estados utilizando el hook `useState`, como `totalRam`, `memoriaEnUso`, `porcentaje`, `libre` para el uso de RAM, y `porcentajeCPU` para el uso de CPU.
   - Estos estados se utilizan para almacenar y actualizar la información sobre el uso de recursos del servidor (RAM y CPU) en tiempo real.

2. **Funciones `getRamUsage` y `getCPUUsage`:**
   - `getRamUsage`: Esta función realiza una solicitud GET a la ruta `/api/ram` para obtener datos relacionados con el uso de RAM del servidor. Cuando se recibe la respuesta del servidor, actualiza los estados `totalRam`, `memoriaEnUso`, `porcentaje` y `libre` con la información obtenida.
   - `getCPUUsage`: Similar a `getRamUsage`, esta función realiza una solicitud GET a la ruta `/api/cpu` para obtener datos relacionados con el uso de CPU del servidor. Actualiza el estado `porcentajeCPU` con la información recibida del servidor.

3. **Hooks `useEffect`:**
   - Se utilizan dos efectos secundarios (`useEffect`) para llamar a las funciones `getRamUsage` y `getCPUUsage` en intervalos regulares.
   - El primer efecto se ejecuta al renderizar el componente (`[]` como dependencia), estableciendo un intervalo de 500ms para llamar a `getRamUsage`.
   - El segundo efecto se ejecuta al renderizar el componente, estableciendo un intervalo de 5000ms para llamar a `getCPUUsage`.

4. **Renderizado del Componente:**
   - El componente `Dashboard` renderiza la estructura del dashboard utilizando elementos JSX.
   - Incluye el componente `Navigator` para la navegación, `Head` para la información de encabezado y `DashboardGraph` para mostrar gráficos de uso de RAM y CPU.
   - Los gráficos se alimentan con los datos almacenados en los estados y se actualizan automáticamente debido a los efectos secundarios que llaman a las funciones de obtención de datos en intervalos regulares.


### HISTORIC

1. **Definición de Estados:**
   - Se definen varios estados utilizando el hook `useState`, como `dataCPU`, `dataRAM` para almacenar datos históricos de CPU y RAM respectivamente, y `cpuGraphic`, `ramGraphic` para almacenar la configuración de gráficos.
   - También se define `serverUrl` para la URL base de la API.

2. **Función `getDataHistoric(value)`:**
   - Esta función hace una solicitud GET a la ruta `/api/${value}_historic` para obtener datos históricos de CPU o RAM según el valor de `value`.
   - Si la respuesta es exitosa, actualiza el estado correspondiente (`dataCPU` o `dataRAM`) con los datos recibidos del servidor.

3. **Funciones `updateRAMUsageHistoric` y `updateCPUUsageHistoric`:**
   - Estas funciones utilizan `getDataHistoric` para obtener datos históricos de RAM y CPU respectivamente.
   - Si los datos recibidos son válidos (un array de datos históricos), actualizan los estados `dataRAM` o `dataCPU` con los datos recibidos.

4. **Hooks `useEffect`:**
   - Se utilizan dos efectos secundarios (`useEffect`) para llamar a las funciones de actualización de datos históricos en el momento adecuado.
   - El primer efecto se ejecuta al renderizar el componente (`[]` como dependencia), llamando a las funciones `updateRAMUsageHistoric` y `updateCPUUsageHistoric` para obtener y actualizar los datos históricos de RAM y CPU.
   - El segundo efecto se ejecuta cuando cambian los estados `dataRAM`, `ramGraphic`, `dataCPU`, `cpuGraphic`, y se encarga de generar la configuración de gráficos para los datos históricos de RAM y CPU y actualizar los estados `ramGraphic` y `cpuGraphic` con esta configuración.

5. **Renderizado del Componente:**
   - El componente `Historic` renderiza la estructura de la página utilizando elementos JSX.
   - Incluye el componente `Navigator` para la navegación y `Head` para la información de encabezado.
   - Renderiza gráficos de uso histórico de RAM y CPU utilizando el componente `HistoricGraph` y la configuración almacenada en los estados `ramGraphic` y `cpuGraphic`.

### TREE

1. **Importaciones:**
   - Importa React y los hooks `useState`, `useEffect` para manejar el estado y los efectos secundarios en componentes funcionales.
   - Importa la librería `Viz.js` para generar gráficos de árboles.
   - Importa los componentes `Navigator`, `Head`, y `TreeGraph` para la navegación, encabezado y visualización del árbol de tareas respectivamente.

2. **Definición de Estados:**
   - Se definen varios estados utilizando el hook `useState`, como `options` para las opciones de selección, `value` para el valor seleccionado, y `treeDot` para el contenido del árbol en formato DOT.
   - `serverUrl` se utiliza para la URL base de la API.

3. **Función `handleChange(selectedValue)`:**
   - Esta función se ejecuta cuando se selecciona una opción en un componente de selección (`Select`).
   - Actualiza el estado `value` con el valor seleccionado y realiza una solicitud GET a la API para obtener el contenido del árbol de tareas correspondiente al valor seleccionado.
   - El contenido del árbol en formato DOT se guarda en el estado `treeDot`.

4. **Estilos Personalizados:**
   - Define estilos personalizados para el componente `Select` utilizando la librería `react-select`.

5. **Hook `useEffect`:**
   - Se utiliza `useEffect` para realizar una solicitud GET a la API y obtener las opciones de selección (`tasks_ids`) al montar el componente.
   - Las opciones se transforman en el formato requerido para el componente `Select` y se actualiza el estado `options` con las opciones transformadas.

6. **Renderizado del Componente:**
   - El componente `Tree` renderiza la estructura de la página utilizando elementos JSX.
   - Incluye el componente `Navigator` para la navegación y `Head` para la información de encabezado.
   - Renderiza el componente `TreeGraph` que muestra el árbol de tareas utilizando el estado `options`, `value`, `handleChange`, y `treeDot`.

### PROCESSSTATE

1. **Definición de Estado:**
   - Se define un estado `currentHighlight` utilizando el hook `useState` para controlar qué nodo y borde resaltar en el diagrama de estado del proceso.

2. **Funciones de Manejo de Eventos:**
   - `handleNewProcess`: Al hacer clic en el botón "New", se resalta el nodo "New" en el diagrama y luego se anima la transición al nodo "Ready" y después al nodo "Running".
   - `handleStopProcess`: Al hacer clic en el botón "Stop", se anima la transición del nodo "Running" al nodo "Ready".
   - `handleResumeProcess`: Al hacer clic en el botón "Resume", se anima la transición del nodo "Ready" al nodo "Running".
   - `handleKillProcess`: Al hacer clic en el botón "Kill", se anima la transición del nodo "Running" al nodo "Kill".

3. **Renderizado del Componente:**
   - El componente `ProcessState` renderiza la estructura de la página utilizando elementos JSX.
   - Incluye el componente `Navigator` para la navegación y `Head` para la información de encabezado.
   - Renderiza el componente `ProcessDiagram` que muestra el diagrama de estado del proceso y recibe el estado `currentHighlight` para resaltar los nodos y bordes correspondientes.

## 4. NGINX
```javascript
worker_processes  1;

events {
  worker_connections  1024;
}

http {
  server {
    listen 80;
    server_name  localhost;

    root   /usr/share/nginx/html;
    index  index.html index.htm;
    include /etc/nginx/mime.types;

    gzip on;
    gzip_min_length 5;
    gzip_proxied expired no-cache no-store private auth;
    gzip_types text/plain text/css application/json application/javascript application/x-javascript text/xml application/xml application/xml+rss text/javascript;

    location / {
      try_files $uri $uri/ /index.html;
    }

    location /api/ {
        proxy_pass http://p1_so1_backend:5000;
    }
  }
}
```


1. `worker_processes 1;`: Esta línea especifica el número de procesos de trabajador que se ejecutarán para manejar las solicitudes entrantes. En este caso, se está configurando para que solo haya un proceso de trabajador.

2. `events { worker_connections 1024; }`: Aquí se define la sección de eventos, donde se especifica el número máximo de conexiones simultáneas que puede manejar cada proceso de trabajador. En este caso, se establece en 1024 conexiones.

3. `http { ... }`: Esto indica que se está configurando la sección HTTP del servidor web.

4. `server { ... }`: Dentro de la sección HTTP, se define un bloque `server` que contiene la configuración para un servidor virtual.

5. `listen 80;`: Esta línea indica al servidor que escuche en el puerto 80 para las solicitudes entrantes.

6. `server_name localhost;`: Aquí se especifica el nombre del servidor, en este caso, `localhost`.

7. `root /usr/share/nginx/html;`: Esta línea establece la ruta raíz del servidor web, es decir, la ubicación donde se buscarán los archivos para servir.

8. `index index.html index.htm;`: Aquí se especifica el orden de los archivos que se buscarán cuando se solicite un directorio. En este caso, primero se buscará `index.html` y luego `index.htm`.

9. `include /etc/nginx/mime.types;`: Esta línea incluye el archivo `mime.types` que contiene las asociaciones entre extensiones de archivo y tipos MIME.

10. `gzip on;`: Habilita la compresión gzip para comprimir el contenido antes de enviarlo al cliente, lo que reduce el tamaño de los datos transferidos y mejora el rendimiento.

11. `gzip_min_length 5;`: Especifica la longitud mínima del archivo que se comprimirá usando gzip.

12. `gzip_proxied expired no-cache no-store private auth;`: Define las condiciones bajo las cuales se debe aplicar la compresión gzip.

13. `gzip_types ...;`: Lista los tipos de contenido que serán comprimidos utilizando gzip.

14. `location / { ... }`: Define la configuración para las solicitudes que coincidan con la ruta `/`. En este caso, se intentará servir el archivo solicitado (`$uri`), y si no se encuentra, se redirigirá a `index.html`.

15. `location /api/ { proxy_pass http://p1_so1_backend:5000; }`: Para las solicitudes que coincidan con la ruta `/api/`, se especifica que se deben pasar a un servidor backend en `http://p1_so1_backend` en el puerto 5000.

## DOCKER COMPOSE

1. **Services (Servicios):**
   - `p1_so1_backend`: Este servicio utiliza la imagen `kritianwhite/proyecto1-p1_so1_backend` para crear un contenedor llamado `p1_so1_backend`. Expone el puerto 5000 dentro del contenedor y también en el host (a través de la declaración `ports`). Se reiniciará siempre que se detenga. Está conectado a la red `app_net` y tiene variables de entorno como `NODE_ENV`, `PORT`, `MYSQL_HOST`, `MYSQL_USER`, `MYSQL_PASSWORD`, y `MYSQL_DATABASE`. Dependiendo de `p1_so1_mysqldb`.

   - `p1_so1_frontend`: Utiliza la imagen `kritianwhite/proyecto1-p1_so1_frontend` para crear un contenedor llamado `p1_so1_frontend`. Se reiniciará siempre que se detenga. Expone el puerto 80 en el host y también en el contenedor. Está conectado a la red `app_net` y depende de `p1_so1_backend`.

   - `p1_so1_mysqldb`: Este servicio utiliza la imagen `mysql:latest` para crear un contenedor llamado `p1_so1_mysqldb`. Se reiniciará en caso de fallo. Define algunas configuraciones adicionales a través de `command`, establece variables de entorno como `MYSQL_ROOT_PASSWORD`, `MYSQL_USER`, `MYSQL_PASSWORD`, y `MYSQL_DATABASE`, y expone el puerto 3306 tanto en el host como en el contenedor. Está conectado a la red `app_net` y utiliza volúmenes para almacenar datos persistentes.

2. **Networks (Redes):**
   - `app_net`: Define una red de tipo puente (`bridge`) llamada `app_net` que conecta todos los contenedores para que puedan comunicarse entre sí.

3. **Volumes (Volúmenes):**
   - `p1-so1-mysql-vol`: Define un volumen llamado `p1-so1-mysql-vol` que se utiliza para almacenar datos persistentes del contenedor `p1_so1_mysqldb`. Este volumen está vinculado al directorio `/var/lib/mysql` dentro del contenedor.

