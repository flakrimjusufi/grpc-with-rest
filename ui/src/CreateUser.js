import React, { useReducer, useState } from 'react';
import axios from "axios";
import './App.css';


const formReducer = (state, event) => {
  if(event.reset) {
   return {
     email: '',
     phoneNumber: '',
     name: ''
   }
 }
  return {
    ...state,
    [event.name]: event.value
  }
}

function CreateUser() {
  const [formData, setFormData] = useReducer(formReducer, {});
  const [submitting, setSubmitting] = useState(false);


  const handleSubmit = event => {
    event.preventDefault();
    setSubmitting(true);

    setTimeout(() => {
      setSubmitting(false);
      setFormData({
       reset: true
     })
    }, 30000);
      axios.post("http://localhost:8090/api/v1/user", {
      name: formData.name,
      email: formData.email,
      phoneNumber: formData.phoneNumber
  }).then(function(response) {
      console.log(response);
  });
  }

  const handleChange = event => {
    const isCheckbox = event.target.type === 'checkbox';
    setFormData({
      name: event.target.name,
      value: isCheckbox ? event.target.checked : event.target.value,
    });
  }

  return(
    <div className="wrapper">
      <h1>Create user</h1>
      {submitting &&
        <div>
          You are submitting the following:
          <ul>
            {Object.entries(formData).map(([name, value]) => (
              <li key={name}><strong>{name}</strong>: {value.toString()}</li>
            ))}
          </ul>
        </div>
      }
      
      <form onSubmit={handleSubmit}>
      <fieldset disabled={submitting}>
          <label>
            <p>Name</p>
            <input name="name" onChange={handleChange} value={formData.name || ''}/>
          </label>
        </fieldset>
        <fieldset disabled={submitting}>
        <label>
            <p>Email</p>
            <input name="email" onChange={handleChange} value={formData.email || ''}/>
          </label>
          <label>
            <p>Phone Number</p>
            <input type="number" name="phoneNumber" onChange={handleChange} step="1" value={formData.phoneNumber || ''}/>
          </label>
        </fieldset>
        <button type="submit" disabled={submitting}>Submit</button>
      </form>
    </div>
  )
}

export default CreateUser;