import { useEffect, useState } from "react";
import NavBar from "../components/NavBar";
import { PostComp } from "../components/PostComp";
import { Post as PostType } from "../utils/type";
import { useParams } from "react-router-dom";
import { CommentList } from "../components/CommentList.tsx";
import { useAuth } from "../context/UseAuth.tsx";
import { CreateCommentForm } from "../components/CreateCommentForm.tsx";

const CommentView = () => {
    const { loggedIn } = useAuth();
    const [postData, setPostData] = useState<PostType | null>(null);
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const { post_id } = useParams();

    const fetchPostByID = async () => {
        try {
            const response = await fetch(`api/posts/${post_id}`, {
                method: "GET",
            });
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }
            const data = await response.json();
            setPostData(data);
        } catch (error) {
            console.error("Error:", error);
        } finally {
            setIsLoading(false);
        }
    };
    useEffect(() => {
        fetchPostByID();
    }, []);
    return (
        <div>
            {postData ? (
                <>
                    <NavBar />
                    {isLoading ? (
                        <h1>Loading...</h1>
                    ) : (
                        <>
                            <PostComp index={null} post={postData} />
                            {loggedIn && (
                                <CreateCommentForm postId={postData.id} />
                            )}
                            <div className="h-2 bg-blue-100 mt-3"></div>
                            {postData.comments ? (
                                <CommentList comments={postData.comments} />
                            ) : (
                                <h1>No Comments Found</h1>
                            )}
                        </>
                    )}
                </>
            ) : (
                <h1>Post Not Found</h1>
            )}
        </div>
    );
};
export default CommentView;
