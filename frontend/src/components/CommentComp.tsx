import { getTimeDif } from "../utils/helpers.ts";
import { Comment as CommentType } from "../utils/type.ts";
import { User } from "./User.tsx";
import { Arrows } from "./Arrows.tsx";

export type CommentProp = {
    comment: CommentType;
};

const CommentComp: React.FC<CommentProp> = ({ comment }) => {
    console.log(comment.upVotes);
    return (
        <div className="grid grid-cols-[1fr,10fr]">
            <Arrows id={-1} type="comments" />
            <div className="col-start-2">
                <span className="text-gray-300 text-xs">
                    <h1 className="text-blue-300 text-xs inline">
                        <User
                            username={comment.author.userName}
                            id={String(comment.authorId)}
                        />
                    </h1>{" "}
                    <span>
                        {comment.upVotes} points {getTimeDif(comment.createdAt)}{" "}
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
export default CommentComp;
