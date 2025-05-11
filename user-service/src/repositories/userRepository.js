const User = require('../models/userModel');

class UserRepository {
    async createUser(userData) {
        return await User.create(userData);
    }

    async findUserByUsername(username) {
        return await User.findOne({ where: { username } });
    }

    async findUserById(id) {
        return await User.findByPk(id);
    }

    async updateUser(id, updateData) {
        return await User.update(updateData, { where: { id } });
    }

    async deleteUser(id) {
        return await User.destroy({ where: { id } });
    }
    
    async getAllUsers() {
        return await User.findAll();
    }

    async getUsersByRole(role) {
        return await User.findAll({ where: { role } });
      }
    
}

module.exports = new UserRepository();
