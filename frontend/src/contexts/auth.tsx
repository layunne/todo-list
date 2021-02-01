import React, { createContext, useState, useEffect } from 'react';
import { User } from '../models/user';
import { LoginData } from '../models/login';
import AuthService from '../services/authService';

const USER_KEY = '@todo-list:user';

interface AuthContextData {
  signed: boolean;
  user: User | null;
  loading: boolean;
  Login(login: LoginData): Promise<void>;
  Logout(): void;
}

const AuthContext = createContext<AuthContextData>({} as AuthContextData);

export const AuthProvider: React.FC = ({ children }) => {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const user = localStorage.getItem(USER_KEY);
    if (user) {
      const userParsed = JSON.parse(user);
      AuthService.setAuthHeader(userParsed.token);

      setUser(userParsed);
      console.log(user);
    }
    setLoading(false);
  }, []);

  async function Login({ username, password }: LoginData) {
    const { data } = await AuthService.login({ username, password });

    const user = {
      id: data.id,
      name: data.name,
      username: data.username,
      token: data.token,
    };
    localStorage.setItem(USER_KEY, JSON.stringify(user));
    AuthService.setAuthHeader(user.token);
    setUser(user);
  }

  function Logout() {
    setUser(null);

    localStorage.removeItem(USER_KEY);
  }

  return (
    <AuthContext.Provider value={{ signed: !!user, user, Login, Logout, loading }}>{children}</AuthContext.Provider>
  );
};

export default AuthContext;
