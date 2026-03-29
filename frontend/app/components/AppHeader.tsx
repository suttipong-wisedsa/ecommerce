"use client";

import { Input, Badge } from "antd";
import { ShoppingCartOutlined } from "@ant-design/icons";
import { useState } from "react";
import { useProductStore } from "../store/product";
import { getProducts } from "../service/product";
import { useCartStore } from "../store/cart";

const { Search } = Input;

export default function AppHeader() {
    // const addToProduct = useProductStore((s) => s.addToProduct);
     const cart = useCartStore((s) => s.cartList);
    const addToProduct = useProductStore((s) => s.addToProduct);
    const searchItem = async (e: any) => {
        setTimeout(() => search(e.target.value), 300)
    }

    const search = async (q: string) => {
        let { data } = await getProducts(q)
          addToProduct(data)
    }

    return (
        <div
            style={{
                display: "flex",
                justifyContent: "space-between",
                alignItems: "center",
                maxWidth: 1200,
                margin: "0 auto",
            }}
        >
            <h2>E-Commerce</h2>

            <Search placeholder="Search product..." style={{ width: 300 }} onChange={searchItem} />

            <Badge count={cart}>
                <ShoppingCartOutlined style={{ fontSize: 24 }} />
            </Badge>
        </div>
    );
}