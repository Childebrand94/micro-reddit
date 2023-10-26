import { Post as PostType } from "../utils/type";
import { shortenUrl, getTimeDif } from "../utils/helpers.ts";
import { useState } from "react";
import { Link } from "react-router-dom";
import { User } from "./User.tsx";

export type PostProps = {
    post: PostType;
    index: number | null;
};

export const PostComp: React.FC<PostProps> = ({ post, index }) => {
    const [points, setPoints] = useState(post.upVotes);

    const handleArrowClick = async (path: string) => {
        try {
            const resp = await fetch(path, {
                method: "PUT",
            });
            if (!resp.ok) {
                throw new Error(`HTTP error! Status: ${resp.status}`);
            }

            const domains: string[] = path.split("/");
            if (domains[domains.length - 1] === "up-vote") {
                setPoints((prePoints) => prePoints + 1);
            } else {
                setPoints((prePoints) => prePoints - 1);
            }
        } catch (error) {
            console.error("Error during fetch:", error);
            throw error;
        }
    };

    return (
        <div
            className={`grid ${
                index ? "grid-cols-[0.5fr,1fr,9fr]" : "grid-cols-1"
            } m-2 gap-2`}
        >
            {index && (
                <div className="w-4 col-start-1">
                    <p className="my-2 text-xl">{index}.</p>
                </div>
            )}

            {index && (
                <div className="flex flex-col col-start-2 my-2">
                    <button className="my-1">
                        <img
                            onClick={() =>
                                handleArrowClick(
                                    `/api/posts/${post.id}/up-vote`,
                                )
                            }
                            className="h-6"
                            src="../../public/assets/arrow-up.png"
                            alt="Up Arrow"
                        />
                    </button>
                    <button>
                        <img
                            onClick={() =>
                                handleArrowClick(
                                    `/api/posts/${post.id}/down-vote`,
                                )
                            }
                            className="h-6 rotate-180"
                            src="../../public/assets/arrow-up.png"
                            alt="Down Arrow"
                        />
                    </button>
                </div>
            )}

            <div className="h-12 col-start-3 my-2 flex flex-col">
                <div className="flex">
                    <Link
                        to={`/posts/${post.id}`}
                        className="font-bold transition duration-200 cursor-pointer break-words"
                    >
                        {post.title}

                        <sub className="text-xs text-gray-400 ml-1">
                            ({shortenUrl(post.url)})
                        </sub>
                    </Link>
                </div>
                <div className="flex">
                    <p className="text-xs">
                        {points ? points : 0}{" "}
                        {points === 1 ? "point" : "points"} posted{" "}
                        {getTimeDif(post.createdAt)} ago by{" "}
                        <User
                            username={post.author.userName}
                            id={String(post.authorId)}
                        />
                        <Link
                            to={`/posts/${post.id}`}
                            className="bg-gray-200 text-black whitespace-nowrap text-xxs ml-2 px-1  rounded-lg hover:bg-gray-400 transition duration-200 cursor-pointer"
                        >
                            {post.comments.length} comments
                        </Link>
                    </p>
                </div>
            </div>
        </div>
    );
};

export default PostComp;
