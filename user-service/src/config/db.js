const { Sequelize } = require('sequelize');

const sequelize = new Sequelize(process.env.DATABASE_URL, {
    dialect: 'postgres',
    logging: false,
    define: {
        schema: 'userservice'
    }
});

module.exports = { sequelize };
