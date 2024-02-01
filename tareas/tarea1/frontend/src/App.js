import React, { useState } from 'react';
import './App.css';

function App() {
  const [mensaje, setMensaje] = useState('');
  const [mostrarDatos, setMostrarDatos] = useState(false);

  const handleClick = async () => {
    try {
      // Realizar una solicitud al servidor Go para obtener el mensaje
      const response = await fetch('http://localhost:8080/data');
      const data = await response.text();

      // Actualizar el estado del mensaje
      setMensaje(data);

      // Alternar el estado de mostrarDatos
      setMostrarDatos(!mostrarDatos);
    } catch (error) {
      console.error('Error al obtener el mensaje:', error);
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <button className='btn-mostrar-ocultar' onClick={handleClick}>
          {mostrarDatos ? 'Ocultar datos' : 'Mostrar datos'}
        </button>
        {mostrarDatos && <h1>{mensaje}</h1>}
      </header>
    </div>
  );
}

export default App;
