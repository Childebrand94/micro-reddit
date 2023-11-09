import { useEffect, useState } from "react";
import { Filter, Post as PostType } from "../utils/type";
import { useFilter } from "../context/UseFilter";
import PostList from "../components/post/PostList";
import NavBar from "../components/nav/NavBar";

const Home = () => {
    const [posts, setPosts] = useState<PostType[]>([]);
    const { updateTrigger, filter } = useFilter();

    const fetchPosts = async (filter: Filter, search: string | null = null) => {
        try {
            let url = `/api/posts?sort=${filter}`;
            if (search && search.trim() !== "") {
                url += `&search=${encodeURIComponent(search.trim())}`;
            }
            console.log(url);

            const response = await fetch(url, {
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
            <NavBar fetchPosts={fetchPosts} />
            <PostList posts={posts} />
        </div>
    );
};
export default Home;
