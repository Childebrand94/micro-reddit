import { Comment as CommentType } from "../../utils/type";
import CommentComp from "./CommentComp";
type props = {
    comments: CommentType[];
    fetchPosts: () => void;
};
export const CommentList: React.FC<props> = ({ comments, fetchPosts }) => {
    return (
        <div className="mt-3 w-full flex flex-col felx-start ">
            {comments.map((comment) => {
                return (
                    <CommentComp
                        fetchPosts={fetchPosts}
                        comment={comment}
                        key={comment.id}
                    />
                );
            })}
        </div>
    );
};
