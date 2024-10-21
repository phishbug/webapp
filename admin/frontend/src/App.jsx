import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import Dashboard from "./components/Dashboard"
import Post from "./components/Post"


import {
  createBrowserRouter,
  RouterProvider,
  Link
} from "react-router-dom";

function App() {

    const router = createBrowserRouter([
      {
        path: "/",
        element: <Dashboard/>,
      },
      {
        path: "/post",
        element: <Post/>,
      },
    ]);

  const [count, setCount] = useState(0)

  return (
    <>
    <div id="dashboard">
        <div className="container-fluid">
            <div className="row">
                <RouterProvider router={router} />
            </div>
        </div>  
    </div>
    </>
  )
}

export default App
