import React, { useState } from "react";

import Navigator from "../components/navigator";
import Head from "../components/head";
import ProcessDiagram from "../components/diagramProcess";



export default function ProcessState() {
    const [currentHighlight, setCurrentHighlight] = useState({ node: '', edge: '' });

    const handleNewProcess = () => {
        // Highlight the "New" node first
        setCurrentHighlight({ node: 'new', edge: '' });

        setTimeout(() => {
            // Animate from "New" to "Ready"
            setCurrentHighlight({ node: 'ready', edge: 'new-ready' });

            setTimeout(() => {
                // Then animate from "Ready" to "Running" and keep it highlighted
                setCurrentHighlight({ node: 'running', edge: 'ready-running' });
            }, 1500); // Delay for the second part of the transition animation
        }, 1500); // Delay to highlight the "New" node before starting the transition
    };

    const handleStopProcess = () => {
        // Animate from "Running" to "Ready" and keep it highlighted
        setCurrentHighlight({ node: 'ready', edge: 'running-ready' });
    };

    const handleResumeProcess = () => {
        // Animate from "Ready" to "Running" and keep it highlighted
        setCurrentHighlight({ node: 'running', edge: 'ready-running' });
    }

    const handleKillProcess = () => {
        // Animate from "Running" to "Kill" and keep it highlighted
        setCurrentHighlight({ node: 'kill', edge: 'running-kill' });
    }

    return (
        <>
            <div className="view-port">
                <Navigator />
                <div className="content">
                    <Head />
                    <div className="container">
                        <div className="dashboard">
                            <h1>DIAGRAM OF PROCESS STATE</h1>
                            <div className="metrics">
                                <div>
                                    <button onClick={handleNewProcess}>New</button>
                                    <button onClick={handleStopProcess}>Stop</button>
                                    <button onClick={handleResumeProcess}>Resume</button>
                                    <button onClick={handleKillProcess}>Kill</button>
                                    <ProcessDiagram currentHighlight={currentHighlight} />
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

        </>
    );
}
