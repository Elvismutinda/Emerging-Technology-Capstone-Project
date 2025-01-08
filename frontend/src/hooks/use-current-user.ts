import { getCookie } from "cookies-next";
import { useEffect, useState } from "react";

const useCurrentUser = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [user, setUser] = useState<any>(null);
  const [token, setToken] = useState<string | null>(null);

  useEffect(() => {
    const token = getCookie("authToken");
    const userData = getCookie("userData");

    if (token && userData) {
      setIsAuthenticated(true);
      setUser(JSON.parse(userData)); // Parse user data from string
      setToken(token as string);
    } else {
      setIsAuthenticated(false);
    }
  }, []);

  return { isAuthenticated, user, token };
};

export default useCurrentUser;

// const useCurrentUser = () => {
//   const token = getCookie("authToken") as string | null;
//   const userData = getCookie("userData");

//   const isAuthenticated = Boolean(token && userData);
//   const user = userData ? JSON.parse(userData as string) : null;

//   return { isAuthenticated, user, token };
// };
