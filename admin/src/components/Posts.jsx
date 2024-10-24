import { useEffect, useState } from 'react';
import axiosInstance from '../axiosInstance';
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Button, Pagination, Box } from '@mui/material';
import { useNavigate, useParams } from 'react-router-dom';

const Posts = () => {
  const { id: id } = useParams(); // Get postId from URL params
  const [posts, setPosts] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [postsPerPage] = useState(5);
  const [totalPosts, setTotalPosts] = useState(0);
  const [deletePostId, setDeletePostId] = useState(null);
  const [modalOpen, setModalOpen] = useState(false);
  const navigate = useNavigate();

 useEffect(() => {
    const fetchPosts = async () => {
      if (id) {
        const response = await axiosInstance.get(`${import.meta.env.VITE_API_URL}/docs/` + id);
        const index = response.data;
        setPosts(response.data)
      }
    };

    fetchPosts();
  }, [id]);

  const handleDelete = async (id) => {
    await axios.delete(`${import.meta.env.VITE_API_URL}/posts/${id}`);
    fetchPosts();
  };

  const handleOpenDeleteModal = (id) => {
    setDeletePostId(id);
    setModalOpen(true);
  };

  const handleCloseDeleteModal = () => {
    setModalOpen(false);
    setDeletePostId(null);
  };

  // Pagination Logic
  const indexOfLastPost = currentPage * postsPerPage;
  const indexOfFirstPost = indexOfLastPost - postsPerPage;
  const currentPosts = posts.slice(indexOfFirstPost, indexOfLastPost);

  return (
    <Box>
      <Button variant="contained" onClick={() => navigate('/add-post')}>Add Post</Button>
      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Title</TableCell>
              <TableCell>Category</TableCell>
              <TableCell>Date</TableCell>
              <TableCell>Actions</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {currentPosts.map((post) => (
              <TableRow key={post.id}>
                <TableCell>{post.title}</TableCell>
                <TableCell>{post.category}</TableCell>
                <TableCell>{post.date}</TableCell>
                <TableCell>
                  <Button onClick={() => navigate(`/add-post/${post.id}`)}>Edit</Button>
                  <Button onClick={() => handleOpenDeleteModal(post.id)}>Delete</Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
      <Pagination
        count={Math.ceil(totalPosts / postsPerPage)}
        page={currentPage}
        onChange={(event, value) => setCurrentPage(value)}
      />
      
      {/* Delete Confirmation Modal */}
      {modalOpen && (
        <div>
          <h3>Are you sure you want to delete this post?</h3>
          <Button onClick={() => handleDelete(deletePostId)}>Yes</Button>
          <Button onClick={handleCloseDeleteModal}>No</Button>
        </div>
      )}
    </Box>
  );
};

export default Posts;
