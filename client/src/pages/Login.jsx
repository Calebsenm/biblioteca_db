import React, { useState } from "react";
import "../styles/login.css";
import { useNavigate } from "react-router-dom";

const Login = () => {

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {

      const response = await fetch("http://localhost:4000/api/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ correo: email, contrasena: password }),
      });

      if (!response.ok) {

        const errorData = await response.json();
        setErrorMessage(errorData.message || "Login failed.");
        return;
      }

      const data = await response.json();
      localStorage.setItem("user", JSON.stringify(data));


      if (data.rol === "administrador") {
        navigate("/admin/dashboard");
      } else {
        navigate("/user/dashboard");
      }
    } catch (error) {
      console.error("Error during login:", error);
      setErrorMessage("An error occurred. Please try again.");
    }
  };

  return (
    <div className="login-container">
      <div className="login-box">
        <h2 className="login-header">Biblioteca</h2> 
        <form onSubmit={handleSubmit}>
          <div className="input-group">
            <input
              type="email"
              placeholder="Email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          </div>
          <div className="input-group">
            <input
              type="password"
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>

          <button type="submit" className="login-btn">
            Log in
          </button>
          <div className="text">
            <div className="create-acount">
              <a href="register">Create Acount</a>
            </div>
            <div className="forgot-password">
              <a href="#">Forgot password?</a>
            </div>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Login;




/*

para recuperar el localestorage
const user = JSON.parse(localStorage.getItem("user"));

if (user) {
  console.log("Usuario logueado:", user);
  console.log("Rol del usuario:", user.rol);
}

*/