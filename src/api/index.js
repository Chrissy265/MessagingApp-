import Header from "./Components/Header/Header.jsx"; 
import ChatHistory from "./ChatHistory.jsx"; 

var socket = new WebSocket("ws://localhost:8080/ws"); 


// Connects to the WebSocket endpoint in question 
//Listen for events such as successful connection
//if any issues such as closed or error socket it will print out to browser 
let connect =cb => {
    console.log("Connecting"); 

    socket.onopen =() => {
        console.log("Successfully Connected");
    };

    socket.onmessage = msg =>{
        console.log(msg);
    };

    socket.onclose = event => {
        console.log("Socket CLosed Connection:", event);
    };

    socket.onerror = error => {
        console.log("Socket Error:", error);
    };
}; 


//Allows to send message from the frontend to backend via WebSocket connection

let sendMsg = msg => {
     console.log("sending msg:", msg); 
     socket.send(msg); //goes to backend
};

export { connect, sendMsg}; 
export default Header; 
export default ChatHistory; 