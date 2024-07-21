import { Link, LinkComponent } from "@tanstack/react-router";
import { CircleUser, Home, Menu, SwordsIcon } from "lucide-react";
import { PropsWithChildren } from "react";
import { Sheet, SheetContent, SheetTrigger } from "./ui/sheet";
import { Button } from "./ui/button";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from "./ui/dropdown-menu";
import { ModeToggle } from "./mode-toggle";

const ShellLink: LinkComponent<'a'> = ({ ...props }) => (
  <Link className="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground [&.active]:bg-muted [&.active]:text-primary transition-all hover:text-primary" {...props} />
);

const MobileShellLink: LinkComponent<'a'> = ({ ...props }) => (
  <Link className="mx-[-0.65rem] flex items-center gap-4 rounded-xl px-3 py-2 text-muted-foreground hover:text-foreground [&.active]:text-foreground [&.active]:bg-muted" {...props} />
);

export const Shell = ({ children }: PropsWithChildren) => {
  return (
    <div className="grid min-h-screen w-full md:grid-cols-[220px_1fr] lg:grid-cols-[280px_1fr]">
      <div className="hidden border-r bg-muted/40 md:block">
        <div className="flex h-14 items-center border-b px-4 lg:h-[60px] lg:px-6">
          <Link href="/" className="flex items-center gap-2 font-semibold">
            <SwordsIcon className="h-6 w-6" />
            <span className="">BattleOrder</span>
          </Link>
        </div>
        <div className="flex-1 py-2 lg:py-4">
          <nav className="grid items-start px-2 text-sm font-medium lg:px-4">
            <ShellLink href="/">
              <Home className="h-4 w-4" />
              Home
            </ShellLink>
          </nav>
        </div>
        <div className="mt-auto p-4">
        </div>
      </div>
      <div className="flex flex-col">
        <header className="flex h-14 items-center gap-4 border-b bg-muted/40 px-4 lg:h-[60px] lg:px-6">
          <Sheet>
            <SheetTrigger asChild>
              <Button
                variant="outline"
                size="icon"
                className="shrink-0 md:hidden"
              >
                <Menu className="h-5 w-5" />
                <span className="sr-only">Toggle navigation menu</span>
              </Button>
            </SheetTrigger>
            <SheetContent side="left" className="flex flex-col">
              <nav className="grid gap-2 text-lg font-medium">
                <Link href="/" className="flex items-center gap-2 text-lg font-semibold">
                  <SwordsIcon className="h-6 w-6" />
                  <span className="sr-only">BattleOrder</span>
                </Link>
                <MobileShellLink>
                  <Home className="h-5 w-5" />
                  Dashboard
                </MobileShellLink>
              </nav>
            </SheetContent>
          </Sheet>
          <div className="w-full flex-1">
          </div>
          <ModeToggle />
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="secondary" size="icon" className="rounded-full">
                <CircleUser className="h-5 w-5" />
                <span className="sr-only">Toggle user menu</span>
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
              <DropdownMenuLabel>My Account</DropdownMenuLabel>
              <DropdownMenuSeparator />
              <DropdownMenuItem>Settings</DropdownMenuItem>
              <DropdownMenuItem>Profile</DropdownMenuItem>
              <DropdownMenuItem>Support</DropdownMenuItem>
              <DropdownMenuSeparator />
              <DropdownMenuItem>Logout</DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </header>
        <main className="flex flex-1 flex-col gap-4 p-4 lg:gap-6 lg:p-6">
          {children}
        </main>
      </div>
    </div>
  );
}
