import { useRouter } from "next/navigation";

const useLogout = () => {
  const router = useRouter();

  const logout = () => {
    // Remove token and user data from localStorage
    localStorage.removeItem("authToken");
    localStorage.removeItem("userData");

    router.push("/login");
  };

  return { logout };
};

export default useLogout;

// usage

// const Profile = () => {
//     const { logout } = useLogout();
  
//     return (
//       <div>
//         <h1>Profile</h1>
//         <button onClick={logout} className="btn-logout">
//           Log Out
//         </button>
//       </div>
//     );
//   };