import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import { BrowserRouter } from 'react-router-dom';

const mount = (el) => {
  ReactDOM.createRoot(el).render(
    <BrowserRouter>
      <App />
    </BrowserRouter>
  );
};

// Export it for Module Federation
export default mount;

// If running standalone (npm start), mount immediately:
if (process.env.NODE_ENV === 'development') {
  const root = document.getElementById('root');
  if (root) {
    mount(root);
  }
}
