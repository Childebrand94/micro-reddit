export const baseUrl = import.meta.env.VITE_API_URL

export const shortenUrl: (url: string) => string = (url) => {
    try {
        const parsedUrl = new URL(url);
        const domainParts = parsedUrl.hostname.split(".");

        if (domainParts.length < 2) {
            return "Invalid URL";
        }

        const lastIndex = domainParts.length - 1;

        return `${domainParts[lastIndex - 1]}.${domainParts[lastIndex]}`;
    } catch (e) {
        return "Invalid URL";
    }
};

export const getTimeDif: (time: string) => string = (time) => {
    const currentDate = new Date();
    const dateFromTimestamp: Date = new Date(time);

    const differenceInMilliseconds: number =
        currentDate.getTime() - dateFromTimestamp.getTime();
    const differenceInSeconds: number = Math.floor(
        differenceInMilliseconds / 1000,
    );
    const differenceInMinutes: number = Math.floor(differenceInSeconds / 60);
    const differenceInHours: number = Math.floor(differenceInMinutes / 60);
    const differenceInDays: number = Math.floor(differenceInHours / 24);

    if (differenceInDays >= 1) {
        return `${differenceInDays} days ago`;
    } else if (differenceInHours >= 1) {
        return `${differenceInHours} hours ago`;
    } else if (differenceInMinutes >= 1) {
        return `${differenceInMinutes} minutes ago`;
    } else if (differenceInSeconds >= 1) {
        return `${differenceInSeconds} seconds ago`;
    } else {
        return `${differenceInMilliseconds} milliseconds ago`;
    }
};

export const debounce = <F extends (...args: any[]) => any>(
    func: F,
    waitFor: number,
): ((...args: Parameters<F>) => void) => {
    let timeout: ReturnType<typeof setTimeout>;

    return function (...args: Parameters<F>) {
        clearTimeout(timeout);
        timeout = setTimeout(() => func(...args), waitFor);
    } as (...args: Parameters<F>) => void;
};
