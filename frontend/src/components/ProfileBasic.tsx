import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { UserPoints, User } from "../utils/type";
import moment from "moment";
import { baseUrl } from "../utils/helpers";

export const ProfileBasic = () => {
    const [userData, setUserData] = useState<User | null>(null);
    const [pointsData, setPointsData] = useState<UserPoints | null>(null);

    const { user_id } = useParams();

    const fetchPoints = async () => {
        try {
            const [respPoints, respUser] = await Promise.all([
                fetch(`${baseUrl}/users/${user_id}/points`, { method: "GET" }),
                fetch(`${baseUrl}/users/${user_id}`, { method: "GET" }),
            ]);
            if (!respPoints.ok || !respUser.ok) {
                throw new Error("Network response was not ok");
            }

            const dataPoints = await respPoints.json();
            const dataUser = await respUser.json();

            setUserData(dataUser);
            setPointsData(dataPoints);
        } catch (error) {
            console.error("Error:", error);
        }
    };
    useEffect(() => {
        fetchPoints();
    }, []);

    const dateStr = userData?.dateJoined;
    const formattedData = moment(dateStr).format("MMMM D, YYYY");

    return (
        <div className="flex flex-col items-center w-full h-full ">
            <div className="bg-white w-5/6 rounded-lg h-5/6 p-4">
            <div className="flex py-1 px-2">
                <div className="mr-2">Karma:</div>
                <div className="font-bold">
                    {pointsData ? pointsData.karma : "unKnown"}
                </div>
            </div>
            <div className="flex py-1 px-2">
                <div className="mr-2">Member since:</div>
                <div className="font-bold">{formattedData}</div>
            </div>
            <div className="flex my-2 border-b-2 border-black w-56">
                <div className="pl-2 font-bold w-full tracking-wide text-lg ">
                    Stats
                </div>
            </div>
            <div className="flex py-1 px-2">
                <div className="pr-2">Total submissions:</div>
                <div className="font-bold">
                    {pointsData ? pointsData.postCount : "Unknown"}
                </div>
            </div>
            <div className="flex py-1 px-2">
                <div className="mr-2">Sites promoted:</div>
                <div className="font-bold">
                    {pointsData ? pointsData.postUpVotes : "Unknown"}
                </div>
            </div>
            <div className="flex py-1 px-2">
                <div className="mr-2">Sites demoted:</div>
                <div className="font-bold">
                    {pointsData ? pointsData.postDownVotes : "Unknown"}
                </div>
            </div>
            </div>
        </div>
    );
};
