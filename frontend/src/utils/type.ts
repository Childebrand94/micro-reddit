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
    parentId: number;
    message: string;
    upVotes: number;
    createdAt: string;
};
export type Post = {
    id: number;
    authorId: number;
    title: string;
    url: string;
    createdAt: string;
    updatedAt: string;
    upVotes: number;
    comments: Comment[];
    author: Author;
};

export type User = {
    id: number;
    firstName: string;
    lastName: string;
    username: string;
    email: string;
    dateJoined: string;
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
export type VoteContextType = {
    points: number;
    setPoints: React.Dispatch<React.SetStateAction<number>>;
    handleArrowClick: (path: string) => Promise<void>;
};

export type HttpMethod = "GET" | "POST" | "PUT" | "DELETE";