import React, { useState } from "react";
import axios from "axios";
import "../../styles/CreateAuthor.css"

function createAuthor() {

    const [authorData, setAuthorData] = useState({
        name: '',
        nationality: '',
    })

    const [error, setError] = useState('');
    const [success, setSuccess] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');
        setSuccess('');

        if (!authorData.name || !authorData.nationality) {
            setError('Todos los campos son obligatorios');
            return;
        }

        try {
            const response = await axios.post('http://localhost:4000/api/admin/autores', authorData, {
                headers: {
                    'Content-Type': 'application/json',
                },
            });

            setSuccess('Autor guardado exitosamente');
            setAuthorData({ name: '', nationality: '' });


        } catch (error) {
            if (error.response) {
                setError(error.response.data.message || 'Error al crear autor');
            } else {
                setError('Error en la conexión con el servidor');
            }
        }
    }

    const handleChange = (e) => {
        const { name, value } = e.target;
        setAuthorData((preData) => ({
            ...preData,
            [name]: value,
        }));
    };

    return (
        <div className='form_container'>
            <h2>Guardar Nuevo Autor</h2>
            <form onSubmit={handleSubmit}>
                <div className="form-group" >
                    <label>Nombre:</label>
                    <input
                        type="text"
                        name="name"
                        value={authorData.name}
                        onChange={handleChange}
                    />
                </div>
                <div className="form-group"  >
                    <label>Dirección:</label>
                    <input
                        type="text"
                        name="nationality"
                        value={authorData.nationality}
                        onChange={handleChange}
                    />
                </div>

                <button type="submit" className="my-button">
                    Crear Autor
                </button>
            </form>

            {error && <div style={{ color: 'red' }}>{error}</div>}
            {success && <div style={{ color: 'green' }}>{success}</div>}
        </div>
    );
}
export default createAuthor; 