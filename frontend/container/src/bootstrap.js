import React, { Suspense, lazy } from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter, Routes, Route, Navigate, Link } from 'react-router-dom';

const UsersApp        = lazy(() => import('UsersMF/App'));
const ReservationsApp = lazy(() => import('ReservationsMF/App'));
const PaymentsApp     = lazy(() => import('PaymentsMF/App'));

const Nav = () => (
  <nav className="flex gap-6 p-4 shadow-md">
    <Link to="/users">Users</Link>
    <Link to="/reservations">Reservations</Link>
    <Link to="/payments">Payments</Link>
  </nav>
);

const Shell = () => (
  <BrowserRouter>
    <Nav />
    <Suspense fallback={<p className="p-4">Loading…</p>}>
      <Routes>
        <Route path="/" element={<Navigate to="/users" replace />} />
        <Route path="/users/*"        element={<UsersApp />} />
        <Route path="/reservations/*" element={<ReservationsApp />} />
        <Route path="/payments/*"     element={<PaymentsApp />} />
      </Routes>
    </Suspense>
  </BrowserRouter>
);

ReactDOM.createRoot(document.getElementById('root')).render(<Shell />);
