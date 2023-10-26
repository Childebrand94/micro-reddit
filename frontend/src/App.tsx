import { HashRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./views/Home";
import CommentView from "./views/CommentView";
import Profile from "./views/Profile";

function App() {
    return (
        <>
            <Router>
                <Routes>
                    <Route path="/" element={<Home />} />
                    <Route path="/posts/:post_id" element={<CommentView />} />
                    <Route path="/profile" element={<Profile />} />
                </Routes>
            </Router>
        </>
    );
}

export default App;
