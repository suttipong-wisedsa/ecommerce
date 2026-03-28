import { create } from "zustand";

type CartItem = {
    product_id: number;
    quantity: number;
};

type CartStore = {
    cartList: number;
    addToCart: (item: CartItem[]) => void;
};

export const useCartStore = create<CartStore>((set) => ({
    cartList: 0,
    addToCart: (item) =>
        // console.log(item)
        set((state) => ({
            cartList: item.length,
        }))
        // set((state) => {
        //       const exist = state.cartList.find(
        //         (c) => c.product_id === item.product_id
        //       );

        //       if (exist) {
        //         return {
        //           cart: state.cartList.map((c) =>
        //             c.product_id === item.product_id
        //               ? { ...c, quantity: c.quantity + item.quantity }
        //               : c
        //           ),
        //         };
        //       }

        //       return { cart: [...state.cartList, item] };
        // }),
}));