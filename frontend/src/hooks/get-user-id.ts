export const getUserId = () => {
  const userData = localStorage.getItem("userData");
  if (userData) {
    try {
      const { id: userId } = JSON.parse(userData);
      return userId;
    } catch (error) {
      console.error("Failed to parse userData:", error);
      return null;
    }
  }
  return null;
};
