import { getCookie } from "cookies-next";
import { useEffect, useState } from "react";

const useUserId = () => {
  const [userId, setUserId] = useState<string | null>(null);

  useEffect(() => {
    const userData = getCookie("userData");

    if (userData) {
      try {
        const { id } = JSON.parse(userData);
        setUserId(id);
      } catch (error) {
        console.error("Failed to parse userData:", error);
        setUserId(null);
      }
    } else {
      setUserId(null);
    }
  }, []);

  return userId;
};

export default useUserId;
