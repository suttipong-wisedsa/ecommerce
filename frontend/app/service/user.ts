import axios from "axios";

const API = process.env.NEXT_PUBLIC_API_URL;

export const loginUser = async (payload: any) => {
  return await axios.post(`${API}/login`, payload);
};

export const registerUser = async (payload: any) => {
  return await axios.post(`${API}/register`, payload);
};