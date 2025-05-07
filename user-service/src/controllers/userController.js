const bcrypt = require('bcrypt');
const jwt = require('jsonwebtoken');
const userRepository = require('../repositories/userRepository');

const secretKey = process.env.JWT_SECRET || 'your_jwt_secret';

exports.register = async (req, res) => {
    try {
        const { username, password, email, role } = req.body;
        const hashedPassword = await bcrypt.hash(password, 10);
        const newUser = await userRepository.createUser({ username, password: hashedPassword, email, role });
        res.status(201).json({ message: 'User registered successfully', user: newUser });
    } catch (err) {
        console.error(err);
        res.status(500).json({ message: 'Error registering user', error: err.message });
    }
};

exports.login = async (req, res) => {
    try {
        const { username, password } = req.body;
        const user = await userRepository.findUserByUsername(username);
        if (!user) return res.status(404).json({ message: 'User not found' });

        const isValid = await bcrypt.compare(password, user.password);
        if (!isValid) return res.status(401).json({ message: 'Invalid credentials' });

        const token = jwt.sign({ id: user.id, username: user.username, role: user.role }, secretKey, { expiresIn: '1h' });
        res.json({ message: 'Login successful', token });
    } catch (err) {
        console.error(err);
        res.status(500).json({ message: 'Error during login', error: err.message });
    }
};

exports.getProfile = async (req, res) => {
    try {
        const user = await userRepository.findUserById(req.user.id);
        if (!user) return res.status(404).json({ message: 'User not found' });
        res.json(user);
    } catch (err) {
        console.error(err);
        res.status(500).json({ message: 'Error retrieving profile', error: err.message });
    }
};

exports.updateProfile = async (req, res) => {
    try {
        const updateData = req.body;
        // Hash password if it is being updated
        if (updateData.password) {
            updateData.password = await bcrypt.hash(updateData.password, 10);
        }
        await userRepository.updateUser(req.user.id, updateData);
        res.json({ message: 'Profile updated successfully' });
    } catch (err) {
        console.error(err);
        res.status(500).json({ message: 'Error updating profile', error: err.message });
    }
};

exports.deleteUser = async (req, res) => {
    try {
        // This endpoint should be protected for admin-only actions.
        const { id } = req.params;
        await userRepository.deleteUser(id);
        res.json({ message: 'User deleted successfully' });
    } catch (err) {
        console.error(err);
        res.status(500).json({ message: 'Error deleting user', error: err.message });
    }
};

exports.getAllUsers = async (req, res) => {
    try {
        const users = await userRepository.getAllUsers();
        res.json(users);
    } catch (err) {
        console.error(err);
        res.status(500).json({ message: 'Error retrieving users', error: err.message });
    }
};
