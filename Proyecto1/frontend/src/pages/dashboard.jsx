import React, { useState, useEffect } from "react";
import Navigator from "../components/navigator";
import Head from "../components/head";
import DashboardGraph from "../components/dashboardGraph";

import './styles/dashboard.css';


export default function Dashboard() {

    const [totalRam, setTotalRam] = useState(0);
    const [memoriaEnUso, setMemoriaEnUso] = useState(0);
    const [porcentaje, setPorcentaje] = useState(0);
    const [libre, setLibre] = useState(0);

    const [porcentajeCPU, setPorcentajeCPU] = useState(0);
    const serverUrl = "/api";

    const getRamUsage = async () => {
        //const serverUrl = "http://localhost:5000";
        //await fetch(`${serverUrl}/${value}ram`);
        try {
            const response = await fetch(`${serverUrl}/ram`, {
                method: 'GET',
            });
            if (response.ok) {
                const data = await response.json();
                setTotalRam(data.informacion_ram.total_memoria);
                setMemoriaEnUso(data.informacion_ram.memoria_utilizada);
                setPorcentaje(data.informacion_ram.porcentaje_utilizado);
                setLibre(data.informacion_ram.memoria_libre);
            } else {
                console.error('Error en la respuesta del servidor:', response.status, response.statusText);
            }
        } catch (error) {
            console.error(error);
        }
    };

    const getCPUUsage = async () => {
        //await fetch(`${serverUrl}/${value}ram`);
        try {
            const response = await fetch(`${serverUrl}/cpu`, {
                method: 'GET',
            });
            if (response.ok) {
                const data = await response.json();
                setPorcentajeCPU(data.informacion_cpu.Cpu_porcentaje);
            } else {
                console.error('Error en la respuesta del servidor:', response.status, response.statusText);
            }
        } catch (error) {
            console.error(error);
        }
    };

    useEffect(() => {
        const interval = setInterval(() => {
            getRamUsage();
        }, 500);

        return () => clearInterval(interval);
    }, []);

    useEffect(() => {
        const interval = setInterval(() => {
            getCPUUsage();
        }, 5000);

        return () => clearInterval(interval);
    }, []);

    return (
        <>
            <div className="view-port">
                <Navigator />
                <div className="content">
                    <Head />
                    <div className="container">
                        <div className="dashboard">
                            <h1>DASHBOARD</h1>
                            <div className="metrics">
                                <DashboardGraph title={"RAM Percentage"} label={"RAM GB"} percentageFree={libre} percentageOcupied={memoriaEnUso} />
                                <DashboardGraph title={"CPU Percentage"} label={"CPU percentage"} percentageFree={100 - porcentajeCPU} percentageOcupied={porcentajeCPU} />
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
}