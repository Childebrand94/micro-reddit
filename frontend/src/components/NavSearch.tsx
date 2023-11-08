import React, { useState, useCallback } from "react";
import { Filter } from "../utils/type";
import { BsSearch } from 'react-icons/bs';
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
    <div className="w-full">
      <div className="relative rounded-xl p-2">
        <input
          className="w-full pl-8 focus:outline-none p-1 rounded-xl"
          type="text"
          placeholder="Search Reddit"
          onChange={handleChange}
          value={searchInput}
        />
        <div className="absolute inset-y-0 left-0 pl-2 flex items-center pointer-events-none">
          <BsSearch className="text-gray-400 ml-2" />
        </div>
      </div>
    </div>
  );
};

export default NavSearch;
