import { HashRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./views/Home";
import CommentView from "./views/CommentView";
import Profile from "./views/Profile";
import { Login } from "./views/Login";
import { AuthProvider } from "./context/AuthProvider";
import { CreatePost } from "./views/CreatePost";

function App() {
    return (
        <>
            <AuthProvider>
                <Router>
                    <Routes>
                        <Route path="/" element={<Home />} />
                        <Route
                            path="/posts/:post_id"
                            element={<CommentView />}
                        />
                        <Route path="/posts" element={<CreatePost />} />
                        <Route path="/users/:user_id" element={<Profile />} />
                        <Route path="/users" element={<Login />} />
                    </Routes>
                </Router>
            </AuthProvider>
        </>
    );
}

export default App;
