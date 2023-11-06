import React from "react";
export type AuthContextType = {
    loggedIn: boolean;
    setLoggedIn: (value: boolean) => void;
    userId: number | null;
};

export const AuthContext = React.createContext<AuthContextType | undefined>(
    undefined,
);
