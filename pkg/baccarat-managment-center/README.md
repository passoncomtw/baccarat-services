# scc-state
Cluster state tracking and notification engine for SCC.

### Usage

```js
SCC_STATE_LOG_LEVEL=1 node server.js
```

### Log levels

 * 3 - log everything
 * 2 - warnings and errors
 * 1 - errors only
 * 0 - log nothing

### Build and deploy to DockerHub

Replace `x.x.x` with the version number.

```
$ docker build -t socketcluster/scc-state:vx.x.x .
```

```
$ docker push socketcluster/scc-state:vx.x.x
```

```
$ docker build . -t passon/baccarat-managment-center --progress=plain -f ./Dockerfile
$ docker run -p 7788:7788 --name baccarat-managment-center -d passon/baccarat-managment-center
```

### Run Docker State Service

```
$ cd pkg/baccarat-managment-center/
$ docker-compose up -d
$ cd -
```
