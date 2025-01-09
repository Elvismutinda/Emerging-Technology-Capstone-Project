import { AvatarProps } from "@radix-ui/react-avatar";
import { Avatar, AvatarFallback } from "./ui/avatar";
import { User } from "lucide-react";

interface UserAvatarProps extends AvatarProps {
  user: any;
}

const UserAvatar = ({ user, ...props }: UserAvatarProps) => {
  return (
    <Avatar {...props}>
        <AvatarFallback>
          <span className="sr-only">{user.username}</span>
          <User className="h-6 w-6" />
        </AvatarFallback>
    </Avatar>
  );
};

export default UserAvatar;
