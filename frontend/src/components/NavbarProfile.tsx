import { useState } from "react";
import { Link, redirect } from "react-router-dom";
import { useAuth } from "../context/UseAuth";
import {FaRedditSquare} from "react-icons/fa"
import { IoIosArrowDown, IoIosArrowUp } from 'react-icons/io';
import { CgProfile } from 'react-icons/cg';
import { IoCreateOutline } from 'react-icons/io5';
import {BiLogOut} from 'react-icons/bi';


export const NavbarProfile = () => {
    const [isDropdownOpen, setIsDropdownOpen] = useState(false);
    const { setLoggedIn, userId } = useAuth();

    const toggleDropdown = () => {
        setIsDropdownOpen(!isDropdownOpen);
    };
    const handleClick = async () => {
        toggleDropdown();
        const url = "/api/users/logout";
        try {
            const response = await fetch(url, {
                method: "Post",
            });

            if (!response.ok) {
                throw new Error("Network response was not ok");
            }

            const data = await response.json();
            console.log(data.message);
        } catch (error) {
            console.log("There was an error submitting the form", error);
        } finally {
            setLoggedIn(false);
            redirect("/api");
        }
    };
  const icons = {
    profile: <CgProfile className="mr-2" />,
    create: <IoCreateOutline className="mr-2" />,
    SignOut: <BiLogOut className="mr-2" />, 
  };

    return (
        <div className="w-9 mr-3">
                <button onClick={toggleDropdown} className="flex text-gray-400 p-1 justify-center items-center hover:border hover:border-blue-200">
                    <FaRedditSquare size={30} className="text-gray-400 mr-1" />
                    {isDropdownOpen? <IoIosArrowUp size={15}/> : <IoIosArrowDown size={15}/> }
                </button>
            {isDropdownOpen ? (
                <div className="absolute flex flex-col bg-blue-100 p-3 right-0 items-start border-2 border-blue-200">
                    <Link
                        onClick={toggleDropdown}
                        className="hover:bg-blue-300 w-full rounded-lg p-1 flex items-center "
                        to={`/users/${userId}`}
                    >
                        {icons.profile}
                         Profile
                    </Link>
                    <Link
                        onClick={toggleDropdown}
                        to="/posts"
                        className="hover:bg-blue-300 w-full rounded-lg p-1 flex items-center" 
                    >
                        {icons.create}
                        Create Post
                    </Link>
                    <button
                        className="hover:bg-blue-300 w-full rounded-lg p-1 text-left flex items-center"
                        onClick={handleClick}
                    >
                        {icons.SignOut}
                        Sign out
                    </button>
                </div>
            ) : (
                <div></div>
            )}
        </div>
    );
};
