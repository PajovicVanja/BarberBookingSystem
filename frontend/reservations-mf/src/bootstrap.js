import './styles.css';
import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import { BrowserRouter } from 'react-router-dom';

const mount = el => {
  ReactDOM.createRoot(el).render(
    <BrowserRouter>
      <App />
    </BrowserRouter>
  );
};

export default mount;

if (process.env.NODE_ENV === 'development') {
  const root = document.getElementById('root');
  if (root) mount(root);
}
