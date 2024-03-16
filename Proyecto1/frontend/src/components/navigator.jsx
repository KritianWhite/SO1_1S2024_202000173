import React from "react";
import { Link } from "react-router-dom";

import './styles/navigator.css';

export default function Navigator() {

    return (
        <>
            <div id="sidebar">
                <header>
                    <a>Project 1 - SO1</a>
                </header>
                <ul className="nav">
                    <li>
                        <Link to="/" className="zmdi zmdi-view-dashboard">Dashboard</Link>
                    </li>
                    <li>
                        <Link to="/historic" className="zmdi zmdi-link">Historic</Link>
                    </li>
                    <li>
                        <Link to="/tree" className="zmdi zmdi-link">Tree</Link>
                    </li>
                    <li>
                        <Link to="/diagramProcess" className="zmdi zmdi-link">Diagram of process state</Link>
                    </li>
                    <li>
                        <a href="#">
                            <i className="zmdi zmdi-info-outline"></i> About
                        </a>
                    </li>
                    <li>
                        <a href="#">
                            <i className="zmdi zmdi-settings"></i> Services
                        </a>
                    </li>
                    <li>
                        <a href="#">
                            <i className="zmdi zmdi-comment-more"></i> Contact
                        </a>
                    </li>
                </ul>
            </div>
        </>
    );

}