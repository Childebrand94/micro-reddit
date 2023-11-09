import { ChangeEvent, useState } from "react";
import { BiLogoReddit } from "react-icons/bi";
import { useFilter } from "../../context/UseFilter";
type FormData = {
    title: string;
    url: string;
};

type Props = {
    toggleExpansion: () => void;
};
export const CreatePostComp: React.FC<Props> = ({ toggleExpansion }) => {
    const [formData, setFormData] = useState<FormData>({
        title: "",
        url: "",
    });

    const { setUpdateTrigger } = useFilter();

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
            setUpdateTrigger((prev) => (prev += 1));
            toggleExpansion();
        } catch (error) {
            console.log("There was an error submitting the form", error);
        }
    };

    return (
        <div className="flex">
            <div>
                <BiLogoReddit size={35} /> 
            </div>
            <form className="flex flex-col w-full" onSubmit={handleSubmit}>
                <label htmlFor="title"></label>
                <input
                    required
                    type="text"
                    id="title"
                    name="title"
                    value={formData.title}
                    onChange={handleChange}
                    placeholder="Title"
                    className="bg-gray-100 w-11/12 my-1 ml-2 px-2 font-normal text-black max-w-lg"
                ></input>
                <label htmlFor="url"></label>
                <input
                    required
                    type="url"
                    id="url"
                    name="url"
                    value={formData.url}
                    onChange={handleChange}
                    placeholder="URL"
                    className="bg-gray-100 w-11/12 ml-2 my-1 px-2 font-normal text-black max-w-lg"
                ></input>
                    <button
                        type="submit"
                        className="bg-blue-300 hover:bg-blue-400 rounded-lg w-16 my-1 ml-2 "
                    >
                        Create
                    </button>
            </form>
        </div>
    );
};
