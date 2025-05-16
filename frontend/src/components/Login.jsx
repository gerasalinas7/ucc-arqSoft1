import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';

function Login() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [errorMessage, setErrorMessage] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      // Enviar solicitud POST al backend para hacer login
      const response = await axios.post('http://localhost:8080/login', {
        username,
        password,
      });

      if (response.data.token) {
        console.log('Login OK');
        navigate('/hello'); // Redirigir a la pantalla Hello World
      } else {
        console.log('Login Incorrecto');
        setErrorMessage('Login Incorrecto');
      }
    } catch (error) {
      console.error('Error de login:', error);
      setErrorMessage('Error al conectar con el servidor');
    }
  };

  return (
    <div>
      <h2>Login</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Username:</label>
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
        </div>
        <div>
          <label>Password:</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <button type="submit">Ingresar</button>
      </form>
      {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
    </div>
  );
}

export default Login;
