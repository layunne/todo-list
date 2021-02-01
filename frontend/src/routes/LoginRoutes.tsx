import React from 'react';
import { BrowserRouter, Route } from 'react-router-dom';

import AuthPage from '../pages/AuthPage';

const LoginRoutes: React.FC = () => {
  return (
    <BrowserRouter>
      <Route path="/" component={AuthPage} />
    </BrowserRouter>
  );
};

export default LoginRoutes;
