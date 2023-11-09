import { Post as PostType } from "../../utils/type";
import PostComp from "./PostComp";
import { BiLogoReddit } from "react-icons/bi";
import { useAuth } from "../../context/UseAuth";
import { useState } from "react";
import { CreatePostComp } from "./CreatePostComp";

type Props = {
    posts: PostType[];
};

const PostList: React.FC<Props> = ({ posts }) => {
    const { loggedIn } = useAuth();

    const [isPostExpanded, setIsPostExpanded] = useState<boolean>(false);

    const handleClick = () => {
        setIsPostExpanded(!isPostExpanded);
    };
    return (
        <div className="sm:px-6 bg-gray-200">
            <div className="max-w-2xl mx-auto flex flex-col justify-center items-center">
            {loggedIn && (
                <div className="sm:w-11/12 border border-gray-300 p-2 bg-white rounded-xl">
                    {isPostExpanded ? (
                        <CreatePostComp toggleExpansion={handleClick} />
                    ) : (
                        <form className="flex ">
                            <BiLogoReddit size={35} />
                            <label htmlFor="createPost"></label>
                            <input
                                onClick={handleClick}
                                type="text"
                                id="createPost"
                                placeholder="Create a post"
                                className="bg-gray-100 w-full ml-2 px-2 font-normal text-black"
                            ></input>
                        </form>
                    )}
                </div>
            )}

            {posts.map((post, i) => {
                return <PostComp index={i + 1} post={post} key={post.id} />;
            })}
        </div>
</div>
    );
};
export default PostList;
