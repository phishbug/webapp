import { useEffect, useState } from 'react';
import axiosInstance from '../axiosInstance';
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Button, Pagination, Box, Badge, Modal, ListItem } from '@mui/material';
import { useNavigate, useParams, Link } from 'react-router-dom';

// Style for the modal
const style = {
  position: 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 400,
  bgcolor: 'background.paper',
  border: '2px solid #000',
  boxShadow: 24,
  p: 4,
};

const Indices = () => {
  const [indicess, setIndicess] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [indicessPerPage] = useState(50);
  const [totalIndicess, setTotalIndicess] = useState(0);
  const [deleteIndicesId, setDeleteIndicesId] = useState(null);
  const [modalOpen, setModalOpen] = useState(false);
  const navigate = useNavigate();

  const fetchIndicess = async () => {
    const response = await axiosInstance.get(`/indicess`);
    setIndicess(response.data);
    setTotalIndicess(response.data.length);
  };

  useEffect(() => {
    fetchIndicess();
  }, []);

  const handleDelete = async (id) => {
    await axios.delete(`${import.meta.env.VITE_API_URL}/indicess/${id}`);
    fetchIndicess();
  };

  const handleOpenDeleteModal = (id) => {
    setDeleteIndicesId(id);
    setModalOpen(true);
  };

  const handleCloseDeleteModal = () => {
    setModalOpen(false);
    setDeleteIndicesId(null);
  };

  // Pagination Logic
  const indexOfLastIndices = currentPage * indicessPerPage;
  const indexOfFirstIndices = indexOfLastIndices - indicessPerPage;
  const currentIndicess = indicess.slice(indexOfFirstIndices, indexOfLastIndices);

  return (
    <Box>
      <Button variant="contained" onClick={() => navigate('/add-indices')}>Add Indices</Button>
      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Name</TableCell>
              <TableCell>No Of Docs</TableCell>
              <TableCell>Date</TableCell>
              <TableCell>Actions</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {currentIndicess.map((indices) => (
              	(indices.health != "green" && 
              		<TableRow key={indices.uuid}>
		                <TableCell> <span className={indices.health}>{indices.index}</span></TableCell>
		                <TableCell>{indices["docs.count"]}</TableCell>
		                <TableCell>{indices.date}</TableCell>
		                <TableCell>
		                  <ListItem component={Link} href="#" to={`/posts/${indices.index}`} >View Post</ListItem>
		                   <ListItem component={Link} href="#" to={`/add-indices/${indices.index}`} >Edit</ListItem>
		                  {indices["docs.count"] <=0 ? <Button onClick={() => handleOpenDeleteModal(indices.id)}>Delete</Button> : <span className="red">Can Not Delete, Please Delete All Docs.</span>}
		                </TableCell>
	              	</TableRow>
	            )
            ))}
          </TableBody>
        </Table>
      </TableContainer>
      <Pagination
        count={Math.ceil(totalIndicess / indicessPerPage)}
        page={currentPage}
        onChange={(event, value) => setCurrentPage(value)}
      />
      
      {/* Delete Confirmation Modal */}
      {modalOpen && (
      	<Modal
	        open={open}
	        // onClose={handleClose}
	        aria-labelledby="child-modal-title"
	        aria-describedby="child-modal-description"
	      >
        <Box sx={{ ...style, width: 200 }}>
        <div>
        
          <h3>Are you sure you want to delete this indices?</h3>
          <Button onClick={() => handleDelete(deleteIndicesId)}>Yes</Button>
          <Button onClick={handleCloseDeleteModal}>No</Button>
        </div>
        </Box>	
        
        </Modal>
      )}

    </Box>
  );
};

export default Indices;
