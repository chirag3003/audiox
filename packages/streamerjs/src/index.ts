console.log("Starting");
import WebSocket, { WebSocketServer } from "ws";

const ws = new WebSocketServer({
    port: 8080,
});

ws.on("connection", e => {
    e.send("hello");
});
