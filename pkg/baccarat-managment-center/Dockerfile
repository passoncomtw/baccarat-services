FROM node:18.15.0-slim

LABEL version="9.2.0"
LABEL description="Docker file for SCC State Server"

RUN mkdir -p /usr/src/
WORKDIR /usr/src/
COPY . /usr/src/

RUN npm install .

EXPOSE 7788
RUN ls /usr/src/

CMD ["npm", "start"]
