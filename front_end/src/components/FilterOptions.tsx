import { useState } from "react";
import { Link } from "react-router-dom";

type filter = "hot" | "top" | "new";

const FilterOptions: React.FC = (): JSX.Element => {
  const [isDropDownOpen, setIsDropDownOpen] = useState<boolean>(false);
  const [activeFilter, setActiveFilter] = useState<filter>("hot");

  const handleDropDown = () => {
    setIsDropDownOpen(!isDropDownOpen);
  };

  const handleLinkClick = (text: filter): void => {
    handleDropDown();
    setActiveFilter(text);
  };

  const options: filter[] = ["hot", "top", "new"];

  return (
    <div className="flex flex-col">
      <div className="flex justify-center text-center w-14 rounded-xl bg-blue-400">
        <button
          className="text-xl p-1"
          onClick={() => setIsDropDownOpen(!isDropDownOpen)}
        >
          {activeFilter.toUpperCase()}
        </button>
      </div>
      {isDropDownOpen && (
        <div className="flex flex-col text-xl">
          {options.map((x, i) => {
            return (
              <Link key={i} onClick={() => handleLinkClick(x)} to={"/"}>
                {x.toUpperCase()}
              </Link>
            );
          })}
        </div>
      )}
    </div>
  );
};
export default FilterOptions;
