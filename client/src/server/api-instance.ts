import axios, { AxiosResponse } from "axios";
import { AccessTokenStorage } from "./storage/access-token-storage";

const Api = axios.create({
  baseURL: "http://localhost:8000",
});

Api.interceptors.request.use(
  (config) => {
    config.headers.Authorization = `Bearer ${AccessTokenStorage.get()}`;
    return config;
  },
  (error) => Promise.reject(error)
);

Api.interceptors.response.use(
  (res: AxiosResponse) => {
    return res;
  },
  (err: AxiosResponse) => {
    return err;
  }
);

export default Api;
