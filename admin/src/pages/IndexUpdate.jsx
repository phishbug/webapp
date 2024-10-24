import { useState, useEffect } from 'react';
import axiosInstance from '../axiosInstance';
import { TextField, Button, Alert, Box } from '@mui/material';
import { useNavigate, useParams } from 'react-router-dom';

const IndexUpdate = () => {
  const { id: id } = useParams(); // Get postId from URL params
  const [settingData, setSettingData] = useState('');
  const [mappingData, setMappingData] = useState('');
  const [name, setName] = useState(id);
  const [indices, setIndices] = useState('');
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

   useEffect(() => {
    const fetchIndex = async () => {
      if (id) {
        const response = await axiosInstance.get(`${import.meta.env.VITE_API_URL}/index/` + id);
        const index = response.data;
        setMappingData(response.data[0])
        setSettingData(response.data[1])
      }
    };

    fetchIndex();
  }, [id]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setSuccess('');

    try {
      const response = await axios.post(`${import.meta.env.VITE_API_URL}/index-update`, {
        data: JSON.parse(jsonData),
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
          label="Setting Data"
          multiline
          rows={10}
          fullWidth
          variant="outlined"
          value={settingData}
          onChange={(e) => setMappingData(e.target.value)}
          required
          sx={{ my: 2 }}
        />

        <TextField
          label="Mapping Data"
          multiline
          rows={10}
          fullWidth
          variant="outlined"
          value={mappingData}
          onChange={(e) => setSettingData(e.target.value)}
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
