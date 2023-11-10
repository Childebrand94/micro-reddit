import React, { useState } from "react";
import { useAuth } from "../../context/UseAuth";
import { LoginWindowState } from "../../utils/type";

type props = {
    fn: (arg: LoginWindowState) => void;
};

const LoginForm: React.FC<props> = ({ fn }) => {
    const { setLoggedIn } = useAuth();
    const [invalidCredential, setInvalidCredential] = useState(false);

    const [formData, setFormData] = useState<Record<string, string>>({
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

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const url = "/api/users/login";
        try {
            const response = await fetch(url, {
                method: "Post",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(formData),
            });

            if (!response.ok) {
                const data = await response.json();
                setInvalidCredential(true);
                throw new Error(`${data.message}`);
            }
            const data = await response.json();
            console.log(data.message);
            setLoggedIn(true);
            window.location.href = "/";
        } catch (error) {
            console.log("There was an error submitting the form", error);
        }
    };

    return (
        <div className="text-center flex flex-col">
            <h1 className="text-3xl font-bold tracking-wide mb-5 text-blue-500">
                Login
            </h1>
            <form onSubmit={handleSubmit}>
                <div className="form-group mb-4">
                    <label htmlFor="email" className="block text-blue-500">
                        {invalidCredential ? (
                            <p className="text-red-500">
                                Invalid Email or Password
                            </p>
                        ) : (
                            <></>
                        )}{" "}
                    </label>
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

                <button
                    type="submit"
                    className="bg-blue-500 text-white py-2 px-5 rounded-md hover:bg-blue-700 transition w-full"
                >
                    Login
                </button>
            </form>
            <div className="mt-3">
                <button
                    className=" underline"
                    onClick={() => fn("forgotPassword")}
                >
                    Forgot Password
                </button>
                <p>New to reddit?</p>
                <button className=" underline" onClick={() => fn("signUp")}>
                    Create your account here
                </button>
            </div>
        </div>
    );
};

export default LoginForm;
