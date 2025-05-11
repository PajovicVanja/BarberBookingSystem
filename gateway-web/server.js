const express = require('express');
const morgan = require('morgan');
const { createProxyMiddleware } = require('http-proxy-middleware');
const cors = require('cors');

const app = express();

app.use(cors());
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

const joinJson = (...objs) => Object.assign({}, ...objs);

// NEW  aggregated route  ──────────────────────────────────────────────
app.get('/api/dashboard/user/:id', async (req, res) => {
  try {
    const uid  = req.params.id;
    const [profile, reservations, payments] = await Promise.all([
      axios.get(`http://user-service:3000/api/users/profile`,  { headers: req.headers, timeout: 2000 }),
      axios.get(`http://reservation-service:8000/api/reservations/user/${uid}`, { timeout: 2000 }),
      axios.get(`http://payment-service:8080/api/payments/user/${uid}`,        { timeout: 2000 })
    ]).then(r => r.map(resp => resp.data));

    res.json(joinJson(
      { profile },
      { reservations },
      { payments }
    ));
  } catch (e) {
    console.error(e.message);
    res.status(502).json({ error: 'aggregation-failed', detail: e.message });
  }
});

app.get('/', (req, res) => {
  res.send('Welcome to the Web API Gateway');
});

const PORT = process.env.PORT || 4000;
app.listen(PORT, () => {
  console.log(`Web Gateway listening on port ${PORT}`);
});
