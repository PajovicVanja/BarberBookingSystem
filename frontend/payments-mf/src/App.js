import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Routes, Route } from 'react-router-dom';

const api = axios.create({
  baseURL: 'http://localhost:4000/api/payments',
});

function CreatePayment() {
  const [form, setForm] = useState({ user_id: '', reservation_id: '', amount: '', payment_method: 'cash' });

  const handleSubmit = async (e) => {
    e.preventDefault();
    await api.post('', form);
    alert('Payment processed!');
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Create Payment</h2>
      <input placeholder="User ID" onChange={(e) => setForm({ ...form, user_id: e.target.value })} required />
      <input placeholder="Reservation ID" onChange={(e) => setForm({ ...form, reservation_id: e.target.value })} required />
      <input type="number" placeholder="Amount" onChange={(e) => setForm({ ...form, amount: e.target.value })} required />
      <select onChange={(e) => setForm({ ...form, payment_method: e.target.value })}>
        <option value="cash">Cash</option>
        <option value="credit_card">Credit Card</option>
      </select>
      <button type="submit">Pay</button>
    </form>
  );
}

function ListPayments() {
  const [payments, setPayments] = useState([]);
  const [userId, setUserId] = useState('');

  const fetchPayments = async () => {
    const res = await api.get(`/user/${userId}`);
    setPayments(res.data);
  };

  return (
    <div>
      <h2>List Payments</h2>
      <input placeholder="User ID" onChange={(e) => setUserId(e.target.value)} />
      <button onClick={fetchPayments}>Fetch</button>
      <ul>
        {payments.map(p => (
          <li key={p.id}>{p.amount}€ - {p.status}</li>
        ))}
      </ul>
    </div>
  );
}

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<CreatePayment />} />
      <Route path="/list" element={<ListPayments />} />
    </Routes>
  );
}
