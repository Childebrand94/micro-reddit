import { Comment as CommentType } from "../utils/type";
import CommentComp from "./CommentComp";
type props = {
    comments: CommentType[];
};
export const CommentList: React.FC<props> = ({ comments }) => {
    return (
        <div>
            {comments.map((c) => {
                return <CommentComp comment={c} key={c.id} />;
            })}
        </div>
    );
};
