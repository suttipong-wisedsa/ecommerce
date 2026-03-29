import axios from "axios";

const API = process.env.NEXT_PUBLIC_API_URL;
const token = localStorage.getItem("token");

export const getProducts = async (e: string | undefined) => {
  return await axios.get(`${API}/products`, {
    params: { q: e },
  });
};

export const postCartProducts = async (payload: any) => {
  return await axios.post(`${API}/products`, payload, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
};

export const getProductsCart = async () => {
  return await axios.get(`${API}/products/cart`);
};