import React from 'react';
import { List, ListItem, ListItemIcon, ListItemText, Drawer, Box, Typography } from '@mui/material';
import { Dashboard, PostAdd, Logout } from '@mui/icons-material';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

const drawerWidth = 240;

const Sidebar = () => {
  const { logout } = useAuth();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate('/login'); // Redirect to login after logout
  };

  return (
    <Drawer
      variant="permanent"
      sx={{
        width: drawerWidth,
        flexShrink: 0,
        '& .MuiDrawer-paper': {
          width: drawerWidth,
          boxSizing: 'border-box',
        },
      }}
    >
      <Box sx={{ display: 'flex', alignItems: 'center', p: 2, bgcolor: '#000', color: 'white' }}>
        <img src="https://phishbug-images.s3.us-west-1.amazonaws.com/favicon-32x32.png" alt="Logo" style={{ width: 40, height: 40, marginRight: 8 }} />
        <Typography variant="h6">Admin Panel</Typography>
      </Box>
      <List>
        <ListItem button component={Link} to="/">
          <ListItemIcon>
            <Dashboard />
          </ListItemIcon>
          <ListItemText primary="Dashboard" />
        </ListItem>
         <ListItem button component={Link}  to="/index">
          <ListItemIcon>
            <PostAdd />
          </ListItemIcon>
          <ListItemText primary="Index" />
        </ListItem>
        
        <ListItem button component={Link} to="/posts">
          <ListItemIcon>
            <PostAdd />
          </ListItemIcon>
          <ListItemText primary="Posts" />
        </ListItem>

        <ListItem button onClick={handleLogout}>
          <ListItemIcon>
            <Logout />
          </ListItemIcon>
          <ListItemText primary="Logout" />
        </ListItem>
      </List>
    </Drawer>
  );
};

export default Sidebar;
