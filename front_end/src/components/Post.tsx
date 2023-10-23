import { Post as PostType } from "../utils/type";

type PostProps = {
  post: PostType;
};

const Post: React.FC<PostProps> = ({ post }) => {
  return (
    <div className="my-2">
      <h1>{post.url}</h1>
      <div className="flex">
        <h3>{post.authorId}</h3>
        <h3>{post.upVotes}</h3>
      </div>
    </div>
  );
};

export default Post;
