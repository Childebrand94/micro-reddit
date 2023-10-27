export const ProfileBasic = () => {
    return (
        <div className="flex flex-col w-full">
            <div className="flex py-1">
                <div className="mr-2">Karma:</div>
                <div className="font-bold">948</div>
            </div>
            <div className="flex py-1">
                <div className="mr-2">Member since:</div>
                <div className="font-bold">July 06, 2005</div>
            </div>
            <div className="flex py-2 bg-gray-200">
                <div className="pl-2 font-bold w-full tracking-wide text-lg ">
                    Stats
                </div>
            </div>
            <div className="flex py-1">
                <div className="pr-2">Total submissions:</div>
                <div className="font-bold">327</div>
            </div>
            <div className="flex py-1">
                <div className="mr-2">Sites promoted:</div>
                <div className="font-bold">422</div>
            </div>
            <div className="flex py-1">
                <div className="mr-2">Sites demoted:</div>
                <div className="font-bold">46</div>
            </div>
        </div>
    );
};
