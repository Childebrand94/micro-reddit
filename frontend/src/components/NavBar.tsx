import { Link } from "react-router-dom";
import FilterOptions from "./FilterOptions";
import NavProfile from "./NavProfile";
import NavSearch from "./NavSearch";

const NavBar = () => {
    return (
        <nav className="flex justify-center items-center max-w-full h-12 bg-blue-100 ">
            <Link to="/">
                <img
                    className="m-1 w-9"
                    src="../../public/assets/logo-reddit.svg"
                    alt="Reddit Logo"
                />
            </Link>
            <FilterOptions />

            <div className="flex items-end">
                <NavSearch />
                <NavProfile />
            </div>
        </nav>
    );
};
export default NavBar;
