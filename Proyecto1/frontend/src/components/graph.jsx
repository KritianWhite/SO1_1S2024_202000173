import React from 'react';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Doughnut } from "react-chartjs-2";

import './styles/graph.css';

ChartJS.register(ArcElement, Tooltip, Legend);

export default function Pies({ title, percentageFree, percentageOcupied }) {

    var options = {
        responsive: true,
        maintainAspectRatio: false,
    };

    var data = {
        labels: ['Percentage free', 'Percentage ocupied'],
        datasets: [
            {
                label: 'Popularidad en Navidad',
                data: [percentageFree, percentageOcupied],
                backgroundColor: [
                    'rgba(255, 99, 132, 0.2)',
                    'rgba(75, 192, 192, 0.2)',
                    'rgba(153, 102, 255, 0.2)',
                ],
                borderColor: [
                    'rgba(54, 162, 235, 1)',
                    'rgba(75, 192, 192, 1)',
                    'rgba(153, 102, 255, 1)',
                ],
                borderWidth: 1,
            },
        ],
    };

    return (
        <>
            <div className='metric'>
                <div className='graphic'>
                <h2>{title}</h2>
                    <Doughnut data={data} options={options} />
                </div>
            </div>
        </>
    )
}