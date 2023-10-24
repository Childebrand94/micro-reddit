import { useState, useEffect } from "react";

type FetchResult<T> = {
  data: T | null;
  loading: boolean;
  error: Error | null | unknown;
};

export const useFetch = <T>(url: string): FetchResult<T> => {
  const [data, setData] = useState<T | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<Error | null | unknown>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(url);
        if (!response.ok) {
          return { data, loading, error };
        }
        const result = await response.json();
        setData(result);
      } catch (error) {
        setError(error);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [url]);

  return { data, loading, error };
};

export default useFetch;
// const [posts, setPosts] = useState<PostType[]>([]);
// const [users, setUsers] = useState<UserType[]>([]);
//
// const fetchPosts = async () => {
//   try {
//     const [postResponse, userResponse] = await Promise.all([
//       fetch("/api/posts"),
//       fetch("/api/users"),
//     ]);
//
//     if (!postResponse.ok || !userResponse.ok) {
//       throw new Error("Network response was not ok");
//     }
//     const postsData = await postResponse.json();
//     const usersData = await userResponse.json();
//
//     setPosts([...posts, ...postsData]);
//     setUsers([...users, ...usersData]);
//   } catch (error) {
//     console.error("Error:", error);
//   }
// };
// useEffect(() => {
//   console.log("Fecthing Posts...");
//   fetchPosts();
// }, []);
