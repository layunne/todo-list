import React, { useState } from 'react';
import { useAuth } from '../../hooks';
import Backdrop from '@material-ui/core/Backdrop';
import CircularProgress from '@material-ui/core/CircularProgress';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';

import * as S from './styled';
import { LoginData } from '../../models/login';

type LoginProps = {
  toRegister: () => void;
};
const Login: React.FC<LoginProps> = ({ toRegister }: LoginProps) => {
  const { Login } = useAuth();
  const [loading, setLoading] = useState<boolean>(false);
  const [user, setUser] = useState<LoginData>({
    username: '',
    password: '',
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (user.username && user.password) {
      setLoading(true);

      try {
        await Login(user);
        console.log('passou');
        setLoading(false);
      } catch (err) {
        console.log(`error: ${err}`);

        setLoading(false);
      }
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.persist();
    setUser((state) => ({
      ...state,
      [e.target.name]: e.target.value,
    }));
  };

  return (
    <>
      <Backdrop open={loading} style={{ color: '#fff', zIndex: 1200 }}>
        <CircularProgress color="inherit" />
      </Backdrop>
      <S.LoginContainer>
        <S.LoginHeader>
          <S.ContentTitle>Sign in</S.ContentTitle>
          <TextField
            label="Username"
            name="username"
            variant="outlined"
            placeholder="Username"
            margin="normal"
            required
            onChange={handleChange}
            value={user.username}
          />
          <TextField
            label="Password"
            name="password"
            variant="outlined"
            placeholder="Password"
            margin="normal"
            required
            type="password"
            onChange={handleChange}
            value={user.password}
          />
          <Button variant="contained" color="primary" onClick={handleSubmit}>
            Login
          </Button>

          <S.RegisterText>Register Now</S.RegisterText>
          <Button variant="contained" color="primary" onClick={toRegister}>
            Register
          </Button>
        </S.LoginHeader>
      </S.LoginContainer>
    </>
  );
};

export default Login;
