import { useState, useEffect } from "react";

// import { useLocation } from "react-router-dom";
import { Post as PostType } from "../utils/type";
import Post from "./Post";

const PostList = () => {
  // const location = useLocation();
  const [posts, setPosts] = useState<PostType[]>([]);

  const fetchPosts = async () => {
    try {
      const response = await fetch("/api/posts", {
        method: "GET",
      });
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      const data = await response.json();
      setPosts([...posts, ...data]);
    } catch (error) {
      console.error("Error:", error);
    }
  };
  useEffect(() => {
    console.log("Fecthing Posts...");
    fetchPosts();
  }, []);

  return (
    <div>
      {posts.map((post, i) => {
        return <Post index={i + 1} post={post} key={post.id} />;
      })}
    </div>
  );
};
export default PostList;
