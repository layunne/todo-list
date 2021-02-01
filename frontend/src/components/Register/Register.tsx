import React, { useState } from 'react';
import Backdrop from '@material-ui/core/Backdrop';
import CircularProgress from '@material-ui/core/CircularProgress';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';

import * as S from './styled';
import { RegisterData } from '../../models/login';
import AuthService from '../../services/authService';

type RegisterProps = {
  toLogin: () => void;
};

const Register: React.FC<RegisterProps> = ({ toLogin }: RegisterProps) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [user, setUser] = useState<RegisterData>({
    username: '',
    password: '',
    name: '',
    confirmPassword: '',
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    setLoading(true);

    try {
      await AuthService.register(user);
      console.log('passou');
      setLoading(false);

      setUser({
        username: '',
        password: '',
        name: '',
        confirmPassword: '',
      });
      toLogin();
    } catch (err) {
      console.log(`error: ${err}`);
      setLoading(false);
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
          <S.ContentTitle>Register</S.ContentTitle>
          <TextField
            label="Name"
            name="name"
            variant="outlined"
            placeholder="Name"
            margin="normal"
            required
            onChange={handleChange}
            value={user.name}
          />
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
          <TextField
            label="Confirm Password"
            name="confirmPassword"
            variant="outlined"
            placeholder="Confirm Password"
            margin="normal"
            required
            type="password"
            onChange={handleChange}
            value={user.confirmPassword}
          />
          <Button variant="contained" color="primary" onClick={handleSubmit}>
            Submit
          </Button>

          <S.RegisterText>Login Now</S.RegisterText>

          <Button variant="contained" color="primary" onClick={toLogin}>
            Login
          </Button>
        </S.LoginHeader>
      </S.LoginContainer>
    </>
  );
};

export default Register;
