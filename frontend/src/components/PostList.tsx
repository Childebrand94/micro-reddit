import { useEffect, useState } from "react";
import { Filter, Post as PostType } from "../utils/type";
import PostComp from "./PostComp";
import { useFilter } from "../context/UseFilter";

const PostList = () => {
    const [posts, setPosts] = useState<PostType[]>([]);
    const { updateTrigger, filter } = useFilter();

    const fetchPosts = async (filter: Filter) => {
        console.log("Fetching Posts...");
        try {
            const response = await fetch(`/api/posts?sort=${filter}`, {
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
        fetchPosts(filter);
    }, [updateTrigger]);

    return (
        <div>
            {posts.map((post, i) => {
                return <PostComp index={i + 1} post={post} key={post.id} />;
            })}
        </div>
    );
};
export default PostList;
