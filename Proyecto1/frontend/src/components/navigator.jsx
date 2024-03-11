import React from "react";


import './styles/navigator.css';

export default function Navigator() {

    return (
        <>
            <div id="sidebar">
                <header>
                    <a>Project 1 - SO1</a>
                </header>
                <ul class="nav">
                    <li>
                        <a href="#">
                            <i class="zmdi zmdi-view-dashboard"></i> Dashboard
                        </a>
                    </li>
                    <li>
                        <a href="#">
                            <i class="zmdi zmdi-link"></i> Shortcuts
                        </a>
                    </li>
                    <li>
                        <a href="#">
                            <i class="zmdi zmdi-widgets"></i> Overview
                        </a>
                    </li>
                    <li>
                        <a href="#">
                            <i class="zmdi zmdi-calendar"></i> Events
                        </a>
                    </li>
                    <li>
                        <a href="#">
                            <i class="zmdi zmdi-info-outline"></i> About
                        </a>
                    </li>
                    <li>
                        <a href="#">
                            <i class="zmdi zmdi-settings"></i> Services
                        </a>
                    </li>
                    <li>
                        <a href="#">
                            <i class="zmdi zmdi-comment-more"></i> Contact
                        </a>
                    </li>
                </ul>
            </div>
        </>
    );

}