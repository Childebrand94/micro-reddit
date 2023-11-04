type arrowPros = {
    id: number;
    type: "posts" | "comments";
};

export const Arrows: React.FC<arrowPros> = ({ id, type }) => {
    const handleArrowClick = async (path: string) => {
        try {
            const resp = await fetch(path, {
                method: "PUT",
            });
            if (!resp.ok) {
                throw new Error(`HTTP error! Status: ${resp.status}`);
            }
        } catch (error) {
            console.error("Error during fetch:", error);
            throw error;
        }
    };

    return (
        <div
            className={`flex flex-col ${
                type === "comments" ? "col-start-1" : "col-start-2"
            } my-2`}
        >
            <button className="my-1">
                <img
                    onClick={() =>
                        handleArrowClick(`/api/${type}/${id}/up-vote`)
                    }
                    className="h-6 hover:scale-110 transition-transform"
                    src="/assets/arrow-up.png"
                    alt="Up Arrow"
                />
            </button>
            <button>
                <img
                    onClick={() =>
                        handleArrowClick(`/api/${type}/${id}/down-vote`)
                    }
                    className="h-6 rotate-180 hover:scale-110 transition-transform"
                    src="/assets/arrow-up.png"
                    alt="Down Arrow"
                />
            </button>
        </div>
    );
};
