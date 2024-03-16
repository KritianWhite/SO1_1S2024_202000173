import React from "react";
import { Line } from 'react-chartjs-2';
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
    Filler,
} from 'chart.js';

import './styles/graph.css';


export default function HistoricGraph({ title, configuration}) {
    ChartJS.register(
        CategoryScale,
        LinearScale,
        PointElement,
        LineElement,
        Title,
        Tooltip,
        Legend,
        Filler
    );

    var misoptions = {
        scales: {
            y: {
                min: 0
            },
            x: {
                ticks: { color: 'rgb(255, 99, 132)' }
            }
        }
    };

    

    return (
        <>
            <div className='metric'>
                <div className='graphic'>
                    <h2>{title}</h2>
                    <Line data={configuration} options={misoptions} />
                </div>
            </div>
        </>
    )
}