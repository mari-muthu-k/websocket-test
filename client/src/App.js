import React, { useState, useEffect } from "react";
import io from "socket.io-client";

const socket = io("http://localhost:5000", {
  transports: ["websocket"],
});

const App = () => {
    const [message, setMessage] = useState("");
    const [messages, setMessages] = useState([]);

    useEffect(() => {
        socket.on("message", (msg) => {
            setMessages([...messages, msg]);
        });

        return () => {
            socket.disconnect();
        };
    }, [messages]);

    const sendMessage = () => {
        socket.emit("message", message);
        setMessage("");
    };

    return (
        <div>
            <h1>Socket.io Chat</h1>
            <div>
                <ul>
                    {messages.map((msg, index) => (
                        <li key={index}>{msg}</li>
                    ))}
                </ul>
            </div>
            <div>
                <input
                    type="text"
                    value={message}
                    onChange={(e) => setMessage(e.target.value)}
                />
                <button onClick={sendMessage}>Send</button>
            </div>
        </div>
    );
};

export default App;