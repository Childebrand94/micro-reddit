import { useAuth } from "../context/UseAuth";
import { redirect } from "react-router-dom";

export const LogoutButton = () => {
    const { setLoggedIn } = useAuth();
    const handleClick = async () => {
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
        <div>
            <button
                className="bg-blue-400 py-1 px-2 rounded-xl text-xl
                hover:bg-blue-700"
                onClick={handleClick}
            >
                Logout
            </button>
        </div>
    );
};
