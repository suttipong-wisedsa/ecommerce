"use client";

import { Row, Col, message } from "antd";
import { useEffect, useState } from "react";
import ProductCard from "@/app/components/data/ProductCard";
import { getProducts, getProductsCart, postCartProducts } from "./service/product";
import { useCartStore } from "@/app/store/cart";

interface Product {
  id: number;
  name: string;
  price: number;
}

export default function Home() {
  const [products, setProducts] = useState<Product[]>([]);
  const [cart, setCart] = useState<Product[]>([]);
  const addCountCart = useCartStore((s) => s.addToCart);

  const fetchProduct = async () => {
    try {
      let { data } = await getProducts()
      setProducts(data)
    } catch (err) {

    }
  }

  useEffect(() => {
    fetchProduct()
  }, []);

  const addToCart = async (product: Product) => {
    try {
      await postCartProducts({
        "product_id": product.id,
        "quantity": 1
      })
      setCart([...cart, product]);
      let { data } = await getProductsCart()
      addCountCart(data)
      message.success("Added to cart");
    } catch (err) {

    }
  };

  return (
    <div style={{ maxWidth: 1200, margin: "0 auto" }}>
      <Row gutter={[16, 16]}>
        {products.map((p) => (
          <Col xs={24} sm={12} md={8} lg={6} key={p.id}>
            <ProductCard product={p} onAdd={addToCart} />
          </Col>
        ))}
      </Row>
    </div>
  );
}