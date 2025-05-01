import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './CreateBook.css';

function CreateBook() {
  const [formData, setFormData] = useState({
    titulo: '',
    genero: '',
    fechapublicacion: '',
    ideditorial: '',
    autores: [],
  });

  const [editorials, setEditorials] = useState([]);
  const [authors, setAuthors] = useState([]);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  useEffect(() => {
    axios.get('http://localhost:4000/api/editoriales')
      .then(response => {
        const editorialesData = response.data.map(editorial => ({
          ...editorial,
          ideditorial: Number(editorial.ideditorial)
        }));
        setEditorials(editorialesData);
      })
      .catch(() => setError('Error al cargar editoriales'));

    axios.get('http://localhost:4000/api/autores')
      .then(response => {
        const autoresData = response.data.map(author => ({
          ...author,
          id: Number(author.id)
        }));
        setAuthors(autoresData);
      })
      .catch(() => setError('Error al cargar autores'));
  }, []);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: name === 'ideditorial' ? Number(value) : value,
    }));
  };

  const handleSelectAutores = (e) => {
    const value = e.target.value;
    setFormData(prev => ({
      ...prev,
      autores: value ? [Number(value)] : []
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setSuccess('');

    // Validación mejorada
    if (!formData.titulo || !formData.genero || !formData.fechapublicacion || 
        !formData.ideditorial || formData.autores.length === 0) {
      setError('Todos los campos son obligatorios');
      return;
    }

    try {
      const requestData = {
        titulo: formData.titulo,
        genero: formData.genero,
        fechapublicacion: formData.fechapublicacion,
        ideditorial: Number(formData.ideditorial),
        autores: formData.autores.map(Number)
      };

      console.log('Datos a enviar:', requestData); // Para depuración

      await axios.post('http://localhost:4000/api/admin/books', requestData, {
        headers: { 'Content-Type': 'application/json' }
      });

      setSuccess('Libro creado exitosamente');
      // Resetear formulario
      setFormData({
        titulo: '',
        genero: '',
        fechapublicacion: '',
        ideditorial: '',
        autores: [],
      });
    } catch (error) {
      console.error('Error al crear libro:', error);
      setError(error.response?.data?.message || 'Error al crear el libro');
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
            value={formData.autores[0]}
            onChange={handleSelectAutores}
            required
          >
            <option value="">Seleccione un autor</option>
            {authors.map(author => (
              <option key={author.id} value={author.id}>
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