import { HashRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./views/Home";
import CommentView from "./views/CommentView";
import Profile from "./views/Profile";
import { Login } from "./views/Login";
import { AuthProvider } from "./context/AuthProvider";
import { CreatePost } from "./views/CreatePost";
import { useFilter } from "./context/UseFilter";
import { useEffect, useState } from "react";
import { Filter, Post as PostType } from "./utils/type";
import { baseUrl } from './utils/helpers'
import 'react-toastify/dist/ReactToastify.css'
import { ToastContainer } from "react-toastify";

function App() {
    const [posts, setPosts] = useState<PostType[] | null>([]);
    const { updateTrigger, filter } = useFilter();


    const fetchPosts = async (filter: Filter, search: string | null = null) => {
        try {
            let url = `${baseUrl}/posts?sort=${filter}`;
            if (search && search.trim() !== "") {
                url += `&search=${encodeURIComponent(search.trim())}`;
            }
            const response = await fetch(url, {
                method: "GET",
                credentials: "include",
            });
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }
            const data = await response.json();

            if (data === null) {
                setPosts(null);
            } else {
                setPosts(data);
            }
        } catch (error) {
            console.error("Error:", error);
        }
    };


    useEffect(() => {
        fetchPosts(filter);
    }, [updateTrigger]);

    return (
        <>
            <AuthProvider>
                <Router>
                    <Routes>
                        <Route
                            path="/"
                            element={
                                <Home fetchPosts={fetchPosts} posts={posts} />
                            }
                        />
                        <Route
                            path="/posts/:post_id"
                            element={<CommentView fetchPosts={fetchPosts} />}
                        />
                        <Route path="/posts" element={<CreatePost />} />
                        <Route
                            path="/users/:user_id"
                            element={<Profile fetchPosts={fetchPosts} />}
                        />
                        <Route path="/users" element={<Login />} />
                    </Routes>
                </Router>
            </AuthProvider>
        </>
    );
}

export default App;
