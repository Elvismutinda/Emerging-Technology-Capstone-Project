import React from 'react'
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
  } from './ui/dropdown-menu'
import useLogout from '@/hooks/use-logout'
import UserAvatar from './UserAvatar';

interface UserAccountNavProps {
    user: any | null;
}

export function UserAccountNav({ user }: UserAccountNavProps) {
  const { logout } = useLogout();

  if (!user) {
    return null;
  }
  return (
    <DropdownMenu>
      <DropdownMenuTrigger>
        <UserAvatar
          user={user}
          className="h-8 w-8"
        />
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        <div className="flex items-center justify-start gap-2 p-2">
          <div className="flex flex-col space-y-1 leading-none">
            {user.username && <p className="font-medium">{user.username}</p>}
            {user.email && (
              <p className="w-[200px] truncate text-sm text-muted-foreground">
                {user.email}
              </p>
            )}
          </div>
        </div>
        <DropdownMenuSeparator />

        <DropdownMenuItem
          className="cursor-pointer"
          onSelect={logout}
        >
          Logout
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
