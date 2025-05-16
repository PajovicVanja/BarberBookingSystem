// users-mf/src/App.js
import './styles.css';                 // ← add this
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Routes, Route, Navigate, useNavigate } from 'react-router-dom';

const api = axios.create({
  baseURL: 'http://localhost:4000/api/users',
});

function decodeToken(token) {
  try {
    return JSON.parse(atob(token.split('.')[1]));
  } catch {
    return {};
  }
}

function Layout({ title, children }) {
  return (
    <div className="container">
      <div className="card">
        <h2 className="icon-scissors">{title}</h2>
        {children}
      </div>
    </div>
  );
}

function Register() {
  const [form, setForm] = useState({
    username: '', email: '', password: '', role: 'customer'
  });
  const navigate = useNavigate();

  const handleSubmit = async e => {
    e.preventDefault();
    await api.post('/register', form);
    alert('Registered! Now login.');
    navigate('/login');
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>
        Username
        <input
          placeholder="Username"
          onChange={e => setForm({ ...form, username: e.target.value })}
          required
        />
      </label>

      <label>
        Email
        <input
          placeholder="Email"
          onChange={e => setForm({ ...form, email: e.target.value })}
          required
        />
      </label>

      <label>
        Password
        <input
          type="password"
          placeholder="Password"
          onChange={e => setForm({ ...form, password: e.target.value })}
          required
        />
      </label>

      <label>
        Role
        <select onChange={e => setForm({ ...form, role: e.target.value })}>
          <option value="customer">Customer</option>
          <option value="barber">Barber</option>
        </select>
      </label>

      <button type="submit">Register</button>
    </form>
  );
}

function Login() {
  const [form, setForm] = useState({ username: '', password: '' });
  const navigate = useNavigate();

  const handleSubmit = async e => {
    e.preventDefault();
    const res = await api.post('/login', form);
    localStorage.setItem('token', res.data.token);
    alert('Login successful');
    navigate('/profile');
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>
        Username
        <input
          placeholder="Username"
          onChange={e => setForm({ ...form, username: e.target.value })}
          required
        />
      </label>

      <label>
        Password
        <input
          type="password"
          placeholder="Password"
          onChange={e => setForm({ ...form, password: e.target.value })}
          required
        />
      </label>

      <button type="submit">Login</button>
    </form>
  );
}

function Profile() {
  const [profile, setProfile] = useState(null);
  const token = localStorage.getItem('token');

  useEffect(() => {
    api.get('/profile', { headers: { Authorization: `Bearer ${token}` } })
       .then(res => setProfile(res.data))
       .catch(console.error);
  }, [token]);

  if (!profile) return <p>Loading...</p>;

  return (
    <>
      <p><strong>Username:</strong> {profile.username}</p>
      <p><strong>Email:</strong>    {profile.email}</p>
      <p><strong>Role:</strong>     {profile.role}</p>
    </>
  );
}

export default function App() {
  return (
    <Routes>
      <Route index element={<Navigate to="login" replace />} />
      <Route path="login"    element={<Layout title="Login"><Login /></Layout>} />
      <Route path="register" element={<Layout title="Register"><Register /></Layout>} />
      <Route path="profile"  element={<Layout title="Profile"><Profile /></Layout>} />
      <Route path="*"        element={<Navigate to="login" replace />} />
    </Routes>
  );
}
