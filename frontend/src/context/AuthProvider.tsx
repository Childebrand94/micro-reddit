import React, { ReactNode, useEffect, useState } from "react";
import { AuthContext } from "./context";
import { baseUrl } from "../utils/helpers";

type AuthProviderProps = {
    children: ReactNode;
};

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
    const [loggedIn, setLoggedIn] = useState<boolean>(true);
    const [userId, setUserId] = useState<number | null>(null);

    const fetchSession = async () => {
        try {
            const response = await fetch(`${baseUrl}/sessions`, {
                method: "GET",
                credentials: 'include',
            });
            if (!response.ok) {
                setLoggedIn(false);
                throw new Error("Network response was not ok");
            }
            const data = await response.json();
            console.log(`this is the data: ${data}`)
            if (data.loggedIn) {
                setLoggedIn(true);
                setUserId(data.userId);

            } else {
                setLoggedIn(false);
                console.log("Please login");
            }
        } catch (error) {
            console.error("Error:", error);
        }
    };
    useEffect(() => {
        fetchSession();
    }, []);

    return (
        <AuthContext.Provider value={{ loggedIn, setLoggedIn, userId, fetchSession }}>
            {children}
        </AuthContext.Provider>
    );
};
