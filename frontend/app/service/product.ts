import axios from "axios";

const API = process.env.NEXT_PUBLIC_API_URL;

export const getProducts = async() => {
  return await axios.get(`${API}/products`);
};

export const postCartProducts = async(payload : any) => {
  return await axios.post(`${API}/products`, payload);
};

export const getProductsCart = async() => {
  return await axios.get(`${API}/products/cart`);
};