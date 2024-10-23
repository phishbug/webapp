import { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { TextField, Button, Box, Typography, Alert } from '@mui/material';

const Login = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null); // Reset the error state before the request
    try {
      const response = await axios.post(`${import.meta.env.VITE_API_URL}/login`, { username, password });
      localStorage.setItem('token', response.data.data.token); // Store the token
      navigate('/'); // Redirect to dashboard
    } catch (error) {
      setError("Login failed. Please check your credentials."); // Set error message
    }
  };

  return (
    <Box
      sx={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        height: '100vh',
        p: 2,
      }}
    >
      <Typography variant="h4" sx={{ mb: 2 }}>Login</Typography>
      {error && <Alert severity="error" sx={{ mb: 2 }}>{error}</Alert>} {/* Display error alert */}
      <form onSubmit={handleSubmit} style={{ width: '100%', maxWidth: '400px' }}>
        <TextField
          label="Username"
          variant="outlined"
          fullWidth
          margin="normal"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          required
        />
        <TextField
          label="Password"
          type="password"
          variant="outlined"
          fullWidth
          margin="normal"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <Button 
          type="submit" 
          variant="contained" 
          color="primary" 
          fullWidth 
          sx={{ mt: 2 }}
        >
          Login
        </Button>
      </form>
    </Box>
  );
};

export default Login;
