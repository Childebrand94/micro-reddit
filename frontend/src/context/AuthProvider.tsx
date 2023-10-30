import React, { ReactNode, useState } from "react";
import { AuthContext } from "./context";

type AuthProviderProps = {
    children: ReactNode;
};

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
    const [loggedIn, setLoggedIn] = useState<boolean>(true);

    return (
        <AuthContext.Provider value={{ loggedIn, setLoggedIn }}>
            {children}
        </AuthContext.Provider>
    );
};
