import { useState } from "react";

type arrowPros = {
    id: string;
    index: number;
};

export const Arrows: React.FC<arrowPros> = ({ id, index }) => {
    const handleArrowClick = async (path: string) => {
        try {
            const resp = await fetch(path, {
                method: "PUT",
            });
            if (!resp.ok) {
                throw new Error(`HTTP error! Status: ${resp.status}`);
            }

            setPoints((prePoints) => prePoints + 1);
        } catch (error) {
            console.error("Error during fetch:", error);
            throw error;
        }
    };

    return (
        <div>
            {index && (
                <div className="flex flex-col col-start-2 my-2">
                    <button className="my-1">
                        <img
                            onClick={() =>
                                handleArrowClick(`/api/posts/${id}/up-vote`)
                            }
                            className="h-6"
                            src="../../public/assets/arrow-up.png"
                            alt="Up Arrow"
                        />
                    </button>
                    <button>
                        <img
                            onClick={() =>
                                handleArrowClick(`/api/posts/${id}/down-vote`)
                            }
                            className="h-6 rotate-180"
                            src="../../public/assets/arrow-up.png"
                            alt="Down Arrow"
                        />
                    </button>
                </div>
            )}
        </div>
    );
};
