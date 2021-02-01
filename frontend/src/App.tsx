import React from 'react';
import Routes from './routes';
import { AuthProvider } from './contexts/auth';
import CssBaseline from '@material-ui/core/CssBaseline';
import { ThemeProvider } from '@material-ui/styles';

import { createMuiTheme } from '@material-ui/core';

const theme = createMuiTheme({
  palette: {
    type: 'dark',
  },
});
function App() {
  return (
    <ThemeProvider theme={theme}>
      <AuthProvider>
        <CssBaseline>
          <Routes />
        </CssBaseline>
      </AuthProvider>
    </ThemeProvider>
  );
}

export default App;
