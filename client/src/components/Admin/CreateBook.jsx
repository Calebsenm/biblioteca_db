import React, { useState, useEffect } from 'react';
import axios from 'axios';
import '../../styles/CreateBook.css';

function CreateBook() {
  const [formData, setFormData] = useState({
    titulo: '',
    genero: '',
    fechapublicacion: '',
    ideditorial: '',
    autores: '',           
  });

  const [editorials, setEditorials] = useState([]);
  const [authors, setAuthors] = useState([]);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  useEffect(() => {
    axios.get('http://localhost:4000/api/editoriales')
      .then(res => setEditorials(res.data))
      .catch(() => setError('Error al cargar editoriales'));

    axios.get('http://localhost:4000/api/autores')
      .then(res => setAuthors(res.data))
      .catch(() => setError('Error al cargar autores'));
  }, []);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: name === 'ideditorial' || name === 'autores'
        ? value  
        : value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setSuccess('');

    const { titulo, genero, fechapublicacion, ideditorial, autores } = formData;


    if (!titulo || !genero || !fechapublicacion || !ideditorial || !autores) {
      setError('Todos los campos son obligatorios');
      return;
    }

    try {
      const requestData = {
        titulo,
        genero,
        fechapublicacion,
        ideditorial: Number(ideditorial),
        autores: [Number(autores)], 
      };

      await axios.post('http://localhost:4000/api/admin/books', requestData, {
        headers: { 'Content-Type': 'application/json' }
      });

      setSuccess('Libro creado exitosamente');
      setFormData({
        titulo: '',
        genero: '',
        fechapublicacion: '',
        ideditorial: '',
        autores: '',
      });
    } catch (err) {
      setError(err.response?.data?.message || 'Error al crear el libro');
    }
  };

  return (
    <div className="form-container">
      <h2>Crear Nuevo Libro</h2>
      <form onSubmit={handleSubmit}>
      <div className="form-group">
          <label>Título:</label>
          <input
            type="text"
            name="titulo"
            value={formData.titulo}
            onChange={handleChange}
            required
          />
        </div>

        <div className="form-group">
          <label>Género:</label>
          <input
            type="text"
            name="genero"
            value={formData.genero}
            onChange={handleChange}
            required
          />
        </div>

        <div className="form-group">
          <label>Fecha de Publicación:</label>
          <input
            type="date"
            name="fechapublicacion"
            value={formData.fechapublicacion}
            onChange={handleChange}
            required
          />
        </div>

        <div className="form-group">
          <label>Editorial:</label>
          <select
            name="ideditorial"
            value={formData.ideditorial}
            onChange={handleChange}
            required
          >
            <option value="">Seleccione una editorial</option>
            {editorials.map(editorial => (
              <option key={editorial.ideditorial} value={editorial.ideditorial}>
                {editorial.nombre}
              </option>
            ))}
          </select>
        </div>

        <div className="form-group">
          <label>Autor:</label>
          <select
            name="autores"                      
            value={formData.autores}          
            onChange={handleChange}           
            required
          >
            <option value="">Seleccione un autor</option>
            {authors.map(author => (
              <option key={author.idautor} value={author.idautor}>
                {author.nombre}
              </option>
            ))}
          </select>
        </div>

        <button type="submit" className="submit-btn">Crear Libro</button>
      </form>

      {error && <div className="error-message">{error}</div>}
      {success && <div className="success-message">{success}</div>}
    </div>
  );
}

export default CreateBook;
