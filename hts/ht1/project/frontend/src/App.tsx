import React, {useState, useEffect} from 'react';
import {Chart as ChartJS, ArcElement, Tooltip, Legend} from "chart.js";
import {Doughnut} from "react-chartjs-2";
import {Ram} from "../wailsjs/go/main/App";
import './App.css';

ChartJS.register(ArcElement, Tooltip, Legend)

function App() {
    const [ramFree, setRamFree] = useState(0);
    const [ramUse, setRamUse] = useState(0);
    async function fetchRam() {
        try {
            const ram = await Ram();
            setRamFree(ram);
            setRamUse(50000-ram);
        } catch (error) {
            console.error("Error al obtener la RAM:", error);
        }
    }

    // Gráfica de dona
    const data = {
        labels: [ 'Libre', 'En Uso' ],
        datasets: [
            {
                label: "USO DE RAM",
                data: [50, 60],
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

    useEffect(() => {
        const intervalo = setInterval(() =>{
            fetchRam();
        }, 500)
        return clearInterval(intervalo);
    }, []);

    return (
        <div>
            <h1>Ram libre: {ramFree}</h1>
            <h1>Ram ocupada: {ramUse}</h1>
            <Doughnut options={options} data={data} />
        </div>
    );
}

export default App