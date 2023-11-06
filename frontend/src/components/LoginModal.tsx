import React from "react";

type Props = {
    onClose: () => void;
};

export const LoginModal: React.FC<Props> = ({ onClose }) => {
    return (
        <div className="fixed inset-0 bg-gray-500 bg-opacity-50 overflow-y-auto h-full w-full flex justify-center items-center">
            <div className="bg-white rounded-lg shadow-xl">
                <div className="max-w-sm m-8 text-center p-6 bg-white rounded-lg border border-gray-200 shadow-sm">
                    <h2 className="text-lg mb-4 font-bold text-gray-900">
                        Please Log In
                    </h2>
                    <p className="mb-4 text-sm text-gray-600">
                        You need to be logged in to perform this action.
                    </p>
                    <button
                        onClick={onClose}
                        className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-300"
                    >
                        Close
                    </button>
                </div>
            </div>
        </div>
    );
};
