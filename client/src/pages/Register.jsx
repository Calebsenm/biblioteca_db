import React, { useState } from 'react';
import axios from 'axios';
import "../styles/Register.css";

function Register() {
  const [formData, setFormData] = useState({
    nombre: '',
    direccion: '',
    telefono: '',
    correo: '',
    fecha_nacimiento: '',
    tipo_socio: 'normal',
    contrasena: '',
    rol: 'usuario',
  });

  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');


  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };


  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setSuccess('');

    if (
      !formData.nombre ||
      !formData.direccion ||
      !formData.telefono ||
      !formData.correo ||
      !formData.fecha_nacimiento ||
      !formData.contrasena
    ) {
      setError('Todos los campos son obligatorios');
      return;
    }

    try {

      const response = await axios.post('http://localhost:4000/api/register', formData, {
        headers: {
          'Content-Type': 'application/json',
        },
      });
      setSuccess(`Usuario registrado exitosamente. ID: ${response.data.id}`);
    } catch (error) {
      if (error.response) {
        setError(error.response.data.message || 'Error al registrar el usuario');
      } else {
        setError('Error en la conexión con el servidor');
      }
    }
  };

  return (
    <div className='register-container'>
      <div className='register-box'>
        <h2 className='register-header'>Registro de Usuario</h2>
        <form onSubmit={handleSubmit}>
          <div className='input-group'>
            <input
              type="text"
              name="nombre"
              placeholder='Nombre'
              value={formData.nombre}
              onChange={handleChange}
            />
          </div>
          <div className='input-group'>

            <input
              type="text"
              name="direccion"
              placeholder='Direccion'
              value={formData.direccion}
              onChange={handleChange}
            />
          </div>
          <div className='input-group'>

            <input
              type="text"
              name="telefono"
              placeholder='Telefono'
              value={formData.telefono}
              onChange={handleChange}
            />
          </div>
          <div className='input-group'>

            <input
              type="email"
              name="correo"
              placeholder='Correo'
              value={formData.correo}
              onChange={handleChange}
            />
          </div>
          <div className='input-group'>

            <input
              type="password"
              name="contrasena"
              placeholder='Contrasena'
              value={formData.contrasena}
              onChange={handleChange}
            />
          </div>
          <div className='input-group'>
            <label>Fecha Nacimiento:</label>
            <input
              type="date"
              name="fecha_nacimiento"
              placeholder='Fecha Nacimiento'
              value={formData.fecha_nacimiento}
              onChange={handleChange}
            />
          </div>
          <div className='input-group'>
            <label>Tipo de Socio:</label>
            <select
              name="tipo_socio"
              className='inputs'
              value={formData.tipo_socio}
              onChange={handleChange}
            >
              <option value="normal">Normal</option>
              <option value="estudiante">Estudiante</option>
              <option value="profesor">Profesor</option>
            </select>
          </div>

          <div className='input-group'>
            <label>Rol:</label>
            <select
              className='inputs'
              name="rol"
              value={formData.rol}
              onChange={handleChange}
            >
              <option value="usuario">Usuario</option>
              <option value="administrador">Administrador</option>
            </select>
          </div>
          <button type="submit" className='login-btn'>Registrar </button>
        </form>

        {error && <div style={{ color: 'red' }}>{error}</div>}
        {success && <div style={{ color: 'green' }}>{success}</div>}
      </div>
    </div>
  );
}

export default Register;
