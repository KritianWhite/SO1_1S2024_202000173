import React, { useState, useEffect } from "react";
import Select from 'react-select'
import Graphviz from 'graphviz-react';


export default function TreeGraph({ options, value, handleChange, treeDot }) {




    return (
        <>
            <div className='metric'>
                <div className='graphic'>
                    <Select
                        options={options}
                        placeholder="Select pid ..."
                        autoFocus value={value}
                        onChange={handleChange}
                    />
                    {/* {error && <p className={styles.error}>{error}</p>} */}
                    {treeDot && (
                        <Graphviz
                            dot={treeDot}
                            options={{
                                zoom: true, // Permite hacer zoom
                                fit: true, // Hace que el gráfico se ajuste al contenedor
                                center: true, // Centra el gráfico en el 
                                width: 500, // Ancho del gráfico en px
                                height: 200 // Alto del gráfico en px
                            }}
                        />
                    )}

                </div>
            </div>
        </>
    )

}