import { Post as PostType } from "../utils/type";
import PostComp from "./PostComp";

type Props = {
    posts: PostType[];
};

const PostList: React.FC<Props> = ({ posts }) => {
    return (
        <div className="sm:px-6">
            {posts.map((post, i) => {
                return <PostComp index={i + 1} post={post} key={post.id} />;
            })}
        </div>
    );
};
export default PostList;
