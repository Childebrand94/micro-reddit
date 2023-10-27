import { useState } from "react";
const NavSearch = () => {
    const [searchInput, setSearchInput] = useState("");

    const handleChange = (e) => {
        e.preventDefault();
        setSearchInput(e.target.value);
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
