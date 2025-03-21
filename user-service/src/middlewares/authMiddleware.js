const jwt = require('jsonwebtoken');
const secretKey = process.env.JWT_SECRET;

module.exports = (req, res, next) => {
    const authHeader = req.headers.authorization;
    if (!authHeader) return res.status(401).json({ message: 'No token provided' });

    const token = authHeader.split(' ')[1];
    if (!token) return res.status(401).json({ message: 'Token format is invalid' });

    jwt.verify(token, secretKey, (err, decoded) => {
        if (err) return res.status(401).json({ message: 'Token is not valid' });
        req.user = decoded;
        next();
    });
};

