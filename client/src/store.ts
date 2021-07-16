import { writable } from "svelte/store";

const messageStore = writable(0);

const socket = new WebSocket("ws://localhost:8081/ws");

console.log("Attempting Connection...");

socket.onopen = () => {
  console.log("Successfully Connected");
};

socket.addEventListener("message", function (event) {
  console.log({ data: event.data });
  messageStore.update((prev) => (prev += +event.data));
});

socket.onclose = (event) => {
  console.log("Socket Closed Connection: ", event);
};

socket.onerror = (error) => {
  console.log("Socket Error: ", error);
};

export default {
  subscribe: messageStore.subscribe,
};
