import { useContext } from 'react';
import AuthContext from '../contexts/auth';

function useAuth() {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth set context error');
  }
  return context;
}

export default useAuth;
