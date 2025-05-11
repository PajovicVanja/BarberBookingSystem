import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Routes, Route } from 'react-router-dom';

// point at your web gateway
const reservationApi = axios.create({
  baseURL: 'http://localhost:4000/api/reservations',
});
const usersApi = axios.create({
  baseURL: 'http://localhost:4000/api/users',
});

// small helper to decode JWT payload
function decodeToken(token) {
  try {
    return JSON.parse(atob(token.split('.')[1]));
  } catch {
    return {};
  }
}

function CreateReservation() {
  const token = localStorage.getItem('token');
  const { id: user_id } = decodeToken(token);
  const [barbers, setBarbers] = useState([]);
  const [form, setForm] = useState({
    user_id: user_id || '',
    barber_id: '',
    appointment_time: '',
  });

  useEffect(() => {
    // fetch all barbers
    if (token) {
      usersApi
        .get('/barbers', { headers: { Authorization: `Bearer ${token}` } })
        .then(res => setBarbers(res.data))
        .catch(console.error);
    }
  }, [token]);

     const handleSubmit = async e => {
         e.preventDefault();
         // build a full ISO timestamp
              // force both IDs to strings, and build a valid ISO datetime
              const payload = {
                user_id:      String(form.user_id),
                barber_id:    String(form.barber_id),
                appointment_time: new Date(form.appointment_time).toISOString(),
              };
         console.log('→ POST /api/reservations payload:', payload);
     
         try {
           const res = await reservationApi.post('', payload, {
             headers: { Authorization: `Bearer ${token}` },
           });
           alert('Reservation created: ' +  JSON.stringify(res.data));
         } catch (err) {
           // Axios wraps the server body on err.response.data
           console.error('❌ Reservation error response:', err.response?.data || err);
           alert('Failed to create reservation:\n'  +
                 JSON.stringify(err.response?.data, null, 2));
         }
       };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Create Reservation</h2>
      {/* show current user but don’t let them change it */}
      <p><strong>User:</strong> {user_id}</p>

      <label>
        Barber:
        <select
          value={form.barber_id}
          onChange={e => setForm({ ...form, barber_id: e.target.value })}
          required
        >
          <option value="">— choose barber —</option>
          {barbers.map(b => (
            <option key={b.id} value={b.id}>
              {b.username}
            </option>
          ))}
        </select>
      </label>

      <label>
        Appointment Time:
        <input
          type="datetime-local"
          value={form.appointment_time}
          onChange={e =>
            setForm({ ...form, appointment_time: e.target.value })
          }
          required
        />
      </label>

      <button type="submit">Create</button>
    </form>
  );
}

function ListReservations() {
  const token = localStorage.getItem('token');
  const { id: user_id } = decodeToken(token);
  const [reservations, setReservations] = useState([]);

  useEffect(() => {
    if (user_id) {
      reservationApi
        .get(`/user/${user_id}`, { headers: { Authorization: `Bearer ${token}` } })
        .then(res => setReservations(res.data))
        .catch(console.error);
    }
  }, [user_id, token]);

  return (
    <div>
      <h2>My Reservations</h2>
      <ul>
        {reservations.map(r => (
          <li key={r.id}>
            {new Date(r.appointment_time).toLocaleString()} — {r.status}
          </li>
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
