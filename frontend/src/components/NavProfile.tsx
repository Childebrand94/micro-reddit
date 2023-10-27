import { Link } from "react-router-dom";

const NavProfile = () => {
    const loggedIn: boolean = true;
    return (
        <div>
            {loggedIn ? (
                <Link
                    to={"/users"}
                    className="bg-blue-400 py-1 px-2 mb-4 rounded-xl text-xl hover:bg-blue-700"
                >
                    LogIn
                </Link>
            ) : (
                <image path="../assets/logo-reddit.svg" />
            )}
        </div>
    );
};
export default NavProfile;
