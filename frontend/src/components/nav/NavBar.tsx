import { Link } from "react-router-dom";
import FilterOptions from "./FilterOptions";
import NavSearch from "./NavSearch";
import { useAuth } from "../../context/UseAuth";
import LogInButton from "../user/LogInButton";
import { Profile } from "./NavProfile";
import { Filter } from "../../utils/type";
import {FcReddit} from "react-icons/fc"

type Props = {
    fetchPosts: (value: Filter, str: string | null) => void;
};

const NavBar: React.FC<Props> = ({ fetchPosts }) => {
    const { loggedIn } = useAuth();

    return (
        <nav className="grid grid-cols-[1fr_2fr_1fr] gap-4 max-w-full h-12 bg-blue-100 px-4">
            <div className="col-start-1 flex items-center justify-between mr-2 max-w-sm">
                <Link to="/">
                    <FcReddit size={45} />   
                </Link>
                <p className="invisible absolute mx-1 sm:relative sm:visible tracking-wide text-xl font-bold font-custom">
                    reddit
                </p>
            <FilterOptions />
            </div>
            <div className="flex items-center col-start-2 w-full">
                <NavSearch fetchPosts={fetchPosts} />
            </div>
            <div className="flex items-center justify-end col-start-3">
                {loggedIn ? <Profile /> : <LogInButton />}
            </div>
        </nav>
    );
};

export default NavBar;