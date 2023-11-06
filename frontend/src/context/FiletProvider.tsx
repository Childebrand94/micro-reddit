import React, { useState, ReactNode } from "react";
import { Filter } from "../utils/type";
import { FilterContext } from "./context";

type FilterProviderProps = {
    children: ReactNode;
};

export const FilterProvider: React.FC<FilterProviderProps> = ({ children }) => {
    const [filter, setFilter] = useState<Filter>("hot");
    const [updateTrigger, setUpdateTrigger] = useState<number>(0);

    return (
        <FilterContext.Provider
            value={{ filter, setFilter, updateTrigger, setUpdateTrigger }}
        >
            {children}
        </FilterContext.Provider>
    );
};
