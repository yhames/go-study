import {WebSocketServer} from "ws";

const wss = new WebSocketServer({port: 8080});

const rooms = {};

wss.on("connection", (socket) => {
  console.log("Connected to WebSocket server...");

  /**
   * Handles incoming messages from clients.
   */
  socket.on("message", (message) => {
    let data = JSON.parse(message);
    switch (data.type) {
      case "join": {
        const roomId = data.roomId;
        if (!rooms[roomId]) {
          rooms[roomId] = new Set();
        }
        rooms[roomId].add(socket);
        socket.roomId = roomId;
        console.log(`Client joined room ${roomId}. Total clients: ${rooms[roomId].size}`);
      }
        break;

      case "signal": {
        const roomId = data.roomId;
        const signalData = data.signalData;
        if (rooms[roomId]) {
          rooms[roomId].forEach(client => {
            if (client !== socket && client.readyState === WebSocket.OPEN) {
              client.send(JSON.stringify(signalData));
            }
          });
        }
      }
        break;

      default:
        console.error("Unknown message type:", data.type);
    }
  });

  /**
   * Handles the closure of a WebSocket connection.
   */
  socket.on("close", () => {
    const id = socket.roomId;
    if (id) {
      rooms[id].delete(socket);
      console.log(`WebSocket closed for room ${id}. Remaining clients: ${rooms[id].size}`);
      if (rooms[id].size === 0) {
        delete rooms[id];
        console.log(`Room ${id} has been deleted.`);
      }
    }
  });

});