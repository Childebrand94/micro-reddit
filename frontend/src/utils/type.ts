type Author = {
    firstName: string;
    lastName: string;
    userName: string;
};
export type Comment = {
    author: Author;
    id: number;
    postId: number;
    authorId: number;
    parentID: {
        Int64: number;
        Valid: boolean;
    };
    message: string;
    upVotes: number;
    createdAt: string;
    usersVoteStatus: VoteOptions;
};
export type Post = {
    id: number;
    authorId: number;
    title: string;
    url: string;
    createdAt: string;
    updatedAt: string;
    upVotes: number;
    comments: Comment[] | null;
    author: Author;
    usersVoteStatus: VoteOptions;
};

export type User = {
    id: number;
    firstName: string;
    lastName: string;
    username: string;
    email: string;
    dateJoined: string;
};

export type UserPoints = {
    postCount: number;
    postUpVotes: number;
    postDownVotes: number;
    commentUpVotes: number;
    commentDownVotes: number;
    karma: number;
};

export type UserID = {
    id: number;
    firstName: string;
    lastName: string;
    username: string;
    email: string;
    dateJoined: string;
    posts: Post[];
    comments: Comment[];
};
// export type VoteContextType = {
// points: number;
// setPoints: React.Dispatch<React.SetStateAction<number>>;
// handleArrowClick: (path: string) => Promise<void>;
// };

export type FormDataType = {
    firstname: string;
    lastname: string;
    email: string;
    username: string;
    password: string;
    retypepassword: string;
};
export type LoginWindowState =
    | "initial"
    | "signUp"
    | "signIn"
    | "forgotPassword";

export type HttpMethod = "GET" | "POST" | "PUT" | "DELETE";

export type Filter = "hot" | "top" | "new";

export type VoteOptions = "upVote" | "downVote" | "noVote";
