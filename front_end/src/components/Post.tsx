import { Post as PostType, User as UserType } from "../utils/type";
import { shortenUrl, getTimeDif } from "../utils/helpers.ts";
import { useEffect, useState } from "react";

type PostProps = {
  post: PostType;
  index: number;
};

export const Post: React.FC<PostProps> = ({ post, index }) => {
  const [user, setUser] = useState<UserType | null>(null);

  const fetchUserByID = async () => {
    try {
      const response = await fetch(`/api/users/${post.authorId}`, {
        method: "GET",
      });
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      const data = await response.json();
      setUser(data);
    } catch (error) {
      console.error("Error:", error);
    }
  };

  useEffect(() => {
    fetchUserByID();
  }, []);

  return (
    <div className="grid grid-cols-[0.5fr,1fr,9fr] m-2 gap-2">
      <div className="w-4">
        <p className="my-2 text-xl">{index}.</p>
      </div>
      <div className="flex flex-col col-start-2 my-2">
        <button className="my-1">
          <img
            className="h-6"
            src="../../public/assets/arrow-up.png"
            alt="Up Arrow"
          />
        </button>
        <button>
          <img
            className="h-6"
            src="../../public/assets/arrow-down.png"
            alt="Down Arrow"
          />
        </button>
      </div>
      <div className="h-12 col-start-3 my-2 flex flex-col">
        <div className="flex">
          <a className=" font-bold transition duration-200 cursor-pointer">
            {post.title}

            <sub className="text-xs text-gray-400 ml-1">
              ({shortenUrl(post.url)})
            </sub>
          </a>
        </div>
        <div className="flex">
          <p className="text-xs">
            posted {getTimeDif(post.createdAt)} ago by {user?.firstName}
            <a className="bg-gray-200 text-black ml-2 px-1  rounded-lg hover:bg-gray-400 transition duration-200 cursor-pointer">
              comments
            </a>
          </p>
        </div>
      </div>
    </div>
  );
};

export default Post;
