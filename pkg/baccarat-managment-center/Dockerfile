FROM node:18.15.0-slim

RUN mkdir -p /usr/src/
WORKDIR /usr/src/
COPY . /usr/src/

RUN npm install .

EXPOSE 3000

CMD ["node", "server.js"]
