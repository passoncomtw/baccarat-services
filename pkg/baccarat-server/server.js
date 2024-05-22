import http from "http";
import eetase from "eetase";
import socketClusterServer from "socketcluster-server";
import express from "express";
import serveStatic from "serve-static";
import path from "path";
import morgan from "morgan";
import { v4 as uuidv4 } from "uuid";
import sccBrokerClient from "scc-broker-client";
import { fileURLToPath } from "url";
import dotenvFlow from "dotenv-flow";

dotenvFlow.config();

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const ENVIRONMENT = process.env.ENV || "dev";
const SOCKETCLUSTER_PORT = process.env.SOCKETCLUSTER_PORT || 8000;
const SOCKETCLUSTER_WS_ENGINE = process.env.SOCKETCLUSTER_WS_ENGINE || "ws";
const SOCKETCLUSTER_SOCKET_CHANNEL_LIMIT =
  Number(process.env.SOCKETCLUSTER_SOCKET_CHANNEL_LIMIT) || 1000;
const SOCKETCLUSTER_LOG_LEVEL = process.env.SOCKETCLUSTER_LOG_LEVEL || 2;

const SCC_INSTANCE_ID = uuidv4();
const SCC_STATE_SERVER_HOST = process.env.SCC_STATE_SERVER_HOST || null;
const SCC_STATE_SERVER_PORT = process.env.SCC_STATE_SERVER_PORT || null;
const SCC_MAPPING_ENGINE = process.env.SCC_MAPPING_ENGINE || null;
const SCC_CLIENT_POOL_SIZE = process.env.SCC_CLIENT_POOL_SIZE || null;
const SCC_AUTH_KEY = process.env.SCC_AUTH_KEY || null;
const SCC_INSTANCE_IP = process.env.SCC_INSTANCE_IP || null;
const SCC_INSTANCE_IP_FAMILY = process.env.SCC_INSTANCE_IP_FAMILY || null;
const SCC_STATE_SERVER_CONNECT_TIMEOUT =
  Number(process.env.SCC_STATE_SERVER_CONNECT_TIMEOUT) || null;
const SCC_STATE_SERVER_ACK_TIMEOUT =
  Number(process.env.SCC_STATE_SERVER_ACK_TIMEOUT) || null;
const SCC_STATE_SERVER_RECONNECT_RANDOMNESS =
  Number(process.env.SCC_STATE_SERVER_RECONNECT_RANDOMNESS) || null;
const SCC_PUB_SUB_BATCH_DURATION =
  Number(process.env.SCC_PUB_SUB_BATCH_DURATION) || null;
const SCC_BROKER_RETRY_DELAY =
  Number(process.env.SCC_BROKER_RETRY_DELAY) || null;

let agOptions = {};

if (process.env.SOCKETCLUSTER_OPTIONS) {
  let envOptions = JSON.parse(process.env.SOCKETCLUSTER_OPTIONS);
  Object.assign(agOptions, envOptions);
}

let httpServer = eetase(http.createServer());
console.log("🚀 ~ agOptions:", agOptions);
let agServer = socketClusterServer.attach(httpServer, agOptions);

agServer.setMiddleware(
  agServer.MIDDLEWARE_INBOUND,
  async (middlewareStream) => {
    for await (let action of middlewareStream) {
      console.log("🚀 ~ forawait ~ action.type:", action.type);

      if (action.type === action.PUBLISH_IN) {
        let authToken = action.socket.authToken;
        if (
          !authToken ||
          !Array.isArray(authToken.channels) ||
          authToken.channels.indexOf(action.channel) === -1
        ) {
          let publishError = new Error(
            `You are not authorized to publish to the ${action.channel} channel`
          );
          publishError.name = "PublishError";
          action.block(publishError);

          continue; // Go to the start of the loop to process the next inbound action.
        }
      }

      // Any unhandled case will be allowed by default.
      action.allow();
    }
  }
);

let expressApp = express();
if (ENVIRONMENT === "dev") {
  // Log every HTTP request. See https://github.com/expressjs/morgan for other
  // available formats.
  expressApp.use(morgan("dev"));
}
expressApp.use(serveStatic(path.resolve(__dirname, "public")));

// Add GET /health-check express route
expressApp.get("/health-check", (req, res) => {
  res.status(200).send("OK");
});

// HTTP request handling loop.
(async () => {
  for await (let requestData of httpServer.listener("request")) {
    expressApp.apply(null, requestData);
  }
})();

// SocketCluster/WebSocket connection handling loop.
(async () => {
  for await (let { socket } of agServer.listener("connection")) {
    // Handle socket connection.
    // socket.setAuthToken("test123");
    for await (let request of socket.procedure("login")) {
      // let passwordHash = sha256(credentials.password);

      // let userQuery = 'SELECT * FROM Users WHERE username = ?';
      // let userData;
      // try {
      //   let rows = await mySQLClient.query(userQuery, [credentials.username]);
      //   userData = rows[0];
      // } catch (error) {
      //   let loginError = new Error(`Could not find a ${credentials.username} user`);
      //   loginError.name = 'LoginError';
      //   request.error(loginError);

      //   return;
      // }

      // let isValidLogin = userData && userData.password === passwordHash;
      // if (!isValidLogin) {
      //   let loginError = new Error('Invalid user credentials');
      //   loginError.name = 'LoginError';
      //   request.error(loginError);

      //   return;
      // }

      // End the 'login' request successfully.
      request.end();

      // This will give the client a token so that they won't
      // have to login again if they lose their connection
      // or revisit the app at a later time.
      socket.setAuthToken({ username: "cluster-1" });
    }
  }
})();

(async () => {
  for await (let { data } of agServer.listener("authentication")) {
    console.log("🚀 ~ forawait ~ data:", data);
  }
})();

httpServer.listen(SOCKETCLUSTER_PORT);

if (SOCKETCLUSTER_LOG_LEVEL >= 1) {
  (async () => {
    for await (let { error } of agServer.listener("error")) {
      console.error(error);
    }
  })();
}

if (SOCKETCLUSTER_LOG_LEVEL >= 2) {
  console.log(
    `   ${colorText("[Active]", 32)} SocketCluster worker with PID ${
      process.pid
    } is listening on port ${SOCKETCLUSTER_PORT}`
  );

  (async () => {
    for await (let { warning } of agServer.listener("warning")) {
      console.warn(warning);
    }
  })();
}

function colorText(message, color) {
  if (color) {
    return `\x1b[${color}m${message}\x1b[0m`;
  }
  return message;
}

if (SCC_STATE_SERVER_HOST) {
  // Setup broker client to connect to SCC.
  let sccClient = sccBrokerClient.attach(agServer.brokerEngine, {
    instanceId: SCC_INSTANCE_ID,
    instancePort: SOCKETCLUSTER_PORT,
    instanceIp: SCC_INSTANCE_IP,
    instanceIpFamily: SCC_INSTANCE_IP_FAMILY,
    pubSubBatchDuration: SCC_PUB_SUB_BATCH_DURATION,
    stateServerHost: SCC_STATE_SERVER_HOST,
    stateServerPort: SCC_STATE_SERVER_PORT,
    mappingEngine: SCC_MAPPING_ENGINE,
    clientPoolSize: SCC_CLIENT_POOL_SIZE,
    authKey: SCC_AUTH_KEY,
    wsEngine: SOCKETCLUSTER_WS_ENGINE,
    socketChannelLimit: SOCKETCLUSTER_SOCKET_CHANNEL_LIMIT,
    stateServerConnectTimeout: SCC_STATE_SERVER_CONNECT_TIMEOUT,
    stateServerAckTimeout: SCC_STATE_SERVER_ACK_TIMEOUT,
    stateServerReconnectRandomness: SCC_STATE_SERVER_RECONNECT_RANDOMNESS,
    brokerRetryDelay: SCC_BROKER_RETRY_DELAY,
  });

  if (SOCKETCLUSTER_LOG_LEVEL >= 1) {
    (async () => {
      for await (let { error } of sccClient.listener("error")) {
        error.name = "SCCError";
      }
    })();
  }
}
