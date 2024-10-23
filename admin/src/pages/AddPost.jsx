import { useEffect, useState } from 'react';
import axios from 'axios';
import ReactQuill from 'react-quill';
import 'react-quill/dist/quill.snow.css';
import { TextField, Button, FormControl, InputLabel, Select, MenuItem, Alert, Box } from '@mui/material';
import { useNavigate, useParams } from 'react-router-dom';

const AddPost = () => {
  const { id: postId } = useParams(); // Get postId from URL params
  const [title, setTitle] = useState('');
  const [category, setCategory] = useState('');
  const [date, setDate] = useState('');
  const [tags, setTags] = useState('');
  const [content, setContent] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    const fetchPost = async () => {
      if (postId) {
        const response = await axios.get(`${import.meta.env.VITE_API_URL}/posts/${postId}`);
        const post = response.data;
        setTitle(post.title);
        setCategory(post.category);
        setDate(post.date);
        setTags(post.tags.join(', '));
        setContent(post.content || '');
      }
    };

    fetchPost();
  }, [postId]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(''); // Reset error message

    const newPost = {
      title,
      category,
      date,
      tags: tags.split(',').map(tag => tag.trim()),
      content,
    };

    try {
      if (postId) {
        // Update post
        await axios.put(`${import.meta.env.VITE_API_URL}/posts/${postId}`, newPost);
      } else {
        // Create post
        await axios.post(`${import.meta.env.VITE_API_URL}/posts`, newPost);
      }
      navigate('/posts'); // Redirect to posts page
    } catch (err) {
      setError('Failed to save the post. Please try again.'); // Set error message
    }
  };

  return (
    <Box>
      <h2>{postId ? 'Edit Post' : 'Add Post'}</h2>
      {error && <Alert severity="error">{error}</Alert>} {/* Display error alert */}
      <form onSubmit={handleSubmit}>
        <TextField
          label="Title"
          fullWidth
          variant="outlined"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
          sx={{ my: 2 }} // Add top and bottom margin
        />
        <FormControl fullWidth margin="normal" sx={{ my: 2 }}>
          <InputLabel>Category</InputLabel>
          <Select
            value={category}
            onChange={(e) => setCategory(e.target.value)}
            required
          >
            <MenuItem value="Category 1">Category 1</MenuItem>
            <MenuItem value="Category 2">Category 2</MenuItem>
            <MenuItem value="Category 3">Category 3</MenuItem>
          </Select>
        </FormControl>
        <TextField
          label="Date"
          type="date"
          fullWidth
          variant="outlined"
          value={date}
          onChange={(e) => setDate(e.target.value)}
          required
          sx={{ my: 2 }} // Add top and bottom margin
        />
        <TextField
          label="Tags (comma separated)"
          fullWidth
          variant="outlined"
          value={tags}
          onChange={(e) => setTags(e.target.value)}
          sx={{ my: 2 }} // Add top and bottom margin
        />
        <ReactQuill
          value={content}
          onChange={setContent}
          placeholder="Write your post content here..."
          style={{ margin: '16px 0' }} // Add margin to Quill editor
        />
        <Button type="submit" variant="contained" color="primary" style={{ marginTop: '16px' }}>
          {postId ? 'Update' : 'Create'}
        </Button>
      </form>
    </Box>
  );
};

export default AddPost;
