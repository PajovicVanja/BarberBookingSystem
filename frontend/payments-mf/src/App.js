import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Routes, Route } from 'react-router-dom';

// point at your web gateway
const paymentsApi = axios.create({
  baseURL: 'http://localhost:4000/api/payments',
});
const reservationsApi = axios.create({
  baseURL: 'http://localhost:4000/api/reservations',
});

// helper to decode JWT payload
function decodeToken(token) {
  try {
    return JSON.parse(atob(token.split('.')[1]));
  } catch {
    return {};
  }
}

function CreatePayment() {
  const token = localStorage.getItem('token');
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
        .get(`/user/${user_id}`)
        .then(res => {
          // only show accepted
          setReservations(res.data.filter(r => r.status === 'accepted'));
        })
        .catch(console.error);
      setForm(f => ({ ...f, user_id }));
    }
  }, [user_id]);

  const handleSubmit = async e => {
    e.preventDefault();
    const payload = {
      user_id:        Number(form.user_id),
      reservation_id: form.reservation_id,
      amount:         parseFloat(form.amount),
      payment_method: form.payment_method,
    };

    try {
      const res = await paymentsApi.post('', payload, {
        headers: { Authorization: `Bearer ${token}` }
      });
      alert('Payment processed: ' + JSON.stringify(res.data));
    } catch (err) {
      console.error('❌ Payment error response:', err.response?.data || err);
      alert('Failed to process payment:\n' +
            JSON.stringify(err.response?.data, null, 2));
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Create Payment</h2>
      <p><strong>User:</strong> {user_id}</p>

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
  );
}

function ListPayments() {
  const token = localStorage.getItem('token');
  const { id: user_id } = decodeToken(token);
  const [payments, setPayments] = useState([]);

  useEffect(() => {
    if (user_id) {
      paymentsApi
        .get(`/user/${user_id}`, { headers: { Authorization: `Bearer ${token}` } })
        .then(res => setPayments(res.data))
        .catch(console.error);
    }
  }, [user_id]);

  return (
    <div>
      <h2>My Payments</h2>
      <ul>
        {payments.map(p => (
          <li key={p.id}>
            €{p.amount.toFixed(2)} — {p.status} (on {new Date(p.created_at).toLocaleString()})
          </li>
        ))}
      </ul>
    </div>
  );
}

function BarberPayments() {
  const token = localStorage.getItem('token');
  const { id: barber_id } = decodeToken(token);
  const [payments, setPayments] = useState([]);

  useEffect(() => {
    paymentsApi
      .get(`/barber/${barber_id}`, { headers: { Authorization: `Bearer ${token}` } })
      .then(res => setPayments(res.data))
      .catch(console.error);
  }, [barber_id]);

  return (
    <div>
      <h2>Barber Payments</h2>
      <ul>
        {payments.map(p => (
          <li key={p.id}>
            <strong>Customer #{p.user_id}</strong> paid <strong>€{p.amount.toFixed(2)}</strong> via <em>{p.payment_method}</em>
            <br/>
            <small>{new Date(p.created_at).toLocaleString()}</small>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default function App() {
  const token = localStorage.getItem('token');
  const { role } = decodeToken(token);

  return (
    <Routes>
      {role === 'barber' ? (
        <Route path="/" element={<BarberPayments />} />
      ) : (
        <>
          <Route path="/" element={<CreatePayment />} />
          <Route path="list" element={<ListPayments />} />
        </>
      )}
    </Routes>
  );
}
