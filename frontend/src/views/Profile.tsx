import NavBar from "../components/NavBar";
import { CommentList } from "../components/CommentList";
import { useEffect, useState } from "react";
import { Comment as CommentType } from "../utils/type";
import { useParams } from "react-router-dom";

const Profile = () => {
    const [commentData, setCommentData] = useState<CommentType[]>([]);
    const { user_id } = useParams();

    const fetchComments = async () => {
        try {
            const response = await fetch(`/api/users/${user_id}/comments`, {
                method: "GET",
            });
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }
            const data = await response.json();
            console.log(data);
            setCommentData([...data]);
        } catch (error) {
            console.error("Error:", error);
        }
    };
    useEffect(() => {
        fetchComments();
    }, []);

    console.log(commentData)
    return (
        <div>
            <NavBar />
            <h1 className="text-blue-300 m-3 font-semibold">UserName</h1>
            <div className="bg-blue-400 w-full h-1"></div>
            {/* <CommentList /> */}
        </div>
    );
};
export default Profile;
