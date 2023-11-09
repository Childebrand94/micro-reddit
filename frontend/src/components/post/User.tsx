import React from "react";
import { Link } from "react-router-dom";

type Props = {
    username: string;
    id: string;
};

export const User: React.FC<Props> = ({ username, id }) => {
    return (
        <Link
            to={`/users/${id}`}
            className="inline font-semibold text-blue-400 hover:underline hover:text-blue-600"
        >
            {username}
        </Link>
    );
};
