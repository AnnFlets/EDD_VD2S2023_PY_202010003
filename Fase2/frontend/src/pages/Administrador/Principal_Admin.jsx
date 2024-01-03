import React from "react";
import "../../styles/otros.css";

function Principal_Admin() {
  const archivos = (e) => {
    e.preventDefault();
    window.open("/principal/admin/archivo", "_self");
  };

  const alumnos = (e) => {
    e.preventDefault();
    window.open("/principal/admin/alumnos", "_self");
  };

  const libros = (e) => {
    e.preventDefault();
    window.open("/principal/admin/libros", "_self");
  };

  const reporte = (e) => {
    e.preventDefault();
    window.open("/principal/admin/reporte", "_self");
  };

  const salir = (e) => {
    e.preventDefault();
    console.log("Listo");
    localStorage.clear();
    window.open("/", "_self");
  };

  return (
    <div className="form-signin1">
      <br/>
      <br/>
      <div className="text-center">
        <form className="card card-body">
        <br/>
          <h1><b>MENÚ ADMINISTRADOR</b></h1>
          <h4 className="h3 mb-3 fw-normal">Bienvenido</h4>
          <br />
          <center>
            <button className="w-50 btn btn-outline-warning" onClick={archivos}>
              Carga Archivos
            </button>
          </center>
          <br />
          <center>
            <button className="w-50 btn btn-outline-success" onClick={alumnos}>
              Ver Estudiantes
            </button>
          </center>
          <br />
          <center>
            <button className="w-50 btn btn-outline-info" onClick={libros}>
              Ver Libros
            </button>
          </center>
          <br />
          <center>
            <button className="w-50 btn btn-outline-secondary" onClick={reporte}>
              Reportes
            </button>
          </center>
          <br />
          <center>
            <button className="w-50 btn btn-outline-danger" onClick={salir}>
              Salir
            </button>
          </center>
          <br/>
          <br/>
        </form>
      </div>
    </div>
  );
}

export default Principal_Admin;