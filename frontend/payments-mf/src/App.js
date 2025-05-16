// payments-mf/src/App.js
import './styles.css';
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Routes, Route } from 'react-router-dom';

const paymentsApi = axios.create({
  baseURL: 'http://localhost:4000/api/payments',
});
const reservationsApi = axios.create({
  baseURL: 'http://localhost:4000/api/reservations',
});

function decodeToken(token) {
  try {
    return JSON.parse(atob(token.split('.')[1]));
  } catch {
    return {};
  }
}

function CreatePayment() {
  const token   = localStorage.getItem('token');
  const { id: user_id } = decodeToken(token);
  const [reservations, setReservations] = useState([]);
  const [form, setForm] = useState({
    user_id: user_id || '',
    reservation_id: '',
    amount: '',
    payment_method: 'cash',
  });

  useEffect(() => {
    if (user_id) {
      reservationsApi
        .get(`/user/${user_id}`, { headers: { Authorization: `Bearer ${token}` } })
        .then(res => {
          setReservations(res.data.filter(r => r.status === 'accepted'));
          setForm(f => ({ ...f, user_id }));
        })
        .catch(console.error);
    }
  }, [user_id, token]);

  const handleSubmit = async e => {
    e.preventDefault();
    try {
      const res = await paymentsApi.post(
        '',
        {
          user_id:        Number(form.user_id),
          reservation_id: form.reservation_id,
          amount:         parseFloat(form.amount),
          payment_method: form.payment_method,
        },
        { headers: { Authorization: `Bearer ${token}` } }
      );
      alert('Payment processed: ' + JSON.stringify(res.data));
    } catch (err) {
      console.error(err);
      alert('Failed to process payment:\n' + JSON.stringify(err.response?.data, null, 2));
    }
  };

  return (
    <div className="container">
      <div className="card">
        <h2 className="icon-scissors">Create Payment</h2>
        <p><strong>User:</strong> {user_id}</p>
        <form onSubmit={handleSubmit}>
          <label>
            Reservation:
            <select
              value={form.reservation_id}
              onChange={e => setForm({ ...form, reservation_id: e.target.value })}
              required
            >
              <option value="">— choose reservation —</option>
              {reservations.map(r => (
                <option key={r.id} value={r.id}>
                  {new Date(r.appointment_time).toLocaleString()} ({r.status})
                </option>
              ))}
            </select>
          </label>
          <label>
            Amount (€):
            <input
              type="number"
              value={form.amount}
              onChange={e => setForm({ ...form, amount: e.target.value })}
              required
            />
          </label>
          <label>
            Method:
            <select
              value={form.payment_method}
              onChange={e => setForm({ ...form, payment_method: e.target.value })}
            >
              <option value="cash">Cash</option>
              <option value="credit_card">Credit Card</option>
            </select>
          </label>
          <button type="submit">Pay</button>
        </form>
      </div>
    </div>
  );
}

function ListPayments() {
  const token   = localStorage.getItem('token');
  const { id: user_id } = decodeToken(token);
  const [payments, setPayments] = useState([]);

  useEffect(() => {
    if (user_id) {
      paymentsApi
        .get(`/user/${user_id}`, { headers: { Authorization: `Bearer ${token}` } })
        .then(res => setPayments(res.data))
        .catch(console.error);
    }
  }, [user_id, token]);

  return (
    <div className="container">
      <div className="card">
        <h2 className="icon-scissors">My Payments</h2>
        <ul>
          {payments.map(p => (
            <li key={p.id}>
              €{p.amount.toFixed(2)} — {p.status} (on {new Date(p.created_at).toLocaleString()})
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}

function BarberPayments() {
  const token   = localStorage.getItem('token');
  const { id: barber_id } = decodeToken(token);
  const [payments, setPayments] = useState([]);

  useEffect(() => {
    if (barber_id) {
      paymentsApi
        .get(`/barber/${barber_id}`, { headers: { Authorization: `Bearer ${token}` } })
        .then(res => setPayments(res.data))
        .catch(console.error);
    }
  }, [barber_id, token]);

  const handleDelete = async id => {
    try {
      await paymentsApi.delete(`/${id}`, { headers: { Authorization: `Bearer ${token}` } });
      setPayments(ps => ps.filter(p => p.id !== id));
    } catch (err) {
      console.error(err);
      alert('Failed to delete payment');
    }
  };

  return (
    <div className="container">
      <div className="card">
        <h2 className="icon-scissors">Barber Payments</h2>
        <ul>
          {payments.map(p => (
            <li key={p.id} style={{ marginBottom: 16 }}>
              <div>
                <strong>Customer #{p.user_id}</strong> paid <strong>€{p.amount.toFixed(2)}</strong> via <em>{p.payment_method}</em>
              </div>
              <div><small>{new Date(p.created_at).toLocaleString()}</small></div>
              <button onClick={() => handleDelete(p.id)}>Delete</button>
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
        <Route index element={<BarberPayments />} />
      ) : (
        <>
          <Route index element={<CreatePayment />} />
          <Route path="list" element={<ListPayments />} />
        </>
      )}
    </Routes>
  );
}
