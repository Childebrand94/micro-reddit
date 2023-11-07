import React, { useState, useCallback } from "react";
import { Filter } from "../utils/type";
import { debounce } from "../utils/helpers";

type Props = {
    fetchPosts: (filter: Filter, search: string | null) => void;
};

const NavSearch: React.FC<Props> = ({ fetchPosts }) => {
    const [searchInput, setSearchInput] = useState<string>("");

    const debouncedFetchPosts = useCallback(
        debounce((searchValue: string) => fetchPosts("hot", searchValue), 500),
        [],
    );

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setSearchInput(e.target.value);
        debouncedFetchPosts(e.target.value);
    };

    return (
        <div>
            <input
                className="h-5/6 mx-4 w-5/6 rounded-xl p-2"
                type="text"
                placeholder="Search here"
                onChange={handleChange}
                value={searchInput}
            />
        </div>
    );
};

export default NavSearch;
