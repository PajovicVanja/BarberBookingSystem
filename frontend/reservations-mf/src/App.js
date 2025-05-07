import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Routes, Route } from 'react-router-dom';

const api = axios.create({
  baseURL: 'http://localhost:4000/api/reservations',
});

function CreateReservation() {
  const [form, setForm] = useState({ user_id: '', barber_id: '', appointment_time: '' });

  const handleSubmit = async (e) => {
    e.preventDefault();
    await api.post('', form);
    alert('Reservation created!');
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Create Reservation</h2>
      <input placeholder="User ID" onChange={(e) => setForm({ ...form, user_id: e.target.value })} required />
      <input placeholder="Barber ID" onChange={(e) => setForm({ ...form, barber_id: e.target.value })} required />
      <input type="datetime-local" onChange={(e) => setForm({ ...form, appointment_time: e.target.value })} required />
      <button type="submit">Create</button>
    </form>
  );
}

function ListReservations() {
  const [reservations, setReservations] = useState([]);
  const [userId, setUserId] = useState('');

  const fetchReservations = async () => {
    const res = await api.get(`/user/${userId}`);
    setReservations(res.data);
  };

  return (
    <div>
      <h2>List Reservations</h2>
      <input placeholder="User ID" onChange={(e) => setUserId(e.target.value)} />
      <button onClick={fetchReservations}>Fetch</button>
      <ul>
        {reservations.map(r => (
          <li key={r.id}>{r.appointment_time} - {r.status}</li>
        ))}
      </ul>
    </div>
  );
}

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<CreateReservation />} />
      <Route path="/list" element={<ListReservations />} />
    </Routes>
  );
}
