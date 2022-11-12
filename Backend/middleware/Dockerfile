FROM node:16.17.0-alpine
WORKDIR /app
COPY package.json ./
RUN npm install
COPY . .
RUN npm run build-ts
EXPOSE 5000
CMD ["npm", "run", "start"]
