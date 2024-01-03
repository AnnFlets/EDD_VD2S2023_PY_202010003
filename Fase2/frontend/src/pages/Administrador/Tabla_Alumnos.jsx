import React, { useState, useEffect } from "react";

function Tabla_Alumnos() {
  const [alumnosRegistrados, SetAlumnosRegistrados] = useState([]);
  useEffect(() => {
    async function peticion() {
      const response = await fetch("http://localhost:4000/tabla-alumnos");
      const result = await response.json();
      SetAlumnosRegistrados(result.Arreglo);
    }
    peticion();
  }, []);

  const salir = (e) => {
    e.preventDefault();
    console.log("Listo");
    window.open("/principal/admin", "_self");
  };

  return (
    <div className="form-signin1">
      <div className="text-center">
        <form className="card card-body">
          <br/>
          <h1 className="h3 mb-3 fw-normal"><b>MENÚ ADMINISTRADOR</b></h1>
          <h4 className="h3 mb-3 fw-normal">Estudiantes registrados</h4>
          <br />
          <center>
            <button className="w-50 btn btn-outline-danger" onClick={salir}>
              Regresar
            </button>
          </center>
          <br />
          <table className="table table-ligth table-striped table-bordered">
            <thead className="table-dark">
              <tr>
                <th scope="col">#</th>
                <th scope="col">Posicion</th>
                <th scope="col">Carnet </th>
                <th scope="col">Password </th>
              </tr>
            </thead>
            <tbody>
              {alumnosRegistrados.map((element, j) => {
                if (element.Estudiante != null) {
                  return (
                    <>
                      <tr key={"alum" + j}>
                        <th scope="row">{j + 1}</th>
                        <td>{element.Llave}</td>
                        <td>{element.Estudiante.Carnet}</td>
                        <td>{element.Estudiante.Password}</td>
                      </tr>
                    </>
                  );
                }
              })}
            </tbody>
          </table>
          <br />
        </form>
      </div>
    </div>
  );
}

export default Tabla_Alumnos;
