import React, { useState } from "react";

function Reporte() {
  const [imagen, setImagen] = useState(
    "https://i.pinimg.com/originals/a4/b4/a3/a4b4a3bc07251390a64e8b8b0630737e.jpg"
  );
  const salir = (e) => {
    e.preventDefault();
    console.log("Listo");
    window.open("/principal/admin", "_self");
  };

  const validar = (data) => {
    console.log(data);
    setImagen(data.imagen.Imagenbase64);
  };

  const reporteGrafo = async (e) => {
    e.preventDefault();
    fetch("http://localhost:4000/reporte-grafo", {})
      .then((response) => response.json())
      .then((data) => validar(data));
  };

  const reporteArbol = async (e) => {
    e.preventDefault();
    fetch("http://localhost:4000/reporte-arbol", {})
      .then((response) => response.json())
      .then((data) => validar(data));
  };

  const reporteBlockchain = async (e) => {
    e.preventDefault();
    fetch("http://localhost:4000/reporte-bloque", {})
      .then((response) => response.json())
      .then((data) => validar(data));
  };

  return (
    <div className="form-signin1">
      <div className="text-center">
        <form className="card card-body">
          <h1 className="h3 mb-3 fw-normal"><b>MENÚ ADMINISTRADOR</b></h1>
          <h4 className="h3 mb-3 fw-normal">Reportes</h4>
          <br />
          <center>
            <button
              className="w-50 btn btn-outline-success"
              onClick={reporteArbol}
            >
              Arbol B
            </button>
          </center>
          <br />
          <center>
            <button
              className="w-50 btn btn-outline-info"
              onClick={reporteGrafo}
            >
              Grafo
            </button>
          </center>
          <br />
          <center>
            <button
              className="w-50 btn btn-outline-warning"
              onClick={reporteBlockchain}
            >
              Árbol Merkle
            </button>
          </center>
          <br />
          <center>
            <button className="w-50 btn btn-outline-danger" onClick={salir}>
              Salir
            </button>
          </center>
          <br />
          <center>
            <img src={imagen} width="350" height="350" alt="some value" />
          </center>
        </form>
      </div>
    </div>
  );
}

export default Reporte;