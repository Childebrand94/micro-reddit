import NavBar from "../components/nav/NavBar";
import { PostComp } from "../components/post/PostComp";
import { useEffect, useState } from "react";
import { Post, User } from "../utils/type";
import { ProfileBasic } from "../components/ProfileBasic";
import { useParams } from "react-router-dom";

const Profile = () => {
    const [userPostData, setUserPostData] = useState<Post[] | null>(null);
    const [userData, setUserData] = useState<User | null>(null);
    const { user_id } = useParams();
    const [toggleView, setToggleView] = useState<boolean>(false);

    const fetchUsersPosts = async () => {
        try {
            const response = await fetch(`/api/users/${user_id}/posts`, {
                method: "GET",
            });
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }
            const data = await response.json();
            setUserPostData(data);
        } catch (error) {
            console.error("Error:", error);
        }
    };
    const fetchUser = async () => {
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
        fetchUsersPosts();
        fetchUser();
    }, [user_id]);

    const handleSubmittedClick = () => {
        setToggleView(!toggleView);
    };

    return (
        <div className="flex flex-col h-screen">
            <NavBar fetchPosts={()=>{}}/>
            <div className="border-b-4 border-blue-400 bg-gray-100 w-full my-3 flex">
                <h1 className="text-blue-700 ml-3 font-bold text-xl tracking-wide">
                    {userData ? userData.username : "Username not found"}
                </h1>
                <div className="ml-4 flex">
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
                <div className="flex flex-grow justify-center bg-gray-200 mt-2">
                    {userPostData !== null ? (
                        userPostData.map((p: Post) => {
                            return (
                                <PostComp key={p.id} post={p} index={null} />
                            );
                        })
                    ) : (
                        <h1>No Posts</h1>
                    )}
                </div>
            ) : (
                <ProfileBasic />
            )}
        </div>
    );
};
export default Profile;
