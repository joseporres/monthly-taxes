# Use an official Node.js runtime as a parent image with a version that supports ESM
FROM node:16

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install the application dependencies
RUN npm install

# Install Vite globally
RUN npm install -g vite

# Copy the local project files to the container
COPY . .

# Expose the port that the application will run on
EXPOSE 3000

# Use npx to run Vite, ensuring it's in the PATH
CMD ["npx", "vite", "preview"]
