import { useState } from "react";
import LoginForm from "../components/user/LoginForm";
import { LoginWindowState } from "../utils/type";
import SignUpForm from "../components/user/SignUpForm";
import { ResetPasswordForm } from "../components/user/ResetPasswordForm";

export const Login = () => {
    const [formWindow, setFormWindow] = useState<LoginWindowState>("signIn");

    const handleClick = (str: LoginWindowState) => {
        setFormWindow(str);
    };

    const renderFormContent = () => {
        switch (formWindow) {
            case "signIn":
                return <LoginForm fn={handleClick} />;
            case "signUp":
                return <SignUpForm />;
            case "forgotPassword":
                return <ResetPasswordForm fn={handleClick} />;
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
