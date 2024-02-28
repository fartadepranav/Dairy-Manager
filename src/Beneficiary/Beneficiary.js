import React from "react";
import { useState } from "react";
import './Beneficiary.css';

export default function Beneficiary() {
  const [name, setName] = useState("");
  const [address, setAddress] = useState("");
  const [phoneNumber, setPhoneNumber] = useState(0);

  const handleSubmit = (event) => {}

  return (
    <form onSubmit={handleSubmit} className="form">
      <div className="form-group">
        <label htmlFor="name">Name:</label>
        <input
          type="text"
          id="name"
          value={name}
          onChange={(event) => setName(event.target.value)}
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
          type="number"
          id="phoneNumber"
          value={phoneNumber}
          onChange={(event) => setPhoneNumber(parseInt(event.target.value))}
        />
      </div>
      <button type="submit">Submit</button>
    </form>
  );
}   