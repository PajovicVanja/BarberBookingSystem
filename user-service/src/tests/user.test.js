// ===== File: src/tests/user.test.js =====
const request = require('supertest');
const app = require('../../server');
const User = require('../../src/models/userModel');
const { sequelize } = require('../../src/config/db');

describe('User Service API', () => {
    beforeAll(async () => {
        await sequelize.createSchema('userservice', { ifNotExists: true });
        await User.sync({ force: true });
    });

    let token;
    let userId;

    test('POST /api/users/register - should register a user', async () => {
        const res = await request(app)
            .post('/api/users/register')
            .send({
                username: 'testuser',
                email: 'testuser@example.com',
                password: 'password123',
                role: 'customer'
            });
        expect(res.statusCode).toEqual(201);
        expect(res.body).toHaveProperty('user');
        userId = res.body.user.id;
    });

    test('POST /api/users/login - should login and return token', async () => {
        const res = await request(app)
            .post('/api/users/login')
            .send({
                username: 'testuser',
                password: 'password123'
            });
        expect(res.statusCode).toEqual(200);
        expect(res.body).toHaveProperty('token');
        token = res.body.token;
    });

    test('GET /api/users/profile - should return user profile', async () => {
        const res = await request(app)
            .get('/api/users/profile')
            .set('Authorization', `Bearer ${token}`);
        expect(res.statusCode).toEqual(200);
        expect(res.body.username).toEqual('testuser');
    });

    test('PATCH /api/users/profile - should update user profile', async () => {
        const res = await request(app)
            .patch('/api/users/profile')
            .set('Authorization', `Bearer ${token}`)
            .send({ username: 'updateduser' });
        expect(res.statusCode).toEqual(200);
    });

    test('DELETE /api/users/:id - should delete a user', async () => {
        // Note: In a production scenario, deletion should be limited to admin roles.
        const res = await request(app)
            .delete(`/api/users/${userId}`)
            .set('Authorization', `Bearer ${token}`);
        expect(res.statusCode).toEqual(200);
    });

    // Close the database connection after all tests finish
    afterAll(async () => {
        await sequelize.close();
    });
});
