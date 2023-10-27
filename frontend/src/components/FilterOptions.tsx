import { useState } from "react";
import { Link } from "react-router-dom";

type filter = "hot" | "top" | "new";

const FilterOptions: React.FC = (): JSX.Element => {
    const [isOpen, setIsOpen] = useState<boolean>(false);
    const [activeFilter, setActiveFilter] = useState<filter>("hot");

    const handleDropDown = () => {
        setIsOpen(!isOpen);
    };

    const handleLinkClick = (text: filter): void => {
        handleDropDown();
        setActiveFilter(text);
    };

    const options: filter[] = ["hot", "top", "new"];

    return (
        <div className="flex flex-col relative">
            <div className="flex justify-center text-center w-14 rounded-xl bg-blue-400">
                <button
                    className="text-xl p-1"
                    onClick={() => setIsOpen(!isOpen)}
                >
                    {activeFilter.toUpperCase()}
                </button>
            </div>

            <div
                className={`flex flex-col text-xl w-14 rounded-xl bg-blue-400 absolute overflow-hidden transition-max-height duration-300 ease-in ${
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