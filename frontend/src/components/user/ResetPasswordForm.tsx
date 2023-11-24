import React, { useState } from "react";
import { LoginWindowState } from "../../utils/type";
import { baseUrl } from "../../utils/helpers";

type Props = {
    fn: (arg: LoginWindowState) => void;
};

export const ResetPasswordForm: React.FC<Props> = ({ fn }) => {
    const [formData, setFormData] = useState<Record<string, string>>({
        email: "",
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
        try {
            const response = await fetch(`${baseUrl}/users`, {
                method: "PUT",
                headers: {
                    "Content-Type": "applications/json",
                },
                body: JSON.stringify(formData),
            });

            if (!response.ok) {
                throw new Error("Failed to update user password");
            }
            const data = await response.json();
            console.log(data);
        } catch (error) {
            console.log("Error:", error);
        } finally {
            fn("signIn");
        }
    };

    return (
        <div className="text-center flex flex-col">
            <h1 className="text-3xl font-bold tracking-wide mb-5 text-blue-500">
                Reset Password
            </h1>
            <form onSubmit={handleSubmit}>
                <div className="form-group mb-4 ">
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
                    <label
                        htmlFor="retypePassword"
                        className="block text-blue-500"
                    ></label>
                    <input
                        type="password"
                        placeholder="Confirm Password"
                        id="retypePassword"
                        name="retypePassword"
                        value={formData.retypepassword}
                        onChange={handleChange}
                        required
                        className="mt-1 p-2 w-full border rounded-md"
                    />
                </div>

                <button
                    type="submit"
                    className="bg-blue-500 text-white py-2 px-5 rounded-md hover:bg-blue-700 transition w-full"
                >
                    Reset
                </button>
            </form>
            <div className="mt-3">
                <p>New to reddit?</p>
                <button className=" underline" onClick={() => fn("signUp")}>
                    Create your account here
                </button>
            </div>
        </div>
    );
};
