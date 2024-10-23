import { useState } from 'react';
import axios from 'axios';
import { TextField, Button, Alert, Box } from '@mui/material';

const IndexUpdate = () => {
  const [jsonData, setJsonData] = useState('');
  const [name, setName] = useState('');
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setSuccess('');

    try {
      const response = await axios.post(`${import.meta.env.VITE_API_URL}/index-update`, {
        data: JSON.parse(jsonData), // Ensure valid JSON
      });
      if (response.status === 200) {
        setSuccess('Index updated successfully!');
      }
    } catch (err) {
      setError('Failed to update index. Please check your input.');
    }
  };

  return (
    <Box>
      <h2>Update Index</h2>
      {error && <Alert severity="error">{error}</Alert>}
      {success && <Alert severity="success">{success}</Alert>}
      <form onSubmit={handleSubmit}>

      	<TextField
          label="Name"
          variant="outlined"
          fullWidth
          margin="normal"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />

        <TextField
          label="JSON Data"
          multiline
          rows={10}
          fullWidth
          variant="outlined"
          value={jsonData}
          onChange={(e) => setJsonData(e.target.value)}
          required
          sx={{ my: 2 }}
        />
        <Button type="submit" variant="contained" color="primary" style={{ marginTop: '16px' }}>
          Update Index
        </Button>
      </form>
    </Box>
  );
};

export default IndexUpdate;
