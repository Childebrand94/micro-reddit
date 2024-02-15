import { useEffect, useState } from "react";

export const LoadingScreen = () => {

    const [message, setMessage] = useState("This can take up to 2 min for the server to wake up...");

    useEffect(() => {
        const messages = [
            "Rousing the hamsters...",
            "Spinning up the hamster wheel...",
            "Brewing some coffee for the server...",
            "Stretching the bytes, please hold on...",
            "Server is doing its morning stretches...",
            "Polishing the pixels, just a sec...",
            "Asking the server to please hurry up...",
            "Our server is slow but our spirits are high!",
            "Warming up the server with a nice cup of tea...",
            "Server's almost ready, doing its final yoga pose...",
        ];

        const intervalId = setInterval(() => {
            setMessage(currentMessage => {
                const nextIndex = (messages.indexOf(currentMessage) + 1) % messages.length;
                return messages[nextIndex];
            });
        }, 10000);

        return () => clearInterval(intervalId);
    }, []);

    return (
        <div className="absolute top-36 left-1/2 transform -translate-x-1/2">
            <p className="text-center">{message}</p>
        </div>
    );
};
