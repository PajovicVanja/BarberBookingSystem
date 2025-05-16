// reservations-mf/src/App.js
import './styles.css';                  
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Routes, Route } from 'react-router-dom';

const reservationApi = axios.create({
  baseURL: 'http://localhost:4000/api/reservations',
});
const usersApi = axios.create({
  baseURL: 'http://localhost:4000/api/users',
});

function decodeToken(token) {
  try {
    return JSON.parse(atob(token.split('.')[1]));
  } catch {
    return {};
  }
}

function CreateReservation() {
  const token   = localStorage.getItem('token');
  const { id: user_id } = decodeToken(token);
  const [barbers, setBarbers] = useState([]);
  const [form, setForm] = useState({
    user_id:          user_id || '',
    barber_id:        '',
    appointment_time: '',
  });

  useEffect(() => {
    if (token) {
      usersApi
        .get('/barbers', { headers: { Authorization: `Bearer ${token}` } })
        .then(res => setBarbers(res.data))
        .catch(console.error);
      setForm(f => ({ ...f, user_id }));
    }
  }, [token, user_id]);

  const handleSubmit = async e => {
    e.preventDefault();
    try {
      const res = await reservationApi.post(
        '',
        {
          user_id:          String(user_id),
          barber_id:        String(form.barber_id),
          appointment_time: new Date(form.appointment_time).toISOString(),
        },
        { headers: { Authorization: `Bearer ${token}` } }
      );
      alert('Reservation created: ' + JSON.stringify(res.data));
    } catch (err) {
      console.error(err);
      alert('Failed to create reservation:\n' +
            JSON.stringify(err.response?.data, null, 2));
    }
  };

  return (
    <div className="container">
      <div className="card">
        <h2 className="icon-scissors">Create Reservation</h2>
        <p><strong>User:</strong> {user_id}</p>
        <form onSubmit={handleSubmit}>
          <label>
            Barber:
            <select
              value={form.barber_id}
              onChange={e => setForm({ ...form, barber_id: e.target.value })}
              required
            >
              <option value="">— choose barber —</option>
              {barbers.map(b => (
                <option key={b.id} value={b.id}>{b.username}</option>
              ))}
            </select>
          </label>

          <label>
            Appointment Time:
            <input
              type="datetime-local"
              value={form.appointment_time}
              onChange={e => setForm({ ...form, appointment_time: e.target.value })}
              required
            />
          </label>

          <button type="submit">Create</button>
        </form>
      </div>
    </div>
  );
}

function ListReservations() {
  const token   = localStorage.getItem('token');
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
    <div className="container">
      <div className="card">
        <h2 className="icon-scissors">My Reservations</h2>
        <ul>
          {reservations.map(r => (
            <li key={r.id} style={{ marginBottom: 12 }}>
              {new Date(r.appointment_time).toLocaleString()} — {r.status}
              {r.status === 'declined' && <div>{r.message}</div>}
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}

function BarberReservations() {
  const token   = localStorage.getItem('token');
  const { id: barber_id } = decodeToken(token);
  const [reservations, setReservations] = useState([]);
  const [msgs, setMsgs] = useState({});

  useEffect(() => {
    if (barber_id) {
      reservationApi
        .get(`/barber/${barber_id}`, { headers: { Authorization: `Bearer ${token}` } })
        .then(res => setReservations(res.data))
        .catch(console.error);
    }
  }, [barber_id, token]);

  function respond(resId, accept) {
    const message = msgs[resId] || '';
    reservationApi
      .patch(`/${resId}`,
             { status: accept ? 'accepted' : 'declined', message },
             { headers: { Authorization: `Bearer ${token}` } })
      .then(() => {
        setReservations(rs =>
          rs.map(r =>
            r.id === resId
              ? { ...r, status: accept ? 'accepted' : 'declined', message }
              : r
          )
        );
      })
      .catch(console.error);
  }

  const handleDelete = async id => {
    try {
      await reservationApi.delete(`/${id}`, {
        headers: { Authorization: `Bearer ${token}` }
      });
      setReservations(rs => rs.filter(r => r.id !== id));
    } catch (err) {
      console.error(err);
      alert('Failed to delete reservation');
    }
  };

  return (
    <div className="container">
      <div className="card">
        <h2 className="icon-scissors">Incoming Reservations</h2>
        <ul style={{ listStyle: 'none', padding: 0 }}>
          {reservations.map(r => (
            <li key={r.id} style={{ marginBottom: 24, borderBottom: '1px solid #eee', paddingBottom: 12 }}>
              <div>
                Customer #{r.user_id} — {new Date(r.appointment_time).toLocaleString()} — {r.status}
              </div>
              {r.status === 'pending' && (
                <div style={{ marginTop: 8 }}>
                  <button onClick={() => respond(r.id, true)} style={{ marginRight: 8 }}>Accept</button>
                  <textarea
                    placeholder="Decline message"
                    value={msgs[r.id] || ''}
                    onChange={e => setMsgs({ ...msgs, [r.id]: e.target.value })}
                    style={{ verticalAlign: 'top', marginRight: 8 }}
                  />
                  <button onClick={() => respond(r.id, false)}>Decline</button>
                </div>
              )}
              <button onClick={() => handleDelete(r.id)} style={{ marginTop: 8 }}>Delete</button>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}

export default function App() {
  const token = localStorage.getItem('token');
  const { role } = decodeToken(token);

  return (
    <Routes>
      {role === 'barber' ? (
        <Route index element={<BarberReservations />} />
      ) : (
        <>
          <Route index element={<CreateReservation />} />
          <Route path="list" element={<ListReservations />} />
        </>
      )}
    </Routes>
  );
}
