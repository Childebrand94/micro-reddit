import { useState } from "react";
import { Comment as CommentType } from "../../utils/type.ts";
import { baseUrl } from "../../utils/helpers.ts";
type props = {
    comment: CommentType;
    fetchPosts: () => void;
    handleClick: () => void;
};

export const CreateChildCommentForm: React.FC<props> = ({
    fetchPosts,
    comment,
    handleClick,
}) => {
    const [formData, setFormData] = useState<string>("");

    const commentData = {
        postId: comment.postId,
        message: formData,
        parentID: comment.id,
        path: comment.path,
    };

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setFormData(event.target.value);
    };
    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        try {
            const response = await fetch(
                `${baseUrl}/posts/${comment.postId}/comments`,
                {
                    method: "Post",
                    credentials: "include",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify(commentData),
                },
            );
            if (!response.ok) {
                throw new Error("Was unable to send comment");
            }
            fetchPosts();
            setFormData("");
            handleClick();
            const data = await response.json();
            console.log(data);
        } catch (error) {
            console.log("There was an error submitting the form", error);
        }
    };
    return (
        <div className="w-11/12 rounded-xl ">
            <form className="flex flex-col" onSubmit={handleSubmit}>
                <label htmlFor="comment"></label>

                <input
                    type="text"
                    id="comment"
                    name="comment"
                    value={formData}
                    onChange={handleChange}
                    placeholder="What are your thoughts?"
                    className=" border-2 p-2 rounded-lg mx-2 mt-2 w-full"
                    required
                ></input>
                <div>
                    <button
                        className="ml-2 mt-2 w-20 bg-blue-400 text-white text-xs py-1 px-2 rounded-sm hover:bg-blue-500 transition"
                        type="submit"
                    >
                        Reply
                    </button>
                    <button
                        onClick={handleClick}
                        className="ml-2 mt-2 w-20 bg-blue-400 text-white text-xs py-1 px-2 rounded-sm hover:bg-blue-500 transition"
                    >
                        Close
                    </button>
                </div>
            </form>
        </div>
    );
};
