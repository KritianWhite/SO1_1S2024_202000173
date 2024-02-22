# App.go

1. Se define un método llamado `Ram()` que pertenece a una estructura `App`. Este método devuelve un entero que representa la cantidad de RAM.

    ```go
    func (a *App) Ram() int {}
    ```

2. Se establece el nombre del archivo `/proc/ram_202000173`. En sistemas basados en Linux, `/proc` es un directorio especial que contiene información sobre el sistema, y `ram_202000173` es un archivo ficticio que se está utilizando en este caso para representar información de la RAM.
    ```go
    // Nombre del archivo
    archivoProc := "/proc/ram_202000173"
    ```

3. Se crea un nuevo comando ejecutable usando `exec.Command()`. En este caso, el comando es `cat` que se usa para imprimir el contenido del archivo especificado.

    ```go
    // Creación del comando cat
    comando := exec.Command("cat", archivoProc)
    ```

4. Se ejecuta el comando usando `comando.Output()`, que devuelve la salida del comando. Si hay un error al ejecutar el comando, se captura y se imprime un mensaje de error.

    ```go
    // Captura de la salida del comando (Contenido del archivo)[Se obtiene en array de bytes]
    salida, err := comando.Output()
    ```

5. La salida del comando se convierte de un array de bytes a una cadena de texto usando `string(salida)`.

    ```go
    // Captura de la salida del comando (Contenido del archivo)[Se obtiene en array de bytes]
    salida, err := comando.Output()
    if err != nil {
        fmt.Println("Error al acedder a la información de la ram.", err)
        return 0
    }
    // Convertir la salida de bytes a una cadena
    salidaStr := string(salida)
    ```

6. Se convierte la cadena de texto resultante a un número entero utilizando `strconv.Atoi()`. Si hay un error durante la conversión, se captura y se imprime un mensaje de error.

    ```go
    // Parsear la cadena a un número entero
	numero, err := strconv.Atoi(salidaStr)
	if err != nil {
		fmt.Println("Error al convertir la salida a entero:", err)
		return 0
	}
    ```

7. Finalmente, el número entero obtenido se imprime y se devuelve como resultado de la función.


# App.tsx

1. **Importaciones**: El código comienza importando las bibliotecas y componentes necesarios para crear la aplicación. Entre ellos están `React`, `useState`, `useEffect` de React para el manejo de estados y efectos, así como también `ChartJS`, `ArcElement`, `Tooltip`, `Legend` de Chart.js y `Doughnut` de `react-chartjs-2` para crear gráficos circulares.

    ```javascript
    import React, { useState, useEffect } from 'react';
    import {Chart as ChartJS, ArcElement, Tooltip, Legend} from "chart.js";
    import { Ram } from "../wailsjs/go/main/App";
    import {Doughnut} from "react-chartjs-2";
    import './App.css';
    ```


2. **Componente App**: Se define el componente funcional `App`, que es el componente principal de la aplicación.

    ```javascript
    function App() {}
    ```

3. **Inicialización de estado**: Se utilizan los hooks `useState` para definir dos variables de estado: `ramFree` y `ramUsed`, que se utilizarán para almacenar el porcentaje de RAM libre y en uso respectivamente.

    ```javascript
    ChartJS.register(ArcElement, Tooltip, Legend)
    const [ramFree, setRamFree] = useState(0);
    const [ramUsed, setRamUsed] = useState(0);
    ```

4. **fetchRamInfo()**: Esta es una función asincrónica que se encarga de obtener la información de la RAM. En este caso, parece estar obteniendo la información de una función `Ram` (que supongo que está definida en otro lugar de la aplicación). Después de obtener la información, calcula el porcentaje de RAM libre y en uso, y actualiza los estados correspondientes.

    ```javascript
    const fetchRamInfo = async () => {
        try {
            const ramInfo = await Ram();
            const ramTotal = 650690;

            // Calculamos los porcentajes
            const ramFreePor = ((ramInfo / ramTotal) * 100).toFixed(2);
            const ramUsedPor = (((ramTotal - ramInfo) / ramTotal) * 100).toFixed(2);

            // Establecemos los estados con los porcentajes limitados
            setRamFree(parseFloat(ramFreePor));
            setRamUsed(parseFloat(ramUsedPor));
        } catch (error) {
            console.error("Error al obtener información de la RAM:", error);
        }
    };
    ```

5. **Gráfica de dona**: Se define el objeto `data` que contiene la configuración de la gráfica de dona. Aquí se especifican las etiquetas de las secciones de la dona (`labels`), los datos correspondientes a cada sección (`data`), y el color y borde de las secciones.

    ```javascript
    // Gráfica de dona
    const data = {
        labels: [ 'Libre', 'En Uso' ],
        datasets: [
            {
                label: "USO DE RAM",
                data: [ramFree, ramUsed],
                backgroundColor: [
                    'rgba(54, 162, 235, 0.9)',
                    'rgba(255, 99, 132, 0.9)',
                ],
                borderColor: [
                    'rgba(54, 162, 235, 1)',
                    'rgba(255, 99, 132, 1)',
                ],
                borderWidth: 1,
            },
        ],
    };
    const options = {
        cutout: 50
    };
    ```

6. **useEffect()**: Se utiliza `useEffect` para ejecutar la función `fetchRamInfo()` cuando el componente se monta por primera vez (es decir, cuando se renderiza por primera vez). Se pasa un arreglo vacío como segundo argumento para indicar que este efecto no depende de ninguna variable y solo debe ejecutarse una vez.

    ```javascript
    useEffect(() => {
        const interval = setInterval(() =>{
            fetchRamInfo();
        }, 500)
        return () => clearInterval(interval)
    }, []);
    ```

7. **Retorno del componente**: Finalmente, en el retorno del componente, se muestra el porcentaje de RAM libre y en uso, y se renderiza el gráfico de dona utilizando el componente `Doughnut` de `react-chartjs-2`, pasándole las opciones y los datos definidos anteriormente.

    ```javascript
    return (
        <>
            <h1>RAM Libre: {ramFree}%</h1>
            <h1>RAM en uso: {ramUsed}%</h1>
            <div className="graphic">
                <Doughnut options={options} data={data} />
            </div>
        </>
    );
    ```
        
## [LINK DEL VIDIO DE YUTUB](https://youtu.be/GWWzC9Nn9no)