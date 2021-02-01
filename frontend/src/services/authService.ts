import { LoginData, RegisterData } from "../models/login";
import api from "./api";

class AuthService {
  
  static login =  async (login : LoginData) =>  {
    return api.post(`/users/login`, login)
  }

  static register =  async (register : RegisterData) =>  {
    return api.post(`/users`, register)
  }

  static setAuthHeader = (token : String) =>  {
    api.defaults.headers.Authorization = token;
  }

}

export default AuthService;