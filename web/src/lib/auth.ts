import { Store, useStore } from "@tanstack/react-store";
import { AuthError, AuthSessionMissingError, User } from "@supabase/supabase-js";
import { supabase } from "./client";

type AuthStore = {
  lastChecked: Date | null;
  isAuthenticated: boolean;
  user: User | null;
};

export const authStore = new Store<AuthStore>({
  lastChecked: null,
  isAuthenticated: false,
  user: null,
});

export const useAuth = () => {
  return useStore(authStore);
};

export const updateAuth = async () => {
  const now = new Date();
  try {
    const { data: { user }, error } = await supabase.auth.getUser();
    if (error) throw error;
    authStore.setState(old => ({ ...old, lastChecked: now, user: user, isAuthenticated: true }));
  } catch (error) {
    if (error instanceof AuthError) {
      if (error.status == 401) {
        authStore.setState(old => ({ ...old, user: null, lastChecked: now, isAuthenticated: false }));
      } else if (error instanceof AuthSessionMissingError) {
        authStore.setState(old => ({ ...old, user: null, lastChecked: now, isAuthenticated: false }));
      } else {
        console.error('Potential real auth failure, not updating state', error);
        authStore.setState(old => ({ ...old, lastChecked: now }));
      }
    } else {
      console.error('Potential real auth failure, not updating state', error);
      authStore.setState(old => ({ ...old, lastChecked: now }));
    }
  }
}
