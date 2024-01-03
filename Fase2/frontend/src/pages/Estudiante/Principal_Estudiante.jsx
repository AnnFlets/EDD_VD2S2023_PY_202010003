import React, { useEffect, useState } from "react";

function Principal_Estudiante() {
  const [cursos, setCursos] = useState([]);

  useEffect(() => {
    async function PedirCursos() {
      const valorLocal = localStorage.getItem("user");
      const response = await fetch("http://localhost:4000/obtener-clases", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          Carnet: valorLocal,
        }),
      });
      const result = await response.json();
      console.log(result);
      setCursos(result.Arreglo);
    }
    PedirCursos();
  }, []);

  const Palabra = () => {
    return (
      <div className="row">
        <div className="row align-items-start">
          {cursos.map((item, i) => (
            <div className="form-signin1 col" key={"CursoEstudiante" + i}>
              <div className="text-center">
                <div className="card card-body">
                  <h1 className="text-left" key={"album" + i} value={i}>
                    {item}
                  </h1>
                  <div>
                    <span
                      className="input-group-text"
                      id="validationTooltipUsernamePrepend"
                    ></span>{" "}
                    <br />
                  </div>
                </div>
                <br />
              </div>
            </div>
          ))}
        </div>
      </div>
    );
  };

  const salir = (e) => {
    e.preventDefault();
    console.log("Listo");
    localStorage.clear();
    window.open("/", "_self");
  };

  const publicaciones = (e) => {
    e.preventDefault();
    localStorage.setItem("cursos", JSON.stringify(cursos));
    window.open("/principal/estudiante/publicacion", "_self");
  };

  const libro = (e) => {
    e.preventDefault();
    localStorage.setItem("cursos", JSON.stringify(cursos));
    window.open("/principal/estudiante/libro", "_self");
  };

  return (
    <div className="form-signin1">
      <br/>
      <div className="text-center">
        <form className="card card-body">
          <br/>
          <h1 className="h3 mb-3 fw-normal"><b>MENÚ ESTUDIANTE</b></h1>
          <br />
          <center>
            <button className="w-50 btn btn-outline-primary" onClick={libro}>
              Ver Libros
            </button>
          </center>
          <br />
          <center>
            <button
              className="w-50 btn btn-outline-info"
              onClick={publicaciones}
            >
              Ver Publicaciones
            </button>
          </center>
          <br />
          <center>
            <button className="w-50 btn btn-outline-danger" onClick={salir}>
              Salir
            </button>
          </center>
          <br />
          <br/>
          <h4 className="h3 mb-3 fw-normal">Cursos asignados</h4>
          {cursos.length > 0 ? <Palabra /> : null}
        </form>
      </div>
    </div>
  );
}

export default Principal_Estudiante;
