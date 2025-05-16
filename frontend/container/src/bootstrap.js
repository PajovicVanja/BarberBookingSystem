// container/src/bootstrap.js
import './styles.css';
import React, { Suspense, lazy } from 'react';
import ReactDOM from 'react-dom/client';
import {
  BrowserRouter,
  Routes,
  Route,
  Navigate,
  Link,
  useNavigate,
  useLocation,
} from 'react-router-dom';

const UsersApp        = lazy(() => import('UsersMF/App'));
const ReservationsApp = lazy(() => import('ReservationsMF/App'));
const PaymentsApp     = lazy(() => import('PaymentsMF/App'));

const Nav = () => {
  const token = localStorage.getItem('token');
  const navigate = useNavigate();
  const location = useLocation();

  const handleLogout = () => {
    localStorage.removeItem('token');
    navigate('/users/login', { replace: true });
  };

  return (
    <nav className="navbar">
      <div className="logo icon-scissors">Barber Booking</div>
      {!token && <Link to="/users/login">Login</Link>}
      {!token && <Link to="/users/register">Register</Link>}

      {token && (
        <>
          <Link to="/users/profile">Profile</Link>
          <button onClick={handleLogout}>Logout</button>
        </>
      )}

      <Link to="/reservations">Reservations</Link>
      <Link to="/payments">Payments</Link>
    </nav>
  );
};

const Shell = () => (
  <BrowserRouter>
    <Nav />
    <Suspense fallback={<p>Loading…</p>}>
      <Routes>
        <Route path="/" element={<Navigate to="/users/login" replace />} />
        <Route path="/users/*"        element={<UsersApp />} />
        <Route path="/reservations/*" element={<ReservationsApp />} />
        <Route path="/payments/*"     element={<PaymentsApp />} />
      </Routes>
    </Suspense>
  </BrowserRouter>
);

ReactDOM.createRoot(document.getElementById('root')).render(<Shell />);
