# Base image
FROM node:18-alpine

# Create working directory
WORKDIR /app

# Copy package files and install dependencies
COPY package*.json ./
RUN npm install

# Copy the rest of the app
COPY . .

# Expose the port your service runs on
EXPOSE 3000

# Start the app
CMD ["npm", "start"]
