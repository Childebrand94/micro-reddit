import React, { useState } from "react";

const SignUpForm: React.FC = () => {
    const [formData, setFormData] = useState<Record<string, string>>({
        firstName: "",
        lastName: "",
        email: "",
        username: "",
        password: "",
        retypePassword: "",
    });

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        setFormData((prevState) => ({
            ...prevState,
            [name]: value,
        }));
    };

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        if (formData.password !== formData.retypePassword) {
            alert("Passwords do not match!");
            return;
        }

        const url = "/api/users";
        try {
            const response = await fetch(url, {
                method: "Post",
                headers: {
                    "Content-Type": "applications/json",
                },
                body: JSON.stringify(formData),
            });

            console.log(response.body);
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }

            const data = await response.json();
            console.log(data);
        } catch (error) {
            console.log("There was an error submitting the form", error);
        }
    };

    return (
        <div className="text-center flex flex-col bg-white p-7 rounded-md ">
            <h1 className="text-3xl font-bold tracking-wide mb-5 text-blue-500">
                Sign Up
            </h1>
            <form onSubmit={handleSubmit}>
                <div className="form-group mb-4">
                    <label
                        htmlFor="firstName"
                        className="block text-blue-500"
                    ></label>
                    <input
                        type="text"
                        placeholder="First Name"
                        id="firstName"
                        name="firstName"
                        value={formData.firstName}
                        onChange={handleChange}
                        required
                        className="mt-1 p-2 w-full border rounded-md"
                    />
                </div>

                <div className="form-group mb-4">
                    <label
                        htmlFor="lastName"
                        className="block text-blue-500"
                    ></label>
                    <input
                        type="text"
                        placeholder="Last Name"
                        id="lastName"
                        name="lastName"
                        value={formData.lastName}
                        onChange={handleChange}
                        required
                        className="mt-1 p-2 w-full border rounded-md"
                    />
                </div>

                <div className="form-group mb-4">
                    <label
                        htmlFor="email"
                        className="block text-blue-500"
                    ></label>
                    <input
                        type="email"
                        placeholder="Email"
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
                        htmlFor="username"
                        className="block text-blue-500"
                    ></label>
                    <input
                        type="username"
                        placeholder="Username"
                        id="username"
                        name="username"
                        value={formData.username}
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
                        placeholder="Password"
                        id="password"
                        name="password"
                        value={formData.password}
                        onChange={handleChange}
                        required
                        className="mt-1 p-2 w-full border rounded-md"
                    />
                </div>

                <div className="form-group mb-4">
                    <label
                        htmlFor="retypePassword"
                        className="block text-blue-500"
                    ></label>
                    <input
                        type="password"
                        placeholder="Password"
                        id="retypePassword"
                        name="retypePassword"
                        value={formData.retypePassword}
                        onChange={handleChange}
                        required
                        className="mt-1 p-2 w-full border rounded-md"
                    />
                </div>

                <button
                    type="submit"
                    className="bg-blue-500 text-white py-2 px-5 rounded-md hover:bg-blue-700 transition"
                >
                    Sign Up
                </button>
            </form>
        </div>
    );
};

export default SignUpForm;
