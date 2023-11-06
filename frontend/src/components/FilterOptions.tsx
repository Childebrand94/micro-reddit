import { useState } from "react";
import { Link } from "react-router-dom";
import { Filter } from "../utils/type";
import { useFilter } from "../context/UseFilter";

const FilterOptions = () => {
    const [isOpen, setIsOpen] = useState<boolean>(false);
    const [activeFilter, setActiveFilter] = useState<Filter>("hot");
    const { setFilter, setUpdateTrigger } = useFilter();

    const handleDropDown = () => {
        setIsOpen(!isOpen);
    };

    const handleLinkClick = (text: Filter): void => {
        handleDropDown();
        setActiveFilter(text);
        setFilter(text);
        setUpdateTrigger((prev) => prev + 1);
    };

    const options: Filter[] = ["hot", "top", "new"];

    return (
        <div className="flex flex-col relative">
            <div className="flex justify-center text-center w-14 p-1 rounded-xl bg-blue-400">
                <button className="text-xl" onClick={() => setIsOpen(!isOpen)}>
                    {activeFilter.toUpperCase()}
                </button>
            </div>

            <div
                className={`flex flex-col text-xl w-14 rounded-xl bg-blue-400 absolute overflow-hidden transition-max-height text-center ease-in ${
                    isOpen ? "max-h-60" : "max-h-0"
                }`}
            >
                {options.map((x, i) => (
                    <Link key={i} onClick={() => handleLinkClick(x)} to={"/"}>
                        {x.toUpperCase()}
                    </Link>
                ))}
            </div>
        </div>
    );
};

export default FilterOptions;
