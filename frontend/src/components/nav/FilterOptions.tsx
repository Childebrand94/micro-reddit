import { useState } from 'react';
import { Link } from 'react-router-dom';
import { FaFireFlameCurved } from 'react-icons/fa6';
import { IoIosArrowDown, IoIosArrowUp } from 'react-icons/io';
import { BsStars, BsPlusCircle } from 'react-icons/bs';
import { Filter } from '../../utils/type';
import { useFilter } from '../../context/UseFilter';

const FilterOptions = () => {
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const [activeFilter, setActiveFilter] = useState<Filter>('hot');
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

  
  const filterIcons = {
    hot: <FaFireFlameCurved className="mr-2" />,
    top: <BsStars className="mr-2" />,
    new: <BsPlusCircle className="mr-2" />, 
  };

  return (
    <div className="flex flex-col relative ml-3 ">
      <button
        className={`text-lg font-bold px-3 py-1 flex items-center justify-between w-full ${isOpen? "rounded-t-lg": "rounded-lg" } bg-blue-400`}
        onClick={handleDropDown}
      >
        
        {filterIcons[activeFilter]}
        <div className='invisible absolute sm:relative sm:visible'>
        {activeFilter.toUpperCase()}
        </div>
        {isOpen ? (
          <IoIosArrowUp className="sm:ml-2" size={15} />
        ) : (
          <IoIosArrowDown className="sm:ml-2" size={15} />
        )}

       
      </button>

      {isOpen && (
        <div
          className="flex flex-col text-xl bg-blue-400 rounded-b-xl transition-all ease-in absolute top-full"
          style={{ minWidth: '100%' }}
        >
          {Object.entries(filterIcons).map(([key, icon]) => {
            if(key === activeFilter) return null;

            return (
            <Link
              key={key}
              className={`py-1 pl-2 font-semibold hover:bg-blue-300 flex items-center`}
              onClick={() => handleLinkClick(key as Filter)}
              to={'/'}
            >
              {icon}
              {key.toUpperCase()}
            </Link>
            )
          })}
        </div>
      )}
    </div>
  );
};

export default FilterOptions;

