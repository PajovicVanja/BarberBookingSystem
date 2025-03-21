const { Sequelize } = require('sequelize');

const sequelize = new Sequelize('postgres://postgres:123123@localhost:5432/ita', {
    dialect: 'postgres',
    logging: false,
    define: {
        schema: 'user-service'  
    }
});

module.exports = { sequelize };
