import { useState } from 'react'
import { Route, Routes } from "react-router-dom";
import Login from "./pages/Login";
import Estudiante from "./pages/Estudiante";
import Tutor from "./pages/Tutor";
import Administrador from "./pages/Administrador";
import TablaAlumnos from "./pages/TablaAlumnos";
import LibrosAdmin from "./pages/LibrosAdmin";
//import './App.css'

function App() {
  return (
    <>
      <Routes>
        {/*Gestión de las rutas*/}
        <Route path="/" element={<Login />} />
        <Route path="/principal/estudiante" element={<Estudiante />} />
        <Route path="/principal/tutor" element={<Tutor />} />
        <Route path="/principal/admin" element={<Administrador />} />
        <Route path="/principal/admin/alumnos" element={<TablaAlumnos />} />
        <Route path="/principal/tutor/libro" element={<LibrosAdmin />} />
      </Routes>
    </>
  )
}

export default App