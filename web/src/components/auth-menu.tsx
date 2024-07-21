import { CircleUser } from "lucide-react";
import { Button } from "./ui/button";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from "./ui/dropdown-menu";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { supabase } from "@/lib/client";
import { Link } from "@tanstack/react-router";

const AnonMenu = () => (
  <>
    <DropdownMenuItem asChild>
      <Link href="/auth/login">Log in</Link>
    </DropdownMenuItem>
    <DropdownMenuItem>Sign up</DropdownMenuItem>
    <DropdownMenuItem>Support</DropdownMenuItem>
  </>
);

const AuthedMenu = () => {
  const client = useQueryClient();
  const { mutate: logout } = useMutation({
    mutationFn: () => supabase.auth.signOut({ scope: "local" }),
    onSuccess: () => client.invalidateQueries({ queryKey: ['auth.user'] }),
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
  const { isLoading, data: user } = useQuery({
    queryKey: ['auth.user'],
    queryFn: async () => {
      const sess = await supabase.auth.getSession();
      if (sess.error) throw sess.error;
      const { data, error } = await supabase.auth.getUser();
      if (error) throw error;
      return data;
    },
  });

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="secondary" size="icon" className="rounded-full">
          <CircleUser className="h-5 w-5" />
          <span className="sr-only">Toggle user menu</span>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        {isLoading && 'Loading...'}
        {!isLoading && (user ? <AuthedMenu /> : <AnonMenu />)}
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
