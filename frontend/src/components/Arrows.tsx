import React from "react";
import { useFilter } from "../context/UseFilter";
import { useNavigate } from "react-router-dom";
import { VoteOptions } from "../utils/type";

type ArrowProps = {
    postId: number;
    commentId: number;
    type: "posts" | "comments";
    usersVote: VoteOptions;
};

export const Arrows: React.FC<ArrowProps> = ({
    postId,
    usersVote,
    commentId,
    type,
}) => {
    const { setUpdateTrigger } = useFilter();
    const navigate = useNavigate();

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
            const resp = await fetch(path, {
                method: "PUT",
            });
            if (!resp.ok) {
                //redirect to sign in
                if (resp.status === 401) {
                    navigate("/users");
                    throw new Error("User must login for this action");
                } else {
                    throw new Error(`Network response was not ok.`);
                }
            }
            setUpdateTrigger((prev: number) => prev + 1);
        } catch (error) {
            console.error("Error during fetch:", error);
        }
    };

    return (
        <div
            className={`flex flex-col ${
                type === "comments"
                    ? "col-start-1 border-l pl-1"
                    : "col-start-2"
            } my-2`}
        >
            <button
                className="my-1 w-6"
                onClick={() =>
                    handleArrowClick(
                        type === "posts" ? postPath.upVote : commentPath.upVote,
                    )
                }
            >
                {usersVote === "upVote" ? (
                    <img
                        className="h-6 hover:scale-110 transition-transform"
                        src="/assets/arrow-up-filled.png"
                        alt="Up Arrow"
                    />
                ) : (
                    <img
                        className="h-6 hover:scale-110 transition-transform"
                        src="/assets/arrow-up.png"
                        alt="Up Arrow"
                    />
                )}
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
                {usersVote === "downVote" ? (
                    <img
                        className="h-6 rotate-180 hover:scale-110 transition-transform"
                        src="/assets/arrow-up-filled.png"
                        alt="Down Arrow"
                    />
                ) : (
                    <img
                        className="h-6 rotate-180 hover:scale-110 transition-transform"
                        src="/assets/arrow-up.png"
                        alt="Down Arrow"
                    />
                )}
            </button>
        </div>
    );
};
