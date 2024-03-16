import React, { useState, useEffect } from "react";

import Navigator from "../components/navigator";
import Head from "../components/head";
import HisctoriGraph from "../components/historicGraph";

export default function Historic() {
    const serverUrl = "/api";
    const [dataCPU, setDataCPU] = useState({ list_cpu: [] });
    const [dataRAM, setDataRAM] = useState({ list_ram: [] });
    const [cpuGraphic, setCpuGraphic] = useState(null);
    const [ramGraphic, setRamGraphic] = useState(null);

    const getDataHistoric = async (value) => {
        try {
            const response = await fetch(`${serverUrl}/${value}_historic`);

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }

            const data = await response.json();
            //console.log(`Data ${value} historic:`, data);
            return data; //.ram_historic || data.cpu_historic || null;
        } catch (error) {
            //console.error(`Error get data ${value} historic:`, error);
            return null;
        }
    };

    const convertirFecha = (fechaOriginal) => {
        // Crear un objeto de fecha
        const fecha = new Date(fechaOriginal);

        // Obtener el día, mes y año
        const dia = fecha.getDate();
        const mes = fecha.getMonth() + 1; // ¡Atención! Los meses comienzan desde 0
        const anio = fecha.getFullYear();

        // Formatear la fecha como "dd/mm/yyyy"
        const fechaFormateada = `${dia < 10 ? '0' : ''}${dia}/${mes < 10 ? '0' : ''}${mes}/${anio}`;

        return fechaFormateada;
    };

    const updateRAMUsageHistoric = async () => {
        const dataRam = await getDataHistoric("ram");
        //console.log("Data RAM:", dataRam ? dataRam.ram_historic : null);
        if (dataRam && Array.isArray(dataRam.ram_historic)) {
            setDataRAM({ list_ram: dataRam.ram_historic }); //update graphic ram
        }
    };

    const updateCPUUsageHistoric = async () => {
        const dataCpu = await getDataHistoric("cpu");
        //console.log("Data CPU:", dataCpu ? dataCpu.cpu_historic : null);
        if (dataCpu && Array.isArray(dataCpu.cpu_historic)) {
            setDataCPU({ list_cpu: dataCpu.cpu_historic }); //update graphic cpu
        }
    };

    // NOTA: La razón de que se están utilizando dos useEffect es porque se necesita 
    // que se actualicen los datos de la RAM y CPU por separado ya que al momento de
    // actualizar los datos de RAM y CPU en el mismo useEffect, está realizando un bucle
    // infinito, por lo que se optó por separarlos en dos useEffect.

    // Primer useEffect para obtener los datos de la RAM y CPU
    useEffect(() => {
        updateRAMUsageHistoric();
        updateCPUUsageHistoric();
    }, []);
    
    // Segundo useEffect para actualizar los datos de la RAM y CPU
    useEffect(() => {
        if (dataRAM.list_ram.length > 0 && dataRAM.list_ram && ramGraphic === null) {
            const labels = dataRAM.list_ram.map((item) => convertirFecha(item.time));
            const dataRam = dataRAM.list_ram.map((item) => item.porcentaje_utilizado);
            const ramGraphicConfig = {
                labels,
                datasets: [
                    {
                        label: "RAM Usage Historic",
                        data: dataRam,
                        fill: true,
                        backgroundColor: "rgba(75,192,192,0.2)",
                        borderColor: "rgba(75,192,192,1)",
                    },
                ],
            };
            setRamGraphic(ramGraphicConfig);
        }

        if (dataCPU.list_cpu.length > 0 && dataCPU.list_cpu && cpuGraphic === null) {
            const labels = dataCPU.list_cpu.map((item) => convertirFecha(item.time));
            const dataCpu = dataCPU.list_cpu.map((item) => item.porcentaje_utilizado);
            const cpuGraphicConfig = {
                labels,
                datasets: [
                    {
                        label: "CPU Usage Historic",
                        data: dataCpu,
                        fill: true,
                        backgroundColor: "rgba(75,192,192,0.2)",
                        borderColor: "rgba(75,192,192,1)",
                    },
                ],
            };
            setCpuGraphic(cpuGraphicConfig);
        }
    }, [dataRAM, ramGraphic, dataCPU, cpuGraphic]);

    return (
        <>
            <div className="view-port">
                <Navigator />
                <div className="content">
                    <Head />
                    <div className="container">
                        <div className="dashboard">
                            <h1>HISTORIC</h1>
                            <div className="metrics">
                                {ramGraphic && (
                                    <HisctoriGraph title="RAM" configuration={ramGraphic} />
                                )}
                                {cpuGraphic && (
                                    <HisctoriGraph title="CPU" configuration={cpuGraphic} />
                                )}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
}
