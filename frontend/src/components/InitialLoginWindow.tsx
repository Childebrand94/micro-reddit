import React from "react";
import { LoginWindowState } from "../utils/type";

type props = {
    fn: (arg: LoginWindowState) => void;
};

export const InitialLoginWindow: React.FC<props> = ({ fn }) => {
    return (
        <div>
            <h1
                className="text-3xl font-bold tracking-wide mb-5 text-blue-500"
                style={{ fontFamily: "'Trebuchet MS', sans-serif" }}
            >
                Welcome to Reddit!
            </h1>
            <button
                onClick={() => fn("signIn")}
                className="m-3 bg-blue-500 text-white py-2 px-8 rounded-md border-2 border-blue-700 hover:bg-blue-700 transition"
                style={{ fontFamily: "'Verdana', sans-serif" }}
            >
                Sign In
            </button>
            <button
                onClick={() => fn("signUp")}
                className="m-3 bg-gray-300 text-blue-500 py-2 px-8 rounded-md border-2 border-gray-500 hover:bg-gray-400 transition"
                style={{ fontFamily: "'Verdana', sans-serif" }}
            >
                Sign Up
            </button>
            <a href="#" className="text-blue-500 underline text-sm">
                Forgot password?
            </a>
        </div>
    );
};
