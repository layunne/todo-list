import styled from 'styled-components';
import Typography from '@material-ui/core/Typography';

export const ContentTitle = styled(Typography)`
  && {
    color: #7f8b88;
    font-size: 25px;
    text-align: center;
    margin-bottom: 30px;
  }
`;

export const RegisterText = styled(Typography)`
  && {
    color: #7f8b88;
    font-size: 25px;
    text-align: center;
    margin-bottom: 30px;
    margin-top: 30px;
  }
`;

export const LoginHeader = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  border: 1px solid;
  padding: 50px;
`;

export const LoginContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 50px
`;
