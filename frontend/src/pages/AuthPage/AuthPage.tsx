import React, { useState } from 'react';

import Register from '../../components/Register';
import Login from '../../components/Login';

const AuthPage: React.FC = () => {
  const [isRegisterPage, setIsRegisterPage] = useState<boolean>(false);

  return (
    <>
      {isRegisterPage ? (
        <Register toLogin={() => setIsRegisterPage(false)} />
      ) : (
        <Login toRegister={() => setIsRegisterPage(true)} />
      )}
    </>
  );
};

export default AuthPage;
