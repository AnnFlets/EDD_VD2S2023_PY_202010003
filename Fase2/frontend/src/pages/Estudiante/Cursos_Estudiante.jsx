import React, { useState, useEffect } from "react";

function Cursos_Estudiante() {
  const [cursos, setCursos] = useState([]);
  const [libros, setLibros] = useState([]);

  useEffect(() => {
    async function PedirCursos() {
      const arregloCursos = localStorage.getItem("cursos");
      const cursosArreglo = JSON.parse(arregloCursos);
      const response = await fetch(
        "http://localhost:4000/obtener-libros-alumno"
      );
      const result = await response.json();
      console.log(result);
      setCursos(cursosArreglo);
      if (result.status === 200) {
        setLibros(result.Arreglo);
      }
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
                    {libros.map((ite, j) => {
                      if (ite.Curso === item) {
                        return (
                          <div key={"p-" + i}>
                            <label className="input-group-text" key={"lib" + j}>
                              {ite.Nombre}
                            </label>
                          </div>
                        );
                      }
                    })}
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
    window.open("/principal/estudiante/publicacion", "_self");
  };

  const principal = (e) => {
    e.preventDefault();
    window.open("/principal/estudiante", "_self");
  };

  return (
    <div className="form-signin1">
      <div className="text-center">
        <form className="card card-body">
          <br/>
          <h1 className="h3 mb-3 fw-normal"><b>MENÚ ESTUDIANTE</b></h1>
          <br />
          <center>
            <button
              className="w-50 btn btn-outline-warning"
              onClick={principal}
            >
              Regresar a Principal
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
          <br/>
          <center>
            <button className="w-50 btn btn-outline-danger" onClick={salir}>
              Salir
            </button>
          </center>
          <br />
          <br/>
          <h4 className="h3 mb-3 fw-normal">Libros</h4>
          {cursos.length > 0 ? <Palabra /> : null}
        </form>
      </div>
    </div>
  );
}

export default Cursos_Estudiante;
