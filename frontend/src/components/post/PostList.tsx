import { Post as PostType } from "../../utils/type";
import PostComp from "./PostComp";
import { BiLogoReddit } from "react-icons/bi";
import { useAuth } from "../../context/UseAuth";
import { useState } from "react";
import { CreatePostComp } from "./CreatePostComp";

type Props = {
    posts: PostType[] | null;
};

const PostList: React.FC<Props> = ({ posts }) => {
    const { loggedIn } = useAuth();

    const [isPostExpanded, setIsPostExpanded] = useState<boolean>(false);

    const handleClick = () => {
        setIsPostExpanded(!isPostExpanded);
    };
    return (
        <div className="sm:px-6 overflow-x-hidden bg-gray-200 flex justify-center  min-h-screen h-full">
            <div className="max-w-lg mx-auto w-full">
                {loggedIn && (
                    <div className="w-full border border-gray-300 p-2 bg-white rounded-xl">
                        {isPostExpanded ? (
                            <CreatePostComp toggleExpansion={handleClick} />
                        ) : (
                            <form className="flex">
                                <BiLogoReddit size={35} />
                                <label htmlFor="createPost"></label>
                                <input
                                    onClick={handleClick}
                                    type="text"
                                    id="createPost"
                                    placeholder="Create a post"
                                    className="bg-gray-100 w-full ml-2 px-2 font-normal text-black"
                                />
                            </form>
                        )}
                    </div>
                )}
                {!posts ? (
                    <div className="text-center mt-4 border border-gray-300 p-2 bg-white w-full h-24 rounded-xl">
                        <h1 className="font-semibold text-xl">
                            No Posts Found
                        </h1>
                    </div>
                ) : (
                    posts.map((post, i) => (
                        <PostComp index={i + 1} post={post} key={post.id} />
                    ))
                )}
            </div>
        </div>
    );
};

export default PostList;
