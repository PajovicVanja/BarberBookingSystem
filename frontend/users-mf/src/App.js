import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Routes, Route, Navigate, useNavigate } from 'react-router-dom';

const api = axios.create({
  baseURL: 'http://localhost:4000/api/users',
});

function Register() {
  const [form, setForm] = useState({ username: '', email: '', password: '', role: 'customer' });
  const navigate = useNavigate();

  const handleSubmit = async e => {
    e.preventDefault();
    await api.post('/register', form);
    alert('Registered! Now login.');
    navigate('/login');    // relative: goes to /users/login
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Register</h2>
      <input placeholder="Username"    onChange={e => setForm({ ...form, username: e.target.value })} required />
      <input placeholder="Email"       onChange={e => setForm({ ...form, email:    e.target.value })} required />
      <input type="password" placeholder="Password" onChange={e => setForm({ ...form, password: e.target.value })} required />
      <select onChange={e => setForm({ ...form, role: e.target.value })}>
        <option value="customer">Customer</option>
        <option value="barber">Barber</option>
      </select>
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
    navigate('/profile');  // relative: goes to /users/profile
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Login</h2>
      <input placeholder="Username"    onChange={e => setForm({ ...form, username: e.target.value })} required />
      <input type="password" placeholder="Password" onChange={e => setForm({ ...form, password: e.target.value })} required />
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
  }, []);

  if (!profile) return <p>Loading...</p>;
  return (
    <div>
      <h2>Profile</h2>
      <p>Username: {profile.username}</p>
      <p>Email:    {profile.email}</p>
      <p>Role:     {profile.role}</p>
    </div>
  );
}

export default function App() {
  return (
    <Routes>
      {/* when at /users      → redirect to /users/login */}
      <Route index       element={<Navigate to="login"    replace />} />

      {/* matches /users/login     */}
      <Route path="login"    element={<Login />} />
      {/* matches /users/register  */}
      <Route path="register" element={<Register />} />
      {/* matches /users/profile   */}
      <Route path="profile"  element={<Profile />} />

      {/* any other /users/* → back to login */}
      <Route path="*" element={<Navigate to="login" replace />} />
    </Routes>
  );
}
