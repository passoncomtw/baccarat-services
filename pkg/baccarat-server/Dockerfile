FROM node:18.15.0-slim

LABEL maintainer="Tomas Lin"
LABEL version="19.1.2"
LABEL description="Docker file for SocketCluster with support for clustering."

RUN mkdir -p /usr/src/
WORKDIR /usr/src/
COPY . /usr/src/

RUN npm install .

EXPOSE 8223

CMD ["npm", "run", "start:docker"]
