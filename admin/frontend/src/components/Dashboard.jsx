
import { useEffect, useState } from 'react';
import { Link } from "react-router-dom";
import axios from 'axios';
import Nav from './Nav';




export default function Dashboard() {

    async function deleteBlog(id){
        try {
            const res = await axios.post('http://localhost:3000/delete/' + id);
            setTimeout(() => {
                window.location = "/";
            }, 1000)
        } catch(e) {

        }
    }

    const badgeClases = [
        'primary',
        'secondary',
        'success',
        'danger',
        'warning',
        "info",
        "light",
        "dark"
    ];

    const [posts, setPosts] = useState([]);

	useEffect(() => {
        const getPosts = async() => {
            try {
                const res = await axios.get('http://localhost:3000');
                setPosts(res.data.hits.hits);
            } catch(e) {

            }
        };
        getPosts();
	}, []);

  
  return (
        <div>
            <Nav/>
            <br/>
            <br/>
            <br/>
            <div className="row">
                <div className="container-fluid">
                <table className="table">
                    <thead>
                        <tr>
                            <th scope="col">#</th>
                            <th scope="col">Title</th>
                            <th scope="col">Keywords</th>
                            <th scope="col">Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                    {posts.map((value, key) => {
                       return <tr key={key + Math.random()}>
                            <th scope="row">{key+1}</th>
                            <td>{value._source.title}</td>
                            <td>
                                {typeof value._source.keywords == 'string'  && value._source.keywords.split(',').map((value, key) => {
                                    return (
                                        <span key={key + 'Badge' + Math.random()}>
                                        <span   className={'badge bg-' + badgeClases[key]}>{value}</span>
                                        <span> </span>
                                        </span>
                                    )
                                })}
                                
                                
                            </td>
                            <td>
                                Edit

                                <a href="#" onClick={(e) => deleteBlog(value._id)} >Delete</a>
                            </td>
                        </tr> 
                    })}
                        
                    </tbody>
                </table>
                </div>
            </div>
        </div>
  );
}
