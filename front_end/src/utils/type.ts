export interface Comment {
  id: number;
  postId: number;
  authorId: number;
  parentId: number;
  message: string;
  upVotes: number;
  createdAt: string;
}

export interface Post {
  id: number;
  authorId: number;
  url: string;
  createdAt: string;
  updatedAt: string;
  upVotes: number;
  comments: Comment[];
}

export interface User {
  id: number;
  firstName: string;
  lastName: string;
  userName: string;
  email: string;
  dateJoined: string;
}
