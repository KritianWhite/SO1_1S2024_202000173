import React, { useState, useEffect } from "react";
import Viz from 'viz.js';

import Navigator from "../components/navigator";
import Head from "../components/head";
import TreeGraph from "../components/treeGraph";

export default function Tree() {

    const [options, setOptions] = useState([]);
    const [value, setValue] = useState(options[0]);
    const [treeDot, setTreeDot] = useState('');
    const serverUrl = "/api";

    const handleChange = async (selectedValue) => {
        setValue(selectedValue);
        console.log(`Option selected:`, selectedValue);
        var numericValue = parseInt(selectedValue.value, 10);
        // Verificación de si la conversión fue exitosa
        if (!isNaN(numericValue)) {
            // Haz algo con el valor numérico
            console.log("Número extraído:", numericValue);
        }

        try {
            const response = await fetch(`${serverUrl}/generateTreeTasks/${numericValue}`,
                { method: 'GET', });
            if (response.ok) {
                const data = await response.json();
                //console.log('data', data);
                setTreeDot(data.treeDot);
            }

        } catch (error) {
            console.error('Error al obtener datos:', error);
        }
    };

    const customStyles = {
        option: (defaultStyles, state) => ({
            // You can log the defaultStyles and state for inspection
            // You don't need to spread the defaultStyles
            ...defaultStyles,
            color: state.isSelected ? "#212529" : "#fff",
            backgroundColor: state.isSelected ? "#a0a0a0" : "#212529",
        }),

        control: (defaultStyles) => ({
            ...defaultStyles,
            // Notice how these are all CSS properties
            backgroundColor: "#212529",
            padding: "10px",
            border: "none",
            boxShadow: "none",
        }),
        singleValue: (defaultStyles) => ({ ...defaultStyles, color: "#fff" }),
    };

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch(`${serverUrl}/tasks_ids`);
                const data = await response.json();

                // Transformar el array de tasks en el formato requerido para Select
                const transformedOptions = data.list_pid.map(task => ({
                    value: `${task}`,
                    label: `pid ${task}`  // Personaliza según tus necesidades
                }));
                setOptions(transformedOptions);
                //console.log('transformedOptions', transformedOptions);

                // Agrega un console.log para mostrar la ejecución de la petición
                //console.log('Petición ejecutada a las', new Date());
            } catch (error) {
                console.error('Error al obtener datos:', error);
            }
        };
        fetchData();

        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []); // El segundo parámetro [] asegura que useEffect se ejecute solo una vez al montar el componente


    return (
        <>
            <div className="view-port">
                <Navigator />
                <div className="content">
                    <Head />
                    <div className="container">
                        <div className="dashboard">
                            <h1>TREE</h1>
                            <div className="metrics">
                                <TreeGraph options={options} value={value} handleChange={handleChange} treeDot={treeDot}  />
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}