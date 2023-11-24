import { useState } from "react";
import { baseUrl } from "../../utils/helpers";

type props = {
    postId: number;
    fetchPosts: () => void;
};

export const CreateCommentForm: React.FC<props> = ({ fetchPosts, postId }) => {
    const [formData, setFormData] = useState<string>("");

    const commentData = {
        postId: postId,
        parentId: null,
        message: formData,
    };

    const handleChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        setFormData(event.target.value);
    };
    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        try {
            const response = await fetch(`${baseUrl}/posts/${postId}/comments`, {
                method: "Post",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(commentData),
            });
            if (!response.ok) {
                throw new Error("Was unable to send comment");
            }
            fetchPosts();
            setFormData("");

            const data = await response.json();
            console.log(data);
        } catch (error) {
            console.log("There was an error submitting the form", error);
        }
    };
    return (
        <div className="w-11/12 rounded-xl ">
            <form className="flex flex-col" onSubmit={handleSubmit}>
                <label htmlFor="text"></label>

                <textarea
                    id="comment"
                    name="comment"
                    rows={4}
                    value={formData}
                    onChange={handleChange}
                    placeholder="What are your thoughts?"
                    className=" border-2 p-2 mx-2 mt-2"
                    required
                ></textarea>
                <button
                    className="ml-2 mt-2 w-20 bg-blue-400 text-white text-xs py-1 px-2 rounded-sm hover:bg-blue-500 transition"
                    type="submit"
                >
                    Comment
                </button>
            </form>
        </div>
    );
};
