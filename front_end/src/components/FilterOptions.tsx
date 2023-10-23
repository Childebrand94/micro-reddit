import { useState } from "react";
import { Link } from "react-router-dom";

const FilterOptions = () => {
  const [isDropDownOpen, setIsDropDownOpen] = useState(false);
  const [displayText, setDisplayText] = useState("HOT");

  const handleDropDown = () => {
    setIsDropDownOpen(!isDropDownOpen);
  };

  const handleLinkClick = (text: string) => {
    handleDropDown();
    setDisplayText(text);
  };

  const options = [
    {
      title: "TOP",
      path: "posts?sort=top",
    },
    {
      title: "NEW",
      path: "posts?sort=new",
    },
    {
      title: "HOT",
      path: "posts?sort=hot",
    },
  ];

  return (
    <div className="flex flex-col">
      <div className="flex justify-center text-center w-14 rounded-xl bg-blue-400">
        <button
          className="text-xl p-1"
          onClick={() => setIsDropDownOpen(!isDropDownOpen)}
        >
          {displayText}
        </button>
      </div>
      {isDropDownOpen && (
        <div className="flex flex-col text-xl">
          {options.map((x) => {
            return (
              <Link onClick={() => handleLinkClick(x.title)} to={"/"}>
                {x.title}
              </Link>
            );
          })}
        </div>
      )}
    </div>
  );
};
export default FilterOptions;
