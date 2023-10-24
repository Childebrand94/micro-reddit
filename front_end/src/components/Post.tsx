import { Post as PostType, User as UserType } from "../utils/type";
import { shortenUrl, getTimeDif } from "../utils/helpers.ts";
import { useEffect, useState } from "react";

type PostProps = {
  post: PostType;
};

export const Post: React.FC<PostProps> = ({ post }) => {
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
    <div className="my-2 flex flex-col">
      <div className="flex">
        <a className="w-9/12 font-bold transition duration-200 cursor-pointer">
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
  );
};

export default Post;
