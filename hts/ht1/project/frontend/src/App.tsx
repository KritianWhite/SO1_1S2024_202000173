import React, { useState, useEffect } from 'react';
import {Chart as ChartJS, ArcElement, Tooltip, Legend} from "chart.js";
import { Ram } from "../wailsjs/go/main/App";
import {Doughnut} from "react-chartjs-2";
import './App.css';

function App() {
    ChartJS.register(ArcElement, Tooltip, Legend)
    const [ramFree, setRamFree] = useState(0);
    const [ramUsed, setRamUsed] = useState(0);

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

    useEffect(() => {
        const interval = setInterval(() =>{
            fetchRamInfo();
        }, 500)
        return () => clearInterval(interval)
    }, []);

    return (
        <>
            <h1>RAM Libre: {ramFree}%</h1>
            <h1>RAM en uso: {ramUsed}%</h1>
            <div className="graphic">
                <Doughnut options={options} data={data} />
            </div>
        </>
    );
}

export default App;
