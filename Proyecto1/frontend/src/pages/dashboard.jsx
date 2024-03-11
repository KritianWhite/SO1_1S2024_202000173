import React from "react";
import Navigator from "../components/navigator";
import Head from "../components/head";
import Graph from "../components/graph";
import './styles/dashboard.css';


export default function Dashboard() {



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
                                <Graph title={"RAM Percentage"} percentageFree={30} percentageOcupied={70} />
                                <Graph title={"CPU Percentage"} percentageFree={70} percentageOcupied={30} /> 
                            </div>
                        </div>

                    </div>
                </div>
            </div>
        </>
    );
}