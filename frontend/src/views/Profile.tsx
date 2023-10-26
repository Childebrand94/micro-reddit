import NavBar from "../components/NavBar";
import { PostComp } from "../components/PostComp";
import { useEffect, useState } from "react";
import { Post, UserID } from "../utils/type";
import { useParams } from "react-router-dom";

const Profile = () => {
    const [userData, setUserData] = useState<UserID | null>(null);
    const { user_id } = useParams();

    const fetchComments = async () => {
        try {
            const response = await fetch(`/api/users/${user_id}`, {
                method: "GET",
            });
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }
            const data = await response.json();
            console.log(data);
            setUserData(data);
        } catch (error) {
            console.error("Error:", error);
        }
    };
    useEffect(() => {
        fetchComments();
    }, []);

    

    return (
        <div>
            <NavBar />
            <h1 className="text-blue-700 m-3 font-bold">{userData?.username}</h1>
            <div className="bg-blue-400 w-full h-1"></div>
            <div>
                {userData !== null ? (
                    userData.posts.map((p: Post) => {
                        return <PostComp key={userData.id} post={p} index={null} />;
                    })
                ) : (
                    <h1>User Not Found</h1>
                )}
            </div>
        </div>
    );
};
export default Profile;
