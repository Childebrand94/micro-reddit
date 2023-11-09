import { useNavigate } from "react-router-dom";
import { Post as PostType } from "../utils/type";
import PostComp from "./PostComp";
import { BiLogoReddit } from "react-icons/bi";
import { useAuth } from "../context/UseAuth";
import { useState } from "react";

type Props = {
    posts: PostType[];
};

const PostList: React.FC<Props> = ({ posts }) => {
    const [isPostExpanded, setIsPostExpanded] = useState<Boolean>(false)
    const {loggedIn} = useAuth()

    const handleCreatePost = () => {
        setIsPostExpanded(true)
    }
    return (
        <div className="sm:px-6 bg-gray-200">
            {loggedIn &&
                <div className="sm:w-11/12 border border-gray-300 p-2 bg-white rounded-xl">
                    { isPostExpanded ? (
                    <form className="flex flex-col">
                        <label htmlFor="title"></label>
                        <input onClick={handleCreatePost} required type="text" id="title" placeholder="Title" className="bg-gray-100 w-11/12 my-1 ml-2 px-2 font-normal text-black"></input>
                        <label htmlFor="url"></label>
                        <input onClick={handleCreatePost} required type="url" id="url" placeholder="URL" className="bg-gray-100 w-11/12 ml-2 my-1 px-2 font-normal text-black"></input>
                        <div className="flex justify-end w-full">
                            <button className="bg-blue-300 hover:bg-blue-400 rounded-lg w-16 my-1 mr-2 ">Create</button>
                        </div>
                    </form>
                    ) 
                    :(
                    <form className="flex ">
                        <BiLogoReddit size={35}/>
                        <label htmlFor="createPost"></label>
                        <input onClick={handleCreatePost} type="text" id="createPost" placeholder="Create a post" className="bg-gray-100 w-full ml-2 px-2 font-normal text-black"></input>
                    </form>

                     )

                    }
                </div>
            }

            {posts.map((post, i) => {
                return <PostComp index={i + 1} post={post} key={post.id} />;
            })}
        </div>
    );
};
export default PostList;
