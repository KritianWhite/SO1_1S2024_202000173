import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

import Dashboard from './pages/dashboard';
import Historic from './pages/historic';
import Tree from './pages/tree';
import DiagramProcess from './pages/processState';
import './App.css';

function App() {
  return (
    <>
      <Router>
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/historic" element={<Historic />} />
          <Route path="/tree" element={<Tree />} />
          <Route path="/diagramProcess" element={<DiagramProcess />} />
        </Routes>
      </Router>
    </>
  );
}

export default App;
