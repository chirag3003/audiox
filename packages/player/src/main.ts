const audio = document.querySelector("audio")!;
audio.play()
let socket = new WebSocket("ws://localhost:8080/echo");
let remoteStream = new MediaStream()
socket.onopen = function () {
    socket.send("hello")
};
socket.onmessage = async function (e) {
    audio.pause();
    audio.src = URL.createObjectURL(e.data);
    await audio.play();
};

if(audio?.srcObject) audio.srcObject = remoteStream


