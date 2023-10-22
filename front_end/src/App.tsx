import { HashRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./views/Home";
import Post from "./views/Post";
import Comment from "./views/Comment";
import Profile from "./views/Profile";

function App() {
  return (
    <>
      <Router>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/profile" element={<Profile />} />
          <Route path="/post" element={<Post />} />
          <Route path="/post/comment" element={<Comment />} />
        </Routes>
      </Router>
    </>
  );
}

export default App;
