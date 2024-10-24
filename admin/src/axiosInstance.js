// src/axiosInstance.js

import axios from 'axios';

// Create an Axios instance
const axiosInstance = axios.create({
  baseURL: `${import.meta.env.VITE_API_URL}`, // Replace with your API base URL
});

// Add a request interceptor
axiosInstance.interceptors.request.use(
  (config) => {
    // Get the token from local storage or any other method you're using
    const token = localStorage.getItem('token'); // or however you store your token

    // If the token exists, add it to the Authorization header
    if (token) {
      config.headers['Authorization'] = `${token}`;
    }

    return config;
  },
  (error) => {
    // Handle the error
    return Promise.reject(error);
  }
);

export default axiosInstance;
