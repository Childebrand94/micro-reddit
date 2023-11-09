import { Comment as CommentType } from "../../utils/type";
import CommentComp from "./CommentComp";
type props = {
    comments: CommentType[];
};
export const CommentList: React.FC<props> = ({ comments }) => {
    return (
        <div className="mt-3 w-full flex flex-col justify-center items-center ">
            {comments.map((c) => {
                return (
                    <CommentComp
                        comment={c}
                        key={c.id}
                    />
                );
            })}
        </div>
    );
};
