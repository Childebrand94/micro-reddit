import { Post as PostType } from "../utils/type";

type PostProps = {
  post: PostType;
};

const Post: React.FC<PostProps> = ({ post }) => {
  return (
    <div>
      <h1>{post.url}</h1>
      <h3>{post.upVotes}</h3>
      <h3>{post.createdAt}</h3>
      <h3>{post.authorId}</h3>
    </div>
  );
};

export default Post;
