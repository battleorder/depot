import { CircleUser } from "lucide-react";
import { Button } from "./ui/button";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from "./ui/dropdown-menu";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { supabase } from "@/lib/client";
import { Link } from "@tanstack/react-router";
import { authStore, updateAuth, useAuth } from "@/lib/auth";
import { useStore } from "@tanstack/react-store";

const AnonMenu = () => (
  <>
    <DropdownMenuItem asChild>
      <Link to="/auth/login">Log in</Link>
    </DropdownMenuItem>
    <DropdownMenuItem>Sign up</DropdownMenuItem>
    <DropdownMenuItem>Support</DropdownMenuItem>
  </>
);

const AuthedMenu = () => {
  const { mutate: logout } = useMutation({
    mutationFn: () => supabase.auth.signOut({ scope: "local" }),
    onSuccess: () => updateAuth(),
  });

  return (
    <>
      <DropdownMenuLabel>My Account</DropdownMenuLabel>
      <DropdownMenuSeparator />
      <DropdownMenuItem>Settings</DropdownMenuItem>
      <DropdownMenuItem>Profile</DropdownMenuItem>
      <DropdownMenuItem>Support</DropdownMenuItem>
      <DropdownMenuSeparator />
      <DropdownMenuItem onClick={() => logout()}>Logout</DropdownMenuItem>
    </>
  );
};

export const AuthMenu = () => {
  const { isAuthenticated } = useStore(authStore);

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="secondary" size="icon" className="rounded-full">
          <CircleUser className="h-5 w-5" />
          <span className="sr-only">Toggle user menu</span>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        {isAuthenticated ? <AuthedMenu /> : <AnonMenu />}
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
