import React, { useState } from "react";

const LoginForm: React.FC = () => {
    const [formData, setFormData] = useState({
        email: "",
        password: "",
    });

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        setFormData((prevState) => ({
            ...prevState,
            [name]: value,
        }));
    };

    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        console.log(formData);
    };

    return (
        <div className="text-center flex flex-col bg-white p-7 rounded-md ">
            <h1 className="text-3xl font-bold tracking-wide mb-5 text-blue-500">
                Login
            </h1>
            <form onSubmit={handleSubmit}>
                <div className="form-group mb-4">
                    <label
                        htmlFor="email"
                        className="block text-blue-500"
                    ></label>
                    <input
                        type="email"
                        placeholder="email"
                        id="email"
                        name="email"
                        value={formData.email}
                        onChange={handleChange}
                        required
                        className="mt-1 p-2 w-full border rounded-md"
                    />
                </div>

                <div className="form-group mb-4">
                    <label
                        htmlFor="password"
                        className="block text-blue-500"
                    ></label>
                    <input
                        type="password"
                        placeholder="password"
                        id="password"
                        name="password"
                        value={formData.password}
                        onChange={handleChange}
                        required
                        className="mt-1 p-2 w-full border rounded-md"
                    />
                </div>

                <button
                    type="submit"
                    className="bg-blue-500 text-white py-2 px-5 rounded-md hover:bg-blue-700 transition"
                >
                    Login
                </button>
            </form>
        </div>
    );
};

export default LoginForm;
