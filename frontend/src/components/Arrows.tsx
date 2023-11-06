import React from "react";

type ArrowProps = {
    postId: number;
    commentId: number;
    type: "posts" | "comments";
};

export const Arrows: React.FC<ArrowProps> = ({ postId, commentId, type }) => {
    // Define the endpoint paths
    const postPath = {
        upVote: `/api/posts/${postId}/up-vote`,
        downVote: `/api/posts/${postId}/down-vote`,
    };

    const commentPath = {
        upVote: `/api/posts/${postId}/comments/${commentId}/up-vote`,
        downVote: `/api/posts/${postId}/comments/${commentId}/down-vote`,
    };

    const handleArrowClick = async (path: string) => {
        try {
            console.log("Sending put request for votes...");
            const resp = await fetch(path, {
                method: "PUT",
            });
            if (!resp.ok) {
                //redirect to sign in 
                throw new Error(`HTTP error! Status: ${resp.status}`);
            }
        } catch (error) {
            console.error("Error during fetch:", error);
        }
    };

    return (
        <div
            className={`flex flex-col ${
                type === "comments" ? "col-start-1" : "col-start-2"
            } my-2`}
        >
            <button
                className="my-1"
                onClick={() =>
                    handleArrowClick(
                        type === "posts" ? postPath.upVote : commentPath.upVote,
                    )
                }
            >
                <img
                    className="h-6 hover:scale-110 transition-transform"
                    src="/assets/arrow-up.png"
                    alt="Up Arrow"
                />
            </button>
            <button
                onClick={() =>
                    handleArrowClick(
                        type === "posts"
                            ? postPath.downVote
                            : commentPath.downVote,
                    )
                }
            >
                <img
                    className="h-6 rotate-180 hover:scale-110 transition-transform"
                    src="/assets/arrow-up.png"
                    alt="Down Arrow"
                />
            </button>
        </div>
    );
};
