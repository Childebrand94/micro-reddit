import { useEffect, useState } from "react";
import NavBar from "../components/nav/NavBar.tsx";
import { PostComp } from "../components/post/PostComp.tsx";
import { Post as PostType, Filter } from "../utils/type";
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
        } catch (error) {
            console.error("Error:", error);
        } finally {
            setIsLoading(false);
        }
    };
    useEffect(() => {
        fetchPostByID();
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

                                {postData.comments ? (
                                    <CommentList
                                        fetchPosts={fetchPostByID}
                                        comments={postData.comments}
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
