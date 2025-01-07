import { useEffect, useState } from "react";

const useCurrentUser = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [user, setUser] = useState<any>(null);

  useEffect(() => {
    // Retrieve token and user data from localStorage
    const token = localStorage.getItem("authToken");
    const userData = localStorage.getItem("userData");

    if (token && userData) {
      setIsAuthenticated(true);
      setUser(JSON.parse(userData)); // Parse user data from string
    } else {
      setIsAuthenticated(false);
    }
  }, []);

  return { isAuthenticated, user };
};

export default useCurrentUser;
