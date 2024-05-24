FROM node:18.15.0-slim

RUN mkdir -p /usr/src/
WORKDIR /usr/src/
COPY . /usr/src/

RUN npm install .

EXPOSE 7788
RUN ls /usr/src/

CMD ["node", "server.js"]
