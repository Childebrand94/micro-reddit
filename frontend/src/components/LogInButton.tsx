import { Link } from "react-router-dom";

const LogInButton = () => {
    return (
        <div>
            <Link
                to={"/users"}
                className="bg-blue-400 py-1 px-2 rounded-xl text-xl hover:bg-blue-700"
            >
                LogIn
            </Link>
        </div>
    );
};
export default LogInButton;
