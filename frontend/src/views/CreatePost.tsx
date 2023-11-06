import { useState, ChangeEvent } from "react";
import { useNavigate } from "react-router-dom";

type FormData = {
    title: string;
    url: string;
};

export const CreatePost = () => {
    const navigate = useNavigate();
    const [formData, setFormData] = useState<FormData>({
        title: "",
        url: "",
    });

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData((prev) => ({
            ...prev,
            [name]: value,
        }));
    };

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        const url = "/api/posts";
        try {
            const response = await fetch(url, {
                method: "Post",
                headers: {
                    "Content-Type": "applications/json",
                },
                body: JSON.stringify(formData),
            });

            if (!response.ok) {
                throw new Error("Network response was not ok");
            }

            const data = await response.json();
            console.log(data.message);
            //change redirect to post that was just made
            navigate("/");
        } catch (error) {
            console.log("There was an error submitting the form", error);
        }
    };

    return (
        <div className="text-center flex flex-col bg-white p-7 rounded-md">
            <h1 className="text-3xl font-bold tracking-wide mb-5 text-blue-500">
                Create Post
            </h1>
            <form onSubmit={handleSubmit}>
                <div className="form-group mb-4">
                    <label htmlFor="title" className="block text-blue-500">
                        Title
                    </label>
                    <input
                        type="text"
                        placeholder="Enter title"
                        id="title"
                        name="title"
                        value={formData.title}
                        onChange={handleChange}
                        required
                        className="mt-1 p-2 w-full border rounded-md"
                    />
                </div>

                <div className="form-group mb-4">
                    <label htmlFor="url" className="block text-blue-500">
                        URL
                    </label>
                    <input
                        type="url"
                        placeholder="Enter the URL"
                        pattern="^(https?://)?([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}(/.*)?$"
                        title="www.example.com"
                        id="url"
                        name="url"
                        value={formData.url}
                        onChange={handleChange}
                        required
                        className="mt-1 p-2 w-full border rounded-md"
                    />
                </div>

                <button
                    type="submit"
                    className="bg-blue-500 text-white py-2 px-5 rounded-md hover:bg-blue-700 transition"
                >
                    Post
                </button>
            </form>
        </div>
    );
};
