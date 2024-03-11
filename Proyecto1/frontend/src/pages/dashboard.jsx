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
                    <div className="container-fluid">
                        <Graph title={"RAM Percentage"}/>
                        <Graph title={"CPU Percentage"}/>
                    </div>
                </div>
            </div>
        </>
    );
}