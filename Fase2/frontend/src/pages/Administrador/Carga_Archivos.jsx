import React from "react";
import "../../styles/otros.css";

function Carga_Archivos() {
  const reportes = (e) => {
    e.preventDefault();
    window.open("/principal/admin/alumnos", "_self");
  };

  const uploadFileTutor = (event) => {
    const file = event.target.files[0];
    const reader = new FileReader();

    reader.onload = (event) => {
      const content = event.target.result;
      const parsedData = parseCSV(content);
      parsedData.map(async (row) => {
        if (row.length > 1) {
          const response = await fetch(
            "http://localhost:4000/registrar-tutor",
            {
              method: "POST",
              headers: {
                "Content-Type": "application/json",
              },
              body: JSON.stringify({
                Carnet: parseInt(row[0]),
                Nombre: row[1],
                Curso: row[2],
                Password: row[3],
              }),
            }
          );
          const result = await response.json();
        }
      });
    };

    reader.onerror = (error) => {
      console.error("Error al leer el archivo:", error);
    };

    reader.readAsText(file);
  };

  const uploadFileAlumno = (event) => {
    const file = event.target.files[0];
    const reader = new FileReader();

    reader.onload = (event) => {
      const content = event.target.result;
      const parsedData = parseCSV(content);
      console.log(parsedData);
      parsedData.map(async (row) => {
        if (row.length > 1) {
          const response = await fetch(
            "http://localhost:4000/registrar-alumno",
            {
              method: "POST",
              headers: {
                "Content-Type": "application/json",
              },
              body: JSON.stringify({
                Carnet: parseInt(row[0]),
                Nombre: row[1],
                Password: row[2],
                Cursos: [row[3], row[4], row[5]],
              }),
            }
          );
          const result = await response.json();
        }
      });
    };

    reader.onerror = (error) => {
      console.error("Error al leer el archivo:", error);
    };

    reader.readAsText(file);
  };

  const parseCSV = (csvContent) => {
    const rows = csvContent.split("\n");
    const encabezado = rows.slice(1);
    const sinRetorno = encabezado.map((row) => row.trim());
    const data = sinRetorno.map((row) => row.split(","));
    return data;
  };

  const onChange1 = (e) => {
    var reader = new FileReader();
    reader.onload = async (e) => {
      var obj = JSON.parse(e.target.result);
      console.log(obj);
      const response = await fetch("http://localhost:4000/registrar-cursos", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          Cursos: obj.Cursos,
        }),
      });
    };
    reader.readAsText(e.target.files[0]);
  };

  const salir = (e) => {
    e.preventDefault();
    console.log("Listo");
    window.open("/principal/admin", "_self");
  };

  return (
    <div className="form-signin1">
      <br/>
      <br/>
      <div className="text-center">
        <form className="card card-body">
          <br/>
          <h1 className="h3 mb-3 fw-normal"><b>MENÚ ADMINISTRADOR</b></h1>
          <h4 className="h3 mb-3 fw-normal">Cargar Archivos</h4>
          <br />
          <div className="input-group mb-3">
            <label className="input-group-text">Cargar Tutores</label>
            <input
              className="form-control"
              id="inputGroupFile01"
              type="file"
              accept=".csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, application/vnd.ms-excel"
              onChange={uploadFileTutor}
            />
          </div>
          <br />
          <div className="input-group mb-3">
            <label className="input-group-text">Cargar Estudiantes</label>
            <input
              className="form-control"
              id="inputGroupFile01"
              type="file"
              accept=".csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, application/vnd.ms-excel"
              onChange={uploadFileAlumno}
            />
          </div>
          <br />
          <div className="input-group mb-3">
            <label className="input-group-text">Cargar Cursos</label>
            <input
              className="form-control"
              id="inputGroupFile02"
              type="file"
              accept="application/json"
              onChange={onChange1}
            />
          </div>
          <br />
          <center>
            <button className="w-50 btn btn-outline-danger" onClick={salir}>
              Regresar
            </button>
          </center>
          <br />
        </form>
      </div>
    </div>
  );
}

export default Carga_Archivos;
