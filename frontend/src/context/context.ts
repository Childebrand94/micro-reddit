import React from "react";
import { Filter } from "../utils/type";
export type AuthContextType = {
    loggedIn: boolean;
    setLoggedIn: (value: boolean) => void;
    userId: number | null;
};

export type PostFilter = {
    filter: Filter;
    setFilter: (value: Filter) => void;
    updateTrigger: number;
    setUpdateTrigger: (value: number | ((prev: number) => number)) => void;
};

export const AuthContext = React.createContext<AuthContextType | undefined>(
    undefined,
);

export const FilterContext = React.createContext<PostFilter | undefined>(
    undefined,
);
