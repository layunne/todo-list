import React from 'react';
import { useAuth } from '../hooks';

import LoginRoutes from './LoginRoutes';
import HomeRoutes from './HomeRoutes';

const Routes: React.FC = () => {
  const { signed } = useAuth();

  return signed ? <HomeRoutes /> : <LoginRoutes />;
};

export default Routes;