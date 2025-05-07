const { DataTypes } = require('sequelize');
const { sequelize } = require('../config/db');

const User = sequelize.define('User', {
    id: {
        type: DataTypes.INTEGER,
        primaryKey: true,
        autoIncrement: true,
    },
    username: {
        type: DataTypes.STRING,
        unique: true,
        allowNull: false,
    },
    password: {
        type: DataTypes.STRING,
        allowNull: false,
    },
    role: {
        type: DataTypes.ENUM('customer', 'barber'),
        allowNull: false,
        defaultValue: 'customer',
    },
    email: {
        type: DataTypes.STRING,
        unique: true,
        allowNull: false,
    }
}, {
    tableName: 'users',
    schema: 'userservice'
});

module.exports = User;