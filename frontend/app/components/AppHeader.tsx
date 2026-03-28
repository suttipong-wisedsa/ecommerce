"use client";

import { Input, Badge } from "antd";
import { ShoppingCartOutlined } from "@ant-design/icons";
import { useState } from "react";
import { useCartStore } from "../store/cart";

const { Search } = Input;

export default function AppHeader() {
    const cart = useCartStore((s) => s.cartList);

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

            <Search placeholder="Search product..." style={{ width: 300 }} />

            <Badge count={cart}>
                <ShoppingCartOutlined style={{ fontSize: 24 }} />
            </Badge>
        </div>
    );
}