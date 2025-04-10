const express = require('express');
const morgan = require('morgan');
const { createProxyMiddleware } = require('http-proxy-middleware');

const app = express();

app.use(morgan('dev'));

// Update proxies using Docker Compose service names instead of localhost.
app.use('/api/payments', createProxyMiddleware({
  target: 'http://payment-service:8080',
  changeOrigin: true,
  pathRewrite: {
    '^/api/payments': '/api/payments'
  }
}));

app.use('/api/reservations', createProxyMiddleware({
  target: 'http://reservation-service:8000',
  changeOrigin: true,
  pathRewrite: {
    '^/api/reservations': '/api/reservations'
  }
}));

app.use('/api/users', createProxyMiddleware({
  target: 'http://user-service:3000',
  changeOrigin: true,
  pathRewrite: {
    '^/api/users': '/api/users'
  }
}));

app.get('/', (req, res) => {
  res.send('Welcome to the Web API Gateway');
});

const PORT = process.env.PORT || 4000;
app.listen(PORT, () => {
  console.log(`Web Gateway listening on port ${PORT}`);
});
