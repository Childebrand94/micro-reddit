import { Link } from "react-router-dom";
import FilterOptions from "./FilterOptions";
import NavSearch from "./NavSearch";
import { useAuth } from "../context/UseAuth";
import LogInButton from "./LogInButton";
import { NavbarProfile } from "./NavbarProfile";

const NavBar = () => {
    const { loggedIn } = useAuth();
    return (
        <nav className="flex justify-between items-center max-w-full h-12 bg-blue-100 px-4">
            <div className="mr-2">
                <Link to="/">
                    <img
                        className="w-12"
                        src="/assets/logo-reddit.svg"
                        alt="Reddit Logo"
                    />
                </Link>
            </div>
            <FilterOptions />

            <div className="flex items-center">
                <NavSearch />
                {loggedIn ? <NavbarProfile /> : <LogInButton />}
            </div>
        </nav>
    );
};

export default NavBar;
