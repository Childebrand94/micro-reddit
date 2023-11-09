import { getTimeDif } from "../../utils/helpers.ts";
import { Comment as CommentType } from "../../utils/type.ts";
import { User } from "../post/User.tsx";
import { Arrows } from "../Arrows.tsx";

type CommentProp = {
    comment: CommentType;
};

const CommentComp: React.FC<CommentProp> = ({ comment }) => {
    return (
        <div className={`bg-white border border-gray-200 w-11/12 rounded-xl my-1 flex  ${comment.parentID["Valid"]  ? "ml-20" : ""} `}>
            <Arrows
                postId={comment.postId}
                commentId={comment.id}
                type="comments"
            
            />
            <div className={` ml-2 col-start-2`}>
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
                <p>{comment.message}</p>
            </div>
        </div>
    );
};
export default CommentComp;
