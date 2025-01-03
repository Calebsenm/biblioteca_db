import React, { useState } from 'react';
import axios from 'axios';
import "./CreateEditorial.css";


function CreateEditorial() {
  const [editorialData, setEditorialData] = useState({
    nombre: '',
    direccion: '',
    paginaWeb: '',
  });

  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  const handleChange = (e) => {
    const { name, value } = e.target;
    setEditorialData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setSuccess('');

    // Validación de campos
    if (!editorialData.nombre || !editorialData.direccion || !editorialData.paginaWeb) {
      setError('Todos los campos son obligatorios');
      return;
    }

    try {
      // Enviar datos para crear editorial
      const response = await axios.post('http://localhost:4000/api/editoriales', editorialData, {
        headers: {
          'Content-Type': 'application/json',
        },
      });

      // Mostrar mensaje de éxito
      setSuccess('Editorial creada exitosamente');
      setEditorialData({ nombre: '', direccion: '', paginaWeb: '' });
    } catch (error) {
      // Manejo de errores
      if (error.response) {
        setError(error.response.data.message || 'Error al crear la editorial');
      } else {
        setError('Error en la conexión con el servidor');
      }
    }
  };

  return (
    <div className="form-container" >
      <h2>Crear Nueva Editorial</h2>
      <form onSubmit={handleSubmit}>
        <div className="form-group" >
          <label>Nombre:</label>
          <input
            type="text"
            name="nombre"
            value={editorialData.nombre}
            onChange={handleChange}
          />
        </div>
        <div className="form-group"  >
          <label>Dirección:</label>
          <input
            type="text"
            name="direccion"
            value={editorialData.direccion}
            onChange={handleChange}
          />
        </div>
        <div className="form-group" >
          <label>Página Web:</label>
          <input
            type="text"
            name="paginaWeb"
            value={editorialData.paginaWeb}
            onChange={handleChange}
          />
        </div>
        
        <button type="submit" className="my-button"> 
          Crear Editorial
        </button>
      </form>

      {error && <div style={{ color: 'red' }}>{error}</div>}
      {success && <div style={{ color: 'green' }}>{success}</div>}
    </div>
  );
}

export default CreateEditorial;