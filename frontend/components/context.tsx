"use client";

import { createContext, use } from "react";
import { isTokenExpired } from "@/lib/token";
import useLocalStorage from "@/hooks/use-localstorage";

type AuthProviderProps = {
  children: React.ReactNode;
};

type AuthProviderState = {
  token: string;
  email: string;
  isAuth: boolean;
  userID: number;
  setAuth: (token: string, email: string, userID: number) => void;
};

const AuthProviderContext = createContext<AuthProviderState>({
  token: "",
  email: "",
  userID: 0,
  isAuth: false,
  setAuth: () => null,
});

export function AuthProvider({ children }: AuthProviderProps) {
  const [email, setEmail] = useLocalStorage("email", "");
  const [token, setToken] = useLocalStorage("token", "");
  const [userID, setUserID] = useLocalStorage("user_id", "");
  const user_id = parseInt(userID);

  const isAuth = !isTokenExpired(token) && user_id !== 0 && email !== "";

  const value = {
    token: token,
    email: email,
    userID: user_id,
    isAuth: isAuth,
    setAuth: (token: string, email: string, userID: number) => {
      setToken(token);
      setEmail(email);
      setUserID(String(userID));
    },
  };

  return (
    <AuthProviderContext.Provider value={value}>
      {children}
    </AuthProviderContext.Provider>
  );
}

export const useAuth = () => {
  const context = use(AuthProviderContext);

  if (context === undefined) {
    throw new Error("useAuth must be use within a AuthProvider");
  }

  return context;
};
