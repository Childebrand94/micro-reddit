import NavBar from "../components/NavBar";
import { PostComp } from "../components/PostComp";
import { useEffect, useState } from "react";
import { Post } from "../utils/type";
import { useParams } from "react-router-dom";
import { ProfileBasic } from "../components/ProfileBasic";

const Profile = () => {
    const [data, setUserData] = useState<Post[] | null>(null);
    const { user_id } = useParams();
    const [toggleView, setToggleView] = useState<boolean>(false);

    const fetchUserData = async () => {
        try {
            const response = await fetch(`/api/users/${user_id}/posts`, {
                method: "GET",
            });
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }
            const data = await response.json();
            setUserData(data);
        } catch (error) {
            console.error("Error:", error);
        }
    };
    useEffect(() => {
        fetchUserData();
    }, []);

    const handleSubmittedClick = () => {
        setToggleView(!toggleView);
    };

    return (
        <div>
            <NavBar />
            <div className="border-b-4 border-blue-400 bg-gray-100 w-full my-3 flex">
                <h1 className="text-blue-700 ml-3 font-bold text-xl tracking-wide">
                    {data ? data[0].author.userName : "Username not found"}
                </h1>
                <div className="ml-4">
                    <button
                        className={`mx-1 px-2 pt-1 ${
                            toggleView
                                ? "bg-gray-100"
                                : "bg-blue-400 text-white"
                        }`}
                        onClick={handleSubmittedClick}
                    >
                        Basic
                    </button>
                    <button
                        className={`mx-1 px-2 pt-1 ${
                            !toggleView
                                ? "bg-gray-100"
                                : "bg-blue-400 text-white"
                        }`}
                        onClick={handleSubmittedClick}
                    >
                        Submitted
                    </button>
                </div>
            </div>

            {toggleView ? (
                <div>
                    {data !== null ? (
                        data.map((p: Post) => {
                            return (
                                <PostComp key={p.id} post={p} index={null} />
                            );
                        })
                    ) : (
                        <h1>User Not Found</h1>
                    )}
                </div>
            ) : (
                <ProfileBasic
                    username={data ? data[0].author.userName : "Unknown"}
                />
            )}
        </div>
    );
};
export default Profile;
