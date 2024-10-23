import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate, useLocation } from 'react-router-dom';
import Sidebar from './components/Sidebar';
import Dashboard from './components/Dashboard';
import Posts from './components/Posts';
import AddPost from './pages/AddPost';
import Login from './components/Login';
import IndexUpdate from './pages/IndexUpdate'; // Import the new component
import { Box } from '@mui/material';
import { AuthProvider, useAuth } from './context/AuthContext';

const ProtectedRoute: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const { isAuthenticated } = useAuth();
  return isAuthenticated ? <>{children}</> : <Navigate to="/login" />;
};

const MainLayout: React.FC = () => {
  const location = useLocation();
  const isLoginPage = location.pathname === '/login';

  return (
    <Box sx={{ display: 'flex' }}>
      {!isLoginPage && <Sidebar />}
      <Box component="main" sx={{ flexGrow: 1, bgcolor: 'background.default', p: 3 }}>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route 
            path="/" 
            element={
              <ProtectedRoute>
                <Dashboard />
              </ProtectedRoute>
            } 
          />
          <Route 
            path="/posts" 
            element={
              <ProtectedRoute>
                <Posts />
              </ProtectedRoute>
            } 
          />
          <Route path="/index" element={<IndexUpdate />} /> {/* Add the route */}
          <Route path="/add-post" element={<ProtectedRoute><AddPost /></ProtectedRoute>} />
          <Route path="/add-post/:id" element={<ProtectedRoute><AddPost /></ProtectedRoute>} />
        </Routes>
      </Box>
    </Box>
  );
};

const App: React.FC = () => {
  return (
    <AuthProvider>
      <Router>
        <MainLayout />
      </Router>
    </AuthProvider>
  );
};

export default App;
