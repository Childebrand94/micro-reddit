import { Filter, Post as PostType } from "../utils/type";
import PostList from "../components/post/PostList";
import NavBar from "../components/nav/NavBar";

type Props = {
    posts: PostType[] | null;
    fetchPosts: (value: Filter, str: string | null) => void;
};
const Home: React.FC<Props> = ({ posts, fetchPosts }) => {
    return (
        <div>
            <NavBar fetchPosts={fetchPosts} />
            <PostList posts={posts} />
        </div>
    );
};
export default Home;
