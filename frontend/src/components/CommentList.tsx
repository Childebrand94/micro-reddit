import { Comment as CommentType } from "../utils/type";
import CommentComp from "./CommentComp";
type props = {
    comments: CommentType[];
    fetchComments: () => void;
};
export const CommentList: React.FC<props> = ({ comments, fetchComments }) => {
    return (
        <div>
            {comments.map((c) => {
                return (
                    <CommentComp
                        comment={c}
                        key={c.id}
                        fetchComments={fetchComments}
                    />
                );
            })}
        </div>
    );
};
