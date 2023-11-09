import { useEffect, useState } from "react";
import NavBar from "../components/NavBar";
import { PostComp } from "../components/PostComp";
import { Post as PostType } from "../utils/type";
import { useParams } from "react-router-dom";
import { CommentList } from "../components/CommentList.tsx";
import { useAuth } from "../context/UseAuth.tsx";
import { CreateCommentForm } from "../components/CreateCommentForm.tsx";
import { useFilter } from "../context/UseFilter.tsx";

const CommentView = () => {
    const { loggedIn } = useAuth();
    const [postData, setPostData] = useState<PostType | null>(null);
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const { post_id } = useParams();
    const { updateTrigger } = useFilter()

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
    }, [updateTrigger]);
    return (
        <div className="h-screen bg-gray-200 flex flex-col" >
            {postData ? (
                <>
                    <NavBar />
                    {isLoading ? (
                        <h1>Loading...</h1>
                    ) : (
                        <div className="flex flex-col items-center mt-3"> 
                            <PostComp
                                index={null}
                                post={postData}
                            />

                            <div className="bg-white rounded-xl p-2 h-full w-11/12 flex flex-col items-center mt-3">
                                {loggedIn && (
                                    <CreateCommentForm
                                        fetchPosts={fetchPostByID}
                                        postId={postData.id}
                                    />
                                )}

                                    {postData.comments ? (
                                        <CommentList
                                            comments={postData.comments}
                                        />
                                    ) : (
                                        <h1 className="text-xl font-bold tracking-wide">Be the first to comment</h1>
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
