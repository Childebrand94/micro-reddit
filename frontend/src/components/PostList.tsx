import { useState, useEffect } from "react";
import { Post as PostType } from "../utils/type";
import PostComp from "./PostComp";

const PostList = () => {
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
            setPosts([...data]);
        } catch (error) {
            console.error("Error:", error);
        }
    };
    useEffect(() => {
        console.log("Fetching Posts...");
      
        fetchPosts()
    }, []);

      console.log(posts)
    return (
        <div>
            {posts.map((post, i) => {
                return <PostComp index={i + 1} post={post} key={post.id} />;
            })}
        </div>
    );
};
export default PostList;
