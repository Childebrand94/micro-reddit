import { useState, useEffect } from "react";

// import { useLocation } from "react-router-dom";
import { Post as PostType, User as UserType } from "../utils/type";
import Post from "./Post";

const PostList = () => {
  // const location = useLocation();
  const [posts, setPosts] = useState<PostType[]>([]);
  const [users, setUsers] = useState<UserType[]>([]);

  const fetchPosts = async () => {
    try {
      const [postResponse, userResponse] = await Promise.all([
        fetch("/api/posts"),
        fetch("/api/users"),
      ]);

      if (!postResponse.ok || !userResponse.ok) {
        throw new Error("Network response was not ok");
      }
      const postsData = await postResponse.json();
      const usersData = await userResponse.json();

      setPosts([...posts, ...postsData]);
      setUsers([...users, ...usersData]);
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
      {posts.map((post) => {
        return <Post post={post} key={post.id} />;
      })}
    </div>
  );
};
export default PostList;
