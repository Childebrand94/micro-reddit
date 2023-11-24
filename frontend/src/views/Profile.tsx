import NavBar from "../components/nav/NavBar";
import { PostComp } from "../components/post/PostComp";
import { useEffect, useState } from "react";
import { Post, User, Filter } from "../utils/type";
import { ProfileBasic } from "../components/ProfileBasic";
import { useParams } from "react-router-dom";
import { useFilter } from "../context/UseFilter";
import { baseUrl } from "../utils/helpers";

type Props = {
    fetchPosts: (value: Filter, str: string | null) => void;
};

const Profile: React.FC<Props> = ({ fetchPosts }) => {
    const [userPostData, setUserPostData] = useState<Post[] | null>(null);
    const [userData, setUserData] = useState<User | null>(null);
    const { user_id } = useParams();
    const [toggleView, setToggleView] = useState<boolean>(false);
    const { updateTrigger } = useFilter();

    const fetchUsersPosts = async () => {
        try {
            const response = await fetch(`${baseUrl}/users/${user_id}/posts`, {
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
            const response = await fetch(`${baseUrl}/users/${user_id}`, {
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
    }, [user_id, updateTrigger]);

    const handleSubmittedClick = () => {
        setToggleView(!toggleView);
    };

    return (
        <div className="flex flex-col bg-gray-200 h-screen">
            <NavBar fetchPosts={fetchPosts} />
            <div className="border-b-4 border-blue-400 bg-gray-100 w-full my-3 flex">
                <h1 className="text-blue-600 ml-3 font-bold text-xl tracking-wide">
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
                <div className="flex flex-col flex-grow items-center w-full bg-gray-200 mt-2">
                    {userPostData !== null ? (
                        userPostData.map((p: Post) => {
                            return (
                                <PostComp key={p.id} post={p} index={null} />
                            );
                        })
                    ) : (
                        <div className="flex justify-center w-full text-center ">
                            <h1 className="text-xl">No Posts Yet...</h1>
                        </div>
                    )}
                </div>
            ) : (
                <ProfileBasic />
            )}
        </div>
    );
};
export default Profile;
