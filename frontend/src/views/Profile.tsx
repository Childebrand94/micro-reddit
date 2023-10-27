import NavBar from "../components/NavBar";
import { PostComp } from "../components/PostComp";
import { useEffect, useState } from "react";
import { Post, UserID } from "../utils/type";
import { useParams } from "react-router-dom";
import { ProfileBasic } from "../components/ProfileBasic";

const Profile = () => {
    const [userData, setUserData] = useState<UserID | null>(null);
    const { user_id } = useParams();
    const [toggleView, setToggleView] = useState<boolean>(false);

    const fetchComments = async () => {
        try {
            const response = await fetch(`/api/users/${user_id}`, {
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
        fetchComments();
    }, []);

    const handleSubmittedClick = () => {
        setToggleView(!toggleView);
    };

    return (
        <div>
            <NavBar />
            <div className="border-b-4 border-blue-400 bg-gray-100 w-full my-3 flex">
                <h1 className="text-blue-700 ml-3 font-bold text-xl tracking-wide">
                    {userData?.username}
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
                    {userData !== null ? (
                        userData.posts.map((p: Post) => {
                            return (
                                <PostComp
                                    key={userData.id}
                                    post={p}
                                    index={null}
                                />
                            );
                        })
                    ) : (
                        <h1>User Not Found</h1>
                    )}
                </div>
            ) : (
                <ProfileBasic />
            )}
        </div>
    );
};
export default Profile;
