FROM node:16.17.0-alpine
WORKDIR /app
COPY package.json ./
RUN npm install
COPY . .
RUN npm run build-ts
EXPOSE 3002
CMD ["npm", "run", "start"]
