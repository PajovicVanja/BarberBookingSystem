require('dotenv').config();
const express = require('express');
const morgan = require('morgan');
const userRoutes = require('./src/routes/userRoutes');
const { sequelize } = require('./src/config/db');
const swaggerUi = require('swagger-ui-express');
const YAML = require('yamljs');

const swaggerDocument = YAML.load('./src/config/swagger.yaml');

const app = express();

app.use(express.json());
app.use(morgan('dev'));

// API routes
app.use('/api/users', userRoutes);
// Swagger UI for API documentation
app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(swaggerDocument));

const PORT = process.env.PORT || 3000;

// Only start the server if this file is run directly (not imported)
if (require.main === module) {
    sequelize
    .createSchema('userservice', { logging: false }) 
    .then(() => sequelize.sync())
    .then(() => {
        console.log('Database connected');
        app.listen(PORT, () => {
            console.log(`User Service running on port ${PORT}`);
        });
    })
    .catch((err) => {
        console.error('Unable to connect to the database:', err);
    });
}



module.exports = app;
