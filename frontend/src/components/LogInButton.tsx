import { Link } from "react-router-dom";

const LogInButton = () => {
    return (
        <div>
            <Link
                to={"/users"}
                className="bg-blue-400 font-semibold tracking-wider py-1 px-2 rounded-xl text-xl hover:bg-blue-500"
            >
                LogIn
            </Link>
        </div>
    );
};
export default LogInButton;
