import { useEffect, useState } from "react";
import NavBar from "../components/nav/NavBar.tsx";
import { PostComp } from "../components/post/PostComp.tsx";
import { Post as PostType, Filter, Comment as CommentType } from "../utils/type";
import { useParams } from "react-router-dom";
import { CommentList } from "../components/comments/CommentList.tsx";
import { useAuth } from "../context/UseAuth.tsx";
import { useFilter } from "../context/UseFilter.tsx";
import { CreateCommentForm } from "../components/comments/CreateCommentForm.tsx";
import { baseUrl } from "../utils/helpers.ts";

type Props = {
    fetchPosts: (value: Filter, str: string | null) => void;
};

const CommentView: React.FC<Props> = ({ fetchPosts }) => {
    const { loggedIn } = useAuth();
    const [postData, setPostData] = useState<PostType | null>(null);
    const [commentData, setCommentData] = useState<CommentType[] | null>(null);
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const { post_id } = useParams();
    const { updateTrigger } = useFilter();

    const fetchPostByID = async () => {
        try {
            const response = await fetch(`${baseUrl}/posts/${post_id}`, {
                method: "GET",
                credentials: "include",
            });
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }
            const data = await response.json();
            setPostData(data);
            setCommentData(data.comments)
        } catch (error) {
            console.error("Error:", error);
        } finally {
            setIsLoading(false);
        }
    };

    const orderComments = () => {
        if (!commentData) {
            return []
        }
        const sortedCommentData = [...commentData].sort((a, b) => {
            const aPathParts = a.path.split('/').map(Number);
            const bPathParts = b.path.split('/').map(Number);

            for (let i = 0; i < Math.min(aPathParts.length, bPathParts.length); i++) {
                if (aPathParts[i] !== bPathParts[i]) {
                    return aPathParts[i] - bPathParts[i];
                }
            }
            return aPathParts.length - bPathParts.length;
        })
        setCommentData(sortedCommentData)
    }


    useEffect(() => {
        fetchPostByID();
        orderComments();
    }, [updateTrigger]);
    return (
        <div className="min-h-screen h-full bg-gray-200 flex flex-col">
            {postData ? (
                <>
                    <NavBar fetchPosts={fetchPosts} />
                    {isLoading ? (
                        <h1>Loading...</h1>
                    ) : (
                        <div className="flex flex-col items-center mt-3">
                            <PostComp index={null} post={postData} />

                            <div className="bg-white rounded-xl p-2 h-full w-11/12 max-w-lg flex flex-col items-center mt-3">
                                {loggedIn && (
                                    <CreateCommentForm
                                        fetchPosts={fetchPostByID}
                                        postId={postData.id}
                                    />
                                )}

                                {commentData ? (
                                    <CommentList
                                        fetchPosts={fetchPostByID}
                                        comments={commentData}
                                    />
                                ) : (
                                    <h1 className="text-xl font-bold tracking-wide">
                                        {loggedIn
                                            ? "Be the first to comment"
                                            : "Please login to comment"}
                                    </h1>
                                )}
                            </div>
                        </div>
                    )}
                </>
            ) : (
                <h1>Post Not Found</h1>
            )}
        </div>
    );
};
export default CommentView;
