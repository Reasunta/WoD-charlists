FROM node:11.10.0
WORKDIR /app
COPY package.json /app
RUN npm install
COPY . /app
CMD ["node","app.js"]