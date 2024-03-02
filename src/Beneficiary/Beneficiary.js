import React from "react";
import { useState } from "react";
import './Beneficiary.css';

export default function Beneficiary() {
  const [name,setName]=useState();
  const [address, setAddress] = useState("");
  const [phoneNumber, setPhoneNumber] = useState(0);

  const handleSubmit = (event) => {
    console.log(event)
    event.preventDefault();

    
    const data = {
      address: address,
      name: name,
      phoneNumber: phoneNumber
    }

    const myHeaders = new Headers();
myHeaders.append("Content-Type", "application/json");

const raw = JSON.stringify(data);

const requestOptions = {
  method: "POST",
  headers: myHeaders,
  body: raw,
  redirect: "follow"
};

fetch("http://localhost:8080/api/addCustomer", requestOptions)
  .then((response) => response.text())
  .then((result) => console.log(result))
  .catch((error) => console.error(error));
  
  }

  return (
    <form onSubmit={handleSubmit} className="form">
      <div className="form-group">
        <label htmlFor="name">Name:</label>
        <input
          type="text"
          id="name"
          onChange={(event) => {
            
            if (event.target.value)
            {
              setName(event.target.value);
              console.log(name);}
            else
            { 
              var txt = document.getElementsByName("err");
              txt.innerHTML = "Please input your name"
            }
            }}
        />
      </div>
      <div className="form-group">
        <label htmlFor="address">Address:</label>
        <input
          type="text"
          id="address"
          value={address}
          onChange={(event) => setAddress(event.target.value)}
        />
      </div>
      <div className="form-group">
        <label htmlFor="phoneNumber">Phone Number:</label>
        <input
          type="text"
          id="phoneNumber"
          onChange={(event) => setPhoneNumber(parseInt(event.target.value))}
        />
      </div>
      <button type="submit">Submit</button>

      <div name="err" style={{color: "white"}}>
      </div>
    </form>
  );
}   