
import React, { useEffect, useState} from 'react';
import axios from 'axios';

const ListPost = (s) => {
  const [loading, setLoading] = useState(true);
  const [data, setData] = useState([])

  useEffect(() => {
  	// console.log("sdfsdfsdf");
    const fetchData = async () =>{
      setLoading(false);
      try {
        const {data: response} = await axios.get('http://localhost:3000');
        setData(response);
      } catch (error) {
        console.error(error.message);
      }
      setLoading(false);
    }

    fetchData();
  }, []);

  return (
    <div>
    {loading && <div>Loading</div>}
    {!loading && (
      <div>
        <h2>Doing stuff with data</h2>
        {data.map(item => (<span>{item.name}</span>))}
      </div>
    )}
    </div>
  )
}

export default ListPost;