import React, { useEffect, useRef } from 'react';
import { DataSet, Network } from 'vis-network/standalone/esm/vis-network';

const ProcessDiagram = ({ currentHighlight }) => {
    const container = useRef(null);
    const networkRef = useRef(null);

    useEffect(() => {
        if (container.current && !networkRef.current) {
            const nodes = new DataSet([
                { id: 'new', label: 'New', color: 'lightblue' }, // Color celeste para 'new'
                { id: 'ready', label: 'Ready', color: 'lightblue' }, // Color celeste para 'ready'
                { id: 'running', label: 'Running', color: 'lightgreen' }, // Color verde para 'running'
                { id: 'kill', label: 'Kill', color: 'red' }, // Color rojo para 'kill'
            ]);

            const edges = new DataSet([
                { id: 'new-ready', from: 'new', to: 'ready', color: 'lightgray', arrows: 'to' },
                { id: 'ready-running', from: 'ready', to: 'running', color: 'lightgray', arrows: 'to' },
                { id: 'running-ready', from: 'running', to: 'ready', color: 'lightgray', arrows: 'to' },
            ]);

            const data = { nodes, edges };
            const options = {};
            networkRef.current = new Network(container.current, data, options);
        }
    }, []);

    useEffect(() => {
        if (networkRef.current) {
            const { nodes, edges } = networkRef.current.body.data;

            // Reset all nodes and edges to their initial color before highlighting the current
            // Este paso es importante para garantizar que los nodos y aristas vuelvan a su color base.
            nodes.get().forEach(node => {
                let color;
                switch (node.id) {
                    case 'new':
                    case 'ready':
                        color = 'lightblue';
                        break;
                    case 'running':
                        color = 'lightgreen';
                        break;
                    case 'kill':
                        color = 'red';
                        break;
                    default:
                        color = 'lightgray'; // Por defecto o para cualquier otro caso
                }
                nodes.update({ id: node.id, color });
            });

            // Solo se actualiza el color de resaltado para el nodo o arista actual sin alterar los colores base definidos anteriormente
            if (currentHighlight.node) {
                nodes.update({ id: currentHighlight.node, color: { background: 'orange' } });
            }
            if (currentHighlight.edge) {
                edges.update({ id: currentHighlight.edge, color: 'red' });
            }
        }
    }, [currentHighlight]);

    return (
        <>
            <div className='metric'>
                <div className='graphic'>
                    <div ref={container} style={{ height: '500px', width: '100%' }} />
                </div>
            </div>
        </>
    );
};


export default ProcessDiagram;