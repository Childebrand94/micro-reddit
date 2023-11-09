import { useState } from "react";
import LoginForm from "../components/user/LoginFrom";
import { LoginWindowState } from "../utils/type";
import SignUpForm from "../components/user/SignUpFrom";

export const Login = () => {
    const [formWindow, setFormWindow] = useState<LoginWindowState>("signIn");

    const handleClick = (str: LoginWindowState) => {
        setFormWindow(str);
    };

    const renderFormContent = () => {
        switch (formWindow) {
            case "signIn":
                return <LoginForm fn={handleClick}/>;
            case "signUp":
                return <SignUpForm />;
            default:
                return null;
        }
    };

    return (
        <div className="h-screen flex justify-center items-center">
            <div className="text-center flex flex-col bg-white p-7 rounded-md ">
                {renderFormContent()}
            </div>
        </div>
    );
};
