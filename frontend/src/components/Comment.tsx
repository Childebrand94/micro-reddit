import { useState } from "react";
import { getTimeDif } from "../utils/helpers.ts";
import { Comment as CommentType } from "../utils/type.ts";

export type CommentProp = {
    comment: CommentType;
    index: number | null;
};

const Comment: React.FC<CommentProp> = ({ comment }) => {
    console.log(comment);
    const [points, setPoints] = useState(comment.upVotes);

    const handleArrowClick = async (path: string) => {
        try {
            const resp = await fetch(path, {
                method: "PUT",
            });
            if (!resp.ok) {
                throw new Error(`HTTP error! Status: ${resp.status}`);
            }

            setPoints((prePoints) => prePoints + 1);
        } catch (error) {
            console.error("Error during fetch:", error);
            throw error;
        }
    };

    return (
        <div className="grid grid-cols-[1fr,10fr]">
            <div className="flex flex-col col-start-1 my-2">
                <button className="my-1 ml-2">
                    <img
                        onClick={() =>
                            handleArrowClick(`/api/posts/${comment.id}/up-vote`)
                        }
                        className="h-4"
                        src="../../public/assets/arrow-up.png"
                        alt="Up Arrow"
                    />
                </button>
                <button className="ml-2">
                    <img
                        onClick={() =>
                            handleArrowClick(
                                `/api/posts/${comment.id}/down-vote`,
                            )
                        }
                        className="h-4 rotate-180"
                        src="../../public/assets/arrow-up.png"
                        alt="Down Arrow"
                    />
                </button>
            </div>
            <div className="col-start-2">
                <span className="text-gray-300 text-xs">
                    <h1 className="text-blue-300 text-xs inline">
                        {comment.author.userName}
                    </h1>{" "}
                    <span>
                        {points} points {getTimeDif(comment.createdAt)}{" "}
                        ago
                    </span>
                </span>
                <p>
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                    Vivamus vel mauris vitae leo consequat ullamcorper. Fusce
                    bibendum, ante ac porttitor dictum, velit odio egestas orci,
                    in facilisis neque tortor eu nibh.
                </p>
            </div>
        </div>
    );
};
export default Comment;
