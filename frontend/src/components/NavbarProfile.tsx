import { useState } from "react";
import { Link, redirect } from "react-router-dom";
import { useAuth } from "../context/UseAuth";

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
                throw new Error("Network respose was not ok");
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

    return (
        <div className="realative w-9">
            <button onClick={toggleDropdown}>
                <img className="pt-2" src="/assets/user.png" />
            </button>
            {isDropdownOpen ? (
                <div className="absolute flex flex-col bg-blue-100 p-3 right-0 items-start">
                    <Link
                        onClick={toggleDropdown}
                        className="hover:bg-blue-300 w-full rounded-lg p-1"
                        to={`/users/${userId}`}
                    >
                        Profile
                    </Link>
                    <Link
                        onClick={toggleDropdown}
                        to="/posts"
                        className="hover:bg-blue-300 w-full rounded-lg p-1"
                    >
                        Create Post
                    </Link>
                    <button
                        className="hover:bg-blue-300 w-full rounded-lg p-1 text-left"
                        onClick={handleClick}
                    >
                        Signout
                    </button>
                </div>
            ) : (
                <div></div>
            )}
        </div>
    );
};
