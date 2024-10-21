import type { FunctionComponent } from "react";
import React, { useState } from "react";

import { Input } from '../components/Forms/input'
import { TagsInput } from "react-tag-input-component";
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

export default function Posts() {
    const navigate = useNavigate();
    // const [selected, setSelected] = useState(["papaya"]);
    const [message, setMessage] = useState();

    const [formData, setFormData] = useState({
        title: '',
        metadata: '',
        body: '',
        keywords: '',
    });

    const handleChange = (e) => {
        console.log();
        const { value, name } = e.target;
        setFormData({
          ...formData,
          [name]: value,
        });

    };
    
    const handleSubmit = async (e) => {
        e.preventDefault();
        console.log(formData);
        try {
          const response = await axios.post('http://localhost:3000/create', formData);
          setMessage('Form submitted successfully!');
          console.log(response.data);
          navigate('/dashboard', { replace: true });
        } catch (error) {
          setMessage('There was an error submitting the form.');
          console.error(error);
        }
    };

  return (
    <div id="posts">

      <form onSubmit={handleSubmit}>
        <h1>Add New Post</h1>
        <hr/>
        <div className="mb-3">
          <label for="exampleInputEmail1" className="form-label">Title</label>
          <input 
              type="text"
              name="title"
              className="form-control"
              value={formData.title}
              onChange={(e) => handleChange(e)}
              required
          />

          <div id="emailHelp" className="form-text">Post Title</div>
        </div>


        <div className="mb-3">
          <label for="exampleInputEmail1" className="form-label">Meta</label>
          <textarea 
              type="text"
              name="metadata"
              className="form-control"
              onChange={(e) => handleChange(e)}
              required
          >
              {formData.metadata}
          </textarea>
          <div id="emailHelp" className="form-text">We'll never share your email with anyone else.</div>
        </div>

          


        <div className="mb-3">
            <label for="exampleInputEmail1" className="form-label">Body</label>
            <textarea type="text"
                  name="body"
                  onChange={(e) => handleChange(e)}
                  required
                  className="form-control"
                  rows="5"
            >
                  
              </textarea>
          <div id="emailHelp" className="form-text">We'll never share your email with anyone else.</div>
        </div>

      
        
        <button type="submit" className="btn btn-primary">Submit</button>
      </form> 
    </div>
  );
}
