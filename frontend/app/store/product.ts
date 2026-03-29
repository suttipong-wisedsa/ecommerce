import { create } from "zustand";

type CartItem = {
    product_id: number;
    quantity: number;
};

type ProductStore = {
    productList: any[];
    addToProduct: (item: any[]) => void;
};

export const useProductStore = create<ProductStore>((set) => ({
    productList: [],
    addToProduct: (item) =>
        set((state) => ({
            productList: [...item],
        }))
}));