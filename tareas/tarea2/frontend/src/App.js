import React, { useRef, useState, useEffect } from 'react';
import './App.css';

function CameraApp() {
  const videoRef = useRef(null);
  const canvasRef = useRef(null);
  const [capturing, setCapturing] = useState(false);

  useEffect(() => {
    getVideoStream();
    // Asegurarse de detener el stream de video cuando el componente se desmonte
    return () => {
      if (videoRef.current && videoRef.current.srcObject) {
        const tracks = videoRef.current.srcObject.getTracks();
        tracks.forEach(track => track.stop());
      }
    };
  }, []);

  const getVideoStream = async () => {
    try {
      const videoStream = await navigator.mediaDevices.getUserMedia({ video: true });
      videoRef.current.srcObject = videoStream;
      setCapturing(true);
    } catch (error) {
      console.error("Error accessing the camera", error);
    }
  };

  const takePhoto = () => {
    const video = videoRef.current;
    const canvas = canvasRef.current;
    const context = canvas.getContext('2d');
    context.drawImage(video, 0, 0, canvas.width, canvas.height);

    // Convertir la imagen a base64
    const imageDataUrl = canvas.toDataURL('image/png');

    // Extraer la parte base64 de la cadena de datos de la imagen
    const base64Image = imageDataUrl.split(';base64,').pop();

    // Enviar la imagen al servidor
    fetch('http://localhost:8000/upload', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ image: base64Image }),
    })
      .then(response => response.text())
      .then(data => {
        console.log('Success:', data);
      })
      .catch((error) => {
        console.error('Error:', error);
      });
  };

  const getPhoto = () => {
    fetch('http://localhost:8000/latest-image')
      .then(response => response.json())
      .then(data => {
        let image = document.getElementById('latest-image');
        // Si no existe el elemento de imagen, lo crea y le asigna un id
        if (!image) {
          image = new Image();
          image.id = 'latest-image'; // Asigna un id al elemento de imagen
          image.className = 'latest-image';
          document.body.appendChild(image);
        }
        // Actualiza el atributo src de la imagen existente
        image.src = `data:image/png;base64,${data.image}`;
      })
      .catch((error) => {
        console.error('Error:', error);
      });
  };
  

  return (
    <div>
      <video ref={videoRef} autoPlay hidden={!capturing}></video>
      <canvas ref={canvasRef} hidden={capturing}></canvas>
      <button onClick={takePhoto} disabled={!capturing}>Tomar Foto</button>
      <button onClick={getPhoto}>Obtener Foto</button>
    </div>
  );
}

export default CameraApp;
