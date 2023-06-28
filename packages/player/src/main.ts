var socket = new WebSocket("ws://localhost:8080/echo");

socket.onopen = function () {
    console.log("aaaaaaaaaaaaaaaf")
    socket.send("hello")
};

socket.onmessage = function (e) {
    console.log(e)
};