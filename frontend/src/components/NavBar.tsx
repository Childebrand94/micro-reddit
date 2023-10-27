import { Link } from "react-router-dom";
import FilterOptions from "./FilterOptions";
import NavProfile from "./NavProfile";
import NavSearch from "./NavSearch";

const NavBar = () => {
    return (
        <nav className="flex justify-between items-center max-w-full h-12 bg-blue-100 px-4">
            <div className="mr-2">
                <Link to="/">
                    <img
                        className="w-12"
                        src="../../public/assets/logo-reddit.svg"
                        alt="Reddit Logo"
                    />
                </Link>
            </div>
            <FilterOptions />

            <div className="flex items-center">
                <NavSearch />
                <NavProfile />
            </div>
        </nav>
    );
};

export default NavBar;
