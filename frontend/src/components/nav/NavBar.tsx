import { Link } from "react-router-dom";
import FilterOptions from "./FilterOptions";
import NavSearch from "./NavSearch";
import { useAuth } from "../../context/UseAuth";
import LogInButton from "../user/LogInButton";
import { Profile } from "./NavProfile";
import { Filter } from "../../utils/type";
import { FcReddit } from "react-icons/fc";
import { useNavigate } from "react-router-dom";

type Props = {
    fetchPosts: (value: Filter, str: string | null) => void;
};

const NavBar: React.FC<Props> = ({ fetchPosts }) => {
    const { loggedIn } = useAuth();
    const navigate = useNavigate()

    const handleClick = () => {
        navigate("/")
        fetchPosts("hot", null)
    }

    return (
        <nav className="grid grid-cols-[1fr_2fr_1fr] gap-4 h-12 bg-blue-100 px-4">
            <div className="col-start-1 flex items-center justify-between mr-2 w-28">
                <button onClick={handleClick}>
                    <FcReddit size={45} />
                </button>
                <p className="invisible absolute mx-1 lg:relative lg:visible tracking-wide text-xl font-bold font-custom">
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
