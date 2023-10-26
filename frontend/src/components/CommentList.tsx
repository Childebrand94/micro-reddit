import CommentComp from "./CommentComp";

export const CommentList = () => {
    return (
        <div>
            {commentData.map((c) => {
                return <CommentComp comment={c} key={c.id} index={null} />;
            })}
        </div>
    );
};
