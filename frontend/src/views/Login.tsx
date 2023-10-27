import { useState } from "react";
import { InitialLoginWindow } from "../components/InitialLoginWindow";
import LoginForm from "../components/LoginFrom";
import SignUpForm from "../components/SignUpFrom";

export type LoginWindowState = "initial" | "signUp" | "signIn";

export const Login = () => {
    const [formWindow, setFormWindow] = useState<LoginWindowState>("initial");

    const handleClick = (str: LoginWindowState) => {
        setFormWindow(str);
    };

    const renderFormContent = () => {
        switch (formWindow) {
            case "initial":
                return <InitialLoginWindow fn={handleClick} />;
            case "signIn":
                return <LoginForm />;
            case "signUp":
                return <SignUpForm />;
            default:
                return null;
        }
    };

    return (
        <div className="h-screen bg-gray-200 flex justify-center items-center">
            <div className="text-center flex flex-col bg-white p-7 rounded-md border-2 border-blue-500">
                {renderFormContent()}
            </div>
        </div>
    );
};
