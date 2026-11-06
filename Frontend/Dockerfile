FROM node:24-alpine

WORKDIR /app

# Install dependencies first: this layer is cached until package.json changes
COPY package*.json ./
RUN npm install

# Then copy the source code
COPY . .

EXPOSE 5173

CMD ["npm", "run", "dev"]
