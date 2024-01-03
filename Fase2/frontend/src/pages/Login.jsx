import React, { useState } from "react";
import "../styles/login.css";
import "bootstrap/dist/css/bootstrap.min.css";

function Login() {
    const [isChecked, setIsChecked] = useState(false);
    const [userName, setUserName] = useState("");
    const [passwordUser, setPasswordUser] = useState("");

    const handleSubmit = async (e) => {
        e.preventDefault();
        const response = await fetch("http://localhost:4000/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                UserName: userName,
                Password: passwordUser,
                Tutor: isChecked,
            }),
        });

        const result = await response.json();
        if (result.rol == 0) {
            alert("Credenciales Incorrectas");
        } else if (result.rol == 1) {
            window.open("/principal/admin", "_self");
            localStorage.setItem("Tipo", "1");
            localStorage.setItem("user", userName);
        } else if (result.rol == 2) {
            window.open("/principal/tutor/libro", "_self");
            localStorage.setItem("Tipo", "2");
            localStorage.setItem("user", userName);
        } else if (result.rol == 3) {
            window.open("/principal/estudiante", "_self");
            localStorage.setItem("Tipo", "3");
            localStorage.setItem("user", userName);
        }
    };
    return (
        <div className="container">
            <br/>
            <br/>
            <br/>
            <br/>
            <div className="row justify-content-center">
                <div className="col-12 col-md-9 col-lg-7 col-xl-6 col-xxl-5">
                    <div className="card border border-light-subtle rounded-4">
                        <div className="card-body p-3 p-md-4 p-xl-5">
                            <div className="row">
                                <div className="col-12">
                                    <div className="mb-5">
                                        <h1 className="text-center">TUTORÍAS ECYS</h1>
                                        <h3 className="text-center">Inicio de sesión</h3>
                                    </div>
                                </div>
                            </div>
                            <form onSubmit={handleSubmit}>
                                <div className="row gy-3 overflow-hidden">
                                    <div className="col-12">
                                        <label>Usuario:</label>
                                        <div className="form-floating mb-3">
                                            <input type="text" className="form-control" placeholder="Nombre de usuario" required value={userName} onChange={(e) => setUserName(e.target.value)} autoFocus/>
                                        </div>
                                    </div>
                                    <div className="col-12">
                                        <label>Contraseña:</label>
                                        <div className="form-floating mb-3">
                                            <input type="password" className="form-control" placeholder="Contraseña" aria-describedby="passwordHelpInline" value={passwordUser} onChange={(e) => setPasswordUser(e.target.value)} autoFocus/>
                                        </div>
                                    </div>
                                    <div className="col-12">
                                        <div className="form-check form-switch text-left">
                                            <input className="form-check-input" type="checkbox" role="switch" id="flexSwitchCheckDefault" value={isChecked} onChange={(e) => setIsChecked(!isChecked)}/>
                                            <label className="form-check-label" htmlFor="flexSwitchCheckDefault">Tutor</label>
                                        </div>
                                    </div>
                                    <div className="col-12">
                                        <div className="d-grid">
                                            <button className="btn btn-outline-primary" type="submit">Iniciar sesión</button>
                                        </div>
                                    </div>
                                </div>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Login